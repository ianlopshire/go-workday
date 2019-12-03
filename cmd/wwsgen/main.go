package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"golang.org/x/tools/imports"

	"aqwari.net/xml/xsdgen"

	"github.com/ianlopshire/go-workday/internal/wwsgen"
	"github.com/ianlopshire/go-workday/internal/wwsgen/wsdl"
)

var wwsVersion = flag.String("version", "v30.1", "Workday Web Service version")
var service = flag.String("service", "", "WWS service to generate (e.g. Absence_Management)")
var pkg = flag.String("pkg", "gen", "package name to be used int the output")
var output = flag.String("output", "./", "output directory")

func main() {

	flag.Parse()

	if service == nil || *service == "" {
		flag.Usage()
		os.Exit(0)
	}

	wsdlReader, err := getWSDL(*service, *wwsVersion)
	if err != nil {
		log.Fatal(err)
	}
	defer wsdlReader.Close()

	xsdReader, err := getXSD(*service, *wwsVersion)
	if err != nil {
		log.Fatal(err)
	}
	defer xsdReader.Close()

	schema := wsdl.Schema{}

	err = xml.NewDecoder(wsdlReader).Decode(&schema)
	if err != nil {
		log.Fatal(err)
	}

	serv, err := wwsgen.BuildService(schema)
	if err != nil {
		log.Fatal(err)
	}

	serv.Pkg = *pkg
	serv.Name = *service
	serv.Version = *wwsVersion

	clientFile, err := os.Create(path.Join(*output, "wwsgen_client.go"))
	if err != nil {
		log.Fatal(err)
	}
	defer clientFile.Close()

	err = wwsgen.PrintService(serv, clientFile)
	if err != nil {
		log.Fatal(err)
	}

	config := new(xsdgen.Config)
	config.Option(xsdgen.DefaultOptions...)
	config.Option(xsdgen.PackageName(*pkg))
	config.Option(xsdgen.Replace("^24", "N24"))

	xsdData, err := ioutil.ReadAll(xsdReader)
	if err != nil {
		log.Fatal(err)
	}

	code, err := config.GenCode(xsdData)
	if err != nil {
		log.Fatal(err)
	}

	ast, err := code.GenAST()
	if err != nil {
		log.Fatal(err)
	}

	data, err := FormattedSource(ast, path.Join(*output, "fixme.go"))
	if err != nil {
		log.Fatal(err)
	}

	typesFile, err := os.Create(path.Join(*output, "wwsgen_types.go"))
	if err != nil {
		log.Fatal(err)
	}
	defer clientFile.Close()

	_, err = typesFile.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

// FormattedSource converts an abstract syntax tree to
// formatted Go source code.
func FormattedSource(file *ast.File, output string) ([]byte, error) {
	var buf bytes.Buffer

	fileset := token.NewFileSet()

	// our *ast.File did not come from a real Go source
	// file. As such, all of its node positions are 0, and
	// the go/printer package will print the package
	// comment between the package statement and
	// the package name. The most straightforward way
	// to work around this is to put the package comment
	// there ourselves.
	if file.Doc != nil {
		for _, v := range file.Doc.List {
			io.WriteString(&buf, v.Text)
			io.WriteString(&buf, "\n")
		}
		file.Doc = nil
	}

	if err := format.Node(&buf, fileset, file); err != nil {
		return nil, err
	}
	out, err := imports.Process(output, buf.Bytes(), nil)
	if err != nil {
		ioutil.WriteFile(output, buf.Bytes(), os.ModePerm)
		return nil, fmt.Errorf("%v", err)
	}
	return out, nil
}

func getXSD(service, version string) (io.ReadCloser, error) {
	url := "https://community.workday.com/sites/default/files/file-hosting/productionapi/" + service + "/" + version + "/" + service + ".xsd"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func getWSDL(service, version string) (io.ReadCloser, error) {
	url := "https://community.workday.com/sites/default/files/file-hosting/productionapi/" + service + "/" + version + "/" + service + ".wsdl"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

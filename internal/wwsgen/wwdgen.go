package wwsgen

import (
	"encoding/xml"
	"html/template"
	"io"
	"log"
	"strings"

	"github.com/ianlopshire/go-workday/internal/wwsgen/wsdl"
	"github.com/pkg/errors"
)

type Service struct {
	Name          string
	Version       string
	Pkg           string
	Documentation string
	Operations    []Operation
}

type Operation struct {
	Name                  string
	GoName                string
	Documentation         string
	InputElementTypeName  string
	InputXMLName          xml.Name
	OutputElementTypeName string
	OutputXMLName         xml.Name
}

func BuildService(schema wsdl.Schema) (service *Service, err error) {
	service = new(Service)
	service.Documentation = schema.Documentation

	messageMap := map[string]wsdl.Message{}
	for _, message := range schema.Messages {
		messageMap[message.Name] = message
	}

	elementMap := map[string]wsdl.Element{}
	for _, element := range schema.Elements {
		elementMap[element.Name] = element
	}

	for _, operation := range schema.Operations {

		inputMessage, ok := messageMap[strings.TrimPrefix(operation.Input.Message, "wd-wsdl:")]
		if !ok {
			log.Println(errors.Errorf("Cannot find input message for operation %v (%v)", operation.Name, operation.Input.Message))
			continue
		}

		inputElement, ok := elementMap[strings.TrimPrefix(inputMessage.Part.Element, "wd:")]
		if !ok {
			log.Println(errors.Errorf("Cannot find input element for operation %v (%v, %v)", operation.Name, operation.Input.Message, inputMessage.Part.Element))
			continue
		}

		outputMessage, ok := messageMap[strings.TrimPrefix(operation.Output.Message, "wd-wsdl:")]
		if !ok {
			log.Println(errors.Errorf("Cannot find output message for operation %v (%v)", operation.Name, operation.Output.Message))
			continue
		}


		outputElement, ok := elementMap[strings.TrimPrefix(outputMessage.Part.Element, "wd:")]
		if !ok && outputMessage.Part.Element != ""{
			log.Println(errors.Errorf("Cannot find output element for operation %v (%v, %v)", operation.Name, operation.Output.Message, outputMessage.Part.Element))
			continue
		}

		service.Operations = append(service.Operations, Operation{
			Name:                 operation.Name,
			GoName:               operation.Name,
			Documentation:        operation.Documentation,
			InputElementTypeName: strings.TrimPrefix(inputElement.Type, "wd:"),
			InputXMLName: xml.Name{
				Space: "urn:com.workday/bsvc",
				Local: strings.TrimPrefix(inputMessage.Part.Element, "wd:"),
			},
			OutputElementTypeName: strings.TrimPrefix(outputElement.Type, "wd:"),
			OutputXMLName: xml.Name{
				Space: "urn:com.workday/bsvc",
				Local: strings.TrimPrefix(outputMessage.Part.Element, "wd:"),
			},
		})
	}
	return service, nil
}

func PrintService(service *Service, wr io.Writer) error {
	tmpl, err := template.New("template").Funcs(funcMap).Parse(templateText)
	if err != nil {
		return err
	}

	err = tmpl.Execute(wr, service)
	if err != nil {
		return err
	}

	return nil
}

var funcMap = template.FuncMap{
	// The name "title" is what the function will be called in the template text.
	"tick":    tick,
	"goName":  goName,
	"comment": comment,
}

func tick() string {
	return "`"
}

func goName(s string) string {
	r := strings.NewReplacer(
		"_", " ",
		"-", " ",
	)
	s = r.Replace(s)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Title(s)
	return s
}

func comment(s string) string {
	lns := strings.Split(s, "\n")
	return "// " + strings.Join(lns, "\n// ")
}

const templateText = `
// Package {{.Pkg}}
//
{{comment .Documentation}}
package {{.Pkg}}

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "{{.Version}}"
	ServiceName = "{{.Name}}"
)

type Client struct {
	*wws.Client
}

{{range .Operations}}
// {{goName .Name}} ({{.Name}})
// 
{{comment .Documentation}}
func (c *Client) {{goName .Name}}(ctx context.Context, input *{{goName .Name}}Input) (output *{{goName .Name}}Output, err error) {
	output = &{{goName .Name}}Output{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type {{goName .Name}}Input struct {
	XMLName string {{tick}}xml:"{{.InputXMLName.Space}} {{.InputXMLName.Local}}"{{tick}}
	{{goName .InputElementTypeName}}
}

type {{goName .Name}}Output struct {
	XMLName string {{tick}}xml:"{{.OutputXMLName.Space}} {{.OutputXMLName.Local}}"{{tick}}
	{{goName .OutputElementTypeName}}
}
{{end}}
`

package raas_test

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/ianlopshire/go-workday/pkg/raas"
	"github.com/ianlopshire/go-workday/pkg/wd/session"
)

func ExampleClient_CallXML() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	client := raas.NewClient(sess)

	def := raas.Definition{
		Owner: "imlopshire",
		Name:  "example_worker_report",
	}

	// Define the layout of the data in the custom report.
	var layout []struct {
		WID            string `xml:"wid"`
		Name           string `xml:"name"`
		YearsOfService int    `xml:"years_of_service"`
	}

	// Define the values for any report prompts.
	query := url.Values{}
	query.Set("workers", "Employee_ID:1001")

	// Fetch the report data and decode into layout.
	err = client.CallXML(context.Background(), def, query, &layout)
	if err != nil {
		log.Fatal(err)
	}

	// Use the resulting data.
	for i := range layout {
		fmt.Printf("%+v", layout[i])
	}
}

func ExampleClient_CallCSV() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	client := raas.NewClient(sess)

	def := raas.Definition{
		Owner: "imlopshire",
		Name:  "example_worker_report",
	}

	// Define the values for any report prompts.
	query := url.Values{}
	query.Set("workers", "Employee_ID:1001")

	// Fetch the report data and decode into layout.
	reader, err := client.CallCSV(context.Background(), def, query)
	if err != nil {
		log.Fatal(err)
	}

	// Use the resulting data.
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(row)
	}
}

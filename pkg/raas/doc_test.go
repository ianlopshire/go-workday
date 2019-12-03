package raas_test

import (
	"context"
	"log"
	"net/http"
	"net/url"

	"github.com/ianlopshire/go-workday/pkg/raas"
	"github.com/ianlopshire/go-workday/pkg/wd"
)

func ExampleClient() {
	// create a new RaaS client
	session := wd.MustSession(wd.EnvironmentImplementation, "exampleCompany", "username", "password")
	client := raas.NewClient(session, http.DefaultClient)

	// define a report to fetch
	def := raas.Definition{
		Owner: "OwnerID",
		Name:  "Example_Report_Name",
	}

	// define the parameters to use when fetching the report
	params := url.Values{}

	// add a reference parameter
	worker := wd.Reference{"Employee_ID": "999999"}
	params.Add(raas.ReferenceParam("worker", worker))

	// define the layout of the report entry
	var ReportEntries []struct {
		WorkerReference wd.Reference `xml:"Worker_Reference"`
		Name            string       `xml:"Name"`
		EffectiveDate   wd.Date      `xml:"Effective_Date"`
	}

	// fetch the report and store the results in ReportEntries
	err := client.Do(context.Background(), def, params, &ReportEntries)
	if err != nil {
		log.Fatal(err)
	}
}

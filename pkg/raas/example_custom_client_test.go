package raas_test

import (
	"context"
	"fmt"
	"log"

	"github.com/ianlopshire/go-workday/pkg/raas"
	"github.com/ianlopshire/go-workday/pkg/wd/session"
)

// Worker conforms to the layout of the data in the custom report.
type Worker struct {
	EmployeeID string `xml:"Employee_ID"`
	FirstName  string `xml:"First_Name"`
	LastName   string `xml:"Last_Name"`
}

// WorkerClient extends raas.Client with the ability to list workers.
type WorkerClient struct {
	*raas.Client
}

var workerDef = raas.Definition{
	Owner: "Owner",
	Name:  "Workers",
}

// List workers fetches a list of workers via RaaS.
func (c *WorkerClient) ListWorkers(ctx context.Context) (workers []Worker, err error) {
	if err := c.CallXML(ctx, workerDef, nil, &workers); err != nil {
		return nil, err
	}
	return workers, nil
}

// ExampleCustomClient demonstrates how a custom clients can be implemented.
func Example_customClient() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	raasClient := raas.NewClient(sess)

	workerClient := &WorkerClient{
		Client: raasClient,
	}

	// Call the custom implemented ListWorkers method.
	workers, err := workerClient.ListWorkers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Use the resulting data.
	for _, worker := range workers {
		fmt.Printf("%+v", worker)
	}
}

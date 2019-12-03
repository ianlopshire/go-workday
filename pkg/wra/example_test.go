package wra_test

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ianlopshire/go-workday/pkg/wd"
	"github.com/ianlopshire/go-workday/pkg/wra"
)

func ExampleClient() {
	// Create a session that provides configuration to the wra client.
	//
	// The Workday REST API requires OAuth to be used for authentication. This means client
	// credentials must be present on the session.
	wds := &wd.Session{
		Environment: wd.EnvironmentImplementation,
		Tenant:      "Tenant",
		ClientCredentials: wd.ClientCredentials{
			ClientID:     "...",
			ClientSecret: "...",
			RefreshToken: "...",
		},
	}

	// Create a WRA client.
	wraClient := wra.NewClient(wds, http.DefaultClient)

	// In this example we will be fetching a list of organizations.

	// Build the request.
	//
	// The client will automatically construct the full url and add any necessary authorization
	// headers.
	req, err := wraClient.NewRequest("GET", "/organizations", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Do the http request.
	//
	// Do handles checking for and decoding errors.
	res, error := wraClient.Do(req)
	if err != nil {
		log.Fatal(error)
	}
	defer res.Body.Close()

	// Define the layout of the response.
	var response struct {
		Total int `json:"total"`
		Data  []struct {
			Descriptor string `json:"descriptor"`
			ID         string `json:"id"`
		} `json:"data"`
	}

	// Decode the response.
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
}

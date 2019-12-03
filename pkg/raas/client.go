package raas

import (
	"context"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ianlopshire/go-workday/pkg/wd"
	"github.com/pkg/errors"
)

// Definition represents a Workday report definition.
type Definition struct {
	Owner string
	Name  string
}

// Client is a Workday Reports as a Service (RaaS).
//
// A Client is meant to be thread safe. A single instance of client is meant to be passed through an entire application.
type Client struct {
	client  *http.Client
	session *wd.Session
}

// NewClient creates a new Client.
//
// If no http.Client is provided, http.DefaultClient will be used.
func NewClient(session *wd.Session, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	return &Client{
		session: session,
		client:  client,
	}
}

// Do fetches report data via https and stores the result in v.
//
// The provided url values will be added to the url when requesting the report.
//
// The v param will be populated by decoding the xml report data into v. The v param must be a non-nil pointer. If v is
// not a non-nil pointer an error will be returned.
func (c *Client) Do(ctx context.Context, def Definition, params url.Values, v interface{}) error {
	req, err := http.NewRequest(http.MethodGet, c.url(def, params), nil)
	if err != nil {
		return errors.Wrap(err, "could not build raas request")
	}
	req = req.WithContext(ctx)
	req.SetBasicAuth(c.session.Credentials.Username, c.session.Credentials.Password)

	res, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "could not complete raas request")
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		message, _ := ioutil.ReadAll(io.LimitReader(res.Body, 5120)) // limit error responses to 5KB
		return errors.Errorf("raas request responded with %v error (%s)", res.StatusCode, message)
	}

	// This unwraps the report entries
	layout := struct {
		XMLName     xml.Name    `xml:"Report_Data"`
		ReportEntry interface{} `xml:"Report_Entry"`
	}{
		ReportEntry: v,
	}

	err = xml.NewDecoder(res.Body).Decode(&layout)
	if err != nil {
		return errors.Wrap(err, "raas request body could not be decoded")
	}

	return nil
}

// url builds the report url.
func (c *Client) url(def Definition, params url.Values) string {
	var b string
	switch c.session.Environment {
	case wd.EnvironmentProduction:
		b = "https://services1.myworkday.com/ccx/service/customreport2"
	default:
		b = "https://wd2-impl-services1.workday.com/ccx/service/customreport2"
	}

	return b + "/" + c.session.Tenant + "/" + def.Owner + "/" + def.Name + "?" + params.Encode()
}

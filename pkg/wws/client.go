package wws

import (
	"bytes"
	"context"
	"encoding/xml"
	"net/http"
	"time"

	"github.com/ianlopshire/go-workday/pkg/wd"
	"github.com/pkg/errors"
)

// Client is a Workday Web Services (WWS) client.
//
// A Client is meant to be thread safe. A single instance of client is meant to be passed through an entire application.
type Client struct {
	session *wd.Session
	client  *http.Client
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

// Do sends the request and decodes the response.
//
// All SOAP structure and security headers are handled automatically.
//
// The response param is populated by decoding the soap body using the standard xml package. It is important that the
// correct xml struct tags are used. The response param must be a non-nil pointer. If response is not a non-nil pointer
// an error will be returned.
//
// The provided version is used to build the service url but is not added to the request. The request version should be
// part of the provided request.
func (c *Client) Do(ctx context.Context, service, version string, request, response interface{}) error {
	requestEnvelope := newEnvelope(time.Now().UnixNano(), c.session.Credentials.Username+"@"+c.session.Tenant, c.session.Credentials.Password, request)

	buff := new(bytes.Buffer)
	err := xml.NewEncoder(buff).Encode(requestEnvelope)
	if err != nil {
		return errors.Wrap(err, "wws request could not be encoded")
	}

	req, err := http.NewRequest(http.MethodPost, c.url(service, version), buff)
	if err != nil {
		return errors.Wrap(err, "wws request could not be created")
	}
	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	res, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "wws http request failed")
	}
	defer res.Body.Close()

	respEnvelope := new(envelope)
	respEnvelope.Body = body{Content: response}
	err = xml.NewDecoder(res.Body).Decode(respEnvelope)
	if err != nil {
		return errors.Wrap(err, "wws response could not be decoded")
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return errors.Wrap(fault, "wss fault")
	}

	return nil
}

// url builds the service url.
func (c *Client) url(service, version string) string {
	var b string
	switch c.session.Environment {
	case wd.EnvironmentProduction:
		b = "https://services1.myworkday.com/ccx/service"
	default:
		b = "https://wd2-impl-services1.workday.com/ccx/service"
	}

	return b + "/" + c.session.Tenant + "/" + service + "/" + version
}

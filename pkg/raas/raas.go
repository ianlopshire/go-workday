// Package raas facilitates interacting with Workday Reports as a Service (RaaS).
package raas

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/ianlopshire/go-workday/pkg/wd"
	"github.com/ianlopshire/go-workday/pkg/wd/credentials"
	"github.com/ianlopshire/go-workday/pkg/wd/session"
)

const (
	DefaultBaseURLProduction     = "https://services1.myworkday.com/ccx/service/customreport2"
	DefaultBaseURLImplementation = "https://wd2-impl-services1.workday.com/ccx/service/customreport2"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (err *APIError) Error() string {
	return fmt.Sprintf("api responded with a %d error: %s", err.StatusCode, err.Message)
}

// Definition represents a Workday custom report definition.
type Definition struct {
	// Owner is the Workday username of the report owner.
	Owner string

	// Name is the name of the custom report in Workday.
	Name string
}

// Client provides general purpose methods for fetching data via Workday RaaS.
type Client struct {
	Backend Backend
	Sess    *session.Session
}

// NewClient creates a new client.
func NewClient(sess *session.Session) *Client {
	return &Client{
		Backend: &APIBackend{},
		Sess:    sess,
	}
}

// CallRaw fetches report data and returns a buffer of the raw data.
func (c *Client) CallRaw(ctx context.Context, def Definition, query url.Values) (*bytes.Buffer, error) {
	r, err := c.Backend.Call(ctx, c.Sess, def, query)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	buff := new(bytes.Buffer)
	if _, err := io.Copy(buff, r); err != nil {
		return nil, err
	}

	return buff, nil
}

// CallXML fetches report data as XML and decodes it into v. v must be a compatible with
// xml.Unmarshal().
func (c *Client) CallXML(ctx context.Context, def Definition, query url.Values, v interface{}) error {
	r, err := c.Backend.Call(ctx, c.Sess, def, query)
	if err != nil {
		return err
	}
	defer r.Close()

	// Unwrap the report entries.
	overlay := struct {
		XMLName     xml.Name    `xml:"Report_Data"`
		ReportEntry interface{} `xml:"Report_Entry"`
	}{
		ReportEntry: v,
	}

	err = xml.NewDecoder(r).Decode(&overlay)
	if err != nil {
		return err
	}

	return nil
}

// CallCSV fetches report data as a CSV and exposes it via a *csv.Reader.
func (c *Client) CallCSV(ctx context.Context, def Definition, query url.Values) (*csv.Reader, error) {
	if query == nil {
		query = url.Values{}
	}
	// Set the format so the response will be returned as CSV data.
	query.Set("format", "csv")

	b, err := c.Backend.Call(ctx, c.Sess, def, query)
	if err != nil {
		return nil, err
	}
	defer b.Close()

	// Copy the data into a buffer so the readCloser from the backend can be closed.
	buff := new(bytes.Buffer)
	if _, err := io.Copy(buff, b); err != nil {
		return nil, err
	}

	return csv.NewReader(buff), nil
}

// Backend defines the low level interface needed to request Workday RaaS report data.
//
// This interface primarily exists for testing.
type Backend interface {
	Call(ctx context.Context, sess *session.Session, def Definition, query url.Values) (io.ReadCloser, error)
}

// APIBackend is a low level RaaS client. Most users should use Client instead.
//
// This is the main implementation of the Backend interface. The zero value of APIBackend
// is usable.
type APIBackend struct {
	// HTTPClient is the http client that will be used to request reports. If HTTPClient
	// is nil http.DefaultClient will be used.
	HTTPClient *http.Client

	// BaseURL, if set, will override the base URL used to request reports. This is useful
	// for testing.
	BaseURL string
}

// Call fetches report data via the Workday RaaS HTTP endpoint. If the endpoint response
// contains an error an *APIError will be returned.
//
// Credentials for authentication are taken from the provided session. Currently only
// basic auth credentials are supported.
//
// Call does not close the http response body.
func (b *APIBackend) Call(ctx context.Context, sess *session.Session, def Definition, query url.Values) (io.ReadCloser, error) {
	url, err := b.BuildURL(sess, def, query)
	if err != nil {
		return nil, err
	}

	creds, err := b.creds(sess)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.SetBasicAuth(creds.Username, creds.Password)

	// Fall back to http.DefaultClient if b.HTTPClient is not set.
	client := b.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	res, err := b.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		defer res.Body.Close()
		message, _ := ioutil.ReadAll(io.LimitReader(res.Body, 5120)) // limit error responses to 5KB
		return nil, &APIError{
			StatusCode: res.StatusCode,
			Message:    string(message),
		}
	}

	return res.Body, nil
}

func (b *APIBackend) creds(sess *session.Session) (*credentials.BasicAuthCredentials, error) {
	if sess.Credentials == nil {
		return nil, errors.New("missing credentials")
	}

	creds, ok := sess.Credentials.CredentialOfType(credentials.TypeBasicAuth)
	if !ok {
		return nil, errors.New("missing credentials: only basic auth credentials are currently supported")
	}
	return creds.(*credentials.BasicAuthCredentials), nil
}

// BuildURL builds the URL for a report.
//
// The URL format is `{baseURL}/{tenant}/{ownerUsername}/{reportName}`. If a base URL
// is not provided it will automatically determined based on the provided session.
func (b *APIBackend) BuildURL(sess *session.Session, def Definition, query url.Values) (string, error) {
	base := b.BaseURL
	if base == "" {
		switch sess.Environment {
		case wd.EnvironmentProduction:
			base = DefaultBaseURLProduction
		case wd.EnvironmentImplementation:
			base = DefaultBaseURLImplementation
		default:
			return "", fmt.Errorf("cannot determine base url for environment %q: base url must be explicitly set for non-standard environments", sess.Environment)
		}
	}

	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, sess.Tenant, def.Owner, def.Name)
	u.RawQuery = query.Encode()
	return u.String(), nil
}

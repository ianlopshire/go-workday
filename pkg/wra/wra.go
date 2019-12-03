// Package wra facilitates interacting with the wra (Workday REST API).
package wra

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ianlopshire/go-workday/internal/oauth"
	"github.com/ianlopshire/go-workday/pkg/wd"
)

// Client is a Workday REST API client.
type Client struct {
	session     *wd.Session
	tokenSource *oauth.TokenSource
	client      *http.Client
}

// New Client creates a new Client.
//
// If the provided http.Client is nil, http.Default client will be used.
func NewClient(session *wd.Session, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	return &Client{
		session: session,
		tokenSource: oauth.NewTokenSource(
			session.ClientCredentials.ClientID,
			session.ClientCredentials.ClientSecret,
			session.ClientCredentials.RefreshToken,
			tokenURL(session),
			client,
		),
		client: client,
	}
}

// NewRequest creates a new wra specific http request.
func (c *Client) NewRequest(method, uri string, body io.Reader) (*http.Request, error) {
	url := apiURL(c.session) + uri

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	return req, nil
}

// Do does the provided http request and checks the response for errors.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		err := ResponseError{
			StatusCode: res.StatusCode,
		}

		_ = json.NewDecoder(res.Body).Decode(&err)
		_ = res.Body.Close()

		return nil, err
	}

	return res, nil
}

func tokenURL(session *wd.Session) string {
	if session.Environment == wd.EnvironmentProduction {
		return "https://services1.myworkday.com/ccx/oauth2/" + session.Tenant + "/token"
	}

	if session.Environment == wd.EnvironmentImplementation {
		return "https://wd2-impl-services1.workday.com/ccx/oauth2/" + session.Tenant + "/token"
	}

	return ""
}

func apiURL(session *wd.Session) string {
	if session.Environment == wd.EnvironmentProduction {
		return "https://services1.myworkday.com/ccx/api/v1/" + session.Tenant
	}

	if session.Environment == wd.EnvironmentImplementation {
		return "https://wd2-impl-services1.workday.com/ccx/api/v1/" + session.Tenant
	}

	return ""
}

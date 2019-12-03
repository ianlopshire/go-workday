// Package credentials provides credential retrieval and management
package credentials

import (
	"os"
)

// Type represents a type of supported credentials.
type Type string

const (
	// CredentialTypeBasicAuth represents basic auth credentials.
	TypeBasicAuth Type = "basic_auth"

	// CredentialTypeOAuthClientCredentials represents OAuth 2.0 client credentials.
	TypeOAuthClientCredentials Type = "client_credentials"
)

// FromEnvironment loads credentials from the environment.
//
//	# Basic Auth
//	WD_BASIC_AUTH_USERNAME
//	WD_BASIC_AUTH_PASSWORD
//
//	# OAuth Client Credentials
//	WD_CLIENT_CREDENTIALS_CLIENT_ID
//	WD_CLIENT_CREDENTIALS_CLIENT_SECRET
//	WD_CLIENT_CREDENTIALS_REFRESH_TOKEN
func FromEnvironment() Providers {
	var creds Providers
	{
		cred := basicAuthCredentialsFromEnvironment()
		if cred != nil {
			creds = append(creds, cred)
		}
	}
	{
		cred := oAuthClientCredentialsFromEnvironment()
		if cred != nil {
			creds = append(creds, cred)
		}
	}

	if len(creds) > 0 {
		return creds
	}
	return nil
}

// Provider defines the interface for a credential provider.
type Provider interface {
	CredentialOfType(t Type) (cred interface{}, ok bool)
}

// Providers allows []Provider to implement the Provider interface.
type Providers []Provider

func (p Providers) CredentialOfType(t Type) (cred interface{}, ok bool) {
	for i := range p {
		if c, ok := p[i].CredentialOfType(t); ok {
			return c, true
		}
	}
	return nil, false
}

// BasicAuthCredentials provides credentials for basic auth.
type BasicAuthCredentials struct {
	Username string
	Password string
}

func (c *BasicAuthCredentials) CredentialOfType(t Type) (cred interface{}, ok bool) {
	if t == TypeBasicAuth {
		return c, true
	}
	return nil, false
}

func basicAuthCredentialsFromEnvironment() *BasicAuthCredentials {
	cred := BasicAuthCredentials{
		Username: os.Getenv("WD_BASIC_AUTH_USERNAME"),
		Password: os.Getenv("WD_BASIC_AUTH_PASSWORD"),
	}
	if cred.Username == "" {
		return nil
	}
	return &cred
}

// OAuthClientCredentials provides credentials for the OAuth 2.0 client credentials flow.
type OAuthClientCredentials struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
}

func (c *OAuthClientCredentials) CredentialOfType(t Type) (cred interface{}, ok bool) {
	if t == TypeOAuthClientCredentials {
		return c, true
	}
	return nil, false
}

func oAuthClientCredentialsFromEnvironment() *OAuthClientCredentials {
	cred := OAuthClientCredentials{
		ClientID:     os.Getenv("WD_CLIENT_CREDENTIALS_CLIENT_ID"),
		ClientSecret: os.Getenv("WD_CLIENT_CREDENTIALS_CLIENT_SECRET"),
		RefreshToken: os.Getenv("WD_CLIENT_CREDENTIALS_REFRESH_TOKEN"),
	}
	if cred.ClientID == "" {
		return nil
	}
	return &cred
}

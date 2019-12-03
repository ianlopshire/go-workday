package wd

import (
	"errors"
	"fmt"
)

const (
	// EnvironmentProduction represents the Workday production environment.
	EnvironmentProduction = "Production"

	// EnvironmentImplementation represents the Workday implementation environment.
	EnvironmentImplementation = "Implementation"
)

// Session provides configuration for Workday service clients.
//
// A single session is meant to be passed to multiple Workday clients across multiple threads. It is not safe to mutate
// a session after it has been passed to a client.
//
// NewSession is the recommended way to create a new session.
type Session struct {
	Environment       string
	Tenant            string
	Credentials       Credentials
	ClientCredentials ClientCredentials
}

// Credentials are Workday login credentials
type Credentials struct {
	Username string
	Password string
}

// ClientCredentials represents Workday OAuth 2 client credentials.
type ClientCredentials struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
}

// NewSession creates a new session. If the provided parameters are not valid an error is returned.
func NewSession(environment, tenant, username, password string) (*Session, error) {
	if environment != EnvironmentProduction && environment != EnvironmentImplementation {
		return nil, fmt.Errorf("unknown environment %v", environment)
	}

	if tenant == "" {
		return nil, errors.New("tenant cannot be empty")
	}

	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	return &Session{
		Environment: environment,
		Tenant:      tenant,
		Credentials: Credentials{
			Username: username,
			Password: password,
		},
	}, nil
}

// MustSession creates a new session. If the provided parameters are not valid MustSession panics.
func MustSession(environment, tenant, username, password string) *Session {
	session, err := NewSession(environment, tenant, username, password)
	if err != nil {
		panic(err)
	}
	return session
}

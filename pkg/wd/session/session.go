// Package session provides configuration for Workday service clients.
package session

import (
	"errors"
	"os"

	"github.com/ianlopshire/go-workday/pkg/wd"
	"github.com/ianlopshire/go-workday/pkg/wd/credentials"
)

// Session provides configuration for Workday service clients.
//
// A single session is meant to be shared by multiple Workday clients and threads. It is
// not safe to mutate a session after it has been passed to a client.
//
// NewSession is the recommended way to create a new session.
type Session struct {
	Environment string
	Tenant      string
	Credentials credentials.Provider
}

// NewSession creates a new session. If no options are provided the session will created
// from the environment. See the FromEnvironment option for details.
func NewSession(options ...Option) (sess *Session, err error) {
	if len(options) <= 0 {
		options = []Option{FromEnvironment}
	}

	for i := range options {
		sess, err = options[i](sess)
		if err != nil {
			return nil, err
		}
	}

	if err := sess.Validate(); err != nil {
		return nil, err
	}

	return sess, nil
}

// MustSession is like NewSession but panics in the cased where NewSession would return an
// error.
func MustSession(options ...Option) (sess *Session) {
	sess, err := NewSession(options...)
	if err != nil {
		panic(err)
	}
	return sess
}

// Validate returns an error if the session is invalid.
func (sess *Session) Validate() error {
	if sess.Environment == "" {
		return errors.New("environment cannot be empty")
	}
	if sess.Tenant == "" {
		return errors.New("tenant cannot be empty")
	}
	return nil
}

// Option defines the function signature used for session options.
type Option func(sess *Session) (*Session, error)

// FromEnvironment is a session option that loads session data from the environment.
//
//	WD_ENVIRONMENT
//	WD_TENANT
//
//	# Basic Auth
//	WD_BASIC_AUTH_USERNAME
//	WD_BASIC_AUTH_PASSWORD
//
//	# OAuth Client Credentials
//	WD_CLIENT_CREDENTIALS_CLIENT_ID
//	WD_CLIENT_CREDENTIALS_CLIENT_SECRET
//	WD_CLIENT_CREDENTIALS_REFRESH_TOKEN
func FromEnvironment(sess *Session) (*Session, error) {
	var config Config
	if tenant := os.Getenv("WD_TENANT"); tenant != "" {
		config.Tenant = wd.String(tenant)
	}
	if environment := os.Getenv("WD_ENVIRONMENT"); environment != "" {
		config.Environment = wd.String(environment)
	}
	if creds := credentials.FromEnvironment(); creds != nil {
		config.Credentials = creds
	}
	return FromConfig(config)(sess)
}

// FromConfig is a session option that loads session from a manual configuration object.
func FromConfig(config Config) Option {
	return func(sess *Session) (*Session, error) {
		if sess == nil {
			sess = new(Session)
		}
		if config.Environment != nil {
			sess.Environment = wd.StringValue(config.Environment)
		}
		if config.Tenant != nil {
			sess.Tenant = wd.StringValue(config.Tenant)
		}
		if config.Credentials != nil {
			sess.Credentials = config.Credentials
		}
		return sess, nil
	}
}

// Config is used to manually configure a session via the FromConfig session option.
type Config struct {
	Environment *string
	Tenant      *string
	Credentials credentials.Provider
}

package session_test

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/ianlopshire/go-workday/pkg/wd"
	"github.com/ianlopshire/go-workday/pkg/wd/credentials"
	"github.com/ianlopshire/go-workday/pkg/wd/session"
)

func TestNewSession(t *testing.T) {
	type args struct {
		options []session.Option
	}
	tests := []struct {
		name     string
		args     args
		env      map[string]string
		wantSess *session.Session
		wantErr  bool
	}{
		{
			name: "default to from env",
			args: args{
				options: nil,
			},
			env: map[string]string{
				"WD_ENVIRONMENT": "environment",
				"WD_TENANT":      "tenant",
			},
			wantSess: &session.Session{
				Environment: "environment",
				Tenant:      "tenant",
				Credentials: nil,
			},
			wantErr: false,
		},
		{
			name: "error from option",
			args: args{
				options: []session.Option{
					session.FromConfig(session.Config{
						Environment: wd.String("environment"),
						Tenant:      wd.String("tenant"),
					}),
					func(sess *session.Session) (s *session.Session, err error) {
						return nil, errors.New("error")
					},
				},
			},
			env:      nil,
			wantSess: nil,
			wantErr:  true,
		},
		{
			name: "invalid",
			args: args{
				options: nil,
			},
			env: map[string]string{
				"WD_ENVIRONMENT": "",
				"WD_TENANT":      "",
			},
			wantSess: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			for k, v := range tt.env {
				os.Setenv(k, v)
			}
			gotSess, err := session.NewSession(tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSess, tt.wantSess) {
				t.Errorf("NewSession() gotSess = %v, want %v", gotSess, tt.wantSess)
			}
		})
	}
}

func TestSession_Validate(t *testing.T) {
	type fields struct {
		Environment string
		Tenant      string
		Credentials credentials.Provider
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "missing environment",
			fields: fields{
				Environment: "",
				Tenant:      "tenant",
				Credentials: nil,
			},
			wantErr: true,
		},
		{
			name: "missing tenant",
			fields: fields{
				Environment: "environment",
				Tenant:      "",
				Credentials: nil,
			},
			wantErr: true,
		},
		{
			name: "valid",
			fields: fields{
				Environment: "environment",
				Tenant:      "tenant",
				Credentials: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sess := &session.Session{
				Environment: tt.fields.Environment,
				Tenant:      tt.fields.Tenant,
				Credentials: tt.fields.Credentials,
			}
			if err := sess.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFromConfig(t *testing.T) {
	type args struct {
		config session.Config
		sess   *session.Session
	}
	tests := []struct {
		name    string
		args    args
		want    *session.Session
		wantErr bool
	}{
		{
			name: "populate all",
			args: args{
				config: session.Config{
					Environment: wd.String("environment"),
					Tenant:      wd.String("tenant"),
					Credentials: credentials.Providers{
						&credentials.BasicAuthCredentials{
							Username: "username",
							Password: "password",
						},
					},
				},
				sess: nil,
			},
			want: &session.Session{
				Environment: "environment",
				Tenant:      "tenant",
				Credentials: credentials.Providers{
					&credentials.BasicAuthCredentials{
						Username: "username",
						Password: "password",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "populate none",
			args: args{
				config: session.Config{},
				sess:   nil,
			},
			want:    &session.Session{},
			wantErr: false,
		},
		{
			name: "populate all",
			args: args{
				config: session.Config{
					Environment: wd.String("environment2"),
					Tenant:      wd.String("tenant2"),
					Credentials: credentials.Providers{
						&credentials.BasicAuthCredentials{
							Username: "username2",
							Password: "password2",
						},
					},
				},
				sess: &session.Session{
					Environment: "environment",
					Tenant:      "tenant",
					Credentials: credentials.Providers{
						&credentials.BasicAuthCredentials{
							Username: "username",
							Password: "password",
						},
					},
				},
			},
			want: &session.Session{
				Environment: "environment2",
				Tenant:      "tenant2",
				Credentials: credentials.Providers{
					&credentials.BasicAuthCredentials{
						Username: "username2",
						Password: "password2",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "override none",
			args: args{
				config: session.Config{},
				sess: &session.Session{
					Environment: "environment",
					Tenant:      "tenant",
					Credentials: credentials.Providers{
						&credentials.BasicAuthCredentials{
							Username: "username",
							Password: "password",
						},
					},
				},
			},
			want: &session.Session{
				Environment: "environment",
				Tenant:      "tenant",
				Credentials: credentials.Providers{
					&credentials.BasicAuthCredentials{
						Username: "username",
						Password: "password",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := session.FromConfig(tt.args.config)(tt.args.sess)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromConfig(config)() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromConfig(config)() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromEnvironment(t *testing.T) {
	type args struct {
		sess *session.Session
	}
	tests := []struct {
		name    string
		args    args
		env     map[string]string
		want    *session.Session
		wantErr bool
	}{
		{
			name: "populate all",
			args: args{
				sess: nil,
			},
			env: map[string]string{
				"WD_BASIC_AUTH_USERNAME": "username",
				"WD_BASIC_AUTH_PASSWORD": "password",
				"WD_ENVIRONMENT":         "environment",
				"WD_TENANT":              "tenant",
			},
			want: &session.Session{
				Environment: "environment",
				Tenant:      "tenant",
				Credentials: credentials.Providers{
					&credentials.BasicAuthCredentials{
						Username: "username",
						Password: "password",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "populate none",
			args: args{
				sess: nil,
			},
			env:     map[string]string{},
			want:    &session.Session{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			for k, v := range tt.env {
				os.Setenv(k, v)
			}
			got, err := session.FromEnvironment(tt.args.sess)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

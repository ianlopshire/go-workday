package credentials_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/ianlopshire/go-workday/pkg/wd/credentials"
)

func TestFromEnvironment(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want credentials.Providers
	}{
		{
			name: "basic auth only",
			env: map[string]string{
				"WD_BASIC_AUTH_USERNAME": "username",
				"WD_BASIC_AUTH_PASSWORD": "password",
			},
			want: credentials.Providers{&credentials.BasicAuthCredentials{
				Username: "username",
				Password: "password",
			}},
		},
		{
			name: "client credentials only",
			env: map[string]string{
				"WD_CLIENT_CREDENTIALS_CLIENT_ID":     "client_id",
				"WD_CLIENT_CREDENTIALS_CLIENT_SECRET": "client_secret",
				"WD_CLIENT_CREDENTIALS_REFRESH_TOKEN": "refresh_token",
			},
			want: credentials.Providers{&credentials.OAuthClientCredentials{
				ClientID:     "client_id",
				ClientSecret: "client_secret",
				RefreshToken: "refresh_token",
			}},
		},
		{
			name: "no creds in env",
			env:  map[string]string{},
			want: nil,
		},
		{
			name: "multiple creds in env",
			env: map[string]string{
				"WD_BASIC_AUTH_USERNAME":              "username",
				"WD_BASIC_AUTH_PASSWORD":              "password",
				"WD_CLIENT_CREDENTIALS_CLIENT_ID":     "client_id",
				"WD_CLIENT_CREDENTIALS_CLIENT_SECRET": "client_secret",
				"WD_CLIENT_CREDENTIALS_REFRESH_TOKEN": "refresh_token",
			},
			want: credentials.Providers{
				&credentials.BasicAuthCredentials{
					Username: "username",
					Password: "password",
				},
				&credentials.OAuthClientCredentials{
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					RefreshToken: "refresh_token",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			for k, v := range tt.env {
				os.Setenv(k, v)
			}
			if got := credentials.FromEnvironment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProviders_CredentialOfType(t *testing.T) {
	type args struct {
		t credentials.Type
	}
	tests := []struct {
		name     string
		p        credentials.Providers
		args     args
		wantCred interface{}
		wantOk   bool
	}{
		{
			name: "ok",
			p: credentials.Providers{
				&credentials.OAuthClientCredentials{
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					RefreshToken: "refresh_token",
				},
				&credentials.BasicAuthCredentials{
					Username: "username",
					Password: "password",
				},
			},
			args: args{
				t: credentials.TypeBasicAuth,
			},
			wantCred: &credentials.BasicAuthCredentials{
				Username: "username",
				Password: "password",
			},
			wantOk: true,
		},
		{
			name: "not ok",
			p: credentials.Providers{
				&credentials.OAuthClientCredentials{
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					RefreshToken: "refresh_token",
				},
			},
			args: args{
				t: credentials.TypeBasicAuth,
			},
			wantCred: nil,
			wantOk:   false,
		},
		{
			name: "nil p",
			p:    nil,
			args: args{
				t: credentials.TypeBasicAuth,
			},
			wantCred: nil,
			wantOk:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCred, gotOk := tt.p.CredentialOfType(tt.args.t)
			if !reflect.DeepEqual(gotCred, tt.wantCred) {
				t.Errorf("CredentialOfType() gotCred = %v, want %v", gotCred, tt.wantCred)
			}
			if gotOk != tt.wantOk {
				t.Errorf("CredentialOfType() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestBasicAuthCredentials_CredentialOfType(t *testing.T) {
	type fields struct {
		Username string
		Password string
	}
	type args struct {
		t credentials.Type
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCred interface{}
		wantOk   bool
	}{
		{
			name: "ok",
			fields: fields{
				Username: "username",
				Password: "password",
			},
			args: args{
				t: credentials.TypeBasicAuth,
			},
			wantCred: &credentials.BasicAuthCredentials{
				Username: "username",
				Password: "password",
			},
			wantOk: true,
		},
		{
			name: "not ok",
			fields: fields{
				Username: "username",
				Password: "password",
			},
			args: args{
				t: credentials.TypeOAuthClientCredentials,
			},
			wantCred: nil,
			wantOk:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &credentials.BasicAuthCredentials{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			gotCred, gotOk := c.CredentialOfType(tt.args.t)
			if !reflect.DeepEqual(gotCred, tt.wantCred) {
				t.Errorf("CredentialOfType() gotCred = %v, want %v", gotCred, tt.wantCred)
			}
			if gotOk != tt.wantOk {
				t.Errorf("CredentialOfType() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestOAuthClientCredentials_CredentialOfType(t *testing.T) {
	type fields struct {
		ClientID     string
		ClientSecret string
		RefreshToken string
	}
	type args struct {
		t credentials.Type
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCred interface{}
		wantOk   bool
	}{
		{
			name: "ok",
			fields: fields{
				ClientID:     "client_id",
				ClientSecret: "client_secret",
				RefreshToken: "refresh_token",
			},
			args: args{
				t: credentials.TypeOAuthClientCredentials,
			},
			wantCred: &credentials.OAuthClientCredentials{
				ClientID:     "client_id",
				ClientSecret: "client_secret",
				RefreshToken: "refresh_token",
			},
			wantOk: true,
		},
		{
			name: "not ok",
			fields: fields{
				ClientID:     "client_id",
				ClientSecret: "client_secret",
				RefreshToken: "refresh_token",
			},
			args: args{
				t: credentials.TypeBasicAuth,
			},
			wantCred: nil,
			wantOk:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &credentials.OAuthClientCredentials{
				ClientID:     tt.fields.ClientID,
				ClientSecret: tt.fields.ClientSecret,
				RefreshToken: tt.fields.RefreshToken,
			}
			gotCred, gotOk := c.CredentialOfType(tt.args.t)
			if !reflect.DeepEqual(gotCred, tt.wantCred) {
				t.Errorf("CredentialOfType() gotCred = %v, want %v", gotCred, tt.wantCred)
			}
			if gotOk != tt.wantOk {
				t.Errorf("CredentialOfType() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

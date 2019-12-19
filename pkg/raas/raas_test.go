package raas_test

import (
	"bytes"
	"context"
	"errors"
	"github.com/ianlopshire/go-workday/pkg/raas/mocks"
	"github.com/ianlopshire/go-workday/pkg/wd"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/ianlopshire/go-workday/pkg/raas"
	"github.com/ianlopshire/go-workday/pkg/wd/credentials"
	"github.com/ianlopshire/go-workday/pkg/wd/session"
)

func TestAPIBackend_Call(t *testing.T) {
	successHandler := func(body []byte) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Write(body)
		}
	}
	type args struct {
		ctx   context.Context
		sess  *session.Session
		def   raas.Definition
		query url.Values
	}
	tests := []struct {
		name    string
		args    args
		handler http.HandlerFunc
		want    []byte
		wantErr bool
	}{
		{
			name: "endpoint success",
			args: args{
				sess: &session.Session{
					Environment: wd.EnvironmentImplementation,
					Tenant:      "tenant",
					Credentials: &credentials.BasicAuthCredentials{
						Username: "username",
						Password: "password",
					},
				},
				def: raas.Definition{
					Owner: "username",
					Name:  "report",
				},
			},
			handler: successHandler([]byte(`...`)),
			want:    []byte(`...`),
			wantErr: false,
		},
		{
			name: "endpoint error",
			args: args{
				sess: &session.Session{
					Environment: wd.EnvironmentImplementation,
					Tenant:      "tenant",
					Credentials: &credentials.BasicAuthCredentials{
						Username: "username",
						Password: "password",
					},
				},
				def: raas.Definition{
					Owner: "username",
					Name:  "report",
				},
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
				w.Write([]byte("Internal Server Error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing creds",
			args: args{
				sess: &session.Session{
					Environment: wd.EnvironmentImplementation,
					Tenant:      "tenant",
					Credentials: nil,
				},
				def: raas.Definition{
					Owner: "username",
					Name:  "report",
				},
			},
			handler: successHandler([]byte(`...`)),
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsupported credentials",
			args: args{
				sess: &session.Session{
					Environment: wd.EnvironmentImplementation,
					Tenant:      "tenant",
					Credentials: &credentials.OAuthClientCredentials{
						ClientID:     "clientID",
						ClientSecret: "clientSecret",
						RefreshToken: "refreshToken",
					},
				},
				def: raas.Definition{
					Owner: "username",
					Name:  "report",
				},
			},
			handler: successHandler([]byte(`...`)),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(tt.handler)
			defer ts.Close()
			b := &raas.APIBackend{
				HTTPClient: ts.Client(),
				BaseURL:    ts.URL,
			}

			ctx := tt.args.ctx
			if ctx == nil {
				ctx = context.Background()
			}

			r, err := b.Call(ctx, tt.args.sess, tt.args.def, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				got, _ := ioutil.ReadAll(r)
				if !bytes.Equal(got, tt.want) {
					t.Errorf("Call() got = %s, want %s", got, tt.want)
				}
			}
		})
	}
}

func TestAPIBackend_BuildURL(t *testing.T) {
	type fields struct {
		BaseURL string
	}
	type args struct {
		sess  *session.Session
		def   raas.Definition
		query url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "base url specified",
			fields: fields{
				BaseURL: "https://example.com",
			},
			args: args{
				sess: &session.Session{
					Environment: wd.EnvironmentImplementation,
					Tenant:      "tenant",
				},
				def: raas.Definition{
					Owner: "username",
					Name:  "report",
				},
				query: url.Values{"format": []string{"json"}},
			},
			want:    "https://example.com/tenant/username/report?format=json",
			wantErr: false,
		},
		{
			name: "nil query",
			fields: fields{
				BaseURL: "https://example.com",
			},
			args: args{
				sess: &session.Session{
					Environment: wd.EnvironmentImplementation,
					Tenant:      "tenant",
				},
				def: raas.Definition{
					Owner: "username",
					Name:  "report",
				},
				query: nil,
			},
			want:    "https://example.com/tenant/username/report",
			wantErr: false,
		},
		{
			name: "default implementation base url",
			fields: fields{
				BaseURL: "",
			},
			args: args{
				sess: &session.Session{
					Environment: wd.EnvironmentImplementation,
					Tenant:      "tenant",
				},
				def: raas.Definition{
					Owner: "username",
					Name:  "report",
				},
				query: url.Values{"format": []string{"json"}},
			},
			want:    "https://wd2-impl-services1.workday.com/ccx/service/customreport2/tenant/username/report?format=json",
			wantErr: false,
		},
		{
			name: "default production base url",
			fields: fields{
				BaseURL: "",
			},
			args: args{
				sess: &session.Session{
					Environment: wd.EnvironmentProduction,
					Tenant:      "tenant",
				},
				def: raas.Definition{
					Owner: "username",
					Name:  "report",
				},
				query: url.Values{"format": []string{"json"}},
			},
			want:    "https://services1.myworkday.com/ccx/service/customreport2/tenant/username/report?format=json",
			wantErr: false,
		},
		{
			name: "invalid base url error",
			fields: fields{
				BaseURL: "%gh&%ij",
			},
			args: args{
				sess: &session.Session{
					Environment: wd.EnvironmentProduction,
					Tenant:      "tenant",
				},
				def: raas.Definition{
					Owner: "username",
					Name:  "report",
				},
				query: url.Values{"format": []string{"json"}},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unknown environment error",
			fields: fields{
				BaseURL: "",
			},
			args: args{
				sess: &session.Session{
					Environment: "not-a-real-environment",
					Tenant:      "tenant",
				},
				def: raas.Definition{
					Owner: "username",
					Name:  "report",
				},
				query: url.Values{"format": []string{"json"}},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &raas.APIBackend{
				BaseURL: tt.fields.BaseURL,
			}
			got, err := b.BuildURL(tt.args.sess, tt.args.def, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIError_Error(t *testing.T) {
	type fields struct {
		StatusCode int
		Message    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base case",
			fields: fields{
				StatusCode: 500,
				Message:    "Internal Server Error",
			},
			want: "api responded with a 500 error: Internal Server Error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &raas.APIError{
				StatusCode: tt.fields.StatusCode,
				Message:    tt.fields.Message,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CallCSV(t *testing.T) {
	type fields struct {
		Sess  *session.Session
		Body  []byte
		Error error
	}
	type args struct {
		ctx   context.Context
		def   raas.Definition
		query url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "base case",
			fields: fields{
				Sess: &session.Session{
					Environment: wd.EnvironmentProduction,
					Tenant:      "Tenant",
					Credentials: &credentials.BasicAuthCredentials{
						Username: "Username",
						Password: "Password",
					},
				},
				Body:  []byte("column1,column2,column3\nvalue1,value2,value3"),
				Error: nil,
			},
			args: args{
				ctx: context.Background(),
				def: raas.Definition{
					Owner: "Owner",
					Name:  "Name",
				},
				query: url.Values{"worker": []string{"Employee_ID!9999"}},
			},
			want: [][]string{
				{"column1", "column2", "column3"},
				{"value1", "value2", "value3"},
			},
			wantErr: false,
		},
		{
			name: "backend error",
			fields: fields{
				Sess: &session.Session{
					Environment: wd.EnvironmentProduction,
					Tenant:      "Tenant",
					Credentials: &credentials.BasicAuthCredentials{
						Username: "Username",
						Password: "Password",
					},
				},
				Body:  nil,
				Error: errors.New("error"),
			},
			args: args{
				ctx: context.Background(),
				def: raas.Definition{
					Owner: "Owner",
					Name:  "Name",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backend := new(mocks.Backend)
			query := url.Values{}
			query.Set("format", "csv")
			for k, v := range tt.args.query {
				query[k] = v
			}
			backend.On("Call", tt.args.ctx, tt.fields.Sess, tt.args.def, query).
				Return(ioutil.NopCloser(bytes.NewReader(tt.fields.Body)), tt.fields.Error)
			c := &raas.Client{
				Backend: backend,
				Sess:    tt.fields.Sess,
			}
			r, err := c.CallCSV(tt.args.ctx, tt.args.def, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("CallCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got, err := r.ReadAll()
				if (err != nil) != tt.wantErr {
					t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("CallCSV() got = %v, want %v", got, tt.want)
				}
			}
			backend.AssertExpectations(t)
		})
	}
}

func TestClient_CallRaw(t *testing.T) {
	type fields struct {
		Error error
		Sess  *session.Session
	}
	type args struct {
		ctx   context.Context
		def   raas.Definition
		query url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "base case",
			fields: fields{
				Sess: &session.Session{
					Environment: wd.EnvironmentProduction,
					Tenant:      "Tenant",
					Credentials: &credentials.BasicAuthCredentials{
						Username: "Username",
						Password: "Password",
					},
				},
				Error: nil,
			},
			args: args{
				ctx: context.Background(),
				def: raas.Definition{
					Owner: "Owner",
					Name:  "Name",
				},
				query: url.Values{"worker": []string{"Employee_ID!9999"}},
			},
			want:    []byte(`...`),
			wantErr: false,
		},
		{
			name: "backend error",
			fields: fields{
				Sess: &session.Session{
					Environment: wd.EnvironmentProduction,
					Tenant:      "Tenant",
					Credentials: &credentials.BasicAuthCredentials{
						Username: "Username",
						Password: "Password",
					},
				},
				Error: errors.New("error"),
			},
			args: args{
				ctx: context.Background(),
				def: raas.Definition{
					Owner: "Owner",
					Name:  "Name",
				},
				query: nil,
			},
			want:    []byte(`...`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backend := new(mocks.Backend)
			backend.On("Call", tt.args.ctx, tt.fields.Sess, tt.args.def, tt.args.query).
				Return(ioutil.NopCloser(bytes.NewReader(tt.want)), tt.fields.Error)
			c := &raas.Client{
				Backend: backend,
				Sess:    tt.fields.Sess,
			}
			got, err := c.CallRaw(tt.args.ctx, tt.args.def, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("CallRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got.Bytes(), tt.want) {
				t.Errorf("CallRaw() got = %v, want %v", got, tt.want)
			}
			backend.AssertExpectations(t)
		})
	}
}

func TestClient_CallXML(t *testing.T) {
	type testData struct {
		S1 string `xml:"S1"`
		I1 int    `xml:"I1"`
	}
	type fields struct {
		Error error
		Body  []byte
		Sess  *session.Session
	}
	type args struct {
		ctx   context.Context
		def   raas.Definition
		query url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []testData
		wantErr bool
	}{
		{
			name: "base case",
			fields: fields{
				Sess: &session.Session{
					Environment: wd.EnvironmentProduction,
					Tenant:      "Tenant",
					Credentials: &credentials.BasicAuthCredentials{
						Username: "Username",
						Password: "Password",
					},
				},
				Body: []byte(`
					<Report_Data>
						<Report_Entry>
							<S1>string 1</S1>
							<I1>1</I1>
						</Report_Entry>
					</Report_Data>`,
				),
				Error: nil,
			},
			args: args{
				ctx: context.Background(),
				def: raas.Definition{
					Owner: "Owner",
					Name:  "Name",
				},
				query: url.Values{"worker": []string{"Employee_ID!9999"}},
			},
			want: []testData{{
				S1: "string 1",
				I1: 1,
			}},
			wantErr: false,
		},
		{
			name: "empty report",
			fields: fields{
				Sess: &session.Session{
					Environment: wd.EnvironmentProduction,
					Tenant:      "Tenant",
					Credentials: &credentials.BasicAuthCredentials{
						Username: "Username",
						Password: "Password",
					},
				},
				Body: []byte(`
					<Report_Data></Report_Data>`,
				),
				Error: nil,
			},
			args: args{
				ctx: context.Background(),
				def: raas.Definition{
					Owner: "Owner",
					Name:  "Name",
				},
				query: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "backend error",
			fields: fields{
				Sess: &session.Session{
					Environment: wd.EnvironmentProduction,
					Tenant:      "Tenant",
					Credentials: &credentials.BasicAuthCredentials{
						Username: "Username",
						Password: "Password",
					},
				},
				Body:  nil,
				Error: errors.New("error"),
			},
			args: args{
				ctx: context.Background(),
				def: raas.Definition{
					Owner: "Owner",
					Name:  "Name",
				},
				query: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "malformed xml",
			fields: fields{
				Sess: &session.Session{
					Environment: wd.EnvironmentProduction,
					Tenant:      "Tenant",
					Credentials: &credentials.BasicAuthCredentials{
						Username: "Username",
						Password: "Password",
					},
				},
				Body: []byte(`
					Report_Data>
						<Report_Entry>
							<S1>string 1</S1>
							<I1>1</I1>
						</Report_Entry>
					</Report_Data>`,
				),
				Error: nil,
			},
			args: args{
				ctx: context.Background(),
				def: raas.Definition{
					Owner: "Owner",
					Name:  "Name",
				},
				query: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backend := new(mocks.Backend)
			backend.On("Call", tt.args.ctx, tt.fields.Sess, tt.args.def, tt.args.query).
				Return(ioutil.NopCloser(bytes.NewReader(tt.fields.Body)), tt.fields.Error)
			c := &raas.Client{
				Backend: backend,
				Sess:    tt.fields.Sess,
			}
			var got []testData
			if err := c.CallXML(tt.args.ctx, tt.args.def, tt.args.query, &got); (err != nil) != tt.wantErr {
				t.Errorf("CallXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CallRaw() got = %v, want %v", got, tt.want)
			}
			backend.AssertExpectations(t)
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		sess *session.Session
	}
	tests := []struct {
		name string
		args args
		want *raas.Client
	}{
		{
			name: "base case",
			args: args{
				sess: &session.Session{
					Environment: wd.EnvironmentImplementation,
					Tenant:      "tenant",
					Credentials: nil,
				},
			},
			want: &raas.Client{
				Backend: &raas.APIBackend{},
				Sess: &session.Session{
					Environment: wd.EnvironmentImplementation,
					Tenant:      "tenant",
					Credentials: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := raas.NewClient(tt.args.sess); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

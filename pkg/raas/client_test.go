package raas

import (
	"net/url"
	"testing"

	"github.com/ianlopshire/go-workday/pkg/wd"
)

func TestClient_URL(t *testing.T) {
	for _, tt := range []struct {
		name       string
		session    *wd.Session
		definition Definition
		params     url.Values
		url        string
	}{
		{
			name: "prod with params",
			session: &wd.Session{
				Environment: wd.EnvironmentProduction,
				Tenant:      "tenant1",
			},
			definition: Definition{
				Name:  "Name1",
				Owner: "Owner1",
			},
			params: url.Values{"worker": []string{"9999"}},
			url:    "https://services1.myworkday.com/ccx/service/customreport2/tenant1/Owner1/Name1?worker=9999",
		},
		{
			name: "impl with params",
			session: &wd.Session{
				Environment: wd.EnvironmentImplementation,
				Tenant:      "tenant1",
			},
			definition: Definition{
				Name:  "Name1",
				Owner: "Owner1",
			},
			params: url.Values{"worker": []string{"9999"}},
			url:    "https://wd2-impl-services1.workday.com/ccx/service/customreport2/tenant1/Owner1/Name1?worker=9999",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				session: tt.session,
			}

			url := c.url(tt.definition, tt.params)

			if url != tt.url {
				t.Errorf("url() have %v, want %v", url, tt.url)
			}
		})
	}
}

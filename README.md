# go-workday

A collection of go packages to aid with Workday integrations.

## Notice

I no longer have access to a Workday tenant to test this package. If you can or know of
someone who can provide a test tenant please reach out.

## Packages

| Name        | Path                  | Description                                        |
| ----------- | --------------------- | -------------------------------------------------- |
| raas        | `/pkg/raas/`          | Reports as a Service (RaaS) client and utilities   |
| wd          | `/pkg/wd/`            | Shared Workday type and utilities                  |
| session     | `/pkg/wd/session`     | Provides configuration for Workday service clients |
| credentials | `/pkg/wd/credentials` | Credential retrieval and management                |

## Usage

### Session

Session provides configuration for Workday service clients.

#### From Environment

The following environment variables can be used to configure a session.

| Name                                  | Description                                           |
| ------------------------------------- | ----------------------------------------------------- |
| `WD_ENVIRONMENT`                      | Workday Environment: `Production` or `Implementation` |
| `WD_TENANT`                           | Workday tenant name                                   |
| **Basic Auth**                        |                                                       |
| `WD_BASIC_AUTH_USERNAME`              | Basic auth username                                   |
| `WD_BASIC_AUTH_PASSWORD`              | Basic auth password                                   |
| **OAuth Client Credentials**          |                                                       |
| `WD_CLIENT_CREDENTIALS_CLIENT_ID`     | OAuth client credentials client ID                    |
| `WD_CLIENT_CREDENTIALS_CLIENT_SECRET` | OAuth client credentials secret                       |
| `WD_CLIENT_CREDENTIALS_REFRESH_TOKEN` | OAuth client credentials refresh token                |


```go
import (
    "github.com/ianlopshire/go-workday/pkg/wd/session"
)

sess, err := session.NewSession()
if err != nil {
    log.fatal(err)
}
```

#### Manual Config

The session can also be manually configured.

```go
import (
	"github.com/ianlopshire/go-workday/pkg/wd"
	"github.com/ianlopshire/go-workday/pkg/wd/credentials"
	"github.com/ianlopshire/go-workday/pkg/wd/session"
)

config := session.Config{
	Environment: wd.String(wd.EnvironmentImplementation),
	Tenant:      wd.String("Tenant"),
	Credentials: &credentials.BasicAuthCredentials{
		Username: "username",
		Password: "password",
	},
}

sess, err := session.NewSession(session.FromConfig(config))
if err != nil {
	log.Fatal(err)
}
```

### Reports as a Service (RaaS)

The RaaS client aids in fetching Workday data via RaaS.

```go
sess, err := session.NewSession()
if err != nil {
	log.Fatal(err)
}
client := raas.NewClient(sess)

def := raas.Definition{
	Owner: "imlopshire",
	Name:  "example_worker_report",
}

// Define the layout of the data in the custom report.
var layout []struct {
	WID            string `xml:"wid"`
	Name           string `xml:"name"`
	YearsOfService int    `xml:"years_of_service"`
}

// Define the values for any report prompts.
query := url.Values{}
query.Set("workers", "Employee_ID:1001")

// Fetch the report data and decode into layout.
err = client.CallXML(context.Background(), def, query, &layout)
if err != nil {
	log.Fatal(err)
}

// Use the resulting data.
for i := range layout {
	fmt.Printf("%+v", layout[i])
}
```
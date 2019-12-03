# WWSGen

WWSGen is a command line utility that automatically generates Workday Web Service clients.

## Basic Usage

```
Usage of wwsgen:
  -output string
    	output directory (default "./")
  -pkg string
    	package name to be used int the output (default "gen")
  -service string
    	WWS service to generate (e.g. Absence_Management)
  -version string
    	Workday Web Service version (default "v30.1")
```

#### Output File Structure

```
gen
├─── wwsgen_client.go   # Contains the auto generated service client.
└─── wwsgen_types.go    # Contains all of the auto generated types.
```

### Go Generate

[Go Generate](https://blog.golang.org/generate) is the preferred method of use for WWSGen.

```go
//go:generate -command wwsgen go run .path/to/cmd/wwsgen/main.go
//go:generate wwsgen --service=Human_Resources --version=v31.2 --pkg=hr

package hr
```

## Overview of Process

WWSGen does two basic operations in order to generate WWS clients.

1. Generate required types via `aqwari.net/xml/xsdgen`
   * Specific changes had to be made to `aqwari.net/xml/xsdgen`. These changes have been vendored at [github.com/ianlopshire/go-xml/tree/empx-wwsgen](https://github.com/ianlopshire/go-xml/tree/empx-wwsgen)
   * Output to `wwsgen_types.go`
2. Generate a service client via `internal/wwsgen`
   * Custom Workday-specific SOAP client auto gen
   * Output to `wwsgen_client.go`

## Known Issues

* Build error due to invalid type usage (See github.com/droyo/go-xml/issues/61)
  * These errors require manual fixing
  * Example of error in commit [f20c5d6](https://github.com/ianlopshire/go-workday/commit/f20c5d6dedce092b3de774c887b8a16ee9a57cb2)
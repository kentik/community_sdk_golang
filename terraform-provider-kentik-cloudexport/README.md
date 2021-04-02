# Terraform Provider for Kentik Cloud Export

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.14.x
- [Go](https://golang.org/doc/install) >= 1.15

## Install

Build and install the plugin so that Terraform can find and use it:

```bash
make install
```

## Test

### Unit tests

Unit tests run the provider against a localhost_apiserver that serves data read from CloudExportTestData.json.  
This allows to:
- avoid the necessity of providing valid API credentials
- avoid creating resources on remote server
- make the test results more reliable

```bash
make test
```

This will:
1. build and run localhost_apiserver that is a stub for kentik apiv6 server
1. run tests (communication with localhost_apiserver)
1. shut down localhost_apiserver


### Acceptance tests

Acceptance tests run the provider against live server

```bash
make testacc
```

You need to provide valid credentials as environment variables so that the provider can communicate with the server:
- KTAPI_AUTH_EMAIL
- KTAPI_AUTH_TOKEN

*Note:* Acceptance tests create real resources.

## Debug

```bash
make build
dlv exec ./terraform-provider-kentik-cloudexport
r -debug
c
# attach with terraform following the just-printed out instruction in your terminal
```

## Using the provider

In folder with Terraform .tf definition file for cloud export resources/data sources:

```bash
terraform init
terraform apply
```

Note: you need to provide kentikapi credentials and also you can provide custom apiserver url, either in .tf file:
```terraform
provider "kentik-cloudexport" {
  email="john@acme.com"
  token="test123"
  # apiurl= "http://localhost:8080" # custom apiserver
}
```

or as environment variables:

```bash
export KTAPI_AUTH_EMAIL="john@acme.com"
export KTAPI_AUTH_TOKEN="test123"
# export KTAPI_URL="http://localhost:8080" # custom apiserver
```

See: [examples](./examples/)  

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.


## TODO

- add acceptance tests that communicate with live api server (once api server "HTTP DELETE" issue is fixed)
- publish to registry.terraform.io

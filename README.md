# community_sdk_golang
[![Go Reference](https://pkg.go.dev/badge/github.com/kentik/community_sdk_golang.svg)](https://pkg.go.dev/github.com/kentik/community_sdk_golang)

This repository is the Kentik Go SDK for the community.

[kentikapi](kentikapi) package contains Go client library for [Kentik REST API](https://kb.kentik.com/v0/Ab09.htm)

## Installation

Install the library:

```bash
go get github.com/kentik/community_sdk_golang/kentikapi
```

## Usage examples

Library usage examples are located in [examples](examples) directory.
Note that they are placed in Go test files (e.g. _users_example_test.go_) to be easily runnable.

Run an example:

```bash
export KTAPI_AUTH_EMAIL=<your kentik api credentials email>
export KTAPI_AUTH_TOKEN=<your kentik api credentials token>

# Run from the repository root
go test -tags examples -count 1 -v ./examples/users_example_test.go
```

To configure timeout for a client call use _context.WithTimeout()_ and pass it to the request function.
[Example use.](./apiv6/examples/cloud_export/main.go)

## Development

Anybody who wants to contribute to development is welcome to provide pull requests.

Run tests: `go test ./...`

Run all tests, including usage examples: `go test -tags examples ./...`

Install linters runner: [golangci-lint local installation](https://golangci-lint.run/usage/install/#local-installation)  
Run golangci-lint: `golangci-lint run ./...`

Subset of the code is generated from API specification available at:
- <https://github.com/kentik/api-schema-public/tree/master/gen/openapiv2/kentik/cloud_export/v202101beta1>
- <https://github.com/kentik/api-schema-public/tree/master/gen/openapiv2/kentik/synthetics/v202101beta1>

The OpenAPI Generator is used for code generation: <https://openapi-generator.tech/>.
Generated code is checked-in to the repository, so that the user can _go get_ the library.

Additional development steps for API v6:
- Access apiv6 directory: cd apiv6
- Generate the client: `./generate_client.sh`
- Generate the stub server: `./generate_server.sh`
- Generate documentation: `./generate_docs.sh`

Note that due to the design of OpenAPI generator, some generated files of the stub server need to be filled manually. They are listed in following files:
- [./apiv6/localhost_apiserver/cloudexport/.openapi-generator-ignore](./apiv6/localhost_apiserver/cloudexport/.openapi-generator-ignore)
- [./apiv6/localhost_apiserver/synthetics/.openapi-generator-ignore](./apiv6/localhost_apiserver/synthetics/.openapi-generator-ignore)

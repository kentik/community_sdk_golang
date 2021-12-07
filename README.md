# community_sdk_golang

[![Go Reference](https://pkg.go.dev/badge/github.com/kentik/community_sdk_golang.svg)](https://pkg.go.dev/github.com/kentik/community_sdk_golang)

This repository is the Kentik Go SDK for the community.

[kentikapi](kentikapi) package contains Go client library for [Kentik APIs](https://kb.kentik.com/v0/Ab09.htm).

## Requirements

- [Go](https://golang.org/doc/install) >= 1.15

## Installation

Install the library:

```bash
go get github.com/kentik/community_sdk_golang/kentikapi
```

## Usage

To use the SDK, import packages of _github.com/kentik/community_sdk_golang_ Go module.

Library documentation: <https://pkg.go.dev/github.com/kentik/community_sdk_golang>

Usage examples: [examples](./examples)
Note that examples are placed in Go test files (e.g. _users_example_test.go_) to be easily runnable.

### Running examples

Run an example:

```bash
export KTAPI_AUTH_EMAIL=<Kentik API authentication email>
export KTAPI_AUTH_TOKEN=<Kentik API authentication token>

# Run from a Go module, e.g. the root of this repository
# Adjust -run parameter to filter example names
go test -tags examples -count 1 -v -run Users github.com/kentik/community_sdk_golang/examples
```

To configure timeout for a client call for Cloud Export or Synthetics use _context.WithTimeout()_ and pass it to the request function. See [an example](examples/cloud_export_example_test.go).

## Development

Anybody who wants to contribute to development is welcome to provide pull requests. To work on the SDK, install tools listed in [requirements section](#requirements).

Optional tools:
- _golangci-lint_: <https://golangci-lint.run/usage/install/#local-installation>

Development steps:
- Compile the code: `go build ./...`
- Run tests: `go test ./...`
- Run all tests, including usage examples: `go test -tags examples ./...`
- Run golangci-lint: `golangci-lint run ./...`
- Format the code: `./tools/fmt.sh`
- Generate code for Synthetics and Cloud Export:
  - Generate the client: `./generate_client.sh`
  - Generate the fake server: `./generate_server.sh`
  - Generate documentation: `./generate_docs.sh`

Subset of the code is generated from API specification available at:
- <https://github.com/kentik/api-schema-public/tree/master/gen/openapiv2/kentik/cloud_export/v202101beta1>
- <https://github.com/kentik/api-schema-public/tree/master/gen/openapiv2/kentik/synthetics/v202101beta1>

The OpenAPI Generator is used for code generation: <https://openapi-generator.tech/>.
Generated code is checked-in to the repository, so that the user can _go get_ the library.

Note that the generated code scaffolding needs to be complemented with actual logic implementation. Here are the lists of generated files of the fake server that need to be filled manually:
- [./kentikapi/fakeapiserver/cloudexport/.openapi-generator-ignore](kentikapi/fakeapiserver/cloudexport/.openapi-generator-ignore)
- [./kentikapi/fakeapiserver/synthetics/.openapi-generator-ignore](kentikapi/fakeapiserver/synthetics/.openapi-generator-ignore)

### Release

Release process for the SDK is based on Git repository tags that follow [semantic versioning](https://semver.org/).

To release the SDK:
1. Make sure that all code that you want to release is in _master_ branch.
1. Navigate to [repository releases page](https://github.com/kentik/community_sdk_golang/releases), click _Draft a new release_ button and put tag version (in _v\[0-9].\[0-9].\[0-9]_ format), name and description.

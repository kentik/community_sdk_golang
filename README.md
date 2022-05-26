# community_sdk_golang

[![Go Reference](https://pkg.go.dev/badge/github.com/kentik/community_sdk_golang.svg)](https://pkg.go.dev/github.com/kentik/community_sdk_golang)

This repository is the Kentik Go SDK for the community. It is distributed as _github.com/kentik/community_sdk_golang_ Go module.

[kentikapi](kentikapi) package contains Go client library for [Kentik APIs](https://kb.kentik.com/v0/Ab09.htm).

## Requirements

- [Go](https://golang.org/doc/install) >= 1.15

## Installation

Install the library:

```bash
go get github.com/kentik/community_sdk_golang/kentikapi
```

## Usage

Library documentation published to [pkg.go.dev](https://pkg.go.dev/github.com/kentik/community_sdk_golang) describes the public API. Top level `kentikapi.Client` object contains multiple services that allow to manipulate corresponding Kentik resources, e.g. _Users_, _CloudExports_.

Usage examples are placed in the [examples](./examples) directory. Note that examples are placed in Go test files (e.g. _users_example_test.go_) to be easily runnable. Head over to [users_example_test.go](./examples/users_example_test.go) to view how to create and use the Client for simple CRUD use case.

### Running examples

Run an example:

```bash
export KTAPI_AUTH_EMAIL=<Kentik API authentication email>
export KTAPI_AUTH_TOKEN=<Kentik API authentication token>

# Run from a Go module, e.g. the root of this repository
# Adjust -run parameter to filter example names
go test -tags examples -count 1 -parallel 1 -v -run Users github.com/kentik/community_sdk_golang/examples
```

### Errors

Library provides a set of [functions](./kentikapi/errors.go) to check error type returned by the client.

## Contributing

Anybody who wants to contribute to development is welcome to provide pull requests. To work on the SDK, install tools listed in [requirements section](#requirements).

Optional tools:
- _golangci-lint_: <https://golangci-lint.run/usage/install/#local-installation>

Development steps:
- Compile the code: `go build -tags examples ./...`
- Run tests: `go test ./...`
- Run all tests, including usage examples: `go test -tags examples ./...`
- Run golangci-lint: `golangci-lint run ./...`
- Format the code: `./tools/fmt.sh`
- Check Go module consistency: `./tools/check-go-mod.sh`

The project uses ADR to record the architectural decisions: [docs/adr/0001-record-architecture-decisions.md](docs/adr/0001-record-architecture-decisions.md)

### Release

Release process for the SDK is based on Git repository tags that follow [semantic versioning](https://semver.org/).

To release the SDK:
1. Make sure that all code that you want to release is in _master_ branch.
1. Navigate to [repository releases page](https://github.com/kentik/community_sdk_golang/releases), click _Draft a new release_ button and put tag version (in _v\[0-9].\[0-9].\[0-9]_ format), name and description.

# community_sdk_golang

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

## Development

Run tests: `go test ./...`

Run all tests, including usage examples: `go test -tags examples ./...`

Install linters runner: [golangci-lint local installation](https://golangci-lint.run/usage/install/#local-installation)  
Run golangci-lint: `golangci-lint run ./...`

## Development state

Finished implementation:
- devices (with interfaces)
- plans

Ongoing implementation:
- users
- my kentik portal

ToDo:
- sites
- tags
- device labels
- custom dimensions (with populators)
- custom applications
- saved filters
- query methods
- alerts
- alerts active

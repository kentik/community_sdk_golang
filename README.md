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
go test -tags examples -v ./examples/users_example_test.go
```

## Development

Run tests: `go test ./...`

Run all tests, including usage examples: `go test -tags examples ./...`

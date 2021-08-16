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
./tools/test.sh -v apiv5 -t TestDemonstrateUsersGetAll
# -v directory 
# -t test function name
# ommiting both will test apiv5 and apiv6 
# ommiting only -t will run all tests in -v
```

## Development

Run tests: `go test ./...`

Run all tests, including usage examples: `go test -tags examples ./...`

Install linters runner: [golangci-lint local installation](https://golangci-lint.run/usage/install/#local-installation)  
Run golangci-lint: `golangci-lint run ./...`

## Development state

Implemented API resources:
- users
- sites
- tags
- devices (with interfaces)
- device labels
- custom dimensions (with populators)
- custom applications
- saved filters
- my kentik portal
- query methods
- plans
- alerts
- alerts active

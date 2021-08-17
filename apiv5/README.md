## Installation

Install the library:

```bash
go get github.com/kentik/community_sdk_golang/apiv5/kentikapi
```

## Usage examples

Library usage examples are located in [examples](examples) directory.
Note that they are placed in Go test files (e.g. _users_example_test.go_) to be easily runnable.

Run an example:

```bash
export KTAPI_AUTH_EMAIL=<your kentik api credentials email>
export KTAPI_AUTH_TOKEN=<your kentik api credentials token>

# Run from the apiv5 directory
go test -tags examples -count=1 -run TestDemonstrateUsersCRUD ./...
```

## Development

Run tests: `go test ./...`

Run all tests, including usage examples: `go test -tags examples ./...`

Install linters runner: [golangci-lint local installation](https://golangci-lint.run/usage/install/#local-installation)  
Run golangci-lint: `golangci-lint run ./...`

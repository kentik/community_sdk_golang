# community_sdk_golang - API v6

This module implements Go client for Kentik API v6.

## Development

Anybody who wants to contribute to development is welcome to provide pull requests.

Subset of the code is generated from API specification available at:
- <https://github.com/kentik/api-schema-public/tree/master/gen/openapiv2/kentik/cloud_export/v202101beta1>
- <https://github.com/kentik/api-schema-public/tree/master/gen/openapiv2/kentik/synthetics/v202101beta1>

The OpenAPI Generator is used for code generation: <https://openapi-generator.tech/>.
Generated code is checked-in to the repository, so that the user can _go get_ the library.

Development steps:
- Run tests: `go test ./...`
- Format the code: `go fmt ./...`
- Generate the client: `./generate_client.sh`
- Generate the stub server: `./generate_server.sh`
- Generate documentation: `./generate_docs.sh`

Note that due to the design of OpenAPI generator, some generated files of the stub server need to be filled manually. They are listed in following files:
- [./localhost_apiserver/cloudexport/.openapi-generator-ignore](./localhost_apiserver/cloudexport/.openapi-generator-ignore)
- [./localhost_apiserver/synthetics/.openapi-generator-ignore](./localhost_apiserver/synthetics/.openapi-generator-ignore)

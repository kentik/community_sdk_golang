#!/usr/bin/env bash
# Generate Go server stub from OpenAPI specification.

openapi_generator_tag="v5.1.1"

cloud_export_spec_filename="cloud_export.openapi.json"
cloud_export_spec_url="https://raw.githubusercontent.com/kentik/api-schema-public/master/gen/openapiv2/kentik/cloud_export/v202101beta1/cloud_export.swagger.json"
synthetics_spec_filename="synthetics.openapi.json"
synthetics_spec_url="https://raw.githubusercontent.com/kentik/api-schema-public/master/gen/openapiv2/kentik/synthetics/v202101beta1/synthetics.swagger.json"

function run() {
    check_prerequisites

    download_openapi_spec
    generate_cloud_export_server
    generate_synthetics_server
    go fmt ./...
}

function check_prerequisites() {
    stage "Checking prerequisites"

    if ! docker --version > /dev/null 2>&1; then
        echo "Please install Docker: https://docs.docker.com/get-docker/"
        exit 1
    fi

    if ! curl --version > /dev/null 2>&1; then
        echo "Please install curl: https://curl.se/"
        exit 1
    fi

    echo "Done"
}

function download_openapi_spec() {
    stage "Downloading OpenAPI specifications"

    curl --location --retry 20 "$cloud_export_spec_url" --output "$cloud_export_spec_filename"
    curl --location --retry 20 "$synthetics_spec_url" --output "$synthetics_spec_filename"

    echo "Done"
}

function generate_cloud_export_server() {
    cloudexport_package="cloudexportstub"
    cloudexport_spec="cloud_export.openapi.json"
    cloudexport_server_output_dir="localhost_apiserver/cloudexport"

    generate_go_server_from_openapi_spec "$cloudexport_spec" "$cloudexport_package" "$cloudexport_server_output_dir"
}

function generate_synthetics_server() {
    synthetics_package="syntheticsstub"
    synthetics_spec="synthetics.openapi.json"
    synthetics_server_output_dir="localhost_apiserver/synthetics"

    generate_go_server_from_openapi_spec "$synthetics_spec" "$synthetics_package" "$synthetics_server_output_dir"
}

function generate_go_server_from_openapi_spec() {
    stage "Generating Go server from OpenAPI specification"

    spec_file="$1"
    package="$2"
    output_dir="$3"

    docker run --rm  -v "$(pwd):/local" \
        "openapitools/openapi-generator-cli:$openapi_generator_tag" generate  \
        -i "/local/$spec_file" \
        -g go-server \
        --additional-properties=enumClassPrefix=true \
        --package-name "$package" \
        -o "/local/$output_dir"


    echo "Changing ownership of generated content to $USER:$USER"
    sudo chown  -R "$USER:$USER"  "$output_dir" # by default the generated output is in user:group root:root

    rm "$output_dir/Dockerfile"
    rm "$output_dir/go.mod"
    rm "$output_dir/main.go"
    rm "$output_dir/README.md"
    rm -rf "$output_dir/api"

    echo "Done"
}

function stage() {
    BOLD_BLUE="\e[1m\e[34m"
    RESET="\e[0m"
    msg="$1"

    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

run

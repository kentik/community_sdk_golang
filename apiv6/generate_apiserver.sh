#!/usr/bin/env bash

function generate_golang_server_from_openapi3_spec() {
    spec_file="$1"
    output_dir="$2"

    docker run --rm  -v "$(pwd)/local" \
        openapitools/openapi-generator-cli generate  \
        -i "/local/$spec_file" \
        -g go-server \
        --package-name cloudexportstub \
        -o "/local/$output_dir"

    echo "Changing ownership of generated content to $USER:$USER"
    sudo chown  -R "$USER:$USER"  "$output" # by default the generated output is in user:group root:root
}

input="api_spec/openapi_3.0.0/cloud_export.openapi.yaml"
output="localhost_apiserver"
generate_golang_server_from_openapi3_spec "$input" "$output"
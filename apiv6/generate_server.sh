#!/usr/bin/env bash

# Generate golang server stub from OpenAPI 3.0.0 spec


function stage() {
    BOLD_BLUE="\e[1m\e[34m"
    RESET="\e[0m"
    msg="$1"

    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

function generate_golang_server_from_openapi3_spec() {
    stage "Generating golang server from openapi spec"

    spec_file="$1"
    package="$2"
    output_dir="$3"

    docker run --rm  -v "$(pwd):/local" \
        openapitools/openapi-generator-cli generate  \
        -i "/local/$spec_file" \
        -g go-server \
        --package-name "$package" \
        -o "/local/$output_dir"


    echo "Changing ownership of generated content to $USER:$USER"
    sudo chown  -R "$USER:$USER"  "$output_dir" # by default the generated output is in user:group root:root
    
    rm "$output_dir/go.mod"
    rm "$output_dir/README.md"
    rm -rf "$output_dir/api"

    echo "Done"
}

# GENERATE CLOUDEXPORT
cloudexport_package="cloudexportstub"
cloudexport_spec="cloud_export.openapi.yaml"
cloudexport_server_output_dir="localhost_apiserver/cloud_export"
generate_golang_server_from_openapi3_spec "$cloudexport_spec" "$cloudexport_package" "$cloudexport_server_output_dir"

# GENERATE SYNTHETICS
synthetics_package="syntheticsstub"
synthetics_spec="synthetics.openapi.yaml"
synthetics_server_output_dir="localhost_apiserver/synthetics"
generate_golang_server_from_openapi3_spec "$synthetics_spec" "$synthetics_package" "$synthetics_server_output_dir"

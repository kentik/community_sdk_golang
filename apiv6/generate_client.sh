#!/usr/bin/env bash

# Generate golang client sdk from OpenAPI 3.0.0 spec


function stage() {
    BOLD_BLUE="\e[1m\e[34m"
    RESET="\e[0m"
    msg="$1"
    
    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

function checkPrerequsites() {
    stage "Checking prerequisites"
    
    if ! docker --version > /dev/null 2>&1; then
        echo "You need to install docker to run the generator"
        exit 1
    fi

    echo "Done"
}

function generate_golang_client_from_openapi3_spec() {
    stage "Generating golang client from openapi spec"

    spec_file="$1"
    package="$2"
    output_dir="$3"

    docker run --rm  -v "$(pwd):/local" \
        openapitools/openapi-generator-cli generate  \
        -i "/local/$spec_file" \
        -g go \
        --additional-properties=enumClassPrefix=true \
        --package-name "$package" \
        -o "/local/$output_dir"


    echo "Done"
}

function change_ownership_to_current_user() {
    dir="$1"
    stage "Changing ownership of $dir to $USER:$USER"

    sudo chown  -R "$USER:$USER"  "$dir" # by default the generated output is in user:group root:root

    echo "Done"
}

function cleanup_non_needed_files() {
    stage "Removing non-needed generated files"

    generated_content_dir="$1"

    rm "$generated_content_dir/go.mod"
    rm "$generated_content_dir/go.sum"
    rm "$generated_content_dir/README.md"
    rm -rf "$generated_content_dir/docs"

    echo "Done"
}

checkPrerequsites

# GENERATE CLOUDEXPORT
cloudexport_package="cloudexport"
cloudexport_spec="cloud_export.openapi.yaml"

cloudexport_client_output_dir="kentikapi/cloudexport"

generate_golang_client_from_openapi3_spec "$cloudexport_spec" "$cloudexport_package" "$cloudexport_client_output_dir"
change_ownership_to_current_user "$cloudexport_client_output_dir"
cleanup_non_needed_files "$cloudexport_client_output_dir"

# GENERATE SYNTHETICS
synthetics_package="synthetics"
synthetics_spec="synthetics.openapi.yaml"

synthetics_client_output_dir="kentikapi/synthetics"

generate_golang_client_from_openapi3_spec "$synthetics_spec" "$synthetics_package" "$synthetics_client_output_dir"
change_ownership_to_current_user "$synthetics_client_output_dir"
cleanup_non_needed_files "$synthetics_client_output_dir"
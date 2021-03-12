#!/usr/bin/env bash

# Generate golang client sdk and markdown docs from OpenAPI 3.0.0 spec


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
        --package-name "$package" \
        -o "/local/$output_dir"

    echo "Done"
}

function generate_markdown_from_openapi3_spec() {
    stage "Generating markdown docs from openapi spec"

    spec_file="$1"
    package="$2"
    output_dir="$3"

    docker run --rm  -v "$(pwd):/local" \
        openapitools/openapi-generator-cli generate  \
        -i "/local/$spec_file" \
        -g markdown \
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

cloudexport_package="cloudexport"
cloudexport_spec="api_spec/openapi_3.0.0/cloud_export.openapi.yaml"

cloudexport_client_output_dir="kentikapi/cloudexport"
cloudexport_docs_output_dir="docs/cloudexport"

checkPrerequsites
generate_golang_client_from_openapi3_spec "$cloudexport_spec" "$cloudexport_package" "$cloudexport_client_output_dir"
generate_markdown_from_openapi3_spec "$cloudexport_spec" "$cloudexport_package" "$cloudexport_docs_output_dir"
change_ownership_to_current_user "$cloudexport_client_output_dir"
change_ownership_to_current_user "$cloudexport_docs_output_dir"
#!/usr/bin/env bash
# Generate markdown docs from OpenAPI specification.

openapi_generator_tag="v5.1.1"

cloud_export_spec_filename="cloud_export.openapi.json"
cloud_export_spec_url="https://raw.githubusercontent.com/kentik/api-schema-public/master/gen/openapiv2/kentik/cloud_export/v202101beta1/cloud_export.swagger.json"
synthetics_spec_filename="synthetics.openapi.json"
synthetics_spec_url="https://raw.githubusercontent.com/kentik/api-schema-public/master/gen/openapiv2/kentik/synthetics/v202101beta1/synthetics.swagger.json"

function run() {
    check_prerequisites

    download_openapi_spec
    generate_cloud_export_docs
    generate_synthetics_docs
}

function check_prerequisites() {
    stage "Checking prerequisites"

    if ! docker --version > /dev/null 2>&1; then
        echo "You need to install Docker to run the generator"
        exit 1
    fi

    echo "Done"
}

function download_openapi_spec() {
    wget "$cloud_export_spec_url" --output-document "$cloud_export_spec_filename"
    wget "$synthetics_spec_url" --output-document "$synthetics_spec_filename"
}

function generate_cloud_export_docs() {
    cloudexport_package="cloudexport"
    cloudexport_spec="cloud_export.openapi.json"
    cloudexport_docs_output_dir="docs/cloudexport"

    generate_markdown_from_openapi3_spec "$cloudexport_spec" "$cloudexport_package" "$cloudexport_docs_output_dir"
    change_ownership_to_current_user "$cloudexport_docs_output_dir"
}

function generate_synthetics_docs() {
    synthetics_package="synthetics"
    synthetics_spec="synthetics.openapi.json"
    synthetics_docs_output_dir="docs/synthetics"

    generate_markdown_from_openapi3_spec "$synthetics_spec" "$synthetics_package" "$synthetics_docs_output_dir"
    change_ownership_to_current_user "$synthetics_docs_output_dir"
}

function generate_markdown_from_openapi3_spec() {
    stage "Generating markdown docs from openapi spec"

    spec_file="$1"
    package="$2"
    output_dir="$3"

    docker run --rm  -v "$(pwd):/local" \
        "openapitools/openapi-generator-cli:$openapi_generator_tag" generate  \
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

function stage() {
    BOLD_BLUE="\e[1m\e[34m"
    RESET="\e[0m"
    msg="$1"

    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

run

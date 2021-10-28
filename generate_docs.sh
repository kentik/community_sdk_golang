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
        echo "Please install Docker: https://docs.docker.com/get-docker/"
        die
    fi

    if ! curl --version > /dev/null 2>&1; then
        echo "Please install curl: https://curl.se/"
        die
    fi

    echo "Done"
}

function download_openapi_spec() {
    stage "Downloading OpenAPI specifications"

    curl --location --retry 20 "$cloud_export_spec_url" --output "$cloud_export_spec_filename" || die
    curl --location --retry 20 "$synthetics_spec_url" --output "$synthetics_spec_filename" || die

    echo "Done"
}

function generate_cloud_export_docs() {
    cloud_export_package="cloudexport"
    cloud_export_docs_output_dir="docs/cloudexport"

    generate_markdown_from_openapi3_spec "$cloud_export_spec_filename" "$cloud_export_package" "$cloud_export_docs_output_dir"
    change_ownership_to_current_user "$cloud_export_docs_output_dir"
}

function generate_synthetics_docs() {
    synthetics_package="synthetics"
    synthetics_docs_output_dir="docs/synthetics"

    generate_markdown_from_openapi3_spec "$synthetics_spec_filename" "$synthetics_package" "$synthetics_docs_output_dir"
    change_ownership_to_current_user "$synthetics_docs_output_dir"
}

function generate_markdown_from_openapi3_spec() {
    stage "Generating markdown docs from OpenAPI spec"

    spec_file="$1"
    package="$2"
    output_dir="$3"

    docker run --rm -v "$(pwd):/local" \
        "openapitools/openapi-generator-cli:$openapi_generator_tag" generate \
        -i "/local/$spec_file" \
        -g markdown \
        --package-name "$package" \
        -o "/local/$output_dir" || die

    echo "Done"
}

function change_ownership_to_current_user() {
    stage "Changing ownership of $dir to $USER"
    dir="$1"

    sudo chown -R "$USER" "$dir" || die # by default the generated output is in user:group root:root

    echo "Done"
}

function stage() {
    BOLD_BLUE="\e[1m\e[34m"
    RESET="\e[0m"
    msg="$1"

    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

function die() {
    echo "Error. Exit 1"
    exit 1
}

run

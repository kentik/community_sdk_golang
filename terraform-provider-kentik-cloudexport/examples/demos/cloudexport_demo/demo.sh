#!/usr/bin/env bash
# Showcase kentik-cloudexport Terraform provider against live Kentik API.

run() {
    check_prerequisites
    cleanup_tf_files

    install_tf_provider
    check_env

    stage "Initialize Terraform"
    pause_and_run pygmentize ./providers.tf
    pause

    pause_and_run terraform init || die
    pause

    stage "Create AWS cloud export"
    pause_and_run pygmentize ./aws.tf
    pause

    pause_and_run terraform plan || die
    pause
    pause_and_run terraform apply -auto-approve || die
    pause

    list_cloud_exports

    stage "Update AWS cloud export"
    pause_and_run sed -i 's/terraform aws cloud export/updated description/g' ./aws.tf
    pause_and_run pygmentize ./aws.tf
    pause

    pause_and_run terraform plan || die
    pause
    pause_and_run terraform apply -auto-approve || die
    pause
    sed -i 's/updated description/terraform aws cloud export/g' ./aws.tf

    list_cloud_exports

    stage "Delete AWS cloud export"
    pause_and_run terraform destroy -auto-approve

    list_cloud_exports
}

check_prerequisites() {
    if ! terraform -v > /dev/null 2>&1; then
        echo "Please install Terraform: https://learn.hashicorp.com/tutorials/terraform/install-cli"
        die
    fi

    if ! pygmentize -V > /dev/null 2>&1; then
        echo "Please install Pygments: https://pygments.org/"
        die
    fi

    if ! curl -V > /dev/null 2>&1; then
        echo "Please install cURL: https://curl.se/"
        die
    fi

    if ! jq -V > /dev/null 2>&1; then
        echo "Please install jq: https://stedolan.github.io/jq/"
        die
    fi
}

cleanup_tf_files() {
    rm -rf .terraform .terraform.lock.hcl terraform.tfstate
}

install_tf_provider() {
    stage "Build & install kentik-cloudexport Terraform provider"

    pushd ../../../ > /dev/null || die
    pause_and_run make install || die
    popd > /dev/null || die

    pause
}

check_env() {
    stage "Check auth env variables"

    if [[ -z "$KTAPI_AUTH_EMAIL" ]]; then
        echo "KTAPI_AUTH_EMAIL env variable must be set to Kentik API account email"
        die
    fi

    if [[ -z "$KTAPI_AUTH_TOKEN" ]]; then
        echo "KTAPI_AUTH_TOKEN env variable must be set to Kentik API authorization token"
        die
    fi

    echo "Print KTAPI_AUTH_EMAIL"
    echo "$KTAPI_AUTH_EMAIL"
    echo "Print KTAPI_AUTH_TOKEN (first 10 chars)"
    echo "${KTAPI_AUTH_TOKEN:0:10}"

    pause
}

list_cloud_exports() {
    read -r -p "Press any key to list Cloud Exports with cURL on https://cloudexports.api.kentik.com/cloud_export/v202101beta1/exports"
    curl --location --request GET --max-time 30 "https://cloudexports.api.kentik.com/cloud_export/v202101beta1/exports" \
        --header "X-CH-Auth-Email: $KTAPI_AUTH_EMAIL" \
        --header "X-CH-Auth-API-Token: $KTAPI_AUTH_TOKEN" | jq
    pause
}

stage() {
    BLUE_BOLD="\e[1m\e[34m"
    RESET="\e[0m"
    msg="$1"

    echo
    echo -e "$BLUE_BOLD$msg$RESET"
}

pause() {
    echo
    read -r -p "Press any key to continue..."
}

pause_and_run() {
    read -r -p "Press any key to run '$*'"
    "$@"
}

die() {
    echo "Error. Exit 1"
    exit 1
}

run

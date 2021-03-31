#!/usr/bin/env bash

function stage() {
    BOLD_BLUE="\e[1m\e[96m"
    RESET="\e[0m"
    msg="$1"
    
    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

function die() {
    echo "Error. Exit 1"
    exit 1
}

if [[ -z "$KTAPI_AUTH_EMAIL" ]]; then
    echo "KTAPI_AUTH_EMAIL env variable must be set to kentikapi account email"
    die
fi

if [[ -z "$KTAPI_AUTH_TOKEN" ]]; then
    echo "KTAPI_AUTH_TOKEN env variable must be set to kentikapi authorization token"
    die
fi

stage "Kentik CloudExport Terraform Provider example - Kentik apiserver"
echo "The provider will connect to live Kentik apiserver"

stage "Build & install plugin"
pushd ../../  > /dev/null || die
make install || die
popd  > /dev/null || die

stage "Terraform init & apply"
rm -rf .terraform .terraform.lock.hcl

# export TF_LOG=ERROR
terraform init || die
terraform apply
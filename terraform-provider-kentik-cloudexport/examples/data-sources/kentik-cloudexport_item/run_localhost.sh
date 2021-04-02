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

apiurl="http://localhost:8080" # localhost apiserver url

stage "Kentik CloudExport Terraform Provider example - localhost apiserver"
echo "Please make sure apiserver at $apiurl is running"

stage "Build & install plugin"
pushd ../../../  > /dev/null || die
make install || die
popd  > /dev/null || die

stage "Terraform init & apply"
rm -rf .terraform .terraform.lock.hcl

# export TF_LOG=ERROR
export KTAPI_URL="$apiurl" 
export KTAPI_AUTH_EMAIL="dummy@acme.com" # authorization not used by localhost_apiserver
export KTAPI_AUTH_TOKEN="dummy"          # authorization not used by localhost_apiserver
terraform init || die
terraform apply
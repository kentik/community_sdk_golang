#!/usr/bin/env bash

function stage() {
    BOLD_BLUE="\e[1m\e[96m"
    RESET="\e[0m"
    msg="$1"
    
    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

function die() {
    echo "Exit 1"
    exit 1
}

stage "Build & install plugin"
pushd ../../../  > /dev/null || die
make install || die
popd  > /dev/null || die

stage "Terraform init & apply"
rm -rf .terraform .terraform.lock.hcl

# export TF_LOG=ERROR
terraform init || die
terraform apply
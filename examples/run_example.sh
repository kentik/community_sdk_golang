#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
REPO_DIR=$(cd -- "$SCRIPT_DIR" && cd ../ && pwd)

source "$REPO_DIR/tools/utility_functions.sh" || exit 1

# INPUT:
#   [none] - run all examples
#   User
#   Device

function stage() {
    echo
    colored_echo BOLD_BLUE "$1"
}

if [[ $1 != "" ]]; then
    stage "Running single example: $1"
    go test -tags examples -count 1 -run "$1"  ./... -v
else
    stage "Running all examples"
    go test -tags examples -count 1 ./... -v
fi

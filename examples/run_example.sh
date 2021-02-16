#!/usr/bin/env bash

# INPUT:
#   none - run all tests
#   test name - run specified test

function stage() {
    BLUE_BOLD="\e[1m\e[34m"
    RESET="\e[0m"
    msg="$1"
    
    echo
    echo -e "$BLUE_BOLD$msg$RESET"
}

if [[ $1 != "" ]]; then
    stage "Running single example: $1"
    go test -tags examples -run $1  ./... -v
else 
    stage "Running all examples"
    go test -tags examples ./... -v
fi
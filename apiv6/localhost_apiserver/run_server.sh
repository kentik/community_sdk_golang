#!/usr/bin/env bash

function stage() {
    BOLD_BLUE="\e[1m\e[96m"
    RESET="\e[0m"
    msg="$1"
    
    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

function checkPrerequsites() {
    stage "Checking prerequisites"
    
    if ! go version > /dev/null 2>&1; then
        echo "You need to install go to run the kentik api server stub"
        exit 1
    fi

    echo "Done"
}

function run() {
    stage "Running localhost api server on :8080"

    go run . -addr ":8080" -storage "CloudExportStorage.json"

    echo "Done"
}

checkPrerequsites
run
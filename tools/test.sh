#!/usr/bin/env bash

# INPUT:
#   [none] - run all examples
#   v - api version: apiv5 or apiv6
#   test name -

function stage() {
    BLUE_BOLD="\e[1m\e[34m"
    RESET="\e[0m"
    msg="$1"

    echo
    echo -e "$BLUE_BOLD$msg$RESET"
}

while getopts v:t: flag
do
  case "${flag}" in
    v) dir=${OPTARG};;
    t) testName=${OPTARG};;
  esac
done

if [ -z "$dir" ] # if $dir is empty
then
    stage "Running all examples"
    cd apiv5
    go test -tags examples -count=1 ./...
    cd ../apiv6
    go test -tags examples ./...
    cd ../
else
    if [ -z "$testName" ]
    then
        stage "Running all examples in $dir"
        cd "$dir"
        go test -tags examples -count=1 ./...
        cd ../
    else
        stage "Running single example: $testName in $dir"
        cd "$dir/examples"
        go test -tags examples -count=1 -run "^$testName$"
        cd ../../
    fi
fi
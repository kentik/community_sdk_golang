#!/usr/bin/env bash

# INPUT:
#   [none] - run all examples
#   User
#   Device

if [[ $1 != "" ]]; then
    echo "Running single example: $1"
    go test -tags examples -count 1 -run "$1"  ./... -v
else
    echo "Running all examples"
    go test -tags examples -count 1 ./... -v
fi

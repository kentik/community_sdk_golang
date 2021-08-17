#!/usr/bin/env bash

# INPUT:
#   [none] - run all examples

pushd apiv5
go test "$@" ./...
popd
pushd apiv6
go test "$@" ./...
popd

#!/usr/bin/env bash
# Install gofumpt and format Go code.

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
REPO_DIR=$(dirname "${SCRIPT_DIR}")

function run() {
    go install mvdan.cc/gofumpt@latest

    GO_FILES=$(find "$REPO_DIR" -type f -name "*.go" -not -path "*gen_*")

    # shellcheck disable=SC2086
    gofumpt -l -w $GO_FILES
}

run

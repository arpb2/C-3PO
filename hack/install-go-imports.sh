#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go get -v -t -d ./...
#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go test -race -coverprofile=coverage.txt -covermode=atomic ./pkg/... ./cmd/...
#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

find_files() {
  find . -not \( \
      \( \
        -wholename './output' \
        -o -wholename './.git' \
        -o -wholename './_output' \
        -o -wholename './_gopath' \
        -o -wholename './release' \
        -o -wholename './target' \
        -o -wholename '*/third_party/*' \
      \) -prune \
    \) -name '*.go'
}

`go list -f {{.Target}} golang.org/x/tools/cmd/goimports` -w `find_files`
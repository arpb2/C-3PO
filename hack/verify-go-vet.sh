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

find_files | xargs "${go vet}" 2>&1
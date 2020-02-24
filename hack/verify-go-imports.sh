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

diff=$(`go list -f {{.Target}} golang.org/x/tools/cmd/goimports` -l `find_files`) || true
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Run ./hack/run-go-imports.sh" >&2
  exit 1
fi
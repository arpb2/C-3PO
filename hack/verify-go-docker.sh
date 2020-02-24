#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

docker build . --file build/docker/Dockerfile --tag c3po:$(date +%s)
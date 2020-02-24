#!/usr/bin/env bash

# Clear any existing docker-compose
sudo rm /usr/local/bin/docker-compose

# Fail instantly on any command fail
set -o errexit
set -o nounset
set -o pipefail

# Install 1.24.1
curl -L https://github.com/docker/compose/releases/download/1.24.1/docker-compose-`uname -s`-`uname -m` > docker-compose
chmod +x docker-compose
sudo mv docker-compose /usr/local/bin

cd build/docker

# Run detached and test
docker-compose up --build -d
sudo docker-compose run go go test ./test/...

cd -
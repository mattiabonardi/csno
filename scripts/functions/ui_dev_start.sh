#!/bin/bash
# Buildkit enables a faster build process
export DOCKER_BUILDKIT=1

# build image
cd ../../functions/ui
docker build -t ui-csno:latest .

cd ../

# open FaaS build
faas-cli build -f ui.yml

# Docker test
docker run --name ui-dev -p 8080:8080 --rm -ti ui:latest 
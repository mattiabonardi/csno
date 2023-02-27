#!/bin/bash
# Buildkit enables a faster build process
export DOCKER_BUILDKIT=1

# open faas cli login
sh faas_cli_login.sh

# Build, push and deploy function
cd ../../functions
faas-cli up -f ./ui.yml
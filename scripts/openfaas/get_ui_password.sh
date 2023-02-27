#!/bin/bash

# Get UI portal password
PASSWORD=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo)
echo "User: admin"
echo "Password: ${PASSWORD}"
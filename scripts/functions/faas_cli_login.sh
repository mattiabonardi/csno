#!/bin/bash

PASSWORD=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo)
faas-cli login --gateway http://localhost:8000 --password $PASSWORD
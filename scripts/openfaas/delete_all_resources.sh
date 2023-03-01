#!/bin/bash

# Remove all openfaas resources
kubectl delete all --all -n openfaas
kubectl delete all --all -n openfaas-fn
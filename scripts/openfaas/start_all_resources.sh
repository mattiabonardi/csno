#!/bin/bash

# Set permission
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml

# Intall OpenFaaS
arkade install openfaas

sleep 10s

# Forwards gateway to machine
kubectl port-forward -n openfaas svc/gateway 8000:8080 &

# Create and expose local registry
kubectl run registry --image=registry:latest --port=5000 --namespace openfaas
sleep 10s
kubectl expose pod registry --namespace openfaas --type=LoadBalancer --port=5000 --target-port=5000
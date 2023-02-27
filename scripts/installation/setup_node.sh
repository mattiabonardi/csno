#!/bin/bash

if [ "$EUID" -ne 0 ]
  then echo "Error: sudo required"
  exit
fi

# Install K3s
curl -sfL https://get.k3s.io | sh -

# Install FaaS CLI
curl -sL https://cli.openfaas.com | sudo sh

# Install Arkade
curl -sLS https://get.arkade.dev | sudo sh

# Install Helm
arkade get helm

# Set permission
sudo chmod 744 /etc/rancher/k3s/k3.yaml
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
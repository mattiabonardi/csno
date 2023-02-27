#!/bin/bash

kubectl get deployments -n openfaas -l "release=openfaas, app=openfaas"
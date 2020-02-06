#!/bin/bash

kubectl get job -n kyma-system simple-job

if [[ $? == 0 ]]; then
  kubectl delete job -n kyma-system simple-job
fi

kubectl apply -f job.yaml



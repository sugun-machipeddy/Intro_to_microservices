#!/bin/sh

kubectl delete deployment api1-deployment
kubectl delete deployment api2-deployment
kubectl delete deployment api3-deployment
kubectl delete service api1-service
kubectl delete service api2-service
kubectl delete service api3-service
istioctl delete gateway api1-gateway
istioctl delete virtualservice api1-virtualservice
istioctl delete destinationrule website

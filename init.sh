#!/bin/bash

kind create cluster --config kind-cluster.yml

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

sleep 5

while true; do
    echo "Checking for ingress creation..."
    if kubectl -n "ingress-nginx" describe pod --selector=app.kubernetes.io/component=controller --request-timeout "90s"  &> /dev/null; then
        echo "Pods in namespace \"ingress-nginx\" of type \"--selector=app.kubernetes.io/component=controller\" exist."
        break
    else
        sleep 5
    fi
done

kubectl wait --namespace ingress-nginx --for=condition=ready pod --selector=app.kubernetes.io/component=controller --timeout=90s

docker build -t bookstore:1.0.0 bookstore

docker build -t books:1.0.0 books

kind load docker-image bookstore:1.0.0

kind load docker-image books:1.0.0

kubectl apply -f books/deployment.yml

kubectl apply -f bookstore/deployment.yml

kubectl wait --namespace ingress-nginx --for=condition=ready pod --selector=app.kubernetes.io/component=controller --timeout=90s

kubectl apply -f ingress.yml

echo $'\n'

echo "Happy disruption!"
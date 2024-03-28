# Disruptor example

Basic example of chaos testing using [xk6-disruptor](https://grafana.com/docs/k6/latest/testing-guides/injecting-faults-with-xk6-disruptor/) extension.

In this example you can find three dummy services:

- Frontend books service which exposes a frontend application showing the list of books available and purchased.
- Books service which exposes an endpoint that list of available books to be purchased.
- Bookstore service which exposes two endpoints to purchase books using their name and to check which books where purchased.

All the data is stored in memory **without using any database** to make it simple and focus on the disruptor.

`⚠️ The ingress deployed will use the host's 80 and 443 ports`

## Requirements

- go: golang binary to at least install xk6 cli - https://go.dev/
- docker: used to create images and to run kind - https://www.docker.com/
- kind: used to create a local kubernetes cluster - https://kind.sigs.k8s.io/
- kubectl: cli used to interact with the kubernetes cluster - https://kubernetes.io/docs/reference/kubectl/
- xk6: custom k6 builder to create k6 binaries with custom extensions - https://github.com/grafana/xk6
- xk6-disruptor: extension to apply faults into Kubernetes clusters - https://grafana.com/docs/k6/latest/testing-guides/injecting-faults-with-xk6-disruptor/

### Prepare the cluster

1. First we create a kind cluster, for that we use kind-cluster.yml configuration file to expose ports on the host machine:

```sh
kind create cluster --config kind-cluster.yml
```

2. Next, we deploy an nginx ingress prepared for kind

```sh
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

3. Wait for the ingress to be ready

```sh
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```

4. Deploy ingress configuration

```sh
kubectl apply -f ingress.yml
```

5. Build the three images

```sh
docker build -t books-frontend:1.0.0 books-frontend
```

```sh
docker build -t bookstore:1.0.0 bookstore
```

```sh
docker build -t books:1.0.0 books
```

6. Deploy the images inside kind cluster

```sh
kind load docker-image books-frontend:1.0.0
```

```sh
kind load docker-image bookstore:1.0.0
```

```sh
kind load docker-image books:1.0.0
```

7. Deploy books-frontend, books and bookstore deployments into kind cluster

```sh
kubectl apply -f books-frontend/deployment.yml
```

```sh
kubectl apply -f books/deployment.yml
```

```sh
kubectl apply -f bookstore/deployment.yml
```

### Prepare a k6 binary with xk6-disruptor extension

1. Install xk6 cli

```sh
go install go.k6.io/xk6/cmd/xk6@latest
```

2. Build k6 binary with xk6-disruptor

```sh
xk6 build --output k6-disruptor --with github.com/grafana/xk6-disruptor
```

### Execute your disruptor test

1. Without disruptors

```sh
./k6-disruptor run --env SKIP_FAULTS=1 disrup-test.js
```

2. With disruptors

```sh
./k6-disruptor run disrup-test.js
```

The disruptor applied on the example test will inject http faults on the **books-service** causing it to have 500ms of latency and an error rate of 0.1.

### Clean up

If you want you can delete the kind cluster with the command:

```sh
kind delete cluster
```

And remove any image left:

```sh
docker image rm books:1.0.0
```

```sh
docker image rm bookstore:1.0.0
```

```sh
docker image rm books-frontend:1.0.0
```
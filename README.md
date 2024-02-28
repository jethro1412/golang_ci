# Github Action Example with Docker, Golang and K8s Deployment

This repository contains example of docker containerized golang application that can be deployed to kubernetes cluster

## Requirements

| Tool | Links |
| ------ | ------ |
| Go | [plugins/dropbox/README.md][PlDb] |
| Docker | [https://docs.docker.com/engine/install/][PlGh] |

## How to run locally?
You can build without docker by running following command
```sh
make run
```
you can access the service by accessing on localhost:8080 in your web browser

## How to build the docker container?
You need to set environment variables `VERSION` in order to tag the docker container and run following command
```sh
make build-docker
```

## How to deploy to kubernetes cluster?
- Build the image and upload it to dockerhub, you can utilize the github action by providing the docker hub username and auth token
- Deploy locally by using following command
```sh
kubectl apply -f deployments/deployment.yaml
```
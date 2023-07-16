
# Go Agent Stash

<p>
	<a  target="_blank">
			<img src="https://github.com/tiangolo/fastapi/workflows/Test/badge.svg?event=push&branch=master" alt="Test">
	</a>
	<a ><img alt="Static Badge" src="https://img.shields.io/badge/docker-processing-red?logo=docker"></a>
	<a ><img alt="Static Badge" src="https://img.shields.io/badge/kubernetes-processing-red?logo=kubernetes"></a>
	<a><img alt="Static Badge" <img alt="Static Badge" src="https://img.shields.io/badge/golang-1.20-brightgreen?logo=go"></a>
</p>

## Introduction

Go Agent Stash is a simple service receive logs from agents - save logs in to filesystem - sync logs to AWS S3

The key features are:
* **Fast transport gRPC**: Using gRPC protocol for fast transfer data between internal service of microservice system.
* **Gateway proxy**: Support gRPC Gateway Proxy to optional get data back by RESTful API or gRPC.
* **Schedule jobs**: Can setup schedule jobs to sync data to S3, clean local file system.
* **Support Docker & Kubernestes**: You can easy deploy to your own system with created scripts.

## Architecture

![Architecture](./architecture.png)

## Requirements
Golang 1.20

## Installation

```console
edit .env.example

make server

---> Your service will start in port=8080(http) & port=9090(gRPC) 
```

You can deploy namespace, deployment, services, ingress into your Kubernetes

```console
export KUBE_CONFIG=<You-.kube-config>

kubectl apply -f k8s/
```

You can build your own docker image

```console
docker run -t go-agent-stash . 
```

## Owner

<table>
  <tr>
    <td align="center"><a href="https://github.com/anthoai97"><img src="https://avatars.githubusercontent.com/u/85854989?v=4" width="100px;" alt="An Quach"/><br /><sub><b>An Thoai</b></sub></a><br /></td>
    </tr>
</table>

## License

This project is licensed under the terms of the MIT license.
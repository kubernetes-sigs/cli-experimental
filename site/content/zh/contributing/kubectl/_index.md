---
title: "Kubectl"
linkTitle: "Kubectl"
type: docs
weight: 1
description: >
    Contribute to Kubectl
---

## Attend a sig-cli meeting

The best way to get started is to attend [sig-cli](https://github.com/kubernetes/community/tree/master/sig-cli)
meetings.  The bug scrub is a great place to pick up an issue to work on.

## Checking out the code

### Install golang

Install the latest version of [go](https://golang.org)

### Get a copy of the code

Fork and clone the [kubernetes](https://github.com/kubernetes/kubernetes) repository

```sh
git clone git@github.com/USER/kubernetes
cd kubernetes
```

### Build the binary

Build the binary using `go build`

```sh
cd cmd/kubectl
go build -v
./kubectl version
```

### Edit the code

The kubectl code is under `staging/src/k8s.io/kubectl`.

- Libraries are under `staging/src/k8s.io/kubectl/pkg`
- Command implementations are under `staging/src/k8s.io/kubectl/pkg/cmd`

## Learning about libraries

Kubectl uses a number of common libraries

- [cobra](https://github.com/spf13/cobra) -- a golang framework for CLIs
- [client-go](https://github.com/kubernetes/client-go) -- libraries for talking to the Kubernetes apiserver
- [api](https://github.com/kubernetes/api) -- Kubernetes types
- [apimachinery](https://github.com/kubernetes/apimachinery) -- Kubernetes apimachinery libraries

## Additional resources

- [Everything You Always Wanted to Know About SIG-CLI but Were Afraid to Ask](https://www.youtube.com/watch?v=QVYQUQd7prE)
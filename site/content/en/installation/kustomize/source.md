---
title: "Go Source"
linkTitle: "Go Source"
weight: 2
type: docs
description: >
    Install Kustomize from the Go source code
---

Requires [Go] to be installed.

## Install the kustomize CLI from source without cloning the repo

For `go version` ≥ `go1.17`

```
GOBIN=$(pwd)/ GO111MODULE=on go install sigs.k8s.io/kustomize/kustomize/v5@latest
```

For `go version` < `go1.17`

```bash
GOBIN=$(pwd)/ GO111MODULE=on go get sigs.k8s.io/kustomize/kustomize/v5
```

## Install the kustomize CLI from local source with cloning the repo

```bash
# Need go 1.13 or higher
unset GOPATH
# see https://golang.org/doc/go1.13#modules
unset GO111MODULES

# clone the repo
git clone git@github.com:kubernetes-sigs/kustomize.git
# get into the repo root
cd kustomize

# Optionally checkout a particular tag if you don't
# want to build at head
git checkout kustomize/v5.0.0

# build the binary -- this installs the binary to your go bin path
make kustomize

# run it
~/go/bin/kustomize version
```

[Go]: https://golang.org

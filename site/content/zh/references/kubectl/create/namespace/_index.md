---
title: "namespace"
linkTitle: "namespace"
weight: 1
type: docs
description: >
    Create a namespace with the specified name.
---

Kubernetes supports multiple virtual clusters backed by the same physical cluster. These virtual clusters are called namespaces.


## Command
```bash
$ kubectl create namespace NAME [--dry-run=server|client|none]
```

## Example

### Command
```bash
$ kubectl create namespace my-namespace
```

### Output
```bash
$ kubectl get namespace
NAME                   STATUS   AGE
default                Active   41s
my-namespace           Active   11s
```



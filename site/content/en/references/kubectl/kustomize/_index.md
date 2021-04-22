---
title: "kustomize"
linkTitle: "kustomize"
weight: 1
type: docs
description: >
    Using kustomization.yaml
---

Print a set of API resources generated from instructions in a kustomization.yaml file.

The argument must be the path to the directory containing the file, or a git repository URL with a path suffix specifying same with respect to the repository root.

## Command
```bash
$ kubectl kustomize <dir>
```

## Example

### Input File
```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-deployment
spec:
  replicas: 5
  template:
    containers:
      - name: the-container
        image: registry/container:latest
```

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

nameSuffix: -dev

resources:
- deployment.yaml
```

### Command
```bash
// deployment.yaml and kustomization.yaml are in the same directory
$ kubectl kustomize
```

### Output
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-deployment-dev
spec:
  replicas: 5
  template:
    containers:
    - image: registry/container:latest
      name: the-container
```
---
title: "namespace"
linkTitle: "namespace"
type: docs
weight: 11
description: >
    Adds namespace to all resources.
---

Will override the existing namespace if it is set on a resource, or add it
if it is not set on a resource.

## Example

### File Input

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-deployment
  namespace: the-namespace
spec:
  replicas: 5
  template:
    containers:
      - name: the-container
        image: registry/conatiner:latest
```

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

namespace: kustomize-namespace

resources:
- deployment.yaml

```

### Build Output

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-deployment
  namespace: kustomize-namespace
spec:
  replicas: 5
  template:
    containers:
    - image: registry/conatiner:latest
      name: the-container
```
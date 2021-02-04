---
title: "nameSuffix"
linkTitle: "nameSuffix"
type: docs
weight: 11
description: >
    Appends the value to the names of all resources and references.
---

As `nameSuffix` is self explanatory, it helps adding suffix to names in the defined yaml files.

**Note:** The suffix is appended before the content hash if the resource type is ConfigMap or Secret.

## Example

### File Input

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
        image: registry/conatiner:latest
```

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

nameSuffix: -custom-suffix

resources:
- deployment.yaml

```

### Build Output

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-deployment-custom-suffix
spec:
  replicas: 5
  template:
    containers:
    - image: registry/conatiner:latest
      name: the-container
```
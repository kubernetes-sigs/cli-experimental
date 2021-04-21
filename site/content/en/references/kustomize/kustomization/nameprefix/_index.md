---
title: "namePrefix"
linkTitle: "namePrefix"
type: docs
weight: 9
description: >
    Prepends the value to the names of all resources and references.
---

As `namePrefix` is self explanatory, it helps adding prefix to names in the defined yaml files.

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

namePrefix: overlook-

resources:
- deployment.yaml

```

### Build Output

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: overlook-the-deployment
spec:
  replicas: 5
  template:
    containers:
    - image: registry/conatiner:latest
      name: the-container
```

{{< alert color="success" title="References" >}}
Apply will propagate the `namePrefix` to any place Resources within the project are referenced by other Resources
including:

- Service references from StatefulSets
- ConfigMap references from PodSpecs
- Secret references from PodSpecs
{{< /alert >}}

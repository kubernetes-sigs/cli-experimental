---
title: "labels"
linkTitle: "labels"
type: docs
weight: 9
description: >
    Add labels and optionally selectors to all resources.
---

A field that allows adding labels without also automatically injecting corresponding selectors.
This can be used instead of the `commonLabels` field, which always adds selectors.

{{% pageinfo color="warning" %}}
Selectors for resources such as Deployments and Services shouldn't be changed once the
resource has been applied to a cluster.

Changing includeSelectors to `true` in live resources could result in failures.
{{% /pageinfo %}}

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

labels:
  - pairs:
      someName: someValue
      owner: alice
      app: bingo
    includeSelectors: true # <-- false by default
```

## Example 1 - selectors NOT modified

### File Input

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

labels:
  - pairs:
      someName: someValue
      owner: alice
      app: bingo

resources:
- deploy.yaml
- service.yaml
```

```yaml
# deploy.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example
```

```yaml
# service.yaml
apiVersion: v1
kind: Service
metadata:
  name: example
```

### Build Output

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    app: bingo
    owner: alice
    someName: someValue
  name: example
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: bingo
    owner: alice
    someName: someValue
  name: example
```


## Example 2 - selectors modified

### File Input

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

labels:
  - pairs:
      someName: someValue
      owner: alice
      app: bingo
    includeSelectors: true 

resources:
- deploy.yaml
- service.yaml
```

```yaml
# deploy.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example
```

```yaml
# service.yaml
apiVersion: v1
kind: Service
metadata:
  name: example
```

### Build Output

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    app: bingo
    owner: alice
    someName: someValue
  name: example
spec:
  selector:
    app: bingo
    owner: alice
    someName: someValue
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: bingo
    owner: alice
    someName: someValue
  name: example
spec:
  selector:
    matchLabels:
      app: bingo
      owner: alice
      someName: someValue
  template:
    metadata:
      labels:
        app: bingo
        owner: alice
        someName: someValue
```

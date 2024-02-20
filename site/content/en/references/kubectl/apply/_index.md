---
title: "apply"
linkTitle: "apply"
weight: 1
type: docs
description: >
    Using apply Command
---

## apply with `YAML files`

Apply can be run directly against Resource Config files or directories using `-f`

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-deployment
spec:
  replicas: 5
  selector:
    matchLabels:
      app: container
  template:
    metadata:
      labels:
        app: container
    spec:
      containers:
      - name: the-container
        image: registry/container:latest
```

```bash
# Apply the Resource Config
kubectl apply -f deployment.yaml
```

This will apply the deployment file on the Kubernetes cluster. You can get the status by using a get command.

```bash
# Get deployments
kubectl get deployments
```

## apply with `Kustomize files`

Though Apply can be run directly against Resource Config files or directories using `-f`, it is recommended
to run Apply against a `kustomization.yaml` using `-k`.  The `kustomization.yaml` allows users to define
configuration that cuts across many Resources (e.g. namespace).

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

# list of Resource Config to be Applied
resources:
- deployment.yaml

# namespace to deploy all Resources to
namespace: default

# labels added to all Resources
commonLabels:
  app: example
  env: test
```

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

Users run Apply on directories containing `kustomization.yaml` files using `-k` or on raw
ResourceConfig files using `-f`.

```bash
# Apply the Resource Config
kubectl apply -k .

# View the Resources
kubectl get -k .
```

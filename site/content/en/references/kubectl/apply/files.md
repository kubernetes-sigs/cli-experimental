---
title: "Files"
linkTitle: "Files"
weight: 1
type: docs
description: >
    Applying Config YAML Files for creating Kubernetes Resources
---

Apply can be run directly against Resource Config files or directories using `-f`

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
  labels:
    component: nginx
    tier: frontend
spec:
  selector:
    matchLabels:
      component: nginx
      tier: frontend
  template:
    metadata:
      labels:
        component: nginx
        tier: frontend
    spec:
      containers:
      - name: nginx
        image: nginx:1.15.4
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
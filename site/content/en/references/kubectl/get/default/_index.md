---
title: "default"
linkTitle: "default"
weight: 2
type: docs
description: >
    Printing the status of the Kubernetes resources
---

## default
If no output format is specified, Get will print a default set of columns.

**Note:** Some columns *may* not directly map to fields on the Resource, but instead may
be a summary of fields.

### Command
```bash
kubectl get deployments nginx
```

### Output
```bash
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
nginx     1         1         1            0           5s
```
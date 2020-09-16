---
title: "Add"
linkTitle: "Add"
weight: 1
type: docs
description: >
    Adding Fields
---

### Adding Fields:

- Fields present in the Resource Config that are missing from the Resource will be added to the
  Resource.
- Fields will be added to the Last Applied Resource Config

```yaml
# deployment.yaml (Resource Config)
apiVersion: apps/v1
kind: Deployment
metadata:
  # ...
  name: nginx-deployment
spec:
  # ...
  minReadySeconds: 3
```

```yaml
# Original Resource
kind: Deployment
metadata:
  # ...
  name: nginx-deployment
spec:
  # ...
status:
  # ...
```

```yaml
# Applied Resource
kind: Deployment
metadata:
  # ...
  name: nginx-deployment
spec:
  # ...
  minReadySeconds: 3
status:
  # ...
```
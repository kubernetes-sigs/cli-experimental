---
title: "Update"
linkTitle: "Update"
weight: 2
type: docs
description: >
    Updating Fields
---

### Updating Fields

- Fields present in the Resource Config that are also present in the Resource will be merged recursively
  until a primitive field is updated, or a field is added / deleted.
- Fields will be updated in the Last Applied Resource Config

```yaml
# deployment.yaml (Resource Config)
apiVersion: apps/v1
kind: Deployment
metadata:
  # ...
  name: nginx-deployment
spec:
  # ...
  replicas: 2
```

```yaml
# Original Resource
kind: Deployment
metadata:
  # ...
  name: nginx-deployment
spec:
  # ...
  # could be defaulted or set by Resource Config
  replicas: 1
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
  # updated
  replicas: 2
status:
  # ...
```

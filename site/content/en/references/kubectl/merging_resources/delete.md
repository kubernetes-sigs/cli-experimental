---
title: "Delete"
linkTitle: "Delete"
weight: 3
type: docs
description: >
    Deleting Fields
---

### Deleting Fields

- Fields present in the **Last Applied Resource Config** that have been removed from the Resource Config
  will be deleted from the Resource.
- Fields set to *null* in the Resource Config that are present in the Resource Config will be deleted from the
  Resource.
- Fields will be removed from the Last Applied Resource Config


```yaml
# deployment.yaml (Resource Config)
apiVersion: apps/v1
kind: Deployment
metadata:
  # ...
  name: nginx-deployment
spec:
  # ...
```

```yaml
# Original Resource
kind: Deployment
metadata:
  # ...
  name: nginx-deployment
  # Containers replicas and minReadySeconds
  kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"apps/v1","kind":"Deployment", "spec":{"replicas": "2", "minReadySeconds": "3", ...}, "metadata": {...}}
spec:
  # ...
  minReadySeconds: 3
  replicas: 2
status:
  # ...
```

```yaml
# Applied Resource
kind: Deployment
metadata:
  # ...
  name: nginx-deployment
  kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"apps/v1","kind":"Deployment", "spec":{...}, "metadata": {...}}
spec:
  # ...
  # deleted and then defaulted, but not in Last Applied
  replicas: 1
  # minReadySeconds deleted
status:
  # ...
```

{{< alert color="success" title="Removing Fields from Resource Config" >}}
Simply removing a field from the Resource Config will *not* transfer the ownership to the cluster.
Instead it will delete the field from the Resource.  If a field is set in the Resource Config and
the user wants to give up ownership (e.g. removing `replicas` from the Resource Config and using
and autoscaler), the user must first remove it from the last Applied Resource Config stored by the
cluster.

This can be performed using `kubectl apply edit-last-applied` to delete the `replicas` field from
the **Last Applied Resource Config**, and then deleting it from the **Resource Config.**
{{< /alert >}}
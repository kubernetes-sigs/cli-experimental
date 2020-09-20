---
title: "patch"
linkTitle: "patch"
weight: 1
type: docs
description: >
    patching a resource.
---

Update field(s) of a resource using strategic merge patch, a JSON merge patch, or a JSON patch.

JSON and YAML formats are accepted.

## Command
```bash
$ kubectl patch (-f FILENAME | TYPE NAME) -p PATCH
```

## Example

### Current State
```bash
$ kubectl get deployments

NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   2/2     2            2           24m
```

### Command
```bash
kubectl patch deployment nginx -p '{"spec":{"replicas":1}}'

deployment.apps/nginx patched
```

This will reduce the number of replicas for nginx from 2 to 1.

### Output
```bash
$ kubectl get deployments

NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   1/1     1            1           26m
```
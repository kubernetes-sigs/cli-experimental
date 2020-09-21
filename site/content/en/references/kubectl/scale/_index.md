---
title: "scale"
linkTitle: "scale"
weight: 1
type: docs
description: >
    Scaling Kubenetes Resources
---

Set a new size for a Deployment, ReplicaSet, Replication Controller, or StatefulSet.

Scale also allows users to specify one or more preconditions for the scale action.

If --current-replicas or --resource-version is specified, it is validated before the scale is attempted, and it is guaranteed that the precondition holds true when the scale is sent to the server.

## Command
```bash
$ kubectl scale [--resource-version=version] [--current-replicas=count] --replicas=COUNT (-f FILENAME | TYPE NAME)
```

## Example

### Current State
```bash
$ kubectl get deployments

NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   1/1     1            1           24s
```

### Command
```bash
$ kubectl scale --replicas=3 deployment/nginx

deployment.apps/nginx scaled
```

### New State
```bash
$ kubectl get deployments

NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   3/3     3            3           87s
```
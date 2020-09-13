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

## More Examples

```bash
# Scale a replicaset named 'foo' to 3.
kubectl scale --replicas=3 rs/foo
```

```sh
# Scale a resource identified by type and name specified in "foo.yaml" to 3.
kubectl scale --replicas=3 -f foo.yaml
```

```sh
# If the deployment named mysql's current size is 2, scale mysql to 3.
kubectl scale --current-replicas=2 --replicas=3 deployment/mysql
```

```sh
# Scale multiple replication controllers.
kubectl scale --replicas=5 rc/foo rc/bar rc/baz
```

```sh
# Scale statefulset named 'web' to 3.
kubectl scale --replicas=3 statefulset/web
```

{{< alert color="success" title="Conditional Scale Update" >}}
It is possible to conditionally update the replicas if and only if the
replicas haven't changed from their last known value using the `--current-replicas` flag.
e.g. `kubectl scale --current-replicas=2 --replicas=3 deployment/mysql`
{{< /alert >}}
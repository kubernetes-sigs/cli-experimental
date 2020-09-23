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

## More Examples

```sh
# Partially update a node using a strategic merge patch. Specify the patch as JSON.
kubectl patch node k8s-node-1 -p '{"spec":{"unschedulable":true}}'
```

```sh
# Partially update a node using a strategic merge patch. Specify the patch as YAML.
kubectl patch node k8s-node-1 -p $'spec:\n unschedulable: true'
```

```sh
# Partially update a node identified by the type and name specified in "node.json" using strategic merge patch.
kubectl patch -f node.json -p '{"spec":{"unschedulable":true}}'
```

```sh
# Update a container's image; spec.containers[*].name is required because it's a merge key.
kubectl patch pod valid-pod -p '{"spec":{"containers":[{"name":"kubernetes-serve-hostname","image":"new image"}]}}'
```

```sh
# Update a container's image using a json patch with positional arrays.
kubectl patch pod valid-pod --type='json' -p='[{"op": "replace", "path": "/spec/containers/0/image", "value":"newimage"}]'
```
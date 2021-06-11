---
title: "delete"
linkTitle: "delete"
weight: 1
type: docs
description: >
    Deleting Kubernetes Resources
---

Delete resources by filenames, stdin, resources and names, or by resources and label selector.

JSON and YAML formats are accepted. Only one type of the arguments may be specified: filenames, resources and names, or resources and label selector.

## Command
```bash
$ kubectl delete ([-f FILENAME] | [-k DIRECTORY] | TYPE [(NAME | -l label | --all)])
```

## Example

### Current state
```bash
$ kubectl get deployment

NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   1/1     1            1           44s

$ kubectl get pods

NAME                     READY   STATUS    RESTARTS   AGE
nginx-6db489d4b7-9wgn9   1/1     Running   0          28s
```

### Command
```bash
$ kubectl delete deployments nginx

deployment.apps "nginx" deleted
```

### Output
```bash
$ kubectl get deployments

No resources found in default namespace.
```
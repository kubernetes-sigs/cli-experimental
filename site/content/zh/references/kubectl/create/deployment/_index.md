---
title: "deployment"
linkTitle: "deployment"
weight: 1
type: docs
description: >
    Create a deployment with the specified name.
---

You describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments.


## Command
```bash
$ kubectl create deployment NAME --image=image -- [COMMAND] [args...]
```

## Example

### Command
```bash
$ kubectl create deployment my-deployment --image=nginx
```

### Output
```bash
$ kubectl get deployments

NAME            READY   UP-TO-DATE   AVAILABLE   AGE
my-deployment   1/1     1            1           35s

$ kubectl get pods

NAME                             READY   STATUS    RESTARTS   AGE
my-deployment-7d6dd5c955-pr4jt   1/1     Running   0          15s
```



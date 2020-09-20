---
title: "run"
linkTitle: "run"
weight: 1
type: docs
description: >
    Using run command
---

Create and run a particular image in a pod. Pods are the smallest deployable units of computing that you can create and manage in Kubernetes.

{{< alert color="warning" title="Important" >}}
`run` command is deprecated
{{< /alert >}}

## Command
```bash
$ kubectl run NAME --image=image [--env="key=value"] [--port=port] [--dry-run=server|client] [--overrides=inline-json] [--command] -- [COMMAND] [args...]
```

## Example

### Command
```bash
kubectl run nginx --image=nginx
```

### Output
```bash
$ kubectl get pods

NAME                     READY   STATUS    RESTARTS   AGE
nginx-6db489d4b7-tsfhq   1/1     Running   0          28s

$ kubectl get deployment

NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   1/1     1            1           44s
```
---
title: "expose"
linkTitle: "expose"
weight: 1
type: docs
description: >
    Expose a resource as a new Kubernetes service.
---

Looks up a deployment, service, replica set, replication controller or pod by name and uses the selector for that resource as the selector for a new service on the specified port. A deployment or replica set will be exposed as a service only if its selector is convertible to a selector that service supports, i.e. when the selector contains only the matchLabels component. Note that if no port is specified via --port and the exposed resource has multiple ports, all will be re-used by the new service. Also if no labels are specified, the new service will re-use the labels from the resource it exposes.

{{< alert color="warning" title="Resources" >}}
Possible resources include (case insensitive):

- pod (po)
- service (svc)
- replicationcontroller (rc)
- deployment (deploy)
- replicaset (rs)

{{< /alert >}}

## Command
```bash
$ kubectl run NAME --image=image [--env="key=value"] [--port=port] [--dry-run=server|client] $ kubectl expose (-f FILENAME | TYPE NAME) [--port=port] [--protocol=TCP|UDP|SCTP] [--target-port=number-or-name] [--name=name] [--external-ip=external-ip-of-service] [--type=type]
```

## Example

### Current state
```bash
$ kubectl get pods

NAME                     READY   STATUS    RESTARTS   AGE
nginx-6db489d4b7-9wgn9   1/1     Running   0          28s

$ kubectl get deployment

NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   1/1     1            1           44s
```

### Command
```bash
$ kubectl expose po nginx-6db489d4b7-9wgn9 --port=80 --target-port=8000
```

### Output
```bash
$ kubectl get services

NAME                     TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
kubernetes               ClusterIP   10.96.0.1       <none>        443/TCP   8m47s
nginx-6db489d4b7-9wgn9   ClusterIP   10.106.133.77   <none>        80/TCP    2m56s
```
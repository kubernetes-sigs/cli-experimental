---
title: "port-forward"
linkTitle: "port-forward"
weight: 1
type: docs
description: >
    Forward one or more local ports to a pod.
---

Forward one or more local ports to a pod. This command requires the node to have 'socat' installed.

Use resource type/name such as deployment/mydeployment to select a pod. Resource type defaults to 'pod' if omitted.

If there are multiple pods matching the criteria, a pod will be selected automatically. The forwarding session ends when the selected pod terminates, and rerun of the command is needed to resume forwarding.

## Command
```bash
$ kubectl port-forward TYPE/NAME [options] [LOCAL_PORT:]REMOTE_PORT [...[LOCAL_PORT_N:]REMOTE_PORT_N]
```

## Forward Multiple Ports

Listen on ports 5000 and 6000 locally, forwarding data to/from ports 5000 and 6000 in the pod

```bash
kubectl port-forward pod/mypod 5000 6000
```

---

## Pod in a Workload

Listen on ports 5000 and 6000 locally, forwarding data to/from ports 5000 and 6000 in a pod selected by the
deployment

```bash
kubectl port-forward deployment/mydeployment 5000 6000
```

---

## Different Local and Remote Ports

Listen on port 8888 locally, forwarding to 5000 in the pod

```bash
kubectl port-forward pod/mypod 8888:5000
```

---

## Random Local Port

Listen on a random port locally, forwarding to 5000 in the pod

```bash
kubectl port-forward pod/mypod :5000
```

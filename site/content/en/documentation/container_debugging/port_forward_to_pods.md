---
title: "Port forward to Pods"
linkTitle: "Port forward to Pods"
weight: 4
type: docs
description: >
    Port forward to Pods
---


{{< alert color="success" title="TL;DR" >}}
- Port Forward local connections to Pods running in a cluster 
{{< /alert >}}

# Port Forward

## Motivation

Connect to ports of Pods running a cluster by port forwarding local ports.

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

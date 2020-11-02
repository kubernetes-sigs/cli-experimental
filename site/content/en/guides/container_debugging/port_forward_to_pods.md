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

{{% alert color="success" title="Operations" %}}
One can also perfrom operations such as, Port Forward to:
- Pod in a Workload
- Different Local and Remote Ports
- Random Local Port
{{% /alert %}}

{{% alert color="warning" title="Command / Examples" %}}
Check out the [reference](/references/kubectl/port-forward/) for commands and examples of `port forwarding`
{{% /alert %}}

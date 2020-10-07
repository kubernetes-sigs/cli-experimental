---
title: "Container Logs"
linkTitle: "Container Logs"
weight: 1
type: docs
description: >
    Dealing with Container Logs
---


{{< alert color="success" title="TL;DR" >}}
- Print the Logs of a Container in a cluster
{{< /alert >}}

# Summarizing Resources

## Motivation

Debugging Workloads by printing out the Logs of containers in a cluster.

## Print Logs for a Container in a Pod

Print the logs for a Pod running a single Container

```bash
kubectl logs echo-c6bc8ccff-nnj52
```

```bash
hello
hello
```

---

{{% alert color="success" title="Operations" %}}
One can also perfrom debugging operations such as:
- Print Logs for all Pods for a Workload
- Follow Logs for a Container
- Printing Logs for a Container that has exited
- Selecting a Container in a Pod 
- Printing Logs After a Time
- Printing Logs Since a Time
- Include Timestamps 

and so on.
{{% /alert %}}

{{% alert color="warning" title="Command / Examples" %}}
Check out the [reference](/cli-experimental/references/kubectl/logs/) for commands and examples.
{{% /alert %}}
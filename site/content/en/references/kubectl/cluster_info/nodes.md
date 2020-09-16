---
title: "Nodes"
linkTitle: "Nodes"
weight: 3
type: docs
description: >
    Print information about Nodes
---

## Nodes

The `kubectl top node` and `kubectl top pod` print information about the
top nodes and pods.

```bash
kubectl top node
```

```bash
  NAME                                 CPU(cores)   CPU%      MEMORY(bytes)   MEMORY%   
  gke-dev-default-pool-e1e7bf6a-cc8b   37m          1%        571Mi           10%       
  gke-dev-default-pool-e1e7bf6a-f0xh   103m         5%        1106Mi          19%       
  gke-dev-default-pool-e1e7bf6a-jfq5   139m         7%        1252Mi          22%       
  gke-dev-default-pool-e1e7bf6a-x37l   112m         5%        982Mi           17%  
```
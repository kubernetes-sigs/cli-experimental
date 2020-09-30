---
title: "custom columns"
linkTitle: "custom columns"
weight: 2
type: docs
description: >
    Print out specific fields as Columns.
---

Print out specific fields as Columns.

**Note:** Custom Columns can also be read from a file using `-o custom-columns-file`.

### Command
```bash
kubectl get deployments -o custom-columns="Name:metadata.name,Replicas:spec.replicas,Strategy:spec.strategy.type"
```

### Output
```bash
Name      Replicas   Strategy
nginx     1          RollingUpdate
```
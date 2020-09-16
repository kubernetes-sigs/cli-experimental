---
title: "Default"
linkTitle: "Default"
weight: 1
type: docs
description: >
    Printing the status in the default way
---

If no output format is specified, Get will print a default set of columns.

**Note:** Some columns *may* not directly map to fields on the Resource, but instead may
be a summary of fields.

```bash
kubectl get deployments nginx
```

```bash
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
nginx     1         1         1            0           5s
```
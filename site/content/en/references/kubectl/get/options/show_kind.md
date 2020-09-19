---
title: "Show Kind"
linkTitle: "Show Kind"
weight: 5
type: docs
description: >
    Print out the Group.Kind as part of the Name column.
---

Print out the Group.Kind as part of the Name column.

**Note:** This can be useful if the user did not specify the group in the command and
they want to know which API is being used.

### Command
```bash
kubectl get deployments --show-kind
```

### Output
```bash
NAME                          DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deployment.extensions/nginx   1         1         1            1           8m
```
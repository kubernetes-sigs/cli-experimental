---
title: "Show Labels"
linkTitle: "Show Labels"
weight: 4
type: docs
description: >
    Print out all labels on each Resource in a single column (last).
---

Print out all labels on each Resource in a single column (last).

### Command
```bash
kubectl get deployment --show-labels
```

### Output
```bash
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE       LABELS
nginx     1         1         1            1           7m        app=nginx
```

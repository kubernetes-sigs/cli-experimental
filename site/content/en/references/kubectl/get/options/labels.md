---
title: "labels"
linkTitle: "labels"
weight: 3
type: docs
description: >
    Print out specific labels each as their own columns
---

Print out specific labels each as their own columns

### Command
```bash
kubectl get deployments -L=app
```

### Output
```bash
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE       APP
nginx     1         1         1            1           8m        nginx
```
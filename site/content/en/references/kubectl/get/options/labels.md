---
title: "Labels"
linkTitle: "Labels"
weight: 3
type: docs
description: >
    Print out specific labels each as their own columns
---

Print out specific labels each as their own columns

```bash
kubectl get deployments -L=app
```

```bash
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE       APP
nginx     1         1         1            1           8m        nginx
```
---
title: "wide"
linkTitle: "wide"
weight: 1
type: docs
description: >
    Print the default columns plus some additional columns.
---

Print the default columns plus some additional columns.

**Note:** Some columns *may* not directly map to fields on the Resource, but instead may
be a summary of fields.

### Command
```bash
kubectl get -o=wide deployments nginx
```

### Output
```bash
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE       CONTAINERS   IMAGES    SELECTOR
nginx     1         1         1            1           26s       nginx        nginx     app=nginx
```

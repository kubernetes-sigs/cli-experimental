---
title: "watch"
linkTitle: "watch"
weight: 6
type: docs
description: >
    Watching for Object Changes
---

Continuously Watch and print Resources as they change

Print Resources as they are updated.

It is possible to have `kubectl get` **continuously watch for changes to objects**, and print the objects
when they are changed or when the watch is reestablished.

### Command
```bash
kubectl get deployments --watch
```

### Output
```bash
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
nginx     1         1         1            1           6h
nginx2    1         1         1            1           21m
```

{{< alert color="warning" title="Watch Timeouts" >}}
Watch **timesout after 5 minutes**, after which kubectl will re-establish the watch and print the
resources.
{{< /alert >}}

It is possible to have `kubectl get` continuously watch for changes to objects **without fetching them first**
using the `--watch-only` flag.

### Command
```bash
kubectl get deployments --watch-only
```

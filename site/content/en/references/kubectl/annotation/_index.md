---
title: "annotation"
linkTitle: "annotation"
weight: 1
type: docs
description: >
    Annotating Kubernetes Resources
---

Update the annotations on one or more resources

All Kubernetes objects support the ability to store additional data with the object as annotations. Annotations are key/value pairs that can be larger than labels and include arbitrary string values such as structured JSON. Tools and system extensions may use annotations to store their own data.

## Command
```bash
$ kubectl annotate [--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version]
```


## Example

### Current State
```bash
$ kubectl get pods
NAME                     READY   STATUS    RESTARTS   AGE
nginx-6db489d4b7-zcc8h   1/1     Running   0          5s
```

### Command
```bash
$ kubectl annotate pods nginx-6db489d4b7-zcc8h description='standard gateway'

pod/nginx-6db489d4b7-zcc8h annotated
```
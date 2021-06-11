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

## More Examples

```sh
# Update pod 'foo' with the annotation 'description' and the value 'my frontend'.
# If the same annotation is set multiple times, only the last value will be applied
kubectl annotate pods foo description='my frontend'
```

```sh
# Update a pod identified by type and name in "pod.json"
kubectl annotate -f pod.json description='my frontend'
```

```sh
# Update pod 'foo' with the annotation 'description' and the value 'my frontend running nginx', overwriting any
existing value.
kubectl annotate --overwrite pods foo description='my frontend running nginx'
```

```sh
# Update all pods in the namespace
kubectl annotate pods --all description='my frontend running nginx'
```

```sh
# Update pod 'foo' only if the resource is unchanged from version 1.
kubectl annotate pods foo description='my frontend running nginx' --resource-version=1
```

```sh
# Update pod 'foo' by removing an annotation named 'description' if it exists.
# Does not require the --overwrite flag.
kubectl annotate pods foo description-
```
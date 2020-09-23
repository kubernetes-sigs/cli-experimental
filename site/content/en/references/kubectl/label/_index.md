---
title: "label"
linkTitle: "label"
weight: 1
type: docs
description: >
    Update the labels on a resource.
---

- A label key and value must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to 63 characters each.
- Optionally, the key can begin with a DNS subdomain prefix and a single '/', like example.com/my-app
- If --overwrite is true, then existing labels can be overwritten, otherwise attempting to overwrite a label will result in an error.
- If --resource-version is specified, then updates will use this resource version, otherwise the existing resource-version will be used.

## Command
```bash
$ kubectl label [--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version]
```

## Example

### Current Status
```bash
$ kubectl get pods
NAME                     READY   STATUS    RESTARTS   AGE
nginx-6db489d4b7-b5nsn   1/1     Running   0          16m
nginx-6db489d4b7-vdhvz   1/1     Running   0          15m
```

### Command
```bash
kubectl label pods nginx-6db489d4b7-b5nsn unhealthy=true
```

### Output
```bash
$ kubectl describe pods nginx-6db489d4b7-b5nsn

Name:         nginx-6db489d4b7-b5nsn
Namespace:    default
Priority:     0
Node:         minikube/172.17.0.32
Start Time:   Sun, 20 Sep 2020 15:09:09 +0000
Labels:       pod-template-hash=6db489d4b7
              run=nginx
              unhealthy=true
Annotations:  <none>
Status:       Running
IP:           172.18.0.6

...

Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  15m   default-scheduler  Successfully assigned default/nginx-6db489d4b7-b5nsn to minikube
  Normal  Pulling    15m   kubelet, minikube  Pulling image "nginx"
  Normal  Pulled     14m   kubelet, minikube  Successfully pulled image "nginx"
  Normal  Created    14m   kubelet, minikube  Created container nginx
  Normal  Started    14m   kubelet, minikube  Started container nginx
```

Notice that the `labels` has `unhealthy=true` as a last entry.

## More Examples

```sh
# Update pod 'foo' with the label 'unhealthy' and the value 'true'.
kubectl label pods foo unhealthy=true
```

```sh
# Update pod 'foo' with the label 'status' and the value 'unhealthy', overwriting any existing value.
kubectl label --overwrite pods foo status=unhealthy
```

```sh
# Update all pods in the namespace
kubectl label pods --all status=unhealthy
```

```sh
# Update a pod identified by the type and name in "pod.json"
kubectl label -f pod.json status=unhealthy
```

```sh
# Update pod 'foo' only if the resource is unchanged from version 1.
kubectl label pods foo status=unhealthy --resource-version=1
```

```sh
# Update pod 'foo' by removing a label named 'bar' if it exists.
# Does not require the --overwrite flag.
kubectl label pods foo bar-
```
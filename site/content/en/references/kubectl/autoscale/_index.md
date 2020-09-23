---
title: "autoscale"
linkTitle: "autoscale"
weight: 1
type: docs
description: >
    Scaling Kubernetes Resources
---

Creates an autoscaler that automatically chooses and sets the number of pods that run in a kubernetes cluster.

Looks up a Deployment, ReplicaSet, StatefulSet, or ReplicationController by name and creates an autoscaler that uses the given resource as a reference. An autoscaler can automatically increase or decrease number of pods deployed within the system as needed.

## Command
```bash
$ kubectl autoscale (-f FILENAME | TYPE NAME | TYPE/NAME) [--min=MINPODS] --max=MAXPODS [--cpu-percent=CPU]
```

[OR]

```bash
$ kubectl hpa (-f FILENAME | TYPE NAME | TYPE/NAME) [--min=MINPODS] --max=MAXPODS [--cpu-percent=CPU]
```
`hpa` stands for `Horizontal Pod Autoscale`

## Example

### Current State
```bash
$ kubectl get pods

NAME                     READY   STATUS    RESTARTS   AGE
nginx-6db489d4b7-zcc8h   1/1     Running   0          5s
```

### Command
```bash
$ kubectl autoscale deployment nginx --min=2 --max=10 --cpu-percent=80

horizontalpodautoscaler.autoscaling/nginx autoscaled
```

[OR]

```bash
$ kubectl hpa deployment nginx --min=2 --max=10 --cpu-percent=80

horizontalpodautoscaler.autoscaling/nginx autoscaled
```

This will make sure to auto-scale horizontally when the CPU usage hits 80%.

### Output
```bash
$ kubectl get pods

NAME                     READY   STATUS    RESTARTS   AGE
nginx-6db489d4b7-2rrrm   1/1     Running   0          15s
nginx-6db489d4b7-vxqwm   1/1     Running   0          53s
```

Notice that the command has an arg that says `--min=2`, the deployment instantaniously auto-scales to 2 pods.
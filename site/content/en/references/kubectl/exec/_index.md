---
title: "exec"
linkTitle: "exec"
weight: 1
type: docs
description: >
    Execute a command in a container
---

Execute a command in a container

## Command
```bash
$ kubectl exec (POD | TYPE/NAME) [-c CONTAINER] [flags] -- COMMAND [args...]
```

## Example I

### Current State
```bash
$ kubectl get pods

NAME                     READY   STATUS    RESTARTS   AGE
nginx-6db489d4b7-qkd5d   1/1     Running   0          20m
```

### Command
```bash
kubectl exec nginx-6db489d4b7-qkd5d -- date
```
Notice that `date` is the command that we are executing on the pod.

### Output
```bash
Mon Sep 21 03:38:53 UTC 2020
```

## Example II

### Current State
```bash
$ kubectl get pods

NAME                     READY   STATUS    RESTARTS   AGE
nginx-6db489d4b7-qkd5d   1/1     Running   0          20m
```

### Command
```bash
kubectl exec nginx-6db489d4b7-qkd5d -- ls
```
Notice that `ls` is the command that we are executing on the pod.

### Output
```bash
bin
boot
dev
docker-entrypoint.d
docker-entrypoint.sh
etc
home
lib
lib64
media
mnt
opt
proc
root
run
sbin
srv
sys
tmp
usr
var
```
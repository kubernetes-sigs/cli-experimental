
---
title: "Executing a command in a container"
linkTitle: "Executing a command in a container"
---


{{< alert color="success" title="TL;DR" >}}
- Execute a Command in a Container
- Get a Shell in a Container
{{< /alert >}}

# Executing Commands

## Motivation

Debugging Workloads by running commands within the Container.  Commands may be a Shell with
a tty.

## Exec Command

Run a command in a Container in the cluster by specifying the **Pod name**.

```bash
kubectl exec nginx-78f5d695bd-czm8z ls
```

```bash
bin  boot  dev	etc  home  lib	lib64  media  mnt  opt	proc  root  run  sbin  srv  sys  tmp  usr  var
```

## Exec Shell

To get a Shell in a Container, use the `-t -i` options to get a tty and attach STDIN.

```bash
kubectl exec -t -i nginx-78f5d695bd-czm8z bash
```

```bash
root@nginx-78f5d695bd-czm8z:/# ls
bin  boot  dev	etc  home  lib	lib64  media  mnt  opt	proc  root  run  sbin  srv  sys  tmp  usr  var
```

{{< alert color="success" title="Specifying the Container" >}}
For Pods running multiple Containers, the Container should be specified with `-c <container-name>`.
{{< /alert >}}
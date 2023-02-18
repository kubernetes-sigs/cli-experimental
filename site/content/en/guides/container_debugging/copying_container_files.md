---
title: "Copying Container Files"
linkTitle: "Copying Container Files"
weight: 2
type: docs
description: >
    Copying Container Files
---


{{< alert color="success" title="TL;DR" >}}
- Copy files to and from Containers in a cluster
{{< /alert >}}

# Copying Container Files

## Motivation

- Copying files from Containers in a cluster to a local filesystem
- Copying files from a local filesystem to Containers in a cluster

{{< alert color="warning" title="Install Tar" >}}
Copy requires that *tar* be installed in the container image.
{{< /alert >}}


## Local to Remote

Copy a local file to a remote Pod in a cluster.

- Local file format is `<path>`
- Remote file format is `<pod-name>:<path>`


```bash
kubectl cp /tmp/foo_dir <some-pod>:/tmp/bar_dir
```


## Remote to Local

Copy a remote file from a Pod to a local file.

- Local file format is `<path>`
- Remote file format is `<pod-name>:<path>`

```bash
kubectl cp <some-pod>:/tmp/foo /tmp/bar
```

{{% alert color="success" title="Operations" %}}
One can also perform operations such as:
- Copy a specific container within a Pod running multiple containers
- Set the Pod namespace by prefixing the Pod name with `<namespace>/`.
{{% /alert %}}

{{% alert color="warning" title="Command / Examples" %}}
Check out the [reference](/references/kubectl/cp/) for commands and examples of `cp`
{{% /alert %}}

---
title: "Copying Container Files"
linkTitle: "Copying Container Files"
---
{{% pageinfo %}}
**Provide feedback at the [survey](https://www.surveymonkey.com/r/JH35X82)**
{{% /pageinfo %}}

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

## Specify the Container

Specify the Container within a Pod running multiple containers.

- `-c <container-name>`

```bash
kubectl cp /tmp/foo <some-pod>:/tmp/bar -c <specific-container>
```

## Namespaces

Set the Pod namespace by prefixing the Pod name with `<namespace>/` .

- `<pod-namespace>/<pod-name>:<path>`

```bash
kubectl cp /tmp/foo <some-namespace>/<some-pod>:/tmp/bar
```
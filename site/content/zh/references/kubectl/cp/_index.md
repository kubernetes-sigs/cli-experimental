---
title: "cp"
linkTitle: "cp"
weight: 1
type: docs
description: >
    Copy files and directories to and from containers
---

Copy files and directories to and from containers.

## Command
```bash
$ kubectl cp <file-spec-src> <file-spec-dest>
```

## Example

### Current State

`files` in local machine. Also notice a file named `simple.txt` which is about to be copied to a pod named `nginx-6db489d4b7-qkd5d`

```bash
$ ls -l

total 8
drwxr-xr-x 2 root root 4096 Mar  1  2020 Desktop
-rw-r--r-- 1 root root    6 Sep 21 03:51 simple.txt
```

`files` in a pod named `nginx-6db489d4b7-qkd5d`
 
```bash
$ kubectl exec nginx-6db489d4b7-qkd5d -- ls -al

total 84
drwxr-xr-x   1 root root 4096 Sep 21 03:17 .
drwxr-xr-x   1 root root 4096 Sep 21 03:17 ..
-rwxr-xr-x   1 root root    0 Sep 21 03:17 .dockerenv
drwxr-xr-x   2 root root 4096 Sep  8 07:00 bin
drwxr-xr-x   2 root root 4096 Jul 10 21:04 boot
drwxr-xr-x   5 root root  360 Sep 21 03:17 dev
...
drwxr-xr-x  10 root root 4096 Sep  8 07:00 usr
drwxr-xr-x   1 root root 4096 Sep  8 07:00 var
```

### Command
```bash
kubectl cp simple.txt nginx-6db489d4b7-qkd5d:.
```
Notice the `.` followed by `nginx-6db489d4b7-qkd5d:` this specifies the current directory .i.e., `root`, if you want to copy the file to another directory - you can use absolute path to paste it there.

**Example:** If you would like to copy `simple.txt` to `/temp` then use command
```bash
kubectl cp simple.txt nginx-6db489d4b7-qkd5d:./temp
```

### Output
```bash
$ kubectl exec nginx-6db489d4b7-qkd5d -- ls -al

total 84
drwxr-xr-x   1 root root 4096 Sep 21 03:17 .
drwxr-xr-x   1 root root 4096 Sep 21 03:17 ..
-rwxr-xr-x   1 root root    0 Sep 21 03:17 .dockerenv
...
-rw-r--r--   1 root root    6 Sep 21 03:54 simple.txt
...
drwxr-xr-x  10 root root 4096 Sep  8 07:00 usr
drwxr-xr-x   1 root root 4096 Sep  8 07:00 var
```

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
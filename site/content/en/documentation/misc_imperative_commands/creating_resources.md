---
title: "Creating Resources"
linkTitle: "Creating Resources"
weight: 2
type: docs
description: >
   Creating Resources
---


{{< alert color="success" title="TL;DR" >}}
- Imperatively Create a Resources
{{< /alert >}}

# Creating Resources

## Motivation

Create Resources directly from the command line for the purposes of development or debugging.
Not for production Application Management.

## Deployment

A Deployment can be created with the `create deployment` command.

```bash
kubectl create deployment my-dep --image=busybox
```

{{< alert color="success" title="Running and Attaching" >}}
It is possible to run a container and immediately attach to it using the `-i -t` flags.  e.g.
`kubectl run -t -i my-dep --image ubuntu -- bash`
{{< /alert >}}

## ConfigMap

Create a configmap based on a file, directory, or specified literal value.

A single configmap may package one or more key/value pairs.

When creating a configmap based on a file, the key will default to the basename of the file, and the value will default
to the file content.  If the basename is an invalid key, you may specify an alternate key.

When creating a configmap based on a directory, each file whose basename is a valid key in the directory will be
packaged into the configmap.  Any directory entries except regular files are ignored (e.g. subdirectories, symlinks,
devices, pipes, etc).


```bash
# Create a new configmap named my-config based on folder bar
kubectl create configmap my-config --from-file=path/to/bar
```

```bash
# Create a new configmap named my-config with specified keys instead of file basenames on disk
kubectl create configmap my-config --from-file=key1=/path/to/bar/file1.txt --from-file=key2=/path/to/bar/file2.txt
  ```

```bash
# Create a new configmap named my-config with key1=config1 and key2=config2
kubectl create configmap my-config --from-literal=key1=config1 --from-literal=key2=config2
```

```bash
# Create a new configmap named my-config from an env file
kubectl create configmap my-config --from-env-file=path/to/bar.env
```

## Secret

Create a new secret named my-secret with keys for each file in folder bar

```bash
kubectl create secret generic my-secret --from-file=path/to/bar
```

{{< alert color="success" title="Bootstrapping Config" >}}
Imperative commands can be used to bootstrap config by using `--dry-run=client -o yaml`.
`kubectl create secret generic my-secret --from-file=path/to/bar --dry-run=client -o yaml`
{{< /alert >}}

## Namespace

Create a new namespace named my-namespace

```bash
kubectl create namespace my-namespace
```

## Auth Resources

### ClusterRole

Create a ClusterRole named "foo" with API Group specified.

```bash
kubectl create clusterrole foo --verb=get,list,watch --resource=rs.extensions
```

### ClusterRoleBinding

Create a role binding to give a user cluster admin permissions.

```bash
kubectl create clusterrolebinding <choose-a-name> --clusterrole=cluster-admin --user=<your-cloud-email-account>
```

{{< alert color="success" title="Required Admin Permissions" >}}
The cluster-admin role maybe required for creating new RBAC bindings.
{{< /alert >}}

### Role

Create a Role named "foo" with API Group specified.

```bash
kubectl create role foo --verb=get,list,watch --resource=rs.extensions
```

### RoleBinding

Create a RoleBinding for user1, user2, and group1 using the admin ClusterRole.

```bash
kubectl create rolebinding admin --clusterrole=admin --user=user1 --user=user2 --group=group1
```

### ServiceAccount

Create a new service account named my-service-account

```bash
kubectl create serviceaccount my-service-account
```
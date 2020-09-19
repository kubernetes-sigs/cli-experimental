---
title: "configmap"
linkTitle: "configmap"
weight: 1
type: docs
description: >
    Creating non-confidential key-value pair data in clusters
---

A ConfigMap is an API object used to store non-confidential data in key-value pairs. Pods can consume ConfigMaps as environment variables, command-line arguments, or as configuration files in a volume.

A ConfigMap allows you to decouple environment-specific configuration from your container images, so that your applications are easily portable.

{{< alert color="warning" title="Warning" >}}
ConfigMap does not provide secrecy or encryption. If the data you want to store are confidential, use a Secret rather than a ConfigMap, or use additional (third party) tools to keep your data private.
{{< /alert >}}

# Command using File

```bash
kubectl create configmap my-config --from-file=path/to/bar
```

## Example

### Input File

```yaml
# application.properties
FOO=Bar
```

### Command

```bash
kubectl create configmap my-config --from-file=application.properties
```

### Output

```bash
$ kubectl get configmap

NAME        DATA   AGE
my-config   1      21s
```

# Command using Literal

```bash
kubectl create configmap my-config --from-literal=key1=config1 --from-literal=key2=config2
```

## Example

### Command

```bash
kubectl create configmap my-config --from-literal=FOO=Bar
```

### Output

```bash
$ kubectl get configmap

NAME        DATA   AGE
my-config   1      21s
```

# Command using env file

```bash
kubectl create configmap my-config --from-env-file=path/to/bar.env
```

## Example

### Input File

```yaml
# tracing.env
ENABLE_TRACING=true
SAMPLER_TYPE=probabilistic
SAMPLER_PARAMETERS=0.1
```

### Command

```bash
kubectl create configmap my-config --from-file=tracing.env
```

### Output

```bash
$ kubectl get configmap

NAME        DATA   AGE
my-config   1      21s
```
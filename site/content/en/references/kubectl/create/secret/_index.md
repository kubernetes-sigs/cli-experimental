---
title: "secrets"
linkTitle: "secrets"
weight: 1
type: docs
description: >
    Creating confidential information in a cluster
---

Kubernetes Secrets let you store and manage sensitive information, such as passwords, OAuth tokens, and ssh keys. Storing confidential information in a Secret is safer and more flexible than putting it verbatim in a Pod definition or in a container image.

{{< alert color="success" title="Note" >}}
Secrets can be created by using any one of the subcommands depending on use case.
- docker-registry
- generic
- tls
{{< /alert >}}

# docker-registry
- Create a secret for use with a Docker registry
```bash
kubectl create secret docker-registry NAME --docker-username=user --docker-password=password --docker-email=email [--docker-server=string] [--from-literal=key1=value1] [--dry-run=server|client|none]
```

## Example

### Command
```bash
kubectl create secret docker-registry my-secret --docker-username=kubectluser --docker-password=somepassword --docker-email=kubectl@kubectl.com --from-literal=token=GGH132YYu8asbbAA
```

### Output
```bash
$ kubectl get secrets

NAME                  TYPE                      DATA   AGE
my-secret             Opaque                    1      14s
```

# generic
- Create a secret from a local file, directory or literal value
```bash
$ kubectl create generic NAME [--type=string] [--from-file=[key=]source] [--from-literal=key1=value1] [--dry-run=server|client|none]
```

## Example

### Input File
```txt
// file-name: simplesecret.txt
kjbfkadbfkabjnaAdjna
```

### Command
```bash
kubectl create secret generic my-secret --from-file=simplesecret.txt
```

### Output
```bash
$ kubectl get secrets

NAME                  TYPE                      DATA   AGE
my-secret             Opaque                    1      14s
```

# tls
- Create a secret from tls certificate and key
```bash
$ kubectl create secret tls NAME --cert=path/to/cert/file --key=path/to/key/file [--dry-run=server|client|none]
```

## Example

### Input File
```yaml
# tls.cert
LS0tLS1CRUd...tCg==
```

```yaml
# tls.key
LS0tLS1CRUd...0tLQo=
```

### Command
```bash
kubectl create secret tls my-secret --cert=tls.cert --ket=tls.key
```

### Output
```bash
$ kubectl get secrets

NAME                  TYPE                      DATA   AGE
my-secret             Opaque                    1      14s
```




---
title: "kustomize create"
linkTitle: "create"
type: docs
weight: 4
description: >
    Create a new kustomization in the current directory
---

The `kustomize create` command will create a new kustomization in the current directory.

When run without any flags the command will create an empty `kustomization.yaml` file that can then be updated manually or with the `kustomize edit` sub-commands.

Example command:

```
kustomize create --namespace=myapp --resources=deployment.yaml,service.yaml --labels=app:myapp
```

Output from example command:

```
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
- deployment.yaml
- service.yaml
namespace: myapp
commonLabels:
  app: myapp
```

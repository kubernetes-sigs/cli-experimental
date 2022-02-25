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

### Detecting resources

> NOTE: Resource detection will not follow symlinks.

```
Flags:
      --annotations string   Add one or more common annotations.
      --autodetect           Search for kubernetes resources in the current directory to be added to the kustomization file.
  -h, --help                 help for create
      --labels string        Add one or more common labels.
      --nameprefix string    Sets the value of the namePrefix field in the kustomization file.
      --namespace string     Set the value of the namespace field in the customization file.
      --namesuffix string    Sets the value of the nameSuffix field in the kustomization file.
      --recursive            Enable recursive directory searching for resource auto-detection.
      --resources string     Name of a file containing a file to add to the kustomization file.

Global Flags:
      --stack-trace   print a stack-trace on error
```
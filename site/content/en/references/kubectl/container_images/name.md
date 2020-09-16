---
title: "Name - Image"
linkTitle: "Name - Image"
weight: 1
type: docs
description: >
    Override or set the Name for Container Images
---

## Setting a Name

The name for an image may be set by specifying `newName` and the name of the old container image.

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
images:
  - name: mycontainerregistry/myimage
    newName: differentregistry/myimage
```

## Setting a Tag

The tag for an image may be set by specifying `newTag` and the name of the container image.
```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
images:
  - name: mycontainerregistry/myimage
    newTag: v1
```
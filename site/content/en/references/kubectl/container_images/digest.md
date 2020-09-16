---
title: "Digest - Image"
linkTitle: "Digest - Image"
weight: 2
type: docs
description: >
   Setting a Digest for Container Images
---

The digest for an image may be set by specifying `digest` and the name of the container image.

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
images:
  - name: alpine
    digest: sha256:24a0c4b4a4c0eb97a1aabb8e29f18e917d05abfe1b7a7c07857230879ce7d3d3
```
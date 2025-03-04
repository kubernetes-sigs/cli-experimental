---
title: "Docker Images"
linkTitle: "Docker Images"
weight: 3
type: docs
description: >
  Install Kustomize by pulling docker images.
---

Starting with Kustomize v3.8.7, docker images are available to run Kustomize.
The images are hosted in kubernetes official repositories.

See [GGCR page] for available images.

```bash
# pull the image
docker pull registry.k8s.io/kustomize/kustomize:v5.6.0

# run 'kustomize version'
docker run registry.k8s.io/kustomize/kustomize:v5.6.0 version
```
[GGCR page]: https://explore.ggcr.dev/?repo=registry.k8s.io/kustomize/kustomize

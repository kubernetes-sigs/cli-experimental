---
title: "Docker Images"
linkTitle: "Docker Images"
weight: 3
type: docs
description: >
  Install Kustomize by pulling docker images.
---

Starting with Kustomize v3.8.7, docker images are available to run Kustomize.
The images are hosted in kubernetes official GCR repositories.

See [GCR page] for available images.

```bash
# pull the image
docker pull k8s.gcr.io/kustomize/kustomize:v5.0.0

# run 'kustomize version'
docker run k8s.gcr.io/kustomize/kustomize:v5.0.0 version
```

[GCR page]: https://us.gcr.io/k8s-artifacts-prod/kustomize/kustomize

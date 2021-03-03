---
title: "Docker 镜像"
linkTitle: "Docker 镜像"
weight: 3
type: docs
description: >
  通过拉取 docker 镜像安装 Kustomize。
---

从 Kustomize v3.8.7 开始，可以使用 docker 镜像来运行 Kustomize。
这些镜像被托管在 kubernetes 官方 GCR 仓库中。

更多可用的镜像，请参考[GCR页面]。

以下命令是如何拉取和运行 kustomize 3.8.7 docker 镜像。

```bash
docker pull k8s.gcr.io/kustomize/kustomize:v3.8.7
docker run k8s.gcr.io/kustomize/kustomize:v3.8.7 version
```

[GCR页面]: https://us.gcr.io/k8s-artifacts-prod/kustomize/kustomize

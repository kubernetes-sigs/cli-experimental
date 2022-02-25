---
title: "kustomize build"
linkTitle: "build"
type: docs
weight: 2
description: >
    Build a kustomization.yaml into the set of Kubernetes resources it describes
---

[Kustomization File reference]: https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/

`kustomize build` is Kustomize's primary command. It recursively builds (aka _hydrates_) the kustomization.yaml you point it to, resulting in a set of Kubernetes resources ready to be deployed. See the [Kustomization File reference] for details on how this works and what a kustomization.yaml can include.
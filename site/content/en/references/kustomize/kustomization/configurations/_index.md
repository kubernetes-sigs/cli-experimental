---
title: "configurations"
linkTitle: "configurations"
type: docs
weight: 6
description: >
    Customizations for transformers.
---

Each entry in this list should be a relative path to
a file that contains transformer customizations.

For example, if you would like to add additional name references, you can do so with the
following kustomization.yaml and configuration.yaml:

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
configurations:
- configuration.yaml
```

```yaml
# configuration.yaml
nameReference:
- kind: Issuer
  fieldSpecs:
    - kind: Ingress
      path: metadata/annotations/certmanager.k8s.io\/issuer
```
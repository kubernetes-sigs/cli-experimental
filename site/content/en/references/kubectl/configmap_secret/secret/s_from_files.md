---
title: "From Files"
linkTitle: "From Files"
weight: 1
type: docs
description: >
    Using Secrets from Files
---

Secret Resources may be generated much like ConfigMaps can. This includes generating them
from literals, files or environment files.

{{< alert color="success" title="Secret Syntax" >}}
Secret type is set using the `type` field.
{{< /alert >}}

**Example:** Generate a `kubernetes.io/tls` Secret from local files

**Input:** The kustomization.yaml file

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
secretGenerator:
- name: app-tls
  files:
    - "secret/tls.cert"
    - "secret/tls.key"
  type: "kubernetes.io/tls"
```

**Applied:** The Resource that is Applied to the cluster

```yaml
apiVersion: v1
kind: Secret
metadata:
  # The name has had a suffix applied
  name: app-tls-4tc9tcbd8k
type: kubernetes.io/tls
# The data has been populated from each command's output
data:
  tls.crt: LS0tLS1CRUd...tCg==
  tls.key: LS0tLS1CRUd...0tLQo=
```
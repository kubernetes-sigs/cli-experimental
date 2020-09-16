---
title: "Namespaces"
linkTitle: "Namespaces"
weight: 1
type: docs
description: >
    Setting the Namespace for all Resources
---

## Setting the Namespace for all Resources

Reference: 

The Namespace for all namespaced Resources declared in the Resource Config may be set with `namespace`.
This sets the namespace for both generated Resources (e.g. ConfigMaps and Secrets) and non-generated
Resources.

**Example:** Set the `namespace` specified in the `kustomization.yaml` on the namespaced Resources.

**Input:** The kustomization.yaml and deployment.yaml files

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: my-namespace
resources:
- deployment.yaml
```

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
```

**Applied:** The Resource that is Applied to the cluster

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx
  name: nginx-deployment
  # The namespace has been added
  namespace: my-namespace
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx
        name: nginx
```

{{< alert color="success" title="Overriding Namespaces" >}}
Setting the namespace will override the namespace on Resources if it is already set.
{{< /alert >}}
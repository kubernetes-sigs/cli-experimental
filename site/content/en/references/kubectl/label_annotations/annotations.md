---
title: "Annotations"
linkTitle: "Annotations"
weight: 1
type: docs
description: >
    Setting Annotations for all Resources
---

## Setting Annotations for all Resources

**Example:** Add the annotations declared in `commonAnnotations` to all Resources in the project.

**Input:** The kustomization.yaml and deployment.yaml files

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
commonAnnotations:
  oncallPager: 800-555-1212
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
  # Annotation added to the Deployment
  annotations:
    oncallPager: 800-555-1212
  labels:
    app: nginx
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      # Annotation also added to PodTemplate
      annotations:
        oncallPager: 800-555-1212
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx
        name: nginx
```

{{< alert color="success" title="Propagating Annotations" >}}
In addition to updating the annotations for each Resource, any fields that contain ObjectMeta
(e.g. PodTemplate) will also have the annotations added.
{{< /alert >}}
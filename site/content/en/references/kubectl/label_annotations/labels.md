---
title: "Labels"
linkTitle: "Labels"
weight: 1
type: docs
description: >
    Setting Labels for all Resources
---

## Setting Labels for all Resources

**Example:** Add the labels declared in `commonLabels` to all Resources in the project.

**Important:** Once set, commonLabels should not be changed so as not to change the Selectors for Services
or Workloads.

**Input:** The kustomization.yaml and deployment.yaml files

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
commonLabels:
  app: foo
  environment: test
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
    bar: baz
spec:
  selector:
    matchLabels:
      app: nginx
      bar: baz
  template:
    metadata:
      labels:
        app: nginx
        bar: baz
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
    app: foo # Label was changed
    environment: test # Label was added
    bar: baz # Label was ignored
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: foo # Selector was changed
      environment: test # Selector was added
      bar: baz # Selector was ignored
  template:
    metadata:
      labels:
        app: foo # Label was changed
        environment: test # Label was added
        bar: baz # Label was ignored
    spec:
      containers:
      - image: nginx
        name: nginx
```

{{< alert color="success" title="Propagating Labels to Selectors" >}}
In addition to updating the labels for each Resource, any selectors will also be updated to target the
labels.  e.g. the selectors for Services in the project will be updated to include the commonLabels
*in addition* to the other labels.

**Note:** Once set, commonLabels should not be changed so as not to change the Selectors for Services
or Workloads.
{{< /alert >}}

{{< alert color="success" title="Common Labels" >}}
The k8s.io documentation defines a set of [Common Labeling Conventions](https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/)
that may be applied to Applications.

**Note:** commonLabels should only be set for **immutable** labels, since they will be applied to Selectors.

Labeling Workload Resources makes it simpler to query Pods - e.g. for the purpose of getting their logs.
{{< /alert >}}

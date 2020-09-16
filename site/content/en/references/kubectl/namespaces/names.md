---
title: "Name prefix / suffix"
linkTitle: "Name prefix / suffix"
weight: 2
type: docs
description: >
    Setting a Name prefix or suffix for all Resources
---

## Setting a Name prefix or suffix for all Resources

A name prefix or suffix can be set for all resources using `namePrefix` or
`nameSuffix`.

**Example:** Prefix the names of all Resources.

**Input:** The kustomization.yaml and deployment.yaml files

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namePrefix: foo-
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
  # The name has been prefixed with "foo-"
  name: foo-nginx-deployment
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

{{< alert color="success" title="Propagation of the Name to Object References" >}}
Resources such as Deployments and StatefulSets may reference other Resources such as
ConfigMaps and Secrets in the Pod Spec.

This sets a name prefix or suffix for both generated Resources (e.g. ConfigMaps 
and Secrets) and non-generated Resources.

The namePrefix or nameSuffix that is applied is propagated to references to updated resources -
e.g. references to Secrets and ConfigMaps are updated with the namePrefix and nameSuffix.
{{< /alert >}}

**Example:** Prefix the names of all Resources.

This will update the ConfigMap reference in the Deployment to have the `foo` prefix.

**Input:** The kustomization.yaml and deployment.yaml files

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namePrefix: foo-
configMapGenerator:
- name: props
  literals:	
  - BAR=baz
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
        env:
        - name: BAR
          valueFrom:
            configMapKeyRef:
              name: props
              key: BAR
```

**Applied:** The Resource that is Applied to the cluster

 ```yaml
apiVersion: v1
data:
  BAR: baz
kind: ConfigMap
metadata:
  creationTimestamp: null
  name: foo-props-44kfh86dgg
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx
  name: foo-nginx-deployment
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
      - env:
        - name: BAR
          valueFrom:
            configMapKeyRef:
              key: BAR
              name: foo-props-44kfh86dgg
        image: nginx
        name: nginx
```

{{< alert color="success" title="References" >}}
Apply will propagate the `namePrefix` to any place Resources within the project are referenced by other Resources
including:

- Service references from StatefulSets
- ConfigMap references from PodSpecs
- Secret references from PodSpecs
{{< /alert >}}
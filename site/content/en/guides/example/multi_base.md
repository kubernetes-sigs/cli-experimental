---
title: "Multibase"
linkTitle: "Multibase"
weight: 1
type: docs
---

`kustomize` encourages defining multiple variants -
e.g. dev, staging and prod,
as overlays on a common base.

It's possible to create an additional overlay to
compose these variants together - just declare the
overlays as the bases of a new kustomization.

This is also a means to apply a common label or
annotation across the variants, if for some reason
the base isn't under your control. It also allows
one to define a left-most namePrefix across the
variants - something that cannot be
done by modifying the common base.

The following demonstrates this using a base
that is just a single pod.

Define a place to work:

```bash
DEMO_HOME = $(mktemp -d)
```

## `/base`
Define a common base:
```bash
$ cd $DEMO_HOME
$ mkdir base
$ cd base
```

Create a Sample Pod File and Kustomize file in `base`
```bash
$ vim kustomization.yaml
```
```yaml 
# kustomization.yaml contents
resources:
- pod.yaml
```
```yaml
# pod.yaml contents
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
  - name: nginx
    image: nginx:latest
```

## `/dev`
Define a dev variant overlaying base:

```bash
$ cd $DEMO_HOME
$ mkdir dev
$ cd dev
```

Create a Kustomize file in `dev`
```yaml
# kustomization.yaml contents
resources:
- ./../base
namePrefix: dev-
```

## `/staging`
Define a staging variant overlaying base:
```bash
$ cd $DEMO_HOME
$ mkdir staging
$ cd staging
```

Create a Kustomize file in `staging`
```yaml
# kustomization.yaml contents
resources:
- ./../base
namePrefix: stag-
```

## `/production`
Define a production variant overlaying base:
```bash
$ cd $DEMO_HOME
$ mkdir production
$ cd production
```
Create a Kustomize file in `production`
```yaml
# kustomization.yaml contents
resources:
- ./../base
namePrefix: prod-
```

## `kustomize @ root dir`
Then define a _Kustomization_ composing three variants together:
```yaml
# kustomization.yaml contents
resources:
- ./dev
- ./staging
- ./production
namePrefix: cluster-a-
```

## `directory structure`
> ```bash
> .
> ├── kustomization.yaml
> ├── base
> │   ├── kustomization.yaml
> │   └── pod.yaml
> ├── dev
> │   └── kustomization.yaml
> ├── production
> │   └── kustomization.yaml
> └── staging
>     └── kustomization.yaml
> ```

Confirm that the `kustomize build` output contains three pod objects from dev, staging and production variants.

## `output`
```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: myapp
  name: cluster-a-dev-myapp-pod
spec:
  containers:
  - image: nginx:latest
    name: nginx
---
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: myapp
  name: cluster-a-prod-myapp-pod
spec:
  containers:
  - image: nginx:latest
    name: nginx
---
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: myapp
  name: cluster-a-stag-myapp-pod
spec:
  containers:
  - image: nginx:latest
    name: nginx
```

Similarly to adding different `namePrefix` in different variants, one can also add different `namespace` and compose those variants in
one _kustomization_. For more details, take a look at:
- [Setting Namespaces](../config_management/namespaces_names.md)
- [When to Use Multiple Namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/#when-to-use-multiple-namespaces)

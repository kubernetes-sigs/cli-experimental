---
title: "Inline Patch"
linkTitle: "Inline Patch"
weight: 2
type: docs
---

[Strategic Merge Patch]: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-api-machinery/strategic-merge-patch.md
[JSON Patch]: https://tools.ietf.org/html/rfc6902

A kustomization file supports patching in three ways:
- patchesStrategicMerge: A list of patch files where each file is parsed as a [Strategic Merge Patch].
- patchesJSON6902: A list of patches and associated targetes, where each file is parsed as a [JSON Patch] and can only be applied to one target resource.
- patches: A list of patches and their associated targets. The patch can be applied to multiple objects. It auto detects whether the patch is a [Strategic Merge Patch] or [JSON Patch].

Since 3.2.0, all three support inline patch, where the patch content is put inside the kustomization file as a single string. With this feature, no separate patch files need to be created.

Make a base kustomization containing a Deployment resource.

Define a place to work:

```bash
DEMO_HOME = $(mktemp -d)
```

### `/base`
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
- deployments.yaml
```
```yaml
# deployments.yaml contents
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy
spec:
  template:
    metadata:
      labels:
        foo: bar
    spec:
      containers:
        - name: nginx
          image: nginx
          args:
          - arg1
          - arg2
```


## `PatchesStrategicMerge`

### `patch`
Create an overlay and add an inline patch in `patchesStrategicMerge` field to the kustomization file
to change the image from `nginx` to `nginx:latest`.

```bash
$ cd $DEMO_HOME
$ mkdir smp_patch
$ cd smp_patch
```

Create a Kustomize file in `smp_patch`
```yaml
# kustomization.yaml contents
resources:
- ../base

patchesStrategicMerge:
- |-
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: deploy
  spec:
    template:
      spec:
        containers:
        - name: nginx
          image: nginx:latest
```

Running `kustomize build`, in the output confirm that image is updated successfully.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy
spec:
  template:
    metadata:
      labels:
        foo: bar
    spec:
      containers:
      - args:
        - arg1
        - arg2
        image: nginx:latest
        name: nginx
```

`$patch: delete` and `$patch: replace` also work in the inline patch. Change the inline patch to delete the container `nginx`.

### `patch: delete`

```bash
$ cd $DEMO_HOME
$ mkdir smp_delete
$ cd smp_delete
```

Create a Kustomize file in `smp_delete`
```yaml
# kustomization.yaml contents
resources:
- ../base

patchesStrategicMerge:
- |-
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: deploy
  spec:
    template:
      spec:
        containers:
        - name: nginx
          $patch: delete
```

Running `kustomize build`, in the output confirm that image is updated successfully.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy
spec:
  template:
    metadata:
      labels:
        foo: bar
    spec:
      containers: []
```

### `patch: replace`

```bash
$ cd $DEMO_HOME
$ mkdir smp_replace
$ cd smp_replace
```

Create a Kustomize file in `smp_replace`
```yaml
# kustomization.yaml contents
resources:
- ../base

patchesStrategicMerge:
- |-
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: deploy
  spec:
    template:
      spec:
        containers:
        - name: nginx
          image: nginx:1.7.9
          $patch: replace
```

Running `kustomize build`, in the output confirm that image is updated successfully. Since we are replacing notice that the arguments set in the base file are gone.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy
spec:
  template:
    metadata:
      labels:
        foo: bar
    spec:
      containers:
      - image: nginx:1.7.9
        name: nginx
```

## `PatchesJson6902`

Create an overlay and add an inline patch in `patchesJSON6902` field to the kustomization file
to change the image from `nginx` to `nginx:latest`.

```bash
$ cd $DEMO_HOME
$ mkdir json
$ cd json
```

Create a Kustomize file in `json`
```yaml
# kustomization.yaml contents
resources:
- ../base

patchesJSON6902:
- target:
    group: apps
    version: v1
    kind: Deployment
    name: deploy
  patch: |-
    - op: replace
      path: /spec/template/spec/containers/0/image
      value: nginx:latest
```

Running `kustomize build`, in the output confirm that image is updated successfully.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy
spec:
  template:
    metadata:
      labels:
        foo: bar
    spec:
      containers:
      - args:
        - arg1
        - arg2
        image: nginx:latest
        name: nginx
```

## `Patches`

Create an overlay and add an inline patch in `patches` field to the kustomization file
to change the image from `nginx` to `nginx:latest`.

```bash
$ cd $DEMO_HOME
$ mkdir patch
$ cd patch
```

Create a Kustomize file in `patch`
```yaml
# kustomization.yaml contents
resources:
- ../base

patches:
- target:
    kind: Deployment
    name: deploy
  patch: |-
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: deploy
    spec:
      template:
        spec:
          containers:
          - name: nginx
            image: nginx:latest
```

Running `kustomize build`, in the output confirm that image is updated successfully.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy
spec:
  template:
    metadata:
      labels:
        foo: bar
    spec:
      containers:
      - args:
        - arg1
        - arg2
        image: nginx:latest
        name: nginx
```
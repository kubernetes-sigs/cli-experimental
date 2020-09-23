---
title: "images"
linkTitle: "images"
type: docs
description: >
    Modify the name, tags and/or digest for images.
---

Images modify the name, tags and/or digest for images without creating patches. 

One can change the `image` in the following ways (Refer the following example to know exactly how this is done):

- `postgres:8` to `my-registry/my-postgres:v1`,
- nginx tag `1.7.9` to `1.8.0`,
- image name `my-demo-app` to `my-app`,
- alpine's tag `3.7` to a digest value

## Example

### File Input

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-deployment
spec:
  template:
    spec:
      containers:
      - name: mypostgresdb
        image: postgres:8
      - name: nginxapp
        image: nginx:1.7.9
      - name: myapp
        image: my-demo-app:latest
      - name: alpine-app
        image: alpine:3.7

```

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

images:
- name: postgres
  newName: my-registry/my-postgres
  newTag: v1
- name: nginx
  newTag: 1.8.0
- name: my-demo-app
  newName: my-app
- name: alpine
  digest: sha256:24a0c4b4a4c0eb97a1aabb8e29f18e917d05abfe1b7a7c07857230879ce7d3d3

resources:
- deployment.yaml

```

### Build Output

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-deployment
spec:
  template:
    spec:
      containers:
      - image: my-registry/my-postgres:v1
        name: mypostgresdb
      - image: nginx:1.8.0
        name: nginxapp
      - image: my-app:latest
        name: myapp
      - image: alpine@sha256:24a0c4b4a4c0eb97a1aabb8e29f18e917d05abfe1b7a7c07857230879ce7d3d3
        name: alpine-app
```
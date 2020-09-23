---
title: "Container Images"
linkTitle: "Container Images"
weight: 4
type: docs
description: >
    Dealing with Application Containers
---

## Motivation

It may be useful to define the tags or digests of container images which are used across many Workloads.

Container image tags and digests are used to refer to a specific version or instance of a container
image - e.g. for the `nginx` container image you might use the tag `1.15.9` or `1.14.2`.

- Update the container image name or tag for multiple Workloads at once
- Increase visibility of the versions of container images being used within
  the project
- Set the image tag from external sources - such as environment variables
- Copy or Fork an existing Project and change the Image Tag for a container
- Change the registry used for an image

{{< alert color="success" title="Note" >}}
Check out the [References](../../../references) to learn how to override or set the Name and Tag for Container Images
{{< /alert >}}

## images

It is possible to set image tags for container images through
the `kustomization.yaml` using the `images` field.  When `images` are
specified, Apply will override the images whose image name matches `name` with a new
tag.


| Field     | Description                                                              | Example Field | Example Result |
|-----------|--------------------------------------------------------------------------|----------| --- |
| `name`    | Match images with this image name| `name: nginx`| |
| `newTag`  | Override the image **tag** or **digest** for images whose image name matches `name`    | `newTag: new` | `nginx:old` -> `nginx:new` |
| `newName` | Override the image **name** for images whose image name matches `name`   | `newName: nginx-special` | `nginx:old` -> `nginx-special:old` |


**Example:** Use `images` in the `kustomization.yaml` to update the container
images in `deployment.yaml`

Apply will set the `nginx` image to have the tag `1.8.0` - e.g. `nginx:1.8.0` and
change the image name to `nginx-special`.
This will set the name and tag for *all* images matching the *name*.

**Input:** The kustomization.yaml and deployment.yaml files

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
images:
  - name: nginx # match images with this name
    newTag: 1.8.0 # override the tag
    newName: nginx-special # override the name
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
      # The image has been changed
      - image: nginx-special:1.8.0
        name: nginx
```

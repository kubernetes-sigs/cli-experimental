---
title: "Tag - Image"
linkTitle: "Tag - Image"
weight: 3
type: docs
description: >
    Override or set the Tag for Container Images
---

## Setting a Tag

The tag for an image may be set by specifying `newTag` and the name of the container image.
```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
images:
  - name: mycontainerregistry/myimage
    newTag: v1
```

## Setting a Tag from the latest commit SHA

A common CI/CD pattern is to tag container images with the git commit SHA of source code.  e.g. if
the image name is `foo` and an image was built for the source code at commit `1bb359ccce344ca5d263cd257958ea035c978fd3`
then the container image would be `foo:1bb359ccce344ca5d263cd257958ea035c978fd3`.

A simple way to push an image that was just built without manually updating the image tags is to
download the [kustomize standalone](https://github.com/kubernetes-sigs/kustomize/) tool and run
`kustomize edit set image` command to update the tags for you.

**Example:** Set the latest git commit SHA as the image tag for `foo` images.

```bash
kustomize edit set image foo:$(git log -n 1 --pretty=format:"%H")
kubectl apply -f .
```

## Setting a Tag from an Environment Variable

It is also possible to set a Tag from an environment variable using the same technique for setting from a commit SHA.

**Example:** Set the tag for the `foo` image to the value in the environment variable `FOO_IMAGE_TAG`.

```bash
kustomize edit set image foo:$FOO_IMAGE_TAG
kubectl apply -f .
```

{{< alert color="success" title="Committing Image Tag Updates" >}}
The `kustomization.yaml` changes *may* be committed back to git so that they
can be audited.  When committing the image tag updates that have already
been pushed by a CI/CD system, be careful not to trigger new builds +
deployments for these changes.
{{< /alert >}}
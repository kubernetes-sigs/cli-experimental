---
title: "resources"
linkTitle: "resources"
type: docs
weight: 20
description: >
    Resources to include.
---

Each entry in this list must be a path to a _file_, or a path (or URL) referring to another
kustomization _directory_, e.g.

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
# Local files
- myNamespace.yaml
- deployment.yaml
- sub-dir/some-deployment.yaml

# Local directories
- ../../commonbase

# Remote URLs
- https://github.com/kubernetes-sigs/kustomize//examples/multibases/?timeout=120&ref=v3.3.1

# Legacy hashicorp/go-getter format
- github.com/kubernets-sigs/kustomize/examples/helloWorld?ref=test-branch
```

Resources will be read and processed in depth-first order.

Files should contain k8s resources in YAML form. A file may contain multiple resources separated by
the document marker `---`.  File paths should be specified _relative_ to the directory holding the
kustomization file containing the `resources` field.

Directory specification can be relative, absolute, or part of a URL.

The URL format is a HTTPS or SSH git clone URL with an optional directory and some query string
parameters. For backwards compatibility, kustomize has also supported a modified
[hashicorp/go-getter] URL format which is no longer recommended. Please refer to [remoteBuild.md]
for more information on remote targets.

[hashicorp/go-getter]: https://github.com/hashicorp/go-getter#url-format
[remoteBuild.md]: https://github.com/kubernetes-sigs/kustomize/blob/master/examples/remoteBuild.md

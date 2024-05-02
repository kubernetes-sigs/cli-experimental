---
title: "helmCharts"
linkTitle: "helmCharts"
type: docs
weight: 8
description: >
    Helm chart inflation generator.
---

[kustomize builtins]: https://kubectl.docs.kubernetes.io/references/kustomize/builtins/#_helmchartinflationgenerator_
[Helm support long term plan]: https://github.com/kubernetes-sigs/kustomize/issues/4401
[Kustomize v5.2.1]: https://github.com/kubernetes-sigs/kustomize/pull/5167#issuecomment-1771738256

## Helm Chart Inflation Generator

Kustomize has limited support for helm chart inflation through the `helmCharts` field.
You can read a detailed description of this field in the docs about [kustomize builtins].

To enable the helm chart inflation generator, you have to specify the `enable-helm` flag as follows:

```sh
kustomize build --enable-helm
```

{{% pageinfo color="info" %}}
Post [Kustomize v5.2.1], the Helm plugin supports OCI-based chart registries too. Just make sure to login to the registry first, for instance for Github Packages it is: `helm registry login https://ghcr.io`.
{{% /pageinfo %}}

## Long term support

The helm chart inflation generator in kustomize is intended to be a limited subset of helm features to help with
getting started with kustomize, and we cannot support the entire helm feature set.

### The current builtin
For enhancements to the helm chart inflation generator feature, we will only support the following changes:

- bug fixes
- critical security issues
- additional fields that are analogous to flags passed to `helm template`, except for flags such as `post-renderer`
  that allow arbitrary commands to be executed

We will not add support for:

- other large features that increase the complexity of the feature and/or have significant security implications

### Future support
The next iteration of the helm inflation generator will take the form of a KRM function, which will have
no such restrictions on what types of features we can add and support. You can see more details in
the [Helm support long term plan].

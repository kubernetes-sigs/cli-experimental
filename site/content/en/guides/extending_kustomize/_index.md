---
title: "Extending Kustomize"
linkTitle: "Extending Kustomize"
type: docs
weight: 6
description: >
    Kustomize plugins guide
---

[Everything is a Transformer]: /references/kustomize/kustomization/#everything-is-a-transformer
[Composition KEP]: https://github.com/kubernetes/enhancements/issues/2299
[Kustomize Plugin Graduation KEP]: https://github.com/kubernetes/enhancements/issues/2953
[KRM Plugin Catalog KEP]: https://github.com/kubernetes/enhancements/issues/2906
[Public KRM Functions Registry KEP]: https://github.com/kubernetes/enhancements/issues/2985

Kustomize offers a plugin framework allowing people to write their own resource _generators_
and _transformers_.

[12-factor]: https://12factor.net

* A _generator_ generates new Kubernetes objects. Examples include a helm chart
   inflator, or a plugin that emits all the
   components (deployment, service, scaler,
   ingress, etc.) needed by someone's [12-factor]
   application, based on a smaller number of free
   variables.

* A _transformer_ plugin makes changes to existing Kubernetes objects. For example, it might apply a custom order to the resource list, perform special
   container command line edits, or any other
   transformation beyond those provided by the
   builtin (`namePrefix`, `commonLabels`, etc.)
   transformers.

[generator options]: https://github.com/kubernetes-sigs/kustomize/tree/master/examples/generatorOptions.md
[transformer configs]: https://github.com/kubernetes-sigs/kustomize/tree/master/examples/transformerconfigs

Before writing a plugin, you should verify whether using advanced [generator options]
or [transformer configs] for the built-in generators and transformers would meet your needs.

## Plugin feature status

All Kustomize plugins are in **alpha**. There are five different ways to build them, some of which are slated for deprecation:
- [Containerized KRM Functions](/guides/extending_kustomize/containerized_krm_functions/)
- [Exec KRM Functions](/guides/extending_kustomize/exec_krm_functions/)
- [Legacy exec plugins (DEPRECATED)](/guides/extending_kustomize/exec_plugins/)
- [Legacy Go plugins (DEPRECATED)](/guides/extending_kustomize/go_plugins/)
- Starlark KRM Functions (DEPRECATED)

 The content on this page focuses on container-based plugins. Plugin developers getting started today are strongly encouraged to build containerized KRM Function plugins. Containers provide some level of plugin security for end users, and this method of plugin development is expected to change the least as Kustomize plugins progress towards stable status.

The vision for the future of Kustomize plugins is detailed in the [Kustomize Plugin Graduation KEP]. The [Composition KEP], [KRM Plugin Catalog KEP] and [Public KRM Functions Registry KEP] go into detail on specific aspects of this plan.


### Security

Kustomize plugins do not run in any kind of
kustomize-provided sandbox.  Thus, there is no notion
of _"plugin security"_ beyond any security inherent to the plugin runtime used. For that reason, plugin authors are strongly encouraged to containerize their plugins and to not require end users to enable disk or network access for them to execute.

End users should always carefully vet each plugin before enabling it, regardless of the plugin type. Under no circumstances should a user enable a plugin from an untrusted source.

All plugins currently require the use of additional flag: `--enable-alpha-plugins`. Without this flag, Kustomize will not load plugins and will fail with a warning about plugin use. 

The use of this flag is an opt-in acknowledging
the unstable (alpha) plugin API, the absence of
plugin provenance, and the fact that a plugin
is not part of kustomize.

To be clear, some kustomize plugin downloaded
from the internet might wonderfully transform
k8s config in a desired manner, while also
quietly doing anything the user could do to the
system running `kustomize build`.

## Plugin configuration

Like all Kustomize features, plugins are configured using [Kubernetes  objects].

```yaml
apiVersion: someteam.example.com/v1
kind: ChartInflator
metadata:
  name: notImportantHere
  annotations:
    config.kubernetes.io/function: |
      container:
        image: example.docker.com/my-functions/chart-inflator:0.1.6
spec:
  chartName: minecraft
```

[Kubernetes objects]: /references/kustomize/glossary#kubernetes-style-object

The `apiVersion` and `kind` fields are required because a kustomize plugin configuration
objects are also [Kubernetes objects].

The `metadata.annotations["config.kubernetes.io/function]` is also required, as it is used to locate the implementation of your plugin. The [KRM Plugin Catalog KEP] may change this.

The `metadata.name` field is also standard Kubernetes object metadata, and is required in most contexts.

Your plugin determines the rest of the configuration object's schema. For example,  `spec.chartName` in the example above will presumably be used by the plugin to determine the Helm chart to fetch and render into resources.
 
### Specification in `kustomization.yaml`

Plugin configuration must be referred to in the `generators` and/or `transformers`
field of your Kustomization. The items in those fields can be either paths to files containing plugin configuration, or the configuration objects themselves inlined as strings.

```yaml
generators:
- relative/path/to/some/chartInflator.yaml
- |-
  apiVersion: someteam.example.com/v1
  kind: ChartInflator
  metadata:
    name: notImportantHere
    annotations:
    config.kubernetes.io/function: |
      container:
        image: example.docker.com/my-functions/chart-inflator:0.1.6
  spec:
    chartName: minecraft

transformers:
- # same options as above
```

Given Kustomization file with the following lines:

```yaml
generators:
- chartInflator.yaml
```

The kustomization process would expect
to find a file called `chartInflator.yaml` in the
kustomization [root](/kustomize/api-reference/glossary#kustomization-root). 
The file `chartInflator.yaml` could contain:

```yaml
apiVersion: someteam.example.com/v1
kind: ChartInflator
metadata:
  name: notImportantHere
  annotations:
    config.kubernetes.io/function: |
      container:
        image: example.docker.com/my-functions/chart-inflator:0.1.6
chartName: minecraft
```

[kustomization]: /kustomize/api-reference/glossary#kustomization
[plugins]: https://github.com/kubernetes-sigs/kustomize/tree/master/plugin/builtin

For more examples of plugin configuration YAML,
browse the unit tests below the [plugins] root.

### Transforming plugin configuration

Both the `transformers` and `generators` fields can also accept paths or URLs containing Kustomization files. 

```yaml
generators:
- relative/path/to/some/kustomization
- /absolute/path/to/some/kustomization
- https://github.com/org/repo/some/kustomization

transformers:
- # same options as above
```

Paths or URLs leading to kustomizations trigger an in-process kustomization run.  Each of the
resulting objects is now further interpreted by
kustomize as a _plugin configuration_ object.

This means you can use Kustomize to manipulate plugin configuration in lower Kustomization layers and run it in overlays by referring to the lower-layer Kustomizations in the transfomers and/or generators fields. See "[Everything is a Transformer]" for more on that pattern.

The [Composition KEP] proposes a new Kind that simplifies the structure required to manipulate plugin configuration before execution.

## Execution

### Plugin orchestration

Plugins are only invoked during a run of the
`kustomize build` command, and only when a targeted Kustomization declaratively configures them.

Generator plugins are run after processing the
`resources` field (which itself can be viewed as a
generator, simply reading objects from disk).

The full set of resources is then passed into the
transformation pipeline, wherein builtin
transformations like `namePrefix` and
`commonLabel` are applied (if they were specified
in the kustomization file), followed by the
user-specified transformers in the `transformers`
field.

The order specified in the `transformers` field is respected, as transformers cannot be expected to
be commutative.

### Required alpha flags

All plugins currently require the use of an additional flag:

> `--enable-alpha-plugins`

Some plugin styles require additional flags to enable them at all, or to enable additional features. Some of these additional flags are not available in `kubectl kustomize`, effectively disabling those plugins in that version of Kustomize. The chart below outlines the flags required as of Kubectl v1.22 and Kustomize v4.4.


| Feature               | Kustomize                                                    | Kubectl Kustomize                                            |
| --------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| legacy go plugins     | `--enable-alpha-plugins`                                     | `--enable-alpha-plugins`                                     |
| legacy exec plugins   | `--enable-alpha-plugins`                                     | `--enable-alpha-plugins`                                     |
| KRM starlark plugins  | `--enable-alpha-plugins --enable-star`                       | DISABLED.                                                    |
| KRM exec plugins      | `--enable-alpha-plugins --enable-exec`                       | DISABLED.                                                    |
| KRM container plugins | `--enable-alpha-plugins` to use; `--mount` `--network` and `--network-name` further configure these types of plugins specifically. | `--enable-alpha-plugins` to use; `--mount` `--network` and `--network-name` further configure these types of plugins specifically. |

## Plugin authoring

There are five ways to build Kustomize plugins, some of which are slated for deprecation:
- [Containerized KRM Functions](/guides/extending_kustomize/containerized_krm_functions/)
- [Exec KRM Functions](/guides/extending_kustomize/exec_krm_functions/)
- [Legacy exec plugins (DEPRECATED)](/guides/extending_kustomize/exec_plugins/)
- [Legacy Go plugins (DEPRECATED)](/guides/extending_kustomize/go_plugins/)
- Starlark KRM Functions (DEPRECATED)

### Generator options

Regardless of how it is built, a plugin can adjust the generator options for the resources it emits by setting one of the following internal annotations.

> NOTE: These annotations are local to kustomize and will not be included in the final output.

**`kustomize.config.k8s.io/needs-hash`**

Resources can be marked as needing to be processed by the internal hash transformer by including the `needs-hash` annotation. When set valid values for the annotation are `"true"` and `"false"` which respectively enable or disable hash suffixing for the resource. Omitting the annotation is equivalent to setting the value `"false"`.

Hashes are determined as follows:

* For `ConfigMap` resources, hashes are based on the values of the `name`, `data`, and `binaryData` fields.
* For `Secret` resources, hashes are based on the values of the `name`, `type`, `data`, and `stringData` fields.
* For any other object type, hashes are based on the entire object content (i.e. all fields).

Example:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm-test
  annotations:
    kustomize.config.k8s.io/needs-hash: "true"
data:
  foo: bar
```

**`kustomize.config.k8s.io/behavior`**

The `behavior` annotation will influence how conflicts are handled for resources emitted by the plugin. Valid values include "create", "merge", and "replace" with "create" being the default.

Example:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm-test
  annotations:
    kustomize.config.k8s.io/behavior: "merge"
data:
  foo: bar
```

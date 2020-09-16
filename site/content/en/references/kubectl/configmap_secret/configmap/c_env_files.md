---
title: "From Environment Files"
linkTitle: "From Environment Files"
weight: 3
type: docs
description: >
    Using ConfigMaps From Environment Files
---

ConfigMap Resources may be generated from key-value pairs much the same as using the literals option
but taking the key-value pairs from an environment file. These generally end in `.env`.
To generate a ConfigMap Resource from an environment file, add an entry to `configMapGenerator` with a
single `env` entry, e.g. `env: config.env`.

{{< alert color="success" title="Environment File Syntax" >}}
- The key/value pairs inside of the environment file are separated by a `=` sign (left side is the key)
- The value of each line will appear as a data item in the ConfigMap keyed by its key.
{{< /alert >}}

**Example:** Create a ConfigMap with 3 data items generated from an environment file.

**Input:** The kustomization.yaml file

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
configMapGenerator:
- name: tracing-options
  env: tracing.env
```

```bash
# tracing.env
ENABLE_TRACING=true
SAMPLER_TYPE=probabilistic
SAMPLER_PARAMETERS=0.1
```

**Applied:** The Resource that is Applied to the cluster

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  # The name has had a suffix applied
  name: tracing-options-6bh8gkdf7k
# The data has been populated from each literal pair
data:
  ENABLE_TRACING: "true"
  SAMPLER_TYPE: "probabilistic"
  SAMPLER_PARAMETERS: "0.1"
```

{{< alert color="success" title="Overriding Base ConfigMap Values" >}}
ConfigMaps Values from Bases may be overridden by adding another generator for the ConfigMap
in the Variant and specifying the `behavior` field.  `behavior` may be
one of `create` (default value), `replace` (replace the base ConfigMap),
or `merge` (add or update the values the ConfigMap).  See [Bases and Variantions](../app_customization/bases_and_variants.md)
for more on using Bases.  e.g. `behavior: "merge"`
{{< /alert >}}
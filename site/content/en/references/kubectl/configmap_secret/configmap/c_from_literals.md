---
title: "From Literals"
linkTitle: "From Literals"
weight: 2
type: docs
description: >
    Using ConfigMaps From Literals
---

ConfigMap Resources may be generated from literal key-value pairs - such as `JAVA_HOME=/opt/java/jdk`.
To generate a ConfigMap Resource from literal key-value pairs, add an entry to `configMapGenerator` with a
list of `literals`.

{{< alert color="success" title="Literal Syntax" >}}
- The key/value are separated by a `=` sign (left side is the key)
- The value of each literal will appear as a data item in the ConfigMap keyed by its key.
{{< /alert >}}

**Example:** Create a ConfigMap with 2 data items generated from literals.

**Input:** The kustomization.yaml file

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
configMapGenerator:
- name: my-java-server-env-vars
  literals:
  - JAVA_HOME=/opt/java/jdk
  - JAVA_TOOL_OPTIONS=-agentlib:hprof
```

**Applied:** The Resource that is Applied to the cluster

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  # The name has had a suffix applied
  name: my-java-server-env-vars-k44mhd6h5f
# The data has been populated from each literal pair
data:
  JAVA_HOME: /opt/java/jdk
  JAVA_TOOL_OPTIONS: -agentlib:hprof
```
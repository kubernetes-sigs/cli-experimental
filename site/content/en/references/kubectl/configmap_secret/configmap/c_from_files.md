---
title: "From Files"
linkTitle: "From Files"
weight: 1
type: docs
description: >
    Using ConfigMaps From Files
---

ConfigMap Resources may be generated from files - such as a java `.properties` file.  To generate a ConfigMap
Resource for a file, add an entry to `configMapGenerator` with the filename.

**Example:** Generate a ConfigMap with a data item containing the contents of a file.

The ConfigMaps will have data values populated from the file contents.  The contents of each file will
appear as a single data item in the ConfigMap keyed by the filename.

**Input:** The kustomization.yaml file

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
configMapGenerator:
- name: my-application-properties
  files:
  - application.properties
```

```yaml
# application.properties
FOO=Bar
```

**Applied:** The Resource that is Applied to the cluster

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  # The name has had a suffix applied
  name: my-application-properties-c79528k426
# The data has been populated from each file's contents
data:
  application.properties: |
    FOO=Bar
```
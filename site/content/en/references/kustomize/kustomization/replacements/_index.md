---
title: "replacements"
linkTitle: "replacements"
type: docs
weight: 18
description: >
    Substitute field(s) in N target(s) with a field from a source.
---

Replacements are used to copy fields from one source into any
number of specified targets.

\
The `replacements` field can support a path to a replacement:

`kustomization.yaml`
```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

replacements:
  - path: replacement.yaml
```
`replacement.yaml`
```yaml
source:
  kind: Deployment
  fieldPath: metadata.name
targets:
  - select:
      name: my-resource
```
\
Alternatively, `replacements` supports inline replacements:

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

replacements:
- source:
    kind: Deployment
    fieldPath: metadata.name
  targets:
  - select:
      name: my-resource
```

### Syntax

The full schema of `replacements` is as follows:

```yaml
replacements:
- source:
    group: string
    version: string
    kind: string
    name: string
    namespace: string
    fieldPath: string
    options:
      delimiter: string
      index: int
      create: bool
  targets:
  - select:
      group: string
      version: string
      kind: string
      name: string
      namespace: string
    reject:
    - group: string
      version: string
      kind: string
      name: string
      namespace: string
    fieldPaths:
    - string
    options:
      delimiter: string
      index: int
      create: bool
```

### Field Descriptions

| Field       | Required| Description | Default |
| -----------: | :----:| ----------- | ---- |
| `source`| :heavy_check_mark: | The source of the value |
| `target`|:heavy_check_mark: | The N fields to write the value to |
| `group` | | The group of the referent |
| `version`|  | The version of the referent
|`kind` | |The kind of the referent
|`name` | |The name of the referent
|`namespace`|  |The namespace of the referent
|`select` |:heavy_check_mark: |Include objects that match this
|`reject`| |Exclude objects that match this
|`fieldPath`|  |The structured path to the source value | `metadata.name`
|`fieldPaths`|  |The structured path(s) to the target nodes | `metadata.name`
|`options`| |Options used to refine interpretation of the field
|`delimiter`|  |Used to split/join the field
|`index`| |Which position in the split to consider | `0`
|`create`|  |If target field is missing, add it | `false`

#### Source
The source field is a selector that determines the source of the value by finding a
match to the specified GVKNN. All the subfields of `source` are optional,
but the source selection must resolve to a single resource.

#### Targets
Replacements will be applied to all targets that are matched by the `select` field and
are NOT matched by the `reject` field, and will be applied to all listed `fieldPaths`.

##### Reject
The reject field is a selector that drops targets selected by select, overruling their selection.

For example, if we wanted to reject all Deployments named my-deploy:

```yaml
reject:
- kind: Deployment
  name: my-deploy
```

This is distinct from the following:

```yaml
reject:
- kind: Deployment
- name: my-deploy
```

The first case would only reject resources that are both of kind Deployment and named my-deploy. The second case would reject all Deployments, and all resources named my-deploy.

We can also reject more than one kind, name, etc. For example:

```yaml
reject:
- kind: Deployment
- kind: StatefulSet
```

#### Delimiter

This field is intended to be used in conjunction with the `index` field for partial string replacement.
For example, say we have a value:

`path: my/path/VALUE`

In our replacement target, we can specify something like:

```yaml
options:
  delimiter: '/'
  index: 2
```

and it would replace VALUE, e.g. `path: my/path/NEW_VALUE`.

#### Index

This field is intended to be used in conjunction with the `delimiter` field described above for partial string
replacement. The default value is 0.

If the index is out of bounds, behavior depends on whether it is in a source or target. In a source, an index out of bounds
will throw an error. For a target, a value less than 0 will cause the target to be prefixed, and a value beyond
the length of the split will cause the target to be suffixed.

If the fields `index` and `delimiter` are specified on sources or targets that are not scalar values (e.g. mapping or list values),
kustomize will throw an error.

#### Field Path format
The fieldPath and fieldPaths fields support a format of a '.'-separated path to a value. For example, the default:

`metadata.name`

You can escape the '.' one of two ways. For example, say we have the following resource:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    config.kubernetes.io/local-config: true # this is what we want to target
```

We can express our path:

1. With a '\': `metadata.annotations.config\.kubernetes\.io/local-config`

2. With '[]': `metadata.annotations.[config.kubernetes.io/local-config]`

Strings are used for mapping nodes. For sequence nodes, we support three options:

1. Index by number: `spec.template.spec.containers.1.image`

2. Index by key-value pair: `spec.template.spec.containers.[name=nginx].image`. If the key-value pair matches multiple elements in the sequence node, all matching elements will be targetted.

3. Index with a wildcard match: `spec.template.spec.containers.*.env.[name=TARGET_ENV].value`. This will target every element in the list.


### Example

For example, suppose one specifies the name of a k8s Secret object in a container's
environment variable as follows:

`job.yaml`
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: hello
spec:
  template:
    spec:
      containers:
        - image: myimage
          name: hello
          env:
            - name: SECRET_TOKEN
              value: SOME_SECRET_NAME
```

Suppose you have the following resources:

`resources.yaml`
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - image: busybox
    name: myapp-container
  restartPolicy: OnFailure
---
apiVersion: v1
kind: Secret
metadata:
  name: my-secret
```

To (1) replace the value of SOME_SECRET_NAME with the name of my-secret, and (2) to add
a restartPolicy copied from my-pod, you can do the following:

`kustomization.yaml`
```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
- resources.yaml
- job.yaml

replacements:
- path: my-replacement.yaml
- source:
    kind: Secret
    name: my-secret
  targets:
  - select:
      name: hello
      kind: Job
    fieldPaths:
    - spec.template.spec.containers.[name=hello].env.[name=SECRET_TOKEN].value
```

`my-replacement.yaml`
```yaml
source:
  kind: Pod
  name: my-pod
  fieldPath: spec.restartPolicy
targets:
- select:
    name: hello
    kind: Job
  fieldPaths:
  - spec.template.spec.restartPolicy
  options:
    create: true
```

The output of `kustomize build` will be:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: my-secret
---
apiVersion: batch/v1
kind: Job
metadata:
  name: hello
spec:
  template:
    spec:
      containers:
      - env:
        - name: SECRET_TOKEN
          value: my-secret # this value is copied from my-secret
        image: myimage
        name: hello
      restartPolicy: OnFailure # this value is copied from my-pod
---
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - image: busybox
    name: myapp-container
  restartPolicy: OnFailure
```

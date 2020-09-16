---
title: "Fields"
linkTitle: "Fields"
weight: 3
type: docs
description: >
    Format and print specific fields from Resources
---

Print the JSON representation of the first Deployment in the list on a single line.

```bash
kubectl get deployment.v1.apps -o=jsonpath='{.items[0]}{"\n"}'

```

```bash
map[apiVersion:apps/v1 kind:Deployment...replicas:1 updatedReplicas:1]]
```

---

Print the `metadata.name` field for the first Deployment in the list.

```bash
kubectl get deployment.v1.apps -o=jsonpath='{.items[0].metadata.name}{"\n"}'
```

```bash
nginx
```

---

For each Deployment, print its `metadata.name` field and a newline afterward.

```bash
kubectl get deployment.v1.apps -o=jsonpath='{range .items[*]}{.metadata.name}{"\n"}{end}'
```

```bash
nginx
nginx2
```

---

For each Deployment, print its `metadata.name` and `.status.availableReplicas`.

```bash
kubectl get deployment.v1.apps -o=jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.status.availableReplicas}{"\n"}{end}'
```
```bash
nginx	1
nginx2	1
```

---

Print the list of Deployments as single line.

```bash
kubectl get deployment.v1.apps -o=jsonpath='{@}{"\n"}'
```

```bash
map[kind:List apiVersion:v1 metadata:map[selfLink: resourceVersion:] items:[map[apiVersion:apps/v1 kind:Deployment...replicas:1 updatedReplicas:1]]]]
```

---

Print each Deployment on a new line.

```bash
kubectl get deployment.v1.apps -o=jsonpath='{range .items[*]}{@}{"\n"}{end}'
```

```bash
map[kind:Deployment...readyReplicas:1]]
map[kind:Deployment...readyReplicas:1]]
```

---

{{< alert color="success" title="Literal Syntax" >}}
On Windows, you must double quote any JSONPath template that contains spaces (not single quote as shown above for bash).
This in turn means that you must use a single quote or escaped double quote around any literals in the template.

For example:

```bash
C:\> kubectl get pods -o=jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.status.startTime}{'\n'}{end}"
```
{{< /alert >}}
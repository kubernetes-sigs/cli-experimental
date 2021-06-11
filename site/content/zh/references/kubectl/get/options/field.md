---
title: "field"
linkTitle: "field"
weight: 1
type: docs
description: >
    Printing the status of the Kubernetes resources using fields
---
## fields

Print the fields from the JSON Path

**Note:**  JSON Path can also be read from a file using `-o custom-columns-file`.

- JSON Path template is composed of JSONPath expressions enclosed by {}. In addition to the original JSONPath syntax, several capabilities are added:
- The `$` operator is optional (the expression starts from the root object by default).
- Use "" to quote text inside JSONPath expressions.
- Use range operator to iterate lists.
- Use negative slice indices to step backwards through a list. Negative indices do not “wrap around” a list. They are valid as long as -index + listLength >= 0.

### JSON Path Symbols Table

| Function	| Description	| Example	| Result |
|---|---|---|---|
| text	| the plain text	| kind is {.kind}	| kind is List |
| @	| the current object	| {@}	| the same as input |
| . or [] |	child operator	| {.kind} or {[‘kind’]}	| List |
| ..	| recursive descent	| {..name}	| 127.0.0.1 127.0.0.2 myself e2e |
| *	| wildcard. Get all objects	| {.items[*].metadata.name}	| [127.0.0.1 127.0.0.2] |
| [start:end :step]	| subscript operator	| {.users[0].name}	| myself |
| [,]	| union operator	| {.items[*][‘metadata.name’, ‘status.capacity’]}	|127.0.0.1 127.0.0.2 map[cpu:4] map[cpu:8] |
| ?()	| filter	| {.users[?(@.name==“e2e”)].user.password}	| secret |
| range, end	| iterate list	| {range .items[*]}[{.metadata.name}, {.status.capacity}] {end}	| [127.0.0.1, map[cpu:4]] [127.0.0.2, map[cpu:8]] |
| “	| quote interpreted string	| {range .items[*]}{.metadata.name}{’\t’} {end} |	127.0.0.1 127.0.0.2|

---


Print the JSON representation of the first Deployment in the list on a single line.

### Command I
```bash
kubectl get deployment.v1.apps -o=jsonpath='{.items[0]}{"\n"}'
```

### Output
```bash
map[apiVersion:apps/v1 kind:Deployment...replicas:1 updatedReplicas:1]]
```

---

Print the `metadata.name` field for the first Deployment in the list.

### Command II
```bash
kubectl get deployment.v1.apps -o=jsonpath='{.items[0].metadata.name}{"\n"}'
```
### Output
```bash
nginx
```

---

For each Deployment, print its `metadata.name` field and a newline afterward.

### Command III
```bash
kubectl get deployment.v1.apps -o=jsonpath='{range .items[*]}{.metadata.name}{"\n"}{end}'
```
### Output
```bash
nginx
nginx2
```

---

For each Deployment, print its `metadata.name` and `.status.availableReplicas`.

### Command IV
```bash
kubectl get deployment.v1.apps -o=jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.status.availableReplicas}{"\n"}{end}'
```
### Output
```bash
nginx	1
nginx2	1
```

---

Print the list of Deployments as single line.

### Command V
```bash
kubectl get deployment.v1.apps -o=jsonpath='{@}{"\n"}'
```
### Output
```bash
map[kind:List apiVersion:v1 metadata:map[selfLink: resourceVersion:] items:[map[apiVersion:apps/v1 kind:Deployment...replicas:1 updatedReplicas:1]]]]
```

---

Print each Deployment on a new line.

### Command VI
```bash
kubectl get deployment.v1.apps -o=jsonpath='{range .items[*]}{@}{"\n"}{end}'
```
### Output
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


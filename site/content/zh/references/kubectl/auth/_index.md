---
title: "auth"
linkTitle: "auth"
weight: 1
type: docs
description: >
    Inspect authorization
---

Inspect if you are authorized to perform an action on / with the Kubernetes resources.

## Command
```bash
$ kubectl auth
```

{{< alert color="success" title="Sub Commands" >}}
- can-i
- reconcile
{{< /alert >}}

## `can-i`
Check whether an action is allowed.

VERB is a logical Kubernetes API verb like 'get', 'list', 'watch', 'delete', etc. TYPE is a Kubernetes resource. Shortcuts and groups will be resolved. NONRESOURCEURL is a partial URL starts with "/". NAME is the name of a particular Kubernetes resource.

## Command
```bash
$ kubectl auth can-i VERB [TYPE | TYPE/NAME | NONRESOURCEURL]
```

## Example I

### Command
```bash
$ kubectl auth can-i create pods --all-namespaces

yes
```

Notice that the command yeilds `yes` as result - which means you are allowed to create pods on all possible namespaces avaiable.

## Example II

### Command
```bash
$ kubectl auth can-i list deployments.apps

yes
```

## `reconcile`
Reconciles rules for RBAC Role, RoleBinding, ClusterRole, and ClusterRole binding objects.

Missing objects are created, and the containing namespace is created for namespaced objects, if required.

Existing roles are updated to include the permissions in the input objects, and remove extra permissions if --remove-extra-permissions is specified.

Existing bindings are updated to include the subjects in the input objects, and remove extra subjects if --remove-extra-subjects is specified.

This is preferred to 'apply' for RBAC resources so that semantically-aware merging of rules and subjects is done.

## Command
```bash
$ kubectl auth reconcile -f FILENAME
```

## Example
`TODO`



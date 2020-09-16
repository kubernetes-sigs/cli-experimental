---
title: "APIs"
linkTitle: "APIs"
weight: 4
type: docs
description: >
    Print information about APIs
---

## APIs

The `kubectl api-versions` and `kubectl api-resources` print information
about the available Kubernetes APIs.  This information is read from the
Discovery Service.

Print the Resource Types available in the cluster.

```bash
kubectl api-resources
```

```bash
NAME                              SHORTNAMES   APIGROUP                       NAMESPACED   KIND
bindings                                                                      true         Binding
componentstatuses                 cs                                          false        ComponentStatus
configmaps                        cm                                          true         ConfigMap
endpoints                         ep                                          true         Endpoints
events                            ev                                          true         Event
limitranges                       limits                                      true         LimitRange
namespaces                        ns                                          false        Namespace
...
```

Print the API versions available in the cluster.

```bash
kubectl api-versions
```

```bash
  admissionregistration.k8s.io/v1beta1
  apiextensions.k8s.io/v1beta1
  apiregistration.k8s.io/v1
  apiregistration.k8s.io/v1beta1
  apps/v1
  apps/v1beta1
  apps/v1beta2
  ...
```

{{< alert color="success" title="Discovery" >}}
The discovery information can be viewed at `127.0.0.1:8001/` by running
`kubectl proxy`.  The Discovery for specific API can be found under either
`/api/v1` or `/apis/<group>/<version>`, depending on the API group -
e.g. `127.0.0.1:8001/apis/apps/v1`
{{< /alert >}}

The `kubectl explain` command can be used to print metadata about specific
Resource types.  This is useful for learning about the type.

```bash
kubectl explain deployment --api-version apps/v1
```

```bash
KIND:     Deployment
VERSION:  apps/v1

DESCRIPTION:
     Deployment enables declarative updates for Pods and ReplicaSets.

FIELDS:
   apiVersion	<string>
     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

   kind	<string>
     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

   metadata	<Object>
     Standard object metadata.

   spec	<Object>
     Specification of the desired behavior of the Deployment.

   status	<Object>
     Most recently observed status of the Deployment.
```

 
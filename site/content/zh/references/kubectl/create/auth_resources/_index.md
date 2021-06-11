---
title: "roles"
linkTitle: "roles"
weight: 1
type: docs
description: >
    Creating Auth Resources
---
## Auth Resources

### ClusterRole

Create a ClusterRole named "foo" with API Group specified.

```bash
kubectl create clusterrole foo --verb=get,list,watch --resource=rs.extensions
```

### ClusterRoleBinding

Create a role binding to give a user cluster admin permissions.

```bash
kubectl create clusterrolebinding <choose-a-name> --clusterrole=cluster-admin --user=<your-cloud-email-account>
```

{{< alert color="success" title="Required Admin Permissions" >}}
The cluster-admin role maybe required for creating new RBAC bindings.
{{< /alert >}}

### Role

Create a Role named "foo" with API Group specified.

```bash
kubectl create role foo --verb=get,list,watch --resource=rs.extensions
```

### RoleBinding

Create a RoleBinding for user1, user2, and group1 using the admin ClusterRole.

```bash
kubectl create rolebinding admin --clusterrole=admin --user=user1 --user=user2 --group=group1
```

### ServiceAccount

Create a new service account named my-service-account

```bash
kubectl create serviceaccount my-service-account
```
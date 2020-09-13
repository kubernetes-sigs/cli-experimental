---
title: "Editing Workloads"
linkTitle: "Editing Workloads"
weight: 4
description: >
   Editing Workloads
---


{{< alert color="success" title="TL;DR" >}}
- Edit a live Resource in an editor
{{< /alert >}}

# Editing Resources

## Motivation

Directly modify a Resource in the cluster by opening its Config in an editor.

## Edit

Edit allows a user to directly edit a Resource in a cluster rather than
editing it through a local file.

```yaml
# Edit the service named 'docker-registry':
kubectl edit svc/docker-registry
```

```yaml
# Use an alternative editor
KUBE_EDITOR="nano" kubectl edit svc/docker-registry
```

```yaml
# Edit the job 'myjob' in JSON using the v1 API format:
kubectl edit job.v1.batch/myjob -o json
```

```yaml
# Edit the deployment 'mydeployment' in YAML and save the modified config in its annotation:
kubectl edit deployment/mydeployment -o yaml --save-config
```

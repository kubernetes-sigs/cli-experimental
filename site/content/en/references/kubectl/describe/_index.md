---
title: "Describe"
linkTitle: "Describe"
weight: 8
type: docs
description: >
    Describing Kubernetes resources
---

```bash
kubectl describe deployments
```

```bash
Name:                   nginx
Namespace:              default
CreationTimestamp:      Thu, 15 Nov 2018 10:58:03 -0800
Labels:                 app=nginx
Annotations:            deployment.kubernetes.io/revision=1
Selector:               app=nginx
Replicas:               1 desired | 1 updated | 1 total | 1 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
Pod Template:
  Labels:  app=nginx
  Containers:
   nginx:
    Image:        nginx
    Port:         <none>
    Host Port:    <none>
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Progressing    True    NewReplicaSetAvailable
  Available      True    MinimumReplicasAvailable
OldReplicaSets:  <none>
NewReplicaSet:   nginx-78f5d695bd (1/1 replicas created)
Events:          <none>
```
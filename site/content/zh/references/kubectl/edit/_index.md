---
title: "edit"
linkTitle: "edit"
weight: 1
type: docs
description: >
    Edit a resource
---

Edit a resource from the default editor.

The edit command allows you to directly edit any API resource you can retrieve via the command line tools. It will open the editor defined by your KUBE_EDITOR, or EDITOR environment variables, or fall back to 'vi' for Linux or 'notepad' for Windows. You can edit multiple objects, although changes are applied one at a time. The command accepts filenames as well as command line arguments, although the files you point to must be previously saved versions of resources.

Editing is done with the API version used to fetch the resource. To edit using a specific API version, fully-qualify the resource, version, and group.

The default format is YAML. To edit in JSON, specify "-o json".

## Command
```bash
$ kubectl edit (RESOURCE/NAME | -f FILENAME)
```

## Example

### Current State

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-dev
  labels:
    app: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
```

```bash
$ kubectl get deployments

NAME        READY   UP-TO-DATE   AVAILABLE   AGE
nginx-dev   1/1     1            1           13m
```

### Command
```bash
$ kubectl edit deployment/nginx-dev
```

### Edit
```yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: nginx-dev
  namespace: default
  selfLink: /apis/apps/v1/namespaces/default/deployments/nginx-dev
  uid: 8799f7a6-e971-4285-bfac-0be1af6557d9
  resourceVersion: '2180'
  generation: 1
  creationTimestamp: '2020-09-20T14:48:35Z'
  annotations:
    deployment.kubernetes.io/revision: '1'
    kubectl.kubernetes.io/last-applied-configuration: >
      {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{},"name":"nginx-dev","namespace":"default"},"spec":{"selector":{"matchLabels":{"app":"nginx"}},"template":{"metadata":{"labels":{"app":"nginx"}},"spec":{"containers":[{"image":"nginx:1.14.2","name":"nginx"}]}}}}
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: 'nginx:1.14.2'
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          imagePullPolicy: IfNotPresent
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      dnsPolicy: ClusterFirst
      securityContext: {}
      schedulerName: default-scheduler
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 25%
      maxSurge: 25%
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
status:
  observedGeneration: 1
  replicas: 1
  updatedReplicas: 1
  readyReplicas: 1
  availableReplicas: 1
  conditions:
    - type: Available
      status: 'True'
      lastUpdateTime: '2020-09-20T14:48:43Z'
      lastTransitionTime: '2020-09-20T14:48:43Z'
      reason: MinimumReplicasAvailable
      message: Deployment has minimum availability.
    - type: Progressing
      status: 'True'
      lastUpdateTime: '2020-09-20T14:48:43Z'
      lastTransitionTime: '2020-09-20T14:48:35Z'
      reason: NewReplicaSetAvailable
      message: ReplicaSet "nginx-dev-59d7cd6545" has successfully progressed.
```

Notice that the number of replicas is changes from 1 to 2
```bash
deployment.apps/nginx-dev edited
```

### Result
```bash
NAME        READY   UP-TO-DATE   AVAILABLE   AGE
nginx-dev   2/2     2            2           105s
```

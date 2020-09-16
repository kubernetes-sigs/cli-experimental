---
title: "Raw"
linkTitle: "Raw"
weight: 2
type: docs
description: >
    Get or List Raw Resources in a cluster as Yaml or Json
---

### YAML

Print the Raw Resource formatting it as YAML.

```bash
kubectl get deployments -o yaml
```

```yaml
apiVersion: v1
items:
- apiVersion: extensions/v1beta1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "1"
    creationTimestamp: 2018-11-15T18:58:03Z
    generation: 1
    labels:
      app: nginx
    name: nginx
    namespace: default
    resourceVersion: "1672574"
    selfLink: /apis/extensions/v1beta1/namespaces/default/deployments/nginx
    uid: 6131547f-e908-11e8-9ff6-42010a8a00d1
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: nginx
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          app: nginx
      spec:
        containers:
        - image: nginx
          imagePullPolicy: Always
          name: nginx
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        terminationGracePeriodSeconds: 30
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: 2018-11-15T18:58:10Z
      lastUpdateTime: 2018-11-15T18:58:10Z
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    - lastTransitionTime: 2018-11-15T18:58:03Z
      lastUpdateTime: 2018-11-15T18:58:10Z
      message: ReplicaSet "nginx-78f5d695bd" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    observedGeneration: 1
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
```

---

### JSON

Print the Raw Resource formatting it as JSON.

```bash
kubectl get deployments -o json
```

```json
{
    "apiVersion": "v1",
    "items": [
        {
            "apiVersion": "extensions/v1beta1",
            "kind": "Deployment",
            "metadata": {
                "annotations": {
                    "deployment.kubernetes.io/revision": "1"
                },
                "creationTimestamp": "2018-11-15T18:58:03Z",
                "generation": 1,
                "labels": {
                    "app": "nginx"
                },
                "name": "nginx",
                "namespace": "default",
                "resourceVersion": "1672574",
                "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/nginx",
                "uid": "6131547f-e908-11e8-9ff6-42010a8a00d1"
            },
            "spec": {
                "progressDeadlineSeconds": 600,
                "replicas": 1,
                "revisionHistoryLimit": 10,
                "selector": {
                    "matchLabels": {
                        "app": "nginx"
                    }
                },
                "strategy": {
                    "rollingUpdate": {
                        "maxSurge": "25%",
                        "maxUnavailable": "25%"
                    },
                    "type": "RollingUpdate"
                },
                "template": {
                    "metadata": {
                        "creationTimestamp": null,
                        "labels": {
                            "app": "nginx"
                        }
                    },
                    "spec": {
                        "containers": [
                            {
                                "image": "nginx",
                                "imagePullPolicy": "Always",
                                "name": "nginx",
                                "resources": {},
                                "terminationMessagePath": "/dev/termination-log",
                                "terminationMessagePolicy": "File"
                            }
                        ],
                        "dnsPolicy": "ClusterFirst",
                        "restartPolicy": "Always",
                        "schedulerName": "default-scheduler",
                        "securityContext": {},
                        "terminationGracePeriodSeconds": 30
                    }
                }
            },
            "status": {
                "availableReplicas": 1,
                "conditions": [
                    {
                        "lastTransitionTime": "2018-11-15T18:58:10Z",
                        "lastUpdateTime": "2018-11-15T18:58:10Z",
                        "message": "Deployment has minimum availability.",
                        "reason": "MinimumReplicasAvailable",
                        "status": "True",
                        "type": "Available"
                    },
                    {
                        "lastTransitionTime": "2018-11-15T18:58:03Z",
                        "lastUpdateTime": "2018-11-15T18:58:10Z",
                        "message": "ReplicaSet \"nginx-78f5d695bd\" has successfully progressed.",
                        "reason": "NewReplicaSetAvailable",
                        "status": "True",
                        "type": "Progressing"
                    }
                ],
                "observedGeneration": 1,
                "readyReplicas": 1,
                "replicas": 1,
                "updatedReplicas": 1
            }
        }
    ],
    "kind": "List",
    "metadata": {
        "resourceVersion": "",
        "selfLink": ""
    }
}
```
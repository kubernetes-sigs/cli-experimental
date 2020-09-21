---
title: "diff"
linkTitle: "diff"
weight: 1
type: docs
description: >
    diffs the online configuration with local config
---

Diff configurations specified by filename or stdin between the current online configuration, and the configuration as it would be if applied.

Output is always YAML.

KUBECTL_EXTERNAL_DIFF environment variable can be used to select your own diff command. By default, the "diff" command available in your path will be run with "-u" (unified diff) and "-N" (treat absent files as empty) options.

{{< alert color="warning" title="Exit status" >}}
- `0` No differences were found. 
- `1` Differences were found. 
- `>1` Kubectl or diff failed with an error.
{{< /alert >}}

## Command
```bash
$ kubectl diff -f FILENAME
```

## Example

### Input File
```yaml
# deployment.yaml - online configuration
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

```yaml
# deployment.yaml - local configuration
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
        image: nginx:latest
```

Notice that the local configuration refers to the latest nginx container from the registry.

### Command
```bash
kubectl diff -f deployment.yaml
```

### Output

```bash
diff -u -N /tmp/LIVE-435797985/apps.v1.Deployment.default.nginx-dev /tmp/MERGED-822429644/apps.v1.Deployment.default.nginx-dev
--- /tmp/LIVE-435797985/apps.v1.Deployment.default.nginx-dev    2020-09-20 14:50:30.160820677 +0000
+++ /tmp/MERGED-822429644/apps.v1.Deployment.default.nginx-dev  2020-09-20 14:50:30.172820784 +0000
@@ -6,7 +6,7 @@
     kubectl.kubernetes.io/last-applied-configuration: |
       {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{},"name":"nginx-dev","namespace":"default"},"spec":{"selector":{"matchLabels":{"app":"nginx"}},"template":{"metadata":{"labels":{"app":"nginx"}},"spec":{"containers":[{"image":"nginx:1.14.2","name":"nginx"}]}}}}
   creationTimestamp: "2020-09-20T14:48:35Z"
-  generation: 1
+  generation: 2
   name: nginx-dev
   namespace: default
   resourceVersion: "2180"
@@ -31,7 +31,7 @@
         app: nginx
     spec:
       containers:
-      - image: nginx:1.14.2
+      - image: nginx:latest
         imagePullPolicy: IfNotPresent
         name: nginx
         resources: {}
exit status 1
```

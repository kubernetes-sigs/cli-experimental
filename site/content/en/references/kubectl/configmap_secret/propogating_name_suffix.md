---
title: "Propagating the Name Suffix"
linkTitle: "Propagating the Name Suffix"
weight: 5
type: docs
description: >
    Letting ConfigMap or Secret know the name of the generated Resource name suffix
---

Workloads that reference the ConfigMap or Secret will need to know the name of the generated Resource
including the suffix, however Apply takes care of this automatically for users.  Apply will identify
references to generated ConfigMaps and Secrets, and update them.

The generated ConfigMap name will be `my-java-server-env-vars` with a suffix unique to its contents.
Changes to the contents will change the name suffix, resulting in the creation of a new ConfigMap,
and transform Workloads to point to this one.

The PodTemplate volume references the ConfigMap by the name specified in the generator (excluding the suffix).
Apply will update the name to include the suffix applied to the ConfigMap name.

**Input:** The kustomization.yaml and deployment.yaml files

```yaml
# kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
configMapGenerator:
- name: my-java-server-env-vars
  literals:
  - JAVA_HOME=/opt/java/jdk
  - JAVA_TOOL_OPTIONS=-agentlib:hprof
resources:
- deployment.yaml
```

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-deployment
  labels:
    app: test
spec:
  selector:
    matchLabels:
      app: test
  template:
    metadata:
      labels:
        app: test
    spec:
      containers:
      - name: container
        image: k8s.gcr.io/busybox
        command: [ "/bin/sh", "-c", "ls /etc/config/" ]
        volumeMounts:
        - name: config-volume
          mountPath: /etc/config
      volumes:
      - name: config-volume
        configMap:
          name: my-java-server-env-vars
```

**Applied:** The Resources that are Applied to the cluster.

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  # The name has been updated to include the suffix
  name: my-java-server-env-vars-k44mhd6h5f
data:
  JAVA_HOME: /opt/java/jdk
  JAVA_TOOL_OPTIONS: -agentlib:hprof
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: test
  name: test-deployment
spec:
  selector:
    matchLabels:
      app: test
  template:
    metadata:
      labels:
        app: test
    spec:
      containers:
      - command:
        - /bin/sh
        - -c
        - ls /etc/config/
        image: k8s.gcr.io/busybox
        name: container
        volumeMounts:
        - mountPath: /etc/config
          name: config-volume
      volumes:
      - configMap:
          # The name has been updated to include the
          # suffix matching the ConfigMap
          name: my-java-server-env-vars-k44mhd6h5f
        name: config-volume
```
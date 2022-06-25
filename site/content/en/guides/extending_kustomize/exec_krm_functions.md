---
title: "Exec KRM functions"
linkTitle: "Exec KRM functions"
type: docs
description: >
    Guide to writing exec KRM functions for use as Kustomize plugins
weight: 2
---

{{% alert color="warning" title="Alpha status warning" %}}
This style of plugin is in alpha and may change significantly.
See the [Kustomize Plugin Graduation KEP](https://github.com/kubernetes/enhancements/issues/2953) for information on the future of Kustomize plugins.
{{% /alert %}}

## Authoring Exec KRM Functions

[KRM Functions Specification]: https://github.com/kubernetes-sigs/kustomize/blob/master/cmd/config/docs/api-conventions/functions-spec.md

An exec KRM Function is any executable written in any language that accepts a ResourceList as input on stdin and emits a ResourceList as output on stdout, in accordance with the [KRM Functions Specification].

### Placement

Exec functions have no built-in security model. Kustomize does not provide an execution sandbox of any kind. Accordingly, plugins intended for distribution should be built as [Containerized KRM Functions](/guides/extending_kustomize/exec_krm_functions/) whenever possible. Exec functions are generally expected to be bespoke to a particular Kustomization and embedded within its directory.

Exec plugins are referenced directly by the metadata of the KRM object used to configure them. The path must be a relative to the Kustomization that imports the plugin config.
For example, given the following directory structure:

```bash
.
├── cmGenerator.yaml # plugin config
├── kustomization.yaml
├── plugins
│   ├── cm_generator.rb 
│   ├── cm_generator_test.rb
│   └── testdata
│       └── ...
└── service.yaml
```

The plugin configuration might look like:

```yaml
apiVersion: myDevOpsTeam
kind: ConfigMapGenerator
metadata:
  name: whatever
  annotations:
    # path is relative to kustomization.yaml
    config.kubernetes.io/function: |
      exec:
        path: ./plugins/cm_generator.rb 
spec:
  altGreeting: Bienvenue
  enableRisky: true
```

## Guided example

This is a (no reading allowed!) 60 second copy/paste guided
example.

This demo writes and uses a somewhat ridiculous
_exec_ plugin (written in bash) that generates a
`ConfigMap`.

Prerequisites:

* linux or osx
* curl
* bash
* Go 1.16

### Make a place to work

```bash
DEMO=$(mktemp -d)
```

### Install kustomize

Per the [instructions](/installation/kustomize/):

```bash
curl -s "https://raw.githubusercontent.com/\
kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash
mkdir -p $DEMO/bin
mv kustomize $DEMO/bin
```

### Create a kustomization

Make a kustomization directory to
hold all your config:

```bash
MYAPP=$DEMO/myapp
mkdir -p $MYAPP
```

Make a deployment config:

```bash
# $MYAPP/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-deployment
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: the-container
        image: monopole/hello:1
        command: ["/hello",
                  "--port=8080",
                  "--enableRiskyFeature=$(ENABLE_RISKY)"]
        ports:
        - containerPort: 8080
        env:
        - name: ALT_GREETING
          valueFrom:
            configMapKeyRef:
              name: the-map
              key: altGreeting
        - name: ENABLE_RISKY
          valueFrom:
            configMapKeyRef:
              name: the-map
              key: enableRisky
```

Make a service config:

```yaml
# $MYAPP/service.yaml
kind: Service
apiVersion: v1
metadata:
  name: the-service
spec:
  type: LoadBalancer
  ports:
  - protocol: TCP
    port: 8666
    targetPort: 8080
```

Now make a config file for the plugin
you're about to write.

This config file is just another Kubernetes resource
object.  The `config.kubernetes.io/function` annotation is used to _find_ the plugin code on your
filesystem (more on this later).

```yaml
# $MYAPP/cmGenerator.yaml
apiVersion: myDevOpsTeam
kind: SillyConfigMapGenerator
metadata:
  name: whatever
  annotations:
    config.kubernetes.io/function: |
      exec: 
        path: ./plugins/silly-generator.sh
spec:
  altGreeting: Bienvenue
  enableRisky: true
```

Finally, make a kustomization file
referencing all of the above:

```yaml
# $MYAPP/kustomization.yaml
commonLabels:
  app: hello
resources:
- deployment.yaml
- service.yaml
generators:
- cmGenerator.yaml
```

Review the files

```bash
$ ls -C1 $MYAPP
cmGenerator.yaml
deployment.yaml
kustomization.yaml
service.yaml
```

### Make a home for plugins

Due to their inherent lack of security, exec plugins should be bespoke to and distributed along with the Kustomization they are for. Plugins intended for reuse should be containerized.

The plugin config defined above in
`$MYAPP/cmGenerator.yaml` specifies:

```yaml
metadata:
  annotations:
    config.kubernetes.io/function: |
      exec: 
        path: ./plugins/silly-generator.sh
```

So let's create that directory:

```bash
mkdir $MYAPP/plugins
```

A plugin gets its own directory to hold itself,
its tests and any supplemental data files it
might need.

### Create the plugin

```bash
# $MYAPP/plugins/silly-generator.sh
#!/bin/bash
resourceList=$(cat) # read the `kind: ResourceList` from stdin
altGreeting=$(echo "$resourceList" | yq e '.functionConfig.spec.altGreeting' - )
enableRisky=$(echo "$resourceList" | yq e '.functionConfig.spec.enableRisky' - )

echo "
kind: ResourceList
items:
- kind: ConfigMap
  apiVersion: v1
  metadata:
    name: the-map
  data:
    altGreeting: "$altGreeting"
    enableRisky: "$enableRisky"
"
```

By definition, an _exec_ plugin must be executable:

```bash
chmod a+x $MYAPP/plugins/silly-generator.sh
```

### Review the layout

```bash
$ tree $DEMO
tmp/tmp.qDYh1kiHqD
└── myapp
    ├── cmGenerator.yaml
    ├── deployment.yaml
    ├── kustomization.yaml
    ├── plugins
    │   └── silly-generator.sh
    └── service.yaml

2 directories, 5 files
```

### Build your app, using the plugin

```bash
$DEMO/bin/kustomize build --enable-alpha-plugins --enable-exec $MYAPP
```

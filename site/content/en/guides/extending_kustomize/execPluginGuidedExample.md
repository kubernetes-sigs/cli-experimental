---
title: "Exec plugin on linux"
linkTitle: "Exec plugin on linux"
type: docs
description: >
    Exec plugin on linux in 60 seconds
---

This is a (no reading allowed!) 60 second copy/paste guided
example.  Full plugin docs [here](..).

This demo writes and uses a somewhat ridiculous
_exec_ plugin (written in bash) that generates a
`ConfigMap`.

This is a guide to try it without damaging your
current setup.

#### requirements

* linux, git, curl, Go 1.13

## Make a place to work

```bash
DEMO=$(mktemp -d)
```

## Create a kustomization

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
                  "--date=$(THE_DATE)",
                  "--enableRiskyFeature=$(ENABLE_RISKY)"]
        ports:
        - containerPort: 8080
        env:
        - name: THE_DATE
          valueFrom:
            configMapKeyRef:
              name: the-map
              key: today
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
EOF
```

Make a service config:

```bash
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
EOF
```

Now make a config file for the plugin
you're about to write.

This config file is just another k8s resource
object.  The values of its `apiVersion` and `kind`
fields are used to _find_ the plugin code on your
filesystem (more on this later).

```bash
# $MYAPP/cmGenerator.yaml
apiVersion: myDevOpsTeam
kind: SillyConfigMapGenerator
metadata:
  name: whatever
argsOneLiner: Bienvenue true
EOF
```

Finally, make a kustomization file
referencing all of the above:

```bash
# $MYAPP/kustomization.yaml
commonLabels:
  app: hello
resources:
- deployment.yaml
- service.yaml
generators:
- cmGenerator.yaml
EOF
```

Review the files

```bash
ls -C1 $MYAPP
```

## Make a home for plugins

By default plugins are searched for in the following locations:

- `$HOME/kustomize/plugin`
- `$XDG_CONFIG_HOME/kustomize/plugin`
- `$KUSTOMIZE_PLUGIN_HOME`

This demo will use the ephemeral directory:

```bash
KUSTOMIZE_PLUGIN_ROOT=$DEMO/plugin
```

The plugin config defined above in
`$MYAPP/cmGenerator.yaml` specifies:

> ```bash
> apiVersion: myDevOpsTeam
> kind: SillyConfigMapGenerator
> ```

This means the plugin must live in a directory
named:

```bash
MY_PLUGIN_DIR=$KUSTOMIZE_PLUGIN_ROOT/myDevOpsTeam/sillyconfigmapgenerator

mkdir -p $MY_PLUGIN_DIR
```

The directory name is the plugin config's
_apiVersion_ followed by its lower-cased _kind_.

A plugin gets its own directory to hold itself,
its tests and any supplemental data files it
might need.

## Create the plugin

There are two kinds of plugins, _exec_ and _Go_.

Make an _exec_ plugin, installing it to the
correct directory and file name.  The file name
must match the plugin's _kind_ (in this case,
`SillyConfigMapGenerator`):

```bash
# $MY_PLUGIN_DIR/SillyConfigMapGenerator
#!/usr/bin/env bash
# Skip the config file name argument.
shift
today=`date +%F`
echo "
kind: ConfigMap
apiVersion: v1
metadata:
  name: the-map
data:
  today: $today
  altGreeting: "$1"
  enableRisky: "$2"
"
EOF
```

By definition, an _exec_ plugin must be executable:

```bash
chmod a+x $MY_PLUGIN_DIR/SillyConfigMapGenerator
```

## Install kustomize

Per the [instructions](/installation/kustomize/):

```bash
curl -s "https://raw.githubusercontent.com/\
kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash
mkdir -p $DEMO/bin
mv kustomize $DEMO/bin
```

## Review the layout

```bash
tree $DEMO
```

## Build your app, using the plugin

```bash
export KUSTOMIZE_PLUGIN_ROOT
$DEMO/bin/kustomize build --enable-alpha-plugins $MYAPP
```

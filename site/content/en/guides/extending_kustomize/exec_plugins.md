---
title: "Exec plugins (deprecated)"
linkTitle: "Exec plugins"
type: docs
description: >
    Guide to writing exec plugins
weight: 3
---

{{% alert color="warning" title="Deprecation warning" %}}
This style of plugin is slated for deprecation.
See the [Kustomize Plugin Graduation KEP](https://github.com/kubernetes/enhancements/issues/2953) for information on the future of Kustomize plugins.
{{% /alert %}}

## Authoring legacy exec plugins

An _exec plugin_ is any executable that accepts a
single argument on its command line - the name of
a YAML file containing its configuration (the file name
provided in the kustomization file).

[helm chart inflator]: https://github.com/kubernetes-sigs/kustomize/tree/master/plugin/someteam.example.com/v1/chartinflator
[bashed config map]: https://github.com/kubernetes-sigs/kustomize/tree/master/plugin/someteam.example.com/v1/bashedconfigmap
[sed transformer]: https://github.com/kubernetes-sigs/kustomize/tree/master/plugin/someteam.example.com/v1/sedtransformer
[hashicorp go-getter]: https://github.com/kubernetes-sigs/kustomize/tree/master/plugin/someteam.example.com/v1/gogetter

### Placement

Each plugin gets its own dedicated directory named

[`XDG_CONFIG_HOME`]: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

```bash
$XDG_CONFIG_HOME/kustomize/plugin
    /${apiVersion}/LOWERCASE(${kind})
```

The default value of [`XDG_CONFIG_HOME`] is
`$HOME/.config`.

The one-plugin-per-directory requirement eases
creation of a plugin bundle (source, tests, plugin
data files, etc.) for sharing.

When loading, kustomize will look for an
_executable_ file called `kind`.

```bash
$XDG_CONFIG_HOME/kustomize/plugin
    /${apiVersion}/LOWERCASE(${kind})/${kind}
```

Failure to find a plugin to load fails the overall
`kustomize build`.

### Examples

* [helm chart inflator] - A generator that inflates a helm chart.
* [bashed config map] -  Super simple configMap generation from bash.
* [sed transformer] - Define your unstructured edits using a
   plugin like this one.
* [hashicorp go-getter] - Download kustomize layes and build it to generate resources

A generator plugin accepts nothing on `stdin`, but emits
generated resources to `stdout`.

A transformer plugin accepts resource YAML on `stdin`,
and emits those resources, presumably transformed, to
`stdout`.

kustomize uses an exec plugin adapter to provide
marshalled resources on `stdin` and capture
`stdout` for further processing.

## Guided example

This is a (no reading allowed!) 60 second copy/paste guided
example.

This demo writes and uses a somewhat ridiculous
_exec_ plugin (written in bash) that generates a
`ConfigMap`.

Prerequisites:
* linux
* git
* curl
* Go 1.13

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

### Make a home for plugins

Plugins must live in a particular place for
kustomize to find them.

This demo will use the ephemeral directory:

```bash
PLUGIN_ROOT=$DEMO/kustomize/plugin
```

The plugin config defined above in
`$MYAPP/cmGenerator.yaml` specifies:

```bash
apiVersion: myDevOpsTeam
kind: SillyConfigMapGenerator
```

This means the plugin must live in a directory
named:

```bash
MY_PLUGIN_DIR=$PLUGIN_ROOT/myDevOpsTeam/sillyconfigmapgenerator

mkdir -p $MY_PLUGIN_DIR
```

The directory name is the plugin config's
_apiVersion_ followed by its lower-cased _kind_.

A plugin gets its own directory to hold itself,
its tests and any supplemental data files it
might need.

### Create the plugin

Make an _exec_ plugin, installing it to the
correct directory and file name.  The file name
must match the plugin's _kind_ (in this case,
`SillyConfigMapGenerator`):

```bash
# $MY_PLUGIN_DIR/SillyConfigMapGenerator
#!/bin/bash
# Skip the config file name argument.
shift
echo "
kind: ConfigMap
apiVersion: v1
metadata:
  name: the-map
data:
  altGreeting: "$1"
  enableRisky: "$2"
"
EOF
```

By definition, an _exec_ plugin must be executable:

```bash
chmod a+x $MY_PLUGIN_DIR/SillyConfigMapGenerator
```

### Review the layout

```bash
tree $DEMO
```

### Build your app

```bash
XDG_CONFIG_HOME=$DEMO $DEMO/bin/kustomize build --enable_alpha_plugins $MYAPP
```

Above, if you had set

```bash
PLUGIN_ROOT=$HOME/.config/kustomize/plugin
```

there would be no need to use `XDG_CONFIG_HOME` in the
_kustomize_ command above.

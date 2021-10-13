---
title: "Containerized KRM Functions"
linkTitle: "Containerized KRM Functions"
type: docs
description: >
    Guide to writing containerized KRM functions for use as Kustomize plugins 
weight: 1
---

{{% alert color="info" title="Alpha status warning" %}}
This is the recommended style of plugin, but you should be aware that it is still in alpha.
See the [Kustomize Plugin Graduation KEP](https://github.com/kubernetes/enhancements/issues/2953) for information on the future of Kustomize plugins.
{{% /alert %}}


## Authoring Containerized KRM Functions

[KRM Functions Specification]: https://github.com/kubernetes-sigs/kustomize/blob/master/cmd/config/docs/api-conventions/functions-spec.md

A containerized KRM Function is any container whose entrypoint accepts a ResourceList as input on stdin and emits a ResourceList as output on stdout, in accordance with the [KRM Functions Specification].

### Configuration

Containerized KRM Function plugins are referenced directly by the metadata of the KRM object used to configure them. 

For example, the plugin configuration might look like:

```yaml
apiVersion: someteam.example.com/v1
kind: ChartInflator
metadata:
  name: notImportantHere
  annotations:
    config.kubernetes.io/function: |
      image: example.docker.com/my-functions/chart-inflator:0.1.6
spec:
  chartName: minecraft
```

## Guided example

This is a (no reading allowed!) 60 second copy/paste guided
example.

This demo writes and uses a somewhat ridiculous
_containerized_ plugin (written in Go) that follows the KRM Function Specification and generates a
`ConfigMap`.

Prerequisites:

* linux or osx
* curl
* bash
* docker
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
object.  The `config.kubernetes.io/function` annotation is used to _find_ the docker container that implements the ValueAnnotator type.

```yaml
# $MYAPP/annotator.yaml
apiVersion: transformers.example.co/v1
kind: ValueAnnotator
metadata:
  name: notImportantHere
  annotations:
    config.kubernetes.io/function: |
      image: example.docker.com/my-functions/valueannotator:1.0.0
value: 'important-data'
```

Finally, make a kustomization file
referencing all of the above:

```yaml
# $MYAPP/kustomization.yaml
resources:
- service.yaml
transformers:
- annotator.yaml
```

Review the files

```bash
ls -C1 $MYAPP
```

### Make a home for plugins

Since Kustomize accesses your code indirectly through the container, you can develop containerized plugins anywhere you'd like within your Go path.

Let's create a new directory for our code:

```bash
mkdir $GOPATH/src/kustomize-plugin-demo
```

### Create the plugin

First initialize the plugin directory with go mod:

```bash
cd $GOPATH/src/kustomize-plugin-demo
```

Create the `main.go` with the demo code:

```go
// $GOPATH/src/kustomize-plugin-demo/main.go
package main

import (
  "os"

  "sigs.k8s.io/kustomize/kyaml/fn/framework"
  "sigs.k8s.io/kustomize/kyaml/fn/framework/command"
  "sigs.k8s.io/kustomize/kyaml/kio"
  "sigs.k8s.io/kustomize/kyaml/yaml"
)

type ValueAnnotator struct {
  Value string `yaml:"value" json:"value"`
}

func main() {
  config := new(ValueAnnotator)
  fn := func(items []*yaml.RNode) ([]*yaml.RNode, error) {
    for i := range items {
      err := items[i].PipeE(yaml.SetAnnotation("custom.io/the-value", config.Value))
      if err != nil {
        return nil, err
      }
    }
    return items, nil
  }
  p := framework.SimpleProcessor{Config: config, Filter: kio.FilterFunc(fn)}
  cmd := command.Build(p, command.StandaloneDisabled, false)
  command.AddGenerateDockerfile(cmd)
  if err := cmd.Execute(); err != nil {
    os.Exit(1)
  }
}
```

Create the go module

```bash
go mod init
go mod tidy
```

### Generate the dockerfile

Because our program uses `command.AddGenerateDockerfile(cmd)`, go can generate a default Dockerfile for us:

```bash
go run main.go gen .
```

### Build the container

```bash
docker build . -t example.docker.com/my-functions/valueannotator:1.0.0
```

Note that since this is a quick local-only demo, we do not need to push the image, and the tag doesn't need to refer to a real docker registry.

### Build your app

```bash
$DEMO/bin/kustomize build --enable-alpha-plugins $MYAPP
```

You should see the "important-data" added to the Service's annotations.

### Clean up

```bash
rm -r $GOPATH/src/kustomize-plugin-demo
```

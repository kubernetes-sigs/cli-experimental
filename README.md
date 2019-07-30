[![Build Status](https://travis-ci.org/kubernetes-sigs/cli-experimental.svg?branch=master)](https://travis-ci.org/kubernetes-sigs/cli-experimental "Travis")
[![Go Report Card](https://goreportcard.com/badge/sigs.k8s.io/cli-experimental)](https://goreportcard.com/report/sigs.k8s.io/cli-experimental)

# cli-experimental

Experimental Kubectl libraries and commands.

## commands
This repo can build a binary `k2` by
```
GO111MODULE=on go build ./cmd/k2
```
It provides following builtin commands as well as adding dynamic commands.
- apply
- status
- prune
- delete

 TODO: expand each of these

 ### dynamic commands

 ## libraries
You can also use cli-experimental as a library. It provides a library to run `apply`,
`status`, `prune` and `delete` for a list of Unstructured resources.

 ```Go
import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/cli-experimental/pkg"
	)
 c := pkg.InitializeCmd(io.Stdout, nil)
var resources []unstructured.Unstructured
 err := c.Apply(resources)
err = c.Status(resources)
err = c.Prune(resources)
err = c.Delete(resources)
```
TODO: add explanation for passing flags

 ## Examples
TODO: add examples

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

- [Slack channel](https://kubernetes.slack.com/messages/sig-cli)
- [Mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-cli)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).

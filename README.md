[![Build Status](https://travis-ci.org/kubernetes-sigs/cli-experimental.svg?branch=master)](https://travis-ci.org/kubernetes-sigs/cli-experimental "Travis")
[![Go Report Card](https://goreportcard.com/badge/sigs.k8s.io/cli-experimental)](https://goreportcard.com/report/sigs.k8s.io/cli-experimental)

# cli-experimental

Experimental Kubectl libraries and commands.

## Kustomize documentation

This repo currently hosts the code for [Kustomize](https://github.com/kubernetes-sigs/kustomize) and [Kubectl](https://github.com/kubernetes/kubectl) docs, which are available at https://kubectl.docs.kubernetes.io. The guide for contributing to the docs is hosted [within the docs themselves](https://kubectl.docs.kubernetes.io/contributing/docs/).

The docs are built with [Hugo](https://gohugo.io/) and deployed with [Netlify](https://app.netlify.com/sites/cli-experimental/deploys). Information about this standard setup for subproject sites can be found [on the community repo](https://github.com/kubernetes/community/blob/master/github-management/subproject-site-requests.md).

Note that kustomize.io and kubectl.io were contributed by [Replicated](https://www.replicated.com/) and are currently still sourced from their repos:
* [replicatedcom/kustomize-www](https://github.com/replicatedcom/kustomize-www)
* [replicatedhq/kustomize-demo-api](https://github.com/replicatedhq/kustomize-demo-api)
* [replicatedhq/kustomize-demo](https://github.com/replicatedhq/kustomize-demo)

Ultimately we want to consolidate the Kustomize documentation and move the code out of the experimental repo. To learn more about that plan and get involved, see [kustomize/#4338](https://github.com/kubernetes-sigs/kustomize/issues/4338).

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

- [Slack channel](https://kubernetes.slack.com/messages/sig-cli)
- [Mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-cli)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).

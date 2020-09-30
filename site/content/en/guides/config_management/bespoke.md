---
title: "Bespoke Application"
linkTitle: "Bespoke Application"
type: docs
weight: 8
description: >
    Workflow for bespoke applications
---

In this workflow, all configuration (resource YAML) files are owned by the user.
No content is incorporated from version control repositories owned by others.

![bespoke config workflow image][workflowBespoke]

#### 1) create a directory in version control

Speculate some overall cluster application called _ldap_;
we want to keep its configuration in its own repo.

> ```
> git init ~/ldap
> ```

#### 2) create a [base]

> ```
> mkdir -p ~/ldap/base
> ```

In this directory, create and commit a [kustomization]
file and a set of [resources].

#### 3) create [overlays]

> ```
> mkdir -p ~/ldap/overlays/staging
> mkdir -p ~/ldap/overlays/production
> ```

Each of these directories needs a [kustomization]
file and one or more [patches].

The _staging_ directory might get a patch
that turns on an experiment flag in a configmap.

The _production_ directory might get a patch
that increases the replica count in a deployment
specified in the base.

#### 4) bring up [variants]

Run kustomize, and pipe the output to [apply].

> ```
> kustomize build ~/ldap/overlays/staging | kubectl apply -f -
> kustomize build ~/ldap/overlays/production | kubectl apply -f -
> ```

You can also use [kubectl-v1.14.0] to apply your [variants].
>
> ```
> kubectl apply -k ~/ldap/overlays/staging
> kubectl apply -k ~/ldap/overlays/production
> ```

[OTS]: /cli-experimental/references/kustomize/glossary#off-the-shelf-configuration
[apply]: /cli-experimental/references/kustomize/glossary#apply
[applying]: /cli-experimental/references/kustomize/glossary#apply
[base]: /cli-experimental/references/kustomize/glossary#base
[fork]: https://guides.github.com/activities/forking/
[variants]: /cli-experimental/references/kustomize/glossary#variant
[kustomization]: /cli-experimental/references/kustomize/glossary#kustomization
[off-the-shelf]: /cli-experimental/references/kustomize/glossary#off-the-shelf-configuration
[overlays]: /cli-experimental/references/kustomize/glossary#overlay
[patch]: /cli-experimental/references/kustomize/glossary#patch
[patches]: /cli-experimental/references/kustomize/glossary#patch
[rebase]: https://git-scm.com/docs/git-rebase
[resources]: /cli-experimental/references/kustomize/glossary#resource
[workflowBespoke]: /cli-experimental/images/workflowBespoke.jpg
[workflowOts]: /cli-experimental/images/workflowOts.jpg
[kubectl-v1.14.0]:https://kubernetes.io/blog/2019/03/25/kubernetes-1-14-release-announcement/

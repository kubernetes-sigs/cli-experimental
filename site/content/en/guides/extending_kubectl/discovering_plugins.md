---
title: "Discovering Plugins"
linkTitle: "Discovering Plugins"
weight: 2
type: docs
description: >
   Discovering Plugins that suits your requirement
---


{{< alert color="success" title="TL;DR" >}}
- [krew.sigs.k8s.io](https://krew.sigs.k8s.io/docs/user-guide/setup/install/) is a kubernetes sub-project to discover and manage plugins
{{< /alert >}}

# Krew

By design, `kubectl` does not install plugins. This task is left to the kubernetes sub-project
[krew.sigs.k8s.io](https://krew.sigs.k8s.io/docs/user-guide/setup/install/) which needs to be installed separately.
Krew helps to

- discover plugins
- get updates for installed plugins
- remove plugins

## Installing krew

Krew should be used as a kubectl plugin. To set yourself up to using krew, you need to do two things:

1. Install git
1. Install krew as described on the project page [krew.sigs.k8s.io](https://krew.sigs.k8s.io/docs/user-guide/setup/install/).
1. Add the krew bin folder to your `PATH` environment variable. For example, in bash `export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"`.

## Krew capabilities

Discover plugins
```bash
kubectl krew search
```

Install a plugin
```bash
kubectl krew install access-matrix
```

Upgrade all installed plugins
```bash
kubectl krew upgrade
```

Show details about a plugin
```bash
kubectl krew info access-matrix
```

Uninstall a plugin
```bash
kubectl krew uninstall access-matrix
```
---
title: "Binaries"
linkTitle: "Binaries"
weight: 3
type: docs
description: >
  Install Kustomize by downloading precompiled binaries.
---

Binaries at various versions for linux, MacOs and Windows are published on the [releases page].

The following [script] detects your OS and downloads the appropriate kustomize binary to your
current working directory.

```bash
curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash
```

**This script doesn't work for ARM architecture.** If you want to install ARM binaries, please
go to the release page to find the URL.

[releases page]: https://github.com/kubernetes-sigs/kustomize/releases
[script]: https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh

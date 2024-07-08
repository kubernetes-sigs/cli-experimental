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
If you are rate limited by GitHub, you can export the variable `GITHUB_TOKEN` with a valid [GITHUB Pat Token] and avoid to be throttled.

```
export GITHUB_TOKEN="github_pat_........."; curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash
```

[releases page]: https://github.com/kubernetes-sigs/kustomize/releases
[script]: https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh
[GITHUB Pat Token]: https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens

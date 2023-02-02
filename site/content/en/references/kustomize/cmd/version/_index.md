---
title: "kustomize version"
linkTitle: "version"
type: docs
weight: 8
description: >
    Prints the kustomize version
---

```
Prints the kustomize version

Usage:
  kustomize version [flags]

Examples:
kustomize version

Flags:
  -h, --help            help for version
  -o, --output string   One of 'yaml' or 'json'.

Global Flags:
      --stack-trace   print a stack-trace on error
```

## Examples

```
$ kustomize version
v5.0.0

kustomize version --short
Flag --short has been deprecated, and will be removed in the future.
{kustomize/v5.0.0  2023-02-02T16:43:10Z   }

$ kustomize version -o yaml
version: kustomize/v5.0.0
gitCommit: 738ca56ccd511a5fcd57b958d6d2019d5b7f2091
buildDate: "2023-02-02T16:43:10Z"
goOs: darwin
goArch: arm64
goVersion: go1.19.5

$ kustomize version -o json
{
  "version": "kustomize/v5.0.0",
  "gitCommit": "738ca56ccd511a5fcd57b958d6d2019d5b7f2091",
  "buildDate": "2023-02-02T16:43:10Z",
  "goOs": "darwin",
  "goArch": "arm64",
  "goVersion": "go1.19.5"
}
```

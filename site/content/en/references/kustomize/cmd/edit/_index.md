---
title: "kustomize edit"
linkTitle: "edit"
type: docs
weight: 5
description: >
    Edits a kustomization file
---

The `kustomize edit command` provides the ability to manipulate an existing Kustomization configuration. As such, you must have either a valid `kustomization.yaml`, `kustomization.yml` or Kustomization resource already defined in order to use the edit command. The available subcommands include:

```
add                       Adds an item to the kustomization file
fix                       Fix the missing fields in kustomization file
remove                    Removes items from the kustomization file
set                       Sets the value of different fields in kustomization file
```

example command:
`kustomize edit add configmap my-configmap --from-file=my-key=file/path --from-literal=my-literal=12345`
---
title: "bases"
linkTitle: "bases"
type: docs
weight: 1
description: >
    Add resources from a kustomization dir.
---

{{% pageinfo color="warning" %}}
The `bases` field was deprecated in v2.1.0
{{% /pageinfo %}}

Move entries into the [resources](/references/kustomize/kustomization/resource)
field.  This allows bases - which are still a
[central concept](/references/kustomize/glossary#base) - to be
ordered relative to other input resources.

---
title: "Diffing Local and Remote Resources"
linkTitle: "Diffing Local and Remote Resources"
weight: 1
type: docs
description: >
   Diffing Local and Remote Resources
---


{{< alert color="success" title="TL;DR" >}}
- View diff of changes before they are Applied to the cluster
{{< /alert >}}

# Diffing Local and Cluster State

## Motivation

The ability to view what changes will be made before applying them to a cluster can be useful.

## Generating a Diff

Use the `diff` program in a user's path to display a diff of the changes that will be
made by Apply.

```sh
kubectl diff -k ./dir/
```

## Setting the Diff Program

The `KUBECTL_EXTERNAL_DIFF` environment variable can be used to select your own diff command.
By default, the "diff" command available in your path will be run with "-u" (unified) and "-N"
(treat new files as empty) options.

```sh
export KUBECTL_EXTERNAL_DIFF=meld; kubectl diff -k ./dir/
```

## Exit status
The following exit values shall be returned:

 `0`
No differences were found.
 `1`
Differences were found.
 `>1`
Kubectl or diff failed with an error.

**Note:** `KUBECTL_EXTERNAL_DIFF`, if used, is expected to follow that convention.
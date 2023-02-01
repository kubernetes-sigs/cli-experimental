---
title: "kustomize localize"
linkTitle: "localize"
type: docs
weight: 9
description: >
  Downloads url content to a local directory, and replaces urls with paths to downloaded content
---

Disclaimer: This is an alpha command. Please see the [command proposal](https://github.com/kubernetes-sigs/kustomize/blob/master/proposals/22-04-localize-command.md)
for full capabilities.

### Description

The `kustomize localize` command makes a recursive copy of a kustomization
in which referenced urls are replaced by local paths to their downloaded content.
The command copies files referenced under kustomization [fields](#fields).
The copy contains files referenced both by the kustomization in question
and by recursively referenced kustomizations.

The purpose of this command is to create a copy on which
`kustomize build`<sup>[[build]](#notes)</sup> produces the same output
without a network connection. The original motivation for this command
is documented [here](https://github.com/kubernetes-sigs/kustomize/issues/3980).
A `kustomize build` use case precluding network use could be a CI/CD pipeline
that only has access to the internal network.

### Usage

The command takes the following form:

<pre>
<b>kustomize localize</b> [<ins>target</ins> [<ins>newDir</ins>]] [--scope <ins>scope</ins>]
</pre>

where

* `target` is the [kustomization directory](https://kubectl.docs.kubernetes.io/references/kustomize/glossary/#kustomization-root) 
to localize. This value can be a local path or a [remote root](https://github.com/kubernetes-sigs/kustomize/blob/master/examples/remoteBuild.md). 
The default value is the current working directory.
* `newDir` is the destination of the "localized" copy that the command creates. 
The destination cannot already exist. 
The command creates the destination directory, but not any of its parents.
The default destination is a directory in the current working directory named:
  * `localized-{target}` for local `target`
  * `localized-{target}-{ref}`<sup>[[ref]](#notes)</sup> for remote `target`.
    See an [example](#example).
* `scope` is the "bounding directory"; in other words, the command 
can only copy files inside this directory. The default is `{target}`. 
This flag cannot be specified for remote `target`, as its value is implicitly 
the repo containing `target`.

### Structure

The localized destination directory is a copy<sup>[[absolute]](#notes), [[symlink]](#notes)</sup>
of `scope`, excluding files that `target` does not reference and
with the addition of downloaded remote content.

Downloaded files are copied to a directory named `localized-files` located in
the same directory as the referencing kustomization. Inside `localized-files`,
the content of remote
* roots are written to path<sup>[[localized root]](#notes)</sup>:

  <pre>
  <ins>host</ins> / <ins>path/to/repo</ins> / <ins>ref</ins> / <ins>path/to/file/in/repo</ins>
  </pre>

* files are written to the following path<sup>[[localized file]](#notes)</sup>
  constructed from its url components:

  <pre>
  <ins>host</ins> / <ins>path</ins>
  </pre>

### Example

Running the following command:

<!--
TODO(annasong): Replace ref with kustomize/v5.0 after release. 
The kustomize/v4.5.7 version is very slow to execute because it fetches
submodules.
-->
```shell
$ kustomize localize https://github.com/kubernetes-sigs/kustomize//api/krusty/testdata/localize/remote?ref=kustomize/v4.5.7&submodules=0&timeout=300
```

in an empty directory named `example` creates the localized destination
with the following contents:

```shell
example
└── localized-remote-kustomize-v4.5.7
    ├── localized-files
    │   └── github.com
    │       └── kubernetes-sigs
    │           └── kustomize
    │               └── kustomize
    │                   └── v4.5.7
    │                       └── api
    │                           └── krusty
    │                               └── testdata
    │                                   └── localize
    │                                       └── simple
    │                                           ├── deployment.yaml
    │                                           ├── service.yaml
    │                                           └── kustomization.yaml
    └── hpa.yaml
```
```shell
# example/localized-remote-kustomize-v4.5.7/kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
commonLabels:
  purpose: remoteReference
kind: Kustomization
resources:
- localized-files/github.com/kubernetes-sigs/kustomize/kustomize/v4.5.7/api/krusty/testdata/localize/simple
- hpa.yaml
```
```shell
# example/localized-remote-kustomize-v4.5.7/localized-files/github.com/kubernetes-sigs/kustomize/kustomize/v4.5.7/api/krusty/testdata/localize/simple/kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namePrefix: localize-
resources:
- deployment.yaml
- service.yaml
```

The [proposal](https://github.com/kubernetes-sigs/kustomize/blob/master/proposals/22-04-localize-command.md)
contains more examples.

### Fields

The command localizes, copying or downloading, files under 
the following kustomization fields<sup>[[resource]](#notes)</sup>:

* `resources`
* `components`
* `openapi.path`
* `configurations`
* `crds`
* `configMapGenerator`<sup>[[gen]](#notes)</sup>
* `secretGenerator`<sup>[[gen]](#notes)</sup>
* `helmCharts`<sup>[[helm]](#notes)</sup>
* `helmGlobals`<sup>[[helm]](#notes)</sup>
* `patches`
* `replacements`

In addition to localizing files<sup>[[plugin]](#notes)</sup> under the following 
plugin fields:

* `generators`
* `transformers`
* `validators`

the command localizes files referenced by the following plugins<sup>[[resource]](#notes)</sup>,
which have a built-in kustomization field counterpart:

* `ConfigMapGenerator`<sup>[[gen]](#notes)</sup>
* `SecretGenerator`<sup>[[gen]](#notes)</sup>
* `HelmChartInflationGenerator`<sup>[[helm]](#notes)</sup>
* `PatchTransformer`
* `PatchJson6902Transformer`
* `PatchStrategicMergeTransformer`
* `ReplacementTransformer`

The command also localizes the following deprecated fields:
* `bases`
* `helmChartInflationGenerator`<sup>[[helm]](#notes)</sup>
* `patchesStrategicMerge`
* `patchesJson6902`

### Notes
* [absolute]: The alpha version of this command does not handle and
throws an error for absolute paths.
<br></br>

* [build]: This command may not catch `build` fallacies in the kustomization, as
this command serves a different purpose than `kustomize build` and does not look
to overstep its scope.
<br></br>
  
  However, this command will fail on a kustomization that requires the 
`kustomize build --load-restrictor LoadRestrictionsNone` flag to run, as
this command copies files following security best practices.<sup>[[symlink]](#notes)</sup>
<br></br>

* [gen]: If a key is not specified for a `ConfigMapGenerator`, `SecretGenerator`
`files` entry, this command does not add one. Because in such a case the 
`build` output key is the file name, the output key will be different on the 
localized copy if the file for said entry is a symlink to a file with a
different name<sup>[[symlink]](#notes)</sup>.
<br></br>

* [helm]: Because helm support in kustomize is [intentionally limited](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/helmcharts/),
this command does not download remote content referenced by helm transformers.
This command does, however, minimally copy local `valuesFile` and `chartHome`, 
albeit without following symlinks in `chartHome`.
<br></br>

* [localized file]: This command is primarily concerned with supporting
remote files of the form `https://raw.githubusercontent.com/kubernetes-sigs/kustomize/kustomize/v4.5.7/proposals/22-04-localize-command.md`,
generated by clicking `Raw` in the GitHub UI. The url path of these files 
consists of 

  <pre>
  <ins>path/to/repo</ins> / <ins>ref</ins> / <ins>path/to/file/in/repo</ins>
  </pre>

  The localized file path format was chosen as `host / path` so that the path 
for these GitHub files would match the localized root path format.
<br></br>

* [localized root]: A [remote root](https://github.com/kubernetes-sigs/kustomize/blob/master/examples/remoteBuild.md)
is a [git url](https://git-scm.com/docs/git-fetch#_git_urls) specifying a repo,
a double slash `//` delimiter, a path to the root inside the repo, and
query string parameters including [[ref]](#notes). The different segments of a
localized remote root path are the:

  * [host](https://git-scm.com/docs/git-fetch#_git_urls) of the git url.
Because [`file`-schemed](https://git-scm.com/docs/git-fetch#_git_urls)
git urls do not have a `host`, their localized paths under the
`localized-files` directory begin with a directory named `file-schemed`
instead of a `host` value.
  * [path/to/repo](https://git-scm.com/docs/git-fetch#_git_urls) of the git url
  * [[ref]](#notes)
  * `path/to/file/in/repo`, which refers to the path to the root after the 
`//` delimiter, concatenated with the relative path from said root
to referenced local files. This path is delinked<sup>[[symlink]](#notes)</sup>.
<br></br>

  This command does not include `scheme`, `userinfo`, or `port`
in the remote root's localized path for the sake of simplicity. Please leave
[feedback](https://github.com/kubernetes-sigs/kustomize/issues/4996)
if omitting these url components affects the correctness of your localized copy.
<br></br>

* [plugin]: The alpha version of this command handles plugin files, but not 
kustomization directories producing plugins. The command throws an error upon
encountering kustomizations under the plugin fields.
<br></br>

* [ref]: This command requires [remote roots](https://github.com/kubernetes-sigs/kustomize/blob/master/examples/remoteBuild.md)
to have a `ref` query string parameter. Ideally, if the ref is a stable tag, 
the content of the remote root is always the same. 
See [[localized root]](#notes) for more on remote roots.
<br></br>

* [resource]: As a byproduct of processing yaml files such as kustomizations,
this command writes their keys in alphabetical order to the destination.
<br></br>

* [symlink]: To avoid `kustomize build` load restriction errors on the 
localized copy, this command copies files using their actual locations, after
following symlinks. `kustomize build` enforces load restrictions using the
"delinked" locations of files. As long as this command preserves the delinked
structure of `scope` in the localized copy, the copy will satisfy 
load requirements.

### Feedback

Please leave your feedback for this command under [the following issue](https://github.com/kubernetes-sigs/kustomize/issues/4996).

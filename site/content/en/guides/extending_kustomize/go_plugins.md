---
title: "Go Plugins (deprecated)"
linkTitle: "Go Plugins"
type: docs
weight: 4
description: >
    Guide to writing Go plugins for Kustomize
---
{{% alert color="warning" title="Deprecation warning" %}}
This style of plugin is slated for deprecation.
See the [Kustomize Plugin Graduation KEP](https://github.com/kubernetes/enhancements/issues/2953) for information on the future of Kustomize plugins.
{{% /alert %}}

[plugin package]: https://golang.org/pkg/plugin
[Go modules]: https://github.com/golang/go/wiki/Modules
[ELF]: https://en.wikipedia.org/wiki/Executable_and_Linkable_Format
[tensorflow plugin]: https://www.tensorflow.org/guide/extend/op
[Go plugin]: https://golang.org/pkg/plugin/

## Authoring Go plugins

#### What is a Go plugin?

A _Go plugin_ is a compilation artifact described
by the Go [plugin package].  It is built with
special flags and cannot run on its own.
It must be loaded into a running Go program.

> A normal program written in Go might be usable
> as _exec plugin_, but is not a _Go plugin_.

Go plugins allow kustomize extensions that run
without the cost marshalling/unmarshalling all
resource data to/from a subprocess for each plugin
run.  The Go plugin API assures a certain level of
consistency to avoid confusing downstream
transformers.

#### Go plugins as Kustomize extensions

Go plugins work as described in the [plugin
package], but fall short of common notions
associated with the word _plugin_.

Be sure to read [Go plugin caveats](#go-plugin-caveats).

A `.go` file can be a [Go plugin] if it declares
'main' as it's package, and exports a symbol to
which useful functions are attached.

It can further be used as a _kustomize_ plugin if
the symbol is named 'KustomizePlugin' and the
attached functions implement the `Configurable`,
`Generator` and `Transformer` interfaces.

A Go plugin for kustomize looks like this:

> ```go
> package main
>
> import (
> "sigs.k8s.io/kustomize/api/resmap"
>   ...
> )
>
> type plugin struct {...}
>
> var KustomizePlugin plugin
>
> func (p *plugin) Config(
>    h *resmap.PluginHelpers,
>    c []byte) error {...}
>
> func (p *plugin) Generate() (resmap.ResMap, error) {...}
>
> func (p *plugin) Transform(m resmap.ResMap) error {...}
> ```

Use of the identifiers `plugin`, `KustomizePlugin`
and implementation of the method signature
`Config` is required.

Implementing the `Generator` or `Transformer`
method allows (respectively) the plugin's config
file to be added to the `generators` or
`transformers` field in the kustomization file.
Do one or the other or both as desired.

### Placement

Each plugin gets its own dedicated directory named

[`XDG_CONFIG_HOME`]: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

```bash
$XDG_CONFIG_HOME/kustomize/plugin
    /${apiVersion}/LOWERCASE(${kind})
```

The default value of [`XDG_CONFIG_HOME`] is
`$HOME/.config`.

The one-plugin-per-directory requirement eases
creation of a plugin bundle (source, tests, plugin
data files, etc.) for sharing.

In the case of a [Go plugin](#go-plugins), it also
allows one to provide a `go.mod` file for the
single plugin, easing resolution of package
version dependency skew.

When loading, kustomize will look for a
file called `${kind}.so` attempt to load it as a Go plugin.

```bash
$XDG_CONFIG_HOME/kustomize/plugin
    /${apiVersion}/LOWERCASE(${kind})/${kind}.so
```

Failure to find a plugin to load fails the overall
`kustomize build`.

[secret generator]: https://github.com/kubernetes-sigs/kustomize/tree/master/plugin/someteam.example.com/v1/secretsfromdatabase
[service generator]: https://github.com/kubernetes-sigs/kustomize/tree/master/plugin/someteam.example.com/v1/someservicegenerator
[string prefixer]: https://github.com/kubernetes-sigs/kustomize/tree/master/plugin/someteam.example.com/v1/stringprefixer
[date prefixer]: https://github.com/kubernetes-sigs/kustomize/tree/master/plugin/someteam.example.com/v1/dateprefixer
[sops encoded secrets]: https://github.com/monopole/sopsencodedsecrets
[SOPSGenerator]: https://github.com/omninonsense/kustomize-sopsgenerator

### Examples

* [service generator] - generate a service from a name and port argument.
* [string prefixer] - uses the value in `metadata/name` as the prefix.
   This particular example exists to show how a plugin can
   transform the behavior of a plugin.  See the
   `TestTransformedTransformers` test in the `target` package.
* [date prefixer] - prefix the current date to resource names, a simple
   example used to modify the string prefixer plugin just mentioned.
* [secret generator] - generate secrets from a toy database.
* [sops encoded secrets] - a more complex secret generator that converts SOPS files into Kubernetes Secrets
* [SOPSGenerator] - another generator that decrypts SOPS files into Secrets
* [All the builtin plugins](https://github.com/kubernetes-sigs/kustomize/tree/master/plugin/builtin).
   User authored plugins are
   on the same footing as builtin operations.

A Go plugin can be both a generator and a
transformer.  The `Generate` method will run along
with all the other generators before the
`Transform` method runs.

Here's a build command that sensibly assumes the
plugin source code sits in the directory where
kustomize expects to find `.so` files:

```bash
d=$XDG_CONFIG_HOME/kustomize/plugin\
/${apiVersion}/LOWERCASE(${kind})

go build -buildmode plugin \
   -o $d/${kind}.so $d/${kind}.go
```

## Go plugin caveats

### The skew problem

Go plugin compilation creates an [ELF] formatted
`.so` file, which by definition has no information
about the provenance of the object code.

Skew between the compilation conditions (versions
of package dependencies, `GOOS`, `GOARCH`) of the
main program ELF and the plugin ELF will cause
plugin load failure, with non-helpful error
messages.

Exec plugins also lack provenance, but won't fail
due to compilation skew.

In either case, the only sensible way to share a
plugin is as some kind of _bundle_ (a git repo
URL, a git archive file, a tar file, etc.)
containing source code, tests and associated data,
unpackable under
`$XDG_CONFIG_HOME/kustomize/plugin`.

In the case of a Go plugin, an _end user_
accepting a shared plugin _must compile both
kustomize and the plugin_.

This means a one-time run of

```bash
# Or whatever is appropriate at time of reading
GOPATH=${whatever} GO111MODULE=on go get sigs.k8s.io/kustomize/api
```

and then a normal development cycle using

```bash
go build -buildmode plugin \
    -o ${wherever}/${kind}.so ${wherever}/${kind}.go
```

with paths and the release version tag (e.g. `v3.0.0`)
adjusted as needed.

For comparison, consider what one
must do to write a [tensorflow plugin].

## Go plugin advantages

### Safety

The Go plugin developer sees the same API offered
to native kustomize operations, assuring certain
semantics, invariants, checks, etc. An exec
plugin sub-process dealing with this via
stdin/stdout will have an easier time screwing
things up for downstream transformers and
consumers.

Minor point: if the plugin reads files via
the kustomize-provided file `Loader` interface, it
will be constrained by kustomize file loading
restrictions.  Of course, nothing but a code audit
prevents a Go plugin from importing the `io` package
and doing whatever it wants.

### Debugging

A Go plugin developer can debug the plugin _in
situ_, setting breakpoints inside the plugin and
elsewhere while running a plugin in feature tests.

To get the best of both worlds (shareability and safety),
a developer can write an `.go` program that functions
as an _exec plugin_, but can be processed by `go generate`
to emit a _Go plugin_ (or vice versa).


## Guided example

[SopsEncodedSecrets repository]: https://github.com/monopole/sopsencodedsecrets
[Go plugin]: https://golang.org/pkg/plugin
[Go plugin caveats]: goPluginCaveats.md

This is a (no reading allowed!) 60 second copy/paste guided
example.

Full plugin docs [here](/guides/extending_kustomize/gopluginguidedexample/).
Be sure to read the [Go plugin caveats](/guides/extending_kustomize/goplugincaveats/).

This demo uses a Go plugin, `SopsEncodedSecrets`,
that lives in the [sopsencodedsecrets repository].
This is an inprocess [Go plugin], not an
sub-process exec plugin that happens to be written
in Go (which is another option for Go authors).

This is a guide to try it without damaging your
current setup.

Prerequisites:
* linux
* git
* curl
* Go 1.13

For encryption:
* gpg

--OR--

* Google cloud (gcloud) install
* a Google account with KMS permission

### Make a place to work

```shell
# Keeping these separate to avoid cluttering the DEMO dir.
DEMO=$(mktemp -d)
tmpGoPath=$(mktemp -d)
```

### Install kustomize

Need v3.0.0 for what follows, and you must _compile_
it (not download the binary from the release page):

```shell
GOPATH=$tmpGoPath go install sigs.k8s.io/kustomize/kustomize
```

### Make a home for plugins

A kustomize plugin is fully determined by
its configuration file and source code.

[required fields]: https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/#required-fields

Kustomize plugin configuration files are formatted
as kubernetes resource objects, meaning
`apiVersion`, `kind` and `metadata` are [required
fields] in these config files.

The kustomize program reads the config file
(because the config file name appears in the
`generators` or `transformers` field in the
kustomization file), then locates the Go plugin's
object code at the following location:

```shell
$XDG_CONFIG_HOME/kustomize/plugin/$apiVersion/$lKind/$kind.so
```

where `lKind` holds the lowercased kind.  The
plugin is then loaded and fed its config, and the
plugin's output becomes part of the overall
`kustomize build` process.

The same plugin might be used multiple times in
one kustomize build, but with different config
files.  Also, kustomize might customize config
data before sending it to the plugin, for whatever
reason.  For these reasons, kustomize owns the
mapping between plugins and config data; it's not
left to plugins to find their own config.

This demo will house the plugin it uses at the
ephemeral directory

```shell
PLUGIN_ROOT=$DEMO/kustomize/plugin
```

and ephemerally set `XDG_CONFIG_HOME` on a command
line below.

### Get the plugin

At this stage in the development of kustomize
plugins, plugin code doesn't know or care what
`apiVersion` or `kind` appears in the config file
sent to it.

The plugin could check these fields, but it's the
remaining fields that provide actual configuration
data, and at this point the successful parsing of
these other fields are the only thing that matters
to a plugin.

This demo uses a plugin called _SopsEncodedSecrets_,
and it lives in the [SopsEncodedSecrets repository].

Somewhat arbitrarily, we'll choose to install
this plugin with

```shell
apiVersion=mygenerators
kind=SopsEncodedSecrets
```

#### Define the plugin's home dir

By convention, the ultimate home of the plugin
code and supplemental data, tests, documentation,
etc. is the lowercase form of its kind.

```shell
lKind=$(echo $kind | awk '{print tolower($0)}')
```

#### Clone the SopsEncodedSecrets plugin repo

In this case, the repo name matches the lowercase
kind already, so we just clone the repo and get
the proper directory name automatically:

```shell
mkdir -p $PLUGIN_ROOT/${apiVersion}
cd $PLUGIN_ROOT/${apiVersion}
git clone git@github.com:monopole/sopsencodedsecrets.git
```

Remember this directory:

```shell
MY_PLUGIN_DIR=$PLUGIN_ROOT/${apiVersion}/${lKind}
```

#### Try the plugin's own test

Plugins may come with their own tests.
This one does, and it hopefully passes:

```shell
cd $MY_PLUGIN_DIR
go test SopsEncodedSecrets_test.go
```

Build the object code for use by kustomize:

```shell
cd $MY_PLUGIN_DIR
GOPATH=$tmpGoPath go build -buildmode plugin -o ${kind}.so ${kind}.go
```

This step may succeed, but kustomize might
ultimately fail to load the plugin because of
dependency [skew].

[skew]: /guides/extending_kustomize/goplugincaveats/
[used in this demo]: #install-kustomize

On load failure

* be sure to build the plugin with the same
   version of Go (_go1.13_) on the same `$GOOS`
   (_linux_) and `$GOARCH` (_amd64_) used to build
   the kustomize being [used in this demo].

* change the plugin's dependencies in its `go.mod`
   to match the versions used by kustomize (check
   kustomize's `go.mod` used in its tagged commit).

Lacking tools and metadata to allow this to be
automated, there won't be a Go plugin ecosystem.

Kustomize has adopted a Go plugin architecture as
to ease accept new generators and transformers
(just write a plugin), and to be sure that native
operations (also constructed and tested as
plugins) are compartmentalized, orderable and
reusable instead of bizarrely woven throughout the
code as a individual special cases.

### Create a kustomization

Make a kustomization directory to
hold all your config:

```shell
MYAPP=$DEMO/myapp
mkdir -p $MYAPP
```

Make a config file for the SopsEncodedSecrets plugin.

Its `apiVersion` and `kind` allow the plugin to be
found:

```shell
cat <<EOF >$MYAPP/secGenerator.yaml
apiVersion: ${apiVersion}
kind: ${kind}
metadata:
  name: mySecretGenerator
name: forbiddenValues
namespace: production
file: myEncryptedData.yaml
keys:
- ROCKET
- CAR
EOF
```

This plugin expects to find more data in
`myEncryptedData.yaml`; we'll get to that shortly.

Make a kustomization file referencing the plugin
config:

```shell
cat <<EOF >$MYAPP/kustomization.yaml
commonLabels:
  app: hello
generators:
- secGenerator.yaml
EOF
```

### Generate encrypted data

First you need to ensure you have an encryption tool installed.

We're going to use [sops](https://github.com/mozilla/sops) to encode a file. Choose either GPG or Google Cloud KMS as the secret provider to continue.

#### GPG

Try this:

```shell
gpg --list-keys
```

If it returns a list, presumably you've already created keys. If not, try import test keys from sops for dev.

```shell
curl https://raw.githubusercontent.com/mozilla/sops/master/pgp/sops_functional_tests_key.asc | gpg --import
SOPS_PGP_FP="1022470DE3F0BC54BC6AB62DE05550BC07FB1A0A"
```

#### Google Cloude KMS

Try this:

```shell
gcloud kms keys list --location global --keyring sops
```

If it succeeds, presumably you've already created keys and placed them in a keyring called sops. If not, do this:

```shell
gcloud kms keyrings create sops --location global
gcloud kms keys create sops-key --location global \
    --keyring sops --purpose encryption
```

Extract your keyLocation for use below:

```shell
keyLocation=$(\
    gcloud kms keys list --location global --keyring sops |\
    grep GOOGLE | cut -d " " -f1)
echo $keyLocation
```

#### Install `sops`

```shell
GOPATH=$tmpGoPath go install go.mozilla.org/sops/cmd/sops
```

#### Create data encrypted with your private key

Create raw data to encrypt:

```shell
cat <<EOF >$MYAPP/myClearData.yaml
VEGETABLE: carrot
ROCKET: saturn-v
FRUIT: apple
CAR: dymaxion
EOF
```

Encrypt the data into file the plugin wants to read:

With PGP

```shell
$tmpGoPath/bin/sops --encrypt \
  --pgp $SOPS_PGP_FP \
  $MYAPP/myClearData.yaml >$MYAPP/myEncryptedData.yaml
```

Or GCP KMS

```shell
$tmpGoPath/bin/sops --encrypt \
  --gcp-kms $keyLocation \
  $MYAPP/myClearData.yaml >$MYAPP/myEncryptedData.yaml
```

### Review the files

```shell
tree $DEMO
```

This should look something like:

> ```shell
> /tmp/tmp.0kIE9VclPt
> ├── kustomize
> │   └── plugin
> │       └── mygenerators
> │           └── sopsencodedsecrets
> │               ├── go.mod
> │               ├── go.sum
> │               ├── LICENSE
> │               ├── README.md
> │               ├── SopsEncodedSecrets.go
> │               ├── SopsEncodedSecrets.so
> │               └── SopsEncodedSecrets_test.go
> └── myapp
>     ├── kustomization.yaml
>     ├── myClearData.yaml
>     ├── myEncryptedData.yaml
>     └── secGenerator.yaml
> ```

### Build your app

```shell
XDG_CONFIG_HOME=$DEMO $tmpGoPath/bin/kustomize build --enable-alpha-plugins $MYAPP
```

This should emit a kubernetes secret, with
encrypted data for the names `ROCKET` and `CAR`.

Above, if you had set

```shell
PLUGIN_ROOT=$HOME/.config/kustomize/plugin
```

there would be no need to use `XDG_CONFIG_HOME` in the
_kustomize_ command above.

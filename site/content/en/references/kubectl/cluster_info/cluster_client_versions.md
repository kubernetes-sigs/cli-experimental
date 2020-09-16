---
title: "Cluster and Client versions"
linkTitle: "Cluster and Client versions"
weight: 1
type: docs
description: >
    Print information about the Cluster and Client versions
---

## Versions

The `kubectl version` prints the client and server versions.  Note that
the client version may not be present for clients built locally from
source.

```bash
kubectl version
```

```bash
Client Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.5", GitCommit:"f01a2bf98249a4db383560443a59bed0c13575df", GitTreeState:"clean", BuildDate:"2018-03-19T19:38:17Z", GoVersion:"go1.9.4", Compiler:"gc", Platform:"darwin/amd64"}
Server Version: version.Info{Major:"1", Minor:"11+", GitVersion:"v1.11.6-gke.2", GitCommit:"04ad69a117f331df6272a343b5d8f9e2aee5ab0c", GitTreeState:"clean", BuildDate:"2019-01-04T16:19:46Z", GoVersion:"go1.10.3b4", Compiler:"gc", Platform:"linux/amd64"}
```

{{< alert color="warning" title="Version Skew" >}}
Kubectl supports +/-1 version skew with the Kubernetes cluster.  Kubectl
versions that are more than 1 version ahead of or behind the cluster are
not guaranteed to be compatible.
{{< /alert >}}

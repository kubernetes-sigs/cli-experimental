---
title: "Proxying traffic to services"
linkTitle: "Proxying traffic to services"
weight: 5
type: docs
description: >
    Proxying traffic to services
---


{{< alert color="success" title="TL;DR" >}}
- Proxy local connections to Services running in the cluster
{{< /alert >}}

# Connecting to Services

## Motivation

Not all Services running a Kubernetes cluster are exposed externally.  However Services
only exposed internally to a cluster with a *clusterIp* are accessible through an
apiserver proxy.

Users may use Proxy to **connect to Kubernetes Services in a cluster that are not
externally exposed**.


**Note:** Services running a type LoadBalancer or type NodePort may be exposed externally and
accessed without the need for a Proxy.

## Connecting to an internal Service

Connect to a internal Service using the Proxy command, and the Service Proxy url.

To visit the nginx service go to the Proxy URL at
`http://127.0.0.1:8001/api/v1/namespaces/default/services/nginx/proxy/`


```bash
kubectl proxy

Starting to serve on 127.0.0.1:8001
```

```bash
curl http://127.0.0.1:8001/api/v1/namespaces/default/services/nginx/proxy/
```

{{< alert color="success" title="Literal Syntax" >}}
To connect to a Service through a proxy the user must build the Proxy URL.  The Proxy URL format is:

```bash
http://<apiserver-address>/api/v1/namespaces/<service-namespace>/services/[https:]<service-name>[:<port-name>]/proxy
```

- The apiserver-address should be the URL printed by the Proxy command
- The Port is optional if you haven’t specified a name for your port
- The Protocol is optional if you are using `http`

{{< /alert >}}

## Builtin Cluster Services

A common usecase is to connect to Services running as part of the cluster itself.  A user can print out these
Services and their Proxy Urls with `kubectl cluster-info`.

```bash
kubectl cluster-info

Kubernetes master is running at https://104.197.5.247
GLBCDefaultBackend is running at https://104.197.5.247/api/v1/namespaces/kube-system/services/default-http-backend:http/proxy
Heapster is running at https://104.197.5.247/api/v1/namespaces/kube-system/services/heapster/proxy
KubeDNS is running at https://104.197.5.247/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://104.197.5.247/api/v1/namespaces/kube-system/services/https:metrics-server:/proxy
```

{{< alert color="success" title="More Info" >}}
For more information on connecting to a cluster, see
[Accessing Clusters](https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/).
{{< /alert >}}


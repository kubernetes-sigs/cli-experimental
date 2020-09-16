---
title: "Control Plane"
linkTitle: "Control Plane"
weight: 2
type: docs
description: >
    Print information about the Control Plane
---

## Control Plane and Addons

The `kubectl cluster-info` prints information about the control plane and
add-ons.

```bash
kubectl cluster-info
```

```bash
  Kubernetes master is running at https://1.1.1.1
  GLBCDefaultBackend is running at https://1.1.1.1/api/v1/namespaces/kube-system/services/default-http-backend:http/proxy
  Heapster is running at https://1.1.1.1/api/v1/namespaces/kube-system/services/heapster/proxy
  KubeDNS is running at https://1.1.1.1/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
  Metrics-server is running at https://1.1.1.1/api/v1/namespaces/kube-system/services/https:metrics-server:/proxy
```

{{< alert color="success" title="Kube Proxy" >}}
The URLs printed by `cluster-info` can be accessed at `127.0.0.1:8001` by
running `kubectl proxy`. 
{{< /alert >}}
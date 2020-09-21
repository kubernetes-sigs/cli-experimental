---
title: "logs"
linkTitle: "logs"
weight: 1
type: docs
description: >
    Getting logs for Kubernetes resources
---

Print the logs for a container in a pod or specified resource. If the pod has only one container, the container name is optional.

## Command
```bash
$ kubectl logs [-f] [-p] (POD | TYPE/NAME) [-c CONTAINER]
```

## Example

### Current State
```bash
$ kubectl get deployments

NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   1/1     1            1           13m
```

### Command
```bash
$ kubectl logs deployment/nginx
```

### Output
```bash
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
```

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

## Print Logs for a Container in a Pod

Print the logs for a Pod running a single Container

```bash
kubectl logs echo-c6bc8ccff-nnj52
```

```bash
hello
hello
```


{{< alert color="success" title="Crash Looping Containers" >}}
If a container is crash looping and you want to print its logs after it
exits, use the `-p` flag to look at the **logs from containers that have
exited**.  e.g. `kubectl logs -p -c ruby web-1`
{{< /alert >}}

---

## Print Logs for all Pods for a Workload

Print the logs for all Pods for a Workload

```bash
# Print logs from all containers matching label
kubectl logs -l app=nginx
```

{{< alert color="success" title="Workloads Logs" >}}
Print all logs from **all containers for a Workload** by passing the
Workload label selector to the `-l` flag.  e.g. if your Workload
label selector is `app=nginx` usie `-l "app=nginx"` to print logs
for all the Pods from that Workload.
{{< /alert >}}

---

## Follow Logs for a Container

Stream logs from a container.


```bash
# Follow logs from container
kubectl logs nginx-78f5d695bd-czm8z -f
```

---

## Printing Logs for a Container that has exited

Print the logs for the previously running container.  This is useful for printing containers that have
crashed or are crash looping.

```bash
# Print logs from exited container
kubectl logs nginx-78f5d695bd-czm8z -p
```

---

## Selecting a Container in a Pod 

Print the logs from a specific container within a Pod.  This is necessary for Pods running multiple
containers.

```bash
# Print logs from the nginx container in the nginx-78f5d695bd-czm8z Pod
kubectl logs nginx-78f5d695bd-czm8z -c nginx
```

---

## Printing Logs After a Time

Print the logs that occurred after an absolute time.

```bash
# Print logs since a date
kubectl logs nginx-78f5d695bd-czm8z --since-time=2018-11-01T15:00:00Z
```

---

## Printing Logs Since a Time

Print the logs that are newer than a duration.

Examples:

- 0s: 0 seconds
- 1m: 1 minute
- 2h: 2 hours

```bash
# Print logs for the past hour
kubectl logs nginx-78f5d695bd-czm8z --since=1h
```

---

## Include Timestamps

Include timestamps in the log lines

```bash
# Print logs with timestamps
kubectl logs -l app=echo --timestamps
```

```bash
2018-11-16T05:26:31.38898405Z hello
2018-11-16T05:27:13.363932497Z hello
```

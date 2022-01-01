---
title: "job"
linkTitle: "job"
weight: 1
type: docs
description: >
    Create a job with the specified name.
---

A Job creates one or more Pods and will continue to retry execution of the Pods until a specified number of them successfully terminate. As pods successfully complete, the Job tracks the successful completions. When a specified number of successful completions is reached, the task (ie, Job) is complete. Deleting a Job will clean up the Pods it created. Suspending a Job will delete its active Pods until the Job is resumed again.

A simple case is to create one Job object in order to reliably run one Pod to completion. The Job object will start a new Pod if the first Pod fails or is deleted (for example due to a node hardware failure or a node reboot).

## Command
```bash
$   kubectl create job NAME --image=image [--from=cronjob/name] -- [COMMAND] [args...] [options]
```

## Example

### Command
```bash
$ kubectl create job my-job --image=nginx
```

### Output
```bash
$ kubectl get jobs

NAME    COMPLETIONS    DURATION    AGE
my-job  1/1            8S          35s

$ kubectl get pods

NAME           READY   STATUS   RESTARTS   AGE 
my-job-mqwpv   1/1     Running  0          15s
```



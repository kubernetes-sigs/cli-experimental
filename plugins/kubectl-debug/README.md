# kubectl-debug plugin

This plugin demonstrates how to use [ephemeral
containers](https://features.k8s.io/277) to debug running pods. The `kubectl
debug` command will add an ephemeral container using an arbitrary container
image to the specified pod which can then be attached and used as a shell.

## Running

```sh
# assumes you have a working KUBECONFIG
$ GO111MODULE="on" go build cmd/kubectl-debug.go
# place the built binary somewhere in your PATH
$ cp ./kubectl-debug /usr/local/bin

# Create a pod to debug
$ cat example.yaml
apiVersion: v1
kind: Pod
metadata:
  name: debugtest
spec:
  shareProcessNamespace: true
  restartPolicy: Always
  containers:
  - name: nginx
    image: nginx:alpine
    stdin: true
    tty: true
  - name: two
    image: gcr.io/google_containers/pause-amd64:3.0
    stdin: true
    tty: true
$ kubectl create -f example.yaml
pod/debugtest created

# Attach a debian debug image
$ kubectl debug debugtest --image alpine --attach
If you don't see a command prompt, try pressing enter.
/ # ps auxww
PID   USER     TIME  COMMAND
    1 root      0:00 /pause
    6 root      0:00 nginx: master process nginx -g daemon off;
   11 101       0:00 nginx: worker process
   12 101       0:00 nginx: worker process
   13 101       0:00 nginx: worker process
   14 101       0:00 nginx: worker process
   15 101       0:00 nginx: worker process
   16 101       0:00 nginx: worker process
   17 101       0:00 nginx: worker process
   18 101       0:00 nginx: worker process
   19 root      0:00 /pause
   24 root      0:00 /bin/sh
   29 root      0:00 ps auxww
/ #
```

## Cleanup

You can "uninstall" this plugin from kubectl by simply removing it from your PATH:

    $ rm /usr/local/bin/kubectl-debug

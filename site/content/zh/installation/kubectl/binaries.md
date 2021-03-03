---
title: "Binaries"
linkTitle: "Binaries"
weight: 1
type: docs
description: >
    Install Kubectl by downloading precompiled binaries.
---

### Install kubectl binary with curl on Linux / macOS

1. Download the latest release with the command:

    ```bash
    curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
    ```

    To download a specific version, replace the `$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)` portion of the command with the specific version.

    For example, to download version v1.19.0 on Linux, type:
    
    ```bash
    curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.19.0/bin/linux/amd64/kubectl
    ```

2. Make the kubectl binary executable.

    ```bash
    chmod +x ./kubectl
    ```

3. Move the binary in to your PATH.

    ```bash
    sudo mv ./kubectl /usr/local/bin/kubectl
    ```
4. Test to ensure the version you installed is up-to-date:

    ```bash
    kubectl version --client
    ```

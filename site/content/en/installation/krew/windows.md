---
title: "Windows"
linkTitle: "Windows"
weight: 1
type: docs
description: >
    Install Krew in a Windows System
---

Steps of installation:

1. Make sure `git` is installed on your system.

1. Download `krew.exe` from the [Releases](https://github.com/kubernetes-sigs/krew/releases) page to a directory.

1. Launch a command-line window (`cmd.exe`) and navigate to that directory.

1. Run the following command to install krew:
    ```bash
    krew install krew
    ```
1. Add `%USERPROFILE%\.krew\bin` directory to your `PATH` environment variable ([how?](https://java.com/en/download/help/path.html))

1. Launch a new command-line window.

1. Verify running `kubectl krew` works.
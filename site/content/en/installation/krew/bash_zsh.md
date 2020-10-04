---
title: "Bash / ZSH"
linkTitle: "Bash / ZSH"
weight: 1
type: docs
description: >
    Install Krew using Bash or ZSH shells
---

Steps of installation:

1. Make sure that `git` is installed.

1. Run this command in your terminal to download and install krew:

    ```bash
    (
        set -x; cd "$(mktemp -d)" &&
        curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/krew.tar.gz" &&
        tar zxvf krew.tar.gz &&
        KREW=./krew-"$(uname | tr '[:upper:]' '[:lower:]')_amd64" &&
        "$KREW" install krew
    )
    ```

1. Add `$HOME/.krew/bin` directory to your PATH environment variable. To do this, update your .bashrc or .zshrc file and append the following line:
    ```bash
        export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"
    ```
    and restart your shell.

1. Verify running `kubectl krew` works.
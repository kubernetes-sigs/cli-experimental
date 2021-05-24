#!/usr/bin/env bash
# Copyright 2019 The Kubernetes Authors.
# SPDX-License-Identifier: Apache-2.0

set -e
TRACE=${TRACE:-""}
if [ -n "$TRACE" ]; then
  set -x
fi

goarch=amd64
goos="unknown"
if [[ "$OSTYPE" == "linux-gnu" ]]; then
  goos="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  goos="darwin"
fi
if [[ "$goos" == "unknown" ]]; then
  echo "OS '$OSTYPE' not supported. Aborting." >&2
  exit 1
fi

if [[ -z "$TMPDIR" ]]; then
  TMPDIR=/tmp
fi
if [[ -z "${KUBEBUILDER_ASSETS}" ]]; then
  export KUBEBUILDER_ASSETS=$kb_root_dir/bin
fi

function header_text {
  echo "##### $@"
}

kb_version=1.19.2
kb_dir=$TMPDIR/kubebuilder

# Skip fetching and untaring the tools by setting the SKIP_FETCH_TOOLS variable
# in your environment to any value:
#
# $ SKIP_FETCH_TOOLS=1 ./test.sh
#
# If you skip fetching tools, this script will use the tools already on your
# machine, but rebuild the kubebuilder and kubebuilder-bin binaries.
SKIP_FETCH_TOOLS=${SKIP_FETCH_TOOLS:-""}

function fetch_kb_tools {
  if [ -n "$SKIP_FETCH_TOOLS" ]; then
    return 0
  fi
  local archiveName="kubebuilder-tools-$kb_version-$goos-$goarch.tar.gz"
  local url="https://storage.googleapis.com/kubebuilder-tools/$archiveName"
  local archivePath="$TMPDIR/$archiveName"
  if [ ! -f $archivePath ]; then
    curl -sL ${url} -o "$archivePath"
  fi
  tar -zvxf "$archivePath" -C "$TMPDIR/"
}

header_text "fetching tools"


fetch_kb_tools

###
###  See https://github.com/kubernetes-sigs/cli-experimental/issues/170
###
#
# go get -v -u github.com/google/wire/cmd/wire
#
# header_text "running go vet"
# go vet ./internal/... ./pkg/... ./cmd/... ./util/...
#
# header_text "running golangci-lint"
# golangci-lint run ./... -D typecheck
#
# header_text "running go test"
# go test ./...
#

echo "passed"
exit 0

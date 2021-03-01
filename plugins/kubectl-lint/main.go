// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/dynamic"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/logs"
	"sigs.k8s.io/cli-experimental/plugins/kubectl-lint/pkglint"
	"sigs.k8s.io/cli-utils/pkg/errors"
	// This is here rather than in the libraries because of
	// https://github.com/kubernetes-sigs/kustomize/issues/2060

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// LintOptions options for lint subcommand
type LintOptions struct {
	FileNameOpts resource.FilenameOptions

	RecordFlags *genericclioptions.RecordFlags
	Recorder    genericclioptions.Recorder

	DynamicClient dynamic.Interface
	Mapper        meta.RESTMapper
	Result        *resource.Result

	genericclioptions.IOStreams
}

// NewLintOptions constructor to type LintOptions
func NewLintOptions(streams genericclioptions.IOStreams) *LintOptions {
	return &LintOptions{
		FileNameOpts: resource.FilenameOptions{},
		RecordFlags:  genericclioptions.NewRecordFlags(),
		Recorder:     genericclioptions.NoopRecorder{},
		IOStreams:    streams,
	}
}

func main() {
	ioStreams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	o := NewLintOptions(ioStreams)
	var f cmdutil.Factory
	lintCmd := &pkglint.Cmd{
		Factory:  f,
		FileNameOptions: &o.FileNameOpts,
		IOStreams:       ioStreams,
	}

	cmd := &cobra.Command{
		Use:                   "kubectl lint [-f FILENAME] [-k DIRECTORY]",
		DisableFlagsInUseLine: true,
		Short:                 i18n.T("Lint resource configuration files."),
		Long: i18n.T(`
Look for common issues with resource configuration.  Emit an error message if kubernetes best practices not followed and exit non-0.
`),
		Example: i18n.T(`#kubectl lint example.yaml and exit non-0 if issues found.
kubectl lint -f example.yaml
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return lintCmd.Run()
		},
	}

	// configure kubectl dependencies and flags
	o.RecordFlags.AddFlags(cmd)
	cmdutil.AddFilenameOptionFlags(cmd, &o.FileNameOpts, "full path to Kubernetes manifests (.yaml) to lint")
	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmd.Execute(); err != nil {
		errors.CheckErr(cmd.ErrOrStderr(), err, "klint")
	}
}

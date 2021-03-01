// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"sigs.k8s.io/cli-experimental/plugins/kubectl-lint/pkglint"
	"sigs.k8s.io/cli-utils/pkg/errors"
	// This is here rather than in the libraries because of
	// https://github.com/kubernetes-sigs/kustomize/issues/2060

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var cmd = &cobra.Command{
	Use:   "lint [-f FILENAME] [-k DIRECTORY]",
	Short: i18n.T("Lint resource configuration files."),
	Long:  i18n.T(`
Look for common issues with resource configuration.  Emit an error message if kubernetes best practices not followed and exit non-0.
`),
	SilenceErrors: true,
	SilenceUsage:  true,
}

func main() {
	ioStreams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	flags := cmd.PersistentFlags()
	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	kubeConfigFlags.AddFlags(flags)
	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)

	f := cmdutil.NewFactory(matchVersionKubeConfigFlags)
	
	cmd.AddCommand(pkglint.NewCmdLint(f, ioStreams))

	if err := cmd.Execute(); err != nil {
		errors.CheckErr(cmd.ErrOrStderr(), err, "kubectl-lint")
	}
}

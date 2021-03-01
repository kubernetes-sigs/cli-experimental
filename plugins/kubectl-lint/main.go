/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
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
	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	f := cmdutil.NewFactory(matchVersionKubeConfigFlags)
	cmd := pkglint.NewCmdLint(f, ioStreams)

	if err := cmd.Execute(); err != nil {
		errors.CheckErr(cmd.ErrOrStderr(), err, "kubectl-lint")
	}
}

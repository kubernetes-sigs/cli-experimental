/*
Copyright 2019 The Kubernetes Authors.
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

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"sigs.k8s.io/cli-experimental/cmd/apply"
	"sigs.k8s.io/cli-experimental/cmd/delete"
	"sigs.k8s.io/cli-experimental/cmd/prune"
	"sigs.k8s.io/cli-experimental/internal/pkg/dy"
	"sigs.k8s.io/cli-experimental/internal/pkg/wirecli/wirek8s"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(args []string, fn func(*cobra.Command)) error {
	rootCmd := &cobra.Command{
		// TODO(Liujingfang1): change this binary to a better name
		Use:   "cli-experimental",
		Short: "kubectl version 2",
		Long: `kubectl version 2
with commands apply, prune, delete and dynamic commands`,
	}
	if fn != nil {
		fn(rootCmd)
	}
	rootCmd.AddCommand(apply.GetApplyCommand(os.Args))
	rootCmd.AddCommand(prune.GetPruneCommand(os.Args))
	rootCmd.AddCommand(delete.GetDeleteCommand(os.Args))
	wirek8s.Flags(rootCmd.PersistentFlags())
	_ = rootCmd.PersistentFlags().Set("namespace", "default")

	// Add dynamic Commands published by CRDs as go-templates
	b, err := dy.InitializeCommandBuilder(rootCmd.OutOrStdout(), args)
	if err != nil {
		return err
	}
	if err := b.Build(rootCmd, nil); err != nil {
		return err
	}

	// Run the Command
	return rootCmd.Execute()
}

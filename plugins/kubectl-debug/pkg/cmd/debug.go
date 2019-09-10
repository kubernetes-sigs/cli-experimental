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
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
	clientset "k8s.io/client-go/kubernetes"
)

var (
	debugExample = `
		# Create a debugging container in pod mypod using defaults (a busybox image named "debugger")
		# and immediately attach to it using "kubectl attach".
		%[1]s debug mypod --attach

		# Create a debugging container in pod mypod using a custom debugging tools image and name.
		%[1]s debug -m gcr.io/verb-images/debug-tools -c debug-tools mypod
`

	errNoContext = fmt.Errorf("no context is currently set, use %q to select a new one", "kubectl config use-context <context>")
)

// DebugOptions provides information required to update
// the current context on a user's KUBECONFIG
type DebugOptions struct {
	configFlags *genericclioptions.ConfigFlags
	builder     *resource.Builder
	clientset   *clientset.Clientset

	args           []string
	attach         bool
	debugContainer v1.EphemeralContainer
	namespace      string

	genericclioptions.IOStreams
}

// NewDebugOptions provides an instance of DebugOptions with default values
func NewDebugOptions(streams genericclioptions.IOStreams) *DebugOptions {
	return &DebugOptions{
		configFlags: genericclioptions.NewConfigFlags(true),

		IOStreams: streams,
	}
}

// NewCmdDebug provides a cobra command wrapping DebugOptions
func NewCmdDebug(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewDebugOptions(streams)

	cmd := &cobra.Command{
		Use:          "debug [pod] [flags]",
		Short:        "Attach a debug container to a running pod",
		Example:      fmt.Sprintf(debugExample, "kubectl"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&o.attach, "attach", false, "Exec `kubectl attach` to attach to container after creation.")
	cmd.Flags().StringVarP(&o.debugContainer.Name, "container", "c", "debugger", "Container name. If omitted, a default will be chosen.")
	cmd.Flags().StringVarP(&o.debugContainer.Image, "image", "m", "busybox", "Container image.")
	cmd.Flags().BoolVarP(&o.debugContainer.Stdin, "stdin", "i", true, "stdin")
	cmd.Flags().BoolVarP(&o.debugContainer.TTY, "tty", "t", true, "tty")
	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

// Complete sets all information required for updating the current context
func (o *DebugOptions) Complete(cmd *cobra.Command, args []string) error {
	var err error

	o.args = args
	o.builder = resource.NewBuilder(o.configFlags)
	o.debugContainer.ImagePullPolicy = "IfNotPresent"
	o.debugContainer.TerminationMessagePolicy = "File"

	o.namespace, _, err = o.configFlags.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}

	config, err := o.configFlags.ToRESTConfig()
	if err != nil {
		return err
	}

	o.clientset, err = clientset.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
}

// Validate ensures that all required arguments and flag values are provided
func (o *DebugOptions) Validate() error {
	if len(o.args) < 1 {
		return fmt.Errorf("pod name required")
	}

	return nil
}

// Run lists all available debugs on a user's KUBECONFIG or updates the
// current context based on a provided debug.
func (o *DebugOptions) Run() error {
	r := o.builder.
		WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
		NamespaceParam(o.namespace).DefaultNamespace().ResourceNames("pods", o.args...).
		Do()
	if err := r.Err(); err != nil {
		return fmt.Errorf("error creating builder: %v", err)
	}

	err := r.Visit(func(info *resource.Info, err error) error {
		if err != nil {
			return err
		}
		klog.V(2).Infof("found object for debug: %v", info.Object)

		pods := o.clientset.CoreV1().Pods(info.Namespace)
		ec, err := pods.GetEphemeralContainers(info.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		klog.V(2).Infof("existing ephemeral containers: %v", ec.EphemeralContainers)

		ec.EphemeralContainers = append(ec.EphemeralContainers, o.debugContainer)
		ec, err = pods.UpdateEphemeralContainers(info.Name, ec)
		if err != nil {
			return err
		}

		if o.attach {
			// Wait for the debug container to start
			w, err := info.Watch(ec.ResourceVersion)
			if err != nil {
				return err
			}
			defer w.Stop()

			// TODO: add timeout
			for {
				e := <-w.ResultChan()
				if e.Type != watch.Modified {
					return fmt.Errorf("expected pod to be modified, but instead got event type: %v", e.Type)
				}

				p, ok := e.Object.(*v1.Pod)
				if !ok {
					return fmt.Errorf("watch did not return a pod: %v", e.Object)
				}

				for _, s := range p.Status.EphemeralContainerStatuses {
					if s.Name != o.debugContainer.Name {
						continue
					}

					klog.V(2).Infof("debug container status is %v", s)
					switch {
					case s.State.Terminated != nil:
						return fmt.Errorf("container exited before attach")
					case s.State.Running != nil:
						return o.kubectlAttach(p)
					}
				}
			}

		}

		return nil
	})

	return err
}

func (o *DebugOptions) kubectlAttach(p *v1.Pod) error {
	cmd := exec.Command(
		"kubectl",
		"-n", p.Namespace,
		"attach", p.Name,
		"-c", o.debugContainer.Name,
		fmt.Sprintf("--stdin=%v", o.debugContainer.Stdin),
		fmt.Sprintf("--tty=%v", o.debugContainer.TTY),
	)
	cmd.Stdin = o.In
	cmd.Stdout = o.Out
	cmd.Stderr = o.ErrOut
	return cmd.Run()
}

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

package apply

import (
	"context"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/cli-experimental/internal/pkg/constants"

	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/cli-experimental/internal/pkg/client"
	"sigs.k8s.io/cli-experimental/internal/pkg/clik8s"
	"sigs.k8s.io/cli-experimental/internal/pkg/util"
	"sigs.k8s.io/kustomize/pkg/inventory"
)

// Apply applies directories
type Apply struct {
	// DynamicClient is the client used to talk
	// with the cluster
	DynamicClient client.Client

	// Out stores the output
	Out io.Writer

	// Resources is a list of resource configurations
	Resources clik8s.ResourceConfigs

	// Commit is a git commit object
	Commit *object.Commit
}

// Result contains the Apply Result
type Result struct {
	Resources clik8s.ResourceConfigs
}

// Do executes the apply
func (a *Apply) Do() (Result, error) {
	fmt.Fprintf(a.Out, "Doing `cli-experimental apply`\n")

	// TODO(Liuijngfang1): add a dry-run for all objects
	// When the dry-run passes, proceed to the actual apply

	for _, u := range normalizeResourceOrdering(a.Resources) {
		if util.HasAnnotation(u, inventory.ContentAnnotation) {
			var err error
			u, err = a.updateInventoryObject(u)
			if err != nil {
				fmt.Fprintf(a.Out, "failed to update inventory object %v\n", err)
			}
		}
		if util.MatchAnnotations(u, map[string]string{
			constants.Presence: constants.EnsureNoExist,
		}) {
			// not applying the resource
			continue
		}
		err := a.DynamicClient.Apply(context.Background(), u)
		if err != nil {
			fmt.Fprintf(a.Out, "failed to apply the object: %s/%s: %v\n", u.GetKind(), u.GetName(), err)
			continue
		}
		fmt.Fprintf(a.Out, "applied %s/%s\n", u.GetKind(), u.GetName())
	}
	return Result{Resources: a.Resources}, nil
}

func (a Apply) updateInventoryObject(u *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	obj := u.DeepCopy()
	err := a.DynamicClient.Get(context.Background(),
		types.NamespacedName{Namespace: u.GetNamespace(), Name: u.GetName()}, obj)
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}
	if errors.IsNotFound(err) {
		return u, nil
	}

	oldAnnotation := obj.GetAnnotations()
	newAnnotation := u.GetAnnotations()
	oldhash, okold := oldAnnotation[inventory.HashAnnotation]
	newhash, oknew := newAnnotation[inventory.HashAnnotation]
	if okold && oknew && oldhash == newhash {
		return obj, nil
	}

	return mergeContentAnnotation(u, obj)
}

func mergeContentAnnotation(newObj, oldObj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	newInv := inventory.NewInventory()
	err := newInv.LoadFromAnnotation(newObj.GetAnnotations())
	if err != nil {
		return nil, err
	}
	oldInv := inventory.NewInventory()
	err = oldInv.LoadFromAnnotation(oldObj.GetAnnotations())
	if err != nil {
		return nil, err
	}
	newInv.Previous.Merge(oldInv.Previous)
	newInv.Previous.Merge(oldInv.Current)

	annotations := newObj.GetAnnotations()
	_ = newInv.UpdateAnnotations(annotations)
	newObj.SetAnnotations(annotations)
	return newObj, nil
}

// normalizeResourceOrdering moves the inventory object to be the first resource
func normalizeResourceOrdering(resources clik8s.ResourceConfigs) []*unstructured.Unstructured {
	var results []*unstructured.Unstructured
	index := -1
	for i, u := range resources {
		if util.HasAnnotation(u, inventory.ContentAnnotation) {
			index = i
		} else {
			results = append(results, u)
		}
	}
	if index >= 0 {
		return append([]*unstructured.Unstructured{resources[index]}, results...)
	}
	return results
}

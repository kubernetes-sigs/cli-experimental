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

package prune

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/cli-experimental/internal/pkg/clik8s"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Prune prunes directories
type Prune struct {
	Client    client.Client
	Out       io.Writer
	Resources clik8s.ResourcePruneConfigs
	Commit    *object.Commit
}

// Result contains the Prune Result
type Result struct {
	Resources clik8s.ResourcePruneConfigs
}

// Do executes the prune
func (o *Prune) Do() (Result, error) {
	if len(o.Resources) != 1 {
		return Result{}, fmt.Errorf("prune only accepts one object, but got %v", o.Resources)
	}
	fmt.Fprintf(o.Out, "Doing `cli-experimental prune`\n")

	u := o.Resources[0]
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(u.GroupVersionKind())
	err := o.Client.Get(context.Background(), client.ObjectKey{
		Namespace: u.GetNamespace(),
		Name:      u.GetName(),
	}, obj)

	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Fprintf(os.Stderr, "retrieving current configuration of %s from server for %v", u.GetName(), err)
			return Result{o.Resources}, err
		}
		fmt.Fprintf(os.Stderr, "couldn't find a configmap for pruning %s", u.GetName())
		return Result{}, nil
	}

	hash := u.GetAnnotations()["current"]
	newObj, err := o.runPrune(hash, obj)
	if err != nil {
		return Result{o.Resources}, err
	}

	// update the configmap object
	err = o.Client.Update(context.Background(), newObj)
	if err != nil {
		return Result{}, err
	}
	fmt.Fprintf(o.Out, "%s/%s updated\n", u.GetKind(), u.GetName())
	return Result{Resources: o.Resources}, nil
}

func (o *Prune) runPrune(hash string, cm *unstructured.Unstructured) (*unstructured.Unstructured, error) {

	toKeep := make(map[string]string)
	old := make(map[string]string)

	refOther := make(map[string][]string)
	refByOther := make(map[string][]string)

	toDelete := []string{}

	all := cm.Object["data"].(map[string]interface{})

	//first round, keep all the items that have the same hash
	for key, value := range all {
		h := value.(string)
		if h == hash {
			toKeep[key] = h
			delete(all, key)
		} else if !strings.Contains(key, "---") {
			old[key] = h
		}

		if strings.Contains(key, "---") {
			objs := strings.Split(key, "---")
			if len(objs) != 2 {
				return nil, fmt.Errorf("not able to get 2 objects from %s", key)
			}
			if _, ok := refByOther[objs[0]]; ok {
				refByOther[objs[0]] = append(refByOther[objs[0]], objs[1])
			} else {
				refByOther[objs[0]] = []string{objs[1]}
			}

			if _, ok := refOther[objs[1]]; ok {
				refOther[objs[1]] = append(refOther[objs[1]], objs[0])
			} else {
				refOther[objs[1]] = []string{objs[0]}
			}
		}
	}

	for key := range old {
		if _, ok := refByOther[key]; !ok {
			toDelete = append(toDelete, key)
			if _, ok2 := refOther[key]; ok2 {
				delete(refOther, key)
			}
		}
	}

	for key, value := range old {
		if refs, ok := refByOther[key]; ok {
			found := false
			for _, ref := range refs {
				if isInMap(ref, toKeep) && o.shouldKeep(key, ref) {
					toKeep[key] = value
					toKeep[key+"---"+ref] = value
					found = true
				}
			}
			if !found {
				toDelete = append(toDelete, key)
			}
		}
	}

	err := o.deleteObjects(toDelete)
	if err != nil {
		return nil, err
	}

	cm.Object["data"] = toKeep
	return cm, nil
}

func isInMap(key string, m map[string]string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

func (o *Prune) shouldKeep(obj1, obj2 string) bool {
	values := strings.Split(obj1, "_")
	switch values[1] {
	case "PersistentVolume":
		return true
	case "ConfigMap", "Secret":
		switch v := strings.Split(obj2, "_"); v[1] {
		case "Deployment", "StetefulSet":
			return o.isUsedInDeployment(obj1, obj2)
		default:
			return false
		}
	default:
		return false
	}
}

func (o *Prune) deleteObjects(items []string) error {
	for _, item := range items {
		if strings.Contains(item, "---") {
			continue
		}
		values := strings.Split(item, "_")
		obj := &unstructured.Unstructured{}
		obj.SetGroupVersionKind(schema.GroupVersionKind{
			Group: values[0],
			Kind:  values[1],
		})
		obj.SetNamespace(values[2])
		obj.SetName(values[3])

		err := o.Client.Delete(context.Background(), obj)
		if err != nil {
			return err
		}
		fmt.Fprintf(o.Out, "%s/%s deleted\n", values[1], values[3])
	}
	return nil
}

func (o *Prune) isUsedInDeployment(obj1, obj2 string) bool {
	strings.Split(obj1, "_")
	strings.Split(obj2, "_")

	values1 := strings.Split(obj1, "_")
	values2 := strings.Split(obj2, "_")

	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(schema.GroupVersionKind{
		Group: values2[0],
		Kind:  values2[1],
	})
	err := o.Client.Get(context.Background(), client.ObjectKey{
		Namespace: values2[2],
		Name:      values2[3],
	}, obj)

	if err != nil {
		return false
	}

	ls := []string{}
	for k, v := range obj.GetLabels() {
		ls = append(ls, k+"="+v)
	}

	rsList := &v1.ReplicaSetList{}
	err = o.Client.List(context.Background(), client.MatchingLabels(obj.GetLabels()), rsList)
	if err != nil {
		return false
	}

	content, err := yaml.Marshal(rsList)
	if err != nil {
		fmt.Printf("failed to marshal the rs \n")
		return false
	}
	return bytes.Contains(content, []byte(values1[3]))

	return false
}

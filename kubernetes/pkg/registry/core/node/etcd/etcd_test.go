/*
Copyright 2015 The Kubernetes Authors.

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

package etcd

import (
	"testing"

	"github.com/openshift/kubernetes/pkg/api"
	"github.com/openshift/kubernetes/pkg/api/resource"
	"github.com/openshift/kubernetes/pkg/fields"
	kubeletclient "github.com/openshift/kubernetes/pkg/kubelet/client"
	"github.com/openshift/kubernetes/pkg/labels"
	"github.com/openshift/kubernetes/pkg/registry/generic"
	"github.com/openshift/kubernetes/pkg/registry/registrytest"
	"github.com/openshift/kubernetes/pkg/runtime"
	etcdtesting "github.com/openshift/kubernetes/pkg/storage/etcd/testing"
)

func newStorage(t *testing.T) (*REST, *etcdtesting.EtcdTestServer) {
	etcdStorage, server := registrytest.NewEtcdStorage(t, "")
	restOptions := generic.RESTOptions{StorageConfig: etcdStorage, Decorator: generic.UndecoratedStorage, DeleteCollectionWorkers: 1}
	storage, err := NewStorage(restOptions, kubeletclient.KubeletClientConfig{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	return storage.Node, server
}

func validNewNode() *api.Node {
	return &api.Node{
		ObjectMeta: api.ObjectMeta{
			Name: "foo",
			Labels: map[string]string{
				"name": "foo",
			},
		},
		Spec: api.NodeSpec{
			ExternalID: "external",
		},
		Status: api.NodeStatus{
			Capacity: api.ResourceList{
				api.ResourceName(api.ResourceCPU):    resource.MustParse("10"),
				api.ResourceName(api.ResourceMemory): resource.MustParse("0"),
			},
		},
	}
}

func TestCreate(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store).ClusterScope()
	node := validNewNode()
	node.ObjectMeta = api.ObjectMeta{GenerateName: "foo"}
	test.TestCreate(
		// valid
		node,
		// invalid
		&api.Node{
			ObjectMeta: api.ObjectMeta{Name: "_-a123-a_"},
		},
	)
}

func TestUpdate(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store).ClusterScope()
	test.TestUpdate(
		// valid
		validNewNode(),
		// updateFunc
		func(obj runtime.Object) runtime.Object {
			object := obj.(*api.Node)
			object.Spec.Unschedulable = !object.Spec.Unschedulable
			return object
		},
	)
}

func TestDelete(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store).ClusterScope()
	test.TestDelete(validNewNode())
}

func TestGet(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store).ClusterScope()
	test.TestGet(validNewNode())
}

func TestList(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store).ClusterScope()
	test.TestList(validNewNode())
}

func TestWatch(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store).ClusterScope()
	test.TestWatch(
		validNewNode(),
		// matching labels
		[]labels.Set{
			{"name": "foo"},
		},
		// not matching labels
		[]labels.Set{
			{"name": "bar"},
			{"foo": "bar"},
		},
		// matching fields
		[]fields.Set{
			{"metadata.name": "foo"},
		},
		// not matchin fields
		[]fields.Set{
			{"metadata.name": "bar"},
		},
	)
}

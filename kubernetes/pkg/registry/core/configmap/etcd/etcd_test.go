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
	"github.com/openshift/kubernetes/pkg/fields"
	"github.com/openshift/kubernetes/pkg/labels"
	"github.com/openshift/kubernetes/pkg/registry/generic"
	"github.com/openshift/kubernetes/pkg/registry/registrytest"
	"github.com/openshift/kubernetes/pkg/runtime"
	etcdtesting "github.com/openshift/kubernetes/pkg/storage/etcd/testing"
)

func newStorage(t *testing.T) (*REST, *etcdtesting.EtcdTestServer) {
	etcdStorage, server := registrytest.NewEtcdStorage(t, "")
	restOptions := generic.RESTOptions{StorageConfig: etcdStorage, Decorator: generic.UndecoratedStorage, DeleteCollectionWorkers: 1}
	return NewREST(restOptions), server
}

func validNewConfigMap() *api.ConfigMap {
	return &api.ConfigMap{
		ObjectMeta: api.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
			Labels: map[string]string{
				"label-1": "value-1",
				"label-2": "value-2",
			},
		},
		Data: map[string]string{
			"test": "data",
		},
	}
}

func TestCreate(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store)

	validConfigMap := validNewConfigMap()
	validConfigMap.ObjectMeta = api.ObjectMeta{
		GenerateName: "foo-",
	}

	test.TestCreate(
		validConfigMap,
		&api.ConfigMap{
			ObjectMeta: api.ObjectMeta{Name: "badName"},
			Data: map[string]string{
				"key": "value",
			},
		},
		&api.ConfigMap{
			ObjectMeta: api.ObjectMeta{Name: "name-2"},
			Data: map[string]string{
				"..dotfile": "do: nothing\n",
			},
		},
	)
}

func TestUpdate(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store)
	test.TestUpdate(
		// valid
		validNewConfigMap(),
		// updateFunc
		func(obj runtime.Object) runtime.Object {
			cfg := obj.(*api.ConfigMap)
			cfg.Data["update-test"] = "value"
			return cfg
		},
		// invalid updateFunc
		func(obj runtime.Object) runtime.Object {
			cfg := obj.(*api.ConfigMap)
			cfg.Data["bad*Key"] = "value"
			return cfg
		},
	)
}

func TestDelete(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store)
	test.TestDelete(validNewConfigMap())
}

func TestGet(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store)
	test.TestGet(validNewConfigMap())
}

func TestList(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store)
	test.TestList(validNewConfigMap())
}

func TestWatch(t *testing.T) {
	storage, server := newStorage(t)
	defer server.Terminate(t)
	defer storage.Store.DestroyFunc()
	test := registrytest.New(t, storage.Store)
	test.TestWatch(
		validNewConfigMap(),
		// matching labels
		[]labels.Set{
			{"label-1": "value-1"},
			{"label-2": "value-2"},
		},
		// not matching labels
		[]labels.Set{
			{"foo": "bar"},
		},
		// matching fields
		[]fields.Set{
			{"metadata.namespace": "default"},
			{"metadata.name": "foo"},
		},
		// not matching fields
		[]fields.Set{
			{"metadata.name": "bar"},
			{"name": "foo"},
		},
	)
}

/*
Copyright 2016 The Kubernetes Authors.

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

package app

import (
	"github.com/golang/glog"
	"github.com/openshift/kubernetes/pkg/api"
	"github.com/openshift/kubernetes/pkg/api/rest"
	"github.com/openshift/kubernetes/pkg/apimachinery/registered"
	"github.com/openshift/kubernetes/pkg/apis/extensions"
	_ "github.com/openshift/kubernetes/pkg/apis/extensions/install"
	"github.com/openshift/kubernetes/pkg/genericapiserver"
	daemonsetetcd "github.com/openshift/kubernetes/pkg/registry/extensions/daemonset/etcd"
	deploymentetcd "github.com/openshift/kubernetes/pkg/registry/extensions/deployment/etcd"
	ingressetcd "github.com/openshift/kubernetes/pkg/registry/extensions/ingress/etcd"
	replicasetetcd "github.com/openshift/kubernetes/pkg/registry/extensions/replicaset/etcd"
)

func installExtensionsAPIs(g *genericapiserver.GenericAPIServer, restOptionsFactory restOptionsFactory) {
	replicaSetStorage := replicasetetcd.NewStorage(restOptionsFactory.NewFor(extensions.Resource("replicasets")))
	deploymentStorage := deploymentetcd.NewStorage(restOptionsFactory.NewFor(extensions.Resource("deployments")))
	ingressStorage, ingressStatusStorage := ingressetcd.NewREST(restOptionsFactory.NewFor(extensions.Resource("ingresses")))
	daemonSetStorage, daemonSetStatusStorage := daemonsetetcd.NewREST(restOptionsFactory.NewFor(extensions.Resource("daemonsets")))

	extensionsResources := map[string]rest.Storage{
		"replicasets":          replicaSetStorage.ReplicaSet,
		"replicasets/status":   replicaSetStorage.Status,
		"replicasets/scale":    replicaSetStorage.Scale,
		"ingresses":            ingressStorage,
		"ingresses/status":     ingressStatusStorage,
		"daemonsets":           daemonSetStorage,
		"daemonsets/status":    daemonSetStatusStorage,
		"deployments":          deploymentStorage.Deployment,
		"deployments/status":   deploymentStorage.Status,
		"deployments/scale":    deploymentStorage.Scale,
		"deployments/rollback": deploymentStorage.Rollback,
	}
	extensionsGroupMeta := registered.GroupOrDie(extensions.GroupName)
	apiGroupInfo := genericapiserver.APIGroupInfo{
		GroupMeta: *extensionsGroupMeta,
		VersionedResourcesStorageMap: map[string]map[string]rest.Storage{
			"v1beta1": extensionsResources,
		},
		OptionsExternalVersion: &registered.GroupOrDie(api.GroupName).GroupVersion,
		Scheme:                 api.Scheme,
		ParameterCodec:         api.ParameterCodec,
		NegotiatedSerializer:   api.Codecs,
	}
	if err := g.InstallAPIGroup(&apiGroupInfo); err != nil {
		glog.Fatalf("Error in registering group versions: %v", err)
	}
}

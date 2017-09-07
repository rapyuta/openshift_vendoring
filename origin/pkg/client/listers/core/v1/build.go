// This file was automatically generated by lister-gen with arguments: --input-dirs=[github.com/openshift/origin/pkg/authorization/api,github.com/openshift/origin/pkg/authorization/api/v1,github.com/openshift/origin/pkg/build/api,github.com/openshift/origin/pkg/build/api/v1,github.com/openshift/origin/pkg/deploy/api,github.com/openshift/origin/pkg/deploy/api/v1,github.com/openshift/origin/pkg/image/api,github.com/openshift/origin/pkg/image/api/v1,github.com/openshift/origin/pkg/oauth/api,github.com/openshift/origin/pkg/oauth/api/v1,github.com/openshift/origin/pkg/project/api,github.com/openshift/origin/pkg/project/api/v1,github.com/openshift/origin/pkg/route/api,github.com/openshift/origin/pkg/route/api/v1,github.com/openshift/origin/pkg/sdn/api,github.com/openshift/origin/pkg/sdn/api/v1,github.com/openshift/origin/pkg/template/api,github.com/openshift/origin/pkg/template/api/v1,github.com/openshift/origin/pkg/user/api,github.com/openshift/origin/pkg/user/api/v1] --logtostderr=true

package v1

import (
	api "github.com/openshift/origin/pkg/build/api"
	v1 "github.com/openshift/origin/pkg/build/api/v1"
	"github.com/openshift/kubernetes/pkg/api/errors"
	"github.com/openshift/kubernetes/pkg/client/cache"
	"github.com/openshift/kubernetes/pkg/labels"
)

// BuildLister helps list Builds.
type BuildLister interface {
	// List lists all Builds in the indexer.
	List(selector labels.Selector) (ret []*v1.Build, err error)
	// Builds returns an object that can list and get Builds.
	Builds(namespace string) BuildNamespaceLister
	BuildListerExpansion
}

// buildLister implements the BuildLister interface.
type buildLister struct {
	indexer cache.Indexer
}

// NewBuildLister returns a new BuildLister.
func NewBuildLister(indexer cache.Indexer) BuildLister {
	return &buildLister{indexer: indexer}
}

// List lists all Builds in the indexer.
func (s *buildLister) List(selector labels.Selector) (ret []*v1.Build, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Build))
	})
	return ret, err
}

// Builds returns an object that can list and get Builds.
func (s *buildLister) Builds(namespace string) BuildNamespaceLister {
	return buildNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BuildNamespaceLister helps list and get Builds.
type BuildNamespaceLister interface {
	// List lists all Builds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Build, err error)
	// Get retrieves the Build from the indexer for a given namespace and name.
	Get(name string) (*v1.Build, error)
	BuildNamespaceListerExpansion
}

// buildNamespaceLister implements the BuildNamespaceLister
// interface.
type buildNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Builds in the indexer for a given namespace.
func (s buildNamespaceLister) List(selector labels.Selector) (ret []*v1.Build, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Build))
	})
	return ret, err
}

// Get retrieves the Build from the indexer for a given namespace and name.
func (s buildNamespaceLister) Get(name string) (*v1.Build, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("build"), name)
	}
	return obj.(*v1.Build), nil
}

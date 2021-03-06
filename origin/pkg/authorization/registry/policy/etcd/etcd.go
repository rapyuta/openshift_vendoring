package etcd

import (
	"github.com/openshift/kubernetes/pkg/fields"
	"github.com/openshift/kubernetes/pkg/labels"
	"github.com/openshift/kubernetes/pkg/registry/generic/registry"
	"github.com/openshift/kubernetes/pkg/runtime"
	"github.com/openshift/kubernetes/pkg/storage"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	"github.com/openshift/origin/pkg/authorization/registry/policy"
	"github.com/openshift/origin/pkg/util/restoptions"
)

type REST struct {
	*registry.Store
}

// NewStorage returns a RESTStorage object that will work against Policy objects.
func NewStorage(optsGetter restoptions.Getter) (*REST, error) {
	store := &registry.Store{
		NewFunc:           func() runtime.Object { return &authorizationapi.Policy{} },
		NewListFunc:       func() runtime.Object { return &authorizationapi.PolicyList{} },
		QualifiedResource: authorizationapi.Resource("policies"),
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*authorizationapi.Policy).Name, nil
		},
		PredicateFunc: func(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
			return policy.Matcher(label, field)
		},

		CreateStrategy: policy.Strategy,
		UpdateStrategy: policy.Strategy,
	}

	if err := restoptions.ApplyOptions(optsGetter, store, true, storage.NoTriggerPublisher); err != nil {
		return nil, err
	}

	return &REST{store}, nil
}

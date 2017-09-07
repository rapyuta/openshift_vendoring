package clusterrolebinding

import (
	kapi "github.com/openshift/kubernetes/pkg/api"
	"github.com/openshift/kubernetes/pkg/api/rest"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
)

// Registry is an interface for things that know how to store RoleBindings.
type Registry interface {
	// ListRoleBindings obtains list of policyRoleBindings that match a selector.
	ListRoleBindings(ctx kapi.Context, options *kapi.ListOptions) (*authorizationapi.RoleBindingList, error)
	// GetRoleBinding retrieves a specific policyRoleBinding.
	GetRoleBinding(ctx kapi.Context, id string) (*authorizationapi.RoleBinding, error)
	// CreateRoleBinding creates a new policyRoleBinding.
	CreateRoleBinding(ctx kapi.Context, policyRoleBinding *authorizationapi.RoleBinding) (*authorizationapi.RoleBinding, error)
	// UpdateRoleBinding updates a policyRoleBinding.
	UpdateRoleBinding(ctx kapi.Context, policyRoleBinding *authorizationapi.RoleBinding) (*authorizationapi.RoleBinding, bool, error)
	// DeleteRoleBinding deletes a policyRoleBinding.
	DeleteRoleBinding(ctx kapi.Context, id string) error
}

// Storage is an interface for a standard REST Storage backend
type Storage interface {
	rest.Getter
	rest.Lister
	rest.CreaterUpdater
	rest.GracefulDeleter

	// CreateRoleBinding creates a new policyRoleBinding.  Skipping the escalation check should only be done during bootstrapping procedures where no users are currently bound.
	CreateRoleBindingWithEscalation(ctx kapi.Context, policyRoleBinding *authorizationapi.RoleBinding) (*authorizationapi.RoleBinding, error)
	// UpdateRoleBinding updates a policyRoleBinding.  Skipping the escalation check should only be done during bootstrapping procedures where no users are currently bound.
	UpdateRoleBindingWithEscalation(ctx kapi.Context, policyRoleBinding *authorizationapi.RoleBinding) (*authorizationapi.RoleBinding, bool, error)
}

// storage puts strong typing around storage calls
type storage struct {
	Storage
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched
// types will panic.
func NewRegistry(s Storage) Registry {
	return &storage{s}
}

func (s *storage) ListRoleBindings(ctx kapi.Context, options *kapi.ListOptions) (*authorizationapi.RoleBindingList, error) {
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}

	return obj.(*authorizationapi.RoleBindingList), nil
}

func (s *storage) CreateRoleBinding(ctx kapi.Context, binding *authorizationapi.RoleBinding) (*authorizationapi.RoleBinding, error) {
	obj, err := s.Create(ctx, binding)
	if err != nil {
		return nil, err
	}
	return obj.(*authorizationapi.RoleBinding), err
}

func (s *storage) UpdateRoleBinding(ctx kapi.Context, binding *authorizationapi.RoleBinding) (*authorizationapi.RoleBinding, bool, error) {
	obj, created, err := s.Update(ctx, binding.Name, rest.DefaultUpdatedObjectInfo(binding, kapi.Scheme))
	if err != nil {
		return nil, created, err
	}
	return obj.(*authorizationapi.RoleBinding), created, err
}

func (s *storage) GetRoleBinding(ctx kapi.Context, name string) (*authorizationapi.RoleBinding, error) {
	obj, err := s.Get(ctx, name)
	if err != nil {
		return nil, err
	}
	return obj.(*authorizationapi.RoleBinding), nil
}

func (s *storage) DeleteRoleBinding(ctx kapi.Context, name string) error {
	_, err := s.Delete(ctx, name, nil)
	return err
}

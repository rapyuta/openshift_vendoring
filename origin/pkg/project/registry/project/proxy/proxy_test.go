package proxy

import (
	"strings"
	"testing"

	kapi "github.com/openshift/kubernetes/pkg/api"
	"github.com/openshift/kubernetes/pkg/api/errors"
	"github.com/openshift/kubernetes/pkg/api/unversioned"
	"github.com/openshift/kubernetes/pkg/auth/user"
	"github.com/openshift/kubernetes/pkg/client/clientset_generated/internalclientset/fake"

	oapi "github.com/openshift/origin/pkg/api"
	"github.com/openshift/origin/pkg/project/api"
)

// mockLister returns the namespaces in the list
type mockLister struct {
	namespaceList *kapi.NamespaceList
}

func (ml *mockLister) List(user user.Info) (*kapi.NamespaceList, error) {
	return ml.namespaceList, nil
}

func TestListProjects(t *testing.T) {
	namespaceList := kapi.NamespaceList{
		Items: []kapi.Namespace{
			{
				ObjectMeta: kapi.ObjectMeta{Name: "foo"},
			},
		},
	}
	mockClient := fake.NewSimpleClientset(&namespaceList)
	storage := REST{
		client: mockClient.Core().Namespaces(),
		lister: &mockLister{&namespaceList},
	}
	user := &user.DefaultInfo{
		Name:   "test-user",
		UID:    "test-uid",
		Groups: []string{"test-groups"},
	}
	ctx := kapi.WithUser(kapi.NewContext(), user)
	response, err := storage.List(ctx, nil)
	if err != nil {
		t.Errorf("%#v should be nil.", err)
	}
	projects := response.(*api.ProjectList)
	if len(projects.Items) != 1 {
		t.Errorf("%#v projects.Items should have len 1.", projects.Items)
	}
	responseProject := projects.Items[0]
	if e, r := responseProject.Name, "foo"; e != r {
		t.Errorf("%#v != %#v.", e, r)
	}
}

func TestCreateProjectBadObject(t *testing.T) {
	storage := REST{}

	obj, err := storage.Create(kapi.NewContext(), &api.ProjectList{})
	if obj != nil {
		t.Errorf("Expected nil, got %v", obj)
	}
	if strings.Index(err.Error(), "not a project:") == -1 {
		t.Errorf("Expected 'not an project' error, got %v", err)
	}
}

func TestCreateInvalidProject(t *testing.T) {
	mockClient := &fake.Clientset{}
	storage := NewREST(mockClient.Core().Namespaces(), &mockLister{}, nil, nil)
	_, err := storage.Create(kapi.NewContext(), &api.Project{
		ObjectMeta: kapi.ObjectMeta{
			Annotations: map[string]string{oapi.OpenShiftDisplayName: "h\t\ni"},
		},
	})
	if !errors.IsInvalid(err) {
		t.Errorf("Expected 'invalid' error, got %v", err)
	}
}

func TestCreateProjectOK(t *testing.T) {
	mockClient := &fake.Clientset{}
	storage := NewREST(mockClient.Core().Namespaces(), &mockLister{}, nil, nil)
	_, err := storage.Create(kapi.NewContext(), &api.Project{
		ObjectMeta: kapi.ObjectMeta{Name: "foo"},
	})
	if err != nil {
		t.Errorf("Unexpected non-nil error: %#v", err)
	}
	if len(mockClient.Actions()) != 1 {
		t.Errorf("Expected client action for create")
	}
	if !mockClient.Actions()[0].Matches("create", "namespaces") {
		t.Errorf("Expected call to create-namespace")
	}
}

func TestGetProjectOK(t *testing.T) {
	mockClient := fake.NewSimpleClientset(&kapi.Namespace{ObjectMeta: kapi.ObjectMeta{Name: "foo"}})
	storage := NewREST(mockClient.Core().Namespaces(), &mockLister{}, nil, nil)
	project, err := storage.Get(kapi.NewContext(), "foo")
	if project == nil {
		t.Error("Unexpected nil project")
	}
	if err != nil {
		t.Errorf("Unexpected non-nil error: %v", err)
	}
	if project.(*api.Project).Name != "foo" {
		t.Errorf("Unexpected project: %#v", project)
	}
}

func TestDeleteProject(t *testing.T) {
	mockClient := &fake.Clientset{}
	storage := REST{
		client: mockClient.Core().Namespaces(),
	}
	obj, err := storage.Delete(kapi.NewContext(), "foo")
	if obj == nil {
		t.Error("Unexpected nil obj")
	}
	if err != nil {
		t.Errorf("Unexpected non-nil error: %#v", err)
	}
	status, ok := obj.(*unversioned.Status)
	if !ok {
		t.Errorf("Expected status type, got: %#v", obj)
	}
	if status.Status != unversioned.StatusSuccess {
		t.Errorf("Expected status=success, got: %#v", status)
	}
	if len(mockClient.Actions()) != 1 {
		t.Errorf("Expected client action for delete")
	}
	if !mockClient.Actions()[0].Matches("delete", "namespaces") {
		t.Errorf("Expected call to delete-namespace")
	}
}

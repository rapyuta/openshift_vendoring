package fake

import (
	v1 "github.com/openshift/origin/pkg/project/api/v1"
	api "github.com/openshift/kubernetes/pkg/api"
	unversioned "github.com/openshift/kubernetes/pkg/api/unversioned"
	api_v1 "github.com/openshift/kubernetes/pkg/api/v1"
	core "github.com/openshift/kubernetes/pkg/client/testing/core"
	labels "github.com/openshift/kubernetes/pkg/labels"
	watch "github.com/openshift/kubernetes/pkg/watch"
)

// FakeProjects implements ProjectInterface
type FakeProjects struct {
	Fake *FakeCoreV1
}

var projectsResource = unversioned.GroupVersionResource{Group: "", Version: "v1", Resource: "projects"}

func (c *FakeProjects) Create(project *v1.Project) (result *v1.Project, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootCreateAction(projectsResource, project), &v1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Project), err
}

func (c *FakeProjects) Update(project *v1.Project) (result *v1.Project, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootUpdateAction(projectsResource, project), &v1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Project), err
}

func (c *FakeProjects) UpdateStatus(project *v1.Project) (*v1.Project, error) {
	obj, err := c.Fake.
		Invokes(core.NewRootUpdateSubresourceAction(projectsResource, "status", project), &v1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Project), err
}

func (c *FakeProjects) Delete(name string, options *api_v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewRootDeleteAction(projectsResource, name), &v1.Project{})
	return err
}

func (c *FakeProjects) DeleteCollection(options *api_v1.DeleteOptions, listOptions api_v1.ListOptions) error {
	action := core.NewRootDeleteCollectionAction(projectsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1.ProjectList{})
	return err
}

func (c *FakeProjects) Get(name string) (result *v1.Project, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootGetAction(projectsResource, name), &v1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Project), err
}

func (c *FakeProjects) List(opts api_v1.ListOptions) (result *v1.ProjectList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootListAction(projectsResource, opts), &v1.ProjectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := core.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ProjectList{}
	for _, item := range obj.(*v1.ProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projects.
func (c *FakeProjects) Watch(opts api_v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewRootWatchAction(projectsResource, opts))
}

// Patch applies the patch and returns the patched project.
func (c *FakeProjects) Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *v1.Project, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootPatchSubresourceAction(projectsResource, name, data, subresources...), &v1.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Project), err
}

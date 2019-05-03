package fake

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeRoles struct {
	Fake *FakeAuthorization
	ns   string
}

var rolesResource = schema.GroupVersionResource{Group: "authorization.openshift.io", Version: "", Resource: "roles"}
var rolesKind = schema.GroupVersionKind{Group: "authorization.openshift.io", Version: "", Kind: "Role"}

func (c *FakeRoles) Get(name string, options v1.GetOptions) (result *authorization.Role, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(rolesResource, c.ns, name), &authorization.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.Role), err
}
func (c *FakeRoles) List(opts v1.ListOptions) (result *authorization.RoleList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(rolesResource, rolesKind, c.ns, opts), &authorization.RoleList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &authorization.RoleList{ListMeta: obj.(*authorization.RoleList).ListMeta}
	for _, item := range obj.(*authorization.RoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(rolesResource, c.ns, opts))
}
func (c *FakeRoles) Create(role *authorization.Role) (result *authorization.Role, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(rolesResource, c.ns, role), &authorization.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.Role), err
}
func (c *FakeRoles) Update(role *authorization.Role) (result *authorization.Role, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(rolesResource, c.ns, role), &authorization.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.Role), err
}
func (c *FakeRoles) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(rolesResource, c.ns, name), &authorization.Role{})
	return err
}
func (c *FakeRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(rolesResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &authorization.RoleList{})
	return err
}
func (c *FakeRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authorization.Role, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(rolesResource, c.ns, name, pt, data, subresources...), &authorization.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.Role), err
}

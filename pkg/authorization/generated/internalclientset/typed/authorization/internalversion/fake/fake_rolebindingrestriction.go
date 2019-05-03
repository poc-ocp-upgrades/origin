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

type FakeRoleBindingRestrictions struct {
	Fake *FakeAuthorization
	ns   string
}

var rolebindingrestrictionsResource = schema.GroupVersionResource{Group: "authorization.openshift.io", Version: "", Resource: "rolebindingrestrictions"}
var rolebindingrestrictionsKind = schema.GroupVersionKind{Group: "authorization.openshift.io", Version: "", Kind: "RoleBindingRestriction"}

func (c *FakeRoleBindingRestrictions) Get(name string, options v1.GetOptions) (result *authorization.RoleBindingRestriction, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(rolebindingrestrictionsResource, c.ns, name), &authorization.RoleBindingRestriction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.RoleBindingRestriction), err
}
func (c *FakeRoleBindingRestrictions) List(opts v1.ListOptions) (result *authorization.RoleBindingRestrictionList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(rolebindingrestrictionsResource, rolebindingrestrictionsKind, c.ns, opts), &authorization.RoleBindingRestrictionList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &authorization.RoleBindingRestrictionList{ListMeta: obj.(*authorization.RoleBindingRestrictionList).ListMeta}
	for _, item := range obj.(*authorization.RoleBindingRestrictionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeRoleBindingRestrictions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(rolebindingrestrictionsResource, c.ns, opts))
}
func (c *FakeRoleBindingRestrictions) Create(roleBindingRestriction *authorization.RoleBindingRestriction) (result *authorization.RoleBindingRestriction, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(rolebindingrestrictionsResource, c.ns, roleBindingRestriction), &authorization.RoleBindingRestriction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.RoleBindingRestriction), err
}
func (c *FakeRoleBindingRestrictions) Update(roleBindingRestriction *authorization.RoleBindingRestriction) (result *authorization.RoleBindingRestriction, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(rolebindingrestrictionsResource, c.ns, roleBindingRestriction), &authorization.RoleBindingRestriction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.RoleBindingRestriction), err
}
func (c *FakeRoleBindingRestrictions) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(rolebindingrestrictionsResource, c.ns, name), &authorization.RoleBindingRestriction{})
	return err
}
func (c *FakeRoleBindingRestrictions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(rolebindingrestrictionsResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &authorization.RoleBindingRestrictionList{})
	return err
}
func (c *FakeRoleBindingRestrictions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authorization.RoleBindingRestriction, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(rolebindingrestrictionsResource, c.ns, name, pt, data, subresources...), &authorization.RoleBindingRestriction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.RoleBindingRestriction), err
}

package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 rbac "k8s.io/kubernetes/pkg/apis/rbac"
)

type FakeRoles struct {
 Fake *FakeRbac
 ns   string
}

var rolesResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "", Resource: "roles"}
var rolesKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "", Kind: "Role"}

func (c *FakeRoles) Get(name string, options v1.GetOptions) (result *rbac.Role, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(rolesResource, c.ns, name), &rbac.Role{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.Role), err
}
func (c *FakeRoles) List(opts v1.ListOptions) (result *rbac.RoleList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(rolesResource, rolesKind, c.ns, opts), &rbac.RoleList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &rbac.RoleList{ListMeta: obj.(*rbac.RoleList).ListMeta}
 for _, item := range obj.(*rbac.RoleList).Items {
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
func (c *FakeRoles) Create(role *rbac.Role) (result *rbac.Role, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(rolesResource, c.ns, role), &rbac.Role{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.Role), err
}
func (c *FakeRoles) Update(role *rbac.Role) (result *rbac.Role, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(rolesResource, c.ns, role), &rbac.Role{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.Role), err
}
func (c *FakeRoles) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(rolesResource, c.ns, name), &rbac.Role{})
 return err
}
func (c *FakeRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(rolesResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &rbac.RoleList{})
 return err
}
func (c *FakeRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac.Role, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(rolesResource, c.ns, name, pt, data, subresources...), &rbac.Role{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.Role), err
}

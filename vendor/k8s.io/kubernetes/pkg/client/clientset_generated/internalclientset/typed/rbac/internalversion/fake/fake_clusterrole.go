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

type FakeClusterRoles struct{ Fake *FakeRbac }

var clusterrolesResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "", Resource: "clusterroles"}
var clusterrolesKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "", Kind: "ClusterRole"}

func (c *FakeClusterRoles) Get(name string, options v1.GetOptions) (result *rbac.ClusterRole, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(clusterrolesResource, name), &rbac.ClusterRole{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.ClusterRole), err
}
func (c *FakeClusterRoles) List(opts v1.ListOptions) (result *rbac.ClusterRoleList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(clusterrolesResource, clusterrolesKind, opts), &rbac.ClusterRoleList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &rbac.ClusterRoleList{ListMeta: obj.(*rbac.ClusterRoleList).ListMeta}
 for _, item := range obj.(*rbac.ClusterRoleList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeClusterRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(clusterrolesResource, opts))
}
func (c *FakeClusterRoles) Create(clusterRole *rbac.ClusterRole) (result *rbac.ClusterRole, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(clusterrolesResource, clusterRole), &rbac.ClusterRole{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.ClusterRole), err
}
func (c *FakeClusterRoles) Update(clusterRole *rbac.ClusterRole) (result *rbac.ClusterRole, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(clusterrolesResource, clusterRole), &rbac.ClusterRole{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.ClusterRole), err
}
func (c *FakeClusterRoles) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(clusterrolesResource, name), &rbac.ClusterRole{})
 return err
}
func (c *FakeClusterRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(clusterrolesResource, listOptions)
 _, err := c.Fake.Invokes(action, &rbac.ClusterRoleList{})
 return err
}
func (c *FakeClusterRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac.ClusterRole, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(clusterrolesResource, name, pt, data, subresources...), &rbac.ClusterRole{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.ClusterRole), err
}

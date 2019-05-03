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

type FakeClusterRoleBindings struct{ Fake *FakeRbac }

var clusterrolebindingsResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "", Resource: "clusterrolebindings"}
var clusterrolebindingsKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "", Kind: "ClusterRoleBinding"}

func (c *FakeClusterRoleBindings) Get(name string, options v1.GetOptions) (result *rbac.ClusterRoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(clusterrolebindingsResource, name), &rbac.ClusterRoleBinding{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.ClusterRoleBinding), err
}
func (c *FakeClusterRoleBindings) List(opts v1.ListOptions) (result *rbac.ClusterRoleBindingList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(clusterrolebindingsResource, clusterrolebindingsKind, opts), &rbac.ClusterRoleBindingList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &rbac.ClusterRoleBindingList{ListMeta: obj.(*rbac.ClusterRoleBindingList).ListMeta}
 for _, item := range obj.(*rbac.ClusterRoleBindingList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeClusterRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(clusterrolebindingsResource, opts))
}
func (c *FakeClusterRoleBindings) Create(clusterRoleBinding *rbac.ClusterRoleBinding) (result *rbac.ClusterRoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(clusterrolebindingsResource, clusterRoleBinding), &rbac.ClusterRoleBinding{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.ClusterRoleBinding), err
}
func (c *FakeClusterRoleBindings) Update(clusterRoleBinding *rbac.ClusterRoleBinding) (result *rbac.ClusterRoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(clusterrolebindingsResource, clusterRoleBinding), &rbac.ClusterRoleBinding{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.ClusterRoleBinding), err
}
func (c *FakeClusterRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(clusterrolebindingsResource, name), &rbac.ClusterRoleBinding{})
 return err
}
func (c *FakeClusterRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(clusterrolebindingsResource, listOptions)
 _, err := c.Fake.Invokes(action, &rbac.ClusterRoleBindingList{})
 return err
}
func (c *FakeClusterRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac.ClusterRoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(clusterrolebindingsResource, name, pt, data, subresources...), &rbac.ClusterRoleBinding{})
 if obj == nil {
  return nil, err
 }
 return obj.(*rbac.ClusterRoleBinding), err
}

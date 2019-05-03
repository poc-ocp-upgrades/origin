package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 core "k8s.io/kubernetes/pkg/apis/core"
)

type FakeServiceAccounts struct {
 Fake *FakeCore
 ns   string
}

var serviceaccountsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "serviceaccounts"}
var serviceaccountsKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "ServiceAccount"}

func (c *FakeServiceAccounts) Get(name string, options v1.GetOptions) (result *core.ServiceAccount, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(serviceaccountsResource, c.ns, name), &core.ServiceAccount{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ServiceAccount), err
}
func (c *FakeServiceAccounts) List(opts v1.ListOptions) (result *core.ServiceAccountList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(serviceaccountsResource, serviceaccountsKind, c.ns, opts), &core.ServiceAccountList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.ServiceAccountList{ListMeta: obj.(*core.ServiceAccountList).ListMeta}
 for _, item := range obj.(*core.ServiceAccountList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeServiceAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(serviceaccountsResource, c.ns, opts))
}
func (c *FakeServiceAccounts) Create(serviceAccount *core.ServiceAccount) (result *core.ServiceAccount, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(serviceaccountsResource, c.ns, serviceAccount), &core.ServiceAccount{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ServiceAccount), err
}
func (c *FakeServiceAccounts) Update(serviceAccount *core.ServiceAccount) (result *core.ServiceAccount, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(serviceaccountsResource, c.ns, serviceAccount), &core.ServiceAccount{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ServiceAccount), err
}
func (c *FakeServiceAccounts) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(serviceaccountsResource, c.ns, name), &core.ServiceAccount{})
 return err
}
func (c *FakeServiceAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(serviceaccountsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.ServiceAccountList{})
 return err
}
func (c *FakeServiceAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ServiceAccount, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(serviceaccountsResource, c.ns, name, pt, data, subresources...), &core.ServiceAccount{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ServiceAccount), err
}

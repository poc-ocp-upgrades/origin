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

type FakeResourceQuotas struct {
 Fake *FakeCore
 ns   string
}

var resourcequotasResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "resourcequotas"}
var resourcequotasKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "ResourceQuota"}

func (c *FakeResourceQuotas) Get(name string, options v1.GetOptions) (result *core.ResourceQuota, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(resourcequotasResource, c.ns, name), &core.ResourceQuota{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ResourceQuota), err
}
func (c *FakeResourceQuotas) List(opts v1.ListOptions) (result *core.ResourceQuotaList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(resourcequotasResource, resourcequotasKind, c.ns, opts), &core.ResourceQuotaList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.ResourceQuotaList{ListMeta: obj.(*core.ResourceQuotaList).ListMeta}
 for _, item := range obj.(*core.ResourceQuotaList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeResourceQuotas) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(resourcequotasResource, c.ns, opts))
}
func (c *FakeResourceQuotas) Create(resourceQuota *core.ResourceQuota) (result *core.ResourceQuota, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(resourcequotasResource, c.ns, resourceQuota), &core.ResourceQuota{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ResourceQuota), err
}
func (c *FakeResourceQuotas) Update(resourceQuota *core.ResourceQuota) (result *core.ResourceQuota, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(resourcequotasResource, c.ns, resourceQuota), &core.ResourceQuota{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ResourceQuota), err
}
func (c *FakeResourceQuotas) UpdateStatus(resourceQuota *core.ResourceQuota) (*core.ResourceQuota, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(resourcequotasResource, "status", c.ns, resourceQuota), &core.ResourceQuota{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ResourceQuota), err
}
func (c *FakeResourceQuotas) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(resourcequotasResource, c.ns, name), &core.ResourceQuota{})
 return err
}
func (c *FakeResourceQuotas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(resourcequotasResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.ResourceQuotaList{})
 return err
}
func (c *FakeResourceQuotas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ResourceQuota, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(resourcequotasResource, c.ns, name, pt, data, subresources...), &core.ResourceQuota{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ResourceQuota), err
}

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

type FakeLimitRanges struct {
 Fake *FakeCore
 ns   string
}

var limitrangesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "limitranges"}
var limitrangesKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "LimitRange"}

func (c *FakeLimitRanges) Get(name string, options v1.GetOptions) (result *core.LimitRange, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(limitrangesResource, c.ns, name), &core.LimitRange{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.LimitRange), err
}
func (c *FakeLimitRanges) List(opts v1.ListOptions) (result *core.LimitRangeList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(limitrangesResource, limitrangesKind, c.ns, opts), &core.LimitRangeList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.LimitRangeList{ListMeta: obj.(*core.LimitRangeList).ListMeta}
 for _, item := range obj.(*core.LimitRangeList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeLimitRanges) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(limitrangesResource, c.ns, opts))
}
func (c *FakeLimitRanges) Create(limitRange *core.LimitRange) (result *core.LimitRange, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(limitrangesResource, c.ns, limitRange), &core.LimitRange{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.LimitRange), err
}
func (c *FakeLimitRanges) Update(limitRange *core.LimitRange) (result *core.LimitRange, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(limitrangesResource, c.ns, limitRange), &core.LimitRange{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.LimitRange), err
}
func (c *FakeLimitRanges) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(limitrangesResource, c.ns, name), &core.LimitRange{})
 return err
}
func (c *FakeLimitRanges) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(limitrangesResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.LimitRangeList{})
 return err
}
func (c *FakeLimitRanges) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.LimitRange, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(limitrangesResource, c.ns, name, pt, data, subresources...), &core.LimitRange{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.LimitRange), err
}

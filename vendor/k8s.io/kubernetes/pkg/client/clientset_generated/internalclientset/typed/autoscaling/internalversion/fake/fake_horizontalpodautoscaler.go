package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
)

type FakeHorizontalPodAutoscalers struct {
 Fake *FakeAutoscaling
 ns   string
}

var horizontalpodautoscalersResource = schema.GroupVersionResource{Group: "autoscaling", Version: "", Resource: "horizontalpodautoscalers"}
var horizontalpodautoscalersKind = schema.GroupVersionKind{Group: "autoscaling", Version: "", Kind: "HorizontalPodAutoscaler"}

func (c *FakeHorizontalPodAutoscalers) Get(name string, options v1.GetOptions) (result *autoscaling.HorizontalPodAutoscaler, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(horizontalpodautoscalersResource, c.ns, name), &autoscaling.HorizontalPodAutoscaler{})
 if obj == nil {
  return nil, err
 }
 return obj.(*autoscaling.HorizontalPodAutoscaler), err
}
func (c *FakeHorizontalPodAutoscalers) List(opts v1.ListOptions) (result *autoscaling.HorizontalPodAutoscalerList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(horizontalpodautoscalersResource, horizontalpodautoscalersKind, c.ns, opts), &autoscaling.HorizontalPodAutoscalerList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &autoscaling.HorizontalPodAutoscalerList{ListMeta: obj.(*autoscaling.HorizontalPodAutoscalerList).ListMeta}
 for _, item := range obj.(*autoscaling.HorizontalPodAutoscalerList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeHorizontalPodAutoscalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(horizontalpodautoscalersResource, c.ns, opts))
}
func (c *FakeHorizontalPodAutoscalers) Create(horizontalPodAutoscaler *autoscaling.HorizontalPodAutoscaler) (result *autoscaling.HorizontalPodAutoscaler, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(horizontalpodautoscalersResource, c.ns, horizontalPodAutoscaler), &autoscaling.HorizontalPodAutoscaler{})
 if obj == nil {
  return nil, err
 }
 return obj.(*autoscaling.HorizontalPodAutoscaler), err
}
func (c *FakeHorizontalPodAutoscalers) Update(horizontalPodAutoscaler *autoscaling.HorizontalPodAutoscaler) (result *autoscaling.HorizontalPodAutoscaler, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(horizontalpodautoscalersResource, c.ns, horizontalPodAutoscaler), &autoscaling.HorizontalPodAutoscaler{})
 if obj == nil {
  return nil, err
 }
 return obj.(*autoscaling.HorizontalPodAutoscaler), err
}
func (c *FakeHorizontalPodAutoscalers) UpdateStatus(horizontalPodAutoscaler *autoscaling.HorizontalPodAutoscaler) (*autoscaling.HorizontalPodAutoscaler, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(horizontalpodautoscalersResource, "status", c.ns, horizontalPodAutoscaler), &autoscaling.HorizontalPodAutoscaler{})
 if obj == nil {
  return nil, err
 }
 return obj.(*autoscaling.HorizontalPodAutoscaler), err
}
func (c *FakeHorizontalPodAutoscalers) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(horizontalpodautoscalersResource, c.ns, name), &autoscaling.HorizontalPodAutoscaler{})
 return err
}
func (c *FakeHorizontalPodAutoscalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(horizontalpodautoscalersResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &autoscaling.HorizontalPodAutoscalerList{})
 return err
}
func (c *FakeHorizontalPodAutoscalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *autoscaling.HorizontalPodAutoscaler, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(horizontalpodautoscalersResource, c.ns, name, pt, data, subresources...), &autoscaling.HorizontalPodAutoscaler{})
 if obj == nil {
  return nil, err
 }
 return obj.(*autoscaling.HorizontalPodAutoscaler), err
}

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

type FakePods struct {
 Fake *FakeCore
 ns   string
}

var podsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "pods"}
var podsKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "Pod"}

func (c *FakePods) Get(name string, options v1.GetOptions) (result *core.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(podsResource, c.ns, name), &core.Pod{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Pod), err
}
func (c *FakePods) List(opts v1.ListOptions) (result *core.PodList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(podsResource, podsKind, c.ns, opts), &core.PodList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.PodList{ListMeta: obj.(*core.PodList).ListMeta}
 for _, item := range obj.(*core.PodList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakePods) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(podsResource, c.ns, opts))
}
func (c *FakePods) Create(pod *core.Pod) (result *core.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(podsResource, c.ns, pod), &core.Pod{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Pod), err
}
func (c *FakePods) Update(pod *core.Pod) (result *core.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(podsResource, c.ns, pod), &core.Pod{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Pod), err
}
func (c *FakePods) UpdateStatus(pod *core.Pod) (*core.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(podsResource, "status", c.ns, pod), &core.Pod{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Pod), err
}
func (c *FakePods) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(podsResource, c.ns, name), &core.Pod{})
 return err
}
func (c *FakePods) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(podsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.PodList{})
 return err
}
func (c *FakePods) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(podsResource, c.ns, name, pt, data, subresources...), &core.Pod{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Pod), err
}

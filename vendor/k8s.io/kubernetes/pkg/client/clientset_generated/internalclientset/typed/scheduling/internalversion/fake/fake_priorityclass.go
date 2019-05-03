package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 scheduling "k8s.io/kubernetes/pkg/apis/scheduling"
)

type FakePriorityClasses struct{ Fake *FakeScheduling }

var priorityclassesResource = schema.GroupVersionResource{Group: "scheduling.k8s.io", Version: "", Resource: "priorityclasses"}
var priorityclassesKind = schema.GroupVersionKind{Group: "scheduling.k8s.io", Version: "", Kind: "PriorityClass"}

func (c *FakePriorityClasses) Get(name string, options v1.GetOptions) (result *scheduling.PriorityClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(priorityclassesResource, name), &scheduling.PriorityClass{})
 if obj == nil {
  return nil, err
 }
 return obj.(*scheduling.PriorityClass), err
}
func (c *FakePriorityClasses) List(opts v1.ListOptions) (result *scheduling.PriorityClassList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(priorityclassesResource, priorityclassesKind, opts), &scheduling.PriorityClassList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &scheduling.PriorityClassList{ListMeta: obj.(*scheduling.PriorityClassList).ListMeta}
 for _, item := range obj.(*scheduling.PriorityClassList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakePriorityClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(priorityclassesResource, opts))
}
func (c *FakePriorityClasses) Create(priorityClass *scheduling.PriorityClass) (result *scheduling.PriorityClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(priorityclassesResource, priorityClass), &scheduling.PriorityClass{})
 if obj == nil {
  return nil, err
 }
 return obj.(*scheduling.PriorityClass), err
}
func (c *FakePriorityClasses) Update(priorityClass *scheduling.PriorityClass) (result *scheduling.PriorityClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(priorityclassesResource, priorityClass), &scheduling.PriorityClass{})
 if obj == nil {
  return nil, err
 }
 return obj.(*scheduling.PriorityClass), err
}
func (c *FakePriorityClasses) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(priorityclassesResource, name), &scheduling.PriorityClass{})
 return err
}
func (c *FakePriorityClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(priorityclassesResource, listOptions)
 _, err := c.Fake.Invokes(action, &scheduling.PriorityClassList{})
 return err
}
func (c *FakePriorityClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *scheduling.PriorityClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(priorityclassesResource, name, pt, data, subresources...), &scheduling.PriorityClass{})
 if obj == nil {
  return nil, err
 }
 return obj.(*scheduling.PriorityClass), err
}

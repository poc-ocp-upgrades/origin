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

type FakeComponentStatuses struct{ Fake *FakeCore }

var componentstatusesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "componentstatuses"}
var componentstatusesKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "ComponentStatus"}

func (c *FakeComponentStatuses) Get(name string, options v1.GetOptions) (result *core.ComponentStatus, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(componentstatusesResource, name), &core.ComponentStatus{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ComponentStatus), err
}
func (c *FakeComponentStatuses) List(opts v1.ListOptions) (result *core.ComponentStatusList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(componentstatusesResource, componentstatusesKind, opts), &core.ComponentStatusList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.ComponentStatusList{ListMeta: obj.(*core.ComponentStatusList).ListMeta}
 for _, item := range obj.(*core.ComponentStatusList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeComponentStatuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(componentstatusesResource, opts))
}
func (c *FakeComponentStatuses) Create(componentStatus *core.ComponentStatus) (result *core.ComponentStatus, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(componentstatusesResource, componentStatus), &core.ComponentStatus{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ComponentStatus), err
}
func (c *FakeComponentStatuses) Update(componentStatus *core.ComponentStatus) (result *core.ComponentStatus, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(componentstatusesResource, componentStatus), &core.ComponentStatus{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ComponentStatus), err
}
func (c *FakeComponentStatuses) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(componentstatusesResource, name), &core.ComponentStatus{})
 return err
}
func (c *FakeComponentStatuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(componentstatusesResource, listOptions)
 _, err := c.Fake.Invokes(action, &core.ComponentStatusList{})
 return err
}
func (c *FakeComponentStatuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ComponentStatus, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(componentstatusesResource, name, pt, data, subresources...), &core.ComponentStatus{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ComponentStatus), err
}

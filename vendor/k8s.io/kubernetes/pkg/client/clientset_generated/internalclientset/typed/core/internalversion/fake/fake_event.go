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

type FakeEvents struct {
 Fake *FakeCore
 ns   string
}

var eventsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "events"}
var eventsKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "Event"}

func (c *FakeEvents) Get(name string, options v1.GetOptions) (result *core.Event, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(eventsResource, c.ns, name), &core.Event{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Event), err
}
func (c *FakeEvents) List(opts v1.ListOptions) (result *core.EventList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(eventsResource, eventsKind, c.ns, opts), &core.EventList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.EventList{ListMeta: obj.(*core.EventList).ListMeta}
 for _, item := range obj.(*core.EventList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeEvents) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(eventsResource, c.ns, opts))
}
func (c *FakeEvents) Create(event *core.Event) (result *core.Event, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(eventsResource, c.ns, event), &core.Event{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Event), err
}
func (c *FakeEvents) Update(event *core.Event) (result *core.Event, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(eventsResource, c.ns, event), &core.Event{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Event), err
}
func (c *FakeEvents) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(eventsResource, c.ns, name), &core.Event{})
 return err
}
func (c *FakeEvents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(eventsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.EventList{})
 return err
}
func (c *FakeEvents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Event, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(eventsResource, c.ns, name, pt, data, subresources...), &core.Event{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Event), err
}

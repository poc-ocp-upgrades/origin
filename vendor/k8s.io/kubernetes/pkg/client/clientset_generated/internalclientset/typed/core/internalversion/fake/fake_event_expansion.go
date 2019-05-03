package fake

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/runtime"
 types "k8s.io/apimachinery/pkg/types"
 core "k8s.io/client-go/testing"
 api "k8s.io/kubernetes/pkg/apis/core"
)

func (c *FakeEvents) CreateWithEventNamespace(event *api.Event) (*api.Event, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := core.NewRootCreateAction(eventsResource, event)
 if c.ns != "" {
  action = core.NewCreateAction(eventsResource, c.ns, event)
 }
 obj, err := c.Fake.Invokes(action, event)
 if obj == nil {
  return nil, err
 }
 return obj.(*api.Event), err
}
func (c *FakeEvents) UpdateWithEventNamespace(event *api.Event) (*api.Event, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := core.NewRootUpdateAction(eventsResource, event)
 if c.ns != "" {
  action = core.NewUpdateAction(eventsResource, c.ns, event)
 }
 obj, err := c.Fake.Invokes(action, event)
 if obj == nil {
  return nil, err
 }
 return obj.(*api.Event), err
}
func (c *FakeEvents) PatchWithEventNamespace(event *api.Event, data []byte) (*api.Event, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pt := types.StrategicMergePatchType
 action := core.NewRootPatchAction(eventsResource, event.Name, pt, data)
 if c.ns != "" {
  action = core.NewPatchAction(eventsResource, c.ns, event.Name, pt, data)
 }
 obj, err := c.Fake.Invokes(action, event)
 if obj == nil {
  return nil, err
 }
 return obj.(*api.Event), err
}
func (c *FakeEvents) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*api.EventList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := core.NewRootListAction(eventsResource, eventsKind, metav1.ListOptions{})
 if c.ns != "" {
  action = core.NewListAction(eventsResource, eventsKind, c.ns, metav1.ListOptions{})
 }
 obj, err := c.Fake.Invokes(action, &api.EventList{})
 if obj == nil {
  return nil, err
 }
 return obj.(*api.EventList), err
}
func (c *FakeEvents) GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := core.GenericActionImpl{}
 action.Verb = "get-field-selector"
 action.Resource = eventsResource
 c.Fake.Invokes(action, nil)
 return fields.Everything()
}

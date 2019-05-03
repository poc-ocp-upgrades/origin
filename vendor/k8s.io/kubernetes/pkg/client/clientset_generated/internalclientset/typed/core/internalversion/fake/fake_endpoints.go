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

type FakeEndpoints struct {
 Fake *FakeCore
 ns   string
}

var endpointsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "endpoints"}
var endpointsKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "Endpoints"}

func (c *FakeEndpoints) Get(name string, options v1.GetOptions) (result *core.Endpoints, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(endpointsResource, c.ns, name), &core.Endpoints{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Endpoints), err
}
func (c *FakeEndpoints) List(opts v1.ListOptions) (result *core.EndpointsList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(endpointsResource, endpointsKind, c.ns, opts), &core.EndpointsList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.EndpointsList{ListMeta: obj.(*core.EndpointsList).ListMeta}
 for _, item := range obj.(*core.EndpointsList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(endpointsResource, c.ns, opts))
}
func (c *FakeEndpoints) Create(endpoints *core.Endpoints) (result *core.Endpoints, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(endpointsResource, c.ns, endpoints), &core.Endpoints{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Endpoints), err
}
func (c *FakeEndpoints) Update(endpoints *core.Endpoints) (result *core.Endpoints, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(endpointsResource, c.ns, endpoints), &core.Endpoints{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Endpoints), err
}
func (c *FakeEndpoints) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(endpointsResource, c.ns, name), &core.Endpoints{})
 return err
}
func (c *FakeEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(endpointsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.EndpointsList{})
 return err
}
func (c *FakeEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Endpoints, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(endpointsResource, c.ns, name, pt, data, subresources...), &core.Endpoints{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Endpoints), err
}

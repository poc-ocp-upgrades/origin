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

type FakeServices struct {
 Fake *FakeCore
 ns   string
}

var servicesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "services"}
var servicesKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "Service"}

func (c *FakeServices) Get(name string, options v1.GetOptions) (result *core.Service, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(servicesResource, c.ns, name), &core.Service{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Service), err
}
func (c *FakeServices) List(opts v1.ListOptions) (result *core.ServiceList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(servicesResource, servicesKind, c.ns, opts), &core.ServiceList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.ServiceList{ListMeta: obj.(*core.ServiceList).ListMeta}
 for _, item := range obj.(*core.ServiceList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(servicesResource, c.ns, opts))
}
func (c *FakeServices) Create(service *core.Service) (result *core.Service, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(servicesResource, c.ns, service), &core.Service{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Service), err
}
func (c *FakeServices) Update(service *core.Service) (result *core.Service, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(servicesResource, c.ns, service), &core.Service{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Service), err
}
func (c *FakeServices) UpdateStatus(service *core.Service) (*core.Service, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(servicesResource, "status", c.ns, service), &core.Service{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Service), err
}
func (c *FakeServices) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(servicesResource, c.ns, name), &core.Service{})
 return err
}
func (c *FakeServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(servicesResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.ServiceList{})
 return err
}
func (c *FakeServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Service, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(servicesResource, c.ns, name, pt, data, subresources...), &core.Service{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Service), err
}

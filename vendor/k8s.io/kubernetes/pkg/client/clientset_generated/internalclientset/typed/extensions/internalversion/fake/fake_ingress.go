package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 extensions "k8s.io/kubernetes/pkg/apis/extensions"
)

type FakeIngresses struct {
 Fake *FakeExtensions
 ns   string
}

var ingressesResource = schema.GroupVersionResource{Group: "extensions", Version: "", Resource: "ingresses"}
var ingressesKind = schema.GroupVersionKind{Group: "extensions", Version: "", Kind: "Ingress"}

func (c *FakeIngresses) Get(name string, options v1.GetOptions) (result *extensions.Ingress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(ingressesResource, c.ns, name), &extensions.Ingress{})
 if obj == nil {
  return nil, err
 }
 return obj.(*extensions.Ingress), err
}
func (c *FakeIngresses) List(opts v1.ListOptions) (result *extensions.IngressList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(ingressesResource, ingressesKind, c.ns, opts), &extensions.IngressList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &extensions.IngressList{ListMeta: obj.(*extensions.IngressList).ListMeta}
 for _, item := range obj.(*extensions.IngressList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeIngresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(ingressesResource, c.ns, opts))
}
func (c *FakeIngresses) Create(ingress *extensions.Ingress) (result *extensions.Ingress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(ingressesResource, c.ns, ingress), &extensions.Ingress{})
 if obj == nil {
  return nil, err
 }
 return obj.(*extensions.Ingress), err
}
func (c *FakeIngresses) Update(ingress *extensions.Ingress) (result *extensions.Ingress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(ingressesResource, c.ns, ingress), &extensions.Ingress{})
 if obj == nil {
  return nil, err
 }
 return obj.(*extensions.Ingress), err
}
func (c *FakeIngresses) UpdateStatus(ingress *extensions.Ingress) (*extensions.Ingress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(ingressesResource, "status", c.ns, ingress), &extensions.Ingress{})
 if obj == nil {
  return nil, err
 }
 return obj.(*extensions.Ingress), err
}
func (c *FakeIngresses) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(ingressesResource, c.ns, name), &extensions.Ingress{})
 return err
}
func (c *FakeIngresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(ingressesResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &extensions.IngressList{})
 return err
}
func (c *FakeIngresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *extensions.Ingress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(ingressesResource, c.ns, name, pt, data, subresources...), &extensions.Ingress{})
 if obj == nil {
  return nil, err
 }
 return obj.(*extensions.Ingress), err
}

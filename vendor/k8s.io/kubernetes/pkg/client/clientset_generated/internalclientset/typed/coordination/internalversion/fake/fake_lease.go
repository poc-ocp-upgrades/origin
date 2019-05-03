package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 coordination "k8s.io/kubernetes/pkg/apis/coordination"
)

type FakeLeases struct {
 Fake *FakeCoordination
 ns   string
}

var leasesResource = schema.GroupVersionResource{Group: "coordination.k8s.io", Version: "", Resource: "leases"}
var leasesKind = schema.GroupVersionKind{Group: "coordination.k8s.io", Version: "", Kind: "Lease"}

func (c *FakeLeases) Get(name string, options v1.GetOptions) (result *coordination.Lease, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(leasesResource, c.ns, name), &coordination.Lease{})
 if obj == nil {
  return nil, err
 }
 return obj.(*coordination.Lease), err
}
func (c *FakeLeases) List(opts v1.ListOptions) (result *coordination.LeaseList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(leasesResource, leasesKind, c.ns, opts), &coordination.LeaseList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &coordination.LeaseList{ListMeta: obj.(*coordination.LeaseList).ListMeta}
 for _, item := range obj.(*coordination.LeaseList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeLeases) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(leasesResource, c.ns, opts))
}
func (c *FakeLeases) Create(lease *coordination.Lease) (result *coordination.Lease, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(leasesResource, c.ns, lease), &coordination.Lease{})
 if obj == nil {
  return nil, err
 }
 return obj.(*coordination.Lease), err
}
func (c *FakeLeases) Update(lease *coordination.Lease) (result *coordination.Lease, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(leasesResource, c.ns, lease), &coordination.Lease{})
 if obj == nil {
  return nil, err
 }
 return obj.(*coordination.Lease), err
}
func (c *FakeLeases) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(leasesResource, c.ns, name), &coordination.Lease{})
 return err
}
func (c *FakeLeases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(leasesResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &coordination.LeaseList{})
 return err
}
func (c *FakeLeases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *coordination.Lease, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(leasesResource, c.ns, name, pt, data, subresources...), &coordination.Lease{})
 if obj == nil {
  return nil, err
 }
 return obj.(*coordination.Lease), err
}

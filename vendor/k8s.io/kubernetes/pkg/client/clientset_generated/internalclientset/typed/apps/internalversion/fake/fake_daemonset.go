package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 apps "k8s.io/kubernetes/pkg/apis/apps"
)

type FakeDaemonSets struct {
 Fake *FakeApps
 ns   string
}

var daemonsetsResource = schema.GroupVersionResource{Group: "apps", Version: "", Resource: "daemonsets"}
var daemonsetsKind = schema.GroupVersionKind{Group: "apps", Version: "", Kind: "DaemonSet"}

func (c *FakeDaemonSets) Get(name string, options v1.GetOptions) (result *apps.DaemonSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(daemonsetsResource, c.ns, name), &apps.DaemonSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.DaemonSet), err
}
func (c *FakeDaemonSets) List(opts v1.ListOptions) (result *apps.DaemonSetList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(daemonsetsResource, daemonsetsKind, c.ns, opts), &apps.DaemonSetList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &apps.DaemonSetList{ListMeta: obj.(*apps.DaemonSetList).ListMeta}
 for _, item := range obj.(*apps.DaemonSetList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeDaemonSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(daemonsetsResource, c.ns, opts))
}
func (c *FakeDaemonSets) Create(daemonSet *apps.DaemonSet) (result *apps.DaemonSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(daemonsetsResource, c.ns, daemonSet), &apps.DaemonSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.DaemonSet), err
}
func (c *FakeDaemonSets) Update(daemonSet *apps.DaemonSet) (result *apps.DaemonSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(daemonsetsResource, c.ns, daemonSet), &apps.DaemonSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.DaemonSet), err
}
func (c *FakeDaemonSets) UpdateStatus(daemonSet *apps.DaemonSet) (*apps.DaemonSet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(daemonsetsResource, "status", c.ns, daemonSet), &apps.DaemonSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.DaemonSet), err
}
func (c *FakeDaemonSets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(daemonsetsResource, c.ns, name), &apps.DaemonSet{})
 return err
}
func (c *FakeDaemonSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(daemonsetsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &apps.DaemonSetList{})
 return err
}
func (c *FakeDaemonSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.DaemonSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(daemonsetsResource, c.ns, name, pt, data, subresources...), &apps.DaemonSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.DaemonSet), err
}

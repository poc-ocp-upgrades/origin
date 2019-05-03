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

type FakeStatefulSets struct {
 Fake *FakeApps
 ns   string
}

var statefulsetsResource = schema.GroupVersionResource{Group: "apps", Version: "", Resource: "statefulsets"}
var statefulsetsKind = schema.GroupVersionKind{Group: "apps", Version: "", Kind: "StatefulSet"}

func (c *FakeStatefulSets) Get(name string, options v1.GetOptions) (result *apps.StatefulSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(statefulsetsResource, c.ns, name), &apps.StatefulSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.StatefulSet), err
}
func (c *FakeStatefulSets) List(opts v1.ListOptions) (result *apps.StatefulSetList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(statefulsetsResource, statefulsetsKind, c.ns, opts), &apps.StatefulSetList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &apps.StatefulSetList{ListMeta: obj.(*apps.StatefulSetList).ListMeta}
 for _, item := range obj.(*apps.StatefulSetList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeStatefulSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(statefulsetsResource, c.ns, opts))
}
func (c *FakeStatefulSets) Create(statefulSet *apps.StatefulSet) (result *apps.StatefulSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(statefulsetsResource, c.ns, statefulSet), &apps.StatefulSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.StatefulSet), err
}
func (c *FakeStatefulSets) Update(statefulSet *apps.StatefulSet) (result *apps.StatefulSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(statefulsetsResource, c.ns, statefulSet), &apps.StatefulSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.StatefulSet), err
}
func (c *FakeStatefulSets) UpdateStatus(statefulSet *apps.StatefulSet) (*apps.StatefulSet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(statefulsetsResource, "status", c.ns, statefulSet), &apps.StatefulSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.StatefulSet), err
}
func (c *FakeStatefulSets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(statefulsetsResource, c.ns, name), &apps.StatefulSet{})
 return err
}
func (c *FakeStatefulSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(statefulsetsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &apps.StatefulSetList{})
 return err
}
func (c *FakeStatefulSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.StatefulSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(statefulsetsResource, c.ns, name, pt, data, subresources...), &apps.StatefulSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.StatefulSet), err
}

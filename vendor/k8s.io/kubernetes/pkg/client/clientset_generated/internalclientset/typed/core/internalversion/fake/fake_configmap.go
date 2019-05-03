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

type FakeConfigMaps struct {
 Fake *FakeCore
 ns   string
}

var configmapsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "configmaps"}
var configmapsKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "ConfigMap"}

func (c *FakeConfigMaps) Get(name string, options v1.GetOptions) (result *core.ConfigMap, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(configmapsResource, c.ns, name), &core.ConfigMap{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ConfigMap), err
}
func (c *FakeConfigMaps) List(opts v1.ListOptions) (result *core.ConfigMapList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(configmapsResource, configmapsKind, c.ns, opts), &core.ConfigMapList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.ConfigMapList{ListMeta: obj.(*core.ConfigMapList).ListMeta}
 for _, item := range obj.(*core.ConfigMapList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeConfigMaps) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(configmapsResource, c.ns, opts))
}
func (c *FakeConfigMaps) Create(configMap *core.ConfigMap) (result *core.ConfigMap, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(configmapsResource, c.ns, configMap), &core.ConfigMap{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ConfigMap), err
}
func (c *FakeConfigMaps) Update(configMap *core.ConfigMap) (result *core.ConfigMap, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(configmapsResource, c.ns, configMap), &core.ConfigMap{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ConfigMap), err
}
func (c *FakeConfigMaps) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(configmapsResource, c.ns, name), &core.ConfigMap{})
 return err
}
func (c *FakeConfigMaps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(configmapsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.ConfigMapList{})
 return err
}
func (c *FakeConfigMaps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ConfigMap, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(configmapsResource, c.ns, name, pt, data, subresources...), &core.ConfigMap{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ConfigMap), err
}

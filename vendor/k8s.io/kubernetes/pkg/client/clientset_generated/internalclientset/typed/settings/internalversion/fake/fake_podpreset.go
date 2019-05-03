package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 settings "k8s.io/kubernetes/pkg/apis/settings"
)

type FakePodPresets struct {
 Fake *FakeSettings
 ns   string
}

var podpresetsResource = schema.GroupVersionResource{Group: "settings.k8s.io", Version: "", Resource: "podpresets"}
var podpresetsKind = schema.GroupVersionKind{Group: "settings.k8s.io", Version: "", Kind: "PodPreset"}

func (c *FakePodPresets) Get(name string, options v1.GetOptions) (result *settings.PodPreset, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(podpresetsResource, c.ns, name), &settings.PodPreset{})
 if obj == nil {
  return nil, err
 }
 return obj.(*settings.PodPreset), err
}
func (c *FakePodPresets) List(opts v1.ListOptions) (result *settings.PodPresetList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(podpresetsResource, podpresetsKind, c.ns, opts), &settings.PodPresetList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &settings.PodPresetList{ListMeta: obj.(*settings.PodPresetList).ListMeta}
 for _, item := range obj.(*settings.PodPresetList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakePodPresets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(podpresetsResource, c.ns, opts))
}
func (c *FakePodPresets) Create(podPreset *settings.PodPreset) (result *settings.PodPreset, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(podpresetsResource, c.ns, podPreset), &settings.PodPreset{})
 if obj == nil {
  return nil, err
 }
 return obj.(*settings.PodPreset), err
}
func (c *FakePodPresets) Update(podPreset *settings.PodPreset) (result *settings.PodPreset, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(podpresetsResource, c.ns, podPreset), &settings.PodPreset{})
 if obj == nil {
  return nil, err
 }
 return obj.(*settings.PodPreset), err
}
func (c *FakePodPresets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(podpresetsResource, c.ns, name), &settings.PodPreset{})
 return err
}
func (c *FakePodPresets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(podpresetsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &settings.PodPresetList{})
 return err
}
func (c *FakePodPresets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *settings.PodPreset, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(podpresetsResource, c.ns, name, pt, data, subresources...), &settings.PodPreset{})
 if obj == nil {
  return nil, err
 }
 return obj.(*settings.PodPreset), err
}

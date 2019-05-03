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

type FakePersistentVolumes struct{ Fake *FakeCore }

var persistentvolumesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "persistentvolumes"}
var persistentvolumesKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "PersistentVolume"}

func (c *FakePersistentVolumes) Get(name string, options v1.GetOptions) (result *core.PersistentVolume, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(persistentvolumesResource, name), &core.PersistentVolume{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PersistentVolume), err
}
func (c *FakePersistentVolumes) List(opts v1.ListOptions) (result *core.PersistentVolumeList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(persistentvolumesResource, persistentvolumesKind, opts), &core.PersistentVolumeList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.PersistentVolumeList{ListMeta: obj.(*core.PersistentVolumeList).ListMeta}
 for _, item := range obj.(*core.PersistentVolumeList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakePersistentVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(persistentvolumesResource, opts))
}
func (c *FakePersistentVolumes) Create(persistentVolume *core.PersistentVolume) (result *core.PersistentVolume, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(persistentvolumesResource, persistentVolume), &core.PersistentVolume{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PersistentVolume), err
}
func (c *FakePersistentVolumes) Update(persistentVolume *core.PersistentVolume) (result *core.PersistentVolume, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(persistentvolumesResource, persistentVolume), &core.PersistentVolume{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PersistentVolume), err
}
func (c *FakePersistentVolumes) UpdateStatus(persistentVolume *core.PersistentVolume) (*core.PersistentVolume, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(persistentvolumesResource, "status", persistentVolume), &core.PersistentVolume{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PersistentVolume), err
}
func (c *FakePersistentVolumes) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(persistentvolumesResource, name), &core.PersistentVolume{})
 return err
}
func (c *FakePersistentVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(persistentvolumesResource, listOptions)
 _, err := c.Fake.Invokes(action, &core.PersistentVolumeList{})
 return err
}
func (c *FakePersistentVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.PersistentVolume, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(persistentvolumesResource, name, pt, data, subresources...), &core.PersistentVolume{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PersistentVolume), err
}

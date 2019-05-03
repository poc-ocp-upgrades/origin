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

type FakePersistentVolumeClaims struct {
 Fake *FakeCore
 ns   string
}

var persistentvolumeclaimsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "persistentvolumeclaims"}
var persistentvolumeclaimsKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "PersistentVolumeClaim"}

func (c *FakePersistentVolumeClaims) Get(name string, options v1.GetOptions) (result *core.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(persistentvolumeclaimsResource, c.ns, name), &core.PersistentVolumeClaim{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PersistentVolumeClaim), err
}
func (c *FakePersistentVolumeClaims) List(opts v1.ListOptions) (result *core.PersistentVolumeClaimList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(persistentvolumeclaimsResource, persistentvolumeclaimsKind, c.ns, opts), &core.PersistentVolumeClaimList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.PersistentVolumeClaimList{ListMeta: obj.(*core.PersistentVolumeClaimList).ListMeta}
 for _, item := range obj.(*core.PersistentVolumeClaimList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakePersistentVolumeClaims) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(persistentvolumeclaimsResource, c.ns, opts))
}
func (c *FakePersistentVolumeClaims) Create(persistentVolumeClaim *core.PersistentVolumeClaim) (result *core.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(persistentvolumeclaimsResource, c.ns, persistentVolumeClaim), &core.PersistentVolumeClaim{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PersistentVolumeClaim), err
}
func (c *FakePersistentVolumeClaims) Update(persistentVolumeClaim *core.PersistentVolumeClaim) (result *core.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(persistentvolumeclaimsResource, c.ns, persistentVolumeClaim), &core.PersistentVolumeClaim{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PersistentVolumeClaim), err
}
func (c *FakePersistentVolumeClaims) UpdateStatus(persistentVolumeClaim *core.PersistentVolumeClaim) (*core.PersistentVolumeClaim, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(persistentvolumeclaimsResource, "status", c.ns, persistentVolumeClaim), &core.PersistentVolumeClaim{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PersistentVolumeClaim), err
}
func (c *FakePersistentVolumeClaims) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(persistentvolumeclaimsResource, c.ns, name), &core.PersistentVolumeClaim{})
 return err
}
func (c *FakePersistentVolumeClaims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(persistentvolumeclaimsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.PersistentVolumeClaimList{})
 return err
}
func (c *FakePersistentVolumeClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(persistentvolumeclaimsResource, c.ns, name, pt, data, subresources...), &core.PersistentVolumeClaim{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PersistentVolumeClaim), err
}

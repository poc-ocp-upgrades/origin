package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 storage "k8s.io/kubernetes/pkg/apis/storage"
)

type FakeStorageClasses struct{ Fake *FakeStorage }

var storageclassesResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "", Resource: "storageclasses"}
var storageclassesKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "", Kind: "StorageClass"}

func (c *FakeStorageClasses) Get(name string, options v1.GetOptions) (result *storage.StorageClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(storageclassesResource, name), &storage.StorageClass{})
 if obj == nil {
  return nil, err
 }
 return obj.(*storage.StorageClass), err
}
func (c *FakeStorageClasses) List(opts v1.ListOptions) (result *storage.StorageClassList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(storageclassesResource, storageclassesKind, opts), &storage.StorageClassList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &storage.StorageClassList{ListMeta: obj.(*storage.StorageClassList).ListMeta}
 for _, item := range obj.(*storage.StorageClassList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeStorageClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(storageclassesResource, opts))
}
func (c *FakeStorageClasses) Create(storageClass *storage.StorageClass) (result *storage.StorageClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(storageclassesResource, storageClass), &storage.StorageClass{})
 if obj == nil {
  return nil, err
 }
 return obj.(*storage.StorageClass), err
}
func (c *FakeStorageClasses) Update(storageClass *storage.StorageClass) (result *storage.StorageClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(storageclassesResource, storageClass), &storage.StorageClass{})
 if obj == nil {
  return nil, err
 }
 return obj.(*storage.StorageClass), err
}
func (c *FakeStorageClasses) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(storageclassesResource, name), &storage.StorageClass{})
 return err
}
func (c *FakeStorageClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(storageclassesResource, listOptions)
 _, err := c.Fake.Invokes(action, &storage.StorageClassList{})
 return err
}
func (c *FakeStorageClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *storage.StorageClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(storageclassesResource, name, pt, data, subresources...), &storage.StorageClass{})
 if obj == nil {
  return nil, err
 }
 return obj.(*storage.StorageClass), err
}

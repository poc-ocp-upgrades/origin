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

type FakeVolumeAttachments struct{ Fake *FakeStorage }

var volumeattachmentsResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "", Resource: "volumeattachments"}
var volumeattachmentsKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "", Kind: "VolumeAttachment"}

func (c *FakeVolumeAttachments) Get(name string, options v1.GetOptions) (result *storage.VolumeAttachment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(volumeattachmentsResource, name), &storage.VolumeAttachment{})
 if obj == nil {
  return nil, err
 }
 return obj.(*storage.VolumeAttachment), err
}
func (c *FakeVolumeAttachments) List(opts v1.ListOptions) (result *storage.VolumeAttachmentList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(volumeattachmentsResource, volumeattachmentsKind, opts), &storage.VolumeAttachmentList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &storage.VolumeAttachmentList{ListMeta: obj.(*storage.VolumeAttachmentList).ListMeta}
 for _, item := range obj.(*storage.VolumeAttachmentList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeVolumeAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(volumeattachmentsResource, opts))
}
func (c *FakeVolumeAttachments) Create(volumeAttachment *storage.VolumeAttachment) (result *storage.VolumeAttachment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(volumeattachmentsResource, volumeAttachment), &storage.VolumeAttachment{})
 if obj == nil {
  return nil, err
 }
 return obj.(*storage.VolumeAttachment), err
}
func (c *FakeVolumeAttachments) Update(volumeAttachment *storage.VolumeAttachment) (result *storage.VolumeAttachment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(volumeattachmentsResource, volumeAttachment), &storage.VolumeAttachment{})
 if obj == nil {
  return nil, err
 }
 return obj.(*storage.VolumeAttachment), err
}
func (c *FakeVolumeAttachments) UpdateStatus(volumeAttachment *storage.VolumeAttachment) (*storage.VolumeAttachment, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(volumeattachmentsResource, "status", volumeAttachment), &storage.VolumeAttachment{})
 if obj == nil {
  return nil, err
 }
 return obj.(*storage.VolumeAttachment), err
}
func (c *FakeVolumeAttachments) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(volumeattachmentsResource, name), &storage.VolumeAttachment{})
 return err
}
func (c *FakeVolumeAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(volumeattachmentsResource, listOptions)
 _, err := c.Fake.Invokes(action, &storage.VolumeAttachmentList{})
 return err
}
func (c *FakeVolumeAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *storage.VolumeAttachment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(volumeattachmentsResource, name, pt, data, subresources...), &storage.VolumeAttachment{})
 if obj == nil {
  return nil, err
 }
 return obj.(*storage.VolumeAttachment), err
}

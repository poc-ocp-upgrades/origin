package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 storage "k8s.io/kubernetes/pkg/apis/storage"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type VolumeAttachmentsGetter interface {
 VolumeAttachments() VolumeAttachmentInterface
}
type VolumeAttachmentInterface interface {
 Create(*storage.VolumeAttachment) (*storage.VolumeAttachment, error)
 Update(*storage.VolumeAttachment) (*storage.VolumeAttachment, error)
 UpdateStatus(*storage.VolumeAttachment) (*storage.VolumeAttachment, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*storage.VolumeAttachment, error)
 List(opts v1.ListOptions) (*storage.VolumeAttachmentList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *storage.VolumeAttachment, err error)
 VolumeAttachmentExpansion
}
type volumeAttachments struct{ client rest.Interface }

func newVolumeAttachments(c *StorageClient) *volumeAttachments {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &volumeAttachments{client: c.RESTClient()}
}
func (c *volumeAttachments) Get(name string, options v1.GetOptions) (result *storage.VolumeAttachment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &storage.VolumeAttachment{}
 err = c.client.Get().Resource("volumeattachments").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *volumeAttachments) List(opts v1.ListOptions) (result *storage.VolumeAttachmentList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &storage.VolumeAttachmentList{}
 err = c.client.Get().Resource("volumeattachments").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *volumeAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("volumeattachments").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *volumeAttachments) Create(volumeAttachment *storage.VolumeAttachment) (result *storage.VolumeAttachment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &storage.VolumeAttachment{}
 err = c.client.Post().Resource("volumeattachments").Body(volumeAttachment).Do().Into(result)
 return
}
func (c *volumeAttachments) Update(volumeAttachment *storage.VolumeAttachment) (result *storage.VolumeAttachment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &storage.VolumeAttachment{}
 err = c.client.Put().Resource("volumeattachments").Name(volumeAttachment.Name).Body(volumeAttachment).Do().Into(result)
 return
}
func (c *volumeAttachments) UpdateStatus(volumeAttachment *storage.VolumeAttachment) (result *storage.VolumeAttachment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &storage.VolumeAttachment{}
 err = c.client.Put().Resource("volumeattachments").Name(volumeAttachment.Name).SubResource("status").Body(volumeAttachment).Do().Into(result)
 return
}
func (c *volumeAttachments) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("volumeattachments").Name(name).Body(options).Do().Error()
}
func (c *volumeAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("volumeattachments").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *volumeAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *storage.VolumeAttachment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &storage.VolumeAttachment{}
 err = c.client.Patch(pt).Resource("volumeattachments").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

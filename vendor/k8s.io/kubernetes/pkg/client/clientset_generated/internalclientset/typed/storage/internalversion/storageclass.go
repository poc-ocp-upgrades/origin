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

type StorageClassesGetter interface{ StorageClasses() StorageClassInterface }
type StorageClassInterface interface {
 Create(*storage.StorageClass) (*storage.StorageClass, error)
 Update(*storage.StorageClass) (*storage.StorageClass, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*storage.StorageClass, error)
 List(opts v1.ListOptions) (*storage.StorageClassList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *storage.StorageClass, err error)
 StorageClassExpansion
}
type storageClasses struct{ client rest.Interface }

func newStorageClasses(c *StorageClient) *storageClasses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &storageClasses{client: c.RESTClient()}
}
func (c *storageClasses) Get(name string, options v1.GetOptions) (result *storage.StorageClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &storage.StorageClass{}
 err = c.client.Get().Resource("storageclasses").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *storageClasses) List(opts v1.ListOptions) (result *storage.StorageClassList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &storage.StorageClassList{}
 err = c.client.Get().Resource("storageclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *storageClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("storageclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *storageClasses) Create(storageClass *storage.StorageClass) (result *storage.StorageClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &storage.StorageClass{}
 err = c.client.Post().Resource("storageclasses").Body(storageClass).Do().Into(result)
 return
}
func (c *storageClasses) Update(storageClass *storage.StorageClass) (result *storage.StorageClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &storage.StorageClass{}
 err = c.client.Put().Resource("storageclasses").Name(storageClass.Name).Body(storageClass).Do().Into(result)
 return
}
func (c *storageClasses) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("storageclasses").Name(name).Body(options).Do().Error()
}
func (c *storageClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("storageclasses").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *storageClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *storage.StorageClass, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &storage.StorageClass{}
 err = c.client.Patch(pt).Resource("storageclasses").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

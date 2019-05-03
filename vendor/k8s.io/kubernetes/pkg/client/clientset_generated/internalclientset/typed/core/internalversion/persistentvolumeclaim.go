package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 core "k8s.io/kubernetes/pkg/apis/core"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type PersistentVolumeClaimsGetter interface {
 PersistentVolumeClaims(namespace string) PersistentVolumeClaimInterface
}
type PersistentVolumeClaimInterface interface {
 Create(*core.PersistentVolumeClaim) (*core.PersistentVolumeClaim, error)
 Update(*core.PersistentVolumeClaim) (*core.PersistentVolumeClaim, error)
 UpdateStatus(*core.PersistentVolumeClaim) (*core.PersistentVolumeClaim, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.PersistentVolumeClaim, error)
 List(opts v1.ListOptions) (*core.PersistentVolumeClaimList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.PersistentVolumeClaim, err error)
 PersistentVolumeClaimExpansion
}
type persistentVolumeClaims struct {
 client rest.Interface
 ns     string
}

func newPersistentVolumeClaims(c *CoreClient, namespace string) *persistentVolumeClaims {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &persistentVolumeClaims{client: c.RESTClient(), ns: namespace}
}
func (c *persistentVolumeClaims) Get(name string, options v1.GetOptions) (result *core.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PersistentVolumeClaim{}
 err = c.client.Get().Namespace(c.ns).Resource("persistentvolumeclaims").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *persistentVolumeClaims) List(opts v1.ListOptions) (result *core.PersistentVolumeClaimList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.PersistentVolumeClaimList{}
 err = c.client.Get().Namespace(c.ns).Resource("persistentvolumeclaims").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *persistentVolumeClaims) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("persistentvolumeclaims").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *persistentVolumeClaims) Create(persistentVolumeClaim *core.PersistentVolumeClaim) (result *core.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PersistentVolumeClaim{}
 err = c.client.Post().Namespace(c.ns).Resource("persistentvolumeclaims").Body(persistentVolumeClaim).Do().Into(result)
 return
}
func (c *persistentVolumeClaims) Update(persistentVolumeClaim *core.PersistentVolumeClaim) (result *core.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PersistentVolumeClaim{}
 err = c.client.Put().Namespace(c.ns).Resource("persistentvolumeclaims").Name(persistentVolumeClaim.Name).Body(persistentVolumeClaim).Do().Into(result)
 return
}
func (c *persistentVolumeClaims) UpdateStatus(persistentVolumeClaim *core.PersistentVolumeClaim) (result *core.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PersistentVolumeClaim{}
 err = c.client.Put().Namespace(c.ns).Resource("persistentvolumeclaims").Name(persistentVolumeClaim.Name).SubResource("status").Body(persistentVolumeClaim).Do().Into(result)
 return
}
func (c *persistentVolumeClaims) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("persistentvolumeclaims").Name(name).Body(options).Do().Error()
}
func (c *persistentVolumeClaims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("persistentvolumeclaims").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *persistentVolumeClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PersistentVolumeClaim{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("persistentvolumeclaims").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

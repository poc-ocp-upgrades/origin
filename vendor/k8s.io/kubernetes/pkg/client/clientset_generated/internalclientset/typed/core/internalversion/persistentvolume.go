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

type PersistentVolumesGetter interface {
 PersistentVolumes() PersistentVolumeInterface
}
type PersistentVolumeInterface interface {
 Create(*core.PersistentVolume) (*core.PersistentVolume, error)
 Update(*core.PersistentVolume) (*core.PersistentVolume, error)
 UpdateStatus(*core.PersistentVolume) (*core.PersistentVolume, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.PersistentVolume, error)
 List(opts v1.ListOptions) (*core.PersistentVolumeList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.PersistentVolume, err error)
 PersistentVolumeExpansion
}
type persistentVolumes struct{ client rest.Interface }

func newPersistentVolumes(c *CoreClient) *persistentVolumes {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &persistentVolumes{client: c.RESTClient()}
}
func (c *persistentVolumes) Get(name string, options v1.GetOptions) (result *core.PersistentVolume, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PersistentVolume{}
 err = c.client.Get().Resource("persistentvolumes").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *persistentVolumes) List(opts v1.ListOptions) (result *core.PersistentVolumeList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.PersistentVolumeList{}
 err = c.client.Get().Resource("persistentvolumes").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *persistentVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("persistentvolumes").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *persistentVolumes) Create(persistentVolume *core.PersistentVolume) (result *core.PersistentVolume, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PersistentVolume{}
 err = c.client.Post().Resource("persistentvolumes").Body(persistentVolume).Do().Into(result)
 return
}
func (c *persistentVolumes) Update(persistentVolume *core.PersistentVolume) (result *core.PersistentVolume, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PersistentVolume{}
 err = c.client.Put().Resource("persistentvolumes").Name(persistentVolume.Name).Body(persistentVolume).Do().Into(result)
 return
}
func (c *persistentVolumes) UpdateStatus(persistentVolume *core.PersistentVolume) (result *core.PersistentVolume, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PersistentVolume{}
 err = c.client.Put().Resource("persistentvolumes").Name(persistentVolume.Name).SubResource("status").Body(persistentVolume).Do().Into(result)
 return
}
func (c *persistentVolumes) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("persistentvolumes").Name(name).Body(options).Do().Error()
}
func (c *persistentVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("persistentvolumes").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *persistentVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.PersistentVolume, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.PersistentVolume{}
 err = c.client.Patch(pt).Resource("persistentvolumes").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

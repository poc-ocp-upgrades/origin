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

type ConfigMapsGetter interface {
 ConfigMaps(namespace string) ConfigMapInterface
}
type ConfigMapInterface interface {
 Create(*core.ConfigMap) (*core.ConfigMap, error)
 Update(*core.ConfigMap) (*core.ConfigMap, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.ConfigMap, error)
 List(opts v1.ListOptions) (*core.ConfigMapList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ConfigMap, err error)
 ConfigMapExpansion
}
type configMaps struct {
 client rest.Interface
 ns     string
}

func newConfigMaps(c *CoreClient, namespace string) *configMaps {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &configMaps{client: c.RESTClient(), ns: namespace}
}
func (c *configMaps) Get(name string, options v1.GetOptions) (result *core.ConfigMap, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ConfigMap{}
 err = c.client.Get().Namespace(c.ns).Resource("configmaps").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *configMaps) List(opts v1.ListOptions) (result *core.ConfigMapList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.ConfigMapList{}
 err = c.client.Get().Namespace(c.ns).Resource("configmaps").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *configMaps) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("configmaps").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *configMaps) Create(configMap *core.ConfigMap) (result *core.ConfigMap, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ConfigMap{}
 err = c.client.Post().Namespace(c.ns).Resource("configmaps").Body(configMap).Do().Into(result)
 return
}
func (c *configMaps) Update(configMap *core.ConfigMap) (result *core.ConfigMap, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ConfigMap{}
 err = c.client.Put().Namespace(c.ns).Resource("configmaps").Name(configMap.Name).Body(configMap).Do().Into(result)
 return
}
func (c *configMaps) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("configmaps").Name(name).Body(options).Do().Error()
}
func (c *configMaps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("configmaps").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *configMaps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ConfigMap, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ConfigMap{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("configmaps").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

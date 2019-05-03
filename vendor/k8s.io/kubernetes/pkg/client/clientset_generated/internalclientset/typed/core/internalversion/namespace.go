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

type NamespacesGetter interface{ Namespaces() NamespaceInterface }
type NamespaceInterface interface {
 Create(*core.Namespace) (*core.Namespace, error)
 Update(*core.Namespace) (*core.Namespace, error)
 UpdateStatus(*core.Namespace) (*core.Namespace, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.Namespace, error)
 List(opts v1.ListOptions) (*core.NamespaceList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Namespace, err error)
 NamespaceExpansion
}
type namespaces struct{ client rest.Interface }

func newNamespaces(c *CoreClient) *namespaces {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &namespaces{client: c.RESTClient()}
}
func (c *namespaces) Get(name string, options v1.GetOptions) (result *core.Namespace, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Namespace{}
 err = c.client.Get().Resource("namespaces").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *namespaces) List(opts v1.ListOptions) (result *core.NamespaceList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.NamespaceList{}
 err = c.client.Get().Resource("namespaces").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *namespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("namespaces").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *namespaces) Create(namespace *core.Namespace) (result *core.Namespace, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Namespace{}
 err = c.client.Post().Resource("namespaces").Body(namespace).Do().Into(result)
 return
}
func (c *namespaces) Update(namespace *core.Namespace) (result *core.Namespace, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Namespace{}
 err = c.client.Put().Resource("namespaces").Name(namespace.Name).Body(namespace).Do().Into(result)
 return
}
func (c *namespaces) UpdateStatus(namespace *core.Namespace) (result *core.Namespace, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Namespace{}
 err = c.client.Put().Resource("namespaces").Name(namespace.Name).SubResource("status").Body(namespace).Do().Into(result)
 return
}
func (c *namespaces) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("namespaces").Name(name).Body(options).Do().Error()
}
func (c *namespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("namespaces").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *namespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Namespace, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Namespace{}
 err = c.client.Patch(pt).Resource("namespaces").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

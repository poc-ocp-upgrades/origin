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

type SecretsGetter interface {
 Secrets(namespace string) SecretInterface
}
type SecretInterface interface {
 Create(*core.Secret) (*core.Secret, error)
 Update(*core.Secret) (*core.Secret, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.Secret, error)
 List(opts v1.ListOptions) (*core.SecretList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Secret, err error)
 SecretExpansion
}
type secrets struct {
 client rest.Interface
 ns     string
}

func newSecrets(c *CoreClient, namespace string) *secrets {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &secrets{client: c.RESTClient(), ns: namespace}
}
func (c *secrets) Get(name string, options v1.GetOptions) (result *core.Secret, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Secret{}
 err = c.client.Get().Namespace(c.ns).Resource("secrets").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *secrets) List(opts v1.ListOptions) (result *core.SecretList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.SecretList{}
 err = c.client.Get().Namespace(c.ns).Resource("secrets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *secrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("secrets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *secrets) Create(secret *core.Secret) (result *core.Secret, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Secret{}
 err = c.client.Post().Namespace(c.ns).Resource("secrets").Body(secret).Do().Into(result)
 return
}
func (c *secrets) Update(secret *core.Secret) (result *core.Secret, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Secret{}
 err = c.client.Put().Namespace(c.ns).Resource("secrets").Name(secret.Name).Body(secret).Do().Into(result)
 return
}
func (c *secrets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("secrets").Name(name).Body(options).Do().Error()
}
func (c *secrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("secrets").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *secrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Secret, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Secret{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("secrets").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

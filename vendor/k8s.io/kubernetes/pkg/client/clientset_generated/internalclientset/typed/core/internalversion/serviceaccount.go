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

type ServiceAccountsGetter interface {
 ServiceAccounts(namespace string) ServiceAccountInterface
}
type ServiceAccountInterface interface {
 Create(*core.ServiceAccount) (*core.ServiceAccount, error)
 Update(*core.ServiceAccount) (*core.ServiceAccount, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.ServiceAccount, error)
 List(opts v1.ListOptions) (*core.ServiceAccountList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ServiceAccount, err error)
 ServiceAccountExpansion
}
type serviceAccounts struct {
 client rest.Interface
 ns     string
}

func newServiceAccounts(c *CoreClient, namespace string) *serviceAccounts {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &serviceAccounts{client: c.RESTClient(), ns: namespace}
}
func (c *serviceAccounts) Get(name string, options v1.GetOptions) (result *core.ServiceAccount, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ServiceAccount{}
 err = c.client.Get().Namespace(c.ns).Resource("serviceaccounts").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *serviceAccounts) List(opts v1.ListOptions) (result *core.ServiceAccountList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.ServiceAccountList{}
 err = c.client.Get().Namespace(c.ns).Resource("serviceaccounts").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *serviceAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("serviceaccounts").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *serviceAccounts) Create(serviceAccount *core.ServiceAccount) (result *core.ServiceAccount, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ServiceAccount{}
 err = c.client.Post().Namespace(c.ns).Resource("serviceaccounts").Body(serviceAccount).Do().Into(result)
 return
}
func (c *serviceAccounts) Update(serviceAccount *core.ServiceAccount) (result *core.ServiceAccount, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ServiceAccount{}
 err = c.client.Put().Namespace(c.ns).Resource("serviceaccounts").Name(serviceAccount.Name).Body(serviceAccount).Do().Into(result)
 return
}
func (c *serviceAccounts) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("serviceaccounts").Name(name).Body(options).Do().Error()
}
func (c *serviceAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("serviceaccounts").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *serviceAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ServiceAccount, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ServiceAccount{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("serviceaccounts").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

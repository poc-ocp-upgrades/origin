package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 rbac "k8s.io/kubernetes/pkg/apis/rbac"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type RolesGetter interface {
 Roles(namespace string) RoleInterface
}
type RoleInterface interface {
 Create(*rbac.Role) (*rbac.Role, error)
 Update(*rbac.Role) (*rbac.Role, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*rbac.Role, error)
 List(opts v1.ListOptions) (*rbac.RoleList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac.Role, err error)
 RoleExpansion
}
type roles struct {
 client rest.Interface
 ns     string
}

func newRoles(c *RbacClient, namespace string) *roles {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &roles{client: c.RESTClient(), ns: namespace}
}
func (c *roles) Get(name string, options v1.GetOptions) (result *rbac.Role, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.Role{}
 err = c.client.Get().Namespace(c.ns).Resource("roles").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *roles) List(opts v1.ListOptions) (result *rbac.RoleList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &rbac.RoleList{}
 err = c.client.Get().Namespace(c.ns).Resource("roles").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *roles) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("roles").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *roles) Create(role *rbac.Role) (result *rbac.Role, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.Role{}
 err = c.client.Post().Namespace(c.ns).Resource("roles").Body(role).Do().Into(result)
 return
}
func (c *roles) Update(role *rbac.Role) (result *rbac.Role, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.Role{}
 err = c.client.Put().Namespace(c.ns).Resource("roles").Name(role.Name).Body(role).Do().Into(result)
 return
}
func (c *roles) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("roles").Name(name).Body(options).Do().Error()
}
func (c *roles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("roles").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *roles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac.Role, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.Role{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("roles").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

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

type RoleBindingsGetter interface {
 RoleBindings(namespace string) RoleBindingInterface
}
type RoleBindingInterface interface {
 Create(*rbac.RoleBinding) (*rbac.RoleBinding, error)
 Update(*rbac.RoleBinding) (*rbac.RoleBinding, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*rbac.RoleBinding, error)
 List(opts v1.ListOptions) (*rbac.RoleBindingList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac.RoleBinding, err error)
 RoleBindingExpansion
}
type roleBindings struct {
 client rest.Interface
 ns     string
}

func newRoleBindings(c *RbacClient, namespace string) *roleBindings {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &roleBindings{client: c.RESTClient(), ns: namespace}
}
func (c *roleBindings) Get(name string, options v1.GetOptions) (result *rbac.RoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.RoleBinding{}
 err = c.client.Get().Namespace(c.ns).Resource("rolebindings").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *roleBindings) List(opts v1.ListOptions) (result *rbac.RoleBindingList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &rbac.RoleBindingList{}
 err = c.client.Get().Namespace(c.ns).Resource("rolebindings").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *roleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("rolebindings").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *roleBindings) Create(roleBinding *rbac.RoleBinding) (result *rbac.RoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.RoleBinding{}
 err = c.client.Post().Namespace(c.ns).Resource("rolebindings").Body(roleBinding).Do().Into(result)
 return
}
func (c *roleBindings) Update(roleBinding *rbac.RoleBinding) (result *rbac.RoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.RoleBinding{}
 err = c.client.Put().Namespace(c.ns).Resource("rolebindings").Name(roleBinding.Name).Body(roleBinding).Do().Into(result)
 return
}
func (c *roleBindings) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("rolebindings").Name(name).Body(options).Do().Error()
}
func (c *roleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("rolebindings").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *roleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac.RoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.RoleBinding{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("rolebindings").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

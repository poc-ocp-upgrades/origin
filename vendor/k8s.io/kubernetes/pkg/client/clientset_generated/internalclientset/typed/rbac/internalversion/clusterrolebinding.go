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

type ClusterRoleBindingsGetter interface {
 ClusterRoleBindings() ClusterRoleBindingInterface
}
type ClusterRoleBindingInterface interface {
 Create(*rbac.ClusterRoleBinding) (*rbac.ClusterRoleBinding, error)
 Update(*rbac.ClusterRoleBinding) (*rbac.ClusterRoleBinding, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*rbac.ClusterRoleBinding, error)
 List(opts v1.ListOptions) (*rbac.ClusterRoleBindingList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac.ClusterRoleBinding, err error)
 ClusterRoleBindingExpansion
}
type clusterRoleBindings struct{ client rest.Interface }

func newClusterRoleBindings(c *RbacClient) *clusterRoleBindings {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &clusterRoleBindings{client: c.RESTClient()}
}
func (c *clusterRoleBindings) Get(name string, options v1.GetOptions) (result *rbac.ClusterRoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.ClusterRoleBinding{}
 err = c.client.Get().Resource("clusterrolebindings").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *clusterRoleBindings) List(opts v1.ListOptions) (result *rbac.ClusterRoleBindingList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &rbac.ClusterRoleBindingList{}
 err = c.client.Get().Resource("clusterrolebindings").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *clusterRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("clusterrolebindings").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *clusterRoleBindings) Create(clusterRoleBinding *rbac.ClusterRoleBinding) (result *rbac.ClusterRoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.ClusterRoleBinding{}
 err = c.client.Post().Resource("clusterrolebindings").Body(clusterRoleBinding).Do().Into(result)
 return
}
func (c *clusterRoleBindings) Update(clusterRoleBinding *rbac.ClusterRoleBinding) (result *rbac.ClusterRoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.ClusterRoleBinding{}
 err = c.client.Put().Resource("clusterrolebindings").Name(clusterRoleBinding.Name).Body(clusterRoleBinding).Do().Into(result)
 return
}
func (c *clusterRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("clusterrolebindings").Name(name).Body(options).Do().Error()
}
func (c *clusterRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("clusterrolebindings").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *clusterRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac.ClusterRoleBinding, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &rbac.ClusterRoleBinding{}
 err = c.client.Patch(pt).Resource("clusterrolebindings").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

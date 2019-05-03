package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 apps "k8s.io/kubernetes/pkg/apis/apps"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type DeploymentsGetter interface {
 Deployments(namespace string) DeploymentInterface
}
type DeploymentInterface interface {
 Create(*apps.Deployment) (*apps.Deployment, error)
 Update(*apps.Deployment) (*apps.Deployment, error)
 UpdateStatus(*apps.Deployment) (*apps.Deployment, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*apps.Deployment, error)
 List(opts v1.ListOptions) (*apps.DeploymentList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.Deployment, err error)
 DeploymentExpansion
}
type deployments struct {
 client rest.Interface
 ns     string
}

func newDeployments(c *AppsClient, namespace string) *deployments {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &deployments{client: c.RESTClient(), ns: namespace}
}
func (c *deployments) Get(name string, options v1.GetOptions) (result *apps.Deployment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.Deployment{}
 err = c.client.Get().Namespace(c.ns).Resource("deployments").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *deployments) List(opts v1.ListOptions) (result *apps.DeploymentList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &apps.DeploymentList{}
 err = c.client.Get().Namespace(c.ns).Resource("deployments").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *deployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("deployments").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *deployments) Create(deployment *apps.Deployment) (result *apps.Deployment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.Deployment{}
 err = c.client.Post().Namespace(c.ns).Resource("deployments").Body(deployment).Do().Into(result)
 return
}
func (c *deployments) Update(deployment *apps.Deployment) (result *apps.Deployment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.Deployment{}
 err = c.client.Put().Namespace(c.ns).Resource("deployments").Name(deployment.Name).Body(deployment).Do().Into(result)
 return
}
func (c *deployments) UpdateStatus(deployment *apps.Deployment) (result *apps.Deployment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.Deployment{}
 err = c.client.Put().Namespace(c.ns).Resource("deployments").Name(deployment.Name).SubResource("status").Body(deployment).Do().Into(result)
 return
}
func (c *deployments) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("deployments").Name(name).Body(options).Do().Error()
}
func (c *deployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("deployments").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *deployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.Deployment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.Deployment{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("deployments").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
 core "k8s.io/kubernetes/pkg/apis/core"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type ReplicationControllersGetter interface {
 ReplicationControllers(namespace string) ReplicationControllerInterface
}
type ReplicationControllerInterface interface {
 Create(*core.ReplicationController) (*core.ReplicationController, error)
 Update(*core.ReplicationController) (*core.ReplicationController, error)
 UpdateStatus(*core.ReplicationController) (*core.ReplicationController, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.ReplicationController, error)
 List(opts v1.ListOptions) (*core.ReplicationControllerList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ReplicationController, err error)
 GetScale(replicationControllerName string, options v1.GetOptions) (*autoscaling.Scale, error)
 UpdateScale(replicationControllerName string, scale *autoscaling.Scale) (*autoscaling.Scale, error)
 ReplicationControllerExpansion
}
type replicationControllers struct {
 client rest.Interface
 ns     string
}

func newReplicationControllers(c *CoreClient, namespace string) *replicationControllers {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &replicationControllers{client: c.RESTClient(), ns: namespace}
}
func (c *replicationControllers) Get(name string, options v1.GetOptions) (result *core.ReplicationController, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ReplicationController{}
 err = c.client.Get().Namespace(c.ns).Resource("replicationcontrollers").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *replicationControllers) List(opts v1.ListOptions) (result *core.ReplicationControllerList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.ReplicationControllerList{}
 err = c.client.Get().Namespace(c.ns).Resource("replicationcontrollers").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *replicationControllers) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("replicationcontrollers").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *replicationControllers) Create(replicationController *core.ReplicationController) (result *core.ReplicationController, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ReplicationController{}
 err = c.client.Post().Namespace(c.ns).Resource("replicationcontrollers").Body(replicationController).Do().Into(result)
 return
}
func (c *replicationControllers) Update(replicationController *core.ReplicationController) (result *core.ReplicationController, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ReplicationController{}
 err = c.client.Put().Namespace(c.ns).Resource("replicationcontrollers").Name(replicationController.Name).Body(replicationController).Do().Into(result)
 return
}
func (c *replicationControllers) UpdateStatus(replicationController *core.ReplicationController) (result *core.ReplicationController, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ReplicationController{}
 err = c.client.Put().Namespace(c.ns).Resource("replicationcontrollers").Name(replicationController.Name).SubResource("status").Body(replicationController).Do().Into(result)
 return
}
func (c *replicationControllers) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("replicationcontrollers").Name(name).Body(options).Do().Error()
}
func (c *replicationControllers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("replicationcontrollers").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *replicationControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ReplicationController, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.ReplicationController{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("replicationcontrollers").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}
func (c *replicationControllers) GetScale(replicationControllerName string, options v1.GetOptions) (result *autoscaling.Scale, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &autoscaling.Scale{}
 err = c.client.Get().Namespace(c.ns).Resource("replicationcontrollers").Name(replicationControllerName).SubResource("scale").VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *replicationControllers) UpdateScale(replicationControllerName string, scale *autoscaling.Scale) (result *autoscaling.Scale, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &autoscaling.Scale{}
 err = c.client.Put().Namespace(c.ns).Resource("replicationcontrollers").Name(replicationControllerName).SubResource("scale").Body(scale).Do().Into(result)
 return
}

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

type NodesGetter interface{ Nodes() NodeInterface }
type NodeInterface interface {
 Create(*core.Node) (*core.Node, error)
 Update(*core.Node) (*core.Node, error)
 UpdateStatus(*core.Node) (*core.Node, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*core.Node, error)
 List(opts v1.ListOptions) (*core.NodeList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Node, err error)
 NodeExpansion
}
type nodes struct{ client rest.Interface }

func newNodes(c *CoreClient) *nodes {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &nodes{client: c.RESTClient()}
}
func (c *nodes) Get(name string, options v1.GetOptions) (result *core.Node, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Node{}
 err = c.client.Get().Resource("nodes").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *nodes) List(opts v1.ListOptions) (result *core.NodeList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &core.NodeList{}
 err = c.client.Get().Resource("nodes").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *nodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("nodes").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *nodes) Create(node *core.Node) (result *core.Node, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Node{}
 err = c.client.Post().Resource("nodes").Body(node).Do().Into(result)
 return
}
func (c *nodes) Update(node *core.Node) (result *core.Node, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Node{}
 err = c.client.Put().Resource("nodes").Name(node.Name).Body(node).Do().Into(result)
 return
}
func (c *nodes) UpdateStatus(node *core.Node) (result *core.Node, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Node{}
 err = c.client.Put().Resource("nodes").Name(node.Name).SubResource("status").Body(node).Do().Into(result)
 return
}
func (c *nodes) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("nodes").Name(name).Body(options).Do().Error()
}
func (c *nodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("nodes").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *nodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Node, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &core.Node{}
 err = c.client.Patch(pt).Resource("nodes").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

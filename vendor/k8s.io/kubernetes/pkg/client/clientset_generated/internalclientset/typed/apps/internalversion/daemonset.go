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

type DaemonSetsGetter interface {
 DaemonSets(namespace string) DaemonSetInterface
}
type DaemonSetInterface interface {
 Create(*apps.DaemonSet) (*apps.DaemonSet, error)
 Update(*apps.DaemonSet) (*apps.DaemonSet, error)
 UpdateStatus(*apps.DaemonSet) (*apps.DaemonSet, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*apps.DaemonSet, error)
 List(opts v1.ListOptions) (*apps.DaemonSetList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.DaemonSet, err error)
 DaemonSetExpansion
}
type daemonSets struct {
 client rest.Interface
 ns     string
}

func newDaemonSets(c *AppsClient, namespace string) *daemonSets {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &daemonSets{client: c.RESTClient(), ns: namespace}
}
func (c *daemonSets) Get(name string, options v1.GetOptions) (result *apps.DaemonSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.DaemonSet{}
 err = c.client.Get().Namespace(c.ns).Resource("daemonsets").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *daemonSets) List(opts v1.ListOptions) (result *apps.DaemonSetList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &apps.DaemonSetList{}
 err = c.client.Get().Namespace(c.ns).Resource("daemonsets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *daemonSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("daemonsets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *daemonSets) Create(daemonSet *apps.DaemonSet) (result *apps.DaemonSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.DaemonSet{}
 err = c.client.Post().Namespace(c.ns).Resource("daemonsets").Body(daemonSet).Do().Into(result)
 return
}
func (c *daemonSets) Update(daemonSet *apps.DaemonSet) (result *apps.DaemonSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.DaemonSet{}
 err = c.client.Put().Namespace(c.ns).Resource("daemonsets").Name(daemonSet.Name).Body(daemonSet).Do().Into(result)
 return
}
func (c *daemonSets) UpdateStatus(daemonSet *apps.DaemonSet) (result *apps.DaemonSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.DaemonSet{}
 err = c.client.Put().Namespace(c.ns).Resource("daemonsets").Name(daemonSet.Name).SubResource("status").Body(daemonSet).Do().Into(result)
 return
}
func (c *daemonSets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("daemonsets").Name(name).Body(options).Do().Error()
}
func (c *daemonSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("daemonsets").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *daemonSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.DaemonSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.DaemonSet{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("daemonsets").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

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

type StatefulSetsGetter interface {
 StatefulSets(namespace string) StatefulSetInterface
}
type StatefulSetInterface interface {
 Create(*apps.StatefulSet) (*apps.StatefulSet, error)
 Update(*apps.StatefulSet) (*apps.StatefulSet, error)
 UpdateStatus(*apps.StatefulSet) (*apps.StatefulSet, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*apps.StatefulSet, error)
 List(opts v1.ListOptions) (*apps.StatefulSetList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.StatefulSet, err error)
 StatefulSetExpansion
}
type statefulSets struct {
 client rest.Interface
 ns     string
}

func newStatefulSets(c *AppsClient, namespace string) *statefulSets {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &statefulSets{client: c.RESTClient(), ns: namespace}
}
func (c *statefulSets) Get(name string, options v1.GetOptions) (result *apps.StatefulSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.StatefulSet{}
 err = c.client.Get().Namespace(c.ns).Resource("statefulsets").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *statefulSets) List(opts v1.ListOptions) (result *apps.StatefulSetList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &apps.StatefulSetList{}
 err = c.client.Get().Namespace(c.ns).Resource("statefulsets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *statefulSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("statefulsets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *statefulSets) Create(statefulSet *apps.StatefulSet) (result *apps.StatefulSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.StatefulSet{}
 err = c.client.Post().Namespace(c.ns).Resource("statefulsets").Body(statefulSet).Do().Into(result)
 return
}
func (c *statefulSets) Update(statefulSet *apps.StatefulSet) (result *apps.StatefulSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.StatefulSet{}
 err = c.client.Put().Namespace(c.ns).Resource("statefulsets").Name(statefulSet.Name).Body(statefulSet).Do().Into(result)
 return
}
func (c *statefulSets) UpdateStatus(statefulSet *apps.StatefulSet) (result *apps.StatefulSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.StatefulSet{}
 err = c.client.Put().Namespace(c.ns).Resource("statefulsets").Name(statefulSet.Name).SubResource("status").Body(statefulSet).Do().Into(result)
 return
}
func (c *statefulSets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("statefulsets").Name(name).Body(options).Do().Error()
}
func (c *statefulSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("statefulsets").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *statefulSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.StatefulSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.StatefulSet{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("statefulsets").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

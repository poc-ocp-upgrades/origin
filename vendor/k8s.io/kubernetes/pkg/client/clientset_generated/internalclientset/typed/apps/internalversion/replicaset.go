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

type ReplicaSetsGetter interface {
 ReplicaSets(namespace string) ReplicaSetInterface
}
type ReplicaSetInterface interface {
 Create(*apps.ReplicaSet) (*apps.ReplicaSet, error)
 Update(*apps.ReplicaSet) (*apps.ReplicaSet, error)
 UpdateStatus(*apps.ReplicaSet) (*apps.ReplicaSet, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*apps.ReplicaSet, error)
 List(opts v1.ListOptions) (*apps.ReplicaSetList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.ReplicaSet, err error)
 ReplicaSetExpansion
}
type replicaSets struct {
 client rest.Interface
 ns     string
}

func newReplicaSets(c *AppsClient, namespace string) *replicaSets {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &replicaSets{client: c.RESTClient(), ns: namespace}
}
func (c *replicaSets) Get(name string, options v1.GetOptions) (result *apps.ReplicaSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.ReplicaSet{}
 err = c.client.Get().Namespace(c.ns).Resource("replicasets").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *replicaSets) List(opts v1.ListOptions) (result *apps.ReplicaSetList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &apps.ReplicaSetList{}
 err = c.client.Get().Namespace(c.ns).Resource("replicasets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *replicaSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("replicasets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *replicaSets) Create(replicaSet *apps.ReplicaSet) (result *apps.ReplicaSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.ReplicaSet{}
 err = c.client.Post().Namespace(c.ns).Resource("replicasets").Body(replicaSet).Do().Into(result)
 return
}
func (c *replicaSets) Update(replicaSet *apps.ReplicaSet) (result *apps.ReplicaSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.ReplicaSet{}
 err = c.client.Put().Namespace(c.ns).Resource("replicasets").Name(replicaSet.Name).Body(replicaSet).Do().Into(result)
 return
}
func (c *replicaSets) UpdateStatus(replicaSet *apps.ReplicaSet) (result *apps.ReplicaSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.ReplicaSet{}
 err = c.client.Put().Namespace(c.ns).Resource("replicasets").Name(replicaSet.Name).SubResource("status").Body(replicaSet).Do().Into(result)
 return
}
func (c *replicaSets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("replicasets").Name(name).Body(options).Do().Error()
}
func (c *replicaSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("replicasets").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *replicaSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.ReplicaSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.ReplicaSet{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("replicasets").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

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

type ControllerRevisionsGetter interface {
 ControllerRevisions(namespace string) ControllerRevisionInterface
}
type ControllerRevisionInterface interface {
 Create(*apps.ControllerRevision) (*apps.ControllerRevision, error)
 Update(*apps.ControllerRevision) (*apps.ControllerRevision, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*apps.ControllerRevision, error)
 List(opts v1.ListOptions) (*apps.ControllerRevisionList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.ControllerRevision, err error)
 ControllerRevisionExpansion
}
type controllerRevisions struct {
 client rest.Interface
 ns     string
}

func newControllerRevisions(c *AppsClient, namespace string) *controllerRevisions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &controllerRevisions{client: c.RESTClient(), ns: namespace}
}
func (c *controllerRevisions) Get(name string, options v1.GetOptions) (result *apps.ControllerRevision, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.ControllerRevision{}
 err = c.client.Get().Namespace(c.ns).Resource("controllerrevisions").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *controllerRevisions) List(opts v1.ListOptions) (result *apps.ControllerRevisionList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &apps.ControllerRevisionList{}
 err = c.client.Get().Namespace(c.ns).Resource("controllerrevisions").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *controllerRevisions) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("controllerrevisions").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *controllerRevisions) Create(controllerRevision *apps.ControllerRevision) (result *apps.ControllerRevision, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.ControllerRevision{}
 err = c.client.Post().Namespace(c.ns).Resource("controllerrevisions").Body(controllerRevision).Do().Into(result)
 return
}
func (c *controllerRevisions) Update(controllerRevision *apps.ControllerRevision) (result *apps.ControllerRevision, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.ControllerRevision{}
 err = c.client.Put().Namespace(c.ns).Resource("controllerrevisions").Name(controllerRevision.Name).Body(controllerRevision).Do().Into(result)
 return
}
func (c *controllerRevisions) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("controllerrevisions").Name(name).Body(options).Do().Error()
}
func (c *controllerRevisions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("controllerrevisions").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *controllerRevisions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.ControllerRevision, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &apps.ControllerRevision{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("controllerrevisions").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 batch "k8s.io/kubernetes/pkg/apis/batch"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type JobsGetter interface {
 Jobs(namespace string) JobInterface
}
type JobInterface interface {
 Create(*batch.Job) (*batch.Job, error)
 Update(*batch.Job) (*batch.Job, error)
 UpdateStatus(*batch.Job) (*batch.Job, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*batch.Job, error)
 List(opts v1.ListOptions) (*batch.JobList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.Job, err error)
 JobExpansion
}
type jobs struct {
 client rest.Interface
 ns     string
}

func newJobs(c *BatchClient, namespace string) *jobs {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &jobs{client: c.RESTClient(), ns: namespace}
}
func (c *jobs) Get(name string, options v1.GetOptions) (result *batch.Job, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &batch.Job{}
 err = c.client.Get().Namespace(c.ns).Resource("jobs").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *jobs) List(opts v1.ListOptions) (result *batch.JobList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &batch.JobList{}
 err = c.client.Get().Namespace(c.ns).Resource("jobs").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *jobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("jobs").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *jobs) Create(job *batch.Job) (result *batch.Job, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &batch.Job{}
 err = c.client.Post().Namespace(c.ns).Resource("jobs").Body(job).Do().Into(result)
 return
}
func (c *jobs) Update(job *batch.Job) (result *batch.Job, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &batch.Job{}
 err = c.client.Put().Namespace(c.ns).Resource("jobs").Name(job.Name).Body(job).Do().Into(result)
 return
}
func (c *jobs) UpdateStatus(job *batch.Job) (result *batch.Job, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &batch.Job{}
 err = c.client.Put().Namespace(c.ns).Resource("jobs").Name(job.Name).SubResource("status").Body(job).Do().Into(result)
 return
}
func (c *jobs) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("jobs").Name(name).Body(options).Do().Error()
}
func (c *jobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("jobs").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *jobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.Job, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &batch.Job{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("jobs").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

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

type CronJobsGetter interface {
 CronJobs(namespace string) CronJobInterface
}
type CronJobInterface interface {
 Create(*batch.CronJob) (*batch.CronJob, error)
 Update(*batch.CronJob) (*batch.CronJob, error)
 UpdateStatus(*batch.CronJob) (*batch.CronJob, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*batch.CronJob, error)
 List(opts v1.ListOptions) (*batch.CronJobList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.CronJob, err error)
 CronJobExpansion
}
type cronJobs struct {
 client rest.Interface
 ns     string
}

func newCronJobs(c *BatchClient, namespace string) *cronJobs {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &cronJobs{client: c.RESTClient(), ns: namespace}
}
func (c *cronJobs) Get(name string, options v1.GetOptions) (result *batch.CronJob, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &batch.CronJob{}
 err = c.client.Get().Namespace(c.ns).Resource("cronjobs").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *cronJobs) List(opts v1.ListOptions) (result *batch.CronJobList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &batch.CronJobList{}
 err = c.client.Get().Namespace(c.ns).Resource("cronjobs").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *cronJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Namespace(c.ns).Resource("cronjobs").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *cronJobs) Create(cronJob *batch.CronJob) (result *batch.CronJob, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &batch.CronJob{}
 err = c.client.Post().Namespace(c.ns).Resource("cronjobs").Body(cronJob).Do().Into(result)
 return
}
func (c *cronJobs) Update(cronJob *batch.CronJob) (result *batch.CronJob, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &batch.CronJob{}
 err = c.client.Put().Namespace(c.ns).Resource("cronjobs").Name(cronJob.Name).Body(cronJob).Do().Into(result)
 return
}
func (c *cronJobs) UpdateStatus(cronJob *batch.CronJob) (result *batch.CronJob, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &batch.CronJob{}
 err = c.client.Put().Namespace(c.ns).Resource("cronjobs").Name(cronJob.Name).SubResource("status").Body(cronJob).Do().Into(result)
 return
}
func (c *cronJobs) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Namespace(c.ns).Resource("cronjobs").Name(name).Body(options).Do().Error()
}
func (c *cronJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Namespace(c.ns).Resource("cronjobs").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *cronJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.CronJob, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &batch.CronJob{}
 err = c.client.Patch(pt).Namespace(c.ns).Resource("cronjobs").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

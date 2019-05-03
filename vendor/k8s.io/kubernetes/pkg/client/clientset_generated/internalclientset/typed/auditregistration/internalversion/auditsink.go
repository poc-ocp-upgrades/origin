package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 auditregistration "k8s.io/kubernetes/pkg/apis/auditregistration"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type AuditSinksGetter interface{ AuditSinks() AuditSinkInterface }
type AuditSinkInterface interface {
 Create(*auditregistration.AuditSink) (*auditregistration.AuditSink, error)
 Update(*auditregistration.AuditSink) (*auditregistration.AuditSink, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*auditregistration.AuditSink, error)
 List(opts v1.ListOptions) (*auditregistration.AuditSinkList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auditregistration.AuditSink, err error)
 AuditSinkExpansion
}
type auditSinks struct{ client rest.Interface }

func newAuditSinks(c *AuditregistrationClient) *auditSinks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &auditSinks{client: c.RESTClient()}
}
func (c *auditSinks) Get(name string, options v1.GetOptions) (result *auditregistration.AuditSink, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &auditregistration.AuditSink{}
 err = c.client.Get().Resource("auditsinks").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *auditSinks) List(opts v1.ListOptions) (result *auditregistration.AuditSinkList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &auditregistration.AuditSinkList{}
 err = c.client.Get().Resource("auditsinks").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *auditSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("auditsinks").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *auditSinks) Create(auditSink *auditregistration.AuditSink) (result *auditregistration.AuditSink, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &auditregistration.AuditSink{}
 err = c.client.Post().Resource("auditsinks").Body(auditSink).Do().Into(result)
 return
}
func (c *auditSinks) Update(auditSink *auditregistration.AuditSink) (result *auditregistration.AuditSink, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &auditregistration.AuditSink{}
 err = c.client.Put().Resource("auditsinks").Name(auditSink.Name).Body(auditSink).Do().Into(result)
 return
}
func (c *auditSinks) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("auditsinks").Name(name).Body(options).Do().Error()
}
func (c *auditSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("auditsinks").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *auditSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auditregistration.AuditSink, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &auditregistration.AuditSink{}
 err = c.client.Patch(pt).Resource("auditsinks").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

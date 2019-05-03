package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type MutatingWebhookConfigurationsGetter interface {
 MutatingWebhookConfigurations() MutatingWebhookConfigurationInterface
}
type MutatingWebhookConfigurationInterface interface {
 Create(*admissionregistration.MutatingWebhookConfiguration) (*admissionregistration.MutatingWebhookConfiguration, error)
 Update(*admissionregistration.MutatingWebhookConfiguration) (*admissionregistration.MutatingWebhookConfiguration, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*admissionregistration.MutatingWebhookConfiguration, error)
 List(opts v1.ListOptions) (*admissionregistration.MutatingWebhookConfigurationList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admissionregistration.MutatingWebhookConfiguration, err error)
 MutatingWebhookConfigurationExpansion
}
type mutatingWebhookConfigurations struct{ client rest.Interface }

func newMutatingWebhookConfigurations(c *AdmissionregistrationClient) *mutatingWebhookConfigurations {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &mutatingWebhookConfigurations{client: c.RESTClient()}
}
func (c *mutatingWebhookConfigurations) Get(name string, options v1.GetOptions) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.MutatingWebhookConfiguration{}
 err = c.client.Get().Resource("mutatingwebhookconfigurations").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *mutatingWebhookConfigurations) List(opts v1.ListOptions) (result *admissionregistration.MutatingWebhookConfigurationList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &admissionregistration.MutatingWebhookConfigurationList{}
 err = c.client.Get().Resource("mutatingwebhookconfigurations").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *mutatingWebhookConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("mutatingwebhookconfigurations").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *mutatingWebhookConfigurations) Create(mutatingWebhookConfiguration *admissionregistration.MutatingWebhookConfiguration) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.MutatingWebhookConfiguration{}
 err = c.client.Post().Resource("mutatingwebhookconfigurations").Body(mutatingWebhookConfiguration).Do().Into(result)
 return
}
func (c *mutatingWebhookConfigurations) Update(mutatingWebhookConfiguration *admissionregistration.MutatingWebhookConfiguration) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.MutatingWebhookConfiguration{}
 err = c.client.Put().Resource("mutatingwebhookconfigurations").Name(mutatingWebhookConfiguration.Name).Body(mutatingWebhookConfiguration).Do().Into(result)
 return
}
func (c *mutatingWebhookConfigurations) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("mutatingwebhookconfigurations").Name(name).Body(options).Do().Error()
}
func (c *mutatingWebhookConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("mutatingwebhookconfigurations").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *mutatingWebhookConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.MutatingWebhookConfiguration{}
 err = c.client.Patch(pt).Resource("mutatingwebhookconfigurations").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

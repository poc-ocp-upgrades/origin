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

type ValidatingWebhookConfigurationsGetter interface {
 ValidatingWebhookConfigurations() ValidatingWebhookConfigurationInterface
}
type ValidatingWebhookConfigurationInterface interface {
 Create(*admissionregistration.ValidatingWebhookConfiguration) (*admissionregistration.ValidatingWebhookConfiguration, error)
 Update(*admissionregistration.ValidatingWebhookConfiguration) (*admissionregistration.ValidatingWebhookConfiguration, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*admissionregistration.ValidatingWebhookConfiguration, error)
 List(opts v1.ListOptions) (*admissionregistration.ValidatingWebhookConfigurationList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admissionregistration.ValidatingWebhookConfiguration, err error)
 ValidatingWebhookConfigurationExpansion
}
type validatingWebhookConfigurations struct{ client rest.Interface }

func newValidatingWebhookConfigurations(c *AdmissionregistrationClient) *validatingWebhookConfigurations {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &validatingWebhookConfigurations{client: c.RESTClient()}
}
func (c *validatingWebhookConfigurations) Get(name string, options v1.GetOptions) (result *admissionregistration.ValidatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.ValidatingWebhookConfiguration{}
 err = c.client.Get().Resource("validatingwebhookconfigurations").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *validatingWebhookConfigurations) List(opts v1.ListOptions) (result *admissionregistration.ValidatingWebhookConfigurationList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &admissionregistration.ValidatingWebhookConfigurationList{}
 err = c.client.Get().Resource("validatingwebhookconfigurations").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *validatingWebhookConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("validatingwebhookconfigurations").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *validatingWebhookConfigurations) Create(validatingWebhookConfiguration *admissionregistration.ValidatingWebhookConfiguration) (result *admissionregistration.ValidatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.ValidatingWebhookConfiguration{}
 err = c.client.Post().Resource("validatingwebhookconfigurations").Body(validatingWebhookConfiguration).Do().Into(result)
 return
}
func (c *validatingWebhookConfigurations) Update(validatingWebhookConfiguration *admissionregistration.ValidatingWebhookConfiguration) (result *admissionregistration.ValidatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.ValidatingWebhookConfiguration{}
 err = c.client.Put().Resource("validatingwebhookconfigurations").Name(validatingWebhookConfiguration.Name).Body(validatingWebhookConfiguration).Do().Into(result)
 return
}
func (c *validatingWebhookConfigurations) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("validatingwebhookconfigurations").Name(name).Body(options).Do().Error()
}
func (c *validatingWebhookConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("validatingwebhookconfigurations").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *validatingWebhookConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admissionregistration.ValidatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.ValidatingWebhookConfiguration{}
 err = c.client.Patch(pt).Resource("validatingwebhookconfigurations").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

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

type InitializerConfigurationsGetter interface {
 InitializerConfigurations() InitializerConfigurationInterface
}
type InitializerConfigurationInterface interface {
 Create(*admissionregistration.InitializerConfiguration) (*admissionregistration.InitializerConfiguration, error)
 Update(*admissionregistration.InitializerConfiguration) (*admissionregistration.InitializerConfiguration, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*admissionregistration.InitializerConfiguration, error)
 List(opts v1.ListOptions) (*admissionregistration.InitializerConfigurationList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admissionregistration.InitializerConfiguration, err error)
 InitializerConfigurationExpansion
}
type initializerConfigurations struct{ client rest.Interface }

func newInitializerConfigurations(c *AdmissionregistrationClient) *initializerConfigurations {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &initializerConfigurations{client: c.RESTClient()}
}
func (c *initializerConfigurations) Get(name string, options v1.GetOptions) (result *admissionregistration.InitializerConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.InitializerConfiguration{}
 err = c.client.Get().Resource("initializerconfigurations").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *initializerConfigurations) List(opts v1.ListOptions) (result *admissionregistration.InitializerConfigurationList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &admissionregistration.InitializerConfigurationList{}
 err = c.client.Get().Resource("initializerconfigurations").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *initializerConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("initializerconfigurations").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *initializerConfigurations) Create(initializerConfiguration *admissionregistration.InitializerConfiguration) (result *admissionregistration.InitializerConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.InitializerConfiguration{}
 err = c.client.Post().Resource("initializerconfigurations").Body(initializerConfiguration).Do().Into(result)
 return
}
func (c *initializerConfigurations) Update(initializerConfiguration *admissionregistration.InitializerConfiguration) (result *admissionregistration.InitializerConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.InitializerConfiguration{}
 err = c.client.Put().Resource("initializerconfigurations").Name(initializerConfiguration.Name).Body(initializerConfiguration).Do().Into(result)
 return
}
func (c *initializerConfigurations) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("initializerconfigurations").Name(name).Body(options).Do().Error()
}
func (c *initializerConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("initializerconfigurations").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *initializerConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admissionregistration.InitializerConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &admissionregistration.InitializerConfiguration{}
 err = c.client.Patch(pt).Resource("initializerconfigurations").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

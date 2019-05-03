package internalversion

import (
 "time"
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 rest "k8s.io/client-go/rest"
 certificates "k8s.io/kubernetes/pkg/apis/certificates"
 scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type CertificateSigningRequestsGetter interface {
 CertificateSigningRequests() CertificateSigningRequestInterface
}
type CertificateSigningRequestInterface interface {
 Create(*certificates.CertificateSigningRequest) (*certificates.CertificateSigningRequest, error)
 Update(*certificates.CertificateSigningRequest) (*certificates.CertificateSigningRequest, error)
 UpdateStatus(*certificates.CertificateSigningRequest) (*certificates.CertificateSigningRequest, error)
 Delete(name string, options *v1.DeleteOptions) error
 DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
 Get(name string, options v1.GetOptions) (*certificates.CertificateSigningRequest, error)
 List(opts v1.ListOptions) (*certificates.CertificateSigningRequestList, error)
 Watch(opts v1.ListOptions) (watch.Interface, error)
 Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *certificates.CertificateSigningRequest, err error)
 CertificateSigningRequestExpansion
}
type certificateSigningRequests struct{ client rest.Interface }

func newCertificateSigningRequests(c *CertificatesClient) *certificateSigningRequests {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &certificateSigningRequests{client: c.RESTClient()}
}
func (c *certificateSigningRequests) Get(name string, options v1.GetOptions) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &certificates.CertificateSigningRequest{}
 err = c.client.Get().Resource("certificatesigningrequests").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
 return
}
func (c *certificateSigningRequests) List(opts v1.ListOptions) (result *certificates.CertificateSigningRequestList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 result = &certificates.CertificateSigningRequestList{}
 err = c.client.Get().Resource("certificatesigningrequests").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
 return
}
func (c *certificateSigningRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if opts.TimeoutSeconds != nil {
  timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
 }
 opts.Watch = true
 return c.client.Get().Resource("certificatesigningrequests").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *certificateSigningRequests) Create(certificateSigningRequest *certificates.CertificateSigningRequest) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &certificates.CertificateSigningRequest{}
 err = c.client.Post().Resource("certificatesigningrequests").Body(certificateSigningRequest).Do().Into(result)
 return
}
func (c *certificateSigningRequests) Update(certificateSigningRequest *certificates.CertificateSigningRequest) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &certificates.CertificateSigningRequest{}
 err = c.client.Put().Resource("certificatesigningrequests").Name(certificateSigningRequest.Name).Body(certificateSigningRequest).Do().Into(result)
 return
}
func (c *certificateSigningRequests) UpdateStatus(certificateSigningRequest *certificates.CertificateSigningRequest) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &certificates.CertificateSigningRequest{}
 err = c.client.Put().Resource("certificatesigningrequests").Name(certificateSigningRequest.Name).SubResource("status").Body(certificateSigningRequest).Do().Into(result)
 return
}
func (c *certificateSigningRequests) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Delete().Resource("certificatesigningrequests").Name(name).Body(options).Do().Error()
}
func (c *certificateSigningRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var timeout time.Duration
 if listOptions.TimeoutSeconds != nil {
  timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
 }
 return c.client.Delete().Resource("certificatesigningrequests").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *certificateSigningRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &certificates.CertificateSigningRequest{}
 err = c.client.Patch(pt).Resource("certificatesigningrequests").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
 return
}

package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 certificates "k8s.io/kubernetes/pkg/apis/certificates"
)

type FakeCertificateSigningRequests struct{ Fake *FakeCertificates }

var certificatesigningrequestsResource = schema.GroupVersionResource{Group: "certificates.k8s.io", Version: "", Resource: "certificatesigningrequests"}
var certificatesigningrequestsKind = schema.GroupVersionKind{Group: "certificates.k8s.io", Version: "", Kind: "CertificateSigningRequest"}

func (c *FakeCertificateSigningRequests) Get(name string, options v1.GetOptions) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(certificatesigningrequestsResource, name), &certificates.CertificateSigningRequest{})
 if obj == nil {
  return nil, err
 }
 return obj.(*certificates.CertificateSigningRequest), err
}
func (c *FakeCertificateSigningRequests) List(opts v1.ListOptions) (result *certificates.CertificateSigningRequestList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(certificatesigningrequestsResource, certificatesigningrequestsKind, opts), &certificates.CertificateSigningRequestList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &certificates.CertificateSigningRequestList{ListMeta: obj.(*certificates.CertificateSigningRequestList).ListMeta}
 for _, item := range obj.(*certificates.CertificateSigningRequestList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeCertificateSigningRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(certificatesigningrequestsResource, opts))
}
func (c *FakeCertificateSigningRequests) Create(certificateSigningRequest *certificates.CertificateSigningRequest) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(certificatesigningrequestsResource, certificateSigningRequest), &certificates.CertificateSigningRequest{})
 if obj == nil {
  return nil, err
 }
 return obj.(*certificates.CertificateSigningRequest), err
}
func (c *FakeCertificateSigningRequests) Update(certificateSigningRequest *certificates.CertificateSigningRequest) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(certificatesigningrequestsResource, certificateSigningRequest), &certificates.CertificateSigningRequest{})
 if obj == nil {
  return nil, err
 }
 return obj.(*certificates.CertificateSigningRequest), err
}
func (c *FakeCertificateSigningRequests) UpdateStatus(certificateSigningRequest *certificates.CertificateSigningRequest) (*certificates.CertificateSigningRequest, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(certificatesigningrequestsResource, "status", certificateSigningRequest), &certificates.CertificateSigningRequest{})
 if obj == nil {
  return nil, err
 }
 return obj.(*certificates.CertificateSigningRequest), err
}
func (c *FakeCertificateSigningRequests) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(certificatesigningrequestsResource, name), &certificates.CertificateSigningRequest{})
 return err
}
func (c *FakeCertificateSigningRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(certificatesigningrequestsResource, listOptions)
 _, err := c.Fake.Invokes(action, &certificates.CertificateSigningRequestList{})
 return err
}
func (c *FakeCertificateSigningRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(certificatesigningrequestsResource, name, pt, data, subresources...), &certificates.CertificateSigningRequest{})
 if obj == nil {
  return nil, err
 }
 return obj.(*certificates.CertificateSigningRequest), err
}

package fake

import (
 core "k8s.io/client-go/testing"
 "k8s.io/kubernetes/pkg/apis/certificates"
)

func (c *FakeCertificateSigningRequests) UpdateApproval(certificateSigningRequest *certificates.CertificateSigningRequest) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(core.NewRootUpdateSubresourceAction(certificatesigningrequestsResource, "approval", certificateSigningRequest), &certificates.CertificateSigningRequest{})
 if obj == nil {
  return nil, err
 }
 return obj.(*certificates.CertificateSigningRequest), err
}

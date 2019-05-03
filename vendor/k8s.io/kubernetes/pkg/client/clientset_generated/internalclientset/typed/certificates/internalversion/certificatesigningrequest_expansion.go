package internalversion

import "k8s.io/kubernetes/pkg/apis/certificates"

type CertificateSigningRequestExpansion interface {
 UpdateApproval(certificateSigningRequest *certificates.CertificateSigningRequest) (result *certificates.CertificateSigningRequest, err error)
}

func (c *certificateSigningRequests) UpdateApproval(certificateSigningRequest *certificates.CertificateSigningRequest) (result *certificates.CertificateSigningRequest, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &certificates.CertificateSigningRequest{}
 err = c.client.Put().Resource("certificatesigningrequests").Name(certificateSigningRequest.Name).Body(certificateSigningRequest).SubResource("approval").Do().Into(result)
 return
}

package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/certificates/internalversion"
)

type FakeCertificates struct{ *testing.Fake }

func (c *FakeCertificates) CertificateSigningRequests() internalversion.CertificateSigningRequestInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeCertificateSigningRequests{c}
}
func (c *FakeCertificates) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}

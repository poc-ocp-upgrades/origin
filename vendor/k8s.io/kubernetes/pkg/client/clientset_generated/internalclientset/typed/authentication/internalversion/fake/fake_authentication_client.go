package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authentication/internalversion"
)

type FakeAuthentication struct{ *testing.Fake }

func (c *FakeAuthentication) TokenReviews() internalversion.TokenReviewInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeTokenReviews{c}
}
func (c *FakeAuthentication) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}

package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authorization/internalversion"
)

type FakeAuthorization struct{ *testing.Fake }

func (c *FakeAuthorization) LocalSubjectAccessReviews(namespace string) internalversion.LocalSubjectAccessReviewInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeLocalSubjectAccessReviews{c, namespace}
}
func (c *FakeAuthorization) SelfSubjectAccessReviews() internalversion.SelfSubjectAccessReviewInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeSelfSubjectAccessReviews{c}
}
func (c *FakeAuthorization) SelfSubjectRulesReviews() internalversion.SelfSubjectRulesReviewInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeSelfSubjectRulesReviews{c}
}
func (c *FakeAuthorization) SubjectAccessReviews() internalversion.SubjectAccessReviewInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeSubjectAccessReviews{c}
}
func (c *FakeAuthorization) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}

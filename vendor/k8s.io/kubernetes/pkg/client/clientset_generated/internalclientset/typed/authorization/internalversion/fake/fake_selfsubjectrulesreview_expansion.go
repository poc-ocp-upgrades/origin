package fake

import (
 core "k8s.io/client-go/testing"
 authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
)

func (c *FakeSelfSubjectRulesReviews) Create(srr *authorizationapi.SelfSubjectRulesReview) (result *authorizationapi.SelfSubjectRulesReview, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(core.NewRootCreateAction(authorizationapi.SchemeGroupVersion.WithResource("selfsubjectrulesreviews"), srr), &authorizationapi.SelfSubjectRulesReview{})
 return obj.(*authorizationapi.SelfSubjectRulesReview), err
}

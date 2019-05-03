package fake

import (
 core "k8s.io/client-go/testing"
 authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
)

func (c *FakeSelfSubjectAccessReviews) Create(sar *authorizationapi.SelfSubjectAccessReview) (result *authorizationapi.SelfSubjectAccessReview, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(core.NewRootCreateAction(authorizationapi.SchemeGroupVersion.WithResource("selfsubjectaccessreviews"), sar), &authorizationapi.SelfSubjectAccessReview{})
 return obj.(*authorizationapi.SelfSubjectAccessReview), err
}

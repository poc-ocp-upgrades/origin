package fake

import (
 authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
 core "k8s.io/client-go/testing"
)

func (c *FakeSubjectAccessReviews) Create(sar *authorizationapi.SubjectAccessReview) (result *authorizationapi.SubjectAccessReview, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(core.NewRootCreateAction(authorizationapi.SchemeGroupVersion.WithResource("subjectaccessreviews"), sar), &authorizationapi.SubjectAccessReview{})
 return obj.(*authorizationapi.SubjectAccessReview), err
}

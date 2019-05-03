package fake

import (
 authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
 core "k8s.io/client-go/testing"
)

func (c *FakeLocalSubjectAccessReviews) Create(sar *authorizationapi.LocalSubjectAccessReview) (result *authorizationapi.LocalSubjectAccessReview, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(core.NewCreateAction(authorizationapi.SchemeGroupVersion.WithResource("localsubjectaccessreviews"), c.ns, sar), &authorizationapi.SubjectAccessReview{})
 return obj.(*authorizationapi.LocalSubjectAccessReview), err
}

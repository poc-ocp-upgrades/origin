package fake

import (
 core "k8s.io/client-go/testing"
 authenticationapi "k8s.io/kubernetes/pkg/apis/authentication"
)

func (c *FakeTokenReviews) Create(tokenReview *authenticationapi.TokenReview) (result *authenticationapi.TokenReview, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(core.NewRootCreateAction(authenticationapi.SchemeGroupVersion.WithResource("tokenreviews"), tokenReview), &authenticationapi.TokenReview{})
 return obj.(*authenticationapi.TokenReview), err
}

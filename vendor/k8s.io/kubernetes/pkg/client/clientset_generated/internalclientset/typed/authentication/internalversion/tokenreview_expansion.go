package internalversion

import (
 authenticationapi "k8s.io/kubernetes/pkg/apis/authentication"
)

type TokenReviewExpansion interface {
 Create(tokenReview *authenticationapi.TokenReview) (result *authenticationapi.TokenReview, err error)
}

func (c *tokenReviews) Create(tokenReview *authenticationapi.TokenReview) (result *authenticationapi.TokenReview, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &authenticationapi.TokenReview{}
 err = c.client.Post().Resource("tokenreviews").Body(tokenReview).Do().Into(result)
 return
}

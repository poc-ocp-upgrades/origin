package internalversion

import (
 rest "k8s.io/client-go/rest"
)

type TokenReviewsGetter interface{ TokenReviews() TokenReviewInterface }
type TokenReviewInterface interface{ TokenReviewExpansion }
type tokenReviews struct{ client rest.Interface }

func newTokenReviews(c *AuthenticationClient) *tokenReviews {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &tokenReviews{client: c.RESTClient()}
}

package internalversion

import (
 rest "k8s.io/client-go/rest"
)

type SelfSubjectAccessReviewsGetter interface {
 SelfSubjectAccessReviews() SelfSubjectAccessReviewInterface
}
type SelfSubjectAccessReviewInterface interface {
 SelfSubjectAccessReviewExpansion
}
type selfSubjectAccessReviews struct{ client rest.Interface }

func newSelfSubjectAccessReviews(c *AuthorizationClient) *selfSubjectAccessReviews {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &selfSubjectAccessReviews{client: c.RESTClient()}
}

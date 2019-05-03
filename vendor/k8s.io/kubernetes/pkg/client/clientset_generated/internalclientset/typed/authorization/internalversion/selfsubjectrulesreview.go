package internalversion

import (
 rest "k8s.io/client-go/rest"
)

type SelfSubjectRulesReviewsGetter interface {
 SelfSubjectRulesReviews() SelfSubjectRulesReviewInterface
}
type SelfSubjectRulesReviewInterface interface {
 SelfSubjectRulesReviewExpansion
}
type selfSubjectRulesReviews struct{ client rest.Interface }

func newSelfSubjectRulesReviews(c *AuthorizationClient) *selfSubjectRulesReviews {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &selfSubjectRulesReviews{client: c.RESTClient()}
}

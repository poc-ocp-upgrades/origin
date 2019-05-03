package internalversion

import (
 rest "k8s.io/client-go/rest"
)

type SubjectAccessReviewsGetter interface {
 SubjectAccessReviews() SubjectAccessReviewInterface
}
type SubjectAccessReviewInterface interface{ SubjectAccessReviewExpansion }
type subjectAccessReviews struct{ client rest.Interface }

func newSubjectAccessReviews(c *AuthorizationClient) *subjectAccessReviews {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &subjectAccessReviews{client: c.RESTClient()}
}

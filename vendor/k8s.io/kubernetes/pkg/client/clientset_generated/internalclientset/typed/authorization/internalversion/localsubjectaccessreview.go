package internalversion

import (
 rest "k8s.io/client-go/rest"
)

type LocalSubjectAccessReviewsGetter interface {
 LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewInterface
}
type LocalSubjectAccessReviewInterface interface {
 LocalSubjectAccessReviewExpansion
}
type localSubjectAccessReviews struct {
 client rest.Interface
 ns     string
}

func newLocalSubjectAccessReviews(c *AuthorizationClient, namespace string) *localSubjectAccessReviews {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &localSubjectAccessReviews{client: c.RESTClient(), ns: namespace}
}

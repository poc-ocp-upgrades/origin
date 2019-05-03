package internalversion

import (
 authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
)

type LocalSubjectAccessReviewExpansion interface {
 Create(sar *authorizationapi.LocalSubjectAccessReview) (result *authorizationapi.LocalSubjectAccessReview, err error)
}

func (c *localSubjectAccessReviews) Create(sar *authorizationapi.LocalSubjectAccessReview) (result *authorizationapi.LocalSubjectAccessReview, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &authorizationapi.LocalSubjectAccessReview{}
 err = c.client.Post().Namespace(c.ns).Resource("localsubjectaccessreviews").Body(sar).Do().Into(result)
 return
}

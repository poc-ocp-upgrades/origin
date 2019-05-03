package internalversion

import (
 authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
)

type SelfSubjectAccessReviewExpansion interface {
 Create(sar *authorizationapi.SelfSubjectAccessReview) (result *authorizationapi.SelfSubjectAccessReview, err error)
}

func (c *selfSubjectAccessReviews) Create(sar *authorizationapi.SelfSubjectAccessReview) (result *authorizationapi.SelfSubjectAccessReview, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &authorizationapi.SelfSubjectAccessReview{}
 err = c.client.Post().Resource("selfsubjectaccessreviews").Body(sar).Do().Into(result)
 return
}

package internalversion

import (
 authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
)

type SubjectAccessReviewExpansion interface {
 Create(sar *authorizationapi.SubjectAccessReview) (result *authorizationapi.SubjectAccessReview, err error)
}

func (c *subjectAccessReviews) Create(sar *authorizationapi.SubjectAccessReview) (result *authorizationapi.SubjectAccessReview, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &authorizationapi.SubjectAccessReview{}
 err = c.client.Post().Resource("subjectaccessreviews").Body(sar).Do().Into(result)
 return
}

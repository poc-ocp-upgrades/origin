package internalversion

import (
 authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
)

type SelfSubjectRulesReviewExpansion interface {
 Create(srr *authorizationapi.SelfSubjectRulesReview) (result *authorizationapi.SelfSubjectRulesReview, err error)
}

func (c *selfSubjectRulesReviews) Create(srr *authorizationapi.SelfSubjectRulesReview) (result *authorizationapi.SelfSubjectRulesReview, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &authorizationapi.SelfSubjectRulesReview{}
 err = c.client.Post().Resource("selfsubjectrulesreviews").Body(srr).Do().Into(result)
 return
}

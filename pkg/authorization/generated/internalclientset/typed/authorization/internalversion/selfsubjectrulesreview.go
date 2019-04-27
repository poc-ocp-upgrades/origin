package internalversion

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	rest "k8s.io/client-go/rest"
)

type SelfSubjectRulesReviewsGetter interface {
	SelfSubjectRulesReviews(namespace string) SelfSubjectRulesReviewInterface
}
type SelfSubjectRulesReviewInterface interface {
	Create(*authorization.SelfSubjectRulesReview) (*authorization.SelfSubjectRulesReview, error)
	SelfSubjectRulesReviewExpansion
}
type selfSubjectRulesReviews struct {
	client	rest.Interface
	ns	string
}

func newSelfSubjectRulesReviews(c *AuthorizationClient, namespace string) *selfSubjectRulesReviews {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &selfSubjectRulesReviews{client: c.RESTClient(), ns: namespace}
}
func (c *selfSubjectRulesReviews) Create(selfSubjectRulesReview *authorization.SelfSubjectRulesReview) (result *authorization.SelfSubjectRulesReview, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &authorization.SelfSubjectRulesReview{}
	err = c.client.Post().Namespace(c.ns).Resource("selfsubjectrulesreviews").Body(selfSubjectRulesReview).Do().Into(result)
	return
}

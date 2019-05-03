package internalversion

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	rest "k8s.io/client-go/rest"
)

type SubjectRulesReviewsGetter interface {
	SubjectRulesReviews(namespace string) SubjectRulesReviewInterface
}
type SubjectRulesReviewInterface interface {
	Create(*authorization.SubjectRulesReview) (*authorization.SubjectRulesReview, error)
	SubjectRulesReviewExpansion
}
type subjectRulesReviews struct {
	client rest.Interface
	ns     string
}

func newSubjectRulesReviews(c *AuthorizationClient, namespace string) *subjectRulesReviews {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &subjectRulesReviews{client: c.RESTClient(), ns: namespace}
}
func (c *subjectRulesReviews) Create(subjectRulesReview *authorization.SubjectRulesReview) (result *authorization.SubjectRulesReview, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &authorization.SubjectRulesReview{}
	err = c.client.Post().Namespace(c.ns).Resource("subjectrulesreviews").Body(subjectRulesReview).Do().Into(result)
	return
}

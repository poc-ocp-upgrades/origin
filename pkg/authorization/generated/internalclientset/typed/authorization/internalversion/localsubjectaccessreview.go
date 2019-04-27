package internalversion

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	rest "k8s.io/client-go/rest"
)

type LocalSubjectAccessReviewsGetter interface {
	LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewInterface
}
type LocalSubjectAccessReviewInterface interface {
	Create(*authorization.LocalSubjectAccessReview) (*authorization.SubjectAccessReviewResponse, error)
	LocalSubjectAccessReviewExpansion
}
type localSubjectAccessReviews struct {
	client	rest.Interface
	ns	string
}

func newLocalSubjectAccessReviews(c *AuthorizationClient, namespace string) *localSubjectAccessReviews {
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
	return &localSubjectAccessReviews{client: c.RESTClient(), ns: namespace}
}
func (c *localSubjectAccessReviews) Create(localSubjectAccessReview *authorization.LocalSubjectAccessReview) (result *authorization.SubjectAccessReviewResponse, err error) {
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
	result = &authorization.SubjectAccessReviewResponse{}
	err = c.client.Post().Namespace(c.ns).Resource("localsubjectaccessreviews").Body(localSubjectAccessReview).Do().Into(result)
	return
}

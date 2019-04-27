package internalversion

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	rest "k8s.io/client-go/rest"
)

type ResourceAccessReviewsGetter interface {
	ResourceAccessReviews() ResourceAccessReviewInterface
}
type ResourceAccessReviewInterface interface {
	Create(*authorization.ResourceAccessReview) (*authorization.ResourceAccessReviewResponse, error)
	ResourceAccessReviewExpansion
}
type resourceAccessReviews struct{ client rest.Interface }

func newResourceAccessReviews(c *AuthorizationClient) *resourceAccessReviews {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &resourceAccessReviews{client: c.RESTClient()}
}
func (c *resourceAccessReviews) Create(resourceAccessReview *authorization.ResourceAccessReview) (result *authorization.ResourceAccessReviewResponse, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &authorization.ResourceAccessReviewResponse{}
	err = c.client.Post().Resource("resourceaccessreviews").Body(resourceAccessReview).Do().Into(result)
	return
}

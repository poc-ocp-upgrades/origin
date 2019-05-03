package internalversion

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	rest "k8s.io/client-go/rest"
)

type LocalResourceAccessReviewsGetter interface {
	LocalResourceAccessReviews(namespace string) LocalResourceAccessReviewInterface
}
type LocalResourceAccessReviewInterface interface {
	Create(*authorization.LocalResourceAccessReview) (*authorization.ResourceAccessReviewResponse, error)
	LocalResourceAccessReviewExpansion
}
type localResourceAccessReviews struct {
	client rest.Interface
	ns     string
}

func newLocalResourceAccessReviews(c *AuthorizationClient, namespace string) *localResourceAccessReviews {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &localResourceAccessReviews{client: c.RESTClient(), ns: namespace}
}
func (c *localResourceAccessReviews) Create(localResourceAccessReview *authorization.LocalResourceAccessReview) (result *authorization.ResourceAccessReviewResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &authorization.ResourceAccessReviewResponse{}
	err = c.client.Post().Namespace(c.ns).Resource("localresourceaccessreviews").Body(localResourceAccessReview).Do().Into(result)
	return
}

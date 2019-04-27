package fake

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakeResourceAccessReviews struct{ Fake *FakeAuthorization }

var resourceaccessreviewsResource = schema.GroupVersionResource{Group: "authorization.openshift.io", Version: "", Resource: "resourceaccessreviews"}
var resourceaccessreviewsKind = schema.GroupVersionKind{Group: "authorization.openshift.io", Version: "", Kind: "ResourceAccessReview"}

func (c *FakeResourceAccessReviews) Create(resourceAccessReview *authorization.ResourceAccessReview) (result *authorization.ResourceAccessReviewResponse, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(resourceaccessreviewsResource, resourceAccessReview), &authorization.ResourceAccessReviewResponse{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.ResourceAccessReviewResponse), err
}

package fake

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakeSubjectAccessReviews struct{ Fake *FakeAuthorization }

var subjectaccessreviewsResource = schema.GroupVersionResource{Group: "authorization.openshift.io", Version: "", Resource: "subjectaccessreviews"}
var subjectaccessreviewsKind = schema.GroupVersionKind{Group: "authorization.openshift.io", Version: "", Kind: "SubjectAccessReview"}

func (c *FakeSubjectAccessReviews) Create(subjectAccessReview *authorization.SubjectAccessReview) (result *authorization.SubjectAccessReviewResponse, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(subjectaccessreviewsResource, subjectAccessReview), &authorization.SubjectAccessReviewResponse{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.SubjectAccessReviewResponse), err
}

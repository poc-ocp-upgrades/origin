package fake

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakeSubjectRulesReviews struct {
	Fake	*FakeAuthorization
	ns	string
}

var subjectrulesreviewsResource = schema.GroupVersionResource{Group: "authorization.openshift.io", Version: "", Resource: "subjectrulesreviews"}
var subjectrulesreviewsKind = schema.GroupVersionKind{Group: "authorization.openshift.io", Version: "", Kind: "SubjectRulesReview"}

func (c *FakeSubjectRulesReviews) Create(subjectRulesReview *authorization.SubjectRulesReview) (result *authorization.SubjectRulesReview, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(subjectrulesreviewsResource, c.ns, subjectRulesReview), &authorization.SubjectRulesReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.SubjectRulesReview), err
}

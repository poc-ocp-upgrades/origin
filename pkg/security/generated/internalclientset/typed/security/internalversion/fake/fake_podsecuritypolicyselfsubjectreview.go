package fake

import (
	security "github.com/openshift/origin/pkg/security/apis/security"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakePodSecurityPolicySelfSubjectReviews struct {
	Fake	*FakeSecurity
	ns	string
}

var podsecuritypolicyselfsubjectreviewsResource = schema.GroupVersionResource{Group: "security.openshift.io", Version: "", Resource: "podsecuritypolicyselfsubjectreviews"}
var podsecuritypolicyselfsubjectreviewsKind = schema.GroupVersionKind{Group: "security.openshift.io", Version: "", Kind: "PodSecurityPolicySelfSubjectReview"}

func (c *FakePodSecurityPolicySelfSubjectReviews) Create(podSecurityPolicySelfSubjectReview *security.PodSecurityPolicySelfSubjectReview) (result *security.PodSecurityPolicySelfSubjectReview, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(podsecuritypolicyselfsubjectreviewsResource, c.ns, podSecurityPolicySelfSubjectReview), &security.PodSecurityPolicySelfSubjectReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*security.PodSecurityPolicySelfSubjectReview), err
}

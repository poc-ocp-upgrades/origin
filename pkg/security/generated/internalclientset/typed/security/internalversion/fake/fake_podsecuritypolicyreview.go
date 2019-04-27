package fake

import (
	security "github.com/openshift/origin/pkg/security/apis/security"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakePodSecurityPolicyReviews struct {
	Fake	*FakeSecurity
	ns	string
}

var podsecuritypolicyreviewsResource = schema.GroupVersionResource{Group: "security.openshift.io", Version: "", Resource: "podsecuritypolicyreviews"}
var podsecuritypolicyreviewsKind = schema.GroupVersionKind{Group: "security.openshift.io", Version: "", Kind: "PodSecurityPolicyReview"}

func (c *FakePodSecurityPolicyReviews) Create(podSecurityPolicyReview *security.PodSecurityPolicyReview) (result *security.PodSecurityPolicyReview, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewCreateAction(podsecuritypolicyreviewsResource, c.ns, podSecurityPolicyReview), &security.PodSecurityPolicyReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*security.PodSecurityPolicyReview), err
}

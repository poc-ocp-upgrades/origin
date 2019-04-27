package internalversion

import (
	security "github.com/openshift/origin/pkg/security/apis/security"
	rest "k8s.io/client-go/rest"
)

type PodSecurityPolicySelfSubjectReviewsGetter interface {
	PodSecurityPolicySelfSubjectReviews(namespace string) PodSecurityPolicySelfSubjectReviewInterface
}
type PodSecurityPolicySelfSubjectReviewInterface interface {
	Create(*security.PodSecurityPolicySelfSubjectReview) (*security.PodSecurityPolicySelfSubjectReview, error)
	PodSecurityPolicySelfSubjectReviewExpansion
}
type podSecurityPolicySelfSubjectReviews struct {
	client	rest.Interface
	ns	string
}

func newPodSecurityPolicySelfSubjectReviews(c *SecurityClient, namespace string) *podSecurityPolicySelfSubjectReviews {
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
	return &podSecurityPolicySelfSubjectReviews{client: c.RESTClient(), ns: namespace}
}
func (c *podSecurityPolicySelfSubjectReviews) Create(podSecurityPolicySelfSubjectReview *security.PodSecurityPolicySelfSubjectReview) (result *security.PodSecurityPolicySelfSubjectReview, err error) {
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
	result = &security.PodSecurityPolicySelfSubjectReview{}
	err = c.client.Post().Namespace(c.ns).Resource("podsecuritypolicyselfsubjectreviews").Body(podSecurityPolicySelfSubjectReview).Do().Into(result)
	return
}

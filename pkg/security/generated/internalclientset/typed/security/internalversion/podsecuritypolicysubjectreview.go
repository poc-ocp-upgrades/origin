package internalversion

import (
	security "github.com/openshift/origin/pkg/security/apis/security"
	rest "k8s.io/client-go/rest"
)

type PodSecurityPolicySubjectReviewsGetter interface {
	PodSecurityPolicySubjectReviews(namespace string) PodSecurityPolicySubjectReviewInterface
}
type PodSecurityPolicySubjectReviewInterface interface {
	Create(*security.PodSecurityPolicySubjectReview) (*security.PodSecurityPolicySubjectReview, error)
	PodSecurityPolicySubjectReviewExpansion
}
type podSecurityPolicySubjectReviews struct {
	client	rest.Interface
	ns	string
}

func newPodSecurityPolicySubjectReviews(c *SecurityClient, namespace string) *podSecurityPolicySubjectReviews {
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
	return &podSecurityPolicySubjectReviews{client: c.RESTClient(), ns: namespace}
}
func (c *podSecurityPolicySubjectReviews) Create(podSecurityPolicySubjectReview *security.PodSecurityPolicySubjectReview) (result *security.PodSecurityPolicySubjectReview, err error) {
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
	result = &security.PodSecurityPolicySubjectReview{}
	err = c.client.Post().Namespace(c.ns).Resource("podsecuritypolicysubjectreviews").Body(podSecurityPolicySubjectReview).Do().Into(result)
	return
}

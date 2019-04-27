package internalversion

import (
	security "github.com/openshift/origin/pkg/security/apis/security"
	rest "k8s.io/client-go/rest"
)

type PodSecurityPolicyReviewsGetter interface {
	PodSecurityPolicyReviews(namespace string) PodSecurityPolicyReviewInterface
}
type PodSecurityPolicyReviewInterface interface {
	Create(*security.PodSecurityPolicyReview) (*security.PodSecurityPolicyReview, error)
	PodSecurityPolicyReviewExpansion
}
type podSecurityPolicyReviews struct {
	client	rest.Interface
	ns	string
}

func newPodSecurityPolicyReviews(c *SecurityClient, namespace string) *podSecurityPolicyReviews {
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
	return &podSecurityPolicyReviews{client: c.RESTClient(), ns: namespace}
}
func (c *podSecurityPolicyReviews) Create(podSecurityPolicyReview *security.PodSecurityPolicyReview) (result *security.PodSecurityPolicyReview, err error) {
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
	result = &security.PodSecurityPolicyReview{}
	err = c.client.Post().Namespace(c.ns).Resource("podsecuritypolicyreviews").Body(podSecurityPolicyReview).Do().Into(result)
	return
}

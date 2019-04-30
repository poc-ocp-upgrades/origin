package util

import (
	"time"
	"github.com/openshift/origin/pkg/util/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

func WaitForEndpointsAvailable(oc *CLI, serviceName string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.Poll(200*time.Millisecond, 3*time.Minute, func() (bool, error) {
		ep, err := oc.KubeClient().CoreV1().Endpoints(oc.Namespace()).Get(serviceName, metav1.GetOptions{})
		if errors.TolerateNotFoundError(err) != nil {
			return false, err
		}
		return (len(ep.Subsets) > 0) && (len(ep.Subsets[0].Addresses) > 0), nil
	})
}

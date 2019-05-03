package policy

import (
	buildv1 "github.com/openshift/api/build/v1"
	buildlister "github.com/openshift/client-go/build/listers/build/v1"
	buildclient "github.com/openshift/origin/pkg/build/client"
	buildutil "github.com/openshift/origin/pkg/build/util"
	"k8s.io/apimachinery/pkg/api/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	"time"
)

type SerialLatestOnlyPolicy struct {
	BuildUpdater buildclient.BuildUpdater
	BuildLister  buildlister.BuildLister
}

func (s *SerialLatestOnlyPolicy) IsRunnable(build *buildv1.Build) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bcName := buildutil.ConfigNameForBuild(build)
	if len(bcName) == 0 {
		return true, nil
	}
	if err := kerrors.NewAggregate(s.cancelPreviousBuilds(build)); err != nil {
		return false, err
	}
	nextBuilds, runningBuilds, err := GetNextConfigBuild(s.BuildLister, build.Namespace, bcName)
	if err != nil || runningBuilds {
		return false, err
	}
	return len(nextBuilds) == 1 && nextBuilds[0].Name == build.Name, err
}
func (s *SerialLatestOnlyPolicy) Handles(policy buildv1.BuildRunPolicy) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return policy == buildv1.BuildRunPolicySerialLatestOnly
}
func (s *SerialLatestOnlyPolicy) cancelPreviousBuilds(build *buildv1.Build) []error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bcName := buildutil.ConfigNameForBuild(build)
	if len(bcName) == 0 {
		return []error{}
	}
	currentBuildNumber, err := buildNumber(build)
	if err != nil {
		return []error{NewNoBuildNumberAnnotationError(build)}
	}
	builds, err := buildutil.BuildConfigBuilds(s.BuildLister, build.Namespace, bcName, func(b *buildv1.Build) bool {
		if buildutil.IsBuildComplete(b) || b.Status.Phase == buildv1.BuildPhaseRunning {
			return false
		}
		buildNumber, _ := buildNumber(b)
		return buildNumber < currentBuildNumber
	})
	if err != nil {
		return []error{err}
	}
	var result = []error{}
	for _, b := range builds {
		err := wait.Poll(500*time.Millisecond, 5*time.Second, func() (bool, error) {
			b = b.DeepCopy()
			b.Status.Cancelled = true
			err := s.BuildUpdater.Update(b.Namespace, b)
			if err != nil && errors.IsConflict(err) {
				klog.V(5).Infof("Error cancelling build %s/%s: %v (will retry)", b.Namespace, b.Name, err)
				return false, nil
			}
			return true, err
		})
		if err != nil {
			result = append(result, err)
		}
	}
	return result
}

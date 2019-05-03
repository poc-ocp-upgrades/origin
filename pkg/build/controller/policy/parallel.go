package policy

import (
	buildv1 "github.com/openshift/api/build/v1"
	buildlister "github.com/openshift/client-go/build/listers/build/v1"
	buildclient "github.com/openshift/origin/pkg/build/client"
	buildutil "github.com/openshift/origin/pkg/build/util"
)

type ParallelPolicy struct {
	BuildLister  buildlister.BuildLister
	BuildUpdater buildclient.BuildUpdater
}

func (s *ParallelPolicy) IsRunnable(build *buildv1.Build) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bcName := buildutil.ConfigNameForBuild(build)
	if len(bcName) == 0 {
		return true, nil
	}
	return !hasRunningSerialBuild(s.BuildLister, build.Namespace, bcName), nil
}
func (s *ParallelPolicy) Handles(policy buildv1.BuildRunPolicy) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return policy == buildv1.BuildRunPolicyParallel
}

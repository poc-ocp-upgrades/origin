package policy

import (
	"fmt"
	buildv1 "github.com/openshift/api/build/v1"
	buildlister "github.com/openshift/client-go/build/listers/build/v1"
	buildclient "github.com/openshift/origin/pkg/build/client"
	buildutil "github.com/openshift/origin/pkg/build/util"
	"k8s.io/klog"
	"strconv"
)

type RunPolicy interface {
	IsRunnable(*buildv1.Build) (bool, error)
	Handles(buildv1.BuildRunPolicy) bool
}

func GetAllRunPolicies(lister buildlister.BuildLister, updater buildclient.BuildUpdater) []RunPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []RunPolicy{&ParallelPolicy{BuildLister: lister, BuildUpdater: updater}, &SerialPolicy{BuildLister: lister, BuildUpdater: updater}, &SerialLatestOnlyPolicy{BuildLister: lister, BuildUpdater: updater}}
}
func ForBuild(build *buildv1.Build, policies []RunPolicy) RunPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buildPolicy := buildRunPolicy(build)
	for _, s := range policies {
		if s.Handles(buildPolicy) {
			klog.V(5).Infof("Using %T run policy for build %s/%s", s, build.Namespace, build.Name)
			return s
		}
	}
	return nil
}
func hasRunningSerialBuild(lister buildlister.BuildLister, namespace, buildConfigName string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var hasRunningBuilds bool
	buildutil.BuildConfigBuilds(lister, namespace, buildConfigName, func(b *buildv1.Build) bool {
		switch b.Status.Phase {
		case buildv1.BuildPhasePending, buildv1.BuildPhaseRunning:
			switch buildRunPolicy(b) {
			case buildv1.BuildRunPolicySerial, buildv1.BuildRunPolicySerialLatestOnly:
				hasRunningBuilds = true
			}
		}
		return false
	})
	return hasRunningBuilds
}
func GetNextConfigBuild(lister buildlister.BuildLister, namespace, buildConfigName string) ([]*buildv1.Build, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		nextBuild           *buildv1.Build
		hasRunningBuilds    bool
		previousBuildNumber int64
	)
	builds, err := buildutil.BuildConfigBuilds(lister, namespace, buildConfigName, func(b *buildv1.Build) bool {
		switch b.Status.Phase {
		case buildv1.BuildPhasePending, buildv1.BuildPhaseRunning:
			hasRunningBuilds = true
		case buildv1.BuildPhaseNew:
			return true
		}
		return false
	})
	if err != nil {
		return nil, hasRunningBuilds, err
	}
	nextParallelBuilds := []*buildv1.Build{}
	for i, b := range builds {
		buildNumber, err := buildNumber(b)
		if err != nil {
			return nil, hasRunningBuilds, err
		}
		if buildRunPolicy(b) == buildv1.BuildRunPolicyParallel {
			nextParallelBuilds = append(nextParallelBuilds, b)
		}
		if previousBuildNumber == 0 || buildNumber < previousBuildNumber {
			nextBuild = builds[i]
			previousBuildNumber = buildNumber
		}
	}
	nextBuilds := []*buildv1.Build{}
	if nextBuild != nil && buildRunPolicy(nextBuild) == buildv1.BuildRunPolicyParallel {
		nextBuilds = nextParallelBuilds
	} else if nextBuild != nil {
		nextBuilds = append(nextBuilds, nextBuild)
	}
	return nextBuilds, hasRunningBuilds, nil
}
func buildNumber(build *buildv1.Build) (int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	annotations := build.GetAnnotations()
	if stringNumber, ok := annotations[buildutil.BuildNumberAnnotation]; ok {
		return strconv.ParseInt(stringNumber, 10, 64)
	}
	return 0, fmt.Errorf("build %s/%s does not have %s annotation", build.Namespace, build.Name, buildutil.BuildNumberAnnotation)
}
func buildRunPolicy(build *buildv1.Build) buildv1.BuildRunPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	labels := build.GetLabels()
	if value, found := labels[buildutil.BuildRunPolicyLabel]; found {
		switch value {
		case "Parallel":
			return buildv1.BuildRunPolicyParallel
		case "Serial":
			return buildv1.BuildRunPolicySerial
		case "SerialLatestOnly":
			return buildv1.BuildRunPolicySerialLatestOnly
		}
	}
	klog.V(5).Infof("Build %s/%s does not have start policy label set, using default (Serial)", build.Namespace, build.Name)
	return buildv1.BuildRunPolicySerial
}

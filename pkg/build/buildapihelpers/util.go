package buildapihelpers

import (
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/origin/pkg/api/apihelpers"
	"k8s.io/apimachinery/pkg/util/validation"
)

const (
	buildPodSuffix           = "build"
	caConfigMapSuffix        = "ca"
	sysConfigConfigMapSuffix = "sys-config"
)

func GetBuildPodName(build *buildv1.Build) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return apihelpers.GetPodName(build.Name, buildPodSuffix)
}
func GetBuildCAConfigMapName(build *buildv1.Build) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return apihelpers.GetConfigMapName(build.Name, caConfigMapSuffix)
}
func GetBuildSystemConfigMapName(build *buildv1.Build) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return apihelpers.GetConfigMapName(build.Name, sysConfigConfigMapSuffix)
}
func StrategyType(strategy buildv1.BuildStrategy) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case strategy.DockerStrategy != nil:
		return "Docker"
	case strategy.CustomStrategy != nil:
		return "Custom"
	case strategy.SourceStrategy != nil:
		return "Source"
	case strategy.JenkinsPipelineStrategy != nil:
		return "JenkinsPipeline"
	}
	return ""
}
func LabelValue(name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(name) <= validation.DNS1123LabelMaxLength {
		return name
	}
	return name[:validation.DNS1123LabelMaxLength]
}

package build

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	buildv1 "github.com/openshift/api/build/v1"
	buildutil "github.com/openshift/origin/pkg/build/util"
)

type buildPodCreationStrategy interface {
	CreateBuildPod(build *buildv1.Build, additionalCAs map[string]string, internalRegistryHost string) (*corev1.Pod, error)
}
type typeBasedFactoryStrategy struct {
	dockerBuildStrategy	buildPodCreationStrategy
	sourceBuildStrategy	buildPodCreationStrategy
	customBuildStrategy	buildPodCreationStrategy
}

func (f *typeBasedFactoryStrategy) CreateBuildPod(build *buildv1.Build, additionalCAs map[string]string, internalRegistryHost string) (*corev1.Pod, error) {
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
	var pod *corev1.Pod
	var err error
	switch {
	case build.Spec.Strategy.DockerStrategy != nil:
		pod, err = f.dockerBuildStrategy.CreateBuildPod(build, additionalCAs, internalRegistryHost)
	case build.Spec.Strategy.SourceStrategy != nil:
		pod, err = f.sourceBuildStrategy.CreateBuildPod(build, additionalCAs, internalRegistryHost)
	case build.Spec.Strategy.CustomStrategy != nil:
		pod, err = f.customBuildStrategy.CreateBuildPod(build, additionalCAs, internalRegistryHost)
	case build.Spec.Strategy.JenkinsPipelineStrategy != nil:
		return nil, fmt.Errorf("creating a build pod for Build %s/%s with the JenkinsPipeline strategy is not supported", build.Namespace, build.Name)
	default:
		return nil, fmt.Errorf("no supported build strategy defined for Build %s/%s", build.Namespace, build.Name)
	}
	if pod != nil {
		if pod.Annotations == nil {
			pod.Annotations = map[string]string{}
		}
		pod.Annotations[buildutil.BuildAnnotation] = build.Name
	}
	return pod, err
}

package internal_helpers

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/api/apihelpers"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	buildPodSuffix = "build"
)

func BuildToPodLogOptions(opts *buildapi.BuildLogOptions) *kapi.PodLogOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &kapi.PodLogOptions{Follow: opts.Follow, SinceSeconds: opts.SinceSeconds, SinceTime: opts.SinceTime, Timestamps: opts.Timestamps, TailLines: opts.TailLines, LimitBytes: opts.LimitBytes}
}
func IsBuildComplete(b *buildapi.Build) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return IsTerminalPhase(b.Status.Phase)
}
func IsTerminalPhase(p buildapi.BuildPhase) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch p {
	case buildapi.BuildPhaseNew, buildapi.BuildPhasePending, buildapi.BuildPhaseRunning:
		return false
	}
	return true
}
func GetBuildPodName(build *buildapi.Build) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apihelpers.GetPodName(build.Name, buildPodSuffix)
}
func StrategyType(strategy buildapi.BuildStrategy) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func GetInputReference(strategy buildapi.BuildStrategy) *kapi.ObjectReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case strategy.SourceStrategy != nil:
		return &strategy.SourceStrategy.From
	case strategy.DockerStrategy != nil:
		return strategy.DockerStrategy.From
	case strategy.CustomStrategy != nil:
		return &strategy.CustomStrategy.From
	default:
		return nil
	}
}
func GetBuildEnv(build *buildapi.Build) []kapi.EnvVar {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case build.Spec.Strategy.SourceStrategy != nil:
		return build.Spec.Strategy.SourceStrategy.Env
	case build.Spec.Strategy.DockerStrategy != nil:
		return build.Spec.Strategy.DockerStrategy.Env
	case build.Spec.Strategy.CustomStrategy != nil:
		return build.Spec.Strategy.CustomStrategy.Env
	case build.Spec.Strategy.JenkinsPipelineStrategy != nil:
		return build.Spec.Strategy.JenkinsPipelineStrategy.Env
	default:
		return nil
	}
}
func SetBuildEnv(build *buildapi.Build, env []kapi.EnvVar) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var oldEnv *[]kapi.EnvVar
	switch {
	case build.Spec.Strategy.SourceStrategy != nil:
		oldEnv = &build.Spec.Strategy.SourceStrategy.Env
	case build.Spec.Strategy.DockerStrategy != nil:
		oldEnv = &build.Spec.Strategy.DockerStrategy.Env
	case build.Spec.Strategy.CustomStrategy != nil:
		oldEnv = &build.Spec.Strategy.CustomStrategy.Env
	case build.Spec.Strategy.JenkinsPipelineStrategy != nil:
		oldEnv = &build.Spec.Strategy.JenkinsPipelineStrategy.Env
	default:
		return
	}
	*oldEnv = env
}
func UpdateBuildEnv(build *buildapi.Build, env []kapi.EnvVar) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	buildEnv := GetBuildEnv(build)
	newEnv := []kapi.EnvVar{}
	for _, e := range buildEnv {
		exists := false
		for _, n := range env {
			if e.Name == n.Name {
				exists = true
				break
			}
		}
		if !exists {
			newEnv = append(newEnv, e)
		}
	}
	newEnv = append(newEnv, env...)
	SetBuildEnv(build, newEnv)
}

type BuildSliceByCreationTimestamp []buildapi.Build

func (b BuildSliceByCreationTimestamp) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(b)
}
func (b BuildSliceByCreationTimestamp) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return b[i].CreationTimestamp.Before(&b[j].CreationTimestamp)
}
func (b BuildSliceByCreationTimestamp) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b[i], b[j] = b[j], b[i]
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

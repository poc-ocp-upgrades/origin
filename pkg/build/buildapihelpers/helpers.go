package buildapihelpers

import (
	goformat "fmt"
	buildv1 "github.com/openshift/api/build/v1"
	corev1 "k8s.io/api/core/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func BuildToPodLogOptions(opts *buildv1.BuildLogOptions) *corev1.PodLogOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &corev1.PodLogOptions{Follow: opts.Follow, SinceSeconds: opts.SinceSeconds, SinceTime: opts.SinceTime, Timestamps: opts.Timestamps, TailLines: opts.TailLines, LimitBytes: opts.LimitBytes}
}
func FindTriggerPolicy(triggerType buildv1.BuildTriggerType, config *buildv1.BuildConfig) (buildTriggers []buildv1.BuildTriggerPolicy) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, specTrigger := range config.Spec.Triggers {
		if specTrigger.Type == triggerType {
			buildTriggers = append(buildTriggers, specTrigger)
		}
	}
	return buildTriggers
}
func HasTriggerType(triggerType buildv1.BuildTriggerType, bc *buildv1.BuildConfig) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	matches := FindTriggerPolicy(triggerType, bc)
	return len(matches) > 0
}
func GetInputReference(strategy buildv1.BuildStrategy) *corev1.ObjectReference {
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

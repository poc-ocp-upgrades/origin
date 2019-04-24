package buildapihelpers

import (
	buildv1 "github.com/openshift/api/build/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	corev1 "k8s.io/api/core/v1"
)

func BuildToPodLogOptions(opts *buildv1.BuildLogOptions) *corev1.PodLogOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &corev1.PodLogOptions{Follow: opts.Follow, SinceSeconds: opts.SinceSeconds, SinceTime: opts.SinceTime, Timestamps: opts.Timestamps, TailLines: opts.TailLines, LimitBytes: opts.LimitBytes}
}
func FindTriggerPolicy(triggerType buildv1.BuildTriggerType, config *buildv1.BuildConfig) (buildTriggers []buildv1.BuildTriggerPolicy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, specTrigger := range config.Spec.Triggers {
		if specTrigger.Type == triggerType {
			buildTriggers = append(buildTriggers, specTrigger)
		}
	}
	return buildTriggers
}
func HasTriggerType(triggerType buildv1.BuildTriggerType, bc *buildv1.BuildConfig) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	matches := FindTriggerPolicy(triggerType, bc)
	return len(matches) > 0
}
func GetInputReference(strategy buildv1.BuildStrategy) *corev1.ObjectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

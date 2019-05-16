package common

import (
	"errors"
	"fmt"
	goformat "fmt"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/origin/pkg/build/buildscheme"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func GetBuildFromPod(pod *corev1.Pod) (*buildv1.Build, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(pod.Spec.Containers) == 0 {
		return nil, errors.New("unable to get build from pod: pod has no containers")
	}
	buildEnvVar := getEnvVar(&pod.Spec.Containers[0], "BUILD")
	if len(buildEnvVar) == 0 {
		return nil, errors.New("unable to get build from pod: BUILD environment variable is empty")
	}
	obj, err := runtime.Decode(buildscheme.Decoder, []byte(buildEnvVar))
	if err != nil {
		return nil, fmt.Errorf("unable to get build from pod: %v", err)
	}
	build, ok := obj.(*buildv1.Build)
	if !ok {
		return nil, fmt.Errorf("unable to get build from pod: %v", errors.New("decoded object is not of type Build"))
	}
	return build, nil
}
func SetBuildInPod(pod *corev1.Pod, build *buildv1.Build) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	encodedBuild, err := runtime.Encode(buildscheme.Encoder, build)
	if err != nil {
		return fmt.Errorf("unable to set build in pod: %v", err)
	}
	for i := range pod.Spec.Containers {
		setEnvVar(&pod.Spec.Containers[i], "BUILD", string(encodedBuild))
	}
	for i := range pod.Spec.InitContainers {
		setEnvVar(&pod.Spec.InitContainers[i], "BUILD", string(encodedBuild))
	}
	return nil
}
func getEnvVar(c *corev1.Container, name string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, envVar := range c.Env {
		if envVar.Name == name {
			return envVar.Value
		}
	}
	return ""
}
func setEnvVar(c *corev1.Container, name, value string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i, envVar := range c.Env {
		if envVar.Name == name {
			c.Env[i] = corev1.EnvVar{Name: name, Value: value}
			return
		}
	}
	c.Env = append(c.Env, corev1.EnvVar{Name: name, Value: value})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

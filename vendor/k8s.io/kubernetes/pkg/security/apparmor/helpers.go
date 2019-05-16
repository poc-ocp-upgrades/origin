package apparmor

import (
	goformat "fmt"
	"k8s.io/api/core/v1"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	ContainerAnnotationKeyPrefix = "container.apparmor.security.beta.kubernetes.io/"
	DefaultProfileAnnotationKey  = "apparmor.security.beta.kubernetes.io/defaultProfileName"
	AllowedProfilesAnnotationKey = "apparmor.security.beta.kubernetes.io/allowedProfileNames"
	ProfileRuntimeDefault        = "runtime/default"
	ProfileNamePrefix            = "localhost/"
	ProfileNameUnconfined        = "unconfined"
)

func isRequired(pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for key, value := range pod.Annotations {
		if strings.HasPrefix(key, ContainerAnnotationKeyPrefix) {
			return value != ProfileNameUnconfined
		}
	}
	return false
}
func GetProfileName(pod *v1.Pod, containerName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return GetProfileNameFromPodAnnotations(pod.Annotations, containerName)
}
func GetProfileNameFromPodAnnotations(annotations map[string]string, containerName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return annotations[ContainerAnnotationKeyPrefix+containerName]
}
func SetProfileName(pod *v1.Pod, containerName, profileName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pod.Annotations == nil {
		pod.Annotations = map[string]string{}
	}
	pod.Annotations[ContainerAnnotationKeyPrefix+containerName] = profileName
	return nil
}
func SetProfileNameFromPodAnnotations(annotations map[string]string, containerName, profileName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if annotations == nil {
		return nil
	}
	annotations[ContainerAnnotationKeyPrefix+containerName] = profileName
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

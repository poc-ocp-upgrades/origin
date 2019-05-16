package seccomp

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	AllowAny                     = "*"
	DefaultProfileAnnotationKey  = "seccomp.security.alpha.kubernetes.io/defaultProfileName"
	AllowedProfilesAnnotationKey = "seccomp.security.alpha.kubernetes.io/allowedProfileNames"
)

type Strategy interface {
	Generate(annotations map[string]string, pod *api.Pod) (string, error)
	ValidatePod(pod *api.Pod) field.ErrorList
	ValidateContainer(pod *api.Pod, container *api.Container) field.ErrorList
}
type strategy struct {
	defaultProfile        string
	allowedProfiles       map[string]bool
	allowedProfilesString string
	allowAnyProfile       bool
}

var _ Strategy = &strategy{}

func NewStrategy(pspAnnotations map[string]string) Strategy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var allowedProfiles map[string]bool
	allowAnyProfile := false
	if allowed, ok := pspAnnotations[AllowedProfilesAnnotationKey]; ok {
		profiles := strings.Split(allowed, ",")
		allowedProfiles = make(map[string]bool, len(profiles))
		for _, p := range profiles {
			if p == AllowAny {
				allowAnyProfile = true
				continue
			}
			allowedProfiles[p] = true
		}
	}
	return &strategy{defaultProfile: pspAnnotations[DefaultProfileAnnotationKey], allowedProfiles: allowedProfiles, allowedProfilesString: pspAnnotations[AllowedProfilesAnnotationKey], allowAnyProfile: allowAnyProfile}
}
func (s *strategy) Generate(annotations map[string]string, pod *api.Pod) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if annotations[api.SeccompPodAnnotationKey] != "" {
		return annotations[api.SeccompPodAnnotationKey], nil
	}
	return s.defaultProfile, nil
}
func (s *strategy) ValidatePod(pod *api.Pod) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	podSpecFieldPath := field.NewPath("pod", "metadata", "annotations").Key(api.SeccompPodAnnotationKey)
	podProfile := pod.Annotations[api.SeccompPodAnnotationKey]
	if !s.allowAnyProfile && len(s.allowedProfiles) == 0 && podProfile != "" {
		allErrs = append(allErrs, field.Forbidden(podSpecFieldPath, "seccomp may not be set"))
		return allErrs
	}
	if !s.profileAllowed(podProfile) {
		msg := fmt.Sprintf("%s is not an allowed seccomp profile. Valid values are %v", podProfile, s.allowedProfilesString)
		allErrs = append(allErrs, field.Forbidden(podSpecFieldPath, msg))
	}
	return allErrs
}
func (s *strategy) ValidateContainer(pod *api.Pod, container *api.Container) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	fieldPath := field.NewPath("pod", "metadata", "annotations").Key(api.SeccompContainerAnnotationKeyPrefix + container.Name)
	containerProfile := profileForContainer(pod, container)
	if !s.allowAnyProfile && len(s.allowedProfiles) == 0 && containerProfile != "" {
		allErrs = append(allErrs, field.Forbidden(fieldPath, "seccomp may not be set"))
		return allErrs
	}
	if !s.profileAllowed(containerProfile) {
		msg := fmt.Sprintf("%s is not an allowed seccomp profile. Valid values are %v", containerProfile, s.allowedProfilesString)
		allErrs = append(allErrs, field.Forbidden(fieldPath, msg))
	}
	return allErrs
}
func (s *strategy) profileAllowed(profile string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(s.allowedProfiles) == 0 && profile == "" {
		return true
	}
	return s.allowAnyProfile || s.allowedProfiles[profile]
}
func profileForContainer(pod *api.Pod, container *api.Container) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	containerProfile, ok := pod.Annotations[api.SeccompContainerAnnotationKeyPrefix+container.Name]
	if ok {
		return containerProfile
	}
	return pod.Annotations[api.SeccompPodAnnotationKey]
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

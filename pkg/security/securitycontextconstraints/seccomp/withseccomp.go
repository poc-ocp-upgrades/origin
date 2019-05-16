package seccomp

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
)

const (
	allowAnyProfile = "*"
)

type withSeccompProfile struct{ allowedProfiles []string }

var _ SeccompStrategy = &withSeccompProfile{}

func NewWithSeccompProfile(allowedProfiles []string) (SeccompStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &withSeccompProfile{allowedProfiles}, nil
}
func (s *withSeccompProfile) Generate(pod *api.Pod) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, p := range s.allowedProfiles {
		if p != allowAnyProfile {
			return p, nil
		}
	}
	return "", nil
}
func (s *withSeccompProfile) ValidatePod(pod *api.Pod) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	fieldPath := field.NewPath("pod", "metadata", "annotations", api.SeccompPodAnnotationKey)
	podProfile, _ := pod.Annotations[api.SeccompPodAnnotationKey]
	if len(s.allowedProfiles) == 0 && podProfile != "" {
		allErrs = append(allErrs, field.Forbidden(fieldPath, "seccomp may not be set"))
		return allErrs
	}
	if !isProfileAllowed(podProfile, s.allowedProfiles) {
		msg := fmt.Sprintf("%s is not a valid seccomp profile. Valid values are %v", podProfile, s.allowedProfiles)
		allErrs = append(allErrs, field.Forbidden(fieldPath, msg))
	}
	return allErrs
}
func (s *withSeccompProfile) ValidateContainer(pod *api.Pod, container *api.Container) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	fieldPath := field.NewPath("pod", "metadata", "annotations", api.SeccompContainerAnnotationKeyPrefix+container.Name)
	containerProfile := profileForContainer(pod, container)
	if len(s.allowedProfiles) == 0 && containerProfile != "" {
		allErrs = append(allErrs, field.Forbidden(fieldPath, "seccomp may not be set"))
		return allErrs
	}
	if !isProfileAllowed(containerProfile, s.allowedProfiles) {
		msg := fmt.Sprintf("%s is not a valid seccomp profile. Valid values are %v", containerProfile, s.allowedProfiles)
		allErrs = append(allErrs, field.Forbidden(fieldPath, msg))
	}
	return allErrs
}
func isProfileAllowed(profile string, allowedProfiles []string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(allowedProfiles) == 0 && profile == "" {
		return true
	}
	for _, p := range allowedProfiles {
		if profile == p || p == allowAnyProfile {
			return true
		}
	}
	return false
}
func profileForContainer(pod *api.Pod, container *api.Container) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	containerProfile, hasContainerProfile := pod.Annotations[api.SeccompContainerAnnotationKeyPrefix+container.Name]
	if hasContainerProfile {
		return containerProfile
	}
	return pod.Annotations[api.SeccompPodAnnotationKey]
}

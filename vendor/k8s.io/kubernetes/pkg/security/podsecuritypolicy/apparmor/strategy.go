package apparmor

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/security/apparmor"
	"k8s.io/kubernetes/pkg/util/maps"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type Strategy interface {
	Generate(annotations map[string]string, container *api.Container) (map[string]string, error)
	Validate(pod *api.Pod, container *api.Container) field.ErrorList
}
type strategy struct {
	defaultProfile        string
	allowedProfiles       map[string]bool
	allowedProfilesString string
}

var _ Strategy = &strategy{}

func NewStrategy(pspAnnotations map[string]string) Strategy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var allowedProfiles map[string]bool
	if allowed, ok := pspAnnotations[apparmor.AllowedProfilesAnnotationKey]; ok {
		profiles := strings.Split(allowed, ",")
		allowedProfiles = make(map[string]bool, len(profiles))
		for _, p := range profiles {
			allowedProfiles[p] = true
		}
	}
	return &strategy{defaultProfile: pspAnnotations[apparmor.DefaultProfileAnnotationKey], allowedProfiles: allowedProfiles, allowedProfilesString: pspAnnotations[apparmor.AllowedProfilesAnnotationKey]}
}
func (s *strategy) Generate(annotations map[string]string, container *api.Container) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	copy := maps.CopySS(annotations)
	if annotations[apparmor.ContainerAnnotationKeyPrefix+container.Name] != "" {
		return copy, nil
	}
	if s.defaultProfile == "" {
		return copy, nil
	}
	if copy == nil {
		copy = map[string]string{}
	}
	copy[apparmor.ContainerAnnotationKeyPrefix+container.Name] = s.defaultProfile
	return copy, nil
}
func (s *strategy) Validate(pod *api.Pod, container *api.Container) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s.allowedProfiles == nil {
		return nil
	}
	allErrs := field.ErrorList{}
	fieldPath := field.NewPath("pod", "metadata", "annotations").Key(apparmor.ContainerAnnotationKeyPrefix + container.Name)
	profile := apparmor.GetProfileNameFromPodAnnotations(pod.Annotations, container.Name)
	if profile == "" {
		if len(s.allowedProfiles) > 0 {
			allErrs = append(allErrs, field.Forbidden(fieldPath, "AppArmor profile must be set"))
			return allErrs
		}
		return nil
	}
	if !s.allowedProfiles[profile] {
		msg := fmt.Sprintf("%s is not an allowed profile. Allowed values: %q", profile, s.allowedProfilesString)
		allErrs = append(allErrs, field.Forbidden(fieldPath, msg))
	}
	return allErrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

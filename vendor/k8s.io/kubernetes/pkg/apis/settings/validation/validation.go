package validation

import (
	goformat "fmt"
	apimachineryvalidation "k8s.io/apimachinery/pkg/api/validation"
	unversionedvalidation "k8s.io/apimachinery/pkg/apis/meta/v1/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	apivalidation "k8s.io/kubernetes/pkg/apis/core/validation"
	"k8s.io/kubernetes/pkg/apis/settings"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ValidatePodPresetName(name string, prefix bool) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apimachineryvalidation.NameIsDNSSubdomain(name, prefix)
}
func ValidatePodPresetSpec(spec *settings.PodPresetSpec, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, unversionedvalidation.ValidateLabelSelector(&spec.Selector, fldPath.Child("selector"))...)
	if spec.Env == nil && spec.EnvFrom == nil && spec.VolumeMounts == nil && spec.Volumes == nil {
		allErrs = append(allErrs, field.Required(fldPath.Child("volumes", "env", "envFrom", "volumeMounts"), "must specify at least one"))
	}
	vols, vErrs := apivalidation.ValidateVolumes(spec.Volumes, fldPath.Child("volumes"))
	allErrs = append(allErrs, vErrs...)
	allErrs = append(allErrs, apivalidation.ValidateEnv(spec.Env, fldPath.Child("env"))...)
	allErrs = append(allErrs, apivalidation.ValidateEnvFrom(spec.EnvFrom, fldPath.Child("envFrom"))...)
	allErrs = append(allErrs, apivalidation.ValidateVolumeMounts(spec.VolumeMounts, nil, vols, nil, fldPath.Child("volumeMounts"))...)
	return allErrs
}
func ValidatePodPreset(pip *settings.PodPreset) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := apivalidation.ValidateObjectMeta(&pip.ObjectMeta, true, ValidatePodPresetName, field.NewPath("metadata"))
	allErrs = append(allErrs, ValidatePodPresetSpec(&pip.Spec, field.NewPath("spec"))...)
	return allErrs
}
func ValidatePodPresetUpdate(pip, oldPip *settings.PodPreset) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := apivalidation.ValidateObjectMetaUpdate(&pip.ObjectMeta, &oldPip.ObjectMeta, field.NewPath("metadata"))
	allErrs = append(allErrs, ValidatePodPresetSpec(&pip.Spec, field.NewPath("spec"))...)
	return allErrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

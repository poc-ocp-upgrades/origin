package customresourcevalidation

import (
	"strings"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	configv1 "github.com/openshift/api/config/v1"
)

func ValidateConfigMapReference(path *field.Path, configMap configv1.ConfigMapNameReference, required bool) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validateConfigMapSecret(path.Child("name"), configMap.Name, required, validation.ValidateConfigMapName)
}
func ValidateSecretReference(path *field.Path, secret configv1.SecretNameReference, required bool) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validateConfigMapSecret(path.Child("name"), secret.Name, required, validation.ValidateSecretName)
}
func validateConfigMapSecret(path *field.Path, name string, required bool, validator validation.ValidateNameFunc) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(name) == 0 {
		if required {
			return field.ErrorList{field.Required(path, "")}
		}
		return nil
	}
	if valErrs := validator(name, false); len(valErrs) > 0 {
		return field.ErrorList{field.Invalid(path, name, strings.Join(valErrs, ", "))}
	}
	return nil
}
func RequireNameCluster(name string, prefix bool) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if name != "cluster" {
		return []string{"must be cluster"}
	}
	return nil
}

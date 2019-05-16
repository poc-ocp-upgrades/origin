package oauth

import (
	configv1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateGoogleIdentityProvider(provider *configv1.GoogleIdentityProvider, mappingMethod configv1.MappingMethodType, fieldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := field.ErrorList{}
	if provider == nil {
		errs = append(errs, field.Required(fieldPath, ""))
		return errs
	}
	errs = append(errs, ValidateOAuthIdentityProvider(provider.ClientID, provider.ClientSecret, fieldPath)...)
	if len(provider.HostedDomain) == 0 && mappingMethod != configv1.MappingMethodLookup {
		errs = append(errs, field.Invalid(fieldPath.Child("hostedDomain"), nil, "hostedDomain must be specified unless lookup is used"))
	}
	return errs
}

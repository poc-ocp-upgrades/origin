package oauth

import (
	"fmt"
	configv1 "github.com/openshift/api/config/v1"
	crvalidation "github.com/openshift/origin/pkg/admission/customresourcevalidation"
	"github.com/openshift/origin/pkg/cmd/server/apis/config/validation/common"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"strings"
)

func ValidateOpenIDIdentityProvider(provider *configv1.OpenIDIdentityProvider, fieldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if provider == nil {
		allErrs = append(allErrs, field.Required(fieldPath, ""))
		return allErrs
	}
	allErrs = append(allErrs, ValidateOAuthIdentityProvider(provider.ClientID, provider.ClientSecret, fieldPath)...)
	if provider.Issuer != strings.TrimRight(provider.Issuer, "/") {
		allErrs = append(allErrs, field.Invalid(fieldPath.Child("issuer"), provider.Issuer, "cannot end with '/'"))
	}
	url, issuerErrs := common.ValidateSecureURL(provider.Issuer, fieldPath.Child("issuer"))
	allErrs = append(allErrs, issuerErrs...)
	if len(url.RawQuery) > 0 || len(url.Fragment) > 0 {
		allErrs = append(allErrs, field.Invalid(fieldPath.Child("issuer"), provider.Issuer, "must not specify query or fragment component"))
	}
	allErrs = append(allErrs, crvalidation.ValidateConfigMapReference(fieldPath.Child("ca"), provider.CA, false)...)
	for i, scope := range provider.ExtraScopes {
		for _, ch := range scope {
			switch {
			case ch == '!':
			case ch >= '#' && ch <= '[':
			case ch >= ']' && ch <= '~':
			default:
				allErrs = append(allErrs, field.Invalid(fieldPath.Child("extraScopes").Index(i), scope, fmt.Sprintf("cannot contain %v", ch)))
			}
		}
	}
	return allErrs
}

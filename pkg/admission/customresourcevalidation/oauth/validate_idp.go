package oauth

import (
	"fmt"
	configv1 "github.com/openshift/api/config/v1"
	crvalidation "github.com/openshift/origin/pkg/admission/customresourcevalidation"
	oauthvalidation "github.com/openshift/origin/pkg/oauth/apis/oauth/validation"
	userapivalidation "github.com/openshift/origin/pkg/user/apis/user/validation"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"strings"
)

var validMappingMethods = sets.NewString(string(configv1.MappingMethodLookup), string(configv1.MappingMethodClaim), string(configv1.MappingMethodAdd))

func validateOAuthSpec(spec configv1.OAuthSpec) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := field.ErrorList{}
	specPath := field.NewPath("spec")
	providerNames := sets.NewString()
	challengeIssuingIdentityProviders := []string{}
	challengeRedirectingIdentityProviders := []string{}
	for i, identityProvider := range spec.IdentityProviders {
		if isUsedAsChallenger(identityProvider.IdentityProviderConfig) {
			if identityProvider.Type == configv1.IdentityProviderTypeRequestHeader {
				challengeRedirectingIdentityProviders = append(challengeRedirectingIdentityProviders, identityProvider.Name)
			} else {
				challengeIssuingIdentityProviders = append(challengeIssuingIdentityProviders, identityProvider.Name)
			}
		}
		identityProviderPath := specPath.Child("identityProvider").Index(i)
		errs = append(errs, ValidateIdentityProvider(identityProvider, identityProviderPath)...)
		if len(identityProvider.Name) > 0 {
			if providerNames.Has(identityProvider.Name) {
				errs = append(errs, field.Invalid(identityProviderPath.Child("name"), identityProvider.Name, "must have a unique name"))
			}
			providerNames.Insert(identityProvider.Name)
		}
	}
	if len(challengeRedirectingIdentityProviders) > 1 {
		errs = append(errs, field.Invalid(specPath.Child("identityProviders"), "<omitted>", fmt.Sprintf("only one identity provider can redirect clients requesting an authentication challenge, found: %v", strings.Join(challengeRedirectingIdentityProviders, ", "))))
	}
	if len(challengeRedirectingIdentityProviders) > 0 && len(challengeIssuingIdentityProviders) > 0 {
		errs = append(errs, field.Invalid(specPath.Child("identityProviders"), "<omitted>", fmt.Sprintf("cannot mix providers that redirect clients requesting auth challenges (%s) with providers issuing challenges to those clients (%s)", strings.Join(challengeRedirectingIdentityProviders, ", "), strings.Join(challengeIssuingIdentityProviders, ", "))))
	}
	timeout := spec.TokenConfig.AccessTokenInactivityTimeoutSeconds
	if timeout > 0 && timeout < oauthvalidation.MinimumInactivityTimeoutSeconds {
		errs = append(errs, field.Invalid(specPath.Child("tokenConfig", "accessTokenInactivityTimeoutSeconds"), timeout, fmt.Sprintf("the minimum acceptable token timeout value is %d seconds", oauthvalidation.MinimumInactivityTimeoutSeconds)))
	}
	if tokenMaxAge := spec.TokenConfig.AccessTokenMaxAgeSeconds; tokenMaxAge < 0 {
		errs = append(errs, field.Invalid(specPath.Child("tokenConfig", "accessTokenMaxAgeSeconds"), tokenMaxAge, "must be a positive integer or 0"))
	}
	errs = append(errs, crvalidation.ValidateSecretReference(specPath.Child("templates", "login"), spec.Templates.Login, false)...)
	errs = append(errs, crvalidation.ValidateSecretReference(specPath.Child("templates", "providerSelection"), spec.Templates.ProviderSelection, false)...)
	errs = append(errs, crvalidation.ValidateSecretReference(specPath.Child("templates", "error"), spec.Templates.Error, false)...)
	return errs
}
func ValidateIdentityProvider(identityProvider configv1.IdentityProvider, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := field.ErrorList{}
	if len(identityProvider.Name) == 0 {
		errs = append(errs, field.Required(fldPath.Child("name"), ""))
	} else if reasons := userapivalidation.ValidateIdentityProviderName(identityProvider.Name); len(reasons) != 0 {
		errs = append(errs, field.Invalid(fldPath.Child("name"), identityProvider.Name, strings.Join(reasons, ", ")))
	}
	if len(identityProvider.MappingMethod) > 0 && !validMappingMethods.Has(string(identityProvider.MappingMethod)) {
		errs = append(errs, field.NotSupported(fldPath.Child("mappingMethod"), identityProvider.MappingMethod, validMappingMethods.List()))
	}
	provider := identityProvider.IdentityProviderConfig
	switch provider.Type {
	case "":
		errs = append(errs, field.Required(fldPath.Child("type"), ""))
	case configv1.IdentityProviderTypeRequestHeader:
		errs = append(errs, ValidateRequestHeaderIdentityProvider(provider.RequestHeader, fldPath)...)
	case configv1.IdentityProviderTypeBasicAuth:
		if provider.BasicAuth == nil {
			errs = append(errs, field.Required(fldPath.Child("basicAuth"), ""))
		} else {
			errs = append(errs, ValidateRemoteConnectionInfo(provider.BasicAuth.OAuthRemoteConnectionInfo, fldPath.Child("basicAuth"))...)
		}
	case configv1.IdentityProviderTypeHTPasswd:
		if provider.HTPasswd == nil {
			errs = append(errs, field.Required(fldPath.Child("htpasswd"), ""))
		} else {
			errs = append(errs, crvalidation.ValidateSecretReference(fldPath.Child("htpasswd", "fileData"), provider.HTPasswd.FileData, true)...)
		}
	case configv1.IdentityProviderTypeLDAP:
		errs = append(errs, ValidateLDAPIdentityProvider(provider.LDAP, fldPath.Child("ldap"))...)
	case configv1.IdentityProviderTypeKeystone:
		errs = append(errs, ValidateKeystoneIdentityProvider(provider.Keystone, fldPath.Child("keystone"))...)
	case configv1.IdentityProviderTypeGitHub:
		errs = append(errs, ValidateGitHubIdentityProvider(provider.GitHub, identityProvider.MappingMethod, fldPath.Child("github"))...)
	case configv1.IdentityProviderTypeGitLab:
		errs = append(errs, ValidateGitLabIdentityProvider(provider.GitLab, fldPath.Child("gitlab"))...)
	case configv1.IdentityProviderTypeGoogle:
		errs = append(errs, ValidateGoogleIdentityProvider(provider.Google, identityProvider.MappingMethod, fldPath.Child("google"))...)
	case configv1.IdentityProviderTypeOpenID:
		errs = append(errs, ValidateOpenIDIdentityProvider(provider.OpenID, fldPath.Child("openID"))...)
	default:
		errs = append(errs, field.Invalid(fldPath.Child("type"), identityProvider.Type, "not a valid provider type"))
	}
	return errs
}
func ValidateOAuthIdentityProvider(clientID string, clientSecretRef configv1.SecretNameReference, fieldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if len(clientID) == 0 {
		allErrs = append(allErrs, field.Required(fieldPath.Child("clientID"), ""))
	}
	allErrs = append(allErrs, crvalidation.ValidateSecretReference(fieldPath.Child("clientSecret"), clientSecretRef, true)...)
	return allErrs
}
func isUsedAsChallenger(idp configv1.IdentityProviderConfig) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch idp.Type {
	case configv1.IdentityProviderTypeBasicAuth, configv1.IdentityProviderTypeGitLab, configv1.IdentityProviderTypeHTPasswd, configv1.IdentityProviderTypeKeystone, configv1.IdentityProviderTypeLDAP, configv1.IdentityProviderTypeOpenID:
		return true
	case configv1.IdentityProviderTypeRequestHeader:
		if idp.RequestHeader == nil {
			return false
		}
		return len(idp.RequestHeader.ChallengeURL) > 0
	default:
		return false
	}
}

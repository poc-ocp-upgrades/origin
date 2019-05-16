package oauth

import (
	"fmt"
	configv1 "github.com/openshift/api/config/v1"
	crvalidation "github.com/openshift/origin/pkg/admission/customresourcevalidation"
	"github.com/openshift/origin/pkg/cmd/server/apis/config/validation/common"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"net/url"
	"path"
	"strings"
)

const (
	urlToken   = "${url}"
	queryToken = "${query}"
)

func ValidateRequestHeaderIdentityProvider(provider *configv1.RequestHeaderIdentityProvider, fieldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := field.ErrorList{}
	if provider == nil {
		errs = append(errs, field.Required(fieldPath, ""))
		return errs
	}
	errs = append(errs, crvalidation.ValidateConfigMapReference(fieldPath.Child("ca"), provider.ClientCA, true)...)
	if len(provider.Headers) == 0 {
		errs = append(errs, field.Required(fieldPath.Child("headers"), ""))
	}
	if len(provider.ChallengeURL) == 0 && len(provider.LoginURL) == 0 {
		errs = append(errs, field.Required(fieldPath, "at least one of challengeURL or loginURL must be specified"))
	}
	if len(provider.ChallengeURL) > 0 {
		u, urlErrs := common.ValidateURL(provider.ChallengeURL, fieldPath.Child("challengeURL"))
		errs = append(errs, urlErrs...)
		if len(urlErrs) == 0 {
			if !hasParamToken(u) {
				errs = append(errs, field.Invalid(field.NewPath("challengeURL"), provider.ChallengeURL, fmt.Sprintf("query does not include %q or %q, redirect will not preserve original authorize parameters", urlToken, queryToken)))
			}
		}
	}
	if len(provider.LoginURL) > 0 {
		u, urlErrs := common.ValidateURL(provider.LoginURL, fieldPath.Child("loginURL"))
		errs = append(errs, urlErrs...)
		if len(urlErrs) == 0 {
			if !hasParamToken(u) {
				errs = append(errs, field.Invalid(fieldPath.Child("loginURL"), provider.LoginURL, fmt.Sprintf("query does not include %q or %q, redirect will not preserve original authorize parameters", urlToken, queryToken)))
			}
			if strings.HasSuffix(u.Path, "/") {
				errs = append(errs, field.Invalid(fieldPath.Child("loginURL"), provider.LoginURL, `path ends with "/", grant approval flows will not function correctly`))
			}
			if _, file := path.Split(u.Path); file != "authorize" {
				errs = append(errs, field.Invalid(fieldPath.Child("loginURL"), provider.LoginURL, `path does not end with "/authorize", grant approval flows will not function correctly`))
			}
		}
	}
	return errs
}
func hasParamToken(u *url.URL) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.Contains(u.RawQuery, urlToken) || strings.Contains(u.RawQuery, queryToken)
}

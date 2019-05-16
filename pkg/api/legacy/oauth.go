package legacy

import (
	oauthv1 "github.com/openshift/api/oauth/v1"
	"github.com/openshift/origin/pkg/oauth/apis/oauth"
	oauthv1helpers "github.com/openshift/origin/pkg/oauth/apis/oauth/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func InstallInternalLegacyOAuth(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	InstallExternalLegacyOAuth(scheme)
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedInternalOAuthTypes, addLegacyOAuthFieldSelectorKeyConversions, oauthv1helpers.RegisterDefaults, oauthv1helpers.RegisterConversions)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func InstallExternalLegacyOAuth(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedOAuthTypes)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func addUngroupifiedOAuthTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	types := []runtime.Object{&oauthv1.OAuthAccessToken{}, &oauthv1.OAuthAccessTokenList{}, &oauthv1.OAuthAuthorizeToken{}, &oauthv1.OAuthAuthorizeTokenList{}, &oauthv1.OAuthClient{}, &oauthv1.OAuthClientList{}, &oauthv1.OAuthClientAuthorization{}, &oauthv1.OAuthClientAuthorizationList{}, &oauthv1.OAuthRedirectReference{}}
	scheme.AddKnownTypes(GroupVersion, types...)
	return nil
}
func addUngroupifiedInternalOAuthTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme.AddKnownTypes(InternalGroupVersion, &oauth.OAuthAccessToken{}, &oauth.OAuthAccessTokenList{}, &oauth.OAuthAuthorizeToken{}, &oauth.OAuthAuthorizeTokenList{}, &oauth.OAuthClient{}, &oauth.OAuthClientList{}, &oauth.OAuthClientAuthorization{}, &oauth.OAuthClientAuthorizationList{}, &oauth.OAuthRedirectReference{})
	return nil
}
func addLegacyOAuthFieldSelectorKeyConversions(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := scheme.AddFieldLabelConversionFunc(GroupVersion.WithKind("OAuthAccessToken"), legacyOAuthAccessTokenFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	if err := scheme.AddFieldLabelConversionFunc(GroupVersion.WithKind("OAuthAuthorizeToken"), legacyOAuthAuthorizeTokenFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	if err := scheme.AddFieldLabelConversionFunc(GroupVersion.WithKind("OAuthClientAuthorization"), legacyOAuthClientAuthorizationFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	return nil
}
func legacyOAuthAccessTokenFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch label {
	case "clientName", "userName", "userUID", "authorizeToken":
		return label, value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}
func legacyOAuthAuthorizeTokenFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch label {
	case "clientName", "userName", "userUID":
		return label, value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}
func legacyOAuthClientAuthorizationFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch label {
	case "clientName", "userName", "userUID":
		return label, value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}

package oauth

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	GroupName = "oauth.openshift.io"
)

var (
	schemeBuilder		= runtime.NewSchemeBuilder(addKnownTypes)
	Install			= schemeBuilder.AddToScheme
	SchemeGroupVersion	= schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	AddToScheme		= schemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(SchemeGroupVersion, &OAuthAccessToken{}, &OAuthAccessTokenList{}, &OAuthAuthorizeToken{}, &OAuthAuthorizeTokenList{}, &OAuthClient{}, &OAuthClientList{}, &OAuthClientAuthorization{}, &OAuthClientAuthorizationList{}, &OAuthRedirectReference{})
	return nil
}

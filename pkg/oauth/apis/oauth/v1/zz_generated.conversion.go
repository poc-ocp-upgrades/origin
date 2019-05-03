package v1

import (
	v1 "github.com/openshift/api/oauth/v1"
	oauth "github.com/openshift/origin/pkg/oauth/apis/oauth"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*v1.ClusterRoleScopeRestriction)(nil), (*oauth.ClusterRoleScopeRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleScopeRestriction_To_oauth_ClusterRoleScopeRestriction(a.(*v1.ClusterRoleScopeRestriction), b.(*oauth.ClusterRoleScopeRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.ClusterRoleScopeRestriction)(nil), (*v1.ClusterRoleScopeRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_ClusterRoleScopeRestriction_To_v1_ClusterRoleScopeRestriction(a.(*oauth.ClusterRoleScopeRestriction), b.(*v1.ClusterRoleScopeRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthAccessToken)(nil), (*oauth.OAuthAccessToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthAccessToken_To_oauth_OAuthAccessToken(a.(*v1.OAuthAccessToken), b.(*oauth.OAuthAccessToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.OAuthAccessToken)(nil), (*v1.OAuthAccessToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_OAuthAccessToken_To_v1_OAuthAccessToken(a.(*oauth.OAuthAccessToken), b.(*v1.OAuthAccessToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthAccessTokenList)(nil), (*oauth.OAuthAccessTokenList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthAccessTokenList_To_oauth_OAuthAccessTokenList(a.(*v1.OAuthAccessTokenList), b.(*oauth.OAuthAccessTokenList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.OAuthAccessTokenList)(nil), (*v1.OAuthAccessTokenList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(a.(*oauth.OAuthAccessTokenList), b.(*v1.OAuthAccessTokenList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthAuthorizeToken)(nil), (*oauth.OAuthAuthorizeToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthAuthorizeToken_To_oauth_OAuthAuthorizeToken(a.(*v1.OAuthAuthorizeToken), b.(*oauth.OAuthAuthorizeToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.OAuthAuthorizeToken)(nil), (*v1.OAuthAuthorizeToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(a.(*oauth.OAuthAuthorizeToken), b.(*v1.OAuthAuthorizeToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthAuthorizeTokenList)(nil), (*oauth.OAuthAuthorizeTokenList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthAuthorizeTokenList_To_oauth_OAuthAuthorizeTokenList(a.(*v1.OAuthAuthorizeTokenList), b.(*oauth.OAuthAuthorizeTokenList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.OAuthAuthorizeTokenList)(nil), (*v1.OAuthAuthorizeTokenList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(a.(*oauth.OAuthAuthorizeTokenList), b.(*v1.OAuthAuthorizeTokenList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthClient)(nil), (*oauth.OAuthClient)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthClient_To_oauth_OAuthClient(a.(*v1.OAuthClient), b.(*oauth.OAuthClient), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.OAuthClient)(nil), (*v1.OAuthClient)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_OAuthClient_To_v1_OAuthClient(a.(*oauth.OAuthClient), b.(*v1.OAuthClient), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthClientAuthorization)(nil), (*oauth.OAuthClientAuthorization)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthClientAuthorization_To_oauth_OAuthClientAuthorization(a.(*v1.OAuthClientAuthorization), b.(*oauth.OAuthClientAuthorization), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.OAuthClientAuthorization)(nil), (*v1.OAuthClientAuthorization)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(a.(*oauth.OAuthClientAuthorization), b.(*v1.OAuthClientAuthorization), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthClientAuthorizationList)(nil), (*oauth.OAuthClientAuthorizationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthClientAuthorizationList_To_oauth_OAuthClientAuthorizationList(a.(*v1.OAuthClientAuthorizationList), b.(*oauth.OAuthClientAuthorizationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.OAuthClientAuthorizationList)(nil), (*v1.OAuthClientAuthorizationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(a.(*oauth.OAuthClientAuthorizationList), b.(*v1.OAuthClientAuthorizationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthClientList)(nil), (*oauth.OAuthClientList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthClientList_To_oauth_OAuthClientList(a.(*v1.OAuthClientList), b.(*oauth.OAuthClientList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.OAuthClientList)(nil), (*v1.OAuthClientList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_OAuthClientList_To_v1_OAuthClientList(a.(*oauth.OAuthClientList), b.(*v1.OAuthClientList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.OAuthRedirectReference)(nil), (*oauth.OAuthRedirectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_OAuthRedirectReference_To_oauth_OAuthRedirectReference(a.(*v1.OAuthRedirectReference), b.(*oauth.OAuthRedirectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.OAuthRedirectReference)(nil), (*v1.OAuthRedirectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_OAuthRedirectReference_To_v1_OAuthRedirectReference(a.(*oauth.OAuthRedirectReference), b.(*v1.OAuthRedirectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RedirectReference)(nil), (*oauth.RedirectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RedirectReference_To_oauth_RedirectReference(a.(*v1.RedirectReference), b.(*oauth.RedirectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.RedirectReference)(nil), (*v1.RedirectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_RedirectReference_To_v1_RedirectReference(a.(*oauth.RedirectReference), b.(*v1.RedirectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ScopeRestriction)(nil), (*oauth.ScopeRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ScopeRestriction_To_oauth_ScopeRestriction(a.(*v1.ScopeRestriction), b.(*oauth.ScopeRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*oauth.ScopeRestriction)(nil), (*v1.ScopeRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_oauth_ScopeRestriction_To_v1_ScopeRestriction(a.(*oauth.ScopeRestriction), b.(*v1.ScopeRestriction), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_ClusterRoleScopeRestriction_To_oauth_ClusterRoleScopeRestriction(in *v1.ClusterRoleScopeRestriction, out *oauth.ClusterRoleScopeRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RoleNames = *(*[]string)(unsafe.Pointer(&in.RoleNames))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.AllowEscalation = in.AllowEscalation
	return nil
}
func Convert_v1_ClusterRoleScopeRestriction_To_oauth_ClusterRoleScopeRestriction(in *v1.ClusterRoleScopeRestriction, out *oauth.ClusterRoleScopeRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterRoleScopeRestriction_To_oauth_ClusterRoleScopeRestriction(in, out, s)
}
func autoConvert_oauth_ClusterRoleScopeRestriction_To_v1_ClusterRoleScopeRestriction(in *oauth.ClusterRoleScopeRestriction, out *v1.ClusterRoleScopeRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RoleNames = *(*[]string)(unsafe.Pointer(&in.RoleNames))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.AllowEscalation = in.AllowEscalation
	return nil
}
func Convert_oauth_ClusterRoleScopeRestriction_To_v1_ClusterRoleScopeRestriction(in *oauth.ClusterRoleScopeRestriction, out *v1.ClusterRoleScopeRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_ClusterRoleScopeRestriction_To_v1_ClusterRoleScopeRestriction(in, out, s)
}
func autoConvert_v1_OAuthAccessToken_To_oauth_OAuthAccessToken(in *v1.OAuthAccessToken, out *oauth.OAuthAccessToken, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	out.Scopes = *(*[]string)(unsafe.Pointer(&in.Scopes))
	out.RedirectURI = in.RedirectURI
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.AuthorizeToken = in.AuthorizeToken
	out.RefreshToken = in.RefreshToken
	out.InactivityTimeoutSeconds = in.InactivityTimeoutSeconds
	return nil
}
func Convert_v1_OAuthAccessToken_To_oauth_OAuthAccessToken(in *v1.OAuthAccessToken, out *oauth.OAuthAccessToken, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthAccessToken_To_oauth_OAuthAccessToken(in, out, s)
}
func autoConvert_oauth_OAuthAccessToken_To_v1_OAuthAccessToken(in *oauth.OAuthAccessToken, out *v1.OAuthAccessToken, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	out.Scopes = *(*[]string)(unsafe.Pointer(&in.Scopes))
	out.RedirectURI = in.RedirectURI
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.AuthorizeToken = in.AuthorizeToken
	out.RefreshToken = in.RefreshToken
	out.InactivityTimeoutSeconds = in.InactivityTimeoutSeconds
	return nil
}
func Convert_oauth_OAuthAccessToken_To_v1_OAuthAccessToken(in *oauth.OAuthAccessToken, out *v1.OAuthAccessToken, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_OAuthAccessToken_To_v1_OAuthAccessToken(in, out, s)
}
func autoConvert_v1_OAuthAccessTokenList_To_oauth_OAuthAccessTokenList(in *v1.OAuthAccessTokenList, out *oauth.OAuthAccessTokenList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]oauth.OAuthAccessToken)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_OAuthAccessTokenList_To_oauth_OAuthAccessTokenList(in *v1.OAuthAccessTokenList, out *oauth.OAuthAccessTokenList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthAccessTokenList_To_oauth_OAuthAccessTokenList(in, out, s)
}
func autoConvert_oauth_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(in *oauth.OAuthAccessTokenList, out *v1.OAuthAccessTokenList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.OAuthAccessToken)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_oauth_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(in *oauth.OAuthAccessTokenList, out *v1.OAuthAccessTokenList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(in, out, s)
}
func autoConvert_v1_OAuthAuthorizeToken_To_oauth_OAuthAuthorizeToken(in *v1.OAuthAuthorizeToken, out *oauth.OAuthAuthorizeToken, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	out.Scopes = *(*[]string)(unsafe.Pointer(&in.Scopes))
	out.RedirectURI = in.RedirectURI
	out.State = in.State
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.CodeChallenge = in.CodeChallenge
	out.CodeChallengeMethod = in.CodeChallengeMethod
	return nil
}
func Convert_v1_OAuthAuthorizeToken_To_oauth_OAuthAuthorizeToken(in *v1.OAuthAuthorizeToken, out *oauth.OAuthAuthorizeToken, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthAuthorizeToken_To_oauth_OAuthAuthorizeToken(in, out, s)
}
func autoConvert_oauth_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(in *oauth.OAuthAuthorizeToken, out *v1.OAuthAuthorizeToken, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	out.Scopes = *(*[]string)(unsafe.Pointer(&in.Scopes))
	out.RedirectURI = in.RedirectURI
	out.State = in.State
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.CodeChallenge = in.CodeChallenge
	out.CodeChallengeMethod = in.CodeChallengeMethod
	return nil
}
func Convert_oauth_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(in *oauth.OAuthAuthorizeToken, out *v1.OAuthAuthorizeToken, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(in, out, s)
}
func autoConvert_v1_OAuthAuthorizeTokenList_To_oauth_OAuthAuthorizeTokenList(in *v1.OAuthAuthorizeTokenList, out *oauth.OAuthAuthorizeTokenList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]oauth.OAuthAuthorizeToken)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_OAuthAuthorizeTokenList_To_oauth_OAuthAuthorizeTokenList(in *v1.OAuthAuthorizeTokenList, out *oauth.OAuthAuthorizeTokenList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthAuthorizeTokenList_To_oauth_OAuthAuthorizeTokenList(in, out, s)
}
func autoConvert_oauth_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(in *oauth.OAuthAuthorizeTokenList, out *v1.OAuthAuthorizeTokenList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.OAuthAuthorizeToken)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_oauth_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(in *oauth.OAuthAuthorizeTokenList, out *v1.OAuthAuthorizeTokenList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(in, out, s)
}
func autoConvert_v1_OAuthClient_To_oauth_OAuthClient(in *v1.OAuthClient, out *oauth.OAuthClient, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Secret = in.Secret
	out.AdditionalSecrets = *(*[]string)(unsafe.Pointer(&in.AdditionalSecrets))
	out.RespondWithChallenges = in.RespondWithChallenges
	out.RedirectURIs = *(*[]string)(unsafe.Pointer(&in.RedirectURIs))
	out.GrantMethod = oauth.GrantHandlerType(in.GrantMethod)
	out.ScopeRestrictions = *(*[]oauth.ScopeRestriction)(unsafe.Pointer(&in.ScopeRestrictions))
	out.AccessTokenMaxAgeSeconds = (*int32)(unsafe.Pointer(in.AccessTokenMaxAgeSeconds))
	out.AccessTokenInactivityTimeoutSeconds = (*int32)(unsafe.Pointer(in.AccessTokenInactivityTimeoutSeconds))
	return nil
}
func Convert_v1_OAuthClient_To_oauth_OAuthClient(in *v1.OAuthClient, out *oauth.OAuthClient, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthClient_To_oauth_OAuthClient(in, out, s)
}
func autoConvert_oauth_OAuthClient_To_v1_OAuthClient(in *oauth.OAuthClient, out *v1.OAuthClient, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Secret = in.Secret
	out.AdditionalSecrets = *(*[]string)(unsafe.Pointer(&in.AdditionalSecrets))
	out.RespondWithChallenges = in.RespondWithChallenges
	out.RedirectURIs = *(*[]string)(unsafe.Pointer(&in.RedirectURIs))
	out.GrantMethod = v1.GrantHandlerType(in.GrantMethod)
	out.ScopeRestrictions = *(*[]v1.ScopeRestriction)(unsafe.Pointer(&in.ScopeRestrictions))
	out.AccessTokenMaxAgeSeconds = (*int32)(unsafe.Pointer(in.AccessTokenMaxAgeSeconds))
	out.AccessTokenInactivityTimeoutSeconds = (*int32)(unsafe.Pointer(in.AccessTokenInactivityTimeoutSeconds))
	return nil
}
func Convert_oauth_OAuthClient_To_v1_OAuthClient(in *oauth.OAuthClient, out *v1.OAuthClient, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_OAuthClient_To_v1_OAuthClient(in, out, s)
}
func autoConvert_v1_OAuthClientAuthorization_To_oauth_OAuthClientAuthorization(in *v1.OAuthClientAuthorization, out *oauth.OAuthClientAuthorization, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.ClientName = in.ClientName
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.Scopes = *(*[]string)(unsafe.Pointer(&in.Scopes))
	return nil
}
func Convert_v1_OAuthClientAuthorization_To_oauth_OAuthClientAuthorization(in *v1.OAuthClientAuthorization, out *oauth.OAuthClientAuthorization, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthClientAuthorization_To_oauth_OAuthClientAuthorization(in, out, s)
}
func autoConvert_oauth_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(in *oauth.OAuthClientAuthorization, out *v1.OAuthClientAuthorization, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.ClientName = in.ClientName
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.Scopes = *(*[]string)(unsafe.Pointer(&in.Scopes))
	return nil
}
func Convert_oauth_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(in *oauth.OAuthClientAuthorization, out *v1.OAuthClientAuthorization, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(in, out, s)
}
func autoConvert_v1_OAuthClientAuthorizationList_To_oauth_OAuthClientAuthorizationList(in *v1.OAuthClientAuthorizationList, out *oauth.OAuthClientAuthorizationList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]oauth.OAuthClientAuthorization)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_OAuthClientAuthorizationList_To_oauth_OAuthClientAuthorizationList(in *v1.OAuthClientAuthorizationList, out *oauth.OAuthClientAuthorizationList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthClientAuthorizationList_To_oauth_OAuthClientAuthorizationList(in, out, s)
}
func autoConvert_oauth_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(in *oauth.OAuthClientAuthorizationList, out *v1.OAuthClientAuthorizationList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.OAuthClientAuthorization)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_oauth_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(in *oauth.OAuthClientAuthorizationList, out *v1.OAuthClientAuthorizationList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(in, out, s)
}
func autoConvert_v1_OAuthClientList_To_oauth_OAuthClientList(in *v1.OAuthClientList, out *oauth.OAuthClientList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]oauth.OAuthClient)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_OAuthClientList_To_oauth_OAuthClientList(in *v1.OAuthClientList, out *oauth.OAuthClientList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthClientList_To_oauth_OAuthClientList(in, out, s)
}
func autoConvert_oauth_OAuthClientList_To_v1_OAuthClientList(in *oauth.OAuthClientList, out *v1.OAuthClientList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.OAuthClient)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_oauth_OAuthClientList_To_v1_OAuthClientList(in *oauth.OAuthClientList, out *v1.OAuthClientList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_OAuthClientList_To_v1_OAuthClientList(in, out, s)
}
func autoConvert_v1_OAuthRedirectReference_To_oauth_OAuthRedirectReference(in *v1.OAuthRedirectReference, out *oauth.OAuthRedirectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_RedirectReference_To_oauth_RedirectReference(&in.Reference, &out.Reference, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_OAuthRedirectReference_To_oauth_OAuthRedirectReference(in *v1.OAuthRedirectReference, out *oauth.OAuthRedirectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_OAuthRedirectReference_To_oauth_OAuthRedirectReference(in, out, s)
}
func autoConvert_oauth_OAuthRedirectReference_To_v1_OAuthRedirectReference(in *oauth.OAuthRedirectReference, out *v1.OAuthRedirectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_oauth_RedirectReference_To_v1_RedirectReference(&in.Reference, &out.Reference, s); err != nil {
		return err
	}
	return nil
}
func Convert_oauth_OAuthRedirectReference_To_v1_OAuthRedirectReference(in *oauth.OAuthRedirectReference, out *v1.OAuthRedirectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_OAuthRedirectReference_To_v1_OAuthRedirectReference(in, out, s)
}
func autoConvert_v1_RedirectReference_To_oauth_RedirectReference(in *v1.RedirectReference, out *oauth.RedirectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Group = in.Group
	out.Kind = in.Kind
	out.Name = in.Name
	return nil
}
func Convert_v1_RedirectReference_To_oauth_RedirectReference(in *v1.RedirectReference, out *oauth.RedirectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RedirectReference_To_oauth_RedirectReference(in, out, s)
}
func autoConvert_oauth_RedirectReference_To_v1_RedirectReference(in *oauth.RedirectReference, out *v1.RedirectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Group = in.Group
	out.Kind = in.Kind
	out.Name = in.Name
	return nil
}
func Convert_oauth_RedirectReference_To_v1_RedirectReference(in *oauth.RedirectReference, out *v1.RedirectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_RedirectReference_To_v1_RedirectReference(in, out, s)
}
func autoConvert_v1_ScopeRestriction_To_oauth_ScopeRestriction(in *v1.ScopeRestriction, out *oauth.ScopeRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ExactValues = *(*[]string)(unsafe.Pointer(&in.ExactValues))
	out.ClusterRole = (*oauth.ClusterRoleScopeRestriction)(unsafe.Pointer(in.ClusterRole))
	return nil
}
func Convert_v1_ScopeRestriction_To_oauth_ScopeRestriction(in *v1.ScopeRestriction, out *oauth.ScopeRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ScopeRestriction_To_oauth_ScopeRestriction(in, out, s)
}
func autoConvert_oauth_ScopeRestriction_To_v1_ScopeRestriction(in *oauth.ScopeRestriction, out *v1.ScopeRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ExactValues = *(*[]string)(unsafe.Pointer(&in.ExactValues))
	out.ClusterRole = (*v1.ClusterRoleScopeRestriction)(unsafe.Pointer(in.ClusterRole))
	return nil
}
func Convert_oauth_ScopeRestriction_To_v1_ScopeRestriction(in *oauth.ScopeRestriction, out *v1.ScopeRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_oauth_ScopeRestriction_To_v1_ScopeRestriction(in, out, s)
}

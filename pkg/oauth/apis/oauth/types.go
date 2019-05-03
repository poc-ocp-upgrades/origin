package oauth

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OAuthAccessToken struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	ClientName               string
	ExpiresIn                int64
	Scopes                   []string
	RedirectURI              string
	UserName                 string
	UserUID                  string
	AuthorizeToken           string
	RefreshToken             string
	InactivityTimeoutSeconds int32
}
type OAuthAuthorizeToken struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	ClientName          string
	ExpiresIn           int64
	Scopes              []string
	RedirectURI         string
	State               string
	UserName            string
	UserUID             string
	CodeChallenge       string
	CodeChallengeMethod string
}
type OAuthClient struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Secret                              string
	AdditionalSecrets                   []string
	RespondWithChallenges               bool
	RedirectURIs                        []string
	GrantMethod                         GrantHandlerType
	ScopeRestrictions                   []ScopeRestriction
	AccessTokenMaxAgeSeconds            *int32
	AccessTokenInactivityTimeoutSeconds *int32
}
type GrantHandlerType string

const (
	GrantHandlerAuto   GrantHandlerType = "auto"
	GrantHandlerPrompt GrantHandlerType = "prompt"
	GrantHandlerDeny   GrantHandlerType = "deny"
)

type ScopeRestriction struct {
	ExactValues []string
	ClusterRole *ClusterRoleScopeRestriction
}
type ClusterRoleScopeRestriction struct {
	RoleNames       []string
	Namespaces      []string
	AllowEscalation bool
}
type OAuthClientAuthorization struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	ClientName string
	UserName   string
	UserUID    string
	Scopes     []string
}
type OAuthAccessTokenList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []OAuthAccessToken
}
type OAuthAuthorizeTokenList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []OAuthAuthorizeToken
}
type OAuthClientList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []OAuthClient
}
type OAuthClientAuthorizationList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []OAuthClientAuthorization
}
type OAuthRedirectReference struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Reference RedirectReference
}
type RedirectReference struct {
	Group string
	Kind  string
	Name  string
}

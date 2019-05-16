package api

import (
	goformat "fmt"
	oauthapi "github.com/openshift/api/oauth/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/authentication/user"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	IdentityDisplayNameKey       = "name"
	IdentityEmailKey             = "email"
	IdentityPreferredUsernameKey = "preferred_username"
)

type UserIdentityInfo interface {
	GetIdentityName() string
	GetProviderName() string
	GetProviderUserName() string
	GetExtra() map[string]string
}
type UserIdentityMapper interface {
	UserFor(identityInfo UserIdentityInfo) (user.Info, error)
}
type Client interface {
	GetId() string
	GetSecret() string
	GetRedirectUri() string
	GetUserData() interface{}
}
type Grant struct {
	Client      Client
	Scope       string
	Expiration  int64
	RedirectURI string
}
type DefaultUserIdentityInfo struct {
	ProviderName     string
	ProviderUserName string
	Extra            map[string]string
}

func NewDefaultUserIdentityInfo(providerName, providerUserName string) *DefaultUserIdentityInfo {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &DefaultUserIdentityInfo{ProviderName: providerName, ProviderUserName: providerUserName, Extra: map[string]string{}}
}
func (i *DefaultUserIdentityInfo) GetIdentityName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return i.ProviderName + ":" + i.ProviderUserName
}
func (i *DefaultUserIdentityInfo) GetProviderName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return i.ProviderName
}
func (i *DefaultUserIdentityInfo) GetProviderUserName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return i.ProviderUserName
}
func (i *DefaultUserIdentityInfo) GetExtra() map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return i.Extra
}

type ProviderInfo struct {
	Name string
	URL  string
}
type OAuthClientGetter interface {
	Get(name string, options metav1.GetOptions) (*oauthapi.OAuthClient, error)
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

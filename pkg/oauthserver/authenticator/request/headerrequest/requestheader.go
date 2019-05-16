package headerrequest

import (
	goformat "fmt"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/identitymapper"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type Config struct {
	IDHeaders                []string
	NameHeaders              []string
	PreferredUsernameHeaders []string
	EmailHeaders             []string
}
type Authenticator struct {
	providerName string
	config       *Config
	mapper       authapi.UserIdentityMapper
}

func NewAuthenticator(providerName string, config *Config, mapper authapi.UserIdentityMapper) *Authenticator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Authenticator{providerName, config, mapper}
}
func (a *Authenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	id := headerValue(req.Header, a.config.IDHeaders)
	if len(id) == 0 {
		return nil, false, nil
	}
	identity := authapi.NewDefaultUserIdentityInfo(a.providerName, id)
	if email := headerValue(req.Header, a.config.EmailHeaders); len(email) > 0 {
		identity.Extra[authapi.IdentityEmailKey] = email
	}
	if name := headerValue(req.Header, a.config.NameHeaders); len(name) > 0 {
		identity.Extra[authapi.IdentityDisplayNameKey] = name
	}
	if preferredUsername := headerValue(req.Header, a.config.PreferredUsernameHeaders); len(preferredUsername) > 0 {
		identity.Extra[authapi.IdentityPreferredUsernameKey] = preferredUsername
	}
	return identitymapper.ResponseFor(a.mapper, identity)
}
func headerValue(h http.Header, headerNames []string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, headerName := range headerNames {
		headerName = strings.TrimSpace(headerName)
		if len(headerName) == 0 {
			continue
		}
		headerValue := h.Get(headerName)
		if len(headerValue) > 0 {
			return headerValue
		}
	}
	return ""
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

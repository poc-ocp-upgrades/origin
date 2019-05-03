package handlers

import (
	"github.com/openshift/origin/pkg/oauthserver/api"
	"k8s.io/apiserver/pkg/authentication/user"
	"net/http"
)

type AuthenticationHandler interface {
	AuthenticationNeeded(client api.Client, w http.ResponseWriter, req *http.Request) (handled bool, err error)
}
type AuthenticationChallenger interface {
	AuthenticationChallenge(req *http.Request) (header http.Header, err error)
}
type AuthenticationRedirector interface {
	AuthenticationRedirect(w http.ResponseWriter, req *http.Request) (err error)
}
type AuthenticationRedirectors struct {
	names         []string
	redirectorMap map[string]AuthenticationRedirector
}

func (ar *AuthenticationRedirectors) Add(name string, redirector AuthenticationRedirector) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ar.redirectorMap == nil {
		ar.redirectorMap = make(map[string]AuthenticationRedirector, 1)
	}
	if _, exists := ar.redirectorMap[name]; exists {
		return
	}
	ar.names = append(ar.names, name)
	ar.redirectorMap[name] = redirector
}
func (ar *AuthenticationRedirectors) Get(name string) (AuthenticationRedirector, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	val, exists := ar.redirectorMap[name]
	return val, exists
}
func (ar *AuthenticationRedirectors) Count() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(ar.names)
}
func (ar *AuthenticationRedirectors) GetNames() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ar.names
}

type AuthenticationErrorHandler interface {
	AuthenticationError(error, http.ResponseWriter, *http.Request) (handled bool, err error)
}
type AuthenticationSelectionHandler interface {
	SelectAuthentication([]api.ProviderInfo, http.ResponseWriter, *http.Request) (selected *api.ProviderInfo, handled bool, err error)
}
type AuthenticationSuccessHandler interface {
	AuthenticationSucceeded(user user.Info, state string, w http.ResponseWriter, req *http.Request) (bool, error)
}
type GrantChecker interface {
	HasAuthorizedClient(user user.Info, grant *api.Grant) (bool, error)
}
type GrantHandler interface {
	GrantNeeded(user user.Info, grant *api.Grant, w http.ResponseWriter, req *http.Request) (granted, handled bool, err error)
}
type GrantErrorHandler interface {
	GrantError(error, http.ResponseWriter, *http.Request) (handled bool, err error)
}
type AuthenticationSuccessHandlers []AuthenticationSuccessHandler

func (all AuthenticationSuccessHandlers) AuthenticationSucceeded(user user.Info, state string, w http.ResponseWriter, req *http.Request) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, h := range all {
		if handled, err := h.AuthenticationSucceeded(user, state, w, req); handled || err != nil {
			return handled, err
		}
	}
	return false, nil
}

type AuthenticationErrorHandlers []AuthenticationErrorHandler

func (all AuthenticationErrorHandlers) AuthenticationError(err error, w http.ResponseWriter, req *http.Request) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	handled := false
	for _, h := range all {
		if handled, err = h.AuthenticationError(err, w, req); handled {
			return handled, err
		}
	}
	return handled, err
}

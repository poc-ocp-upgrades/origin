package session

import (
	"github.com/openshift/origin/pkg/oauthserver/authenticator/password/bootstrap"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
	"net/http"
	"time"
)

func NewBootstrapAuthenticator(delegate SessionAuthenticator, getter bootstrap.BootstrapUserDataGetter, store Store) SessionAuthenticator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &bootstrapAuthenticator{delegate: delegate, getter: getter, store: store}
}

type bootstrapAuthenticator struct {
	delegate SessionAuthenticator
	getter   bootstrap.BootstrapUserDataGetter
	store    Store
}

func (b *bootstrapAuthenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authResponse, ok, err := b.delegate.AuthenticateRequest(req)
	if err != nil || !ok || authResponse.User.GetName() != bootstrap.BootstrapUser {
		return authResponse, ok, err
	}
	data, ok, err := b.getter.Get()
	if err != nil || !ok {
		return nil, ok, err
	}
	if data.UID != authResponse.User.GetUID() {
		return nil, false, nil
	}
	return authResponse, true, nil
}
func (b *bootstrapAuthenticator) AuthenticationSucceeded(user user.Info, state string, w http.ResponseWriter, req *http.Request) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if user.GetName() != bootstrap.BootstrapUser {
		return b.delegate.AuthenticationSucceeded(user, state, w, req)
	}
	return false, putUser(b.store, w, user, time.Hour)
}
func (b *bootstrapAuthenticator) InvalidateAuthentication(w http.ResponseWriter, user user.Info) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if user.GetName() != bootstrap.BootstrapUser {
		return b.delegate.InvalidateAuthentication(w, user)
	}
	return nil
}

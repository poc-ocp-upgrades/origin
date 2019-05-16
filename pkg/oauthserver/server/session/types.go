package session

import (
	"github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
	"net/http"
)

type Store interface {
	Get(r *http.Request) Values
	Put(w http.ResponseWriter, v Values) error
}
type Values map[interface{}]interface{}

func (v Values) GetString(key string) (string, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	str, _ := v[key].(string)
	return str, len(str) != 0
}
func (v Values) GetInt64(key string) (int64, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	i, _ := v[key].(int64)
	return i, i != 0
}

type SessionInvalidator interface {
	InvalidateAuthentication(w http.ResponseWriter, user user.Info) error
}
type SessionAuthenticator interface {
	authenticator.Request
	handlers.AuthenticationSuccessHandler
	SessionInvalidator
}

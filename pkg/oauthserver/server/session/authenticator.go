package session

import (
	goformat "fmt"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

const (
	userNameKey = "user.name"
	userUIDKey  = "user.uid"
	expKey      = "exp"
)

type sessionAuthenticator struct {
	store  Store
	maxAge time.Duration
}

func NewAuthenticator(store Store, maxAge time.Duration) SessionAuthenticator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &sessionAuthenticator{store: store, maxAge: maxAge}
}
func (a *sessionAuthenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	values := a.store.Get(req)
	expires, ok := values.GetInt64(expKey)
	if !ok {
		return nil, false, nil
	}
	if expires < time.Now().Unix() {
		return nil, false, nil
	}
	name, ok := values.GetString(userNameKey)
	if !ok {
		return nil, false, nil
	}
	uid, ok := values.GetString(userUIDKey)
	if !ok {
		return nil, false, nil
	}
	return &authenticator.Response{User: &user.DefaultInfo{Name: name, UID: uid}}, true, nil
}
func (a *sessionAuthenticator) AuthenticationSucceeded(user user.Info, state string, w http.ResponseWriter, req *http.Request) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false, putUser(a.store, w, user, a.maxAge)
}
func (a *sessionAuthenticator) InvalidateAuthentication(w http.ResponseWriter, _ user.Info) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return putUser(a.store, w, &user.DefaultInfo{}, 0)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

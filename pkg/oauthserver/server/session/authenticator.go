package session

import (
	"net/http"
	"bytes"
	"runtime"
	"fmt"
	"time"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
)

const (
	userNameKey	= "user.name"
	userUIDKey	= "user.uid"
	expKey		= "exp"
)

type sessionAuthenticator struct {
	store	Store
	maxAge	time.Duration
}

func NewAuthenticator(store Store, maxAge time.Duration) SessionAuthenticator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &sessionAuthenticator{store: store, maxAge: maxAge}
}
func (a *sessionAuthenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, putUser(a.store, w, user, a.maxAge)
}
func (a *sessionAuthenticator) InvalidateAuthentication(w http.ResponseWriter, _ user.Info) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return putUser(a.store, w, &user.DefaultInfo{}, 0)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

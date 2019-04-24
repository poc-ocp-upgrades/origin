package denypassword

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apiserver/pkg/authentication/authenticator"
)

type denyPasswordAuthenticator struct{}

func New() authenticator.Password {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &denyPasswordAuthenticator{}
}
func (a denyPasswordAuthenticator) AuthenticatePassword(ctx context.Context, username, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, false, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

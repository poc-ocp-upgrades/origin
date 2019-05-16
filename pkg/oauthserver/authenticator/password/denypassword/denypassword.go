package denypassword

import (
	"context"
	goformat "fmt"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type denyPasswordAuthenticator struct{}

func New() authenticator.Password {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &denyPasswordAuthenticator{}
}
func (a denyPasswordAuthenticator) AuthenticatePassword(ctx context.Context, username, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, false, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

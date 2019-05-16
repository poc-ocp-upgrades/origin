package identitymapper

import (
	goformat "fmt"
	userapi "github.com/openshift/api/user/v1"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type DefaultUserInitStrategy struct{}

func NewDefaultUserInitStrategy() Initializer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &DefaultUserInitStrategy{}
}
func (*DefaultUserInitStrategy) InitializeUser(identity *userapi.Identity, user *userapi.User) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(user.FullName) == 0 {
		user.FullName = identity.Extra[authapi.IdentityDisplayNameKey]
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

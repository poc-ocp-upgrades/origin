package allowanypassword

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

package allowanypassword

import (
	"context"
	goformat "fmt"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/identitymapper"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type alwaysAcceptPasswordAuthenticator struct {
	providerName   string
	identityMapper authapi.UserIdentityMapper
}

func New(providerName string, identityMapper authapi.UserIdentityMapper) authenticator.Password {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &alwaysAcceptPasswordAuthenticator{providerName, identityMapper}
}
func (a alwaysAcceptPasswordAuthenticator) AuthenticatePassword(ctx context.Context, username, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	username = strings.TrimSpace(username)
	if username == "" || password == "" {
		return nil, false, nil
	}
	identity := authapi.NewDefaultUserIdentityInfo(a.providerName, username)
	return identitymapper.ResponseFor(a.identityMapper, identity)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

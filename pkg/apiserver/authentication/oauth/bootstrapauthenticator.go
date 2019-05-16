package oauth

import (
	"context"
	goformat "fmt"
	userapi "github.com/openshift/api/user/v1"
	oauthclient "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/password/bootstrap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	kauthenticator "k8s.io/apiserver/pkg/authentication/authenticator"
	kuser "k8s.io/apiserver/pkg/authentication/user"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type bootstrapAuthenticator struct {
	tokens    oauthclient.OAuthAccessTokenInterface
	getter    bootstrap.BootstrapUserDataGetter
	validator OAuthTokenValidator
}

func NewBootstrapAuthenticator(tokens oauthclient.OAuthAccessTokenInterface, getter bootstrap.BootstrapUserDataGetter, validators ...OAuthTokenValidator) kauthenticator.Token {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &bootstrapAuthenticator{tokens: tokens, getter: getter, validator: OAuthTokenValidators(validators)}
}
func (a *bootstrapAuthenticator) AuthenticateToken(ctx context.Context, name string) (*kauthenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	token, err := a.tokens.Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, false, errLookup
	}
	if token.UserName != bootstrap.BootstrapUser {
		return nil, false, nil
	}
	data, ok, err := a.getter.Get()
	if err != nil || !ok {
		return nil, ok, err
	}
	fakeUser := &userapi.User{ObjectMeta: metav1.ObjectMeta{UID: types.UID(data.UID)}}
	if err := a.validator.Validate(token, fakeUser); err != nil {
		return nil, false, err
	}
	return &kauthenticator.Response{User: &kuser.DefaultInfo{Name: bootstrap.BootstrapUser, Groups: []string{bootstrappolicy.ClusterAdminGroup}, Extra: map[string][]string{authorizationapi.ScopesKey: token.Scopes}}}, true, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

package google

import (
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external/openid"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	googleAuthorizeURL = "https://accounts.google.com/o/oauth2/auth"
	googleTokenURL     = "https://www.googleapis.com/oauth2/v3/token"
	googleUserInfoURL  = "https://www.googleapis.com/oauth2/v3/userinfo"
	googleHostedDomain = "hd"
)

var googleOAuthScopes = []string{"openid", "email", "profile"}

func NewProvider(providerName, clientID, clientSecret, hostedDomain string, transport http.RoundTripper) (external.Provider, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config := openid.Config{ClientID: clientID, ClientSecret: clientSecret, AuthorizeURL: googleAuthorizeURL, TokenURL: googleTokenURL, UserInfoURL: googleUserInfoURL, Scopes: googleOAuthScopes, IDClaims: []string{"sub"}, PreferredUsernameClaims: []string{"preferred_username", "email"}, EmailClaims: []string{"email"}, NameClaims: []string{"name", "email"}}
	if len(hostedDomain) > 0 {
		config.ExtraAuthorizeParameters = map[string]string{googleHostedDomain: hostedDomain}
		config.IDTokenValidator = func(idToken map[string]interface{}) error {
			hdClaim, ok := idToken[googleHostedDomain].(string)
			if !ok {
				return errors.New("id_token did not contain a hd claim")
			}
			if hdClaim != hostedDomain {
				return fmt.Errorf("id_token hd claim (%s) did not match hostedDomain (%s)", hdClaim, hostedDomain)
			}
			return nil
		}
	}
	return openid.NewProvider(providerName, transport, config)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

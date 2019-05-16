package openid

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/RangelReale/osincli"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
	"net/http"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	subjectClaim = "sub"
)

type TokenValidator func(map[string]interface{}) error
type Config struct {
	ClientID                 string
	ClientSecret             string
	Scopes                   []string
	ExtraAuthorizeParameters map[string]string
	AuthorizeURL             string
	TokenURL                 string
	UserInfoURL              string
	IDClaims                 []string
	PreferredUsernameClaims  []string
	EmailClaims              []string
	NameClaims               []string
	IDTokenValidator         TokenValidator
}
type provider struct {
	providerName string
	transport    http.RoundTripper
	Config
}

func NewProvider(providerName string, transport http.RoundTripper, config Config) (external.Provider, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(config.ClientID) == 0 {
		return nil, errors.New("ClientID is required")
	}
	if len(config.ClientSecret) == 0 {
		return nil, errors.New("ClientSecret is required")
	}
	if len(config.AuthorizeURL) == 0 {
		return nil, errors.New("authorize URL is required")
	} else if u, err := url.Parse(config.AuthorizeURL); err != nil {
		return nil, errors.New("authorize URL is invalid")
	} else if u.Scheme != "https" {
		return nil, errors.New("authorize URL must use https scheme")
	}
	if len(config.TokenURL) == 0 {
		return nil, errors.New("token URL is required")
	} else if u, err := url.Parse(config.TokenURL); err != nil {
		return nil, errors.New("token URL is invalid")
	} else if u.Scheme != "https" {
		return nil, errors.New("token URL must use https scheme")
	}
	if len(config.UserInfoURL) > 0 {
		if u, err := url.Parse(config.UserInfoURL); err != nil {
			return nil, errors.New("UserInfo URL is invalid")
		} else if u.Scheme != "https" {
			return nil, errors.New("UserInfo URL must use https scheme")
		}
	}
	if !sets.NewString(config.Scopes...).Has("openid") {
		return nil, errors.New("scopes must include openid")
	}
	if len(config.IDClaims) == 0 {
		return nil, errors.New("IDClaims must specify at least one claim")
	}
	return provider{providerName: providerName, transport: transport, Config: config}, nil
}
func (p provider) NewConfig() (*osincli.ClientConfig, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config := &osincli.ClientConfig{ClientId: p.ClientID, ClientSecret: p.ClientSecret, ErrorsInStatusCode: true, SendClientSecretInParams: true, AuthorizeUrl: p.AuthorizeURL, TokenUrl: p.TokenURL, Scope: strings.Join(p.Scopes, " ")}
	return config, nil
}
func (p provider) GetTransport() (http.RoundTripper, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return p.transport, nil
}
func (p provider) AddCustomParameters(req *osincli.AuthorizeRequest) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for k, v := range p.ExtraAuthorizeParameters {
		req.CustomParameters[k] = v
	}
}
func (p provider) GetUserIdentity(data *osincli.AccessData) (authapi.UserIdentityInfo, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	idToken, ok := getClaimValue(data.ResponseData, "id_token")
	if !ok {
		return nil, false, fmt.Errorf("no id_token returned in %#v", data.ResponseData)
	}
	idTokenClaims, err := decodeJWT(idToken)
	if err != nil {
		return nil, false, err
	}
	if p.IDTokenValidator != nil {
		if err := p.IDTokenValidator(idTokenClaims); err != nil {
			return nil, false, err
		}
	}
	idTokenSubject, ok := getClaimValue(idTokenClaims, subjectClaim)
	if !ok {
		return nil, false, fmt.Errorf("id_token did not contain a 'sub' claim: %#v", idTokenClaims)
	}
	claims := idTokenClaims
	if len(p.UserInfoURL) != 0 {
		userInfoClaims, err := fetchUserInfo(p.UserInfoURL, data.AccessToken, p.transport)
		if err != nil {
			return nil, false, err
		}
		userInfoSubject, ok := getClaimValue(userInfoClaims, subjectClaim)
		if !ok {
			return nil, false, fmt.Errorf("userinfo response did not contain a 'sub' claim: %#v", userInfoClaims)
		}
		if userInfoSubject != idTokenSubject {
			return nil, false, fmt.Errorf("userinfo 'sub' claim (%s) did not match id_token 'sub' claim (%s)", userInfoSubject, idTokenSubject)
		}
		for k, v := range userInfoClaims {
			claims[k] = v
		}
	}
	klog.V(5).Infof("openid claims: %#v", claims)
	id, ok := getClaimValue(claims, p.IDClaims...)
	if !ok {
		return nil, false, fmt.Errorf("could not retrieve id claim for %#v from %#v", p.IDClaims, claims)
	}
	identity := authapi.NewDefaultUserIdentityInfo(p.providerName, id)
	if preferredUsername, ok := getClaimValue(claims, p.PreferredUsernameClaims...); ok {
		identity.Extra[authapi.IdentityPreferredUsernameKey] = preferredUsername
	}
	if email, ok := getClaimValue(claims, p.EmailClaims...); ok {
		identity.Extra[authapi.IdentityEmailKey] = email
	}
	if name, ok := getClaimValue(claims, p.NameClaims...); ok {
		identity.Extra[authapi.IdentityDisplayNameKey] = name
	}
	klog.V(4).Infof("identity=%#v", identity)
	return identity, true, nil
}
func getClaimValue(data map[string]interface{}, claims ...string) (string, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, claim := range claims {
		s, _ := data[claim].(string)
		if len(s) > 0 {
			return s, true
		}
	}
	return "", false
}
func fetchUserInfo(url, accessToken string, transport http.RoundTripper) (map[string]interface{}, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response from UserInfo: %d, WWW-Authenticate=%s", resp.StatusCode, resp.Header.Get("WWW-Authenticate"))
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return getJSON(data)
}
func decodeJWT(jwt string) (map[string]interface{}, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	jwtParts := strings.Split(jwt, ".")
	if len(jwtParts) != 3 {
		return nil, fmt.Errorf("invalid JSON Web Token: expected 3 parts, got %d", len(jwtParts))
	}
	encodedPayload := jwtParts[1]
	if l := len(encodedPayload) % 4; l != 0 {
		encodedPayload += strings.Repeat("=", 4-l)
	}
	decodedPayload, err := base64.URLEncoding.DecodeString(encodedPayload)
	if err != nil {
		return nil, fmt.Errorf("error decoding payload: %v", err)
	}
	return getJSON(decodedPayload)
}
func getJSON(in []byte) (map[string]interface{}, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var data map[string]interface{}
	if err := json.Unmarshal(in, &data); err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	return data, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

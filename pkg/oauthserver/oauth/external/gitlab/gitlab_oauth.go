package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/RangelReale/osincli"
	"k8s.io/klog"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external"
)

const (
	gitlabUserAPIPath	= "/api/v3/user"
	gitlabOAuthScope	= "api"
)

type provider struct {
	providerName	string
	transport	http.RoundTripper
	authorizeURL	string
	tokenURL	string
	userAPIURL	string
	clientID	string
	clientSecret	string
}
type gitlabUser struct {
	ID		uint64
	Username	string
	Email		string
	Name		string
}

func NewOAuthProvider(providerName, URL, clientID, clientSecret string, transport http.RoundTripper) (external.Provider, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u, err := url.Parse(URL)
	if err != nil {
		return nil, errors.New("Host URL is invalid")
	}
	return &provider{providerName: providerName, transport: transport, authorizeURL: appendPath(*u, gitlabAuthorizePath), tokenURL: appendPath(*u, gitlabTokenPath), userAPIURL: appendPath(*u, gitlabUserAPIPath), clientID: clientID, clientSecret: clientSecret}, nil
}
func (p *provider) GetTransport() (http.RoundTripper, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.transport, nil
}
func (p *provider) NewConfig() (*osincli.ClientConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &osincli.ClientConfig{ClientId: p.clientID, ClientSecret: p.clientSecret, ErrorsInStatusCode: true, SendClientSecretInParams: true, AuthorizeUrl: p.authorizeURL, TokenUrl: p.tokenURL, Scope: gitlabOAuthScope}
	return config, nil
}
func (p *provider) AddCustomParameters(req *osincli.AuthorizeRequest) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (p *provider) GetUserIdentity(data *osincli.AccessData) (authapi.UserIdentityInfo, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	req, _ := http.NewRequest("GET", p.userAPIURL, nil)
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", data.AccessToken))
	client := http.DefaultClient
	if p.transport != nil {
		client = &http.Client{Transport: p.transport}
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, false, err
	}
	userdata := gitlabUser{}
	err = json.Unmarshal(body, &userdata)
	if err != nil {
		return nil, false, err
	}
	if userdata.ID == 0 {
		return nil, false, errors.New("Could not retrieve GitLab id")
	}
	identity := authapi.NewDefaultUserIdentityInfo(p.providerName, fmt.Sprintf("%d", userdata.ID))
	if len(userdata.Name) > 0 {
		identity.Extra[authapi.IdentityDisplayNameKey] = userdata.Name
	}
	if len(userdata.Username) > 0 {
		identity.Extra[authapi.IdentityPreferredUsernameKey] = userdata.Username
	}
	if len(userdata.Email) > 0 {
		identity.Extra[authapi.IdentityEmailKey] = userdata.Email
	}
	klog.V(4).Infof("Got identity=%#v", identity)
	return identity, true, nil
}

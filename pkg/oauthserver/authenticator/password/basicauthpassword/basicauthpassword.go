package basicauthpassword

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	goformat "fmt"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/identitymapper"
	"io/ioutil"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type Authenticator struct {
	providerName string
	url          string
	client       *http.Client
	mapper       authapi.UserIdentityMapper
}
type RemoteUserData struct {
	Subject           string `json:"sub"`
	Name              string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
	Email             string `json:"email"`
}
type RemoteError struct{ Error string }

var RedirectAttemptedError = errors.New("Redirect attempted")

func New(providerName string, url string, transport http.RoundTripper, mapper authapi.UserIdentityMapper) authenticator.Password {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if transport == nil {
		transport = http.DefaultTransport
	}
	client := &http.Client{Transport: transport}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return RedirectAttemptedError
	}
	return &Authenticator{providerName, url, client, mapper}
}
func (a *Authenticator) AuthenticatePassword(ctx context.Context, username, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, err := http.NewRequest("GET", a.url, nil)
	if err != nil {
		return nil, false, err
	}
	if strings.Contains(username, ":") {
		return nil, false, fmt.Errorf("invalid username")
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Accept", "application/json")
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, false, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}
	remoteError := RemoteError{}
	json.Unmarshal(body, &remoteError)
	if remoteError.Error != "" {
		return nil, false, errors.New(remoteError.Error)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("An error occurred while authenticating (%d)", resp.StatusCode)
	}
	remoteUserData := RemoteUserData{}
	err = json.Unmarshal(body, &remoteUserData)
	if err != nil {
		return nil, false, err
	}
	if len(remoteUserData.Subject) == 0 {
		return nil, false, errors.New("Could not retrieve user data")
	}
	identity := authapi.NewDefaultUserIdentityInfo(a.providerName, remoteUserData.Subject)
	if len(remoteUserData.Name) > 0 {
		identity.Extra[authapi.IdentityDisplayNameKey] = remoteUserData.Name
	}
	if len(remoteUserData.PreferredUsername) > 0 {
		identity.Extra[authapi.IdentityPreferredUsernameKey] = remoteUserData.PreferredUsername
	}
	if len(remoteUserData.Email) > 0 {
		identity.Extra[authapi.IdentityEmailKey] = remoteUserData.Email
	}
	return identitymapper.ResponseFor(a.mapper, identity)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

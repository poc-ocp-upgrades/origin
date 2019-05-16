package keystonepassword

import (
	"context"
	"fmt"
	goformat "fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	tokens3 "github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/identitymapper"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/klog"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"runtime/debug"
	gotime "time"
)

type keystonePasswordAuthenticator struct {
	providerName        string
	url                 string
	client              *http.Client
	domainName          string
	identityMapper      authapi.UserIdentityMapper
	useKeystoneIdentity bool
}

func New(providerName string, url string, transport http.RoundTripper, domainName string, identityMapper authapi.UserIdentityMapper, useKeystoneIdentity bool) authenticator.Password {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if transport == nil {
		transport = http.DefaultTransport
	}
	client := &http.Client{Transport: transport}
	return &keystonePasswordAuthenticator{providerName, url, client, domainName, identityMapper, useKeystoneIdentity}
}
func getUserIDv3(client *gophercloud.ProviderClient, options tokens3.AuthOptionsBuilder, eo gophercloud.EndpointOpts) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	v3Client, err := openstack.NewIdentityV3(client, eo)
	if err != nil {
		return "", err
	}
	result := tokens3.Create(v3Client, options)
	if result.Err != nil {
		return "", result.Err
	}
	user, err := result.ExtractUser()
	if err != nil {
		return "", err
	}
	return user.ID, nil
}
func (a keystonePasswordAuthenticator) AuthenticatePassword(ctx context.Context, username, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer func() {
		if e := recover(); e != nil {
			utilruntime.HandleError(fmt.Errorf("Recovered panic: %v, %s", e, debug.Stack()))
		}
	}()
	if len(password) == 0 {
		return nil, false, nil
	}
	opts := gophercloud.AuthOptions{IdentityEndpoint: a.url, Username: username, Password: password, DomainName: a.domainName}
	client, err := openstack.NewClient(opts.IdentityEndpoint)
	if err != nil {
		klog.Warningf("Failed: Initializing openstack authentication client: %v", err)
		return nil, false, err
	}
	client.HTTPClient = *a.client
	userid, err := getUserIDv3(client, &opts, gophercloud.EndpointOpts{})
	if err != nil {
		if _, ok := err.(gophercloud.ErrDefault401); ok {
			return nil, false, nil
		}
		klog.Warningf("Failed: Calling openstack AuthenticateV3: %v", err)
		return nil, false, err
	}
	providerUserID := username
	if a.useKeystoneIdentity {
		providerUserID = userid
	}
	identity := authapi.NewDefaultUserIdentityInfo(a.providerName, providerUserID)
	identity.Extra[authapi.IdentityPreferredUsernameKey] = username
	return identitymapper.ResponseFor(a.identityMapper, identity)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

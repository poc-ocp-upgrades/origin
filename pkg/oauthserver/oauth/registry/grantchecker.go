package registry

import (
	stderrors "errors"
	goformat "fmt"
	oauth "github.com/openshift/api/oauth/v1"
	oauthclient "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	"github.com/openshift/origin/pkg/oauth/scope"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kuser "k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var errEmptyUID = stderrors.New("user from request has empty UID and thus cannot perform a grant flow")

type ClientAuthorizationGrantChecker struct {
	client oauthclient.OAuthClientAuthorizationInterface
}

func NewClientAuthorizationGrantChecker(client oauthclient.OAuthClientAuthorizationInterface) *ClientAuthorizationGrantChecker {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ClientAuthorizationGrantChecker{client}
}
func (c *ClientAuthorizationGrantChecker) HasAuthorizedClient(user kuser.Info, grant *api.Grant) (approved bool, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(user.GetUID()) == 0 {
		return false, errEmptyUID
	}
	id := user.GetName() + ":" + grant.Client.GetId()
	var authorization *oauth.OAuthClientAuthorization
	if err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		authorization, err = c.getClientAuthorization(id, user)
		return err
	}); err != nil {
		return false, err
	}
	if authorization == nil || !scope.Covers(authorization.Scopes, scope.Split(grant.Scope)) {
		return false, nil
	}
	return true, nil
}
func (c *ClientAuthorizationGrantChecker) getClientAuthorization(name string, user kuser.Info) (*oauth.OAuthClientAuthorization, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authorization, err := c.client.Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if user.GetUID() != authorization.UserUID {
		klog.Infof("%#v does not match stored client authorization %#v, attempting to delete stale authorization", user, authorization)
		if err := c.client.Delete(name, metav1.NewPreconditionDeleteOptions(string(authorization.UID))); err != nil && !errors.IsNotFound(err) {
			return nil, err
		}
		return nil, nil
	}
	return authorization, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

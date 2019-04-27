package legacy

import (
	"testing"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/openshift/origin/pkg/api/apihelpers/apitesting"
	oauthapi "github.com/openshift/origin/pkg/oauth/apis/oauth"
)

func TestOAuthFieldSelectorConversions(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	install := func(scheme *runtime.Scheme) error {
		InstallInternalLegacyOAuth(scheme)
		return nil
	}
	apitesting.FieldKeyCheck{SchemeBuilder: []func(*runtime.Scheme) error{install}, Kind: GroupVersion.WithKind("OAuthAccessToken"), AllowedExternalFieldKeys: []string{"clientName", "userName", "userUID", "authorizeToken"}, FieldKeyEvaluatorFn: oauthapi.OAuthAccessTokenFieldSelector}.Check(t)
	apitesting.FieldKeyCheck{SchemeBuilder: []func(*runtime.Scheme) error{install}, Kind: GroupVersion.WithKind("OAuthAuthorizeToken"), AllowedExternalFieldKeys: []string{"clientName", "userName", "userUID"}, FieldKeyEvaluatorFn: oauthapi.OAuthAuthorizeTokenFieldSelector}.Check(t)
	apitesting.FieldKeyCheck{SchemeBuilder: []func(*runtime.Scheme) error{install}, Kind: GroupVersion.WithKind("OAuthClientAuthorization"), AllowedExternalFieldKeys: []string{"clientName", "userName", "userUID"}, FieldKeyEvaluatorFn: oauthapi.OAuthClientAuthorizationFieldSelector}.Check(t)
}

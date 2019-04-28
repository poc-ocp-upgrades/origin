package v1

import (
	"testing"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "github.com/openshift/api/oauth/v1"
	"github.com/openshift/origin/pkg/api/apihelpers/apitesting"
	oauthapi "github.com/openshift/origin/pkg/oauth/apis/oauth"
)

func TestFieldSelectorConversions(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	apitesting.FieldKeyCheck{SchemeBuilder: []func(*runtime.Scheme) error{Install}, Kind: v1.GroupVersion.WithKind("OAuthAccessToken"), AllowedExternalFieldKeys: []string{"clientName", "userName", "userUID", "authorizeToken"}, FieldKeyEvaluatorFn: oauthapi.OAuthAccessTokenFieldSelector}.Check(t)
	apitesting.FieldKeyCheck{SchemeBuilder: []func(*runtime.Scheme) error{Install}, Kind: v1.GroupVersion.WithKind("OAuthAuthorizeToken"), AllowedExternalFieldKeys: []string{"clientName", "userName", "userUID"}, FieldKeyEvaluatorFn: oauthapi.OAuthAuthorizeTokenFieldSelector}.Check(t)
	apitesting.FieldKeyCheck{SchemeBuilder: []func(*runtime.Scheme) error{Install}, Kind: v1.GroupVersion.WithKind("OAuthClientAuthorization"), AllowedExternalFieldKeys: []string{"clientName", "userName", "userUID"}, FieldKeyEvaluatorFn: oauthapi.OAuthClientAuthorizationFieldSelector}.Check(t)
}

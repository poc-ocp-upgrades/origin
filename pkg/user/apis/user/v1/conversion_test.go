package v1

import (
	"testing"
	"k8s.io/apimachinery/pkg/runtime"
	v1 "github.com/openshift/api/user/v1"
	"github.com/openshift/origin/pkg/api/apihelpers/apitesting"
	userapi "github.com/openshift/origin/pkg/user/apis/user"
)

func TestFieldSelectorConversions(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	apitesting.FieldKeyCheck{SchemeBuilder: []func(*runtime.Scheme) error{Install}, Kind: v1.GroupVersion.WithKind("Identity"), AllowedExternalFieldKeys: []string{"providerName", "providerUserName", "user.name", "user.uid"}, FieldKeyEvaluatorFn: userapi.IdentityFieldSelector}.Check(t)
}

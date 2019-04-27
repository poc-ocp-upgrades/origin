package legacy

import (
	"testing"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/openshift/origin/pkg/api/apihelpers/apitesting"
	userapi "github.com/openshift/origin/pkg/user/apis/user"
)

func TestUserFieldSelectorConversions(t *testing.T) {
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
		InstallInternalLegacyUser(scheme)
		return nil
	}
	apitesting.FieldKeyCheck{SchemeBuilder: []func(*runtime.Scheme) error{install}, Kind: GroupVersion.WithKind("Identity"), AllowedExternalFieldKeys: []string{"providerName", "providerUserName", "user.name", "user.uid"}, FieldKeyEvaluatorFn: userapi.IdentityFieldSelector}.Check(t)
}

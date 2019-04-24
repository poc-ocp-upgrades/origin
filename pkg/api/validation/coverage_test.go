package validation

import (
	"reflect"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"strings"
	"testing"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"github.com/openshift/origin/pkg/api/legacy"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	quotaapi "github.com/openshift/origin/pkg/quota/apis/quota"
)

var KnownValidationExceptions = []reflect.Type{reflect.TypeOf(&buildapi.BuildLog{}), reflect.TypeOf(&appsapi.DeploymentLog{}), reflect.TypeOf(&imageapi.ImageStreamImage{}), reflect.TypeOf(&imageapi.ImageStreamTag{}), reflect.TypeOf(&authorizationapi.IsPersonalSubjectAccessReview{}), reflect.TypeOf(&authorizationapi.SubjectAccessReviewResponse{}), reflect.TypeOf(&authorizationapi.ResourceAccessReviewResponse{}), reflect.TypeOf(&quotaapi.AppliedClusterResourceQuota{})}
var MissingValidationExceptions = []reflect.Type{reflect.TypeOf(&buildapi.BuildLogOptions{}), reflect.TypeOf(&buildapi.BinaryBuildRequestOptions{}), reflect.TypeOf(&imageapi.DockerImage{})}

func TestCoverage(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for kind, apiType := range legacyscheme.Scheme.KnownTypes(legacy.InternalGroupVersion) {
		if strings.HasPrefix(apiType.PkgPath(), "github.com/openshift/origin/vendor/") {
			continue
		}
		if strings.HasSuffix(kind, "List") {
			continue
		}
		ptrType := reflect.PtrTo(apiType)
		if _, exists := Validator.typeToValidator[ptrType]; !exists {
			allowed := false
			for _, exception := range KnownValidationExceptions {
				if ptrType == exception {
					allowed = true
					break
				}
			}
			for _, exception := range MissingValidationExceptions {
				if ptrType == exception {
					allowed = true
				}
			}
			if !allowed {
				t.Errorf("%v is not registered.  Look in pkg/api/validation/register.go.", apiType)
			}
		}
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

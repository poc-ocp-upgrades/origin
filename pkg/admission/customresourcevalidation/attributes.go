package customresourcevalidation

import (
	goformat "fmt"
	authorizationv1 "github.com/openshift/api/authorization/v1"
	configv1 "github.com/openshift/api/config/v1"
	quotav1 "github.com/openshift/api/quota/v1"
	securityv1 "github.com/openshift/api/security/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/admission"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type unstructuredUnpackingAttributes struct{ admission.Attributes }

func (a *unstructuredUnpackingAttributes) GetObject() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return toBestObjectPossible(a.Attributes.GetObject())
}
func (a *unstructuredUnpackingAttributes) GetOldObject() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return toBestObjectPossible(a.Attributes.GetOldObject())
}
func toBestObjectPossible(orig runtime.Object) runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	unstructuredOrig, ok := orig.(runtime.Unstructured)
	if !ok {
		return orig
	}
	targetObj, err := supportedObjectsScheme.New(unstructuredOrig.GetObjectKind().GroupVersionKind())
	if err != nil {
		utilruntime.HandleError(err)
		return unstructuredOrig
	}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredOrig.UnstructuredContent(), targetObj); err != nil {
		utilruntime.HandleError(err)
		return unstructuredOrig
	}
	return targetObj
}

var supportedObjectsScheme = runtime.NewScheme()

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	utilruntime.Must(configv1.Install(supportedObjectsScheme))
	utilruntime.Must(quotav1.Install(supportedObjectsScheme))
	utilruntime.Must(securityv1.Install(supportedObjectsScheme))
	utilruntime.Must(authorizationv1.Install(supportedObjectsScheme))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

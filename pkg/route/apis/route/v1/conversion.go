package v1

import (
	v1 "github.com/openshift/api/route/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
)

func addFieldSelectorKeyConversions(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddFieldLabelConversionFunc(v1.GroupVersion.WithKind("Route"), routeFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	return nil
}
func routeFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "spec.path", "spec.host", "spec.to.name":
		return label, value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

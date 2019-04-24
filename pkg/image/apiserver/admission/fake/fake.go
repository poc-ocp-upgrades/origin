package fake

import (
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
)

type ImageStreamLimitVerifier struct {
	ImageStreamEvaluator	func(ns string, is *imageapi.ImageStream) error
	Err			error
}

func (f *ImageStreamLimitVerifier) VerifyLimits(ns string, is *imageapi.ImageStream) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if f.ImageStreamEvaluator != nil {
		return f.ImageStreamEvaluator(ns, is)
	}
	return f.Err
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

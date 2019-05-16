package fake

import (
	goformat "fmt"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type ImageStreamLimitVerifier struct {
	ImageStreamEvaluator func(ns string, is *imageapi.ImageStream) error
	Err                  error
}

func (f *ImageStreamLimitVerifier) VerifyLimits(ns string, is *imageapi.ImageStream) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if f.ImageStreamEvaluator != nil {
		return f.ImageStreamEvaluator(ns, is)
	}
	return f.Err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

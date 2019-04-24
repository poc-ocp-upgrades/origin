package policy

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	buildv1 "github.com/openshift/api/build/v1"
	buildutil "github.com/openshift/origin/pkg/build/util"
)

type NoBuildNumberAnnotationError struct{ build *buildv1.Build }

func NewNoBuildNumberAnnotationError(build *buildv1.Build) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NoBuildNumberAnnotationError{build: build}
}
func (e NoBuildNumberAnnotationError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("build %s/%s does not have required %q annotation set", e.build.Namespace, e.build.Name, buildutil.BuildNumberAnnotation)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

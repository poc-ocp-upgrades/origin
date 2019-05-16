package policy

import (
	"fmt"
	goformat "fmt"
	buildv1 "github.com/openshift/api/build/v1"
	buildutil "github.com/openshift/origin/pkg/build/util"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type NoBuildNumberAnnotationError struct{ build *buildv1.Build }

func NewNoBuildNumberAnnotationError(build *buildv1.Build) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NoBuildNumberAnnotationError{build: build}
}
func (e NoBuildNumberAnnotationError) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("build %s/%s does not have required %q annotation set", e.build.Namespace, e.build.Name, buildutil.BuildNumberAnnotation)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

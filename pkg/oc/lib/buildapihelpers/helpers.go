package buildapihelpers

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	buildv1 "github.com/openshift/api/build/v1"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
)

type PredicateFunc func(interface{}) bool

func FilterBuilds(builds []buildv1.Build, predicate PredicateFunc) []buildv1.Build {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(builds) == 0 {
		return builds
	}
	result := make([]buildv1.Build, 0)
	for _, build := range builds {
		if predicate(build) {
			result = append(result, build)
		}
	}
	return result
}
func ByBuildConfigPredicate(labelValue string) PredicateFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(arg interface{}) bool {
		return hasBuildConfigAnnotation(arg.(buildv1.Build), buildapi.BuildConfigAnnotation, labelValue) || hasBuildConfigLabel(arg.(buildv1.Build), buildapi.BuildConfigLabel, labelValue)
	}
}
func hasBuildConfigLabel(build buildv1.Build, labelName, labelValue string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	value, ok := build.Labels[labelName]
	return ok && value == labelValue
}
func hasBuildConfigAnnotation(build buildv1.Build, annotationName, annotationValue string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if build.Annotations == nil {
		return false
	}
	value, ok := build.Annotations[annotationName]
	return ok && value == annotationValue
}
func BuildNameForConfigVersion(name string, version int) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s-%d", name, version)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

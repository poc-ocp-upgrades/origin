package release

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

const (
	annotationReleaseFromRelease		= "release.openshift.io/from-release"
	annotationReleaseFromImageStream	= "release.openshift.io/from-image-stream"
	annotationReleaseOperator		= "io.openshift.release.operator"
	annotationReleaseOverride		= "io.openshift.release.override"
	annotationReleaseBaseImageDigest	= "io.openshift.release.base-image-digest"
	annotationBuildVersions			= "io.openshift.build.versions"
	annotationBuildSourceRef		= "io.openshift.build.commit.ref"
	annotationBuildSourceCommit		= "io.openshift.build.commit.id"
	annotationBuildSourceLocation		= "io.openshift.build.source-location"
	urlGithubPrefix				= "https://github.com/"
)

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

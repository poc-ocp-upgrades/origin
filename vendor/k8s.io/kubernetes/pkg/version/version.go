package version

import (
	"fmt"
	apimachineryversion "k8s.io/apimachinery/pkg/version"
	"runtime"
)

func Get() apimachineryversion.Info {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apimachineryversion.Info{Major: gitMajor, Minor: gitMinor, GitVersion: gitVersion, GitCommit: gitCommit, GitTreeState: gitTreeState, BuildDate: buildDate, GoVersion: runtime.Version(), Compiler: runtime.Compiler, Platform: fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)}
}

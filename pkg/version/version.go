package version

import (
	"fmt"
	goformat "fmt"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/version"
	goos "os"
	"runtime"
	godefaultruntime "runtime"
	gotime "time"
)

var (
	commitFromGit  string
	versionFromGit = "unknown"
	majorFromGit   string
	minorFromGit   string
	buildDate      string
	gitTreeState   string
)

func Get() version.Info {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return version.Info{Major: majorFromGit, Minor: minorFromGit, GitCommit: commitFromGit, GitVersion: versionFromGit, GitTreeState: gitTreeState, BuildDate: buildDate, GoVersion: runtime.Version(), Compiler: runtime.Compiler, Platform: fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)}
}
func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	buildInfo := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "openshift_build_info", Help: "A metric with a constant '1' value labeled by major, minor, git commit & git version from which OpenShift was built."}, []string{"major", "minor", "gitCommit", "gitVersion"})
	buildInfo.WithLabelValues(majorFromGit, minorFromGit, commitFromGit, versionFromGit).Set(1)
	prometheus.MustRegister(buildInfo)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

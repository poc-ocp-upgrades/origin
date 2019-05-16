package exec

import (
	goformat "fmt"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/probe"
	"k8s.io/utils/exec"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func New() Prober {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return execProber{}
}

type Prober interface {
	Probe(e exec.Cmd) (probe.Result, string, error)
}
type execProber struct{}

func (pr execProber) Probe(e exec.Cmd) (probe.Result, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, err := e.CombinedOutput()
	klog.V(4).Infof("Exec probe response: %q", string(data))
	if err != nil {
		exit, ok := err.(exec.ExitError)
		if ok {
			if exit.ExitStatus() == 0 {
				return probe.Success, string(data), nil
			}
			return probe.Failure, string(data), nil
		}
		return probe.Unknown, "", err
	}
	return probe.Success, string(data), nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

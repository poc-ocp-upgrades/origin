package tcp

import (
	goformat "fmt"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/probe"
	"net"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	"time"
	gotime "time"
)

func New() Prober {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return tcpProber{}
}

type Prober interface {
	Probe(host string, port int, timeout time.Duration) (probe.Result, string, error)
}
type tcpProber struct{}

func (pr tcpProber) Probe(host string, port int, timeout time.Duration) (probe.Result, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return DoTCPProbe(net.JoinHostPort(host, strconv.Itoa(port)), timeout)
}
func DoTCPProbe(addr string, timeout time.Duration) (probe.Result, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return probe.Failure, err.Error(), nil
	}
	err = conn.Close()
	if err != nil {
		klog.Errorf("Unexpected error closing TCP probe socket: %v (%#v)", err, err)
	}
	return probe.Success, "", nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

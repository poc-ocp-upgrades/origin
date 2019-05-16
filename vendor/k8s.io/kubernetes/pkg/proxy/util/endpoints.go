package util

import (
	"fmt"
	goformat "fmt"
	"k8s.io/klog"
	"net"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	gotime "time"
)

func IPPart(s string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	len := 32
	if ip.To4() == nil {
		len = 128
	}
	return fmt.Sprintf("%s/%d", ip.String(), len)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

package winuserspace

import (
	goformat "fmt"
	"k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/proxy"
	"net"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type LoadBalancer interface {
	NextEndpoint(service proxy.ServicePortName, srcAddr net.Addr, sessionAffinityReset bool) (string, error)
	NewService(service proxy.ServicePortName, sessionAffinityType v1.ServiceAffinity, stickyMaxAgeMinutes int) error
	DeleteService(service proxy.ServicePortName)
	CleanupStaleStickySessions(service proxy.ServicePortName)
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

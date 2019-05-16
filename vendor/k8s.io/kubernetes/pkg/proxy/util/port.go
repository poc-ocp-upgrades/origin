package util

import (
	"fmt"
	"k8s.io/klog"
	"net"
	"strconv"
)

type LocalPort struct {
	Description string
	IP          string
	Port        int
	Protocol    string
}

func (lp *LocalPort) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ipPort := net.JoinHostPort(lp.IP, strconv.Itoa(lp.Port))
	return fmt.Sprintf("%q (%s/%s)", lp.Description, ipPort, lp.Protocol)
}

type Closeable interface{ Close() error }
type PortOpener interface {
	OpenLocalPort(lp *LocalPort) (Closeable, error)
}

func RevertPorts(replacementPortsMap, originalPortsMap map[LocalPort]Closeable) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for k, v := range replacementPortsMap {
		if originalPortsMap[k] == nil {
			klog.V(2).Infof("Closing local port %s", k.String())
			v.Close()
		}
	}
}

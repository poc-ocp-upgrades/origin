package app

import (
	"github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/windows/service"
)

const (
	serviceName = "kube-proxy"
)

func initForOS(windowsService bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if windowsService {
		return service.InitService(serviceName)
	}
	return nil
}
func (o *Options) addOSFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.BoolVar(&o.WindowsService, "windows-service", o.WindowsService, "Enable Windows Service Control Manager API integration")
}

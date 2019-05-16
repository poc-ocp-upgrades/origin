package cloudprovider

import (
	goformat "fmt"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers/aws"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers/azure"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers/cloudstack"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers/gce"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers/openstack"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers/ovirt"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers/photon"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

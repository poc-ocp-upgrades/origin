package clusterresourceoverride

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "autoscaling.openshift.io/ClusterResourceOverride"
const ConfigKind = "ClusterResourceOverrideConfig"

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

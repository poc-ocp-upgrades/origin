package persistentvolume

import (
	goformat "fmt"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/features"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func DropDisabledFields(pvSpec *api.PersistentVolumeSpec, oldPVSpec *api.PersistentVolumeSpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
		pvSpec.VolumeMode = nil
		if oldPVSpec != nil {
			oldPVSpec.VolumeMode = nil
		}
	}
	if !utilfeature.DefaultFeatureGate.Enabled(features.CSIPersistentVolume) {
		if oldPVSpec == nil || oldPVSpec.PersistentVolumeSource.CSI == nil {
			pvSpec.PersistentVolumeSource.CSI = nil
		}
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

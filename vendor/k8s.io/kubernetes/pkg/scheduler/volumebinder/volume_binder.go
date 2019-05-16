package volumebinder

import (
	goformat "fmt"
	"k8s.io/api/core/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	storageinformers "k8s.io/client-go/informers/storage/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/controller/volume/persistentvolume"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type VolumeBinder struct {
	Binder persistentvolume.SchedulerVolumeBinder
}

func NewVolumeBinder(client clientset.Interface, pvcInformer coreinformers.PersistentVolumeClaimInformer, pvInformer coreinformers.PersistentVolumeInformer, storageClassInformer storageinformers.StorageClassInformer, bindTimeout time.Duration) *VolumeBinder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &VolumeBinder{Binder: persistentvolume.NewVolumeBinder(client, pvcInformer, pvInformer, storageClassInformer, bindTimeout)}
}
func NewFakeVolumeBinder(config *persistentvolume.FakeVolumeBinderConfig) *VolumeBinder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &VolumeBinder{Binder: persistentvolume.NewFakeVolumeBinder(config)}
}
func (b *VolumeBinder) DeletePodBindings(pod *v1.Pod) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cache := b.Binder.GetBindingsCache()
	if cache != nil && pod != nil {
		cache.DeleteBindings(pod)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

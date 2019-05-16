package expand

import (
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller/volume/events"
	"k8s.io/kubernetes/pkg/controller/volume/expand/cache"
	"k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util"
	"time"
)

type PVCPopulator interface{ Run(stopCh <-chan struct{}) }
type pvcPopulator struct {
	loopPeriod      time.Duration
	resizeMap       cache.VolumeResizeMap
	pvcLister       corelisters.PersistentVolumeClaimLister
	pvLister        corelisters.PersistentVolumeLister
	kubeClient      clientset.Interface
	volumePluginMgr *volume.VolumePluginMgr
	recorder        record.EventRecorder
}

func NewPVCPopulator(loopPeriod time.Duration, resizeMap cache.VolumeResizeMap, pvcLister corelisters.PersistentVolumeClaimLister, pvLister corelisters.PersistentVolumeLister, volumePluginMgr *volume.VolumePluginMgr, kubeClient clientset.Interface) PVCPopulator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	populator := &pvcPopulator{loopPeriod: loopPeriod, pvcLister: pvcLister, pvLister: pvLister, resizeMap: resizeMap, volumePluginMgr: volumePluginMgr, kubeClient: kubeClient}
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	populator.recorder = eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "volume_expand"})
	return populator
}
func (populator *pvcPopulator) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	wait.Until(populator.Sync, populator.loopPeriod, stopCh)
}
func (populator *pvcPopulator) Sync() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pvcs, err := populator.pvcLister.List(labels.Everything())
	if err != nil {
		klog.Errorf("Listing PVCs failed in populator : %v", err)
		return
	}
	for _, pvc := range pvcs {
		pv, err := getPersistentVolume(pvc, populator.pvLister)
		if err != nil {
			klog.V(5).Infof("Error getting persistent volume for PVC %q : %v", pvc.UID, err)
			continue
		}
		pvcSize := pvc.Spec.Resources.Requests[v1.ResourceStorage]
		pvcStatusSize := pvc.Status.Capacity[v1.ResourceStorage]
		volumeSpec := volume.NewSpecFromPersistentVolume(pv, false)
		volumePlugin, err := populator.volumePluginMgr.FindExpandablePluginBySpec(volumeSpec)
		if (err != nil || volumePlugin == nil) && pvcStatusSize.Cmp(pvcSize) < 0 {
			err = fmt.Errorf("didn't find a plugin capable of expanding the volume; " + "waiting for an external controller to process this PVC")
			eventType := v1.EventTypeNormal
			if err != nil {
				eventType = v1.EventTypeWarning
			}
			populator.recorder.Event(pvc, eventType, events.ExternalExpanding, fmt.Sprintf("Ignoring the PVC: %v.", err))
			klog.V(3).Infof("Ignoring the PVC %q (uid: %q) : %v.", util.GetPersistentVolumeClaimQualifiedName(pvc), pvc.UID, err)
			continue
		}
		populator.resizeMap.AddPVCUpdate(pvc, pv)
	}
}

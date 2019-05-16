package populator

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	corelisters "k8s.io/client-go/listers/core/v1"
	kcache "k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/cache"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/util"
	"k8s.io/kubernetes/pkg/volume"
	volutil "k8s.io/kubernetes/pkg/volume/util"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type DesiredStateOfWorldPopulator interface{ Run(stopCh <-chan struct{}) }

func NewDesiredStateOfWorldPopulator(loopSleepDuration time.Duration, listPodsRetryDuration time.Duration, podLister corelisters.PodLister, desiredStateOfWorld cache.DesiredStateOfWorld, volumePluginMgr *volume.VolumePluginMgr, pvcLister corelisters.PersistentVolumeClaimLister, pvLister corelisters.PersistentVolumeLister) DesiredStateOfWorldPopulator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &desiredStateOfWorldPopulator{loopSleepDuration: loopSleepDuration, listPodsRetryDuration: listPodsRetryDuration, podLister: podLister, desiredStateOfWorld: desiredStateOfWorld, volumePluginMgr: volumePluginMgr, pvcLister: pvcLister, pvLister: pvLister}
}

type desiredStateOfWorldPopulator struct {
	loopSleepDuration     time.Duration
	podLister             corelisters.PodLister
	desiredStateOfWorld   cache.DesiredStateOfWorld
	volumePluginMgr       *volume.VolumePluginMgr
	pvcLister             corelisters.PersistentVolumeClaimLister
	pvLister              corelisters.PersistentVolumeLister
	listPodsRetryDuration time.Duration
	timeOfLastListPods    time.Time
}

func (dswp *desiredStateOfWorldPopulator) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	wait.Until(dswp.populatorLoopFunc(), dswp.loopSleepDuration, stopCh)
}
func (dswp *desiredStateOfWorldPopulator) populatorLoopFunc() func() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func() {
		dswp.findAndRemoveDeletedPods()
		if time.Since(dswp.timeOfLastListPods) < dswp.listPodsRetryDuration {
			klog.V(5).Infof("Skipping findAndAddActivePods(). Not permitted until %v (listPodsRetryDuration %v).", dswp.timeOfLastListPods.Add(dswp.listPodsRetryDuration), dswp.listPodsRetryDuration)
			return
		}
		dswp.findAndAddActivePods()
	}
}
func (dswp *desiredStateOfWorldPopulator) findAndRemoveDeletedPods() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for dswPodUID, dswPodToAdd := range dswp.desiredStateOfWorld.GetPodToAdd() {
		dswPodKey, err := kcache.MetaNamespaceKeyFunc(dswPodToAdd.Pod)
		if err != nil {
			klog.Errorf("MetaNamespaceKeyFunc failed for pod %q (UID %q) with: %v", dswPodKey, dswPodUID, err)
			continue
		}
		namespace, name, err := kcache.SplitMetaNamespaceKey(dswPodKey)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("error splitting dswPodKey %q: %v", dswPodKey, err))
			continue
		}
		informerPod, err := dswp.podLister.Pods(namespace).Get(name)
		switch {
		case errors.IsNotFound(err):
		case err != nil:
			klog.Errorf("podLister Get failed for pod %q (UID %q) with %v", dswPodKey, dswPodUID, err)
			continue
		default:
			volumeActionFlag := util.DetermineVolumeAction(informerPod, dswp.desiredStateOfWorld, true)
			if volumeActionFlag {
				informerPodUID := volutil.GetUniquePodName(informerPod)
				if informerPodUID == dswPodUID {
					klog.V(10).Infof("Verified pod %q (UID %q) from dsw exists in pod informer.", dswPodKey, dswPodUID)
					continue
				}
			}
		}
		klog.V(1).Infof("Removing pod %q (UID %q) from dsw because it does not exist in pod informer.", dswPodKey, dswPodUID)
		dswp.desiredStateOfWorld.DeletePod(dswPodUID, dswPodToAdd.VolumeName, dswPodToAdd.NodeName)
	}
}
func (dswp *desiredStateOfWorldPopulator) findAndAddActivePods() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pods, err := dswp.podLister.List(labels.Everything())
	if err != nil {
		klog.Errorf("podLister List failed: %v", err)
		return
	}
	dswp.timeOfLastListPods = time.Now()
	for _, pod := range pods {
		if volutil.IsPodTerminated(pod, pod.Status) {
			continue
		}
		util.ProcessPodVolumes(pod, true, dswp.desiredStateOfWorld, dswp.volumePluginMgr, dswp.pvcLister, dswp.pvLister)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

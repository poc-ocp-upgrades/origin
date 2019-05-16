package expand

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller/volume/expand/cache"
	"k8s.io/kubernetes/pkg/util/goroutinemap/exponentialbackoff"
	"k8s.io/kubernetes/pkg/volume/util"
	"k8s.io/kubernetes/pkg/volume/util/operationexecutor"
	"time"
)

type SyncVolumeResize interface{ Run(stopCh <-chan struct{}) }
type syncResize struct {
	loopPeriod  time.Duration
	resizeMap   cache.VolumeResizeMap
	opsExecutor operationexecutor.OperationExecutor
	kubeClient  clientset.Interface
}

func NewSyncVolumeResize(loopPeriod time.Duration, opsExecutor operationexecutor.OperationExecutor, resizeMap cache.VolumeResizeMap, kubeClient clientset.Interface) SyncVolumeResize {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rc := &syncResize{loopPeriod: loopPeriod, opsExecutor: opsExecutor, resizeMap: resizeMap, kubeClient: kubeClient}
	return rc
}
func (rc *syncResize) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	wait.Until(rc.Sync, rc.loopPeriod, stopCh)
}
func (rc *syncResize) Sync() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, pvcWithResizeRequest := range rc.resizeMap.GetPVCsWithResizeRequest() {
		uniqueVolumeKey := v1.UniqueVolumeName(pvcWithResizeRequest.UniquePVCKey())
		updatedClaim, err := markPVCResizeInProgress(pvcWithResizeRequest, rc.kubeClient)
		if err != nil {
			klog.V(5).Infof("Error setting PVC %s in progress with error : %v", pvcWithResizeRequest.QualifiedName(), err)
			continue
		}
		if updatedClaim != nil {
			pvcWithResizeRequest.PVC = updatedClaim
		}
		if rc.opsExecutor.IsOperationPending(uniqueVolumeKey, "") {
			klog.V(10).Infof("Operation for PVC %v is already pending", pvcWithResizeRequest.QualifiedName())
			continue
		}
		growFuncError := rc.opsExecutor.ExpandVolume(pvcWithResizeRequest, rc.resizeMap)
		if growFuncError != nil && !exponentialbackoff.IsExponentialBackoff(growFuncError) {
			klog.Errorf("Error growing pvc %s with %v", pvcWithResizeRequest.QualifiedName(), growFuncError)
		}
		if growFuncError == nil {
			klog.V(5).Infof("Started opsExecutor.ExpandVolume for volume %s", pvcWithResizeRequest.QualifiedName())
		}
	}
}
func markPVCResizeInProgress(pvcWithResizeRequest *cache.PVCWithResizeRequest, kubeClient clientset.Interface) (*v1.PersistentVolumeClaim, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	progressCondition := v1.PersistentVolumeClaimCondition{Type: v1.PersistentVolumeClaimResizing, Status: v1.ConditionTrue, LastTransitionTime: metav1.Now()}
	conditions := []v1.PersistentVolumeClaimCondition{progressCondition}
	newPVC := pvcWithResizeRequest.PVC.DeepCopy()
	newPVC = util.MergeResizeConditionOnPVC(newPVC, conditions)
	return util.PatchPVCStatus(pvcWithResizeRequest.PVC, newPVC, kubeClient)
}

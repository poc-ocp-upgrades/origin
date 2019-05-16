package cache

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	commontypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/volume/util"
	"k8s.io/kubernetes/pkg/volume/util/types"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

type VolumeResizeMap interface {
	AddPVCUpdate(pvc *v1.PersistentVolumeClaim, pv *v1.PersistentVolume)
	DeletePVC(pvc *v1.PersistentVolumeClaim)
	GetPVCsWithResizeRequest() []*PVCWithResizeRequest
	MarkAsResized(*PVCWithResizeRequest, resource.Quantity) error
	UpdatePVSize(*PVCWithResizeRequest, resource.Quantity) error
	MarkForFSResize(*PVCWithResizeRequest) error
}
type volumeResizeMap struct {
	pvcrs      map[types.UniquePVCName]*PVCWithResizeRequest
	kubeClient clientset.Interface
	sync.Mutex
}
type PVCWithResizeRequest struct {
	PVC              *v1.PersistentVolumeClaim
	PersistentVolume *v1.PersistentVolume
	CurrentSize      resource.Quantity
	ExpectedSize     resource.Quantity
}

func (pvcr *PVCWithResizeRequest) UniquePVCKey() types.UniquePVCName {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return types.UniquePVCName(pvcr.PVC.UID)
}
func (pvcr *PVCWithResizeRequest) QualifiedName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return util.GetPersistentVolumeClaimQualifiedName(pvcr.PVC)
}
func NewVolumeResizeMap(kubeClient clientset.Interface) VolumeResizeMap {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resizeMap := &volumeResizeMap{}
	resizeMap.pvcrs = make(map[types.UniquePVCName]*PVCWithResizeRequest)
	resizeMap.kubeClient = kubeClient
	return resizeMap
}
func (resizeMap *volumeResizeMap) AddPVCUpdate(pvc *v1.PersistentVolumeClaim, pv *v1.PersistentVolume) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pv.Spec.ClaimRef == nil || pvc.Namespace != pv.Spec.ClaimRef.Namespace || pvc.Name != pv.Spec.ClaimRef.Name {
		klog.V(4).Infof("Persistent Volume is not bound to PVC being updated : %s", util.ClaimToClaimKey(pvc))
		return
	}
	if pvc.Status.Phase != v1.ClaimBound {
		return
	}
	pvcSize := pvc.Spec.Resources.Requests[v1.ResourceStorage]
	pvcStatusSize := pvc.Status.Capacity[v1.ResourceStorage]
	if pvcStatusSize.Cmp(pvcSize) >= 0 {
		return
	}
	klog.V(4).Infof("Adding pvc %s with Size %s/%s for resizing", util.ClaimToClaimKey(pvc), pvcSize.String(), pvcStatusSize.String())
	pvcRequest := &PVCWithResizeRequest{PVC: pvc, CurrentSize: pvcStatusSize, ExpectedSize: pvcSize, PersistentVolume: pv}
	resizeMap.Lock()
	defer resizeMap.Unlock()
	resizeMap.pvcrs[types.UniquePVCName(pvc.UID)] = pvcRequest
}
func (resizeMap *volumeResizeMap) GetPVCsWithResizeRequest() []*PVCWithResizeRequest {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resizeMap.Lock()
	defer resizeMap.Unlock()
	pvcrs := []*PVCWithResizeRequest{}
	for _, pvcr := range resizeMap.pvcrs {
		pvcrs = append(pvcrs, pvcr)
	}
	resizeMap.pvcrs = map[types.UniquePVCName]*PVCWithResizeRequest{}
	return pvcrs
}
func (resizeMap *volumeResizeMap) DeletePVC(pvc *v1.PersistentVolumeClaim) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pvcUniqueName := types.UniquePVCName(pvc.UID)
	klog.V(5).Infof("Removing PVC %v from resize map", pvcUniqueName)
	resizeMap.Lock()
	defer resizeMap.Unlock()
	delete(resizeMap.pvcrs, pvcUniqueName)
}
func (resizeMap *volumeResizeMap) MarkAsResized(pvcr *PVCWithResizeRequest, newSize resource.Quantity) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	emptyCondition := []v1.PersistentVolumeClaimCondition{}
	err := resizeMap.updatePVCCapacityAndConditions(pvcr, newSize, emptyCondition)
	if err != nil {
		klog.V(4).Infof("Error updating PV spec capacity for volume %q with : %v", pvcr.QualifiedName(), err)
		return err
	}
	return nil
}
func (resizeMap *volumeResizeMap) MarkForFSResize(pvcr *PVCWithResizeRequest) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pvcCondition := v1.PersistentVolumeClaimCondition{Type: v1.PersistentVolumeClaimFileSystemResizePending, Status: v1.ConditionTrue, LastTransitionTime: metav1.Now(), Message: "Waiting for user to (re-)start a pod to finish file system resize of volume on node."}
	conditions := []v1.PersistentVolumeClaimCondition{pvcCondition}
	newPVC := pvcr.PVC.DeepCopy()
	newPVC = util.MergeResizeConditionOnPVC(newPVC, conditions)
	_, err := util.PatchPVCStatus(pvcr.PVC, newPVC, resizeMap.kubeClient)
	return err
}
func (resizeMap *volumeResizeMap) UpdatePVSize(pvcr *PVCWithResizeRequest, newSize resource.Quantity) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldPv := pvcr.PersistentVolume
	pvClone := oldPv.DeepCopy()
	oldData, err := json.Marshal(pvClone)
	if err != nil {
		return fmt.Errorf("Unexpected error marshaling PV : %q with error %v", pvClone.Name, err)
	}
	pvClone.Spec.Capacity[v1.ResourceStorage] = newSize
	newData, err := json.Marshal(pvClone)
	if err != nil {
		return fmt.Errorf("Unexpected error marshaling PV : %q with error %v", pvClone.Name, err)
	}
	patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, pvClone)
	if err != nil {
		return fmt.Errorf("Error Creating two way merge patch for  PV : %q with error %v", pvClone.Name, err)
	}
	_, updateErr := resizeMap.kubeClient.CoreV1().PersistentVolumes().Patch(pvClone.Name, commontypes.StrategicMergePatchType, patchBytes)
	if updateErr != nil {
		klog.V(4).Infof("Error updating pv %q with error : %v", pvClone.Name, updateErr)
		return updateErr
	}
	return nil
}
func (resizeMap *volumeResizeMap) updatePVCCapacityAndConditions(pvcr *PVCWithResizeRequest, newSize resource.Quantity, pvcConditions []v1.PersistentVolumeClaimCondition) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newPVC := pvcr.PVC.DeepCopy()
	newPVC.Status.Capacity[v1.ResourceStorage] = newSize
	newPVC = util.MergeResizeConditionOnPVC(newPVC, pvcConditions)
	_, err := util.PatchPVCStatus(pvcr.PVC, newPVC, resizeMap.kubeClient)
	return err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

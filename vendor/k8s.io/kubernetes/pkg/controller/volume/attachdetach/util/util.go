package util

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/cache"
	"k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func CreateVolumeSpec(podVolume v1.Volume, podNamespace string, pvcLister corelisters.PersistentVolumeClaimLister, pvLister corelisters.PersistentVolumeLister) (*volume.Spec, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pvcSource := podVolume.VolumeSource.PersistentVolumeClaim; pvcSource != nil {
		klog.V(10).Infof("Found PVC, ClaimName: %q/%q", podNamespace, pvcSource.ClaimName)
		pvName, pvcUID, err := getPVCFromCacheExtractPV(podNamespace, pvcSource.ClaimName, pvcLister)
		if err != nil {
			return nil, fmt.Errorf("error processing PVC %q/%q: %v", podNamespace, pvcSource.ClaimName, err)
		}
		klog.V(10).Infof("Found bound PV for PVC (ClaimName %q/%q pvcUID %v): pvName=%q", podNamespace, pvcSource.ClaimName, pvcUID, pvName)
		volumeSpec, err := getPVSpecFromCache(pvName, pvcSource.ReadOnly, pvcUID, pvLister)
		if err != nil {
			return nil, fmt.Errorf("error processing PVC %q/%q: %v", podNamespace, pvcSource.ClaimName, err)
		}
		klog.V(10).Infof("Extracted volumeSpec (%v) from bound PV (pvName %q) and PVC (ClaimName %q/%q pvcUID %v)", volumeSpec.Name(), pvName, podNamespace, pvcSource.ClaimName, pvcUID)
		return volumeSpec, nil
	}
	clonedPodVolume := podVolume.DeepCopy()
	return volume.NewSpecFromVolume(clonedPodVolume), nil
}
func getPVCFromCacheExtractPV(namespace string, name string, pvcLister corelisters.PersistentVolumeClaimLister) (string, types.UID, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pvc, err := pvcLister.PersistentVolumeClaims(namespace).Get(name)
	if err != nil {
		return "", "", fmt.Errorf("failed to find PVC %s/%s in PVCInformer cache: %v", namespace, name, err)
	}
	if pvc.Status.Phase != v1.ClaimBound || pvc.Spec.VolumeName == "" {
		return "", "", fmt.Errorf("PVC %s/%s has non-bound phase (%q) or empty pvc.Spec.VolumeName (%q)", namespace, name, pvc.Status.Phase, pvc.Spec.VolumeName)
	}
	return pvc.Spec.VolumeName, pvc.UID, nil
}
func getPVSpecFromCache(name string, pvcReadOnly bool, expectedClaimUID types.UID, pvLister corelisters.PersistentVolumeLister) (*volume.Spec, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pv, err := pvLister.Get(name)
	if err != nil {
		return nil, fmt.Errorf("failed to find PV %q in PVInformer cache: %v", name, err)
	}
	if pv.Spec.ClaimRef == nil {
		return nil, fmt.Errorf("found PV object %q but it has a nil pv.Spec.ClaimRef indicating it is not yet bound to the claim", name)
	}
	if pv.Spec.ClaimRef.UID != expectedClaimUID {
		return nil, fmt.Errorf("found PV object %q but its pv.Spec.ClaimRef.UID (%q) does not point to claim.UID (%q)", name, pv.Spec.ClaimRef.UID, expectedClaimUID)
	}
	clonedPV := pv.DeepCopy()
	return volume.NewSpecFromPersistentVolume(clonedPV, pvcReadOnly), nil
}
func DetermineVolumeAction(pod *v1.Pod, desiredStateOfWorld cache.DesiredStateOfWorld, defaultAction bool) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pod == nil || len(pod.Spec.Volumes) <= 0 {
		return defaultAction
	}
	nodeName := types.NodeName(pod.Spec.NodeName)
	keepTerminatedPodVolume := desiredStateOfWorld.GetKeepTerminatedPodVolumesForNode(nodeName)
	if util.IsPodTerminated(pod, pod.Status) {
		return keepTerminatedPodVolume
	}
	return defaultAction
}
func ProcessPodVolumes(pod *v1.Pod, addVolumes bool, desiredStateOfWorld cache.DesiredStateOfWorld, volumePluginMgr *volume.VolumePluginMgr, pvcLister corelisters.PersistentVolumeClaimLister, pvLister corelisters.PersistentVolumeLister) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pod == nil {
		return
	}
	if len(pod.Spec.Volumes) <= 0 {
		klog.V(10).Infof("Skipping processing of pod %q/%q: it has no volumes.", pod.Namespace, pod.Name)
		return
	}
	nodeName := types.NodeName(pod.Spec.NodeName)
	if nodeName == "" {
		klog.V(10).Infof("Skipping processing of pod %q/%q: it is not scheduled to a node.", pod.Namespace, pod.Name)
		return
	} else if !desiredStateOfWorld.NodeExists(nodeName) {
		klog.V(4).Infof("Skipping processing of pod %q/%q: it is scheduled to node %q which is not managed by the controller.", pod.Namespace, pod.Name, nodeName)
		return
	}
	for _, podVolume := range pod.Spec.Volumes {
		volumeSpec, err := CreateVolumeSpec(podVolume, pod.Namespace, pvcLister, pvLister)
		if err != nil {
			klog.V(10).Infof("Error processing volume %q for pod %q/%q: %v", podVolume.Name, pod.Namespace, pod.Name, err)
			continue
		}
		attachableVolumePlugin, err := volumePluginMgr.FindAttachablePluginBySpec(volumeSpec)
		if err != nil || attachableVolumePlugin == nil {
			klog.V(10).Infof("Skipping volume %q for pod %q/%q: it does not implement attacher interface. err=%v", podVolume.Name, pod.Namespace, pod.Name, err)
			continue
		}
		uniquePodName := util.GetUniquePodName(pod)
		if addVolumes {
			_, err := desiredStateOfWorld.AddPod(uniquePodName, pod, volumeSpec, nodeName)
			if err != nil {
				klog.V(10).Infof("Failed to add volume %q for pod %q/%q to desiredStateOfWorld. %v", podVolume.Name, pod.Namespace, pod.Name, err)
			}
		} else {
			uniqueVolumeName, err := util.GetUniqueVolumeNameFromSpec(attachableVolumePlugin, volumeSpec)
			if err != nil {
				klog.V(10).Infof("Failed to delete volume %q for pod %q/%q from desiredStateOfWorld. GetUniqueVolumeNameFromSpec failed with %v", podVolume.Name, pod.Namespace, pod.Name, err)
				continue
			}
			desiredStateOfWorld.DeletePod(uniquePodName, uniqueVolumeName, nodeName)
		}
	}
	return
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

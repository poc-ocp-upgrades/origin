package cache

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util"
	"k8s.io/kubernetes/pkg/volume/util/operationexecutor"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	"time"
	gotime "time"
)

type ActualStateOfWorld interface {
	operationexecutor.ActualStateOfWorldAttacherUpdater
	AddVolumeNode(uniqueName v1.UniqueVolumeName, volumeSpec *volume.Spec, nodeName types.NodeName, devicePath string) (v1.UniqueVolumeName, error)
	SetVolumeMountedByNode(volumeName v1.UniqueVolumeName, nodeName types.NodeName, mounted bool) error
	SetNodeStatusUpdateNeeded(nodeName types.NodeName)
	ResetDetachRequestTime(volumeName v1.UniqueVolumeName, nodeName types.NodeName)
	SetDetachRequestTime(volumeName v1.UniqueVolumeName, nodeName types.NodeName) (time.Duration, error)
	DeleteVolumeNode(volumeName v1.UniqueVolumeName, nodeName types.NodeName)
	VolumeNodeExists(volumeName v1.UniqueVolumeName, nodeName types.NodeName) bool
	GetAttachedVolumes() []AttachedVolume
	GetAttachedVolumesForNode(nodeName types.NodeName) []AttachedVolume
	GetAttachedVolumesPerNode() map[types.NodeName][]operationexecutor.AttachedVolume
	GetNodesForVolume(volumeName v1.UniqueVolumeName) []types.NodeName
	GetVolumesToReportAttached() map[types.NodeName][]v1.AttachedVolume
	GetNodesToUpdateStatusFor() map[types.NodeName]nodeToUpdateStatusFor
}
type AttachedVolume struct {
	operationexecutor.AttachedVolume
	MountedByNode       bool
	DetachRequestedTime time.Time
}

func NewActualStateOfWorld(volumePluginMgr *volume.VolumePluginMgr) ActualStateOfWorld {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &actualStateOfWorld{attachedVolumes: make(map[v1.UniqueVolumeName]attachedVolume), nodesToUpdateStatusFor: make(map[types.NodeName]nodeToUpdateStatusFor), volumePluginMgr: volumePluginMgr}
}

type actualStateOfWorld struct {
	attachedVolumes        map[v1.UniqueVolumeName]attachedVolume
	nodesToUpdateStatusFor map[types.NodeName]nodeToUpdateStatusFor
	volumePluginMgr        *volume.VolumePluginMgr
	sync.RWMutex
}
type attachedVolume struct {
	volumeName      v1.UniqueVolumeName
	spec            *volume.Spec
	nodesAttachedTo map[types.NodeName]nodeAttachedTo
	devicePath      string
}
type nodeAttachedTo struct {
	nodeName              types.NodeName
	mountedByNode         bool
	mountedByNodeSetCount uint
	detachRequestedTime   time.Time
}
type nodeToUpdateStatusFor struct {
	nodeName                  types.NodeName
	statusUpdateNeeded        bool
	volumesToReportAsAttached map[v1.UniqueVolumeName]v1.UniqueVolumeName
}

func (asw *actualStateOfWorld) MarkVolumeAsAttached(uniqueName v1.UniqueVolumeName, volumeSpec *volume.Spec, nodeName types.NodeName, devicePath string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, err := asw.AddVolumeNode(uniqueName, volumeSpec, nodeName, devicePath)
	return err
}
func (asw *actualStateOfWorld) MarkVolumeAsDetached(volumeName v1.UniqueVolumeName, nodeName types.NodeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.DeleteVolumeNode(volumeName, nodeName)
}
func (asw *actualStateOfWorld) RemoveVolumeFromReportAsAttached(volumeName v1.UniqueVolumeName, nodeName types.NodeName) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.Lock()
	defer asw.Unlock()
	return asw.removeVolumeFromReportAsAttached(volumeName, nodeName)
}
func (asw *actualStateOfWorld) AddVolumeToReportAsAttached(volumeName v1.UniqueVolumeName, nodeName types.NodeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.Lock()
	defer asw.Unlock()
	asw.addVolumeToReportAsAttached(volumeName, nodeName)
}
func (asw *actualStateOfWorld) AddVolumeNode(uniqueName v1.UniqueVolumeName, volumeSpec *volume.Spec, nodeName types.NodeName, devicePath string) (v1.UniqueVolumeName, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.Lock()
	defer asw.Unlock()
	var volumeName v1.UniqueVolumeName
	if volumeSpec != nil {
		attachableVolumePlugin, err := asw.volumePluginMgr.FindAttachablePluginBySpec(volumeSpec)
		if err != nil || attachableVolumePlugin == nil {
			return "", fmt.Errorf("failed to get AttachablePlugin from volumeSpec for volume %q err=%v", volumeSpec.Name(), err)
		}
		volumeName, err = util.GetUniqueVolumeNameFromSpec(attachableVolumePlugin, volumeSpec)
		if err != nil {
			return "", fmt.Errorf("failed to GetUniqueVolumeNameFromSpec for volumeSpec %q err=%v", volumeSpec.Name(), err)
		}
	} else {
		volumeName = uniqueName
	}
	volumeObj, volumeExists := asw.attachedVolumes[volumeName]
	if !volumeExists {
		volumeObj = attachedVolume{volumeName: volumeName, spec: volumeSpec, nodesAttachedTo: make(map[types.NodeName]nodeAttachedTo), devicePath: devicePath}
	} else {
		volumeObj.devicePath = devicePath
		volumeObj.spec = volumeSpec
		klog.V(2).Infof("Volume %q is already added to attachedVolume list to node %q, update device path %q", volumeName, nodeName, devicePath)
	}
	asw.attachedVolumes[volumeName] = volumeObj
	_, nodeExists := volumeObj.nodesAttachedTo[nodeName]
	if !nodeExists {
		volumeObj.nodesAttachedTo[nodeName] = nodeAttachedTo{nodeName: nodeName, mountedByNode: true, mountedByNodeSetCount: 0, detachRequestedTime: time.Time{}}
	} else {
		klog.V(5).Infof("Volume %q is already added to attachedVolume list to the node %q", volumeName, nodeName)
	}
	asw.addVolumeToReportAsAttached(volumeName, nodeName)
	return volumeName, nil
}
func (asw *actualStateOfWorld) SetVolumeMountedByNode(volumeName v1.UniqueVolumeName, nodeName types.NodeName, mounted bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.Lock()
	defer asw.Unlock()
	volumeObj, nodeObj, err := asw.getNodeAndVolume(volumeName, nodeName)
	if err != nil {
		return fmt.Errorf("Failed to SetVolumeMountedByNode with error: %v", err)
	}
	if mounted {
		nodeObj.mountedByNodeSetCount = nodeObj.mountedByNodeSetCount + 1
	}
	nodeObj.mountedByNode = mounted
	volumeObj.nodesAttachedTo[nodeName] = nodeObj
	klog.V(4).Infof("SetVolumeMountedByNode volume %v to the node %q mounted %t", volumeName, nodeName, mounted)
	return nil
}
func (asw *actualStateOfWorld) ResetDetachRequestTime(volumeName v1.UniqueVolumeName, nodeName types.NodeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.Lock()
	defer asw.Unlock()
	volumeObj, nodeObj, err := asw.getNodeAndVolume(volumeName, nodeName)
	if err != nil {
		klog.Errorf("Failed to ResetDetachRequestTime with error: %v", err)
		return
	}
	nodeObj.detachRequestedTime = time.Time{}
	volumeObj.nodesAttachedTo[nodeName] = nodeObj
}
func (asw *actualStateOfWorld) SetDetachRequestTime(volumeName v1.UniqueVolumeName, nodeName types.NodeName) (time.Duration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.Lock()
	defer asw.Unlock()
	volumeObj, nodeObj, err := asw.getNodeAndVolume(volumeName, nodeName)
	if err != nil {
		return 0, fmt.Errorf("Failed to set detach request time with error: %v", err)
	}
	if nodeObj.detachRequestedTime.IsZero() {
		nodeObj.detachRequestedTime = time.Now()
		volumeObj.nodesAttachedTo[nodeName] = nodeObj
		klog.V(4).Infof("Set detach request time to current time for volume %v on node %q", volumeName, nodeName)
	}
	return time.Since(nodeObj.detachRequestedTime), nil
}
func (asw *actualStateOfWorld) getNodeAndVolume(volumeName v1.UniqueVolumeName, nodeName types.NodeName) (attachedVolume, nodeAttachedTo, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumeObj, volumeExists := asw.attachedVolumes[volumeName]
	if volumeExists {
		nodeObj, nodeExists := volumeObj.nodesAttachedTo[nodeName]
		if nodeExists {
			return volumeObj, nodeObj, nil
		}
	}
	return attachedVolume{}, nodeAttachedTo{}, fmt.Errorf("volume %v is no longer attached to the node %q", volumeName, nodeName)
}
func (asw *actualStateOfWorld) removeVolumeFromReportAsAttached(volumeName v1.UniqueVolumeName, nodeName types.NodeName) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeToUpdate, nodeToUpdateExists := asw.nodesToUpdateStatusFor[nodeName]
	if nodeToUpdateExists {
		_, nodeToUpdateVolumeExists := nodeToUpdate.volumesToReportAsAttached[volumeName]
		if nodeToUpdateVolumeExists {
			nodeToUpdate.statusUpdateNeeded = true
			delete(nodeToUpdate.volumesToReportAsAttached, volumeName)
			asw.nodesToUpdateStatusFor[nodeName] = nodeToUpdate
			return nil
		}
	}
	return fmt.Errorf("volume %q does not exist in volumesToReportAsAttached list or node %q does not exist in nodesToUpdateStatusFor list", volumeName, nodeName)
}
func (asw *actualStateOfWorld) addVolumeToReportAsAttached(volumeName v1.UniqueVolumeName, nodeName types.NodeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, _, err := asw.getNodeAndVolume(volumeName, nodeName); err != nil {
		klog.V(4).Infof("Volume %q is no longer attached to node %q", volumeName, nodeName)
		return
	}
	nodeToUpdate, nodeToUpdateExists := asw.nodesToUpdateStatusFor[nodeName]
	if !nodeToUpdateExists {
		nodeToUpdate = nodeToUpdateStatusFor{nodeName: nodeName, statusUpdateNeeded: true, volumesToReportAsAttached: make(map[v1.UniqueVolumeName]v1.UniqueVolumeName)}
		asw.nodesToUpdateStatusFor[nodeName] = nodeToUpdate
		klog.V(4).Infof("Add new node %q to nodesToUpdateStatusFor", nodeName)
	}
	_, nodeToUpdateVolumeExists := nodeToUpdate.volumesToReportAsAttached[volumeName]
	if !nodeToUpdateVolumeExists {
		nodeToUpdate.statusUpdateNeeded = true
		nodeToUpdate.volumesToReportAsAttached[volumeName] = volumeName
		asw.nodesToUpdateStatusFor[nodeName] = nodeToUpdate
		klog.V(4).Infof("Report volume %q as attached to node %q", volumeName, nodeName)
	}
}
func (asw *actualStateOfWorld) updateNodeStatusUpdateNeeded(nodeName types.NodeName, needed bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeToUpdate, nodeToUpdateExists := asw.nodesToUpdateStatusFor[nodeName]
	if !nodeToUpdateExists {
		errMsg := fmt.Sprintf("Failed to set statusUpdateNeeded to needed %t, because nodeName=%q does not exist", needed, nodeName)
		return fmt.Errorf(errMsg)
	}
	nodeToUpdate.statusUpdateNeeded = needed
	asw.nodesToUpdateStatusFor[nodeName] = nodeToUpdate
	return nil
}
func (asw *actualStateOfWorld) SetNodeStatusUpdateNeeded(nodeName types.NodeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.Lock()
	defer asw.Unlock()
	if err := asw.updateNodeStatusUpdateNeeded(nodeName, true); err != nil {
		klog.Warningf("Failed to update statusUpdateNeeded field in actual state of world: %v", err)
	}
}
func (asw *actualStateOfWorld) DeleteVolumeNode(volumeName v1.UniqueVolumeName, nodeName types.NodeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.Lock()
	defer asw.Unlock()
	volumeObj, volumeExists := asw.attachedVolumes[volumeName]
	if !volumeExists {
		return
	}
	_, nodeExists := volumeObj.nodesAttachedTo[nodeName]
	if nodeExists {
		delete(asw.attachedVolumes[volumeName].nodesAttachedTo, nodeName)
	}
	if len(volumeObj.nodesAttachedTo) == 0 {
		delete(asw.attachedVolumes, volumeName)
	}
	asw.removeVolumeFromReportAsAttached(volumeName, nodeName)
}
func (asw *actualStateOfWorld) VolumeNodeExists(volumeName v1.UniqueVolumeName, nodeName types.NodeName) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.RLock()
	defer asw.RUnlock()
	volumeObj, volumeExists := asw.attachedVolumes[volumeName]
	if volumeExists {
		if _, nodeExists := volumeObj.nodesAttachedTo[nodeName]; nodeExists {
			return true
		}
	}
	return false
}
func (asw *actualStateOfWorld) GetAttachedVolumes() []AttachedVolume {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.RLock()
	defer asw.RUnlock()
	attachedVolumes := make([]AttachedVolume, 0, len(asw.attachedVolumes))
	for _, volumeObj := range asw.attachedVolumes {
		for _, nodeObj := range volumeObj.nodesAttachedTo {
			attachedVolumes = append(attachedVolumes, getAttachedVolume(&volumeObj, &nodeObj))
		}
	}
	return attachedVolumes
}
func (asw *actualStateOfWorld) GetAttachedVolumesForNode(nodeName types.NodeName) []AttachedVolume {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.RLock()
	defer asw.RUnlock()
	attachedVolumes := make([]AttachedVolume, 0, len(asw.attachedVolumes))
	for _, volumeObj := range asw.attachedVolumes {
		for actualNodeName, nodeObj := range volumeObj.nodesAttachedTo {
			if actualNodeName == nodeName {
				attachedVolumes = append(attachedVolumes, getAttachedVolume(&volumeObj, &nodeObj))
				break
			}
		}
	}
	return attachedVolumes
}
func (asw *actualStateOfWorld) GetAttachedVolumesPerNode() map[types.NodeName][]operationexecutor.AttachedVolume {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.RLock()
	defer asw.RUnlock()
	attachedVolumesPerNode := make(map[types.NodeName][]operationexecutor.AttachedVolume)
	for _, volumeObj := range asw.attachedVolumes {
		for nodeName, nodeObj := range volumeObj.nodesAttachedTo {
			volumes := attachedVolumesPerNode[nodeName]
			volumes = append(volumes, getAttachedVolume(&volumeObj, &nodeObj).AttachedVolume)
			attachedVolumesPerNode[nodeName] = volumes
		}
	}
	return attachedVolumesPerNode
}
func (asw *actualStateOfWorld) GetNodesForVolume(volumeName v1.UniqueVolumeName) []types.NodeName {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.RLock()
	defer asw.RUnlock()
	volumeObj, volumeExists := asw.attachedVolumes[volumeName]
	if !volumeExists || len(volumeObj.nodesAttachedTo) == 0 {
		return []types.NodeName{}
	}
	nodes := []types.NodeName{}
	for k := range volumeObj.nodesAttachedTo {
		nodes = append(nodes, k)
	}
	return nodes
}
func (asw *actualStateOfWorld) GetVolumesToReportAttached() map[types.NodeName][]v1.AttachedVolume {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	asw.RLock()
	defer asw.RUnlock()
	volumesToReportAttached := make(map[types.NodeName][]v1.AttachedVolume)
	for nodeName, nodeToUpdateObj := range asw.nodesToUpdateStatusFor {
		if nodeToUpdateObj.statusUpdateNeeded {
			attachedVolumes := make([]v1.AttachedVolume, len(nodeToUpdateObj.volumesToReportAsAttached))
			i := 0
			for _, volume := range nodeToUpdateObj.volumesToReportAsAttached {
				attachedVolumes[i] = v1.AttachedVolume{Name: volume, DevicePath: asw.attachedVolumes[volume].devicePath}
				i++
			}
			volumesToReportAttached[nodeToUpdateObj.nodeName] = attachedVolumes
		}
		if err := asw.updateNodeStatusUpdateNeeded(nodeName, false); err != nil {
			klog.Errorf("Failed to update statusUpdateNeeded field when getting volumes: %v", err)
		}
	}
	return volumesToReportAttached
}
func (asw *actualStateOfWorld) GetNodesToUpdateStatusFor() map[types.NodeName]nodeToUpdateStatusFor {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return asw.nodesToUpdateStatusFor
}
func getAttachedVolume(attachedVolume *attachedVolume, nodeAttachedTo *nodeAttachedTo) AttachedVolume {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return AttachedVolume{AttachedVolume: operationexecutor.AttachedVolume{VolumeName: attachedVolume.volumeName, VolumeSpec: attachedVolume.spec, NodeName: nodeAttachedTo.nodeName, DevicePath: attachedVolume.devicePath, PluginIsAttachable: true}, MountedByNode: nodeAttachedTo.mountedByNode, DetachRequestedTime: nodeAttachedTo.detachRequestedTime}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

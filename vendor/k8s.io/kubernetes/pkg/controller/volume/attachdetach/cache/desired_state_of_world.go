package cache

import (
	"fmt"
	"k8s.io/api/core/v1"
	k8stypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util"
	"k8s.io/kubernetes/pkg/volume/util/operationexecutor"
	"k8s.io/kubernetes/pkg/volume/util/types"
	"sync"
)

type DesiredStateOfWorld interface {
	AddNode(nodeName k8stypes.NodeName, keepTerminatedPodVolumes bool)
	AddPod(podName types.UniquePodName, pod *v1.Pod, volumeSpec *volume.Spec, nodeName k8stypes.NodeName) (v1.UniqueVolumeName, error)
	DeleteNode(nodeName k8stypes.NodeName) error
	DeletePod(podName types.UniquePodName, volumeName v1.UniqueVolumeName, nodeName k8stypes.NodeName)
	NodeExists(nodeName k8stypes.NodeName) bool
	VolumeExists(volumeName v1.UniqueVolumeName, nodeName k8stypes.NodeName) bool
	GetVolumesToAttach() []VolumeToAttach
	GetPodToAdd() map[types.UniquePodName]PodToAdd
	GetKeepTerminatedPodVolumesForNode(k8stypes.NodeName) bool
	SetMultiAttachError(v1.UniqueVolumeName, k8stypes.NodeName)
	GetVolumePodsOnNodes(nodes []k8stypes.NodeName, volumeName v1.UniqueVolumeName) []*v1.Pod
}
type VolumeToAttach struct {
	operationexecutor.VolumeToAttach
}
type PodToAdd struct {
	Pod        *v1.Pod
	VolumeName v1.UniqueVolumeName
	NodeName   k8stypes.NodeName
}

func NewDesiredStateOfWorld(volumePluginMgr *volume.VolumePluginMgr) DesiredStateOfWorld {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &desiredStateOfWorld{nodesManaged: make(map[k8stypes.NodeName]nodeManaged), volumePluginMgr: volumePluginMgr}
}

type desiredStateOfWorld struct {
	nodesManaged    map[k8stypes.NodeName]nodeManaged
	volumePluginMgr *volume.VolumePluginMgr
	sync.RWMutex
}
type nodeManaged struct {
	nodeName                 k8stypes.NodeName
	volumesToAttach          map[v1.UniqueVolumeName]volumeToAttach
	keepTerminatedPodVolumes bool
}
type volumeToAttach struct {
	multiAttachErrorReported bool
	volumeName               v1.UniqueVolumeName
	spec                     *volume.Spec
	scheduledPods            map[types.UniquePodName]pod
}
type pod struct {
	podName types.UniquePodName
	podObj  *v1.Pod
}

func (dsw *desiredStateOfWorld) AddNode(nodeName k8stypes.NodeName, keepTerminatedPodVolumes bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.Lock()
	defer dsw.Unlock()
	if _, nodeExists := dsw.nodesManaged[nodeName]; !nodeExists {
		dsw.nodesManaged[nodeName] = nodeManaged{nodeName: nodeName, volumesToAttach: make(map[v1.UniqueVolumeName]volumeToAttach), keepTerminatedPodVolumes: keepTerminatedPodVolumes}
	}
}
func (dsw *desiredStateOfWorld) AddPod(podName types.UniquePodName, podToAdd *v1.Pod, volumeSpec *volume.Spec, nodeName k8stypes.NodeName) (v1.UniqueVolumeName, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.Lock()
	defer dsw.Unlock()
	nodeObj, nodeExists := dsw.nodesManaged[nodeName]
	if !nodeExists {
		return "", fmt.Errorf("no node with the name %q exists in the list of managed nodes", nodeName)
	}
	attachableVolumePlugin, err := dsw.volumePluginMgr.FindAttachablePluginBySpec(volumeSpec)
	if err != nil || attachableVolumePlugin == nil {
		return "", fmt.Errorf("failed to get AttachablePlugin from volumeSpec for volume %q err=%v", volumeSpec.Name(), err)
	}
	volumeName, err := util.GetUniqueVolumeNameFromSpec(attachableVolumePlugin, volumeSpec)
	if err != nil {
		return "", fmt.Errorf("failed to get UniqueVolumeName from volumeSpec for plugin=%q and volume=%q err=%v", attachableVolumePlugin.GetPluginName(), volumeSpec.Name(), err)
	}
	volumeObj, volumeExists := nodeObj.volumesToAttach[volumeName]
	if !volumeExists {
		volumeObj = volumeToAttach{multiAttachErrorReported: false, volumeName: volumeName, spec: volumeSpec, scheduledPods: make(map[types.UniquePodName]pod)}
		dsw.nodesManaged[nodeName].volumesToAttach[volumeName] = volumeObj
	}
	if _, podExists := volumeObj.scheduledPods[podName]; !podExists {
		dsw.nodesManaged[nodeName].volumesToAttach[volumeName].scheduledPods[podName] = pod{podName: podName, podObj: podToAdd}
	}
	return volumeName, nil
}
func (dsw *desiredStateOfWorld) DeleteNode(nodeName k8stypes.NodeName) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.Lock()
	defer dsw.Unlock()
	nodeObj, nodeExists := dsw.nodesManaged[nodeName]
	if !nodeExists {
		return nil
	}
	if len(nodeObj.volumesToAttach) > 0 {
		return fmt.Errorf("failed to delete node %q from list of nodes managed by attach/detach controller--the node still contains %v volumes in its list of volumes to attach", nodeName, len(nodeObj.volumesToAttach))
	}
	delete(dsw.nodesManaged, nodeName)
	return nil
}
func (dsw *desiredStateOfWorld) DeletePod(podName types.UniquePodName, volumeName v1.UniqueVolumeName, nodeName k8stypes.NodeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.Lock()
	defer dsw.Unlock()
	nodeObj, nodeExists := dsw.nodesManaged[nodeName]
	if !nodeExists {
		return
	}
	volumeObj, volumeExists := nodeObj.volumesToAttach[volumeName]
	if !volumeExists {
		return
	}
	if _, podExists := volumeObj.scheduledPods[podName]; !podExists {
		return
	}
	delete(dsw.nodesManaged[nodeName].volumesToAttach[volumeName].scheduledPods, podName)
	if len(volumeObj.scheduledPods) == 0 {
		delete(dsw.nodesManaged[nodeName].volumesToAttach, volumeName)
	}
}
func (dsw *desiredStateOfWorld) NodeExists(nodeName k8stypes.NodeName) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.RLock()
	defer dsw.RUnlock()
	_, nodeExists := dsw.nodesManaged[nodeName]
	return nodeExists
}
func (dsw *desiredStateOfWorld) VolumeExists(volumeName v1.UniqueVolumeName, nodeName k8stypes.NodeName) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.RLock()
	defer dsw.RUnlock()
	nodeObj, nodeExists := dsw.nodesManaged[nodeName]
	if nodeExists {
		if _, volumeExists := nodeObj.volumesToAttach[volumeName]; volumeExists {
			return true
		}
	}
	return false
}
func (dsw *desiredStateOfWorld) SetMultiAttachError(volumeName v1.UniqueVolumeName, nodeName k8stypes.NodeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.Lock()
	defer dsw.Unlock()
	nodeObj, nodeExists := dsw.nodesManaged[nodeName]
	if nodeExists {
		if volumeObj, volumeExists := nodeObj.volumesToAttach[volumeName]; volumeExists {
			volumeObj.multiAttachErrorReported = true
			dsw.nodesManaged[nodeName].volumesToAttach[volumeName] = volumeObj
		}
	}
}
func (dsw *desiredStateOfWorld) GetKeepTerminatedPodVolumesForNode(nodeName k8stypes.NodeName) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.RLock()
	defer dsw.RUnlock()
	if nodeName == "" {
		return false
	}
	if node, ok := dsw.nodesManaged[nodeName]; ok {
		return node.keepTerminatedPodVolumes
	}
	return false
}
func (dsw *desiredStateOfWorld) GetVolumesToAttach() []VolumeToAttach {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.RLock()
	defer dsw.RUnlock()
	volumesToAttach := make([]VolumeToAttach, 0, len(dsw.nodesManaged))
	for nodeName, nodeObj := range dsw.nodesManaged {
		for volumeName, volumeObj := range nodeObj.volumesToAttach {
			volumesToAttach = append(volumesToAttach, VolumeToAttach{VolumeToAttach: operationexecutor.VolumeToAttach{MultiAttachErrorReported: volumeObj.multiAttachErrorReported, VolumeName: volumeName, VolumeSpec: volumeObj.spec, NodeName: nodeName, ScheduledPods: getPodsFromMap(volumeObj.scheduledPods)}})
		}
	}
	return volumesToAttach
}
func getPodsFromMap(podMap map[types.UniquePodName]pod) []*v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pods := make([]*v1.Pod, 0, len(podMap))
	for _, pod := range podMap {
		pods = append(pods, pod.podObj)
	}
	return pods
}
func (dsw *desiredStateOfWorld) GetPodToAdd() map[types.UniquePodName]PodToAdd {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.RLock()
	defer dsw.RUnlock()
	pods := make(map[types.UniquePodName]PodToAdd)
	for nodeName, nodeObj := range dsw.nodesManaged {
		for volumeName, volumeObj := range nodeObj.volumesToAttach {
			for podUID, pod := range volumeObj.scheduledPods {
				pods[podUID] = PodToAdd{Pod: pod.podObj, VolumeName: volumeName, NodeName: nodeName}
			}
		}
	}
	return pods
}
func (dsw *desiredStateOfWorld) GetVolumePodsOnNodes(nodes []k8stypes.NodeName, volumeName v1.UniqueVolumeName) []*v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsw.RLock()
	defer dsw.RUnlock()
	pods := []*v1.Pod{}
	for _, nodeName := range nodes {
		node, ok := dsw.nodesManaged[nodeName]
		if !ok {
			continue
		}
		volume, ok := node.volumesToAttach[volumeName]
		if !ok {
			continue
		}
		for _, pod := range volume.scheduledPods {
			pods = append(pods, pod.podObj)
		}
	}
	return pods
}

package predicates

import (
	"errors"
	"fmt"
	"k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/sets"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	corelisters "k8s.io/client-go/listers/core/v1"
	storagelisters "k8s.io/client-go/listers/storage/v1"
	"k8s.io/klog"
	v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
	v1qos "k8s.io/kubernetes/pkg/apis/core/v1/helper/qos"
	"k8s.io/kubernetes/pkg/features"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	priorityutil "k8s.io/kubernetes/pkg/scheduler/algorithm/priorities/util"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
	schedutil "k8s.io/kubernetes/pkg/scheduler/util"
	"k8s.io/kubernetes/pkg/scheduler/volumebinder"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
	"os"
	"regexp"
	"strconv"
)

const (
	MatchInterPodAffinityPred           = "MatchInterPodAffinity"
	CheckVolumeBindingPred              = "CheckVolumeBinding"
	CheckNodeConditionPred              = "CheckNodeCondition"
	GeneralPred                         = "GeneralPredicates"
	HostNamePred                        = "HostName"
	PodFitsHostPortsPred                = "PodFitsHostPorts"
	MatchNodeSelectorPred               = "MatchNodeSelector"
	PodFitsResourcesPred                = "PodFitsResources"
	NoDiskConflictPred                  = "NoDiskConflict"
	PodToleratesNodeTaintsPred          = "PodToleratesNodeTaints"
	CheckNodeUnschedulablePred          = "CheckNodeUnschedulable"
	PodToleratesNodeNoExecuteTaintsPred = "PodToleratesNodeNoExecuteTaints"
	CheckNodeLabelPresencePred          = "CheckNodeLabelPresence"
	CheckServiceAffinityPred            = "CheckServiceAffinity"
	MaxEBSVolumeCountPred               = "MaxEBSVolumeCount"
	MaxGCEPDVolumeCountPred             = "MaxGCEPDVolumeCount"
	MaxAzureDiskVolumeCountPred         = "MaxAzureDiskVolumeCount"
	MaxCinderVolumeCountPred            = "MaxCinderVolumeCount"
	MaxCSIVolumeCountPred               = "MaxCSIVolumeCountPred"
	NoVolumeZoneConflictPred            = "NoVolumeZoneConflict"
	CheckNodeMemoryPressurePred         = "CheckNodeMemoryPressure"
	CheckNodeDiskPressurePred           = "CheckNodeDiskPressure"
	CheckNodePIDPressurePred            = "CheckNodePIDPressure"
	DefaultMaxGCEPDVolumes              = 16
	DefaultMaxAzureDiskVolumes          = 16
	KubeMaxPDVols                       = "KUBE_MAX_PD_VOLS"
	EBSVolumeFilterType                 = "EBS"
	GCEPDVolumeFilterType               = "GCE"
	AzureDiskVolumeFilterType           = "AzureDisk"
	CinderVolumeFilterType              = "Cinder"
)

var (
	predicatesOrdering = []string{CheckNodeConditionPred, CheckNodeUnschedulablePred, GeneralPred, HostNamePred, PodFitsHostPortsPred, MatchNodeSelectorPred, PodFitsResourcesPred, NoDiskConflictPred, PodToleratesNodeTaintsPred, PodToleratesNodeNoExecuteTaintsPred, CheckNodeLabelPresencePred, CheckServiceAffinityPred, MaxEBSVolumeCountPred, MaxGCEPDVolumeCountPred, MaxCSIVolumeCountPred, MaxAzureDiskVolumeCountPred, MaxCinderVolumeCountPred, CheckVolumeBindingPred, NoVolumeZoneConflictPred, CheckNodeMemoryPressurePred, CheckNodePIDPressurePred, CheckNodeDiskPressurePred, MatchInterPodAffinityPred}
)

type NodeInfo interface {
	GetNodeInfo(nodeID string) (*v1.Node, error)
}
type PersistentVolumeInfo interface {
	GetPersistentVolumeInfo(pvID string) (*v1.PersistentVolume, error)
}
type CachedPersistentVolumeInfo struct {
	corelisters.PersistentVolumeLister
}

func Ordering() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return predicatesOrdering
}
func SetPredicatesOrdering(names []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	predicatesOrdering = names
}
func (c *CachedPersistentVolumeInfo) GetPersistentVolumeInfo(pvID string) (*v1.PersistentVolume, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Get(pvID)
}

type PersistentVolumeClaimInfo interface {
	GetPersistentVolumeClaimInfo(namespace string, name string) (*v1.PersistentVolumeClaim, error)
}
type CachedPersistentVolumeClaimInfo struct {
	corelisters.PersistentVolumeClaimLister
}

func (c *CachedPersistentVolumeClaimInfo) GetPersistentVolumeClaimInfo(namespace string, name string) (*v1.PersistentVolumeClaim, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.PersistentVolumeClaims(namespace).Get(name)
}

type CachedNodeInfo struct{ corelisters.NodeLister }

func (c *CachedNodeInfo) GetNodeInfo(id string) (*v1.Node, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node, err := c.Get(id)
	if apierrors.IsNotFound(err) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("error retrieving node '%v' from cache: %v", id, err)
	}
	return node, nil
}

type StorageClassInfo interface {
	GetStorageClassInfo(className string) (*storagev1.StorageClass, error)
}
type CachedStorageClassInfo struct {
	storagelisters.StorageClassLister
}

func (c *CachedStorageClassInfo) GetStorageClassInfo(className string) (*storagev1.StorageClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Get(className)
}
func isVolumeConflict(volume v1.Volume, pod *v1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if volume.GCEPersistentDisk == nil && volume.AWSElasticBlockStore == nil && volume.RBD == nil && volume.ISCSI == nil {
		return false
	}
	for _, existingVolume := range pod.Spec.Volumes {
		if volume.GCEPersistentDisk != nil && existingVolume.GCEPersistentDisk != nil {
			disk, existingDisk := volume.GCEPersistentDisk, existingVolume.GCEPersistentDisk
			if disk.PDName == existingDisk.PDName && !(disk.ReadOnly && existingDisk.ReadOnly) {
				return true
			}
		}
		if volume.AWSElasticBlockStore != nil && existingVolume.AWSElasticBlockStore != nil {
			if volume.AWSElasticBlockStore.VolumeID == existingVolume.AWSElasticBlockStore.VolumeID {
				return true
			}
		}
		if volume.ISCSI != nil && existingVolume.ISCSI != nil {
			iqn := volume.ISCSI.IQN
			eiqn := existingVolume.ISCSI.IQN
			if iqn == eiqn && !(volume.ISCSI.ReadOnly && existingVolume.ISCSI.ReadOnly) {
				return true
			}
		}
		if volume.RBD != nil && existingVolume.RBD != nil {
			mon, pool, image := volume.RBD.CephMonitors, volume.RBD.RBDPool, volume.RBD.RBDImage
			emon, epool, eimage := existingVolume.RBD.CephMonitors, existingVolume.RBD.RBDPool, existingVolume.RBD.RBDImage
			if haveOverlap(mon, emon) && pool == epool && image == eimage && !(volume.RBD.ReadOnly && existingVolume.RBD.ReadOnly) {
				return true
			}
		}
	}
	return false
}
func NoDiskConflict(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, v := range pod.Spec.Volumes {
		for _, ev := range nodeInfo.Pods() {
			if isVolumeConflict(v, ev) {
				return false, []algorithm.PredicateFailureReason{ErrDiskConflict}, nil
			}
		}
	}
	return true, nil, nil
}

type MaxPDVolumeCountChecker struct {
	filter               VolumeFilter
	volumeLimitKey       v1.ResourceName
	maxVolumeFunc        func(node *v1.Node) int
	pvInfo               PersistentVolumeInfo
	pvcInfo              PersistentVolumeClaimInfo
	randomVolumeIDPrefix string
}
type VolumeFilter struct {
	FilterVolume           func(vol *v1.Volume) (id string, relevant bool)
	FilterPersistentVolume func(pv *v1.PersistentVolume) (id string, relevant bool)
}

func NewMaxPDVolumeCountPredicate(filterName string, pvInfo PersistentVolumeInfo, pvcInfo PersistentVolumeClaimInfo) algorithm.FitPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var filter VolumeFilter
	var volumeLimitKey v1.ResourceName
	switch filterName {
	case EBSVolumeFilterType:
		filter = EBSVolumeFilter
		volumeLimitKey = v1.ResourceName(volumeutil.EBSVolumeLimitKey)
	case GCEPDVolumeFilterType:
		filter = GCEPDVolumeFilter
		volumeLimitKey = v1.ResourceName(volumeutil.GCEVolumeLimitKey)
	case AzureDiskVolumeFilterType:
		filter = AzureDiskVolumeFilter
		volumeLimitKey = v1.ResourceName(volumeutil.AzureVolumeLimitKey)
	case CinderVolumeFilterType:
		filter = CinderVolumeFilter
		volumeLimitKey = v1.ResourceName(volumeutil.CinderVolumeLimitKey)
	default:
		klog.Fatalf("Wrong filterName, Only Support %v %v %v ", EBSVolumeFilterType, GCEPDVolumeFilterType, AzureDiskVolumeFilterType)
		return nil
	}
	c := &MaxPDVolumeCountChecker{filter: filter, volumeLimitKey: volumeLimitKey, maxVolumeFunc: getMaxVolumeFunc(filterName), pvInfo: pvInfo, pvcInfo: pvcInfo, randomVolumeIDPrefix: rand.String(32)}
	return c.predicate
}
func getMaxVolumeFunc(filterName string) func(node *v1.Node) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(node *v1.Node) int {
		maxVolumesFromEnv := getMaxVolLimitFromEnv()
		if maxVolumesFromEnv > 0 {
			return maxVolumesFromEnv
		}
		var nodeInstanceType string
		for k, v := range node.ObjectMeta.Labels {
			if k == kubeletapis.LabelInstanceType {
				nodeInstanceType = v
			}
		}
		switch filterName {
		case EBSVolumeFilterType:
			return getMaxEBSVolume(nodeInstanceType)
		case GCEPDVolumeFilterType:
			return DefaultMaxGCEPDVolumes
		case AzureDiskVolumeFilterType:
			return DefaultMaxAzureDiskVolumes
		case CinderVolumeFilterType:
			return volumeutil.DefaultMaxCinderVolumes
		default:
			return -1
		}
	}
}
func getMaxEBSVolume(nodeInstanceType string) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ok, _ := regexp.MatchString(volumeutil.EBSNitroLimitRegex, nodeInstanceType); ok {
		return volumeutil.DefaultMaxEBSNitroVolumeLimit
	}
	return volumeutil.DefaultMaxEBSVolumes
}
func getMaxVolLimitFromEnv() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if rawMaxVols := os.Getenv(KubeMaxPDVols); rawMaxVols != "" {
		if parsedMaxVols, err := strconv.Atoi(rawMaxVols); err != nil {
			klog.Errorf("Unable to parse maximum PD volumes value, using default: %v", err)
		} else if parsedMaxVols <= 0 {
			klog.Errorf("Maximum PD volumes must be a positive value, using default ")
		} else {
			return parsedMaxVols
		}
	}
	return -1
}
func (c *MaxPDVolumeCountChecker) filterVolumes(volumes []v1.Volume, namespace string, filteredVolumes map[string]bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range volumes {
		vol := &volumes[i]
		if id, ok := c.filter.FilterVolume(vol); ok {
			filteredVolumes[id] = true
		} else if vol.PersistentVolumeClaim != nil {
			pvcName := vol.PersistentVolumeClaim.ClaimName
			if pvcName == "" {
				return fmt.Errorf("PersistentVolumeClaim had no name")
			}
			pvID := fmt.Sprintf("%s-%s/%s", c.randomVolumeIDPrefix, namespace, pvcName)
			pvc, err := c.pvcInfo.GetPersistentVolumeClaimInfo(namespace, pvcName)
			if err != nil || pvc == nil {
				klog.V(4).Infof("Unable to look up PVC info for %s/%s, assuming PVC matches predicate when counting limits: %v", namespace, pvcName, err)
				filteredVolumes[pvID] = true
				continue
			}
			pvName := pvc.Spec.VolumeName
			if pvName == "" {
				klog.V(4).Infof("PVC %s/%s is not bound, assuming PVC matches predicate when counting limits", namespace, pvcName)
				filteredVolumes[pvID] = true
				continue
			}
			pv, err := c.pvInfo.GetPersistentVolumeInfo(pvName)
			if err != nil || pv == nil {
				klog.V(4).Infof("Unable to look up PV info for %s/%s/%s, assuming PV matches predicate when counting limits: %v", namespace, pvcName, pvName, err)
				filteredVolumes[pvID] = true
				continue
			}
			if id, ok := c.filter.FilterPersistentVolume(pv); ok {
				filteredVolumes[id] = true
			}
		}
	}
	return nil
}
func (c *MaxPDVolumeCountChecker) predicate(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(pod.Spec.Volumes) == 0 {
		return true, nil, nil
	}
	newVolumes := make(map[string]bool)
	if err := c.filterVolumes(pod.Spec.Volumes, pod.Namespace, newVolumes); err != nil {
		return false, nil, err
	}
	if len(newVolumes) == 0 {
		return true, nil, nil
	}
	existingVolumes := make(map[string]bool)
	for _, existingPod := range nodeInfo.Pods() {
		if err := c.filterVolumes(existingPod.Spec.Volumes, existingPod.Namespace, existingVolumes); err != nil {
			return false, nil, err
		}
	}
	numExistingVolumes := len(existingVolumes)
	for k := range existingVolumes {
		if _, ok := newVolumes[k]; ok {
			delete(newVolumes, k)
		}
	}
	numNewVolumes := len(newVolumes)
	maxAttachLimit := c.maxVolumeFunc(nodeInfo.Node())
	if utilfeature.DefaultFeatureGate.Enabled(features.AttachVolumeLimit) {
		volumeLimits := nodeInfo.VolumeLimits()
		if maxAttachLimitFromAllocatable, ok := volumeLimits[c.volumeLimitKey]; ok {
			maxAttachLimit = int(maxAttachLimitFromAllocatable)
		}
	}
	if numExistingVolumes+numNewVolumes > maxAttachLimit {
		return false, []algorithm.PredicateFailureReason{ErrMaxVolumeCountExceeded}, nil
	}
	if nodeInfo != nil && nodeInfo.TransientInfo != nil && utilfeature.DefaultFeatureGate.Enabled(features.BalanceAttachedNodeVolumes) {
		nodeInfo.TransientInfo.TransientLock.Lock()
		defer nodeInfo.TransientInfo.TransientLock.Unlock()
		nodeInfo.TransientInfo.TransNodeInfo.AllocatableVolumesCount = maxAttachLimit - numExistingVolumes
		nodeInfo.TransientInfo.TransNodeInfo.RequestedVolumes = numNewVolumes
	}
	return true, nil, nil
}

var EBSVolumeFilter = VolumeFilter{FilterVolume: func(vol *v1.Volume) (string, bool) {
	if vol.AWSElasticBlockStore != nil {
		return vol.AWSElasticBlockStore.VolumeID, true
	}
	return "", false
}, FilterPersistentVolume: func(pv *v1.PersistentVolume) (string, bool) {
	if pv.Spec.AWSElasticBlockStore != nil {
		return pv.Spec.AWSElasticBlockStore.VolumeID, true
	}
	return "", false
}}
var GCEPDVolumeFilter = VolumeFilter{FilterVolume: func(vol *v1.Volume) (string, bool) {
	if vol.GCEPersistentDisk != nil {
		return vol.GCEPersistentDisk.PDName, true
	}
	return "", false
}, FilterPersistentVolume: func(pv *v1.PersistentVolume) (string, bool) {
	if pv.Spec.GCEPersistentDisk != nil {
		return pv.Spec.GCEPersistentDisk.PDName, true
	}
	return "", false
}}
var AzureDiskVolumeFilter = VolumeFilter{FilterVolume: func(vol *v1.Volume) (string, bool) {
	if vol.AzureDisk != nil {
		return vol.AzureDisk.DiskName, true
	}
	return "", false
}, FilterPersistentVolume: func(pv *v1.PersistentVolume) (string, bool) {
	if pv.Spec.AzureDisk != nil {
		return pv.Spec.AzureDisk.DiskName, true
	}
	return "", false
}}
var CinderVolumeFilter = VolumeFilter{FilterVolume: func(vol *v1.Volume) (string, bool) {
	if vol.Cinder != nil {
		return vol.Cinder.VolumeID, true
	}
	return "", false
}, FilterPersistentVolume: func(pv *v1.PersistentVolume) (string, bool) {
	if pv.Spec.Cinder != nil {
		return pv.Spec.Cinder.VolumeID, true
	}
	return "", false
}}

type VolumeZoneChecker struct {
	pvInfo    PersistentVolumeInfo
	pvcInfo   PersistentVolumeClaimInfo
	classInfo StorageClassInfo
}

func NewVolumeZonePredicate(pvInfo PersistentVolumeInfo, pvcInfo PersistentVolumeClaimInfo, classInfo StorageClassInfo) algorithm.FitPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &VolumeZoneChecker{pvInfo: pvInfo, pvcInfo: pvcInfo, classInfo: classInfo}
	return c.predicate
}
func (c *VolumeZoneChecker) predicate(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(pod.Spec.Volumes) == 0 {
		return true, nil, nil
	}
	node := nodeInfo.Node()
	if node == nil {
		return false, nil, fmt.Errorf("node not found")
	}
	nodeConstraints := make(map[string]string)
	for k, v := range node.ObjectMeta.Labels {
		if k != kubeletapis.LabelZoneFailureDomain && k != kubeletapis.LabelZoneRegion {
			continue
		}
		nodeConstraints[k] = v
	}
	if len(nodeConstraints) == 0 {
		return true, nil, nil
	}
	namespace := pod.Namespace
	manifest := &(pod.Spec)
	for i := range manifest.Volumes {
		volume := &manifest.Volumes[i]
		if volume.PersistentVolumeClaim != nil {
			pvcName := volume.PersistentVolumeClaim.ClaimName
			if pvcName == "" {
				return false, nil, fmt.Errorf("PersistentVolumeClaim had no name")
			}
			pvc, err := c.pvcInfo.GetPersistentVolumeClaimInfo(namespace, pvcName)
			if err != nil {
				return false, nil, err
			}
			if pvc == nil {
				return false, nil, fmt.Errorf("PersistentVolumeClaim was not found: %q", pvcName)
			}
			pvName := pvc.Spec.VolumeName
			if pvName == "" {
				if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
					scName := v1helper.GetPersistentVolumeClaimClass(pvc)
					if len(scName) > 0 {
						class, _ := c.classInfo.GetStorageClassInfo(scName)
						if class != nil {
							if class.VolumeBindingMode == nil {
								return false, nil, fmt.Errorf("VolumeBindingMode not set for StorageClass %q", scName)
							}
							if *class.VolumeBindingMode == storagev1.VolumeBindingWaitForFirstConsumer {
								continue
							}
						}
					}
				}
				return false, nil, fmt.Errorf("PersistentVolumeClaim is not bound: %q", pvcName)
			}
			pv, err := c.pvInfo.GetPersistentVolumeInfo(pvName)
			if err != nil {
				return false, nil, err
			}
			if pv == nil {
				return false, nil, fmt.Errorf("PersistentVolume not found: %q", pvName)
			}
			for k, v := range pv.ObjectMeta.Labels {
				if k != kubeletapis.LabelZoneFailureDomain && k != kubeletapis.LabelZoneRegion {
					continue
				}
				nodeV, _ := nodeConstraints[k]
				volumeVSet, err := volumeutil.LabelZonesToSet(v)
				if err != nil {
					klog.Warningf("Failed to parse label for %q: %q. Ignoring the label. err=%v. ", k, v, err)
					continue
				}
				if !volumeVSet.Has(nodeV) {
					klog.V(10).Infof("Won't schedule pod %q onto node %q due to volume %q (mismatch on %q)", pod.Name, node.Name, pvName, k)
					return false, []algorithm.PredicateFailureReason{ErrVolumeZoneConflict}, nil
				}
			}
		}
	}
	return true, nil, nil
}
func GetResourceRequest(pod *v1.Pod) *schedulercache.Resource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := &schedulercache.Resource{}
	for _, container := range pod.Spec.Containers {
		result.Add(container.Resources.Requests)
	}
	for _, container := range pod.Spec.InitContainers {
		result.SetMaxResource(container.Resources.Requests)
	}
	return result
}
func podName(pod *v1.Pod) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pod.Namespace + "/" + pod.Name
}
func PodFitsResources(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := nodeInfo.Node()
	if node == nil {
		return false, nil, fmt.Errorf("node not found")
	}
	var predicateFails []algorithm.PredicateFailureReason
	allowedPodNumber := nodeInfo.AllowedPodNumber()
	if len(nodeInfo.Pods())+1 > allowedPodNumber {
		predicateFails = append(predicateFails, NewInsufficientResourceError(v1.ResourcePods, 1, int64(len(nodeInfo.Pods())), int64(allowedPodNumber)))
	}
	ignoredExtendedResources := sets.NewString()
	var podRequest *schedulercache.Resource
	if predicateMeta, ok := meta.(*predicateMetadata); ok {
		podRequest = predicateMeta.podRequest
		if predicateMeta.ignoredExtendedResources != nil {
			ignoredExtendedResources = predicateMeta.ignoredExtendedResources
		}
	} else {
		podRequest = GetResourceRequest(pod)
	}
	if podRequest.MilliCPU == 0 && podRequest.Memory == 0 && podRequest.EphemeralStorage == 0 && len(podRequest.ScalarResources) == 0 {
		return len(predicateFails) == 0, predicateFails, nil
	}
	allocatable := nodeInfo.AllocatableResource()
	if allocatable.MilliCPU < podRequest.MilliCPU+nodeInfo.RequestedResource().MilliCPU {
		predicateFails = append(predicateFails, NewInsufficientResourceError(v1.ResourceCPU, podRequest.MilliCPU, nodeInfo.RequestedResource().MilliCPU, allocatable.MilliCPU))
	}
	if allocatable.Memory < podRequest.Memory+nodeInfo.RequestedResource().Memory {
		predicateFails = append(predicateFails, NewInsufficientResourceError(v1.ResourceMemory, podRequest.Memory, nodeInfo.RequestedResource().Memory, allocatable.Memory))
	}
	if allocatable.EphemeralStorage < podRequest.EphemeralStorage+nodeInfo.RequestedResource().EphemeralStorage {
		predicateFails = append(predicateFails, NewInsufficientResourceError(v1.ResourceEphemeralStorage, podRequest.EphemeralStorage, nodeInfo.RequestedResource().EphemeralStorage, allocatable.EphemeralStorage))
	}
	for rName, rQuant := range podRequest.ScalarResources {
		if v1helper.IsExtendedResourceName(rName) {
			if ignoredExtendedResources.Has(string(rName)) {
				continue
			}
		}
		if allocatable.ScalarResources[rName] < rQuant+nodeInfo.RequestedResource().ScalarResources[rName] {
			predicateFails = append(predicateFails, NewInsufficientResourceError(rName, podRequest.ScalarResources[rName], nodeInfo.RequestedResource().ScalarResources[rName], allocatable.ScalarResources[rName]))
		}
	}
	if klog.V(10) {
		if len(predicateFails) == 0 {
			klog.Infof("Schedule Pod %+v on Node %+v is allowed, Node is running only %v out of %v Pods.", podName(pod), node.Name, len(nodeInfo.Pods()), allowedPodNumber)
		}
	}
	return len(predicateFails) == 0, predicateFails, nil
}
func nodeMatchesNodeSelectorTerms(node *v1.Node, nodeSelectorTerms []v1.NodeSelectorTerm) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nodeFields := map[string]string{}
	for k, f := range algorithm.NodeFieldSelectorKeys {
		nodeFields[k] = f(node)
	}
	return v1helper.MatchNodeSelectorTerms(nodeSelectorTerms, labels.Set(node.Labels), fields.Set(nodeFields))
}
func podMatchesNodeSelectorAndAffinityTerms(pod *v1.Pod, node *v1.Node) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(pod.Spec.NodeSelector) > 0 {
		selector := labels.SelectorFromSet(pod.Spec.NodeSelector)
		if !selector.Matches(labels.Set(node.Labels)) {
			return false
		}
	}
	nodeAffinityMatches := true
	affinity := pod.Spec.Affinity
	if affinity != nil && affinity.NodeAffinity != nil {
		nodeAffinity := affinity.NodeAffinity
		if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
			return true
		}
		if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution != nil {
			nodeSelectorTerms := nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms
			klog.V(10).Infof("Match for RequiredDuringSchedulingIgnoredDuringExecution node selector terms %+v", nodeSelectorTerms)
			nodeAffinityMatches = nodeAffinityMatches && nodeMatchesNodeSelectorTerms(node, nodeSelectorTerms)
		}
	}
	return nodeAffinityMatches
}
func PodMatchNodeSelector(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := nodeInfo.Node()
	if node == nil {
		return false, nil, fmt.Errorf("node not found")
	}
	if podMatchesNodeSelectorAndAffinityTerms(pod, node) {
		return true, nil, nil
	}
	return false, []algorithm.PredicateFailureReason{ErrNodeSelectorNotMatch}, nil
}
func PodFitsHost(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(pod.Spec.NodeName) == 0 {
		return true, nil, nil
	}
	node := nodeInfo.Node()
	if node == nil {
		return false, nil, fmt.Errorf("node not found")
	}
	if pod.Spec.NodeName == node.Name {
		return true, nil, nil
	}
	return false, []algorithm.PredicateFailureReason{ErrPodNotMatchHostName}, nil
}

type NodeLabelChecker struct {
	labels   []string
	presence bool
}

func NewNodeLabelPredicate(labels []string, presence bool) algorithm.FitPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	labelChecker := &NodeLabelChecker{labels: labels, presence: presence}
	return labelChecker.CheckNodeLabelPresence
}
func (n *NodeLabelChecker) CheckNodeLabelPresence(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := nodeInfo.Node()
	if node == nil {
		return false, nil, fmt.Errorf("node not found")
	}
	var exists bool
	nodeLabels := labels.Set(node.Labels)
	for _, label := range n.labels {
		exists = nodeLabels.Has(label)
		if (exists && !n.presence) || (!exists && n.presence) {
			return false, []algorithm.PredicateFailureReason{ErrNodeLabelPresenceViolated}, nil
		}
	}
	return true, nil, nil
}

type ServiceAffinity struct {
	podLister     algorithm.PodLister
	serviceLister algorithm.ServiceLister
	nodeInfo      NodeInfo
	labels        []string
}

func (s *ServiceAffinity) serviceAffinityMetadataProducer(pm *predicateMetadata) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pm.pod == nil {
		klog.Errorf("Cannot precompute service affinity, a pod is required to calculate service affinity.")
		return
	}
	pm.serviceAffinityInUse = true
	var errSvc, errList error
	pm.serviceAffinityMatchingPodServices, errSvc = s.serviceLister.GetPodServices(pm.pod)
	selector := CreateSelectorFromLabels(pm.pod.Labels)
	allMatches, errList := s.podLister.List(selector)
	if errSvc != nil || errList != nil {
		klog.Errorf("Some Error were found while precomputing svc affinity: \nservices:%v , \npods:%v", errSvc, errList)
	}
	pm.serviceAffinityMatchingPodList = FilterPodsByNamespace(allMatches, pm.pod.Namespace)
}
func NewServiceAffinityPredicate(podLister algorithm.PodLister, serviceLister algorithm.ServiceLister, nodeInfo NodeInfo, labels []string) (algorithm.FitPredicate, PredicateMetadataProducer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	affinity := &ServiceAffinity{podLister: podLister, serviceLister: serviceLister, nodeInfo: nodeInfo, labels: labels}
	return affinity.checkServiceAffinity, affinity.serviceAffinityMetadataProducer
}
func (s *ServiceAffinity) checkServiceAffinity(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var services []*v1.Service
	var pods []*v1.Pod
	if pm, ok := meta.(*predicateMetadata); ok && (pm.serviceAffinityMatchingPodList != nil || pm.serviceAffinityMatchingPodServices != nil) {
		services = pm.serviceAffinityMatchingPodServices
		pods = pm.serviceAffinityMatchingPodList
	} else {
		pm = &predicateMetadata{pod: pod}
		s.serviceAffinityMetadataProducer(pm)
		pods, services = pm.serviceAffinityMatchingPodList, pm.serviceAffinityMatchingPodServices
	}
	filteredPods := nodeInfo.FilterOutPods(pods)
	node := nodeInfo.Node()
	if node == nil {
		return false, nil, fmt.Errorf("node not found")
	}
	affinityLabels := FindLabelsInSet(s.labels, labels.Set(pod.Spec.NodeSelector))
	if len(s.labels) > len(affinityLabels) {
		if len(services) > 0 {
			if len(filteredPods) > 0 {
				nodeWithAffinityLabels, err := s.nodeInfo.GetNodeInfo(filteredPods[0].Spec.NodeName)
				if err != nil {
					return false, nil, err
				}
				AddUnsetLabelsToMap(affinityLabels, s.labels, labels.Set(nodeWithAffinityLabels.Labels))
			}
		}
	}
	if CreateSelectorFromLabels(affinityLabels).Matches(labels.Set(node.Labels)) {
		return true, nil, nil
	}
	return false, []algorithm.PredicateFailureReason{ErrServiceAffinityViolated}, nil
}
func PodFitsHostPorts(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var wantPorts []*v1.ContainerPort
	if predicateMeta, ok := meta.(*predicateMetadata); ok {
		wantPorts = predicateMeta.podPorts
	} else {
		wantPorts = schedutil.GetContainerPorts(pod)
	}
	if len(wantPorts) == 0 {
		return true, nil, nil
	}
	existingPorts := nodeInfo.UsedPorts()
	if portsConflict(existingPorts, wantPorts) {
		return false, []algorithm.PredicateFailureReason{ErrPodNotFitsHostPorts}, nil
	}
	return true, nil, nil
}
func haveOverlap(a1, a2 []string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := map[string]bool{}
	for _, val := range a1 {
		m[val] = true
	}
	for _, val := range a2 {
		if _, ok := m[val]; ok {
			return true
		}
	}
	return false
}
func GeneralPredicates(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var predicateFails []algorithm.PredicateFailureReason
	fit, reasons, err := noncriticalPredicates(pod, meta, nodeInfo)
	if err != nil {
		return false, predicateFails, err
	}
	if !fit {
		predicateFails = append(predicateFails, reasons...)
	}
	fit, reasons, err = EssentialPredicates(pod, meta, nodeInfo)
	if err != nil {
		return false, predicateFails, err
	}
	if !fit {
		predicateFails = append(predicateFails, reasons...)
	}
	return len(predicateFails) == 0, predicateFails, nil
}
func noncriticalPredicates(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var predicateFails []algorithm.PredicateFailureReason
	fit, reasons, err := PodFitsResources(pod, meta, nodeInfo)
	if err != nil {
		return false, predicateFails, err
	}
	if !fit {
		predicateFails = append(predicateFails, reasons...)
	}
	return len(predicateFails) == 0, predicateFails, nil
}
func EssentialPredicates(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var predicateFails []algorithm.PredicateFailureReason
	fit, reasons, err := PodFitsHost(pod, meta, nodeInfo)
	if err != nil {
		return false, predicateFails, err
	}
	if !fit {
		predicateFails = append(predicateFails, reasons...)
	}
	fit, reasons, err = PodFitsHostPorts(pod, meta, nodeInfo)
	if err != nil {
		return false, predicateFails, err
	}
	if !fit {
		predicateFails = append(predicateFails, reasons...)
	}
	fit, reasons, err = PodMatchNodeSelector(pod, meta, nodeInfo)
	if err != nil {
		return false, predicateFails, err
	}
	if !fit {
		predicateFails = append(predicateFails, reasons...)
	}
	return len(predicateFails) == 0, predicateFails, nil
}

type PodAffinityChecker struct {
	info      NodeInfo
	podLister algorithm.PodLister
}

func NewPodAffinityPredicate(info NodeInfo, podLister algorithm.PodLister) algorithm.FitPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	checker := &PodAffinityChecker{info: info, podLister: podLister}
	return checker.InterPodAffinityMatches
}
func (c *PodAffinityChecker) InterPodAffinityMatches(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := nodeInfo.Node()
	if node == nil {
		return false, nil, fmt.Errorf("node not found")
	}
	if failedPredicates, error := c.satisfiesExistingPodsAntiAffinity(pod, meta, nodeInfo); failedPredicates != nil {
		failedPredicates := append([]algorithm.PredicateFailureReason{ErrPodAffinityNotMatch}, failedPredicates)
		return false, failedPredicates, error
	}
	affinity := pod.Spec.Affinity
	if affinity == nil || (affinity.PodAffinity == nil && affinity.PodAntiAffinity == nil) {
		return true, nil, nil
	}
	if failedPredicates, error := c.satisfiesPodsAffinityAntiAffinity(pod, meta, nodeInfo, affinity); failedPredicates != nil {
		failedPredicates := append([]algorithm.PredicateFailureReason{ErrPodAffinityNotMatch}, failedPredicates)
		return false, failedPredicates, error
	}
	if klog.V(10) {
		klog.Infof("Schedule Pod %+v on Node %+v is allowed, pod (anti)affinity constraints satisfied", podName(pod), node.Name)
	}
	return true, nil, nil
}
func (c *PodAffinityChecker) podMatchesPodAffinityTerms(pod, targetPod *v1.Pod, nodeInfo *schedulercache.NodeInfo, terms []v1.PodAffinityTerm) (bool, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(terms) == 0 {
		return false, false, fmt.Errorf("terms array is empty")
	}
	props, err := getAffinityTermProperties(pod, terms)
	if err != nil {
		return false, false, err
	}
	if !podMatchesAllAffinityTermProperties(targetPod, props) {
		return false, false, nil
	}
	targetPodNode, err := c.info.GetNodeInfo(targetPod.Spec.NodeName)
	if err != nil {
		return false, false, err
	}
	for _, term := range terms {
		if len(term.TopologyKey) == 0 {
			return false, false, fmt.Errorf("empty topologyKey is not allowed except for PreferredDuringScheduling pod anti-affinity")
		}
		if !priorityutil.NodesHaveSameTopologyKey(nodeInfo.Node(), targetPodNode, term.TopologyKey) {
			return false, true, nil
		}
	}
	return true, true, nil
}
func GetPodAffinityTerms(podAffinity *v1.PodAffinity) (terms []v1.PodAffinityTerm) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if podAffinity != nil {
		if len(podAffinity.RequiredDuringSchedulingIgnoredDuringExecution) != 0 {
			terms = podAffinity.RequiredDuringSchedulingIgnoredDuringExecution
		}
	}
	return terms
}
func GetPodAntiAffinityTerms(podAntiAffinity *v1.PodAntiAffinity) (terms []v1.PodAffinityTerm) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if podAntiAffinity != nil {
		if len(podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution) != 0 {
			terms = podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution
		}
	}
	return terms
}
func getMatchingAntiAffinityTopologyPairsOfPod(newPod *v1.Pod, existingPod *v1.Pod, node *v1.Node) (*topologyPairsMaps, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	affinity := existingPod.Spec.Affinity
	if affinity == nil || affinity.PodAntiAffinity == nil {
		return nil, nil
	}
	topologyMaps := newTopologyPairsMaps()
	for _, term := range GetPodAntiAffinityTerms(affinity.PodAntiAffinity) {
		namespaces := priorityutil.GetNamespacesFromPodAffinityTerm(existingPod, &term)
		selector, err := metav1.LabelSelectorAsSelector(term.LabelSelector)
		if err != nil {
			return nil, err
		}
		if priorityutil.PodMatchesTermsNamespaceAndSelector(newPod, namespaces, selector) {
			if topologyValue, ok := node.Labels[term.TopologyKey]; ok {
				pair := topologyPair{key: term.TopologyKey, value: topologyValue}
				topologyMaps.addTopologyPair(pair, existingPod)
			}
		}
	}
	return topologyMaps, nil
}
func (c *PodAffinityChecker) getMatchingAntiAffinityTopologyPairsOfPods(pod *v1.Pod, existingPods []*v1.Pod) (*topologyPairsMaps, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	topologyMaps := newTopologyPairsMaps()
	for _, existingPod := range existingPods {
		existingPodNode, err := c.info.GetNodeInfo(existingPod.Spec.NodeName)
		if err != nil {
			if apierrors.IsNotFound(err) {
				klog.Errorf("Node not found, %v", existingPod.Spec.NodeName)
				continue
			}
			return nil, err
		}
		existingPodTopologyMaps, err := getMatchingAntiAffinityTopologyPairsOfPod(pod, existingPod, existingPodNode)
		if err != nil {
			return nil, err
		}
		topologyMaps.appendMaps(existingPodTopologyMaps)
	}
	return topologyMaps, nil
}
func (c *PodAffinityChecker) satisfiesExistingPodsAntiAffinity(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := nodeInfo.Node()
	if node == nil {
		return ErrExistingPodsAntiAffinityRulesNotMatch, fmt.Errorf("Node is nil")
	}
	var topologyMaps *topologyPairsMaps
	if predicateMeta, ok := meta.(*predicateMetadata); ok {
		topologyMaps = predicateMeta.topologyPairsAntiAffinityPodsMap
	} else {
		filteredPods, err := c.podLister.FilteredList(nodeInfo.Filter, labels.Everything())
		if err != nil {
			errMessage := fmt.Sprintf("Failed to get all pods, %+v", err)
			klog.Error(errMessage)
			return ErrExistingPodsAntiAffinityRulesNotMatch, errors.New(errMessage)
		}
		if topologyMaps, err = c.getMatchingAntiAffinityTopologyPairsOfPods(pod, filteredPods); err != nil {
			errMessage := fmt.Sprintf("Failed to get all terms that pod %+v matches, err: %+v", podName(pod), err)
			klog.Error(errMessage)
			return ErrExistingPodsAntiAffinityRulesNotMatch, errors.New(errMessage)
		}
	}
	for topologyKey, topologyValue := range node.Labels {
		if topologyMaps.topologyPairToPods[topologyPair{key: topologyKey, value: topologyValue}] != nil {
			klog.V(10).Infof("Cannot schedule pod %+v onto node %v", podName(pod), node.Name)
			return ErrExistingPodsAntiAffinityRulesNotMatch, nil
		}
	}
	if klog.V(10) {
		klog.Infof("Schedule Pod %+v on Node %+v is allowed, existing pods anti-affinity terms satisfied.", podName(pod), node.Name)
	}
	return nil, nil
}
func (c *PodAffinityChecker) nodeMatchesAllTopologyTerms(pod *v1.Pod, topologyPairs *topologyPairsMaps, nodeInfo *schedulercache.NodeInfo, terms []v1.PodAffinityTerm) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := nodeInfo.Node()
	for _, term := range terms {
		if topologyValue, ok := node.Labels[term.TopologyKey]; ok {
			pair := topologyPair{key: term.TopologyKey, value: topologyValue}
			if _, ok := topologyPairs.topologyPairToPods[pair]; !ok {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
func (c *PodAffinityChecker) nodeMatchesAnyTopologyTerm(pod *v1.Pod, topologyPairs *topologyPairsMaps, nodeInfo *schedulercache.NodeInfo, terms []v1.PodAffinityTerm) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := nodeInfo.Node()
	for _, term := range terms {
		if topologyValue, ok := node.Labels[term.TopologyKey]; ok {
			pair := topologyPair{key: term.TopologyKey, value: topologyValue}
			if _, ok := topologyPairs.topologyPairToPods[pair]; ok {
				return true
			}
		}
	}
	return false
}
func (c *PodAffinityChecker) satisfiesPodsAffinityAntiAffinity(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo, affinity *v1.Affinity) (algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := nodeInfo.Node()
	if node == nil {
		return ErrPodAffinityRulesNotMatch, fmt.Errorf("Node is nil")
	}
	if predicateMeta, ok := meta.(*predicateMetadata); ok {
		topologyPairsPotentialAffinityPods := predicateMeta.topologyPairsPotentialAffinityPods
		if affinityTerms := GetPodAffinityTerms(affinity.PodAffinity); len(affinityTerms) > 0 {
			matchExists := c.nodeMatchesAllTopologyTerms(pod, topologyPairsPotentialAffinityPods, nodeInfo, affinityTerms)
			if !matchExists {
				if !(len(topologyPairsPotentialAffinityPods.topologyPairToPods) == 0 && targetPodMatchesAffinityOfPod(pod, pod)) {
					klog.V(10).Infof("Cannot schedule pod %+v onto node %v, because of PodAffinity", podName(pod), node.Name)
					return ErrPodAffinityRulesNotMatch, nil
				}
			}
		}
		topologyPairsPotentialAntiAffinityPods := predicateMeta.topologyPairsPotentialAntiAffinityPods
		if antiAffinityTerms := GetPodAntiAffinityTerms(affinity.PodAntiAffinity); len(antiAffinityTerms) > 0 {
			matchExists := c.nodeMatchesAnyTopologyTerm(pod, topologyPairsPotentialAntiAffinityPods, nodeInfo, antiAffinityTerms)
			if matchExists {
				klog.V(10).Infof("Cannot schedule pod %+v onto node %v, because of PodAntiAffinity", podName(pod), node.Name)
				return ErrPodAntiAffinityRulesNotMatch, nil
			}
		}
	} else {
		filteredPods, err := c.podLister.FilteredList(nodeInfo.Filter, labels.Everything())
		if err != nil {
			return ErrPodAffinityRulesNotMatch, err
		}
		affinityTerms := GetPodAffinityTerms(affinity.PodAffinity)
		antiAffinityTerms := GetPodAntiAffinityTerms(affinity.PodAntiAffinity)
		matchFound, termsSelectorMatchFound := false, false
		for _, targetPod := range filteredPods {
			if !matchFound && len(affinityTerms) > 0 {
				affTermsMatch, termsSelectorMatch, err := c.podMatchesPodAffinityTerms(pod, targetPod, nodeInfo, affinityTerms)
				if err != nil {
					errMessage := fmt.Sprintf("Cannot schedule pod %+v onto node %v, because of PodAffinity, err: %v", podName(pod), node.Name, err)
					klog.Error(errMessage)
					return ErrPodAffinityRulesNotMatch, errors.New(errMessage)
				}
				if termsSelectorMatch {
					termsSelectorMatchFound = true
				}
				if affTermsMatch {
					matchFound = true
				}
			}
			if len(antiAffinityTerms) > 0 {
				antiAffTermsMatch, _, err := c.podMatchesPodAffinityTerms(pod, targetPod, nodeInfo, antiAffinityTerms)
				if err != nil || antiAffTermsMatch {
					klog.V(10).Infof("Cannot schedule pod %+v onto node %v, because of PodAntiAffinityTerm, err: %v", podName(pod), node.Name, err)
					return ErrPodAntiAffinityRulesNotMatch, nil
				}
			}
		}
		if !matchFound && len(affinityTerms) > 0 {
			if termsSelectorMatchFound {
				klog.V(10).Infof("Cannot schedule pod %+v onto node %v, because of PodAffinity", podName(pod), node.Name)
				return ErrPodAffinityRulesNotMatch, nil
			}
			if !targetPodMatchesAffinityOfPod(pod, pod) {
				klog.V(10).Infof("Cannot schedule pod %+v onto node %v, because of PodAffinity", podName(pod), node.Name)
				return ErrPodAffinityRulesNotMatch, nil
			}
		}
	}
	if klog.V(10) {
		klog.Infof("Schedule Pod %+v on Node %+v is allowed, pod affinity/anti-affinity constraints satisfied.", podName(pod), node.Name)
	}
	return nil, nil
}
func CheckNodeUnschedulablePredicate(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if nodeInfo == nil || nodeInfo.Node() == nil {
		return false, []algorithm.PredicateFailureReason{ErrNodeUnknownCondition}, nil
	}
	podToleratesUnschedulable := v1helper.TolerationsTolerateTaint(pod.Spec.Tolerations, &v1.Taint{Key: schedulerapi.TaintNodeUnschedulable, Effect: v1.TaintEffectNoSchedule})
	if nodeInfo.Node().Spec.Unschedulable && !podToleratesUnschedulable {
		return false, []algorithm.PredicateFailureReason{ErrNodeUnschedulable}, nil
	}
	return true, nil, nil
}
func PodToleratesNodeTaints(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if nodeInfo == nil || nodeInfo.Node() == nil {
		return false, []algorithm.PredicateFailureReason{ErrNodeUnknownCondition}, nil
	}
	return podToleratesNodeTaints(pod, nodeInfo, func(t *v1.Taint) bool {
		return t.Effect == v1.TaintEffectNoSchedule || t.Effect == v1.TaintEffectNoExecute
	})
}
func PodToleratesNodeNoExecuteTaints(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return podToleratesNodeTaints(pod, nodeInfo, func(t *v1.Taint) bool {
		return t.Effect == v1.TaintEffectNoExecute
	})
}
func podToleratesNodeTaints(pod *v1.Pod, nodeInfo *schedulercache.NodeInfo, filter func(t *v1.Taint) bool) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	taints, err := nodeInfo.Taints()
	if err != nil {
		return false, nil, err
	}
	if v1helper.TolerationsTolerateTaintsWithFilter(pod.Spec.Tolerations, taints, filter) {
		return true, nil, nil
	}
	return false, []algorithm.PredicateFailureReason{ErrTaintsTolerationsNotMatch}, nil
}
func isPodBestEffort(pod *v1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v1qos.GetPodQOS(pod) == v1.PodQOSBestEffort
}
func CheckNodeMemoryPressurePredicate(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var podBestEffort bool
	if predicateMeta, ok := meta.(*predicateMetadata); ok {
		podBestEffort = predicateMeta.podBestEffort
	} else {
		podBestEffort = isPodBestEffort(pod)
	}
	if !podBestEffort {
		return true, nil, nil
	}
	if nodeInfo.MemoryPressureCondition() == v1.ConditionTrue {
		return false, []algorithm.PredicateFailureReason{ErrNodeUnderMemoryPressure}, nil
	}
	return true, nil, nil
}
func CheckNodeDiskPressurePredicate(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if nodeInfo.DiskPressureCondition() == v1.ConditionTrue {
		return false, []algorithm.PredicateFailureReason{ErrNodeUnderDiskPressure}, nil
	}
	return true, nil, nil
}
func CheckNodePIDPressurePredicate(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if nodeInfo.PIDPressureCondition() == v1.ConditionTrue {
		return false, []algorithm.PredicateFailureReason{ErrNodeUnderPIDPressure}, nil
	}
	return true, nil, nil
}
func CheckNodeConditionPredicate(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	reasons := []algorithm.PredicateFailureReason{}
	if nodeInfo == nil || nodeInfo.Node() == nil {
		return false, []algorithm.PredicateFailureReason{ErrNodeUnknownCondition}, nil
	}
	node := nodeInfo.Node()
	for _, cond := range node.Status.Conditions {
		if cond.Type == v1.NodeReady && cond.Status != v1.ConditionTrue {
			reasons = append(reasons, ErrNodeNotReady)
		} else if cond.Type == v1.NodeOutOfDisk && cond.Status != v1.ConditionFalse {
			reasons = append(reasons, ErrNodeOutOfDisk)
		} else if cond.Type == v1.NodeNetworkUnavailable && cond.Status != v1.ConditionFalse {
			reasons = append(reasons, ErrNodeNetworkUnavailable)
		}
	}
	if node.Spec.Unschedulable {
		reasons = append(reasons, ErrNodeUnschedulable)
	}
	return len(reasons) == 0, reasons, nil
}

type VolumeBindingChecker struct{ binder *volumebinder.VolumeBinder }

func NewVolumeBindingPredicate(binder *volumebinder.VolumeBinder) algorithm.FitPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &VolumeBindingChecker{binder: binder}
	return c.predicate
}
func (c *VolumeBindingChecker) predicate(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		return true, nil, nil
	}
	node := nodeInfo.Node()
	if node == nil {
		return false, nil, fmt.Errorf("node not found")
	}
	unboundSatisfied, boundSatisfied, err := c.binder.Binder.FindPodVolumes(pod, node)
	if err != nil {
		return false, nil, err
	}
	failReasons := []algorithm.PredicateFailureReason{}
	if !boundSatisfied {
		klog.V(5).Infof("Bound PVs not satisfied for pod %v/%v, node %q", pod.Namespace, pod.Name, node.Name)
		failReasons = append(failReasons, ErrVolumeNodeConflict)
	}
	if !unboundSatisfied {
		klog.V(5).Infof("Couldn't find matching PVs for pod %v/%v, node %q", pod.Namespace, pod.Name, node.Name)
		failReasons = append(failReasons, ErrVolumeBindConflict)
	}
	if len(failReasons) > 0 {
		return false, failReasons, nil
	}
	klog.V(5).Infof("All PVCs found matches for pod %v/%v, node %q", pod.Namespace, pod.Name, node.Name)
	return true, nil, nil
}

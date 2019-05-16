package persistentvolume

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/tools/cache"
	v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
	"k8s.io/kubernetes/pkg/features"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	gotime "time"
)

type persistentVolumeOrderedIndex struct{ store cache.Indexer }

func newPersistentVolumeOrderedIndex() persistentVolumeOrderedIndex {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return persistentVolumeOrderedIndex{cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{"accessmodes": accessModesIndexFunc})}
}
func accessModesIndexFunc(obj interface{}) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pv, ok := obj.(*v1.PersistentVolume); ok {
		modes := v1helper.GetAccessModesAsString(pv.Spec.AccessModes)
		return []string{modes}, nil
	}
	return []string{""}, fmt.Errorf("object is not a persistent volume: %v", obj)
}
func (pvIndex *persistentVolumeOrderedIndex) listByAccessModes(modes []v1.PersistentVolumeAccessMode) ([]*v1.PersistentVolume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pv := &v1.PersistentVolume{Spec: v1.PersistentVolumeSpec{AccessModes: modes}}
	objs, err := pvIndex.store.Index("accessmodes", pv)
	if err != nil {
		return nil, err
	}
	volumes := make([]*v1.PersistentVolume, len(objs))
	for i, obj := range objs {
		volumes[i] = obj.(*v1.PersistentVolume)
	}
	return volumes, nil
}
func (pvIndex *persistentVolumeOrderedIndex) findByClaim(claim *v1.PersistentVolumeClaim, delayBinding bool) (*v1.PersistentVolume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allPossibleModes := pvIndex.allPossibleMatchingAccessModes(claim.Spec.AccessModes)
	for _, modes := range allPossibleModes {
		volumes, err := pvIndex.listByAccessModes(modes)
		if err != nil {
			return nil, err
		}
		bestVol, err := findMatchingVolume(claim, volumes, nil, nil, delayBinding)
		if err != nil {
			return nil, err
		}
		if bestVol != nil {
			return bestVol, nil
		}
	}
	return nil, nil
}
func findMatchingVolume(claim *v1.PersistentVolumeClaim, volumes []*v1.PersistentVolume, node *v1.Node, excludedVolumes map[string]*v1.PersistentVolume, delayBinding bool) (*v1.PersistentVolume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var smallestVolume *v1.PersistentVolume
	var smallestVolumeQty resource.Quantity
	requestedQty := claim.Spec.Resources.Requests[v1.ResourceName(v1.ResourceStorage)]
	requestedClass := v1helper.GetPersistentVolumeClaimClass(claim)
	var selector labels.Selector
	if claim.Spec.Selector != nil {
		internalSelector, err := metav1.LabelSelectorAsSelector(claim.Spec.Selector)
		if err != nil {
			return nil, fmt.Errorf("error creating internal label selector for claim: %v: %v", claimToClaimKey(claim), err)
		}
		selector = internalSelector
	}
	for _, volume := range volumes {
		if _, ok := excludedVolumes[volume.Name]; ok {
			continue
		}
		volumeQty := volume.Spec.Capacity[v1.ResourceStorage]
		isMismatch, err := checkVolumeModeMismatches(&claim.Spec, &volume.Spec)
		if err != nil {
			return nil, fmt.Errorf("error checking if volumeMode was a mismatch: %v", err)
		}
		if isMismatch {
			continue
		}
		if utilfeature.DefaultFeatureGate.Enabled(features.StorageObjectInUseProtection) {
			if volume.ObjectMeta.DeletionTimestamp != nil {
				continue
			}
		}
		nodeAffinityValid := true
		if node != nil {
			err := volumeutil.CheckNodeAffinity(volume, node.Labels)
			if err != nil {
				nodeAffinityValid = false
			}
		}
		if isVolumeBoundToClaim(volume, claim) {
			if volumeQty.Cmp(requestedQty) < 0 {
				continue
			}
			if !nodeAffinityValid {
				return nil, nil
			}
			return volume, nil
		}
		if node == nil && delayBinding {
			continue
		}
		if volume.Status.Phase != v1.VolumeAvailable {
			continue
		} else if volume.Spec.ClaimRef != nil {
			continue
		} else if selector != nil && !selector.Matches(labels.Set(volume.Labels)) {
			continue
		}
		if v1helper.GetPersistentVolumeClass(volume) != requestedClass {
			continue
		}
		if !nodeAffinityValid {
			continue
		}
		if node != nil {
			if !checkAccessModes(claim, volume) {
				continue
			}
		}
		if volumeQty.Cmp(requestedQty) >= 0 {
			if smallestVolume == nil || smallestVolumeQty.Cmp(volumeQty) > 0 {
				smallestVolume = volume
				smallestVolumeQty = volumeQty
			}
		}
	}
	if smallestVolume != nil {
		return smallestVolume, nil
	}
	return nil, nil
}
func checkVolumeModeMismatches(pvcSpec *v1.PersistentVolumeClaimSpec, pvSpec *v1.PersistentVolumeSpec) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
		return false, nil
	}
	requestedVolumeMode := v1.PersistentVolumeFilesystem
	if pvcSpec.VolumeMode != nil {
		requestedVolumeMode = *pvcSpec.VolumeMode
	}
	pvVolumeMode := v1.PersistentVolumeFilesystem
	if pvSpec.VolumeMode != nil {
		pvVolumeMode = *pvSpec.VolumeMode
	}
	return requestedVolumeMode != pvVolumeMode, nil
}
func (pvIndex *persistentVolumeOrderedIndex) findBestMatchForClaim(claim *v1.PersistentVolumeClaim, delayBinding bool) (*v1.PersistentVolume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pvIndex.findByClaim(claim, delayBinding)
}
func (pvIndex *persistentVolumeOrderedIndex) allPossibleMatchingAccessModes(requestedModes []v1.PersistentVolumeAccessMode) [][]v1.PersistentVolumeAccessMode {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	matchedModes := [][]v1.PersistentVolumeAccessMode{}
	keys := pvIndex.store.ListIndexFuncValues("accessmodes")
	for _, key := range keys {
		indexedModes := v1helper.GetAccessModesFromString(key)
		if volumeutil.AccessModesContainedInAll(indexedModes, requestedModes) {
			matchedModes = append(matchedModes, indexedModes)
		}
	}
	sort.Sort(byAccessModes{matchedModes})
	return matchedModes
}

type byAccessModes struct {
	modes [][]v1.PersistentVolumeAccessMode
}

func (c byAccessModes) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(c.modes[i]) < len(c.modes[j])
}
func (c byAccessModes) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.modes[i], c.modes[j] = c.modes[j], c.modes[i]
}
func (c byAccessModes) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(c.modes)
}
func claimToClaimKey(claim *v1.PersistentVolumeClaim) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s/%s", claim.Namespace, claim.Name)
}
func claimrefToClaimKey(claimref *v1.ObjectReference) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s/%s", claimref.Namespace, claimref.Name)
}
func checkAccessModes(claim *v1.PersistentVolumeClaim, volume *v1.PersistentVolume) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pvModesMap := map[v1.PersistentVolumeAccessMode]bool{}
	for _, mode := range volume.Spec.AccessModes {
		pvModesMap[mode] = true
	}
	for _, mode := range claim.Spec.AccessModes {
		_, ok := pvModesMap[mode]
		if !ok {
			return false
		}
	}
	return true
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

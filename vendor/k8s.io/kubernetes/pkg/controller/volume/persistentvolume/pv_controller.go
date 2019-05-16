package persistentvolume

import (
	"fmt"
	"k8s.io/api/core/v1"
	storage "k8s.io/api/storage/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	corelisters "k8s.io/client-go/listers/core/v1"
	storagelisters "k8s.io/client-go/listers/storage/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	ref "k8s.io/client-go/tools/reference"
	"k8s.io/client-go/util/workqueue"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
	"k8s.io/kubernetes/pkg/controller/volume/events"
	"k8s.io/kubernetes/pkg/controller/volume/persistentvolume/metrics"
	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/util/goroutinemap"
	"k8s.io/kubernetes/pkg/util/goroutinemap/exponentialbackoff"
	vol "k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util"
	"k8s.io/kubernetes/pkg/volume/util/recyclerclient"
	"reflect"
	"strings"
	"time"
)

const annBindCompleted = "pv.kubernetes.io/bind-completed"
const annBoundByController = "pv.kubernetes.io/bound-by-controller"
const annDynamicallyProvisioned = "pv.kubernetes.io/provisioned-by"
const annStorageProvisioner = "volume.beta.kubernetes.io/storage-provisioner"
const annSelectedNode = "volume.kubernetes.io/selected-node"
const notSupportedProvisioner = "kubernetes.io/no-provisioner"
const CloudVolumeCreatedForClaimNamespaceTag = "kubernetes.io/created-for/pvc/namespace"
const CloudVolumeCreatedForClaimNameTag = "kubernetes.io/created-for/pvc/name"
const CloudVolumeCreatedForVolumeNameTag = "kubernetes.io/created-for/pv/name"
const createProvisionedPVRetryCount = 5
const createProvisionedPVInterval = 10 * time.Second

type PersistentVolumeController struct {
	volumeLister                  corelisters.PersistentVolumeLister
	volumeListerSynced            cache.InformerSynced
	claimLister                   corelisters.PersistentVolumeClaimLister
	claimListerSynced             cache.InformerSynced
	classLister                   storagelisters.StorageClassLister
	classListerSynced             cache.InformerSynced
	podLister                     corelisters.PodLister
	podListerSynced               cache.InformerSynced
	NodeLister                    corelisters.NodeLister
	NodeListerSynced              cache.InformerSynced
	kubeClient                    clientset.Interface
	eventRecorder                 record.EventRecorder
	cloud                         cloudprovider.Interface
	volumePluginMgr               vol.VolumePluginMgr
	enableDynamicProvisioning     bool
	clusterName                   string
	resyncPeriod                  time.Duration
	volumes                       persistentVolumeOrderedIndex
	claims                        cache.Store
	claimQueue                    *workqueue.Type
	volumeQueue                   *workqueue.Type
	runningOperations             goroutinemap.GoRoutineMap
	preOperationHook              func(operationName string)
	createProvisionedPVRetryCount int
	createProvisionedPVInterval   time.Duration
}

func (ctrl *PersistentVolumeController) syncClaim(claim *v1.PersistentVolumeClaim) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("synchronizing PersistentVolumeClaim[%s]: %s", claimToClaimKey(claim), getClaimStatusForLogging(claim))
	if !metav1.HasAnnotation(claim.ObjectMeta, annBindCompleted) {
		return ctrl.syncUnboundClaim(claim)
	} else {
		return ctrl.syncBoundClaim(claim)
	}
}
func checkVolumeSatisfyClaim(volume *v1.PersistentVolume, claim *v1.PersistentVolumeClaim) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requestedQty := claim.Spec.Resources.Requests[v1.ResourceName(v1.ResourceStorage)]
	requestedSize := requestedQty.Value()
	if utilfeature.DefaultFeatureGate.Enabled(features.StorageObjectInUseProtection) {
		if volume.ObjectMeta.DeletionTimestamp != nil {
			return fmt.Errorf("the volume is marked for deletion")
		}
	}
	volumeQty := volume.Spec.Capacity[v1.ResourceStorage]
	volumeSize := volumeQty.Value()
	if volumeSize < requestedSize {
		return fmt.Errorf("requested PV is too small")
	}
	requestedClass := v1helper.GetPersistentVolumeClaimClass(claim)
	if v1helper.GetPersistentVolumeClass(volume) != requestedClass {
		return fmt.Errorf("storageClassName does not match")
	}
	isMismatch, err := checkVolumeModeMismatches(&claim.Spec, &volume.Spec)
	if err != nil {
		return fmt.Errorf("error checking volumeMode: %v", err)
	}
	if isMismatch {
		return fmt.Errorf("incompatible volumeMode")
	}
	if !checkAccessModes(claim, volume) {
		return fmt.Errorf("incompatible accessMode")
	}
	return nil
}
func (ctrl *PersistentVolumeController) shouldDelayBinding(claim *v1.PersistentVolumeClaim) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		return false, nil
	}
	if _, ok := claim.Annotations[annSelectedNode]; ok {
		return false, nil
	}
	className := v1helper.GetPersistentVolumeClaimClass(claim)
	if className == "" {
		return false, nil
	}
	class, err := ctrl.classLister.Get(className)
	if err != nil {
		return false, nil
	}
	if class.VolumeBindingMode == nil {
		return false, fmt.Errorf("VolumeBindingMode not set for StorageClass %q", className)
	}
	return *class.VolumeBindingMode == storage.VolumeBindingWaitForFirstConsumer, nil
}
func (ctrl *PersistentVolumeController) syncUnboundClaim(claim *v1.PersistentVolumeClaim) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if claim.Spec.VolumeName == "" {
		delayBinding, err := ctrl.shouldDelayBinding(claim)
		if err != nil {
			return err
		}
		volume, err := ctrl.volumes.findBestMatchForClaim(claim, delayBinding)
		if err != nil {
			klog.V(2).Infof("synchronizing unbound PersistentVolumeClaim[%s]: Error finding PV for claim: %v", claimToClaimKey(claim), err)
			return fmt.Errorf("Error finding PV for claim %q: %v", claimToClaimKey(claim), err)
		}
		if volume == nil {
			klog.V(4).Infof("synchronizing unbound PersistentVolumeClaim[%s]: no volume found", claimToClaimKey(claim))
			switch {
			case delayBinding:
				ctrl.eventRecorder.Event(claim, v1.EventTypeNormal, events.WaitForFirstConsumer, "waiting for first consumer to be created before binding")
			case v1helper.GetPersistentVolumeClaimClass(claim) != "":
				if err = ctrl.provisionClaim(claim); err != nil {
					return err
				}
				return nil
			default:
				ctrl.eventRecorder.Event(claim, v1.EventTypeNormal, events.FailedBinding, "no persistent volumes available for this claim and no storage class is set")
			}
			if _, err = ctrl.updateClaimStatus(claim, v1.ClaimPending, nil); err != nil {
				return err
			}
			return nil
		} else {
			klog.V(4).Infof("synchronizing unbound PersistentVolumeClaim[%s]: volume %q found: %s", claimToClaimKey(claim), volume.Name, getVolumeStatusForLogging(volume))
			if err = ctrl.bind(volume, claim); err != nil {
				return err
			}
			return nil
		}
	} else {
		klog.V(4).Infof("synchronizing unbound PersistentVolumeClaim[%s]: volume %q requested", claimToClaimKey(claim), claim.Spec.VolumeName)
		obj, found, err := ctrl.volumes.store.GetByKey(claim.Spec.VolumeName)
		if err != nil {
			return err
		}
		if !found {
			klog.V(4).Infof("synchronizing unbound PersistentVolumeClaim[%s]: volume %q requested and not found, will try again next time", claimToClaimKey(claim), claim.Spec.VolumeName)
			if _, err = ctrl.updateClaimStatus(claim, v1.ClaimPending, nil); err != nil {
				return err
			}
			return nil
		} else {
			volume, ok := obj.(*v1.PersistentVolume)
			if !ok {
				return fmt.Errorf("Cannot convert object from volume cache to volume %q!?: %+v", claim.Spec.VolumeName, obj)
			}
			klog.V(4).Infof("synchronizing unbound PersistentVolumeClaim[%s]: volume %q requested and found: %s", claimToClaimKey(claim), claim.Spec.VolumeName, getVolumeStatusForLogging(volume))
			if volume.Spec.ClaimRef == nil {
				klog.V(4).Infof("synchronizing unbound PersistentVolumeClaim[%s]: volume is unbound, binding", claimToClaimKey(claim))
				if err = checkVolumeSatisfyClaim(volume, claim); err != nil {
					klog.V(4).Infof("Can't bind the claim to volume %q: %v", volume.Name, err)
					msg := fmt.Sprintf("Cannot bind to requested volume %q: %s", volume.Name, err)
					ctrl.eventRecorder.Event(claim, v1.EventTypeWarning, events.VolumeMismatch, msg)
					if _, err = ctrl.updateClaimStatus(claim, v1.ClaimPending, nil); err != nil {
						return err
					}
				} else if err = ctrl.bind(volume, claim); err != nil {
					return err
				}
				return nil
			} else if isVolumeBoundToClaim(volume, claim) {
				klog.V(4).Infof("synchronizing unbound PersistentVolumeClaim[%s]: volume already bound, finishing the binding", claimToClaimKey(claim))
				if err = ctrl.bind(volume, claim); err != nil {
					return err
				}
				return nil
			} else {
				if !metav1.HasAnnotation(claim.ObjectMeta, annBoundByController) {
					klog.V(4).Infof("synchronizing unbound PersistentVolumeClaim[%s]: volume already bound to different claim by user, will retry later", claimToClaimKey(claim))
					if _, err = ctrl.updateClaimStatus(claim, v1.ClaimPending, nil); err != nil {
						return err
					}
					return nil
				} else {
					klog.V(4).Infof("synchronizing unbound PersistentVolumeClaim[%s]: volume already bound to different claim %q by controller, THIS SHOULD NEVER HAPPEN", claimToClaimKey(claim), claimrefToClaimKey(volume.Spec.ClaimRef))
					return fmt.Errorf("Invalid binding of claim %q to volume %q: volume already claimed by %q", claimToClaimKey(claim), claim.Spec.VolumeName, claimrefToClaimKey(volume.Spec.ClaimRef))
				}
			}
		}
	}
}
func (ctrl *PersistentVolumeController) syncBoundClaim(claim *v1.PersistentVolumeClaim) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if claim.Spec.VolumeName == "" {
		if _, err := ctrl.updateClaimStatusWithEvent(claim, v1.ClaimLost, nil, v1.EventTypeWarning, "ClaimLost", "Bound claim has lost reference to PersistentVolume. Data on the volume is lost!"); err != nil {
			return err
		}
		return nil
	}
	obj, found, err := ctrl.volumes.store.GetByKey(claim.Spec.VolumeName)
	if err != nil {
		return err
	}
	if !found {
		if _, err = ctrl.updateClaimStatusWithEvent(claim, v1.ClaimLost, nil, v1.EventTypeWarning, "ClaimLost", "Bound claim has lost its PersistentVolume. Data on the volume is lost!"); err != nil {
			return err
		}
		return nil
	} else {
		volume, ok := obj.(*v1.PersistentVolume)
		if !ok {
			return fmt.Errorf("Cannot convert object from volume cache to volume %q!?: %#v", claim.Spec.VolumeName, obj)
		}
		klog.V(4).Infof("synchronizing bound PersistentVolumeClaim[%s]: volume %q found: %s", claimToClaimKey(claim), claim.Spec.VolumeName, getVolumeStatusForLogging(volume))
		if volume.Spec.ClaimRef == nil {
			klog.V(4).Infof("synchronizing bound PersistentVolumeClaim[%s]: volume is unbound, fixing", claimToClaimKey(claim))
			if err = ctrl.bind(volume, claim); err != nil {
				return err
			}
			return nil
		} else if volume.Spec.ClaimRef.UID == claim.UID {
			klog.V(4).Infof("synchronizing bound PersistentVolumeClaim[%s]: claim is already correctly bound", claimToClaimKey(claim))
			if err = ctrl.bind(volume, claim); err != nil {
				return err
			}
			return nil
		} else {
			if _, err = ctrl.updateClaimStatusWithEvent(claim, v1.ClaimLost, nil, v1.EventTypeWarning, "ClaimMisbound", "Two claims are bound to the same volume, this one is bound incorrectly"); err != nil {
				return err
			}
			return nil
		}
	}
}
func (ctrl *PersistentVolumeController) syncVolume(volume *v1.PersistentVolume) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("synchronizing PersistentVolume[%s]: %s", volume.Name, getVolumeStatusForLogging(volume))
	if volume.Spec.ClaimRef == nil {
		klog.V(4).Infof("synchronizing PersistentVolume[%s]: volume is unused", volume.Name)
		if _, err := ctrl.updateVolumePhase(volume, v1.VolumeAvailable, ""); err != nil {
			return err
		}
		return nil
	} else {
		if volume.Spec.ClaimRef.UID == "" {
			klog.V(4).Infof("synchronizing PersistentVolume[%s]: volume is pre-bound to claim %s", volume.Name, claimrefToClaimKey(volume.Spec.ClaimRef))
			if _, err := ctrl.updateVolumePhase(volume, v1.VolumeAvailable, ""); err != nil {
				return err
			}
			return nil
		}
		klog.V(4).Infof("synchronizing PersistentVolume[%s]: volume is bound to claim %s", volume.Name, claimrefToClaimKey(volume.Spec.ClaimRef))
		var claim *v1.PersistentVolumeClaim
		claimName := claimrefToClaimKey(volume.Spec.ClaimRef)
		obj, found, err := ctrl.claims.GetByKey(claimName)
		if err != nil {
			return err
		}
		if !found && metav1.HasAnnotation(volume.ObjectMeta, annBoundByController) {
			if volume.Status.Phase != v1.VolumeReleased && volume.Status.Phase != v1.VolumeFailed {
				obj, err = ctrl.claimLister.PersistentVolumeClaims(volume.Spec.ClaimRef.Namespace).Get(volume.Spec.ClaimRef.Name)
				if err != nil && !apierrs.IsNotFound(err) {
					return err
				}
				found = !apierrs.IsNotFound(err)
				if !found {
					obj, err = ctrl.kubeClient.CoreV1().PersistentVolumeClaims(volume.Spec.ClaimRef.Namespace).Get(volume.Spec.ClaimRef.Name, metav1.GetOptions{})
					if err != nil && !apierrs.IsNotFound(err) {
						return err
					}
					found = !apierrs.IsNotFound(err)
				}
			}
		}
		if !found {
			klog.V(4).Infof("synchronizing PersistentVolume[%s]: claim %s not found", volume.Name, claimrefToClaimKey(volume.Spec.ClaimRef))
		} else {
			var ok bool
			claim, ok = obj.(*v1.PersistentVolumeClaim)
			if !ok {
				return fmt.Errorf("Cannot convert object from volume cache to volume %q!?: %#v", claim.Spec.VolumeName, obj)
			}
			klog.V(4).Infof("synchronizing PersistentVolume[%s]: claim %s found: %s", volume.Name, claimrefToClaimKey(volume.Spec.ClaimRef), getClaimStatusForLogging(claim))
		}
		if claim != nil && claim.UID != volume.Spec.ClaimRef.UID {
			klog.V(4).Infof("synchronizing PersistentVolume[%s]: claim %s has different UID, the old one must have been deleted", volume.Name, claimrefToClaimKey(volume.Spec.ClaimRef))
			claim = nil
		}
		if claim == nil {
			if volume.Status.Phase != v1.VolumeReleased && volume.Status.Phase != v1.VolumeFailed {
				klog.V(2).Infof("volume %q is released and reclaim policy %q will be executed", volume.Name, volume.Spec.PersistentVolumeReclaimPolicy)
				if volume, err = ctrl.updateVolumePhase(volume, v1.VolumeReleased, ""); err != nil {
					return err
				}
			}
			if err = ctrl.reclaimVolume(volume); err != nil {
				return err
			}
			return nil
		} else if claim.Spec.VolumeName == "" {
			if isMismatch, err := checkVolumeModeMismatches(&claim.Spec, &volume.Spec); err != nil || isMismatch {
				volumeMsg := fmt.Sprintf("Cannot bind PersistentVolume to requested PersistentVolumeClaim %q due to incompatible volumeMode.", claim.Name)
				ctrl.eventRecorder.Event(volume, v1.EventTypeWarning, events.VolumeMismatch, volumeMsg)
				claimMsg := fmt.Sprintf("Cannot bind PersistentVolume %q to requested PersistentVolumeClaim due to incompatible volumeMode.", volume.Name)
				ctrl.eventRecorder.Event(claim, v1.EventTypeWarning, events.VolumeMismatch, claimMsg)
				return nil
			}
			if metav1.HasAnnotation(volume.ObjectMeta, annBoundByController) {
				klog.V(4).Infof("synchronizing PersistentVolume[%s]: volume not bound yet, waiting for syncClaim to fix it", volume.Name)
			} else {
				klog.V(4).Infof("synchronizing PersistentVolume[%s]: volume was bound and got unbound (by user?), waiting for syncClaim to fix it", volume.Name)
			}
			ctrl.claimQueue.Add(claimToClaimKey(claim))
			return nil
		} else if claim.Spec.VolumeName == volume.Name {
			klog.V(4).Infof("synchronizing PersistentVolume[%s]: all is bound", volume.Name)
			if _, err = ctrl.updateVolumePhase(volume, v1.VolumeBound, ""); err != nil {
				return err
			}
			return nil
		} else {
			if metav1.HasAnnotation(volume.ObjectMeta, annDynamicallyProvisioned) && volume.Spec.PersistentVolumeReclaimPolicy == v1.PersistentVolumeReclaimDelete {
				if volume.Status.Phase != v1.VolumeReleased && volume.Status.Phase != v1.VolumeFailed {
					klog.V(2).Infof("dynamically volume %q is released and it will be deleted", volume.Name)
					if volume, err = ctrl.updateVolumePhase(volume, v1.VolumeReleased, ""); err != nil {
						return err
					}
				}
				if err = ctrl.reclaimVolume(volume); err != nil {
					return err
				}
				return nil
			} else {
				if metav1.HasAnnotation(volume.ObjectMeta, annBoundByController) {
					klog.V(4).Infof("synchronizing PersistentVolume[%s]: volume is bound by controller to a claim that is bound to another volume, unbinding", volume.Name)
					if err = ctrl.unbindVolume(volume); err != nil {
						return err
					}
					return nil
				} else {
					klog.V(4).Infof("synchronizing PersistentVolume[%s]: volume is bound by user to a claim that is bound to another volume, waiting for the claim to get unbound", volume.Name)
					if err = ctrl.unbindVolume(volume); err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
}
func (ctrl *PersistentVolumeController) updateClaimStatus(claim *v1.PersistentVolumeClaim, phase v1.PersistentVolumeClaimPhase, volume *v1.PersistentVolume) (*v1.PersistentVolumeClaim, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("updating PersistentVolumeClaim[%s] status: set phase %s", claimToClaimKey(claim), phase)
	dirty := false
	claimClone := claim.DeepCopy()
	if claim.Status.Phase != phase {
		claimClone.Status.Phase = phase
		dirty = true
	}
	if volume == nil {
		if claim.Status.AccessModes != nil {
			claimClone.Status.AccessModes = nil
			dirty = true
		}
		if claim.Status.Capacity != nil {
			claimClone.Status.Capacity = nil
			dirty = true
		}
	} else {
		if !reflect.DeepEqual(claim.Status.AccessModes, volume.Spec.AccessModes) {
			claimClone.Status.AccessModes = volume.Spec.AccessModes
			dirty = true
		}
		if claim.Status.Phase != phase {
			volumeCap, ok := volume.Spec.Capacity[v1.ResourceStorage]
			if !ok {
				return nil, fmt.Errorf("PersistentVolume %q is without a storage capacity", volume.Name)
			}
			claimCap, ok := claim.Status.Capacity[v1.ResourceStorage]
			if !ok || volumeCap.Cmp(claimCap) != 0 {
				claimClone.Status.Capacity = volume.Spec.Capacity
				dirty = true
			}
		}
	}
	if !dirty {
		klog.V(4).Infof("updating PersistentVolumeClaim[%s] status: phase %s already set", claimToClaimKey(claim), phase)
		return claim, nil
	}
	newClaim, err := ctrl.kubeClient.CoreV1().PersistentVolumeClaims(claimClone.Namespace).UpdateStatus(claimClone)
	if err != nil {
		klog.V(4).Infof("updating PersistentVolumeClaim[%s] status: set phase %s failed: %v", claimToClaimKey(claim), phase, err)
		return newClaim, err
	}
	_, err = ctrl.storeClaimUpdate(newClaim)
	if err != nil {
		klog.V(4).Infof("updating PersistentVolumeClaim[%s] status: cannot update internal cache: %v", claimToClaimKey(claim), err)
		return newClaim, err
	}
	klog.V(2).Infof("claim %q entered phase %q", claimToClaimKey(claim), phase)
	return newClaim, nil
}
func (ctrl *PersistentVolumeController) updateClaimStatusWithEvent(claim *v1.PersistentVolumeClaim, phase v1.PersistentVolumeClaimPhase, volume *v1.PersistentVolume, eventtype, reason, message string) (*v1.PersistentVolumeClaim, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("updating updateClaimStatusWithEvent[%s]: set phase %s", claimToClaimKey(claim), phase)
	if claim.Status.Phase == phase {
		klog.V(4).Infof("updating updateClaimStatusWithEvent[%s]: phase %s already set", claimToClaimKey(claim), phase)
		return claim, nil
	}
	newClaim, err := ctrl.updateClaimStatus(claim, phase, volume)
	if err != nil {
		return nil, err
	}
	klog.V(3).Infof("claim %q changed status to %q: %s", claimToClaimKey(claim), phase, message)
	ctrl.eventRecorder.Event(newClaim, eventtype, reason, message)
	return newClaim, nil
}
func (ctrl *PersistentVolumeController) updateVolumePhase(volume *v1.PersistentVolume, phase v1.PersistentVolumePhase, message string) (*v1.PersistentVolume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("updating PersistentVolume[%s]: set phase %s", volume.Name, phase)
	if volume.Status.Phase == phase {
		klog.V(4).Infof("updating PersistentVolume[%s]: phase %s already set", volume.Name, phase)
		return volume, nil
	}
	volumeClone := volume.DeepCopy()
	volumeClone.Status.Phase = phase
	volumeClone.Status.Message = message
	newVol, err := ctrl.kubeClient.CoreV1().PersistentVolumes().UpdateStatus(volumeClone)
	if err != nil {
		klog.V(4).Infof("updating PersistentVolume[%s]: set phase %s failed: %v", volume.Name, phase, err)
		return newVol, err
	}
	_, err = ctrl.storeVolumeUpdate(newVol)
	if err != nil {
		klog.V(4).Infof("updating PersistentVolume[%s]: cannot update internal cache: %v", volume.Name, err)
		return newVol, err
	}
	klog.V(2).Infof("volume %q entered phase %q", volume.Name, phase)
	return newVol, err
}
func (ctrl *PersistentVolumeController) updateVolumePhaseWithEvent(volume *v1.PersistentVolume, phase v1.PersistentVolumePhase, eventtype, reason, message string) (*v1.PersistentVolume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("updating updateVolumePhaseWithEvent[%s]: set phase %s", volume.Name, phase)
	if volume.Status.Phase == phase {
		klog.V(4).Infof("updating updateVolumePhaseWithEvent[%s]: phase %s already set", volume.Name, phase)
		return volume, nil
	}
	newVol, err := ctrl.updateVolumePhase(volume, phase, message)
	if err != nil {
		return nil, err
	}
	klog.V(3).Infof("volume %q changed status to %q: %s", volume.Name, phase, message)
	ctrl.eventRecorder.Event(newVol, eventtype, reason, message)
	return newVol, nil
}
func (ctrl *PersistentVolumeController) bindVolumeToClaim(volume *v1.PersistentVolume, claim *v1.PersistentVolumeClaim) (*v1.PersistentVolume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("updating PersistentVolume[%s]: binding to %q", volume.Name, claimToClaimKey(claim))
	volumeClone, dirty, err := ctrl.getBindVolumeToClaim(volume, claim)
	if err != nil {
		return nil, err
	}
	if dirty {
		return ctrl.updateBindVolumeToClaim(volumeClone, claim, true)
	}
	klog.V(4).Infof("updating PersistentVolume[%s]: already bound to %q", volume.Name, claimToClaimKey(claim))
	return volume, nil
}
func (ctrl *PersistentVolumeController) updateBindVolumeToClaim(volumeClone *v1.PersistentVolume, claim *v1.PersistentVolumeClaim, updateCache bool) (*v1.PersistentVolume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(2).Infof("claim %q bound to volume %q", claimToClaimKey(claim), volumeClone.Name)
	newVol, err := ctrl.kubeClient.CoreV1().PersistentVolumes().Update(volumeClone)
	if err != nil {
		klog.V(4).Infof("updating PersistentVolume[%s]: binding to %q failed: %v", volumeClone.Name, claimToClaimKey(claim), err)
		return newVol, err
	}
	if updateCache {
		_, err = ctrl.storeVolumeUpdate(newVol)
		if err != nil {
			klog.V(4).Infof("updating PersistentVolume[%s]: cannot update internal cache: %v", volumeClone.Name, err)
			return newVol, err
		}
	}
	klog.V(4).Infof("updating PersistentVolume[%s]: bound to %q", newVol.Name, claimToClaimKey(claim))
	return newVol, nil
}
func (ctrl *PersistentVolumeController) getBindVolumeToClaim(volume *v1.PersistentVolume, claim *v1.PersistentVolumeClaim) (*v1.PersistentVolume, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dirty := false
	shouldSetBoundByController := false
	if !isVolumeBoundToClaim(volume, claim) {
		shouldSetBoundByController = true
	}
	volumeClone := volume.DeepCopy()
	if volume.Spec.ClaimRef == nil || volume.Spec.ClaimRef.Name != claim.Name || volume.Spec.ClaimRef.Namespace != claim.Namespace || volume.Spec.ClaimRef.UID != claim.UID {
		claimRef, err := ref.GetReference(scheme.Scheme, claim)
		if err != nil {
			return nil, false, fmt.Errorf("Unexpected error getting claim reference: %v", err)
		}
		volumeClone.Spec.ClaimRef = claimRef
		dirty = true
	}
	if shouldSetBoundByController && !metav1.HasAnnotation(volumeClone.ObjectMeta, annBoundByController) {
		metav1.SetMetaDataAnnotation(&volumeClone.ObjectMeta, annBoundByController, "yes")
		dirty = true
	}
	return volumeClone, dirty, nil
}
func (ctrl *PersistentVolumeController) bindClaimToVolume(claim *v1.PersistentVolumeClaim, volume *v1.PersistentVolume) (*v1.PersistentVolumeClaim, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("updating PersistentVolumeClaim[%s]: binding to %q", claimToClaimKey(claim), volume.Name)
	dirty := false
	shouldBind := false
	if volume.Name != claim.Spec.VolumeName {
		shouldBind = true
	}
	claimClone := claim.DeepCopy()
	if shouldBind {
		dirty = true
		claimClone.Spec.VolumeName = volume.Name
		if !metav1.HasAnnotation(claimClone.ObjectMeta, annBoundByController) {
			metav1.SetMetaDataAnnotation(&claimClone.ObjectMeta, annBoundByController, "yes")
		}
	}
	if !metav1.HasAnnotation(claimClone.ObjectMeta, annBindCompleted) {
		metav1.SetMetaDataAnnotation(&claimClone.ObjectMeta, annBindCompleted, "yes")
		dirty = true
	}
	if dirty {
		klog.V(2).Infof("volume %q bound to claim %q", volume.Name, claimToClaimKey(claim))
		newClaim, err := ctrl.kubeClient.CoreV1().PersistentVolumeClaims(claim.Namespace).Update(claimClone)
		if err != nil {
			klog.V(4).Infof("updating PersistentVolumeClaim[%s]: binding to %q failed: %v", claimToClaimKey(claim), volume.Name, err)
			return newClaim, err
		}
		_, err = ctrl.storeClaimUpdate(newClaim)
		if err != nil {
			klog.V(4).Infof("updating PersistentVolumeClaim[%s]: cannot update internal cache: %v", claimToClaimKey(claim), err)
			return newClaim, err
		}
		klog.V(4).Infof("updating PersistentVolumeClaim[%s]: bound to %q", claimToClaimKey(claim), volume.Name)
		return newClaim, nil
	}
	klog.V(4).Infof("updating PersistentVolumeClaim[%s]: already bound to %q", claimToClaimKey(claim), volume.Name)
	return claim, nil
}
func (ctrl *PersistentVolumeController) bind(volume *v1.PersistentVolume, claim *v1.PersistentVolumeClaim) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	var updatedClaim *v1.PersistentVolumeClaim
	var updatedVolume *v1.PersistentVolume
	klog.V(4).Infof("binding volume %q to claim %q", volume.Name, claimToClaimKey(claim))
	if updatedVolume, err = ctrl.bindVolumeToClaim(volume, claim); err != nil {
		klog.V(3).Infof("error binding volume %q to claim %q: failed saving the volume: %v", volume.Name, claimToClaimKey(claim), err)
		return err
	}
	volume = updatedVolume
	if updatedVolume, err = ctrl.updateVolumePhase(volume, v1.VolumeBound, ""); err != nil {
		klog.V(3).Infof("error binding volume %q to claim %q: failed saving the volume status: %v", volume.Name, claimToClaimKey(claim), err)
		return err
	}
	volume = updatedVolume
	if updatedClaim, err = ctrl.bindClaimToVolume(claim, volume); err != nil {
		klog.V(3).Infof("error binding volume %q to claim %q: failed saving the claim: %v", volume.Name, claimToClaimKey(claim), err)
		return err
	}
	claim = updatedClaim
	if updatedClaim, err = ctrl.updateClaimStatus(claim, v1.ClaimBound, volume); err != nil {
		klog.V(3).Infof("error binding volume %q to claim %q: failed saving the claim status: %v", volume.Name, claimToClaimKey(claim), err)
		return err
	}
	claim = updatedClaim
	klog.V(4).Infof("volume %q bound to claim %q", volume.Name, claimToClaimKey(claim))
	klog.V(4).Infof("volume %q status after binding: %s", volume.Name, getVolumeStatusForLogging(volume))
	klog.V(4).Infof("claim %q status after binding: %s", claimToClaimKey(claim), getClaimStatusForLogging(claim))
	return nil
}
func (ctrl *PersistentVolumeController) unbindVolume(volume *v1.PersistentVolume) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("updating PersistentVolume[%s]: rolling back binding from %q", volume.Name, claimrefToClaimKey(volume.Spec.ClaimRef))
	volumeClone := volume.DeepCopy()
	if metav1.HasAnnotation(volume.ObjectMeta, annBoundByController) {
		volumeClone.Spec.ClaimRef = nil
		delete(volumeClone.Annotations, annBoundByController)
		if len(volumeClone.Annotations) == 0 {
			volumeClone.Annotations = nil
		}
	} else {
		volumeClone.Spec.ClaimRef.UID = ""
	}
	newVol, err := ctrl.kubeClient.CoreV1().PersistentVolumes().Update(volumeClone)
	if err != nil {
		klog.V(4).Infof("updating PersistentVolume[%s]: rollback failed: %v", volume.Name, err)
		return err
	}
	_, err = ctrl.storeVolumeUpdate(newVol)
	if err != nil {
		klog.V(4).Infof("updating PersistentVolume[%s]: cannot update internal cache: %v", volume.Name, err)
		return err
	}
	klog.V(4).Infof("updating PersistentVolume[%s]: rolled back", newVol.Name)
	_, err = ctrl.updateVolumePhase(newVol, v1.VolumeAvailable, "")
	return err
}
func (ctrl *PersistentVolumeController) reclaimVolume(volume *v1.PersistentVolume) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch volume.Spec.PersistentVolumeReclaimPolicy {
	case v1.PersistentVolumeReclaimRetain:
		klog.V(4).Infof("reclaimVolume[%s]: policy is Retain, nothing to do", volume.Name)
	case v1.PersistentVolumeReclaimRecycle:
		klog.V(4).Infof("reclaimVolume[%s]: policy is Recycle", volume.Name)
		opName := fmt.Sprintf("recycle-%s[%s]", volume.Name, string(volume.UID))
		ctrl.scheduleOperation(opName, func() error {
			ctrl.recycleVolumeOperation(volume)
			return nil
		})
	case v1.PersistentVolumeReclaimDelete:
		klog.V(4).Infof("reclaimVolume[%s]: policy is Delete", volume.Name)
		opName := fmt.Sprintf("delete-%s[%s]", volume.Name, string(volume.UID))
		startTime := time.Now()
		ctrl.scheduleOperation(opName, func() error {
			pluginName, err := ctrl.deleteVolumeOperation(volume)
			timeTaken := time.Since(startTime).Seconds()
			metrics.RecordVolumeOperationMetric(pluginName, "delete", timeTaken, err)
			return err
		})
	default:
		if _, err := ctrl.updateVolumePhaseWithEvent(volume, v1.VolumeFailed, v1.EventTypeWarning, "VolumeUnknownReclaimPolicy", "Volume has unrecognized PersistentVolumeReclaimPolicy"); err != nil {
			return err
		}
	}
	return nil
}
func (ctrl *PersistentVolumeController) recycleVolumeOperation(volume *v1.PersistentVolume) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("recycleVolumeOperation [%s] started", volume.Name)
	newVolume, err := ctrl.kubeClient.CoreV1().PersistentVolumes().Get(volume.Name, metav1.GetOptions{})
	if err != nil {
		klog.V(3).Infof("error reading persistent volume %q: %v", volume.Name, err)
		return
	}
	needsReclaim, err := ctrl.isVolumeReleased(newVolume)
	if err != nil {
		klog.V(3).Infof("error reading claim for volume %q: %v", volume.Name, err)
		return
	}
	if !needsReclaim {
		klog.V(3).Infof("volume %q no longer needs recycling, skipping", volume.Name)
		return
	}
	pods, used, err := ctrl.isVolumeUsed(newVolume)
	if err != nil {
		klog.V(3).Infof("can't recycle volume %q: %v", volume.Name, err)
		return
	}
	claimName := claimrefToClaimKey(volume.Spec.ClaimRef)
	_, claimCached, err := ctrl.claims.GetByKey(claimName)
	if err != nil {
		klog.V(3).Infof("error getting the claim %s from cache", claimName)
		return
	}
	if used && !claimCached {
		msg := fmt.Sprintf("Volume is used by pods: %s", strings.Join(pods, ","))
		klog.V(3).Infof("can't recycle volume %q: %s", volume.Name, msg)
		ctrl.eventRecorder.Event(volume, v1.EventTypeNormal, events.VolumeFailedRecycle, msg)
		return
	}
	volume = newVolume
	spec := vol.NewSpecFromPersistentVolume(volume, false)
	plugin, err := ctrl.volumePluginMgr.FindRecyclablePluginBySpec(spec)
	if err != nil {
		if _, err = ctrl.updateVolumePhaseWithEvent(volume, v1.VolumeFailed, v1.EventTypeWarning, events.VolumeFailedRecycle, "No recycler plugin found for the volume!"); err != nil {
			klog.V(4).Infof("recycleVolumeOperation [%s]: failed to mark volume as failed: %v", volume.Name, err)
			return
		}
		return
	}
	recorder := ctrl.newRecyclerEventRecorder(volume)
	if err = plugin.Recycle(volume.Name, spec, recorder); err != nil {
		strerr := fmt.Sprintf("Recycle failed: %s", err)
		if _, err = ctrl.updateVolumePhaseWithEvent(volume, v1.VolumeFailed, v1.EventTypeWarning, events.VolumeFailedRecycle, strerr); err != nil {
			klog.V(4).Infof("recycleVolumeOperation [%s]: failed to mark volume as failed: %v", volume.Name, err)
			return
		}
		return
	}
	klog.V(2).Infof("volume %q recycled", volume.Name)
	ctrl.eventRecorder.Event(volume, v1.EventTypeNormal, events.VolumeRecycled, "Volume recycled")
	if err = ctrl.unbindVolume(volume); err != nil {
		klog.V(3).Infof("recycleVolumeOperation [%s]: failed to make recycled volume 'Available' (%v), we will recycle the volume again", volume.Name, err)
		return
	}
	return
}
func (ctrl *PersistentVolumeController) deleteVolumeOperation(volume *v1.PersistentVolume) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("deleteVolumeOperation [%s] started", volume.Name)
	newVolume, err := ctrl.kubeClient.CoreV1().PersistentVolumes().Get(volume.Name, metav1.GetOptions{})
	if err != nil {
		klog.V(3).Infof("error reading persistent volume %q: %v", volume.Name, err)
		return "", nil
	}
	needsReclaim, err := ctrl.isVolumeReleased(newVolume)
	if err != nil {
		klog.V(3).Infof("error reading claim for volume %q: %v", volume.Name, err)
		return "", nil
	}
	if !needsReclaim {
		klog.V(3).Infof("volume %q no longer needs deletion, skipping", volume.Name)
		return "", nil
	}
	pluginName, deleted, err := ctrl.doDeleteVolume(volume)
	if err != nil {
		klog.V(3).Infof("deletion of volume %q failed: %v", volume.Name, err)
		if vol.IsDeletedVolumeInUse(err) {
			ctrl.eventRecorder.Event(volume, v1.EventTypeNormal, events.VolumeDelete, err.Error())
		} else {
			if _, err := ctrl.updateVolumePhaseWithEvent(volume, v1.VolumeFailed, v1.EventTypeWarning, events.VolumeFailedDelete, err.Error()); err != nil {
				klog.V(4).Infof("deleteVolumeOperation [%s]: failed to mark volume as failed: %v", volume.Name, err)
				return pluginName, err
			}
		}
		return pluginName, err
	}
	if !deleted {
		return pluginName, nil
	}
	klog.V(4).Infof("deleteVolumeOperation [%s]: success", volume.Name)
	if err = ctrl.kubeClient.CoreV1().PersistentVolumes().Delete(volume.Name, nil); err != nil {
		klog.V(3).Infof("failed to delete volume %q from database: %v", volume.Name, err)
		return pluginName, nil
	}
	return pluginName, nil
}
func (ctrl *PersistentVolumeController) isVolumeReleased(volume *v1.PersistentVolume) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if volume.Spec.ClaimRef == nil {
		klog.V(4).Infof("isVolumeReleased[%s]: ClaimRef is nil", volume.Name)
		return false, nil
	}
	if volume.Spec.ClaimRef.UID == "" {
		klog.V(4).Infof("isVolumeReleased[%s]: ClaimRef is not bound", volume.Name)
		return false, nil
	}
	var claim *v1.PersistentVolumeClaim
	claimName := claimrefToClaimKey(volume.Spec.ClaimRef)
	obj, found, err := ctrl.claims.GetByKey(claimName)
	if err != nil {
		return false, err
	}
	if !found {
	} else {
		var ok bool
		claim, ok = obj.(*v1.PersistentVolumeClaim)
		if !ok {
			return false, fmt.Errorf("Cannot convert object from claim cache to claim!?: %#v", obj)
		}
	}
	if claim != nil && claim.UID == volume.Spec.ClaimRef.UID {
		if len(claim.Spec.VolumeName) > 0 && claim.Spec.VolumeName != volume.Name {
			return true, nil
		}
		klog.V(4).Infof("isVolumeReleased[%s]: ClaimRef is still valid, volume is not released", volume.Name)
		return false, nil
	}
	klog.V(2).Infof("isVolumeReleased[%s]: volume is released", volume.Name)
	return true, nil
}
func (ctrl *PersistentVolumeController) isVolumeUsed(pv *v1.PersistentVolume) ([]string, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pv.Spec.ClaimRef == nil {
		return nil, false, nil
	}
	claimName := pv.Spec.ClaimRef.Name
	podNames := sets.NewString()
	pods, err := ctrl.podLister.Pods(pv.Spec.ClaimRef.Namespace).List(labels.Everything())
	if err != nil {
		return nil, false, fmt.Errorf("error listing pods: %s", err)
	}
	for _, pod := range pods {
		if util.IsPodTerminated(pod, pod.Status) {
			continue
		}
		for i := range pod.Spec.Volumes {
			usedPV := &pod.Spec.Volumes[i]
			if usedPV.PersistentVolumeClaim != nil && usedPV.PersistentVolumeClaim.ClaimName == claimName {
				podNames.Insert(pod.Namespace + "/" + pod.Name)
			}
		}
	}
	return podNames.List(), podNames.Len() != 0, nil
}
func (ctrl *PersistentVolumeController) doDeleteVolume(volume *v1.PersistentVolume) (string, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("doDeleteVolume [%s]", volume.Name)
	var err error
	plugin, err := ctrl.findDeletablePlugin(volume)
	if err != nil {
		return "", false, err
	}
	if plugin == nil {
		klog.V(3).Infof("external deleter for volume %q requested, ignoring", volume.Name)
		return "", false, nil
	}
	pluginName := plugin.GetPluginName()
	klog.V(5).Infof("found a deleter plugin %q for volume %q", pluginName, volume.Name)
	spec := vol.NewSpecFromPersistentVolume(volume, false)
	deleter, err := plugin.NewDeleter(spec)
	if err != nil {
		return pluginName, false, fmt.Errorf("Failed to create deleter for volume %q: %v", volume.Name, err)
	}
	opComplete := util.OperationCompleteHook(pluginName, "volume_delete")
	err = deleter.Delete()
	opComplete(&err)
	if err != nil {
		return pluginName, false, err
	}
	klog.V(2).Infof("volume %q deleted", volume.Name)
	return pluginName, true, nil
}
func (ctrl *PersistentVolumeController) provisionClaim(claim *v1.PersistentVolumeClaim) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctrl.enableDynamicProvisioning {
		return nil
	}
	klog.V(4).Infof("provisionClaim[%s]: started", claimToClaimKey(claim))
	opName := fmt.Sprintf("provision-%s[%s]", claimToClaimKey(claim), string(claim.UID))
	startTime := time.Now()
	ctrl.scheduleOperation(opName, func() error {
		pluginName, err := ctrl.provisionClaimOperation(claim)
		timeTaken := time.Since(startTime).Seconds()
		metrics.RecordVolumeOperationMetric(pluginName, "provision", timeTaken, err)
		return err
	})
	return nil
}
func (ctrl *PersistentVolumeController) provisionClaimOperation(claim *v1.PersistentVolumeClaim) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	claimClass := v1helper.GetPersistentVolumeClaimClass(claim)
	klog.V(4).Infof("provisionClaimOperation [%s] started, class: %q", claimToClaimKey(claim), claimClass)
	plugin, storageClass, err := ctrl.findProvisionablePlugin(claim)
	if err != nil {
		ctrl.eventRecorder.Event(claim, v1.EventTypeWarning, events.ProvisioningFailed, err.Error())
		klog.V(2).Infof("error finding provisioning plugin for claim %s: %v", claimToClaimKey(claim), err)
		return "", err
	}
	var pluginName string
	if plugin != nil {
		pluginName = plugin.GetPluginName()
	}
	newClaim, err := ctrl.setClaimProvisioner(claim, storageClass)
	if err != nil {
		klog.V(2).Infof("error saving claim %s: %v", claimToClaimKey(claim), err)
		return pluginName, err
	}
	claim = newClaim
	if plugin == nil {
		msg := fmt.Sprintf("waiting for a volume to be created, either by external provisioner %q or manually created by system administrator", storageClass.Provisioner)
		ctrl.eventRecorder.Event(claim, v1.EventTypeNormal, events.ExternalProvisioning, msg)
		klog.V(3).Infof("provisioning claim %q: %s", claimToClaimKey(claim), msg)
		return pluginName, nil
	}
	pvName := ctrl.getProvisionedVolumeNameForClaim(claim)
	volume, err := ctrl.kubeClient.CoreV1().PersistentVolumes().Get(pvName, metav1.GetOptions{})
	if err == nil && volume != nil {
		klog.V(4).Infof("provisionClaimOperation [%s]: volume already exists, skipping", claimToClaimKey(claim))
		return pluginName, err
	}
	claimRef, err := ref.GetReference(scheme.Scheme, claim)
	if err != nil {
		klog.V(3).Infof("unexpected error getting claim reference: %v", err)
		return pluginName, err
	}
	tags := make(map[string]string)
	tags[CloudVolumeCreatedForClaimNamespaceTag] = claim.Namespace
	tags[CloudVolumeCreatedForClaimNameTag] = claim.Name
	tags[CloudVolumeCreatedForVolumeNameTag] = pvName
	options := vol.VolumeOptions{PersistentVolumeReclaimPolicy: *storageClass.ReclaimPolicy, MountOptions: storageClass.MountOptions, CloudTags: &tags, ClusterName: ctrl.clusterName, PVName: pvName, PVC: claim, Parameters: storageClass.Parameters}
	if !plugin.SupportsMountOption() && len(options.MountOptions) > 0 {
		strerr := fmt.Sprintf("Mount options are not supported by the provisioner but StorageClass %q has mount options %v", storageClass.Name, options.MountOptions)
		klog.V(2).Infof("Mount options are not supported by the provisioner but claim %q's StorageClass %q has mount options %v", claimToClaimKey(claim), storageClass.Name, options.MountOptions)
		ctrl.eventRecorder.Event(claim, v1.EventTypeWarning, events.ProvisioningFailed, strerr)
		return pluginName, fmt.Errorf("provisioner %q doesn't support mount options", plugin.GetPluginName())
	}
	provisioner, err := plugin.NewProvisioner(options)
	if err != nil {
		strerr := fmt.Sprintf("Failed to create provisioner: %v", err)
		klog.V(2).Infof("failed to create provisioner for claim %q with StorageClass %q: %v", claimToClaimKey(claim), storageClass.Name, err)
		ctrl.eventRecorder.Event(claim, v1.EventTypeWarning, events.ProvisioningFailed, strerr)
		return pluginName, err
	}
	var selectedNode *v1.Node = nil
	if nodeName, ok := claim.Annotations[annSelectedNode]; ok {
		selectedNode, err = ctrl.NodeLister.Get(nodeName)
		if err != nil {
			strerr := fmt.Sprintf("Failed to get target node: %v", err)
			klog.V(3).Infof("unexpected error getting target node %q for claim %q: %v", nodeName, claimToClaimKey(claim), err)
			ctrl.eventRecorder.Event(claim, v1.EventTypeWarning, events.ProvisioningFailed, strerr)
			return pluginName, err
		}
	}
	allowedTopologies := storageClass.AllowedTopologies
	opComplete := util.OperationCompleteHook(plugin.GetPluginName(), "volume_provision")
	volume, err = provisioner.Provision(selectedNode, allowedTopologies)
	opComplete(&err)
	if err != nil {
		ctrl.rescheduleProvisioning(claim)
		strerr := fmt.Sprintf("Failed to provision volume with StorageClass %q: %v", storageClass.Name, err)
		klog.V(2).Infof("failed to provision volume for claim %q with StorageClass %q: %v", claimToClaimKey(claim), storageClass.Name, err)
		ctrl.eventRecorder.Event(claim, v1.EventTypeWarning, events.ProvisioningFailed, strerr)
		return pluginName, err
	}
	klog.V(3).Infof("volume %q for claim %q created", volume.Name, claimToClaimKey(claim))
	if volume.Name == "" {
		volume.Name = pvName
	}
	volume.Spec.ClaimRef = claimRef
	volume.Status.Phase = v1.VolumeBound
	volume.Spec.StorageClassName = claimClass
	metav1.SetMetaDataAnnotation(&volume.ObjectMeta, annBoundByController, "yes")
	metav1.SetMetaDataAnnotation(&volume.ObjectMeta, annDynamicallyProvisioned, plugin.GetPluginName())
	for i := 0; i < ctrl.createProvisionedPVRetryCount; i++ {
		klog.V(4).Infof("provisionClaimOperation [%s]: trying to save volume %s", claimToClaimKey(claim), volume.Name)
		var newVol *v1.PersistentVolume
		if newVol, err = ctrl.kubeClient.CoreV1().PersistentVolumes().Create(volume); err == nil || apierrs.IsAlreadyExists(err) {
			if err != nil {
				klog.V(3).Infof("volume %q for claim %q already exists, reusing", volume.Name, claimToClaimKey(claim))
				err = nil
			} else {
				klog.V(3).Infof("volume %q for claim %q saved", volume.Name, claimToClaimKey(claim))
				_, updateErr := ctrl.storeVolumeUpdate(newVol)
				if updateErr != nil {
					klog.V(4).Infof("provisionClaimOperation [%s]: cannot update internal cache: %v", volume.Name, updateErr)
				}
			}
			break
		}
		klog.V(3).Infof("failed to save volume %q for claim %q: %v", volume.Name, claimToClaimKey(claim), err)
		time.Sleep(ctrl.createProvisionedPVInterval)
	}
	if err != nil {
		strerr := fmt.Sprintf("Error creating provisioned PV object for claim %s: %v. Deleting the volume.", claimToClaimKey(claim), err)
		klog.V(3).Info(strerr)
		ctrl.eventRecorder.Event(claim, v1.EventTypeWarning, events.ProvisioningFailed, strerr)
		var deleteErr error
		var deleted bool
		for i := 0; i < ctrl.createProvisionedPVRetryCount; i++ {
			_, deleted, deleteErr = ctrl.doDeleteVolume(volume)
			if deleteErr == nil && deleted {
				klog.V(4).Infof("provisionClaimOperation [%s]: cleaning volume %s succeeded", claimToClaimKey(claim), volume.Name)
				break
			}
			if !deleted {
				klog.Errorf("Error finding internal deleter for volume plugin %q", plugin.GetPluginName())
				break
			}
			klog.V(3).Infof("failed to delete volume %q: %v", volume.Name, deleteErr)
			time.Sleep(ctrl.createProvisionedPVInterval)
		}
		if deleteErr != nil {
			strerr := fmt.Sprintf("Error cleaning provisioned volume for claim %s: %v. Please delete manually.", claimToClaimKey(claim), deleteErr)
			klog.V(2).Info(strerr)
			ctrl.eventRecorder.Event(claim, v1.EventTypeWarning, events.ProvisioningCleanupFailed, strerr)
		}
	} else {
		klog.V(2).Infof("volume %q provisioned for claim %q", volume.Name, claimToClaimKey(claim))
		msg := fmt.Sprintf("Successfully provisioned volume %s using %s", volume.Name, plugin.GetPluginName())
		ctrl.eventRecorder.Event(claim, v1.EventTypeNormal, events.ProvisioningSucceeded, msg)
	}
	return pluginName, nil
}
func (ctrl *PersistentVolumeController) rescheduleProvisioning(claim *v1.PersistentVolumeClaim) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, ok := claim.Annotations[annSelectedNode]; !ok {
		return
	}
	newClaim := claim.DeepCopy()
	delete(newClaim.Annotations, annSelectedNode)
	if _, err := ctrl.kubeClient.CoreV1().PersistentVolumeClaims(newClaim.Namespace).Update(newClaim); err != nil {
		klog.V(4).Infof("Failed to delete annotation 'annSelectedNode' for PersistentVolumeClaim %q: %v", claimToClaimKey(newClaim), err)
		return
	}
	if _, err := ctrl.storeClaimUpdate(newClaim); err != nil {
		klog.V(4).Infof("Updating PersistentVolumeClaim %q: cannot update internal cache: %v", claimToClaimKey(newClaim), err)
	}
}
func (ctrl *PersistentVolumeController) getProvisionedVolumeNameForClaim(claim *v1.PersistentVolumeClaim) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "pvc-" + string(claim.UID)
}
func (ctrl *PersistentVolumeController) scheduleOperation(operationName string, operation func() error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("scheduleOperation[%s]", operationName)
	if ctrl.preOperationHook != nil {
		ctrl.preOperationHook(operationName)
	}
	err := ctrl.runningOperations.Run(operationName, operation)
	if err != nil {
		switch {
		case goroutinemap.IsAlreadyExists(err):
			klog.V(4).Infof("operation %q is already running, skipping", operationName)
		case exponentialbackoff.IsExponentialBackoff(err):
			klog.V(4).Infof("operation %q postponed due to exponential backoff", operationName)
		default:
			klog.Errorf("error scheduling operation %q: %v", operationName, err)
		}
	}
}
func (ctrl *PersistentVolumeController) newRecyclerEventRecorder(volume *v1.PersistentVolume) recyclerclient.RecycleEventRecorder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(eventtype, message string) {
		ctrl.eventRecorder.Eventf(volume, eventtype, events.RecyclerPod, "Recycler pod: %s", message)
	}
}
func (ctrl *PersistentVolumeController) findProvisionablePlugin(claim *v1.PersistentVolumeClaim) (vol.ProvisionableVolumePlugin, *storage.StorageClass, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	claimClass := v1helper.GetPersistentVolumeClaimClass(claim)
	class, err := ctrl.classLister.Get(claimClass)
	if err != nil {
		return nil, nil, err
	}
	plugin, err := ctrl.volumePluginMgr.FindProvisionablePluginByName(class.Provisioner)
	if err != nil {
		if !strings.HasPrefix(class.Provisioner, "kubernetes.io/") {
			return nil, class, nil
		}
		return nil, class, err
	}
	return plugin, class, nil
}
func (ctrl *PersistentVolumeController) findDeletablePlugin(volume *v1.PersistentVolume) (vol.DeletableVolumePlugin, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var plugin vol.DeletableVolumePlugin
	if metav1.HasAnnotation(volume.ObjectMeta, annDynamicallyProvisioned) {
		provisionPluginName := volume.Annotations[annDynamicallyProvisioned]
		if provisionPluginName != "" {
			plugin, err := ctrl.volumePluginMgr.FindDeletablePluginByName(provisionPluginName)
			if err != nil {
				if !strings.HasPrefix(provisionPluginName, "kubernetes.io/") {
					return nil, nil
				}
				return nil, err
			}
			return plugin, nil
		}
	}
	spec := vol.NewSpecFromPersistentVolume(volume, false)
	plugin, err := ctrl.volumePluginMgr.FindDeletablePluginBySpec(spec)
	if err != nil {
		return nil, fmt.Errorf("Error getting deleter volume plugin for volume %q: %v", volume.Name, err)
	}
	return plugin, nil
}

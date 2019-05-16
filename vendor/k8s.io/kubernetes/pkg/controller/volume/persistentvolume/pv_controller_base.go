package persistentvolume

import (
	"fmt"
	"k8s.io/api/core/v1"
	storage "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"
	storageinformers "k8s.io/client-go/informers/storage/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/volume/persistentvolume/metrics"
	"k8s.io/kubernetes/pkg/util/goroutinemap"
	vol "k8s.io/kubernetes/pkg/volume"
	"strconv"
	"time"
)

type ControllerParameters struct {
	KubeClient                clientset.Interface
	SyncPeriod                time.Duration
	VolumePlugins             []vol.VolumePlugin
	Cloud                     cloudprovider.Interface
	ClusterName               string
	VolumeInformer            coreinformers.PersistentVolumeInformer
	ClaimInformer             coreinformers.PersistentVolumeClaimInformer
	ClassInformer             storageinformers.StorageClassInformer
	PodInformer               coreinformers.PodInformer
	NodeInformer              coreinformers.NodeInformer
	EventRecorder             record.EventRecorder
	EnableDynamicProvisioning bool
}

func NewController(p ControllerParameters) (*PersistentVolumeController, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	eventRecorder := p.EventRecorder
	if eventRecorder == nil {
		broadcaster := record.NewBroadcaster()
		broadcaster.StartLogging(klog.Infof)
		broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: p.KubeClient.CoreV1().Events("")})
		eventRecorder = broadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "persistentvolume-controller"})
	}
	controller := &PersistentVolumeController{volumes: newPersistentVolumeOrderedIndex(), claims: cache.NewStore(cache.DeletionHandlingMetaNamespaceKeyFunc), kubeClient: p.KubeClient, eventRecorder: eventRecorder, runningOperations: goroutinemap.NewGoRoutineMap(true), cloud: p.Cloud, enableDynamicProvisioning: p.EnableDynamicProvisioning, clusterName: p.ClusterName, createProvisionedPVRetryCount: createProvisionedPVRetryCount, createProvisionedPVInterval: createProvisionedPVInterval, claimQueue: workqueue.NewNamed("claims"), volumeQueue: workqueue.NewNamed("volumes"), resyncPeriod: p.SyncPeriod}
	if err := controller.volumePluginMgr.InitPlugins(p.VolumePlugins, nil, controller); err != nil {
		return nil, fmt.Errorf("Could not initialize volume plugins for PersistentVolume Controller: %v", err)
	}
	p.VolumeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		controller.enqueueWork(controller.volumeQueue, obj)
	}, UpdateFunc: func(oldObj, newObj interface{}) {
		controller.enqueueWork(controller.volumeQueue, newObj)
	}, DeleteFunc: func(obj interface{}) {
		controller.enqueueWork(controller.volumeQueue, obj)
	}})
	controller.volumeLister = p.VolumeInformer.Lister()
	controller.volumeListerSynced = p.VolumeInformer.Informer().HasSynced
	p.ClaimInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		controller.enqueueWork(controller.claimQueue, obj)
	}, UpdateFunc: func(oldObj, newObj interface{}) {
		controller.enqueueWork(controller.claimQueue, newObj)
	}, DeleteFunc: func(obj interface{}) {
		controller.enqueueWork(controller.claimQueue, obj)
	}})
	controller.claimLister = p.ClaimInformer.Lister()
	controller.claimListerSynced = p.ClaimInformer.Informer().HasSynced
	controller.classLister = p.ClassInformer.Lister()
	controller.classListerSynced = p.ClassInformer.Informer().HasSynced
	controller.podLister = p.PodInformer.Lister()
	controller.podListerSynced = p.PodInformer.Informer().HasSynced
	controller.NodeLister = p.NodeInformer.Lister()
	controller.NodeListerSynced = p.NodeInformer.Informer().HasSynced
	return controller, nil
}
func (ctrl *PersistentVolumeController) initializeCaches(volumeLister corelisters.PersistentVolumeLister, claimLister corelisters.PersistentVolumeClaimLister) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumeList, err := volumeLister.List(labels.Everything())
	if err != nil {
		klog.Errorf("PersistentVolumeController can't initialize caches: %v", err)
		return
	}
	for _, volume := range volumeList {
		volumeClone := volume.DeepCopy()
		if _, err = ctrl.storeVolumeUpdate(volumeClone); err != nil {
			klog.Errorf("error updating volume cache: %v", err)
		}
	}
	claimList, err := claimLister.List(labels.Everything())
	if err != nil {
		klog.Errorf("PersistentVolumeController can't initialize caches: %v", err)
		return
	}
	for _, claim := range claimList {
		if _, err = ctrl.storeClaimUpdate(claim.DeepCopy()); err != nil {
			klog.Errorf("error updating claim cache: %v", err)
		}
	}
	klog.V(4).Infof("controller initialized")
}
func (ctrl *PersistentVolumeController) enqueueWork(queue workqueue.Interface, obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if unknown, ok := obj.(cache.DeletedFinalStateUnknown); ok && unknown.Obj != nil {
		obj = unknown.Obj
	}
	objName, err := controller.KeyFunc(obj)
	if err != nil {
		klog.Errorf("failed to get key from object: %v", err)
		return
	}
	klog.V(5).Infof("enqueued %q for sync", objName)
	queue.Add(objName)
}
func (ctrl *PersistentVolumeController) storeVolumeUpdate(volume interface{}) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return storeObjectUpdate(ctrl.volumes.store, volume, "volume")
}
func (ctrl *PersistentVolumeController) storeClaimUpdate(claim interface{}) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return storeObjectUpdate(ctrl.claims, claim, "claim")
}
func (ctrl *PersistentVolumeController) updateVolume(volume *v1.PersistentVolume) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	new, err := ctrl.storeVolumeUpdate(volume)
	if err != nil {
		klog.Errorf("%v", err)
	}
	if !new {
		return
	}
	err = ctrl.syncVolume(volume)
	if err != nil {
		if errors.IsConflict(err) {
			klog.V(3).Infof("could not sync volume %q: %+v", volume.Name, err)
		} else {
			klog.Errorf("could not sync volume %q: %+v", volume.Name, err)
		}
	}
}
func (ctrl *PersistentVolumeController) deleteVolume(volume *v1.PersistentVolume) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_ = ctrl.volumes.store.Delete(volume)
	klog.V(4).Infof("volume %q deleted", volume.Name)
	if volume.Spec.ClaimRef == nil {
		return
	}
	claimKey := claimrefToClaimKey(volume.Spec.ClaimRef)
	klog.V(5).Infof("deleteVolume[%s]: scheduling sync of claim %q", volume.Name, claimKey)
	ctrl.claimQueue.Add(claimKey)
}
func (ctrl *PersistentVolumeController) updateClaim(claim *v1.PersistentVolumeClaim) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	new, err := ctrl.storeClaimUpdate(claim)
	if err != nil {
		klog.Errorf("%v", err)
	}
	if !new {
		return
	}
	err = ctrl.syncClaim(claim)
	if err != nil {
		if errors.IsConflict(err) {
			klog.V(3).Infof("could not sync claim %q: %+v", claimToClaimKey(claim), err)
		} else {
			klog.Errorf("could not sync volume %q: %+v", claimToClaimKey(claim), err)
		}
	}
}
func (ctrl *PersistentVolumeController) deleteClaim(claim *v1.PersistentVolumeClaim) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_ = ctrl.claims.Delete(claim)
	klog.V(4).Infof("claim %q deleted", claimToClaimKey(claim))
	volumeName := claim.Spec.VolumeName
	if volumeName == "" {
		klog.V(5).Infof("deleteClaim[%q]: volume not bound", claimToClaimKey(claim))
		return
	}
	klog.V(5).Infof("deleteClaim[%q]: scheduling sync of volume %s", claimToClaimKey(claim), volumeName)
	ctrl.volumeQueue.Add(volumeName)
}
func (ctrl *PersistentVolumeController) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	defer ctrl.claimQueue.ShutDown()
	defer ctrl.volumeQueue.ShutDown()
	klog.Infof("Starting persistent volume controller")
	defer klog.Infof("Shutting down persistent volume controller")
	if !controller.WaitForCacheSync("persistent volume", stopCh, ctrl.volumeListerSynced, ctrl.claimListerSynced, ctrl.classListerSynced, ctrl.podListerSynced, ctrl.NodeListerSynced) {
		return
	}
	ctrl.initializeCaches(ctrl.volumeLister, ctrl.claimLister)
	go wait.Until(ctrl.resync, ctrl.resyncPeriod, stopCh)
	go wait.Until(ctrl.volumeWorker, time.Second, stopCh)
	go wait.Until(ctrl.claimWorker, time.Second, stopCh)
	metrics.Register(ctrl.volumes.store, ctrl.claims)
	<-stopCh
}
func (ctrl *PersistentVolumeController) volumeWorker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	workFunc := func() bool {
		keyObj, quit := ctrl.volumeQueue.Get()
		if quit {
			return true
		}
		defer ctrl.volumeQueue.Done(keyObj)
		key := keyObj.(string)
		klog.V(5).Infof("volumeWorker[%s]", key)
		_, name, err := cache.SplitMetaNamespaceKey(key)
		if err != nil {
			klog.V(4).Infof("error getting name of volume %q to get volume from informer: %v", key, err)
			return false
		}
		volume, err := ctrl.volumeLister.Get(name)
		if err == nil {
			ctrl.updateVolume(volume)
			return false
		}
		if !errors.IsNotFound(err) {
			klog.V(2).Infof("error getting volume %q from informer: %v", key, err)
			return false
		}
		volumeObj, found, err := ctrl.volumes.store.GetByKey(key)
		if err != nil {
			klog.V(2).Infof("error getting volume %q from cache: %v", key, err)
			return false
		}
		if !found {
			klog.V(2).Infof("deletion of volume %q was already processed", key)
			return false
		}
		volume, ok := volumeObj.(*v1.PersistentVolume)
		if !ok {
			klog.Errorf("expected volume, got %+v", volumeObj)
			return false
		}
		ctrl.deleteVolume(volume)
		return false
	}
	for {
		if quit := workFunc(); quit {
			klog.Infof("volume worker queue shutting down")
			return
		}
	}
}
func (ctrl *PersistentVolumeController) claimWorker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	workFunc := func() bool {
		keyObj, quit := ctrl.claimQueue.Get()
		if quit {
			return true
		}
		defer ctrl.claimQueue.Done(keyObj)
		key := keyObj.(string)
		klog.V(5).Infof("claimWorker[%s]", key)
		namespace, name, err := cache.SplitMetaNamespaceKey(key)
		if err != nil {
			klog.V(4).Infof("error getting namespace & name of claim %q to get claim from informer: %v", key, err)
			return false
		}
		claim, err := ctrl.claimLister.PersistentVolumeClaims(namespace).Get(name)
		if err == nil {
			ctrl.updateClaim(claim)
			return false
		}
		if !errors.IsNotFound(err) {
			klog.V(2).Infof("error getting claim %q from informer: %v", key, err)
			return false
		}
		claimObj, found, err := ctrl.claims.GetByKey(key)
		if err != nil {
			klog.V(2).Infof("error getting claim %q from cache: %v", key, err)
			return false
		}
		if !found {
			klog.V(2).Infof("deletion of claim %q was already processed", key)
			return false
		}
		claim, ok := claimObj.(*v1.PersistentVolumeClaim)
		if !ok {
			klog.Errorf("expected claim, got %+v", claimObj)
			return false
		}
		ctrl.deleteClaim(claim)
		return false
	}
	for {
		if quit := workFunc(); quit {
			klog.Infof("claim worker queue shutting down")
			return
		}
	}
}
func (ctrl *PersistentVolumeController) resync() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("resyncing PV controller")
	pvcs, err := ctrl.claimLister.List(labels.NewSelector())
	if err != nil {
		klog.Warningf("cannot list claims: %s", err)
		return
	}
	for _, pvc := range pvcs {
		ctrl.enqueueWork(ctrl.claimQueue, pvc)
	}
	pvs, err := ctrl.volumeLister.List(labels.NewSelector())
	if err != nil {
		klog.Warningf("cannot list persistent volumes: %s", err)
		return
	}
	for _, pv := range pvs {
		ctrl.enqueueWork(ctrl.volumeQueue, pv)
	}
}
func (ctrl *PersistentVolumeController) setClaimProvisioner(claim *v1.PersistentVolumeClaim, class *storage.StorageClass) (*v1.PersistentVolumeClaim, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if val, ok := claim.Annotations[annStorageProvisioner]; ok && val == class.Provisioner {
		return claim, nil
	}
	claimClone := claim.DeepCopy()
	metav1.SetMetaDataAnnotation(&claimClone.ObjectMeta, annStorageProvisioner, class.Provisioner)
	newClaim, err := ctrl.kubeClient.CoreV1().PersistentVolumeClaims(claim.Namespace).Update(claimClone)
	if err != nil {
		return newClaim, err
	}
	_, err = ctrl.storeClaimUpdate(newClaim)
	if err != nil {
		return newClaim, err
	}
	return newClaim, nil
}
func getClaimStatusForLogging(claim *v1.PersistentVolumeClaim) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	bound := metav1.HasAnnotation(claim.ObjectMeta, annBindCompleted)
	boundByController := metav1.HasAnnotation(claim.ObjectMeta, annBoundByController)
	return fmt.Sprintf("phase: %s, bound to: %q, bindCompleted: %v, boundByController: %v", claim.Status.Phase, claim.Spec.VolumeName, bound, boundByController)
}
func getVolumeStatusForLogging(volume *v1.PersistentVolume) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	boundByController := metav1.HasAnnotation(volume.ObjectMeta, annBoundByController)
	claimName := ""
	if volume.Spec.ClaimRef != nil {
		claimName = fmt.Sprintf("%s/%s (uid: %s)", volume.Spec.ClaimRef.Namespace, volume.Spec.ClaimRef.Name, volume.Spec.ClaimRef.UID)
	}
	return fmt.Sprintf("phase: %s, bound to: %q, boundByController: %v", volume.Status.Phase, claimName, boundByController)
}
func isVolumeBoundToClaim(volume *v1.PersistentVolume, claim *v1.PersistentVolumeClaim) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if volume.Spec.ClaimRef == nil {
		return false
	}
	if claim.Name != volume.Spec.ClaimRef.Name || claim.Namespace != volume.Spec.ClaimRef.Namespace {
		return false
	}
	if volume.Spec.ClaimRef.UID != "" && claim.UID != volume.Spec.ClaimRef.UID {
		return false
	}
	return true
}
func storeObjectUpdate(store cache.Store, obj interface{}, className string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objName, err := controller.KeyFunc(obj)
	if err != nil {
		return false, fmt.Errorf("Couldn't get key for object %+v: %v", obj, err)
	}
	oldObj, found, err := store.Get(obj)
	if err != nil {
		return false, fmt.Errorf("Error finding %s %q in controller cache: %v", className, objName, err)
	}
	objAccessor, err := meta.Accessor(obj)
	if err != nil {
		return false, err
	}
	if !found {
		klog.V(4).Infof("storeObjectUpdate: adding %s %q, version %s", className, objName, objAccessor.GetResourceVersion())
		if err = store.Add(obj); err != nil {
			return false, fmt.Errorf("Error adding %s %q to controller cache: %v", className, objName, err)
		}
		return true, nil
	}
	oldObjAccessor, err := meta.Accessor(oldObj)
	if err != nil {
		return false, err
	}
	objResourceVersion, err := strconv.ParseInt(objAccessor.GetResourceVersion(), 10, 64)
	if err != nil {
		return false, fmt.Errorf("Error parsing ResourceVersion %q of %s %q: %s", objAccessor.GetResourceVersion(), className, objName, err)
	}
	oldObjResourceVersion, err := strconv.ParseInt(oldObjAccessor.GetResourceVersion(), 10, 64)
	if err != nil {
		return false, fmt.Errorf("Error parsing old ResourceVersion %q of %s %q: %s", oldObjAccessor.GetResourceVersion(), className, objName, err)
	}
	if oldObjResourceVersion > objResourceVersion {
		klog.V(4).Infof("storeObjectUpdate: ignoring %s %q version %s", className, objName, objAccessor.GetResourceVersion())
		return false, nil
	}
	klog.V(4).Infof("storeObjectUpdate updating %s %q with version %s", className, objName, objAccessor.GetResourceVersion())
	if err = store.Update(obj); err != nil {
		return false, fmt.Errorf("Error updating %s %q in controller cache: %v", className, objName, err)
	}
	return true, nil
}

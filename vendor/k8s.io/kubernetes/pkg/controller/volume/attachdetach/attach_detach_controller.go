package attachdetach

import (
	"fmt"
	goformat "fmt"
	authenticationv1 "k8s.io/api/authentication/v1"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	kcache "k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	cloudprovider "k8s.io/cloud-provider"
	csiclient "k8s.io/csi-api/pkg/client/clientset/versioned"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/cache"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/metrics"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/populator"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/reconciler"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/statusupdater"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/util"
	"k8s.io/kubernetes/pkg/util/mount"
	"k8s.io/kubernetes/pkg/volume"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
	"k8s.io/kubernetes/pkg/volume/util/operationexecutor"
	"k8s.io/kubernetes/pkg/volume/util/volumepathhandler"
	"net"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type TimerConfig struct {
	ReconcilerLoopPeriod                              time.Duration
	ReconcilerMaxWaitForUnmountDuration               time.Duration
	DesiredStateOfWorldPopulatorLoopSleepPeriod       time.Duration
	DesiredStateOfWorldPopulatorListPodsRetryDuration time.Duration
}

var DefaultTimerConfig TimerConfig = TimerConfig{ReconcilerLoopPeriod: 100 * time.Millisecond, ReconcilerMaxWaitForUnmountDuration: 6 * time.Minute, DesiredStateOfWorldPopulatorLoopSleepPeriod: 1 * time.Minute, DesiredStateOfWorldPopulatorListPodsRetryDuration: 3 * time.Minute}

type AttachDetachController interface {
	Run(stopCh <-chan struct{})
	GetDesiredStateOfWorld() cache.DesiredStateOfWorld
}

func NewAttachDetachController(kubeClient clientset.Interface, csiClient csiclient.Interface, podInformer coreinformers.PodInformer, nodeInformer coreinformers.NodeInformer, pvcInformer coreinformers.PersistentVolumeClaimInformer, pvInformer coreinformers.PersistentVolumeInformer, cloud cloudprovider.Interface, plugins []volume.VolumePlugin, prober volume.DynamicPluginProber, disableReconciliationSync bool, reconcilerSyncDuration time.Duration, timerConfig TimerConfig) (AttachDetachController, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	adc := &attachDetachController{kubeClient: kubeClient, csiClient: csiClient, pvcLister: pvcInformer.Lister(), pvcsSynced: pvcInformer.Informer().HasSynced, pvLister: pvInformer.Lister(), pvsSynced: pvInformer.Informer().HasSynced, podLister: podInformer.Lister(), podsSynced: podInformer.Informer().HasSynced, podIndexer: podInformer.Informer().GetIndexer(), nodeLister: nodeInformer.Lister(), nodesSynced: nodeInformer.Informer().HasSynced, cloud: cloud, pvcQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "pvcs")}
	if err := adc.volumePluginMgr.InitPlugins(plugins, prober, adc); err != nil {
		return nil, fmt.Errorf("Could not initialize volume plugins for Attach/Detach Controller: %+v", err)
	}
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "attachdetach-controller"})
	blkutil := volumepathhandler.NewBlockVolumePathHandler()
	adc.desiredStateOfWorld = cache.NewDesiredStateOfWorld(&adc.volumePluginMgr)
	adc.actualStateOfWorld = cache.NewActualStateOfWorld(&adc.volumePluginMgr)
	adc.attacherDetacher = operationexecutor.NewOperationExecutor(operationexecutor.NewOperationGenerator(kubeClient, &adc.volumePluginMgr, recorder, false, blkutil))
	adc.nodeStatusUpdater = statusupdater.NewNodeStatusUpdater(kubeClient, nodeInformer.Lister(), adc.actualStateOfWorld)
	adc.reconciler = reconciler.NewReconciler(timerConfig.ReconcilerLoopPeriod, timerConfig.ReconcilerMaxWaitForUnmountDuration, reconcilerSyncDuration, disableReconciliationSync, adc.desiredStateOfWorld, adc.actualStateOfWorld, adc.attacherDetacher, adc.nodeStatusUpdater, recorder)
	adc.desiredStateOfWorldPopulator = populator.NewDesiredStateOfWorldPopulator(timerConfig.DesiredStateOfWorldPopulatorLoopSleepPeriod, timerConfig.DesiredStateOfWorldPopulatorListPodsRetryDuration, podInformer.Lister(), adc.desiredStateOfWorld, &adc.volumePluginMgr, pvcInformer.Lister(), pvInformer.Lister())
	podInformer.Informer().AddEventHandler(kcache.ResourceEventHandlerFuncs{AddFunc: adc.podAdd, UpdateFunc: adc.podUpdate, DeleteFunc: adc.podDelete})
	adc.podIndexer.AddIndexers(kcache.Indexers{pvcKeyIndex: indexByPVCKey})
	nodeInformer.Informer().AddEventHandler(kcache.ResourceEventHandlerFuncs{AddFunc: adc.nodeAdd, UpdateFunc: adc.nodeUpdate, DeleteFunc: adc.nodeDelete})
	pvcInformer.Informer().AddEventHandler(kcache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		adc.enqueuePVC(obj)
	}, UpdateFunc: func(old, new interface{}) {
		adc.enqueuePVC(new)
	}})
	return adc, nil
}

const (
	pvcKeyIndex string = "pvcKey"
)

func indexByPVCKey(obj interface{}) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod, ok := obj.(*v1.Pod)
	if !ok {
		return []string{}, nil
	}
	if len(pod.Spec.NodeName) == 0 || volumeutil.IsPodTerminated(pod, pod.Status) {
		return []string{}, nil
	}
	keys := []string{}
	for _, podVolume := range pod.Spec.Volumes {
		if pvcSource := podVolume.VolumeSource.PersistentVolumeClaim; pvcSource != nil {
			keys = append(keys, fmt.Sprintf("%s/%s", pod.Namespace, pvcSource.ClaimName))
		}
	}
	return keys, nil
}

type attachDetachController struct {
	kubeClient                   clientset.Interface
	csiClient                    csiclient.Interface
	pvcLister                    corelisters.PersistentVolumeClaimLister
	pvcsSynced                   kcache.InformerSynced
	pvLister                     corelisters.PersistentVolumeLister
	pvsSynced                    kcache.InformerSynced
	podLister                    corelisters.PodLister
	podsSynced                   kcache.InformerSynced
	podIndexer                   kcache.Indexer
	nodeLister                   corelisters.NodeLister
	nodesSynced                  kcache.InformerSynced
	cloud                        cloudprovider.Interface
	volumePluginMgr              volume.VolumePluginMgr
	desiredStateOfWorld          cache.DesiredStateOfWorld
	actualStateOfWorld           cache.ActualStateOfWorld
	attacherDetacher             operationexecutor.OperationExecutor
	reconciler                   reconciler.Reconciler
	nodeStatusUpdater            statusupdater.NodeStatusUpdater
	desiredStateOfWorldPopulator populator.DesiredStateOfWorldPopulator
	recorder                     record.EventRecorder
	pvcQueue                     workqueue.RateLimitingInterface
}

func (adc *attachDetachController) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer runtime.HandleCrash()
	defer adc.pvcQueue.ShutDown()
	klog.Infof("Starting attach detach controller")
	defer klog.Infof("Shutting down attach detach controller")
	if !controller.WaitForCacheSync("attach detach", stopCh, adc.podsSynced, adc.nodesSynced, adc.pvcsSynced, adc.pvsSynced) {
		return
	}
	err := adc.populateActualStateOfWorld()
	if err != nil {
		klog.Errorf("Error populating the actual state of world: %v", err)
	}
	err = adc.populateDesiredStateOfWorld()
	if err != nil {
		klog.Errorf("Error populating the desired state of world: %v", err)
	}
	go adc.reconciler.Run(stopCh)
	go adc.desiredStateOfWorldPopulator.Run(stopCh)
	go wait.Until(adc.pvcWorker, time.Second, stopCh)
	metrics.Register(adc.pvcLister, adc.pvLister, adc.podLister, adc.actualStateOfWorld, adc.desiredStateOfWorld, &adc.volumePluginMgr)
	<-stopCh
}
func (adc *attachDetachController) populateActualStateOfWorld() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(5).Infof("Populating ActualStateOfworld")
	nodes, err := adc.nodeLister.List(labels.Everything())
	if err != nil {
		return err
	}
	for _, node := range nodes {
		nodeName := types.NodeName(node.Name)
		for _, attachedVolume := range node.Status.VolumesAttached {
			uniqueName := attachedVolume.Name
			err = adc.actualStateOfWorld.MarkVolumeAsAttached(uniqueName, nil, nodeName, attachedVolume.DevicePath)
			if err != nil {
				klog.Errorf("Failed to mark the volume as attached: %v", err)
				continue
			}
			adc.processVolumesInUse(nodeName, node.Status.VolumesInUse)
			adc.addNodeToDswp(node, types.NodeName(node.Name))
		}
	}
	return nil
}
func (adc *attachDetachController) getNodeVolumeDevicePath(volumeName v1.UniqueVolumeName, nodeName types.NodeName) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var devicePath string
	var found bool
	node, err := adc.nodeLister.Get(string(nodeName))
	if err != nil {
		return devicePath, err
	}
	for _, attachedVolume := range node.Status.VolumesAttached {
		if volumeName == attachedVolume.Name {
			devicePath = attachedVolume.DevicePath
			found = true
			break
		}
	}
	if !found {
		err = fmt.Errorf("Volume %s not found on node %s", volumeName, nodeName)
	}
	return devicePath, err
}
func (adc *attachDetachController) populateDesiredStateOfWorld() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(5).Infof("Populating DesiredStateOfworld")
	pods, err := adc.podLister.List(labels.Everything())
	if err != nil {
		return err
	}
	for _, pod := range pods {
		podToAdd := pod
		adc.podAdd(podToAdd)
		for _, podVolume := range podToAdd.Spec.Volumes {
			volumeSpec, err := util.CreateVolumeSpec(podVolume, podToAdd.Namespace, adc.pvcLister, adc.pvLister)
			if err != nil {
				klog.Errorf("Error creating spec for volume %q, pod %q/%q: %v", podVolume.Name, podToAdd.Namespace, podToAdd.Name, err)
				continue
			}
			nodeName := types.NodeName(podToAdd.Spec.NodeName)
			plugin, err := adc.volumePluginMgr.FindAttachablePluginBySpec(volumeSpec)
			if err != nil || plugin == nil {
				klog.V(10).Infof("Skipping volume %q for pod %q/%q: it does not implement attacher interface. err=%v", podVolume.Name, podToAdd.Namespace, podToAdd.Name, err)
				continue
			}
			volumeName, err := volumeutil.GetUniqueVolumeNameFromSpec(plugin, volumeSpec)
			if err != nil {
				klog.Errorf("Failed to find unique name for volume %q, pod %q/%q: %v", podVolume.Name, podToAdd.Namespace, podToAdd.Name, err)
				continue
			}
			if adc.actualStateOfWorld.VolumeNodeExists(volumeName, nodeName) {
				devicePath, err := adc.getNodeVolumeDevicePath(volumeName, nodeName)
				if err != nil {
					klog.Errorf("Failed to find device path: %v", err)
					continue
				}
				err = adc.actualStateOfWorld.MarkVolumeAsAttached(volumeName, volumeSpec, nodeName, devicePath)
				if err != nil {
					klog.Errorf("Failed to update volume spec for node %s: %v", nodeName, err)
				}
			}
		}
	}
	return nil
}
func (adc *attachDetachController) podAdd(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod, ok := obj.(*v1.Pod)
	if pod == nil || !ok {
		return
	}
	if pod.Spec.NodeName == "" {
		return
	}
	volumeActionFlag := util.DetermineVolumeAction(pod, adc.desiredStateOfWorld, true)
	util.ProcessPodVolumes(pod, volumeActionFlag, adc.desiredStateOfWorld, &adc.volumePluginMgr, adc.pvcLister, adc.pvLister)
}
func (adc *attachDetachController) GetDesiredStateOfWorld() cache.DesiredStateOfWorld {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return adc.desiredStateOfWorld
}
func (adc *attachDetachController) podUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod, ok := newObj.(*v1.Pod)
	if pod == nil || !ok {
		return
	}
	if pod.Spec.NodeName == "" {
		return
	}
	volumeActionFlag := util.DetermineVolumeAction(pod, adc.desiredStateOfWorld, true)
	util.ProcessPodVolumes(pod, volumeActionFlag, adc.desiredStateOfWorld, &adc.volumePluginMgr, adc.pvcLister, adc.pvLister)
}
func (adc *attachDetachController) podDelete(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod, ok := obj.(*v1.Pod)
	if pod == nil || !ok {
		return
	}
	util.ProcessPodVolumes(pod, false, adc.desiredStateOfWorld, &adc.volumePluginMgr, adc.pvcLister, adc.pvLister)
}
func (adc *attachDetachController) nodeAdd(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	node, ok := obj.(*v1.Node)
	if node == nil || !ok {
		return
	}
	nodeName := types.NodeName(node.Name)
	adc.nodeUpdate(nil, obj)
	adc.actualStateOfWorld.SetNodeStatusUpdateNeeded(nodeName)
}
func (adc *attachDetachController) nodeUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	node, ok := newObj.(*v1.Node)
	if node == nil || !ok {
		return
	}
	nodeName := types.NodeName(node.Name)
	adc.addNodeToDswp(node, nodeName)
	adc.processVolumesInUse(nodeName, node.Status.VolumesInUse)
}
func (adc *attachDetachController) nodeDelete(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	node, ok := obj.(*v1.Node)
	if node == nil || !ok {
		return
	}
	nodeName := types.NodeName(node.Name)
	if err := adc.desiredStateOfWorld.DeleteNode(nodeName); err != nil {
		klog.Infof("error removing node %q from desired-state-of-world: %v", nodeName, err)
	}
	adc.processVolumesInUse(nodeName, node.Status.VolumesInUse)
}
func (adc *attachDetachController) enqueuePVC(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := kcache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
		return
	}
	adc.pvcQueue.Add(key)
}
func (adc *attachDetachController) pvcWorker() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for adc.processNextItem() {
	}
}
func (adc *attachDetachController) processNextItem() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	keyObj, shutdown := adc.pvcQueue.Get()
	if shutdown {
		return false
	}
	defer adc.pvcQueue.Done(keyObj)
	if err := adc.syncPVCByKey(keyObj.(string)); err != nil {
		adc.pvcQueue.AddRateLimited(keyObj)
		runtime.HandleError(fmt.Errorf("Failed to sync pvc %q, will retry again: %v", keyObj.(string), err))
		return true
	}
	adc.pvcQueue.Forget(keyObj)
	return true
}
func (adc *attachDetachController) syncPVCByKey(key string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(5).Infof("syncPVCByKey[%s]", key)
	namespace, name, err := kcache.SplitMetaNamespaceKey(key)
	if err != nil {
		klog.V(4).Infof("error getting namespace & name of pvc %q to get pvc from informer: %v", key, err)
		return nil
	}
	pvc, err := adc.pvcLister.PersistentVolumeClaims(namespace).Get(name)
	if apierrors.IsNotFound(err) {
		klog.V(4).Infof("error getting pvc %q from informer: %v", key, err)
		return nil
	}
	if err != nil {
		return err
	}
	if pvc.Status.Phase != v1.ClaimBound || pvc.Spec.VolumeName == "" {
		return nil
	}
	objs, err := adc.podIndexer.ByIndex(pvcKeyIndex, key)
	if err != nil {
		return err
	}
	for _, obj := range objs {
		pod, ok := obj.(*v1.Pod)
		if !ok {
			continue
		}
		volumeActionFlag := util.DetermineVolumeAction(pod, adc.desiredStateOfWorld, true)
		util.ProcessPodVolumes(pod, volumeActionFlag, adc.desiredStateOfWorld, &adc.volumePluginMgr, adc.pvcLister, adc.pvLister)
	}
	return nil
}
func (adc *attachDetachController) processVolumesInUse(nodeName types.NodeName, volumesInUse []v1.UniqueVolumeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("processVolumesInUse for node %q", nodeName)
	for _, attachedVolume := range adc.actualStateOfWorld.GetAttachedVolumesForNode(nodeName) {
		mounted := false
		for _, volumeInUse := range volumesInUse {
			if attachedVolume.VolumeName == volumeInUse {
				mounted = true
				break
			}
		}
		err := adc.actualStateOfWorld.SetVolumeMountedByNode(attachedVolume.VolumeName, nodeName, mounted)
		if err != nil {
			klog.Warningf("SetVolumeMountedByNode(%q, %q, %v) returned an error: %v", attachedVolume.VolumeName, nodeName, mounted, err)
		}
	}
}
func (adc *attachDetachController) GetPluginDir(podUID string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (adc *attachDetachController) GetVolumeDevicePluginDir(podUID string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (adc *attachDetachController) GetPodsDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (adc *attachDetachController) GetPodVolumeDir(podUID types.UID, pluginName, volumeName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (adc *attachDetachController) GetPodPluginDir(podUID types.UID, pluginName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (adc *attachDetachController) GetPodVolumeDeviceDir(podUID types.UID, pluginName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (adc *attachDetachController) GetKubeClient() clientset.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return adc.kubeClient
}
func (adc *attachDetachController) NewWrapperMounter(volName string, spec volume.Spec, pod *v1.Pod, opts volume.VolumeOptions) (volume.Mounter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("NewWrapperMounter not supported by Attach/Detach controller's VolumeHost implementation")
}
func (adc *attachDetachController) NewWrapperUnmounter(volName string, spec volume.Spec, podUID types.UID) (volume.Unmounter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("NewWrapperUnmounter not supported by Attach/Detach controller's VolumeHost implementation")
}
func (adc *attachDetachController) GetCloudProvider() cloudprovider.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return adc.cloud
}
func (adc *attachDetachController) GetMounter(pluginName string) mount.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (adc *attachDetachController) GetHostName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (adc *attachDetachController) GetHostIP() (net.IP, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("GetHostIP() not supported by Attach/Detach controller's VolumeHost implementation")
}
func (adc *attachDetachController) GetNodeAllocatable() (v1.ResourceList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v1.ResourceList{}, nil
}
func (adc *attachDetachController) GetSecretFunc() func(namespace, name string) (*v1.Secret, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(_, _ string) (*v1.Secret, error) {
		return nil, fmt.Errorf("GetSecret unsupported in attachDetachController")
	}
}
func (adc *attachDetachController) GetConfigMapFunc() func(namespace, name string) (*v1.ConfigMap, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(_, _ string) (*v1.ConfigMap, error) {
		return nil, fmt.Errorf("GetConfigMap unsupported in attachDetachController")
	}
}
func (adc *attachDetachController) GetServiceAccountTokenFunc() func(_, _ string, _ *authenticationv1.TokenRequest) (*authenticationv1.TokenRequest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(_, _ string, _ *authenticationv1.TokenRequest) (*authenticationv1.TokenRequest, error) {
		return nil, fmt.Errorf("GetServiceAccountToken unsupported in attachDetachController")
	}
}
func (adc *attachDetachController) DeleteServiceAccountTokenFunc() func(types.UID) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(types.UID) {
		klog.Errorf("DeleteServiceAccountToken unsupported in attachDetachController")
	}
}
func (adc *attachDetachController) GetExec(pluginName string) mount.Exec {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return mount.NewOsExec()
}
func (adc *attachDetachController) addNodeToDswp(node *v1.Node, nodeName types.NodeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, exists := node.Annotations[volumeutil.ControllerManagedAttachAnnotation]; exists {
		keepTerminatedPodVolumes := false
		if t, ok := node.Annotations[volumeutil.KeepTerminatedPodVolumesAnnotation]; ok {
			keepTerminatedPodVolumes = (t == "true")
		}
		adc.desiredStateOfWorld.AddNode(nodeName, keepTerminatedPodVolumes)
	}
}
func (adc *attachDetachController) GetNodeLabels() (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("GetNodeLabels() unsupported in Attach/Detach controller")
}
func (adc *attachDetachController) GetNodeName() types.NodeName {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (adc *attachDetachController) GetEventRecorder() record.EventRecorder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return adc.recorder
}
func (adc *attachDetachController) GetCSIClient() csiclient.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return adc.csiClient
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

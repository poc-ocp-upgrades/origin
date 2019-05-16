package persistentvolume

import (
	"fmt"
	authenticationv1 "k8s.io/api/authentication/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	cloudprovider "k8s.io/cloud-provider"
	csiclientset "k8s.io/csi-api/pkg/client/clientset/versioned"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/util/mount"
	vol "k8s.io/kubernetes/pkg/volume"
	"net"
)

var _ vol.VolumeHost = &PersistentVolumeController{}

func (ctrl *PersistentVolumeController) GetPluginDir(pluginName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (ctrl *PersistentVolumeController) GetVolumeDevicePluginDir(pluginName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (ctrl *PersistentVolumeController) GetPodsDir() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (ctrl *PersistentVolumeController) GetPodVolumeDir(podUID types.UID, pluginName string, volumeName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (ctrl *PersistentVolumeController) GetPodPluginDir(podUID types.UID, pluginName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (ctrl *PersistentVolumeController) GetPodVolumeDeviceDir(ppodUID types.UID, pluginName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (ctrl *PersistentVolumeController) GetKubeClient() clientset.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ctrl.kubeClient
}
func (ctrl *PersistentVolumeController) NewWrapperMounter(volName string, spec vol.Spec, pod *v1.Pod, opts vol.VolumeOptions) (vol.Mounter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("PersistentVolumeController.NewWrapperMounter is not implemented")
}
func (ctrl *PersistentVolumeController) NewWrapperUnmounter(volName string, spec vol.Spec, podUID types.UID) (vol.Unmounter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("PersistentVolumeController.NewWrapperMounter is not implemented")
}
func (ctrl *PersistentVolumeController) GetCloudProvider() cloudprovider.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ctrl.cloud
}
func (ctrl *PersistentVolumeController) GetMounter(pluginName string) mount.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (ctrl *PersistentVolumeController) GetHostName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (ctrl *PersistentVolumeController) GetHostIP() (net.IP, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("PersistentVolumeController.GetHostIP() is not implemented")
}
func (ctrl *PersistentVolumeController) GetNodeAllocatable() (v1.ResourceList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v1.ResourceList{}, nil
}
func (ctrl *PersistentVolumeController) GetSecretFunc() func(namespace, name string) (*v1.Secret, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(_, _ string) (*v1.Secret, error) {
		return nil, fmt.Errorf("GetSecret unsupported in PersistentVolumeController")
	}
}
func (ctrl *PersistentVolumeController) GetConfigMapFunc() func(namespace, name string) (*v1.ConfigMap, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(_, _ string) (*v1.ConfigMap, error) {
		return nil, fmt.Errorf("GetConfigMap unsupported in PersistentVolumeController")
	}
}
func (ctrl *PersistentVolumeController) GetServiceAccountTokenFunc() func(_, _ string, _ *authenticationv1.TokenRequest) (*authenticationv1.TokenRequest, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(_, _ string, _ *authenticationv1.TokenRequest) (*authenticationv1.TokenRequest, error) {
		return nil, fmt.Errorf("GetServiceAccountToken unsupported in PersistentVolumeController")
	}
}
func (ctrl *PersistentVolumeController) DeleteServiceAccountTokenFunc() func(types.UID) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(types.UID) {
		klog.Errorf("DeleteServiceAccountToken unsupported in PersistentVolumeController")
	}
}
func (adc *PersistentVolumeController) GetExec(pluginName string) mount.Exec {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return mount.NewOsExec()
}
func (ctrl *PersistentVolumeController) GetNodeLabels() (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, fmt.Errorf("GetNodeLabels() unsupported in PersistentVolumeController")
}
func (ctrl *PersistentVolumeController) GetNodeName() types.NodeName {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (ctrl *PersistentVolumeController) GetEventRecorder() record.EventRecorder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ctrl.eventRecorder
}
func (ctrl *PersistentVolumeController) GetCSIClient() csiclientset.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}

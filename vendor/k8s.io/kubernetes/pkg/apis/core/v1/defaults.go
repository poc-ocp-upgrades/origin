package v1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/util/parsers"
	utilpointer "k8s.io/utils/pointer"
	"time"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return RegisterDefaults(scheme)
}
func SetDefaults_ResourceList(obj *v1.ResourceList) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for key, val := range *obj {
		const milliScale = -3
		val.RoundUp(milliScale)
		(*obj)[v1.ResourceName(key)] = val
	}
}
func SetDefaults_ReplicationController(obj *v1.ReplicationController) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var labels map[string]string
	if obj.Spec.Template != nil {
		labels = obj.Spec.Template.Labels
	}
	if labels != nil {
		if len(obj.Spec.Selector) == 0 {
			obj.Spec.Selector = labels
		}
		if len(obj.Labels) == 0 {
			obj.Labels = labels
		}
	}
	if obj.Spec.Replicas == nil {
		obj.Spec.Replicas = new(int32)
		*obj.Spec.Replicas = 1
	}
}
func SetDefaults_Volume(obj *v1.Volume) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if utilpointer.AllPtrFieldsNil(&obj.VolumeSource) {
		obj.VolumeSource = v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}}
	}
}
func SetDefaults_ContainerPort(obj *v1.ContainerPort) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Protocol == "" {
		obj.Protocol = v1.ProtocolTCP
	}
}
func SetDefaults_Container(obj *v1.Container) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.ImagePullPolicy == "" {
		_, tag, _, _ := parsers.ParseImageName(obj.Image)
		if tag == "latest" {
			obj.ImagePullPolicy = v1.PullAlways
		} else {
			obj.ImagePullPolicy = v1.PullIfNotPresent
		}
	}
	if obj.TerminationMessagePath == "" {
		obj.TerminationMessagePath = v1.TerminationMessagePathDefault
	}
	if obj.TerminationMessagePolicy == "" {
		obj.TerminationMessagePolicy = v1.TerminationMessageReadFile
	}
}
func SetDefaults_Service(obj *v1.Service) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Spec.SessionAffinity == "" {
		obj.Spec.SessionAffinity = v1.ServiceAffinityNone
	}
	if obj.Spec.SessionAffinity == v1.ServiceAffinityNone {
		obj.Spec.SessionAffinityConfig = nil
	}
	if obj.Spec.SessionAffinity == v1.ServiceAffinityClientIP {
		if obj.Spec.SessionAffinityConfig == nil || obj.Spec.SessionAffinityConfig.ClientIP == nil || obj.Spec.SessionAffinityConfig.ClientIP.TimeoutSeconds == nil {
			timeoutSeconds := v1.DefaultClientIPServiceAffinitySeconds
			obj.Spec.SessionAffinityConfig = &v1.SessionAffinityConfig{ClientIP: &v1.ClientIPConfig{TimeoutSeconds: &timeoutSeconds}}
		}
	}
	if obj.Spec.Type == "" {
		obj.Spec.Type = v1.ServiceTypeClusterIP
	}
	for i := range obj.Spec.Ports {
		sp := &obj.Spec.Ports[i]
		if sp.Protocol == "" {
			sp.Protocol = v1.ProtocolTCP
		}
		if sp.TargetPort == intstr.FromInt(0) || sp.TargetPort == intstr.FromString("") {
			sp.TargetPort = intstr.FromInt(int(sp.Port))
		}
	}
	if (obj.Spec.Type == v1.ServiceTypeNodePort || obj.Spec.Type == v1.ServiceTypeLoadBalancer) && obj.Spec.ExternalTrafficPolicy == "" {
		obj.Spec.ExternalTrafficPolicy = v1.ServiceExternalTrafficPolicyTypeCluster
	}
}
func SetDefaults_Pod(obj *v1.Pod) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i := range obj.Spec.Containers {
		if obj.Spec.Containers[i].Resources.Limits != nil {
			if obj.Spec.Containers[i].Resources.Requests == nil {
				obj.Spec.Containers[i].Resources.Requests = make(v1.ResourceList)
			}
			for key, value := range obj.Spec.Containers[i].Resources.Limits {
				if _, exists := obj.Spec.Containers[i].Resources.Requests[key]; !exists {
					obj.Spec.Containers[i].Resources.Requests[key] = *(value.Copy())
				}
			}
		}
	}
	for i := range obj.Spec.InitContainers {
		if obj.Spec.InitContainers[i].Resources.Limits != nil {
			if obj.Spec.InitContainers[i].Resources.Requests == nil {
				obj.Spec.InitContainers[i].Resources.Requests = make(v1.ResourceList)
			}
			for key, value := range obj.Spec.InitContainers[i].Resources.Limits {
				if _, exists := obj.Spec.InitContainers[i].Resources.Requests[key]; !exists {
					obj.Spec.InitContainers[i].Resources.Requests[key] = *(value.Copy())
				}
			}
		}
	}
	if obj.Spec.EnableServiceLinks == nil {
		enableServiceLinks := v1.DefaultEnableServiceLinks
		obj.Spec.EnableServiceLinks = &enableServiceLinks
	}
}
func SetDefaults_PodSpec(obj *v1.PodSpec) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.DNSPolicy == "" {
		obj.DNSPolicy = v1.DNSClusterFirst
	}
	if obj.RestartPolicy == "" {
		obj.RestartPolicy = v1.RestartPolicyAlways
	}
	if obj.HostNetwork {
		defaultHostNetworkPorts(&obj.Containers)
		defaultHostNetworkPorts(&obj.InitContainers)
	}
	if obj.SecurityContext == nil {
		obj.SecurityContext = &v1.PodSecurityContext{}
	}
	if obj.TerminationGracePeriodSeconds == nil {
		period := int64(v1.DefaultTerminationGracePeriodSeconds)
		obj.TerminationGracePeriodSeconds = &period
	}
	if obj.SchedulerName == "" {
		obj.SchedulerName = v1.DefaultSchedulerName
	}
}
func SetDefaults_Probe(obj *v1.Probe) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.TimeoutSeconds == 0 {
		obj.TimeoutSeconds = 1
	}
	if obj.PeriodSeconds == 0 {
		obj.PeriodSeconds = 10
	}
	if obj.SuccessThreshold == 0 {
		obj.SuccessThreshold = 1
	}
	if obj.FailureThreshold == 0 {
		obj.FailureThreshold = 3
	}
}
func SetDefaults_SecretVolumeSource(obj *v1.SecretVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.DefaultMode == nil {
		perm := int32(v1.SecretVolumeSourceDefaultMode)
		obj.DefaultMode = &perm
	}
}
func SetDefaults_ConfigMapVolumeSource(obj *v1.ConfigMapVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.DefaultMode == nil {
		perm := int32(v1.ConfigMapVolumeSourceDefaultMode)
		obj.DefaultMode = &perm
	}
}
func SetDefaults_DownwardAPIVolumeSource(obj *v1.DownwardAPIVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.DefaultMode == nil {
		perm := int32(v1.DownwardAPIVolumeSourceDefaultMode)
		obj.DefaultMode = &perm
	}
}
func SetDefaults_Secret(obj *v1.Secret) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Type == "" {
		obj.Type = v1.SecretTypeOpaque
	}
}
func SetDefaults_ProjectedVolumeSource(obj *v1.ProjectedVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.DefaultMode == nil {
		perm := int32(v1.ProjectedVolumeSourceDefaultMode)
		obj.DefaultMode = &perm
	}
}
func SetDefaults_ServiceAccountTokenProjection(obj *v1.ServiceAccountTokenProjection) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hour := int64(time.Hour.Seconds())
	if obj.ExpirationSeconds == nil {
		obj.ExpirationSeconds = &hour
	}
}
func SetDefaults_PersistentVolume(obj *v1.PersistentVolume) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Status.Phase == "" {
		obj.Status.Phase = v1.VolumePending
	}
	if obj.Spec.PersistentVolumeReclaimPolicy == "" {
		obj.Spec.PersistentVolumeReclaimPolicy = v1.PersistentVolumeReclaimRetain
	}
	if obj.Spec.VolumeMode == nil && utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
		obj.Spec.VolumeMode = new(v1.PersistentVolumeMode)
		*obj.Spec.VolumeMode = v1.PersistentVolumeFilesystem
	}
}
func SetDefaults_PersistentVolumeClaim(obj *v1.PersistentVolumeClaim) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Status.Phase == "" {
		obj.Status.Phase = v1.ClaimPending
	}
	if obj.Spec.VolumeMode == nil && utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
		obj.Spec.VolumeMode = new(v1.PersistentVolumeMode)
		*obj.Spec.VolumeMode = v1.PersistentVolumeFilesystem
	}
}
func SetDefaults_ISCSIVolumeSource(obj *v1.ISCSIVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.ISCSIInterface == "" {
		obj.ISCSIInterface = "default"
	}
}
func SetDefaults_ISCSIPersistentVolumeSource(obj *v1.ISCSIPersistentVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.ISCSIInterface == "" {
		obj.ISCSIInterface = "default"
	}
}
func SetDefaults_AzureDiskVolumeSource(obj *v1.AzureDiskVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.CachingMode == nil {
		obj.CachingMode = new(v1.AzureDataDiskCachingMode)
		*obj.CachingMode = v1.AzureDataDiskCachingReadWrite
	}
	if obj.Kind == nil {
		obj.Kind = new(v1.AzureDataDiskKind)
		*obj.Kind = v1.AzureSharedBlobDisk
	}
	if obj.FSType == nil {
		obj.FSType = new(string)
		*obj.FSType = "ext4"
	}
	if obj.ReadOnly == nil {
		obj.ReadOnly = new(bool)
		*obj.ReadOnly = false
	}
}
func SetDefaults_Endpoints(obj *v1.Endpoints) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i := range obj.Subsets {
		ss := &obj.Subsets[i]
		for i := range ss.Ports {
			ep := &ss.Ports[i]
			if ep.Protocol == "" {
				ep.Protocol = v1.ProtocolTCP
			}
		}
	}
}
func SetDefaults_HTTPGetAction(obj *v1.HTTPGetAction) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Path == "" {
		obj.Path = "/"
	}
	if obj.Scheme == "" {
		obj.Scheme = v1.URISchemeHTTP
	}
}
func SetDefaults_NamespaceStatus(obj *v1.NamespaceStatus) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Phase == "" {
		obj.Phase = v1.NamespaceActive
	}
}
func SetDefaults_NodeStatus(obj *v1.NodeStatus) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Allocatable == nil && obj.Capacity != nil {
		obj.Allocatable = make(v1.ResourceList, len(obj.Capacity))
		for key, value := range obj.Capacity {
			obj.Allocatable[key] = *(value.Copy())
		}
		obj.Allocatable = obj.Capacity
	}
}
func SetDefaults_ObjectFieldSelector(obj *v1.ObjectFieldSelector) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.APIVersion == "" {
		obj.APIVersion = "v1"
	}
}
func SetDefaults_LimitRangeItem(obj *v1.LimitRangeItem) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Type == v1.LimitTypeContainer {
		if obj.Default == nil {
			obj.Default = make(v1.ResourceList)
		}
		if obj.DefaultRequest == nil {
			obj.DefaultRequest = make(v1.ResourceList)
		}
		for key, value := range obj.Max {
			if _, exists := obj.Default[key]; !exists {
				obj.Default[key] = *(value.Copy())
			}
		}
		for key, value := range obj.Default {
			if _, exists := obj.DefaultRequest[key]; !exists {
				obj.DefaultRequest[key] = *(value.Copy())
			}
		}
		for key, value := range obj.Min {
			if _, exists := obj.DefaultRequest[key]; !exists {
				obj.DefaultRequest[key] = *(value.Copy())
			}
		}
	}
}
func SetDefaults_ConfigMap(obj *v1.ConfigMap) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Data == nil {
		obj.Data = make(map[string]string)
	}
}
func defaultHostNetworkPorts(containers *[]v1.Container) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i := range *containers {
		for j := range (*containers)[i].Ports {
			if (*containers)[i].Ports[j].HostPort == 0 {
				(*containers)[i].Ports[j].HostPort = (*containers)[i].Ports[j].ContainerPort
			}
		}
	}
}
func SetDefaults_RBDVolumeSource(obj *v1.RBDVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.RBDPool == "" {
		obj.RBDPool = "rbd"
	}
	if obj.RadosUser == "" {
		obj.RadosUser = "admin"
	}
	if obj.Keyring == "" {
		obj.Keyring = "/etc/ceph/keyring"
	}
}
func SetDefaults_RBDPersistentVolumeSource(obj *v1.RBDPersistentVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.RBDPool == "" {
		obj.RBDPool = "rbd"
	}
	if obj.RadosUser == "" {
		obj.RadosUser = "admin"
	}
	if obj.Keyring == "" {
		obj.Keyring = "/etc/ceph/keyring"
	}
}
func SetDefaults_ScaleIOVolumeSource(obj *v1.ScaleIOVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.StorageMode == "" {
		obj.StorageMode = "ThinProvisioned"
	}
	if obj.FSType == "" {
		obj.FSType = "xfs"
	}
}
func SetDefaults_ScaleIOPersistentVolumeSource(obj *v1.ScaleIOPersistentVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.StorageMode == "" {
		obj.StorageMode = "ThinProvisioned"
	}
	if obj.FSType == "" {
		obj.FSType = "xfs"
	}
}
func SetDefaults_HostPathVolumeSource(obj *v1.HostPathVolumeSource) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	typeVol := v1.HostPathUnset
	if obj.Type == nil {
		obj.Type = &typeVol
	}
}

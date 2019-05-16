package features

import (
	goformat "fmt"
	apiextensionsfeatures "k8s.io/apiextensions-apiserver/pkg/features"
	genericfeatures "k8s.io/apiserver/pkg/features"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	AppArmor                                    utilfeature.Feature = "AppArmor"
	DynamicKubeletConfig                        utilfeature.Feature = "DynamicKubeletConfig"
	ExperimentalHostUserNamespaceDefaultingGate utilfeature.Feature = "ExperimentalHostUserNamespaceDefaulting"
	ExperimentalCriticalPodAnnotation           utilfeature.Feature = "ExperimentalCriticalPodAnnotation"
	DevicePlugins                               utilfeature.Feature = "DevicePlugins"
	TaintBasedEvictions                         utilfeature.Feature = "TaintBasedEvictions"
	RotateKubeletServerCertificate              utilfeature.Feature = "RotateKubeletServerCertificate"
	RotateKubeletClientCertificate              utilfeature.Feature = "RotateKubeletClientCertificate"
	PersistentLocalVolumes                      utilfeature.Feature = "PersistentLocalVolumes"
	LocalStorageCapacityIsolation               utilfeature.Feature = "LocalStorageCapacityIsolation"
	ExpandPersistentVolumes                     utilfeature.Feature = "ExpandPersistentVolumes"
	ExpandInUsePersistentVolumes                utilfeature.Feature = "ExpandInUsePersistentVolumes"
	DebugContainers                             utilfeature.Feature = "DebugContainers"
	PodShareProcessNamespace                    utilfeature.Feature = "PodShareProcessNamespace"
	PodPriority                                 utilfeature.Feature = "PodPriority"
	EnableEquivalenceClassCache                 utilfeature.Feature = "EnableEquivalenceClassCache"
	TaintNodesByCondition                       utilfeature.Feature = "TaintNodesByCondition"
	MountPropagation                            utilfeature.Feature = "MountPropagation"
	QOSReserved                                 utilfeature.Feature = "QOSReserved"
	CPUManager                                  utilfeature.Feature = "CPUManager"
	CPUCFSQuotaPeriod                           utilfeature.Feature = "CustomCPUCFSQuotaPeriod"
	HugePages                                   utilfeature.Feature = "HugePages"
	Sysctls                                     utilfeature.Feature = "Sysctls"
	ServiceNodeExclusion                        utilfeature.Feature = "ServiceNodeExclusion"
	MountContainers                             utilfeature.Feature = "MountContainers"
	VolumeScheduling                            utilfeature.Feature = "VolumeScheduling"
	CSIPersistentVolume                         utilfeature.Feature = "CSIPersistentVolume"
	CSIDriverRegistry                           utilfeature.Feature = "CSIDriverRegistry"
	CSINodeInfo                                 utilfeature.Feature = "CSINodeInfo"
	CustomPodDNS                                utilfeature.Feature = "CustomPodDNS"
	BlockVolume                                 utilfeature.Feature = "BlockVolume"
	StorageObjectInUseProtection                utilfeature.Feature = "StorageObjectInUseProtection"
	ResourceLimitsPriorityFunction              utilfeature.Feature = "ResourceLimitsPriorityFunction"
	SupportIPVSProxyMode                        utilfeature.Feature = "SupportIPVSProxyMode"
	SupportPodPidsLimit                         utilfeature.Feature = "SupportPodPidsLimit"
	HyperVContainer                             utilfeature.Feature = "HyperVContainer"
	ScheduleDaemonSetPods                       utilfeature.Feature = "ScheduleDaemonSetPods"
	TokenRequest                                utilfeature.Feature = "TokenRequest"
	TokenRequestProjection                      utilfeature.Feature = "TokenRequestProjection"
	BoundServiceAccountTokenVolume              utilfeature.Feature = "BoundServiceAccountTokenVolume"
	CRIContainerLogRotation                     utilfeature.Feature = "CRIContainerLogRotation"
	GCERegionalPersistentDisk                   utilfeature.Feature = "GCERegionalPersistentDisk"
	RunAsGroup                                  utilfeature.Feature = "RunAsGroup"
	VolumeSubpath                               utilfeature.Feature = "VolumeSubpath"
	AttachVolumeLimit                           utilfeature.Feature = "AttachVolumeLimit"
	BalanceAttachedNodeVolumes                  utilfeature.Feature = "BalanceAttachedNodeVolumes"
	PodReadinessGates                           utilfeature.Feature = "PodReadinessGates"
	VolumeSubpathEnvExpansion                   utilfeature.Feature = "VolumeSubpathEnvExpansion"
	KubeletPluginsWatcher                       utilfeature.Feature = "KubeletPluginsWatcher"
	ResourceQuotaScopeSelectors                 utilfeature.Feature = "ResourceQuotaScopeSelectors"
	CSIBlockVolume                              utilfeature.Feature = "CSIBlockVolume"
	RuntimeClass                                utilfeature.Feature = "RuntimeClass"
	NodeLease                                   utilfeature.Feature = "NodeLease"
	SCTPSupport                                 utilfeature.Feature = "SCTPSupport"
	VolumeSnapshotDataSource                    utilfeature.Feature = "VolumeSnapshotDataSource"
	ProcMountType                               utilfeature.Feature = "ProcMountType"
	TTLAfterFinished                            utilfeature.Feature = "TTLAfterFinished"
	KubeletPodResources                         utilfeature.Feature = "KubeletPodResources"
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	utilfeature.DefaultFeatureGate.Add(defaultKubernetesFeatureGates)
}

var defaultKubernetesFeatureGates = map[utilfeature.Feature]utilfeature.FeatureSpec{AppArmor: {Default: true, PreRelease: utilfeature.Beta}, DynamicKubeletConfig: {Default: true, PreRelease: utilfeature.Beta}, ExperimentalHostUserNamespaceDefaultingGate: {Default: false, PreRelease: utilfeature.Beta}, ExperimentalCriticalPodAnnotation: {Default: false, PreRelease: utilfeature.Alpha}, DevicePlugins: {Default: true, PreRelease: utilfeature.Beta}, TaintBasedEvictions: {Default: true, PreRelease: utilfeature.Beta}, RotateKubeletServerCertificate: {Default: true, PreRelease: utilfeature.Beta}, RotateKubeletClientCertificate: {Default: true, PreRelease: utilfeature.Beta}, PersistentLocalVolumes: {Default: true, PreRelease: utilfeature.Beta}, LocalStorageCapacityIsolation: {Default: true, PreRelease: utilfeature.Beta}, HugePages: {Default: true, PreRelease: utilfeature.Beta}, Sysctls: {Default: true, PreRelease: utilfeature.Beta}, DebugContainers: {Default: false, PreRelease: utilfeature.Alpha}, PodShareProcessNamespace: {Default: true, PreRelease: utilfeature.Beta}, PodPriority: {Default: true, PreRelease: utilfeature.Beta}, EnableEquivalenceClassCache: {Default: false, PreRelease: utilfeature.Alpha}, TaintNodesByCondition: {Default: true, PreRelease: utilfeature.Beta}, MountPropagation: {Default: true, PreRelease: utilfeature.GA}, QOSReserved: {Default: false, PreRelease: utilfeature.Alpha}, ExpandPersistentVolumes: {Default: true, PreRelease: utilfeature.Beta}, ExpandInUsePersistentVolumes: {Default: false, PreRelease: utilfeature.Alpha}, AttachVolumeLimit: {Default: true, PreRelease: utilfeature.Beta}, CPUManager: {Default: true, PreRelease: utilfeature.Beta}, CPUCFSQuotaPeriod: {Default: false, PreRelease: utilfeature.Alpha}, ServiceNodeExclusion: {Default: false, PreRelease: utilfeature.Alpha}, MountContainers: {Default: false, PreRelease: utilfeature.Alpha}, VolumeScheduling: {Default: true, PreRelease: utilfeature.GA}, CSIPersistentVolume: {Default: true, PreRelease: utilfeature.GA}, CSIDriverRegistry: {Default: false, PreRelease: utilfeature.Alpha}, CSINodeInfo: {Default: false, PreRelease: utilfeature.Alpha}, CustomPodDNS: {Default: true, PreRelease: utilfeature.Beta}, BlockVolume: {Default: true, PreRelease: utilfeature.Beta}, StorageObjectInUseProtection: {Default: true, PreRelease: utilfeature.GA}, ResourceLimitsPriorityFunction: {Default: false, PreRelease: utilfeature.Alpha}, SupportIPVSProxyMode: {Default: true, PreRelease: utilfeature.GA}, SupportPodPidsLimit: {Default: false, PreRelease: utilfeature.Alpha}, HyperVContainer: {Default: false, PreRelease: utilfeature.Alpha}, ScheduleDaemonSetPods: {Default: true, PreRelease: utilfeature.Beta}, TokenRequest: {Default: true, PreRelease: utilfeature.Beta}, TokenRequestProjection: {Default: true, PreRelease: utilfeature.Beta}, BoundServiceAccountTokenVolume: {Default: false, PreRelease: utilfeature.Alpha}, CRIContainerLogRotation: {Default: true, PreRelease: utilfeature.Beta}, GCERegionalPersistentDisk: {Default: true, PreRelease: utilfeature.GA}, RunAsGroup: {Default: false, PreRelease: utilfeature.Alpha}, VolumeSubpath: {Default: true, PreRelease: utilfeature.GA}, BalanceAttachedNodeVolumes: {Default: false, PreRelease: utilfeature.Alpha}, PodReadinessGates: {Default: true, PreRelease: utilfeature.Beta}, VolumeSubpathEnvExpansion: {Default: false, PreRelease: utilfeature.Alpha}, KubeletPluginsWatcher: {Default: true, PreRelease: utilfeature.GA}, ResourceQuotaScopeSelectors: {Default: true, PreRelease: utilfeature.Beta}, CSIBlockVolume: {Default: false, PreRelease: utilfeature.Alpha}, RuntimeClass: {Default: false, PreRelease: utilfeature.Alpha}, NodeLease: {Default: false, PreRelease: utilfeature.Alpha}, SCTPSupport: {Default: false, PreRelease: utilfeature.Alpha}, VolumeSnapshotDataSource: {Default: false, PreRelease: utilfeature.Alpha}, ProcMountType: {Default: false, PreRelease: utilfeature.Alpha}, TTLAfterFinished: {Default: false, PreRelease: utilfeature.Alpha}, KubeletPodResources: {Default: false, PreRelease: utilfeature.Alpha}, genericfeatures.StreamingProxyRedirects: {Default: true, PreRelease: utilfeature.Beta}, genericfeatures.AdvancedAuditing: {Default: true, PreRelease: utilfeature.GA}, genericfeatures.DynamicAuditing: {Default: false, PreRelease: utilfeature.Alpha}, genericfeatures.APIResponseCompression: {Default: false, PreRelease: utilfeature.Alpha}, genericfeatures.Initializers: {Default: false, PreRelease: utilfeature.Alpha}, genericfeatures.APIListChunking: {Default: true, PreRelease: utilfeature.Beta}, genericfeatures.DryRun: {Default: true, PreRelease: utilfeature.Beta}, apiextensionsfeatures.CustomResourceValidation: {Default: true, PreRelease: utilfeature.Beta}, apiextensionsfeatures.CustomResourceSubresources: {Default: true, PreRelease: utilfeature.Beta}, apiextensionsfeatures.CustomResourceWebhookConversion: {Default: false, PreRelease: utilfeature.Alpha}, apiextensionsfeatures.CustomResourcePublishOpenAPI: {Default: true, PreRelease: utilfeature.Alpha}}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

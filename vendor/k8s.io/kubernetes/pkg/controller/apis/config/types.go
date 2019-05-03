package config

import (
 apimachineryconfig "k8s.io/apimachinery/pkg/apis/config"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 apiserverconfig "k8s.io/apiserver/pkg/apis/config"
)

type GroupResource struct {
 Group    string
 Resource string
}
type KubeControllerManagerConfiguration struct {
 metav1.TypeMeta
 Generic                          GenericControllerManagerConfiguration
 KubeCloudShared                  KubeCloudSharedConfiguration
 AttachDetachController           AttachDetachControllerConfiguration
 CSRSigningController             CSRSigningControllerConfiguration
 DaemonSetController              DaemonSetControllerConfiguration
 DeploymentController             DeploymentControllerConfiguration
 DeprecatedController             DeprecatedControllerConfiguration
 EndpointController               EndpointControllerConfiguration
 GarbageCollectorController       GarbageCollectorControllerConfiguration
 HPAController                    HPAControllerConfiguration
 JobController                    JobControllerConfiguration
 NamespaceController              NamespaceControllerConfiguration
 NodeIPAMController               NodeIPAMControllerConfiguration
 NodeLifecycleController          NodeLifecycleControllerConfiguration
 PersistentVolumeBinderController PersistentVolumeBinderControllerConfiguration
 PodGCController                  PodGCControllerConfiguration
 ReplicaSetController             ReplicaSetControllerConfiguration
 ReplicationController            ReplicationControllerConfiguration
 ResourceQuotaController          ResourceQuotaControllerConfiguration
 SAController                     SAControllerConfiguration
 ServiceController                ServiceControllerConfiguration
 TTLAfterFinishedController       TTLAfterFinishedControllerConfiguration
}
type GenericControllerManagerConfiguration struct {
 Port                    int32
 Address                 string
 MinResyncPeriod         metav1.Duration
 ClientConnection        apimachineryconfig.ClientConnectionConfiguration
 ControllerStartInterval metav1.Duration
 LeaderElection          apiserverconfig.LeaderElectionConfiguration
 Controllers             []string
 Debugging               apiserverconfig.DebuggingConfiguration
}
type KubeCloudSharedConfiguration struct {
 CloudProvider                CloudProviderConfiguration
 ExternalCloudVolumePlugin    string
 UseServiceAccountCredentials bool
 AllowUntaggedCloud           bool
 RouteReconciliationPeriod    metav1.Duration
 NodeMonitorPeriod            metav1.Duration
 ClusterName                  string
 ClusterCIDR                  string
 AllocateNodeCIDRs            bool
 CIDRAllocatorType            string
 ConfigureCloudRoutes         bool
 NodeSyncPeriod               metav1.Duration
}
type AttachDetachControllerConfiguration struct {
 DisableAttachDetachReconcilerSync bool
 ReconcilerSyncLoopPeriod          metav1.Duration
}
type CloudProviderConfiguration struct {
 Name            string
 CloudConfigFile string
}
type CSRSigningControllerConfiguration struct {
 ClusterSigningCertFile string
 ClusterSigningKeyFile  string
 ClusterSigningDuration metav1.Duration
}
type DaemonSetControllerConfiguration struct{ ConcurrentDaemonSetSyncs int32 }
type DeploymentControllerConfiguration struct {
 ConcurrentDeploymentSyncs      int32
 DeploymentControllerSyncPeriod metav1.Duration
}
type DeprecatedControllerConfiguration struct {
 DeletingPodsQPS    float32
 DeletingPodsBurst  int32
 RegisterRetryCount int32
}
type EndpointControllerConfiguration struct{ ConcurrentEndpointSyncs int32 }
type GarbageCollectorControllerConfiguration struct {
 EnableGarbageCollector bool
 ConcurrentGCSyncs      int32
 GCIgnoredResources     []GroupResource
}
type HPAControllerConfiguration struct {
 HorizontalPodAutoscalerSyncPeriod                   metav1.Duration
 HorizontalPodAutoscalerUpscaleForbiddenWindow       metav1.Duration
 HorizontalPodAutoscalerDownscaleForbiddenWindow     metav1.Duration
 HorizontalPodAutoscalerDownscaleStabilizationWindow metav1.Duration
 HorizontalPodAutoscalerTolerance                    float64
 HorizontalPodAutoscalerUseRESTClients               bool
 HorizontalPodAutoscalerCPUInitializationPeriod      metav1.Duration
 HorizontalPodAutoscalerInitialReadinessDelay        metav1.Duration
}
type JobControllerConfiguration struct{ ConcurrentJobSyncs int32 }
type NamespaceControllerConfiguration struct {
 NamespaceSyncPeriod      metav1.Duration
 ConcurrentNamespaceSyncs int32
}
type NodeIPAMControllerConfiguration struct {
 ServiceCIDR      string
 NodeCIDRMaskSize int32
}
type NodeLifecycleControllerConfiguration struct {
 EnableTaintManager        bool
 NodeEvictionRate          float32
 SecondaryNodeEvictionRate float32
 NodeStartupGracePeriod    metav1.Duration
 NodeMonitorGracePeriod    metav1.Duration
 PodEvictionTimeout        metav1.Duration
 LargeClusterSizeThreshold int32
 UnhealthyZoneThreshold    float32
}
type PersistentVolumeBinderControllerConfiguration struct {
 PVClaimBinderSyncPeriod metav1.Duration
 VolumeConfiguration     VolumeConfiguration
}
type PodGCControllerConfiguration struct{ TerminatedPodGCThreshold int32 }
type ReplicaSetControllerConfiguration struct{ ConcurrentRSSyncs int32 }
type ReplicationControllerConfiguration struct{ ConcurrentRCSyncs int32 }
type ResourceQuotaControllerConfiguration struct {
 ResourceQuotaSyncPeriod      metav1.Duration
 ConcurrentResourceQuotaSyncs int32
}
type SAControllerConfiguration struct {
 ServiceAccountKeyFile  string
 ConcurrentSATokenSyncs int32
 RootCAFile             string
}
type ServiceControllerConfiguration struct{ ConcurrentServiceSyncs int32 }
type VolumeConfiguration struct {
 EnableHostPathProvisioning            bool
 EnableDynamicProvisioning             bool
 PersistentVolumeRecyclerConfiguration PersistentVolumeRecyclerConfiguration
 FlexVolumePluginDir                   string
}
type PersistentVolumeRecyclerConfiguration struct {
 MaximumRetry                int32
 MinimumTimeoutNFS           int32
 PodTemplateFilePathNFS      string
 IncrementTimeoutNFS         int32
 PodTemplateFilePathHostPath string
 MinimumTimeoutHostPath      int32
 IncrementTimeoutHostPath    int32
}
type TTLAfterFinishedControllerConfiguration struct{ ConcurrentTTLSyncs int32 }

package policy

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/util/intstr"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type PodDisruptionBudgetSpec struct {
 MinAvailable   *intstr.IntOrString
 Selector       *metav1.LabelSelector
 MaxUnavailable *intstr.IntOrString
}
type PodDisruptionBudgetStatus struct {
 ObservedGeneration    int64
 DisruptedPods         map[string]metav1.Time
 PodDisruptionsAllowed int32
 CurrentHealthy        int32
 DesiredHealthy        int32
 ExpectedPods          int32
}
type PodDisruptionBudget struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   PodDisruptionBudgetSpec
 Status PodDisruptionBudgetStatus
}
type PodDisruptionBudgetList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []PodDisruptionBudget
}
type Eviction struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 DeleteOptions *metav1.DeleteOptions
}
type PodSecurityPolicy struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec PodSecurityPolicySpec
}
type PodSecurityPolicySpec struct {
 Privileged                      bool
 DefaultAddCapabilities          []api.Capability
 RequiredDropCapabilities        []api.Capability
 AllowedCapabilities             []api.Capability
 Volumes                         []FSType
 HostNetwork                     bool
 HostPorts                       []HostPortRange
 HostPID                         bool
 HostIPC                         bool
 SELinux                         SELinuxStrategyOptions
 RunAsUser                       RunAsUserStrategyOptions
 RunAsGroup                      *RunAsGroupStrategyOptions
 SupplementalGroups              SupplementalGroupsStrategyOptions
 FSGroup                         FSGroupStrategyOptions
 ReadOnlyRootFilesystem          bool
 DefaultAllowPrivilegeEscalation *bool
 AllowPrivilegeEscalation        bool
 AllowedHostPaths                []AllowedHostPath
 AllowedFlexVolumes              []AllowedFlexVolume
 AllowedUnsafeSysctls            []string
 ForbiddenSysctls                []string
 AllowedProcMountTypes           []api.ProcMountType
}
type AllowedHostPath struct {
 PathPrefix string
 ReadOnly   bool
}
type HostPortRange struct {
 Min int32
 Max int32
}

var AllowAllCapabilities api.Capability = "*"

type FSType string

var (
 AzureFile             FSType = "azureFile"
 Flocker               FSType = "flocker"
 FlexVolume            FSType = "flexVolume"
 HostPath              FSType = "hostPath"
 EmptyDir              FSType = "emptyDir"
 GCEPersistentDisk     FSType = "gcePersistentDisk"
 AWSElasticBlockStore  FSType = "awsElasticBlockStore"
 GitRepo               FSType = "gitRepo"
 Secret                FSType = "secret"
 NFS                   FSType = "nfs"
 ISCSI                 FSType = "iscsi"
 Glusterfs             FSType = "glusterfs"
 PersistentVolumeClaim FSType = "persistentVolumeClaim"
 RBD                   FSType = "rbd"
 Cinder                FSType = "cinder"
 CephFS                FSType = "cephFS"
 DownwardAPI           FSType = "downwardAPI"
 FC                    FSType = "fc"
 ConfigMap             FSType = "configMap"
 VsphereVolume         FSType = "vsphereVolume"
 Quobyte               FSType = "quobyte"
 AzureDisk             FSType = "azureDisk"
 PhotonPersistentDisk  FSType = "photonPersistentDisk"
 StorageOS             FSType = "storageos"
 Projected             FSType = "projected"
 PortworxVolume        FSType = "portworxVolume"
 ScaleIO               FSType = "scaleIO"
 CSI                   FSType = "csi"
 All                   FSType = "*"
)

type AllowedFlexVolume struct{ Driver string }
type SELinuxStrategyOptions struct {
 Rule           SELinuxStrategy
 SELinuxOptions *api.SELinuxOptions
}
type SELinuxStrategy string

const (
 SELinuxStrategyMustRunAs SELinuxStrategy = "MustRunAs"
 SELinuxStrategyRunAsAny  SELinuxStrategy = "RunAsAny"
)

type RunAsUserStrategyOptions struct {
 Rule   RunAsUserStrategy
 Ranges []IDRange
}
type RunAsGroupStrategyOptions struct {
 Rule   RunAsGroupStrategy
 Ranges []IDRange
}
type IDRange struct {
 Min int64
 Max int64
}
type RunAsUserStrategy string

const (
 RunAsUserStrategyMustRunAs        RunAsUserStrategy = "MustRunAs"
 RunAsUserStrategyMustRunAsNonRoot RunAsUserStrategy = "MustRunAsNonRoot"
 RunAsUserStrategyRunAsAny         RunAsUserStrategy = "RunAsAny"
)

type RunAsGroupStrategy string

const (
 RunAsGroupStrategyMayRunAs  RunAsGroupStrategy = "MayRunAs"
 RunAsGroupStrategyMustRunAs RunAsGroupStrategy = "MustRunAs"
 RunAsGroupStrategyRunAsAny  RunAsGroupStrategy = "RunAsAny"
)

type FSGroupStrategyOptions struct {
 Rule   FSGroupStrategyType
 Ranges []IDRange
}
type FSGroupStrategyType string

const (
 FSGroupStrategyMayRunAs  FSGroupStrategyType = "MayRunAs"
 FSGroupStrategyMustRunAs FSGroupStrategyType = "MustRunAs"
 FSGroupStrategyRunAsAny  FSGroupStrategyType = "RunAsAny"
)

type SupplementalGroupsStrategyOptions struct {
 Rule   SupplementalGroupsStrategyType
 Ranges []IDRange
}
type SupplementalGroupsStrategyType string

const (
 SupplementalGroupsStrategyMayRunAs  SupplementalGroupsStrategyType = "MayRunAs"
 SupplementalGroupsStrategyMustRunAs SupplementalGroupsStrategyType = "MustRunAs"
 SupplementalGroupsStrategyRunAsAny  SupplementalGroupsStrategyType = "RunAsAny"
)

type PodSecurityPolicyList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []PodSecurityPolicy
}

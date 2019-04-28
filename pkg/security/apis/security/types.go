package security

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/core"
)

var AllowAllCapabilities core.Capability = "*"

type SecurityContextConstraints struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Priority			*int32
	AllowPrivilegedContainer	bool
	DefaultAddCapabilities		[]core.Capability
	RequiredDropCapabilities	[]core.Capability
	AllowedCapabilities		[]core.Capability
	Volumes				[]FSType
	AllowedFlexVolumes		[]AllowedFlexVolume
	AllowHostNetwork		bool
	AllowHostPorts			bool
	AllowHostPID			bool
	AllowHostIPC			bool
	DefaultAllowPrivilegeEscalation	*bool
	AllowPrivilegeEscalation	*bool
	SELinuxContext			SELinuxContextStrategyOptions
	RunAsUser			RunAsUserStrategyOptions
	SupplementalGroups		SupplementalGroupsStrategyOptions
	FSGroup				FSGroupStrategyOptions
	ReadOnlyRootFilesystem		bool
	SeccompProfiles			[]string
	Users				[]string
	Groups				[]string
	AllowedUnsafeSysctls		[]string
	ForbiddenSysctls		[]string
}
type FSType string

var (
	FSTypeAzureFile			FSType	= "azureFile"
	FSTypeAzureDisk			FSType	= "azureDisk"
	FSTypeFlocker			FSType	= "flocker"
	FSTypeFlexVolume		FSType	= "flexVolume"
	FSTypeHostPath			FSType	= "hostPath"
	FSTypeEmptyDir			FSType	= "emptyDir"
	FSTypeGCEPersistentDisk		FSType	= "gcePersistentDisk"
	FSTypeAWSElasticBlockStore	FSType	= "awsElasticBlockStore"
	FSTypeGitRepo			FSType	= "gitRepo"
	FSTypeSecret			FSType	= "secret"
	FSTypeNFS			FSType	= "nfs"
	FSTypeISCSI			FSType	= "iscsi"
	FSTypeGlusterfs			FSType	= "glusterfs"
	FSTypePersistentVolumeClaim	FSType	= "persistentVolumeClaim"
	FSTypeRBD			FSType	= "rbd"
	FSTypeCinder			FSType	= "cinder"
	FSTypeCephFS			FSType	= "cephFS"
	FSTypeDownwardAPI		FSType	= "downwardAPI"
	FSTypeFC			FSType	= "fc"
	FSTypeConfigMap			FSType	= "configMap"
	FSTypeVsphereVolume		FSType	= "vsphere"
	FSTypeQuobyte			FSType	= "quobyte"
	FSTypePhotonPersistentDisk	FSType	= "photonPersistentDisk"
	FSProjected			FSType	= "projected"
	FSPortworxVolume		FSType	= "portworxVolume"
	FSScaleIO			FSType	= "scaleIO"
	FSStorageOS			FSType	= "storageOS"
	FSTypeAll			FSType	= "*"
	FSTypeNone			FSType	= "none"
)

type AllowedFlexVolume struct{ Driver string }
type SELinuxContextStrategyOptions struct {
	Type		SELinuxContextStrategyType
	SELinuxOptions	*core.SELinuxOptions
}
type RunAsUserStrategyOptions struct {
	Type		RunAsUserStrategyType
	UID		*int64
	UIDRangeMin	*int64
	UIDRangeMax	*int64
}
type FSGroupStrategyOptions struct {
	Type	FSGroupStrategyType
	Ranges	[]IDRange
}
type SupplementalGroupsStrategyOptions struct {
	Type	SupplementalGroupsStrategyType
	Ranges	[]IDRange
}
type IDRange struct {
	Min	int64
	Max	int64
}
type SELinuxContextStrategyType string
type RunAsUserStrategyType string
type SupplementalGroupsStrategyType string
type FSGroupStrategyType string

const (
	SELinuxStrategyMustRunAs		SELinuxContextStrategyType	= "MustRunAs"
	SELinuxStrategyRunAsAny			SELinuxContextStrategyType	= "RunAsAny"
	RunAsUserStrategyMustRunAs		RunAsUserStrategyType		= "MustRunAs"
	RunAsUserStrategyMustRunAsRange		RunAsUserStrategyType		= "MustRunAsRange"
	RunAsUserStrategyMustRunAsNonRoot	RunAsUserStrategyType		= "MustRunAsNonRoot"
	RunAsUserStrategyRunAsAny		RunAsUserStrategyType		= "RunAsAny"
	FSGroupStrategyMustRunAs		FSGroupStrategyType		= "MustRunAs"
	FSGroupStrategyRunAsAny			FSGroupStrategyType		= "RunAsAny"
	SupplementalGroupsStrategyMustRunAs	SupplementalGroupsStrategyType	= "MustRunAs"
	SupplementalGroupsStrategyRunAsAny	SupplementalGroupsStrategyType	= "RunAsAny"
)

type SecurityContextConstraintsList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]SecurityContextConstraints
}
type PodSecurityPolicySubjectReview struct {
	metav1.TypeMeta
	Spec	PodSecurityPolicySubjectReviewSpec
	Status	PodSecurityPolicySubjectReviewStatus
}
type PodSecurityPolicySubjectReviewSpec struct {
	Template	core.PodTemplateSpec
	User		string
	Groups		[]string
}
type PodSecurityPolicySubjectReviewStatus struct {
	AllowedBy	*core.ObjectReference
	Reason		string
	Template	core.PodTemplateSpec
}
type PodSecurityPolicySelfSubjectReview struct {
	metav1.TypeMeta
	Spec	PodSecurityPolicySelfSubjectReviewSpec
	Status	PodSecurityPolicySubjectReviewStatus
}
type PodSecurityPolicySelfSubjectReviewSpec struct{ Template core.PodTemplateSpec }
type PodSecurityPolicyReview struct {
	metav1.TypeMeta
	Spec	PodSecurityPolicyReviewSpec
	Status	PodSecurityPolicyReviewStatus
}
type PodSecurityPolicyReviewSpec struct {
	Template		core.PodTemplateSpec
	ServiceAccountNames	[]string
}
type PodSecurityPolicyReviewStatus struct {
	AllowedServiceAccounts []ServiceAccountPodSecurityPolicyReviewStatus
}
type ServiceAccountPodSecurityPolicyReviewStatus struct {
	PodSecurityPolicySubjectReviewStatus
	Name	string
}
type RangeAllocation struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Range	string
	Data	[]byte
}
type RangeAllocationList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]RangeAllocation
}

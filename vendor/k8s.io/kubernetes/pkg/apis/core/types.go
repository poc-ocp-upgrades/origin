package core

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	NamespaceDefault              = "default"
	NamespaceAll                  = ""
	NamespaceNone                 = ""
	NamespaceSystem               = "kube-system"
	NamespacePublic               = "kube-public"
	NamespaceNodeLease            = "kube-node-lease"
	TerminationMessagePathDefault = "/dev/termination-log"
)

type Volume struct {
	Name string
	VolumeSource
}
type VolumeSource struct {
	HostPath              *HostPathVolumeSource
	EmptyDir              *EmptyDirVolumeSource
	GCEPersistentDisk     *GCEPersistentDiskVolumeSource
	AWSElasticBlockStore  *AWSElasticBlockStoreVolumeSource
	GitRepo               *GitRepoVolumeSource
	Secret                *SecretVolumeSource
	NFS                   *NFSVolumeSource
	ISCSI                 *ISCSIVolumeSource
	Glusterfs             *GlusterfsVolumeSource
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource
	RBD                   *RBDVolumeSource
	Quobyte               *QuobyteVolumeSource
	FlexVolume            *FlexVolumeSource
	Cinder                *CinderVolumeSource
	CephFS                *CephFSVolumeSource
	Flocker               *FlockerVolumeSource
	DownwardAPI           *DownwardAPIVolumeSource
	FC                    *FCVolumeSource
	AzureFile             *AzureFileVolumeSource
	ConfigMap             *ConfigMapVolumeSource
	VsphereVolume         *VsphereVirtualDiskVolumeSource
	AzureDisk             *AzureDiskVolumeSource
	PhotonPersistentDisk  *PhotonPersistentDiskVolumeSource
	Projected             *ProjectedVolumeSource
	PortworxVolume        *PortworxVolumeSource
	ScaleIO               *ScaleIOVolumeSource
	StorageOS             *StorageOSVolumeSource
}
type PersistentVolumeSource struct {
	GCEPersistentDisk    *GCEPersistentDiskVolumeSource
	AWSElasticBlockStore *AWSElasticBlockStoreVolumeSource
	HostPath             *HostPathVolumeSource
	Glusterfs            *GlusterfsPersistentVolumeSource
	NFS                  *NFSVolumeSource
	RBD                  *RBDPersistentVolumeSource
	Quobyte              *QuobyteVolumeSource
	ISCSI                *ISCSIPersistentVolumeSource
	FlexVolume           *FlexPersistentVolumeSource
	Cinder               *CinderPersistentVolumeSource
	CephFS               *CephFSPersistentVolumeSource
	FC                   *FCVolumeSource
	Flocker              *FlockerVolumeSource
	AzureFile            *AzureFilePersistentVolumeSource
	VsphereVolume        *VsphereVirtualDiskVolumeSource
	AzureDisk            *AzureDiskVolumeSource
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSource
	PortworxVolume       *PortworxVolumeSource
	ScaleIO              *ScaleIOPersistentVolumeSource
	Local                *LocalVolumeSource
	StorageOS            *StorageOSPersistentVolumeSource
	CSI                  *CSIPersistentVolumeSource
}
type PersistentVolumeClaimVolumeSource struct {
	ClaimName string
	ReadOnly  bool
}

const (
	BetaStorageClassAnnotation = "volume.beta.kubernetes.io/storage-class"
	MountOptionAnnotation      = "volume.beta.kubernetes.io/mount-options"
)

type PersistentVolume struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   PersistentVolumeSpec
	Status PersistentVolumeStatus
}
type PersistentVolumeSpec struct {
	Capacity ResourceList
	PersistentVolumeSource
	AccessModes                   []PersistentVolumeAccessMode
	ClaimRef                      *ObjectReference
	PersistentVolumeReclaimPolicy PersistentVolumeReclaimPolicy
	StorageClassName              string
	MountOptions                  []string
	VolumeMode                    *PersistentVolumeMode
	NodeAffinity                  *VolumeNodeAffinity
}
type VolumeNodeAffinity struct{ Required *NodeSelector }
type PersistentVolumeReclaimPolicy string

const (
	PersistentVolumeReclaimRecycle PersistentVolumeReclaimPolicy = "Recycle"
	PersistentVolumeReclaimDelete  PersistentVolumeReclaimPolicy = "Delete"
	PersistentVolumeReclaimRetain  PersistentVolumeReclaimPolicy = "Retain"
)

type PersistentVolumeMode string

const (
	PersistentVolumeBlock      PersistentVolumeMode = "Block"
	PersistentVolumeFilesystem PersistentVolumeMode = "Filesystem"
)

type PersistentVolumeStatus struct {
	Phase   PersistentVolumePhase
	Message string
	Reason  string
}
type PersistentVolumeList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []PersistentVolume
}
type PersistentVolumeClaim struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   PersistentVolumeClaimSpec
	Status PersistentVolumeClaimStatus
}
type PersistentVolumeClaimList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []PersistentVolumeClaim
}
type PersistentVolumeClaimSpec struct {
	AccessModes      []PersistentVolumeAccessMode
	Selector         *metav1.LabelSelector
	Resources        ResourceRequirements
	VolumeName       string
	StorageClassName *string
	VolumeMode       *PersistentVolumeMode
	DataSource       *TypedLocalObjectReference
}
type PersistentVolumeClaimConditionType string

const (
	PersistentVolumeClaimResizing                PersistentVolumeClaimConditionType = "Resizing"
	PersistentVolumeClaimFileSystemResizePending PersistentVolumeClaimConditionType = "FileSystemResizePending"
)

type PersistentVolumeClaimCondition struct {
	Type               PersistentVolumeClaimConditionType
	Status             ConditionStatus
	LastProbeTime      metav1.Time
	LastTransitionTime metav1.Time
	Reason             string
	Message            string
}
type PersistentVolumeClaimStatus struct {
	Phase       PersistentVolumeClaimPhase
	AccessModes []PersistentVolumeAccessMode
	Capacity    ResourceList
	Conditions  []PersistentVolumeClaimCondition
}
type PersistentVolumeAccessMode string

const (
	ReadWriteOnce PersistentVolumeAccessMode = "ReadWriteOnce"
	ReadOnlyMany  PersistentVolumeAccessMode = "ReadOnlyMany"
	ReadWriteMany PersistentVolumeAccessMode = "ReadWriteMany"
)

type PersistentVolumePhase string

const (
	VolumePending   PersistentVolumePhase = "Pending"
	VolumeAvailable PersistentVolumePhase = "Available"
	VolumeBound     PersistentVolumePhase = "Bound"
	VolumeReleased  PersistentVolumePhase = "Released"
	VolumeFailed    PersistentVolumePhase = "Failed"
)

type PersistentVolumeClaimPhase string

const (
	ClaimPending PersistentVolumeClaimPhase = "Pending"
	ClaimBound   PersistentVolumeClaimPhase = "Bound"
	ClaimLost    PersistentVolumeClaimPhase = "Lost"
)

type HostPathType string

const (
	HostPathUnset             HostPathType = ""
	HostPathDirectoryOrCreate HostPathType = "DirectoryOrCreate"
	HostPathDirectory         HostPathType = "Directory"
	HostPathFileOrCreate      HostPathType = "FileOrCreate"
	HostPathFile              HostPathType = "File"
	HostPathSocket            HostPathType = "Socket"
	HostPathCharDev           HostPathType = "CharDevice"
	HostPathBlockDev          HostPathType = "BlockDevice"
)

type HostPathVolumeSource struct {
	Path string
	Type *HostPathType
}
type EmptyDirVolumeSource struct {
	Medium    StorageMedium
	SizeLimit *resource.Quantity
}
type StorageMedium string

const (
	StorageMediumDefault   StorageMedium = ""
	StorageMediumMemory    StorageMedium = "Memory"
	StorageMediumHugePages StorageMedium = "HugePages"
)

type Protocol string

const (
	ProtocolTCP  Protocol = "TCP"
	ProtocolUDP  Protocol = "UDP"
	ProtocolSCTP Protocol = "SCTP"
)

type GCEPersistentDiskVolumeSource struct {
	PDName    string
	FSType    string
	Partition int32
	ReadOnly  bool
}
type ISCSIVolumeSource struct {
	TargetPortal      string
	IQN               string
	Lun               int32
	ISCSIInterface    string
	FSType            string
	ReadOnly          bool
	Portals           []string
	DiscoveryCHAPAuth bool
	SessionCHAPAuth   bool
	SecretRef         *LocalObjectReference
	InitiatorName     *string
}
type ISCSIPersistentVolumeSource struct {
	TargetPortal      string
	IQN               string
	Lun               int32
	ISCSIInterface    string
	FSType            string
	ReadOnly          bool
	Portals           []string
	DiscoveryCHAPAuth bool
	SessionCHAPAuth   bool
	SecretRef         *SecretReference
	InitiatorName     *string
}
type FCVolumeSource struct {
	TargetWWNs []string
	Lun        *int32
	FSType     string
	ReadOnly   bool
	WWIDs      []string
}
type FlexPersistentVolumeSource struct {
	Driver    string
	FSType    string
	SecretRef *SecretReference
	ReadOnly  bool
	Options   map[string]string
}
type FlexVolumeSource struct {
	Driver    string
	FSType    string
	SecretRef *LocalObjectReference
	ReadOnly  bool
	Options   map[string]string
}
type AWSElasticBlockStoreVolumeSource struct {
	VolumeID  string
	FSType    string
	Partition int32
	ReadOnly  bool
}
type GitRepoVolumeSource struct {
	Repository string
	Revision   string
	Directory  string
}
type SecretVolumeSource struct {
	SecretName  string
	Items       []KeyToPath
	DefaultMode *int32
	Optional    *bool
}
type SecretProjection struct {
	LocalObjectReference
	Items    []KeyToPath
	Optional *bool
}
type NFSVolumeSource struct {
	Server   string
	Path     string
	ReadOnly bool
}
type QuobyteVolumeSource struct {
	Registry string
	Volume   string
	ReadOnly bool
	User     string
	Group    string
}
type GlusterfsVolumeSource struct {
	EndpointsName string
	Path          string
	ReadOnly      bool
}
type GlusterfsPersistentVolumeSource struct {
	EndpointsName      string
	Path               string
	ReadOnly           bool
	EndpointsNamespace *string
}
type RBDVolumeSource struct {
	CephMonitors []string
	RBDImage     string
	FSType       string
	RBDPool      string
	RadosUser    string
	Keyring      string
	SecretRef    *LocalObjectReference
	ReadOnly     bool
}
type RBDPersistentVolumeSource struct {
	CephMonitors []string
	RBDImage     string
	FSType       string
	RBDPool      string
	RadosUser    string
	Keyring      string
	SecretRef    *SecretReference
	ReadOnly     bool
}
type CinderVolumeSource struct {
	VolumeID  string
	FSType    string
	ReadOnly  bool
	SecretRef *LocalObjectReference
}
type CinderPersistentVolumeSource struct {
	VolumeID  string
	FSType    string
	ReadOnly  bool
	SecretRef *SecretReference
}
type CephFSVolumeSource struct {
	Monitors   []string
	Path       string
	User       string
	SecretFile string
	SecretRef  *LocalObjectReference
	ReadOnly   bool
}
type SecretReference struct {
	Name      string
	Namespace string
}
type CephFSPersistentVolumeSource struct {
	Monitors   []string
	Path       string
	User       string
	SecretFile string
	SecretRef  *SecretReference
	ReadOnly   bool
}
type FlockerVolumeSource struct {
	DatasetName string
	DatasetUUID string
}
type DownwardAPIVolumeSource struct {
	Items       []DownwardAPIVolumeFile
	DefaultMode *int32
}
type DownwardAPIVolumeFile struct {
	Path             string
	FieldRef         *ObjectFieldSelector
	ResourceFieldRef *ResourceFieldSelector
	Mode             *int32
}
type DownwardAPIProjection struct{ Items []DownwardAPIVolumeFile }
type AzureFileVolumeSource struct {
	SecretName string
	ShareName  string
	ReadOnly   bool
}
type AzureFilePersistentVolumeSource struct {
	SecretName      string
	ShareName       string
	ReadOnly        bool
	SecretNamespace *string
}
type VsphereVirtualDiskVolumeSource struct {
	VolumePath        string
	FSType            string
	StoragePolicyName string
	StoragePolicyID   string
}
type PhotonPersistentDiskVolumeSource struct {
	PdID   string
	FSType string
}
type PortworxVolumeSource struct {
	VolumeID string
	FSType   string
	ReadOnly bool
}
type AzureDataDiskCachingMode string
type AzureDataDiskKind string

const (
	AzureDataDiskCachingNone      AzureDataDiskCachingMode = "None"
	AzureDataDiskCachingReadOnly  AzureDataDiskCachingMode = "ReadOnly"
	AzureDataDiskCachingReadWrite AzureDataDiskCachingMode = "ReadWrite"
	AzureSharedBlobDisk           AzureDataDiskKind        = "Shared"
	AzureDedicatedBlobDisk        AzureDataDiskKind        = "Dedicated"
	AzureManagedDisk              AzureDataDiskKind        = "Managed"
)

type AzureDiskVolumeSource struct {
	DiskName    string
	DataDiskURI string
	CachingMode *AzureDataDiskCachingMode
	FSType      *string
	ReadOnly    *bool
	Kind        *AzureDataDiskKind
}
type ScaleIOVolumeSource struct {
	Gateway          string
	System           string
	SecretRef        *LocalObjectReference
	SSLEnabled       bool
	ProtectionDomain string
	StoragePool      string
	StorageMode      string
	VolumeName       string
	FSType           string
	ReadOnly         bool
}
type ScaleIOPersistentVolumeSource struct {
	Gateway          string
	System           string
	SecretRef        *SecretReference
	SSLEnabled       bool
	ProtectionDomain string
	StoragePool      string
	StorageMode      string
	VolumeName       string
	FSType           string
	ReadOnly         bool
}
type StorageOSVolumeSource struct {
	VolumeName      string
	VolumeNamespace string
	FSType          string
	ReadOnly        bool
	SecretRef       *LocalObjectReference
}
type StorageOSPersistentVolumeSource struct {
	VolumeName      string
	VolumeNamespace string
	FSType          string
	ReadOnly        bool
	SecretRef       *ObjectReference
}
type ConfigMapVolumeSource struct {
	LocalObjectReference
	Items       []KeyToPath
	DefaultMode *int32
	Optional    *bool
}
type ConfigMapProjection struct {
	LocalObjectReference
	Items    []KeyToPath
	Optional *bool
}
type ServiceAccountTokenProjection struct {
	Audience          string
	ExpirationSeconds int64
	Path              string
}
type ProjectedVolumeSource struct {
	Sources     []VolumeProjection
	DefaultMode *int32
}
type VolumeProjection struct {
	Secret              *SecretProjection
	DownwardAPI         *DownwardAPIProjection
	ConfigMap           *ConfigMapProjection
	ServiceAccountToken *ServiceAccountTokenProjection
}
type KeyToPath struct {
	Key  string
	Path string
	Mode *int32
}
type LocalVolumeSource struct {
	Path   string
	FSType *string
}
type CSIPersistentVolumeSource struct {
	Driver                     string
	VolumeHandle               string
	ReadOnly                   bool
	FSType                     string
	VolumeAttributes           map[string]string
	ControllerPublishSecretRef *SecretReference
	NodeStageSecretRef         *SecretReference
	NodePublishSecretRef       *SecretReference
}
type ContainerPort struct {
	Name          string
	HostPort      int32
	ContainerPort int32
	Protocol      Protocol
	HostIP        string
}
type VolumeMount struct {
	Name             string
	ReadOnly         bool
	MountPath        string
	SubPath          string
	MountPropagation *MountPropagationMode
}
type MountPropagationMode string

const (
	MountPropagationNone            MountPropagationMode = "None"
	MountPropagationHostToContainer MountPropagationMode = "HostToContainer"
	MountPropagationBidirectional   MountPropagationMode = "Bidirectional"
)

type VolumeDevice struct {
	Name       string
	DevicePath string
}
type EnvVar struct {
	Name      string
	Value     string
	ValueFrom *EnvVarSource
}
type EnvVarSource struct {
	FieldRef         *ObjectFieldSelector
	ResourceFieldRef *ResourceFieldSelector
	ConfigMapKeyRef  *ConfigMapKeySelector
	SecretKeyRef     *SecretKeySelector
}
type ObjectFieldSelector struct {
	APIVersion string
	FieldPath  string
}
type ResourceFieldSelector struct {
	ContainerName string
	Resource      string
	Divisor       resource.Quantity
}
type ConfigMapKeySelector struct {
	LocalObjectReference
	Key      string
	Optional *bool
}
type SecretKeySelector struct {
	LocalObjectReference
	Key      string
	Optional *bool
}
type EnvFromSource struct {
	Prefix       string
	ConfigMapRef *ConfigMapEnvSource
	SecretRef    *SecretEnvSource
}
type ConfigMapEnvSource struct {
	LocalObjectReference
	Optional *bool
}
type SecretEnvSource struct {
	LocalObjectReference
	Optional *bool
}
type HTTPHeader struct {
	Name  string
	Value string
}
type HTTPGetAction struct {
	Path        string
	Port        intstr.IntOrString
	Host        string
	Scheme      URIScheme
	HTTPHeaders []HTTPHeader
}
type URIScheme string

const (
	URISchemeHTTP  URIScheme = "HTTP"
	URISchemeHTTPS URIScheme = "HTTPS"
)

type TCPSocketAction struct {
	Port intstr.IntOrString
	Host string
}
type ExecAction struct{ Command []string }
type Probe struct {
	Handler
	InitialDelaySeconds int32
	TimeoutSeconds      int32
	PeriodSeconds       int32
	SuccessThreshold    int32
	FailureThreshold    int32
}
type PullPolicy string

const (
	PullAlways       PullPolicy = "Always"
	PullNever        PullPolicy = "Never"
	PullIfNotPresent PullPolicy = "IfNotPresent"
)

type TerminationMessagePolicy string

const (
	TerminationMessageReadFile              TerminationMessagePolicy = "File"
	TerminationMessageFallbackToLogsOnError TerminationMessagePolicy = "FallbackToLogsOnError"
)

type Capability string
type Capabilities struct {
	Add  []Capability
	Drop []Capability
}
type ResourceRequirements struct {
	Limits   ResourceList
	Requests ResourceList
}
type Container struct {
	Name                     string
	Image                    string
	Command                  []string
	Args                     []string
	WorkingDir               string
	Ports                    []ContainerPort
	EnvFrom                  []EnvFromSource
	Env                      []EnvVar
	Resources                ResourceRequirements
	VolumeMounts             []VolumeMount
	VolumeDevices            []VolumeDevice
	LivenessProbe            *Probe
	ReadinessProbe           *Probe
	Lifecycle                *Lifecycle
	TerminationMessagePath   string
	TerminationMessagePolicy TerminationMessagePolicy
	ImagePullPolicy          PullPolicy
	SecurityContext          *SecurityContext
	Stdin                    bool
	StdinOnce                bool
	TTY                      bool
}
type Handler struct {
	Exec      *ExecAction
	HTTPGet   *HTTPGetAction
	TCPSocket *TCPSocketAction
}
type Lifecycle struct {
	PostStart *Handler
	PreStop   *Handler
}
type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

type ContainerStateWaiting struct {
	Reason  string
	Message string
}
type ContainerStateRunning struct{ StartedAt metav1.Time }
type ContainerStateTerminated struct {
	ExitCode    int32
	Signal      int32
	Reason      string
	Message     string
	StartedAt   metav1.Time
	FinishedAt  metav1.Time
	ContainerID string
}
type ContainerState struct {
	Waiting    *ContainerStateWaiting
	Running    *ContainerStateRunning
	Terminated *ContainerStateTerminated
}
type ContainerStatus struct {
	Name                 string
	State                ContainerState
	LastTerminationState ContainerState
	Ready                bool
	RestartCount         int32
	Image                string
	ImageID              string
	ContainerID          string
}
type PodPhase string

const (
	PodPending   PodPhase = "Pending"
	PodRunning   PodPhase = "Running"
	PodSucceeded PodPhase = "Succeeded"
	PodFailed    PodPhase = "Failed"
	PodUnknown   PodPhase = "Unknown"
)

type PodConditionType string

const (
	PodScheduled           PodConditionType = "PodScheduled"
	PodReady               PodConditionType = "Ready"
	PodInitialized         PodConditionType = "Initialized"
	PodReasonUnschedulable                  = "Unschedulable"
	ContainersReady        PodConditionType = "ContainersReady"
)

type PodCondition struct {
	Type               PodConditionType
	Status             ConditionStatus
	LastProbeTime      metav1.Time
	LastTransitionTime metav1.Time
	Reason             string
	Message            string
}
type RestartPolicy string

const (
	RestartPolicyAlways    RestartPolicy = "Always"
	RestartPolicyOnFailure RestartPolicy = "OnFailure"
	RestartPolicyNever     RestartPolicy = "Never"
)

type PodList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Pod
}
type DNSPolicy string

const (
	DNSClusterFirstWithHostNet DNSPolicy = "ClusterFirstWithHostNet"
	DNSClusterFirst            DNSPolicy = "ClusterFirst"
	DNSDefault                 DNSPolicy = "Default"
	DNSNone                    DNSPolicy = "None"
)

type NodeSelector struct{ NodeSelectorTerms []NodeSelectorTerm }
type NodeSelectorTerm struct {
	MatchExpressions []NodeSelectorRequirement
	MatchFields      []NodeSelectorRequirement
}
type NodeSelectorRequirement struct {
	Key      string
	Operator NodeSelectorOperator
	Values   []string
}
type NodeSelectorOperator string

const (
	NodeSelectorOpIn           NodeSelectorOperator = "In"
	NodeSelectorOpNotIn        NodeSelectorOperator = "NotIn"
	NodeSelectorOpExists       NodeSelectorOperator = "Exists"
	NodeSelectorOpDoesNotExist NodeSelectorOperator = "DoesNotExist"
	NodeSelectorOpGt           NodeSelectorOperator = "Gt"
	NodeSelectorOpLt           NodeSelectorOperator = "Lt"
)

type TopologySelectorTerm struct {
	MatchLabelExpressions []TopologySelectorLabelRequirement
}
type TopologySelectorLabelRequirement struct {
	Key    string
	Values []string
}
type Affinity struct {
	NodeAffinity    *NodeAffinity
	PodAffinity     *PodAffinity
	PodAntiAffinity *PodAntiAffinity
}
type PodAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTerm
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm
}
type PodAntiAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTerm
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm
}
type WeightedPodAffinityTerm struct {
	Weight          int32
	PodAffinityTerm PodAffinityTerm
}
type PodAffinityTerm struct {
	LabelSelector *metav1.LabelSelector
	Namespaces    []string
	TopologyKey   string
}
type NodeAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  *NodeSelector
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredSchedulingTerm
}
type PreferredSchedulingTerm struct {
	Weight     int32
	Preference NodeSelectorTerm
}
type Taint struct {
	Key       string
	Value     string
	Effect    TaintEffect
	TimeAdded *metav1.Time
}
type TaintEffect string

const (
	TaintEffectNoSchedule       TaintEffect = "NoSchedule"
	TaintEffectPreferNoSchedule TaintEffect = "PreferNoSchedule"
	TaintEffectNoExecute        TaintEffect = "NoExecute"
)

type Toleration struct {
	Key               string
	Operator          TolerationOperator
	Value             string
	Effect            TaintEffect
	TolerationSeconds *int64
}
type TolerationOperator string

const (
	TolerationOpExists TolerationOperator = "Exists"
	TolerationOpEqual  TolerationOperator = "Equal"
)

type PodReadinessGate struct{ ConditionType PodConditionType }
type PodSpec struct {
	Volumes                       []Volume
	InitContainers                []Container
	Containers                    []Container
	RestartPolicy                 RestartPolicy
	TerminationGracePeriodSeconds *int64
	ActiveDeadlineSeconds         *int64
	DNSPolicy                     DNSPolicy
	NodeSelector                  map[string]string
	ServiceAccountName            string
	AutomountServiceAccountToken  *bool
	NodeName                      string
	SecurityContext               *PodSecurityContext
	ImagePullSecrets              []LocalObjectReference
	Hostname                      string
	Subdomain                     string
	Affinity                      *Affinity
	SchedulerName                 string
	Tolerations                   []Toleration
	HostAliases                   []HostAlias
	PriorityClassName             string
	Priority                      *int32
	DNSConfig                     *PodDNSConfig
	ReadinessGates                []PodReadinessGate
	RuntimeClassName              *string
	EnableServiceLinks            *bool
}
type HostAlias struct {
	IP        string
	Hostnames []string
}
type Sysctl struct {
	Name  string
	Value string
}
type PodSecurityContext struct {
	HostNetwork           bool
	HostPID               bool
	HostIPC               bool
	ShareProcessNamespace *bool
	SELinuxOptions        *SELinuxOptions
	RunAsUser             *int64
	RunAsGroup            *int64
	RunAsNonRoot          *bool
	SupplementalGroups    []int64
	FSGroup               *int64
	Sysctls               []Sysctl
}
type PodQOSClass string

const (
	PodQOSGuaranteed PodQOSClass = "Guaranteed"
	PodQOSBurstable  PodQOSClass = "Burstable"
	PodQOSBestEffort PodQOSClass = "BestEffort"
)

type PodDNSConfig struct {
	Nameservers []string
	Searches    []string
	Options     []PodDNSConfigOption
}
type PodDNSConfigOption struct {
	Name  string
	Value *string
}
type PodStatus struct {
	Phase                 PodPhase
	Conditions            []PodCondition
	Message               string
	Reason                string
	NominatedNodeName     string
	HostIP                string
	PodIP                 string
	StartTime             *metav1.Time
	QOSClass              PodQOSClass
	InitContainerStatuses []ContainerStatus
	ContainerStatuses     []ContainerStatus
}
type PodStatusResult struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Status PodStatus
}
type Pod struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   PodSpec
	Status PodStatus
}
type PodTemplateSpec struct {
	metav1.ObjectMeta
	Spec PodSpec
}
type PodTemplate struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Template PodTemplateSpec
}
type PodTemplateList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []PodTemplate
}
type ReplicationControllerSpec struct {
	Replicas        int32
	MinReadySeconds int32
	Selector        map[string]string
	Template        *PodTemplateSpec
}
type ReplicationControllerStatus struct {
	Replicas             int32
	FullyLabeledReplicas int32
	ReadyReplicas        int32
	AvailableReplicas    int32
	ObservedGeneration   int64
	Conditions           []ReplicationControllerCondition
}
type ReplicationControllerConditionType string

const (
	ReplicationControllerReplicaFailure ReplicationControllerConditionType = "ReplicaFailure"
)

type ReplicationControllerCondition struct {
	Type               ReplicationControllerConditionType
	Status             ConditionStatus
	LastTransitionTime metav1.Time
	Reason             string
	Message            string
}
type ReplicationController struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ReplicationControllerSpec
	Status ReplicationControllerStatus
}
type ReplicationControllerList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ReplicationController
}

const (
	ClusterIPNone = "None"
)

type ServiceList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Service
}
type ServiceAffinity string

const (
	ServiceAffinityClientIP ServiceAffinity = "ClientIP"
	ServiceAffinityNone     ServiceAffinity = "None"
)
const (
	DefaultClientIPServiceAffinitySeconds int32 = 10800
	MaxClientIPServiceAffinitySeconds     int32 = 86400
)

type SessionAffinityConfig struct{ ClientIP *ClientIPConfig }
type ClientIPConfig struct{ TimeoutSeconds *int32 }
type ServiceType string

const (
	ServiceTypeClusterIP    ServiceType = "ClusterIP"
	ServiceTypeNodePort     ServiceType = "NodePort"
	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
	ServiceTypeExternalName ServiceType = "ExternalName"
)

type ServiceExternalTrafficPolicyType string

const (
	ServiceExternalTrafficPolicyTypeLocal   ServiceExternalTrafficPolicyType = "Local"
	ServiceExternalTrafficPolicyTypeCluster ServiceExternalTrafficPolicyType = "Cluster"
)

type ServiceStatus struct{ LoadBalancer LoadBalancerStatus }
type LoadBalancerStatus struct{ Ingress []LoadBalancerIngress }
type LoadBalancerIngress struct {
	IP       string
	Hostname string
}
type ServiceSpec struct {
	Type                     ServiceType
	Ports                    []ServicePort
	Selector                 map[string]string
	ClusterIP                string
	ExternalName             string
	ExternalIPs              []string
	LoadBalancerIP           string
	SessionAffinity          ServiceAffinity
	SessionAffinityConfig    *SessionAffinityConfig
	LoadBalancerSourceRanges []string
	ExternalTrafficPolicy    ServiceExternalTrafficPolicyType
	HealthCheckNodePort      int32
	PublishNotReadyAddresses bool
}
type ServicePort struct {
	Name       string
	Protocol   Protocol
	Port       int32
	TargetPort intstr.IntOrString
	NodePort   int32
}
type Service struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ServiceSpec
	Status ServiceStatus
}
type ServiceAccount struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Secrets                      []ObjectReference
	ImagePullSecrets             []LocalObjectReference
	AutomountServiceAccountToken *bool
}
type ServiceAccountList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ServiceAccount
}
type Endpoints struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Subsets []EndpointSubset
}
type EndpointSubset struct {
	Addresses         []EndpointAddress
	NotReadyAddresses []EndpointAddress
	Ports             []EndpointPort
}
type EndpointAddress struct {
	IP        string
	Hostname  string
	NodeName  *string
	TargetRef *ObjectReference
}
type EndpointPort struct {
	Name     string
	Port     int32
	Protocol Protocol
}
type EndpointsList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Endpoints
}
type NodeSpec struct {
	PodCIDR             string
	ProviderID          string
	Unschedulable       bool
	Taints              []Taint
	ConfigSource        *NodeConfigSource
	DoNotUse_ExternalID string
}
type NodeConfigSource struct{ ConfigMap *ConfigMapNodeConfigSource }
type ConfigMapNodeConfigSource struct {
	Namespace        string
	Name             string
	UID              types.UID
	ResourceVersion  string
	KubeletConfigKey string
}
type DaemonEndpoint struct{ Port int32 }
type NodeDaemonEndpoints struct{ KubeletEndpoint DaemonEndpoint }
type NodeSystemInfo struct {
	MachineID               string
	SystemUUID              string
	BootID                  string
	KernelVersion           string
	OSImage                 string
	ContainerRuntimeVersion string
	KubeletVersion          string
	KubeProxyVersion        string
	OperatingSystem         string
	Architecture            string
}
type NodeConfigStatus struct {
	Assigned      *NodeConfigSource
	Active        *NodeConfigSource
	LastKnownGood *NodeConfigSource
	Error         string
}
type NodeStatus struct {
	Capacity        ResourceList
	Allocatable     ResourceList
	Phase           NodePhase
	Conditions      []NodeCondition
	Addresses       []NodeAddress
	DaemonEndpoints NodeDaemonEndpoints
	NodeInfo        NodeSystemInfo
	Images          []ContainerImage
	VolumesInUse    []UniqueVolumeName
	VolumesAttached []AttachedVolume
	Config          *NodeConfigStatus
}
type UniqueVolumeName string
type AttachedVolume struct {
	Name       UniqueVolumeName
	DevicePath string
}
type AvoidPods struct{ PreferAvoidPods []PreferAvoidPodsEntry }
type PreferAvoidPodsEntry struct {
	PodSignature PodSignature
	EvictionTime metav1.Time
	Reason       string
	Message      string
}
type PodSignature struct{ PodController *metav1.OwnerReference }
type ContainerImage struct {
	Names     []string
	SizeBytes int64
}
type NodePhase string

const (
	NodePending    NodePhase = "Pending"
	NodeRunning    NodePhase = "Running"
	NodeTerminated NodePhase = "Terminated"
)

type NodeConditionType string

const (
	NodeReady              NodeConditionType = "Ready"
	NodeOutOfDisk          NodeConditionType = "OutOfDisk"
	NodeMemoryPressure     NodeConditionType = "MemoryPressure"
	NodeDiskPressure       NodeConditionType = "DiskPressure"
	NodeNetworkUnavailable NodeConditionType = "NetworkUnavailable"
)

type NodeCondition struct {
	Type               NodeConditionType
	Status             ConditionStatus
	LastHeartbeatTime  metav1.Time
	LastTransitionTime metav1.Time
	Reason             string
	Message            string
}
type NodeAddressType string

const (
	NodeHostName    NodeAddressType = "Hostname"
	NodeExternalIP  NodeAddressType = "ExternalIP"
	NodeInternalIP  NodeAddressType = "InternalIP"
	NodeExternalDNS NodeAddressType = "ExternalDNS"
	NodeInternalDNS NodeAddressType = "InternalDNS"
)

type NodeAddress struct {
	Type    NodeAddressType
	Address string
}
type NodeResources struct{ Capacity ResourceList }
type ResourceName string

const (
	ResourceCPU              ResourceName = "cpu"
	ResourceMemory           ResourceName = "memory"
	ResourceStorage          ResourceName = "storage"
	ResourceEphemeralStorage ResourceName = "ephemeral-storage"
)
const (
	ResourceDefaultNamespacePrefix  = "kubernetes.io/"
	ResourceHugePagesPrefix         = "hugepages-"
	ResourceAttachableVolumesPrefix = "attachable-volumes-"
)

type ResourceList map[ResourceName]resource.Quantity
type Node struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   NodeSpec
	Status NodeStatus
}
type NodeList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Node
}
type NamespaceSpec struct{ Finalizers []FinalizerName }
type FinalizerName string

const (
	FinalizerKubernetes FinalizerName = "kubernetes"
)

type NamespaceStatus struct{ Phase NamespacePhase }
type NamespacePhase string

const (
	NamespaceActive      NamespacePhase = "Active"
	NamespaceTerminating NamespacePhase = "Terminating"
)

type Namespace struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   NamespaceSpec
	Status NamespaceStatus
}
type NamespaceList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Namespace
}
type Binding struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Target ObjectReference
}
type Preconditions struct{ UID *types.UID }
type PodLogOptions struct {
	metav1.TypeMeta
	Container    string
	Follow       bool
	Previous     bool
	SinceSeconds *int64
	SinceTime    *metav1.Time
	Timestamps   bool
	TailLines    *int64
	LimitBytes   *int64
}
type PodAttachOptions struct {
	metav1.TypeMeta
	Stdin     bool
	Stdout    bool
	Stderr    bool
	TTY       bool
	Container string
}
type PodExecOptions struct {
	metav1.TypeMeta
	Stdin     bool
	Stdout    bool
	Stderr    bool
	TTY       bool
	Container string
	Command   []string
}
type PodPortForwardOptions struct {
	metav1.TypeMeta
	Ports []int32
}
type PodProxyOptions struct {
	metav1.TypeMeta
	Path string
}
type NodeProxyOptions struct {
	metav1.TypeMeta
	Path string
}
type ServiceProxyOptions struct {
	metav1.TypeMeta
	Path string
}
type ObjectReference struct {
	Kind            string
	Namespace       string
	Name            string
	UID             types.UID
	APIVersion      string
	ResourceVersion string
	FieldPath       string
}
type LocalObjectReference struct{ Name string }
type TypedLocalObjectReference struct {
	APIGroup *string
	Kind     string
	Name     string
}
type SerializedReference struct {
	metav1.TypeMeta
	Reference ObjectReference
}
type EventSource struct {
	Component string
	Host      string
}

const (
	EventTypeNormal  string = "Normal"
	EventTypeWarning string = "Warning"
)

type Event struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	InvolvedObject      ObjectReference
	Reason              string
	Message             string
	Source              EventSource
	FirstTimestamp      metav1.Time
	LastTimestamp       metav1.Time
	Count               int32
	Type                string
	EventTime           metav1.MicroTime
	Series              *EventSeries
	Action              string
	Related             *ObjectReference
	ReportingController string
	ReportingInstance   string
}
type EventSeries struct {
	Count            int32
	LastObservedTime metav1.MicroTime
	State            EventSeriesState
}
type EventSeriesState string

const (
	EventSeriesStateOngoing  EventSeriesState = "Ongoing"
	EventSeriesStateFinished EventSeriesState = "Finished"
	EventSeriesStateUnknown  EventSeriesState = "Unknown"
)

type EventList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Event
}
type List metainternalversion.List
type LimitType string

const (
	LimitTypePod                   LimitType = "Pod"
	LimitTypeContainer             LimitType = "Container"
	LimitTypePersistentVolumeClaim LimitType = "PersistentVolumeClaim"
)

type LimitRangeItem struct {
	Type                 LimitType
	Max                  ResourceList
	Min                  ResourceList
	Default              ResourceList
	DefaultRequest       ResourceList
	MaxLimitRequestRatio ResourceList
}
type LimitRangeSpec struct{ Limits []LimitRangeItem }
type LimitRange struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec LimitRangeSpec
}
type LimitRangeList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []LimitRange
}

const (
	ResourcePods                     ResourceName = "pods"
	ResourceServices                 ResourceName = "services"
	ResourceReplicationControllers   ResourceName = "replicationcontrollers"
	ResourceQuotas                   ResourceName = "resourcequotas"
	ResourceSecrets                  ResourceName = "secrets"
	ResourceConfigMaps               ResourceName = "configmaps"
	ResourcePersistentVolumeClaims   ResourceName = "persistentvolumeclaims"
	ResourceServicesNodePorts        ResourceName = "services.nodeports"
	ResourceServicesLoadBalancers    ResourceName = "services.loadbalancers"
	ResourceRequestsCPU              ResourceName = "requests.cpu"
	ResourceRequestsMemory           ResourceName = "requests.memory"
	ResourceRequestsStorage          ResourceName = "requests.storage"
	ResourceRequestsEphemeralStorage ResourceName = "requests.ephemeral-storage"
	ResourceLimitsCPU                ResourceName = "limits.cpu"
	ResourceLimitsMemory             ResourceName = "limits.memory"
	ResourceLimitsEphemeralStorage   ResourceName = "limits.ephemeral-storage"
)
const (
	ResourceRequestsHugePagesPrefix = "requests.hugepages-"
	DefaultResourceRequestsPrefix   = "requests."
)

type ResourceQuotaScope string

const (
	ResourceQuotaScopeTerminating    ResourceQuotaScope = "Terminating"
	ResourceQuotaScopeNotTerminating ResourceQuotaScope = "NotTerminating"
	ResourceQuotaScopeBestEffort     ResourceQuotaScope = "BestEffort"
	ResourceQuotaScopeNotBestEffort  ResourceQuotaScope = "NotBestEffort"
	ResourceQuotaScopePriorityClass  ResourceQuotaScope = "PriorityClass"
)

type ResourceQuotaSpec struct {
	Hard          ResourceList
	Scopes        []ResourceQuotaScope
	ScopeSelector *ScopeSelector
}
type ScopeSelector struct {
	MatchExpressions []ScopedResourceSelectorRequirement
}
type ScopedResourceSelectorRequirement struct {
	ScopeName ResourceQuotaScope
	Operator  ScopeSelectorOperator
	Values    []string
}
type ScopeSelectorOperator string

const (
	ScopeSelectorOpIn           ScopeSelectorOperator = "In"
	ScopeSelectorOpNotIn        ScopeSelectorOperator = "NotIn"
	ScopeSelectorOpExists       ScopeSelectorOperator = "Exists"
	ScopeSelectorOpDoesNotExist ScopeSelectorOperator = "DoesNotExist"
)

type ResourceQuotaStatus struct {
	Hard ResourceList
	Used ResourceList
}
type ResourceQuota struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ResourceQuotaSpec
	Status ResourceQuotaStatus
}
type ResourceQuotaList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ResourceQuota
}
type Secret struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Data map[string][]byte
	Type SecretType
}

const MaxSecretSize = 1 * 1024 * 1024

type SecretType string

const (
	SecretTypeOpaque              SecretType = "Opaque"
	SecretTypeServiceAccountToken SecretType = "kubernetes.io/service-account-token"
	ServiceAccountNameKey                    = "kubernetes.io/service-account.name"
	ServiceAccountUIDKey                     = "kubernetes.io/service-account.uid"
	ServiceAccountTokenKey                   = "token"
	ServiceAccountKubeconfigKey              = "kubernetes.kubeconfig"
	ServiceAccountRootCAKey                  = "ca.crt"
	ServiceAccountNamespaceKey               = "namespace"
	SecretTypeDockercfg           SecretType = "kubernetes.io/dockercfg"
	DockerConfigKey                          = ".dockercfg"
	SecretTypeDockerConfigJson    SecretType = "kubernetes.io/dockerconfigjson"
	DockerConfigJsonKey                      = ".dockerconfigjson"
	SecretTypeBasicAuth           SecretType = "kubernetes.io/basic-auth"
	BasicAuthUsernameKey                     = "username"
	BasicAuthPasswordKey                     = "password"
	SecretTypeSSHAuth             SecretType = "kubernetes.io/ssh-auth"
	SSHAuthPrivateKey                        = "ssh-privatekey"
	SecretTypeTLS                 SecretType = "kubernetes.io/tls"
	TLSCertKey                               = "tls.crt"
	TLSPrivateKeyKey                         = "tls.key"
	SecretTypeBootstrapToken      SecretType = "bootstrap.kubernetes.io/token"
)

type SecretList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Secret
}
type ConfigMap struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Data       map[string]string
	BinaryData map[string][]byte
}
type ConfigMapList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ConfigMap
}

const (
	ExecStdinParam             = "input"
	ExecStdoutParam            = "output"
	ExecStderrParam            = "error"
	ExecTTYParam               = "tty"
	ExecCommandParam           = "command"
	StreamType                 = "streamType"
	StreamTypeStdin            = "stdin"
	StreamTypeStdout           = "stdout"
	StreamTypeStderr           = "stderr"
	StreamTypeData             = "data"
	StreamTypeError            = "error"
	StreamTypeResize           = "resize"
	PortHeader                 = "port"
	PortForwardRequestIDHeader = "requestID"
)

type ComponentConditionType string

const (
	ComponentHealthy ComponentConditionType = "Healthy"
)

type ComponentCondition struct {
	Type    ComponentConditionType
	Status  ConditionStatus
	Message string
	Error   string
}
type ComponentStatus struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Conditions []ComponentCondition
}
type ComponentStatusList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ComponentStatus
}
type SecurityContext struct {
	Capabilities             *Capabilities
	Privileged               *bool
	SELinuxOptions           *SELinuxOptions
	RunAsUser                *int64
	RunAsGroup               *int64
	RunAsNonRoot             *bool
	ReadOnlyRootFilesystem   *bool
	AllowPrivilegeEscalation *bool
	ProcMount                *ProcMountType
}
type ProcMountType string

const (
	DefaultProcMount  ProcMountType = "Default"
	UnmaskedProcMount ProcMountType = "Unmasked"
)

type SELinuxOptions struct {
	User  string
	Role  string
	Type  string
	Level string
}
type RangeAllocation struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Range string
	Data  []byte
}

const (
	DefaultSchedulerName                        = "default-scheduler"
	DefaultHardPodAffinitySymmetricWeight int32 = 1
)

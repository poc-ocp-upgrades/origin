package config

import (
	"github.com/openshift/origin/pkg/build/apis/build"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kubernetes/pkg/apis/core"
)

const (
	AllVersions = "*"
)
const (
	DefaultIngressIPNetworkCIDR = "172.29.0.0/16"
)

var (
	KnownKubernetesAPILevels		= []string{"v1beta1", "v1beta2", "v1beta3", "v1"}
	KnownOpenShiftAPILevels			= []string{"v1beta1", "v1beta3", "v1"}
	DefaultKubernetesAPILevels		= []string{"v1"}
	DefaultOpenShiftAPILevels		= []string{"v1"}
	DeadKubernetesAPILevels			= []string{"v1beta1", "v1beta2", "v1beta3"}
	DeadOpenShiftAPILevels			= []string{"v1beta1", "v1beta3"}
	KnownKubernetesStorageVersionLevels	= []string{"v1", "v1beta3"}
	KnownOpenShiftStorageVersionLevels	= []string{"v1", "v1beta3"}
	DefaultOpenShiftStorageVersionLevel	= "v1"
	DeadKubernetesStorageVersionLevels	= []string{"v1beta3"}
	DeadOpenShiftStorageVersionLevels	= []string{"v1beta1", "v1beta3"}
	APIGroupKube				= ""
	APIGroupApps				= "apps"
	APIGroupAdmission			= "admission.k8s.io"
	APIGroupAdmissionRegistration		= "admissionregistration.k8s.io"
	APIGroupAPIExtensions			= "apiextensions.k8s.io"
	APIGroupAPIRegistration			= "apiregistration.k8s.io"
	APIGroupAuthentication			= "authentication.k8s.io"
	APIGroupAuthorization			= "authorization.k8s.io"
	APIGroupExtensions			= "extensions"
	APIGroupEvents				= "events.k8s.io"
	APIGroupImagePolicy			= "imagepolicy.k8s.io"
	APIGroupAutoscaling			= "autoscaling"
	APIGroupBatch				= "batch"
	APIGroupCertificates			= "certificates.k8s.io"
	APIGroupCoordination			= "coordination.k8s.io"
	APIGroupNetworking			= "networking.k8s.io"
	APIGroupPolicy				= "policy"
	APIGroupStorage				= "storage.k8s.io"
	APIGroupComponentConfig			= "componentconfig"
	APIGroupAuthorizationRbac		= "rbac.authorization.k8s.io"
	APIGroupSettings			= "settings.k8s.io"
	APIGroupScheduling			= "scheduling.k8s.io"
	OriginAPIGroupCore			= ""
	OriginAPIGroupAuthorization		= "authorization.openshift.io"
	OriginAPIGroupBuild			= "build.openshift.io"
	OriginAPIGroupDeploy			= "apps.openshift.io"
	OriginAPIGroupTemplate			= "template.openshift.io"
	OriginAPIGroupImage			= "image.openshift.io"
	OriginAPIGroupProject			= "project.openshift.io"
	OriginAPIGroupProjectRequestLimit	= "requestlimit.project.openshift.io"
	OriginAPIGroupUser			= "user.openshift.io"
	OriginAPIGroupOAuth			= "oauth.openshift.io"
	OriginAPIGroupRoute			= "route.openshift.io"
	OriginAPIGroupNetwork			= "network.openshift.io"
	OriginAPIGroupQuota			= "quota.openshift.io"
	OriginAPIGroupSecurity			= "security.openshift.io"
	KubeAPIGroupsToAllowedVersions		= map[string][]string{APIGroupKube: {"v1"}, APIGroupExtensions: {"v1beta1"}, APIGroupEvents: {"v1beta1"}, APIGroupApps: {"v1", "v1beta1", "v1beta2"}, APIGroupAdmission: {}, APIGroupAdmissionRegistration: {"v1beta1"}, APIGroupAPIExtensions: {"v1beta1"}, APIGroupAPIRegistration: {"v1", "v1beta1"}, APIGroupAuthentication: {"v1", "v1beta1"}, APIGroupAuthorization: {"v1", "v1beta1"}, APIGroupAuthorizationRbac: {"v1", "v1beta1"}, APIGroupAutoscaling: {"v1", "v2beta1", "v2beta2"}, APIGroupBatch: {"v1", "v1beta1"}, APIGroupCertificates: {"v1beta1"}, APIGroupCoordination: {"v1beta1"}, APIGroupImagePolicy: {}, APIGroupNetworking: {"v1"}, APIGroupPolicy: {"v1beta1"}, APIGroupStorage: {"v1", "v1beta1"}, APIGroupSettings: {}, APIGroupScheduling: {"v1beta1"}}
	OriginAPIGroupsToAllowedVersions	= map[string][]string{OriginAPIGroupAuthorization: {"v1"}, OriginAPIGroupBuild: {"v1"}, OriginAPIGroupDeploy: {"v1"}, OriginAPIGroupTemplate: {"v1"}, OriginAPIGroupImage: {"v1"}, OriginAPIGroupProject: {"v1"}, OriginAPIGroupUser: {"v1"}, OriginAPIGroupOAuth: {"v1"}, OriginAPIGroupNetwork: {"v1"}, OriginAPIGroupRoute: {"v1"}, OriginAPIGroupQuota: {"v1"}, OriginAPIGroupSecurity: {"v1"}}
	KubeDefaultDisabledVersions		= map[string][]string{APIGroupKube: {"v1beta3"}, APIGroupExtensions: {}, APIGroupAutoscaling: {"v2alpha1"}, APIGroupBatch: {"v2alpha1"}, APIGroupImagePolicy: {"v1alpha1"}, APIGroupPolicy: {}, APIGroupApps: {}, APIGroupAdmission: {"v1beta1"}, APIGroupAdmissionRegistration: {"v1alpha1"}, APIGroupAuthorizationRbac: {"v1alpha1"}, APIGroupSettings: {"v1alpha1"}, APIGroupScheduling: {"v1alpha1"}, APIGroupStorage: {"v1alpha1"}}
	KnownKubeAPIGroups			= sets.StringKeySet(KubeAPIGroupsToAllowedVersions)
	KnownOriginAPIGroups			= sets.StringKeySet(OriginAPIGroupsToAllowedVersions)
)

type ExtendedArguments map[string][]string
type NodeConfig struct {
	metav1.TypeMeta
	NodeName			string
	NodeIP				string
	ServingInfo			ServingInfo
	MasterKubeConfig		string
	MasterClientConnectionOverrides	*ClientConnectionOverrides
	DNSDomain			string
	DNSIP				string
	DNSBindAddress			string
	DNSNameservers			[]string
	DNSRecursiveResolvConf		string
	NetworkConfig			NodeNetworkConfig
	VolumeDirectory			string
	ImageConfig			ImageConfig
	AllowDisabledDocker		bool
	PodManifestConfig		*PodManifestConfig
	AuthConfig			NodeAuthConfig
	DockerConfig			DockerConfig
	KubeletArguments		ExtendedArguments
	ProxyArguments			ExtendedArguments
	IPTablesSyncPeriod		string
	EnableUnidling			bool
	VolumeConfig			NodeVolumeConfig
}
type NodeVolumeConfig struct{ LocalQuota LocalQuota }
type MasterVolumeConfig struct{ DynamicProvisioningEnabled bool }
type LocalQuota struct{ PerFSGroup *resource.Quantity }
type NodeNetworkConfig struct {
	NetworkPluginName	string
	MTU			uint32
}
type NodeAuthConfig struct {
	AuthenticationCacheTTL	string
	AuthenticationCacheSize	int
	AuthorizationCacheTTL	string
	AuthorizationCacheSize	int
}
type DockerConfig struct {
	ExecHandlerName		DockerExecHandlerType
	DockerShimSocket	string
	DockershimRootDirectory	string
}
type DockerExecHandlerType string

const (
	DockerExecHandlerNative		DockerExecHandlerType	= "native"
	DockerExecHandlerNsenter	DockerExecHandlerType	= "nsenter"
)

type MasterConfig struct {
	metav1.TypeMeta
	ServingInfo		HTTPServingInfo
	AuthConfig		MasterAuthConfig
	AggregatorConfig	AggregatorConfig
	CORSAllowedOrigins	[]string
	APILevels		[]string
	MasterPublicURL		string
	AdmissionConfig		AdmissionConfig
	Controllers		string
	ControllerConfig	ControllerConfig
	EtcdStorageConfig	EtcdStorageConfig
	EtcdClientInfo		EtcdConnectionInfo
	KubeletClientInfo	KubeletConnectionInfo
	KubernetesMasterConfig	KubernetesMasterConfig
	EtcdConfig		*EtcdConfig
	OAuthConfig		*OAuthConfig
	DNSConfig		*DNSConfig
	ServiceAccountConfig	ServiceAccountConfig
	MasterClients		MasterClients
	ImageConfig		ImageConfig
	ImagePolicyConfig	ImagePolicyConfig
	PolicyConfig		PolicyConfig
	ProjectConfig		ProjectConfig
	RoutingConfig		RoutingConfig
	NetworkConfig		MasterNetworkConfig
	VolumeConfig		MasterVolumeConfig
	JenkinsPipelineConfig	JenkinsPipelineConfig
	AuditConfig		AuditConfig
	DisableOpenAPI		bool
}
type MasterAuthConfig struct {
	RequestHeader			*RequestHeaderAuthenticationOptions
	WebhookTokenAuthenticators	[]WebhookTokenAuthenticator
	OAuthMetadataFile		string
}
type RequestHeaderAuthenticationOptions struct {
	ClientCA		string
	ClientCommonNames	[]string
	UsernameHeaders		[]string
	GroupHeaders		[]string
	ExtraHeaderPrefixes	[]string
}
type AggregatorConfig struct{ ProxyClientInfo CertInfo }
type LogFormatType string
type WebHookModeType string

const (
	LogFormatLegacy		LogFormatType	= "legacy"
	LogFormatJson		LogFormatType	= "json"
	WebHookModeBatch	WebHookModeType	= "batch"
	WebHookModeBlocking	WebHookModeType	= "blocking"
)

type AuditConfig struct {
	Enabled				bool
	AuditFilePath			string
	InternalAuditFilePath		string
	MaximumFileRetentionDays	int
	MaximumRetainedFiles		int
	MaximumFileSizeMegabytes	int
	PolicyFile			string
	PolicyConfiguration		runtime.Object
	LogFormat			LogFormatType
	WebHookKubeConfig		string
	WebHookMode			WebHookModeType
}
type JenkinsPipelineConfig struct {
	AutoProvisionEnabled	*bool
	TemplateNamespace	string
	TemplateName		string
	ServiceName		string
	Parameters		map[string]string
}
type ImagePolicyConfig struct {
	MaxImagesBulkImportedPerRepository		int
	DisableScheduledImport				bool
	ScheduledImageImportMinimumIntervalSeconds	int
	MaxScheduledImageImportsPerMinute		int
	AllowedRegistriesForImport			*AllowedRegistries
	InternalRegistryHostname			string
	ExternalRegistryHostnames			[]string
	AdditionalTrustedCA				string
}
type AllowedRegistries []RegistryLocation
type RegistryLocation struct {
	DomainName	string
	Insecure	bool
}
type ProjectConfig struct {
	DefaultNodeSelector	string
	ProjectRequestMessage	string
	ProjectRequestTemplate	string
	SecurityAllocator	*SecurityAllocator
}
type RoutingConfig struct{ Subdomain string }
type SecurityAllocator struct {
	UIDAllocatorRange	string
	MCSAllocatorRange	string
	MCSLabelsPerProject	int
}
type PolicyConfig struct{ UserAgentMatchingConfig UserAgentMatchingConfig }
type UserAgentMatchingConfig struct {
	RequiredClients		[]UserAgentMatchRule
	DeniedClients		[]UserAgentDenyRule
	DefaultRejectionMessage	string
}
type UserAgentMatchRule struct {
	Regex		string
	HTTPVerbs	[]string
}
type UserAgentDenyRule struct {
	UserAgentMatchRule
	RejectionMessage	string
}
type MasterNetworkConfig struct {
	NetworkPluginName		string
	DeprecatedClusterNetworkCIDR	string
	ClusterNetworks			[]ClusterNetworkEntry
	DeprecatedHostSubnetLength	uint32
	ServiceNetworkCIDR		string
	ExternalIPNetworkCIDRs		[]string
	IngressIPNetworkCIDR		string
	VXLANPort			uint32
}
type ClusterNetworkEntry struct {
	CIDR			string
	HostSubnetLength	uint32
}
type ImageConfig struct {
	Format	string
	Latest	bool
}
type RemoteConnectionInfo struct {
	URL		string
	CA		string
	ClientCert	CertInfo
}
type KubeletConnectionInfo struct {
	Port		uint
	CA		string
	ClientCert	CertInfo
}
type EtcdConnectionInfo struct {
	URLs		[]string
	CA		string
	ClientCert	CertInfo
}
type EtcdStorageConfig struct {
	KubernetesStorageVersion	string
	KubernetesStoragePrefix		string
	OpenShiftStorageVersion		string
	OpenShiftStoragePrefix		string
}
type ServingInfo struct {
	BindAddress		string
	BindNetwork		string
	ServerCert		CertInfo
	ClientCA		string
	NamedCertificates	[]NamedCertificate
	MinTLSVersion		string
	CipherSuites		[]string
}
type NamedCertificate struct {
	Names	[]string
	CertInfo
}
type HTTPServingInfo struct {
	ServingInfo
	MaxRequestsInFlight	int
	RequestTimeoutSeconds	int
}
type MasterClients struct {
	OpenShiftLoopbackKubeConfig			string
	OpenShiftLoopbackClientConnectionOverrides	*ClientConnectionOverrides
}
type ClientConnectionOverrides struct {
	AcceptContentTypes	string
	ContentType		string
	QPS			float32
	Burst			int32
}
type DNSConfig struct {
	BindAddress		string
	BindNetwork		string
	AllowRecursiveQueries	bool
}
type WebhookTokenAuthenticator struct {
	ConfigFile	string
	CacheTTL	string
}
type OAuthConfig struct {
	MasterCA			*string
	MasterURL			string
	MasterPublicURL			string
	AssetPublicURL			string
	AlwaysShowProviderSelection	bool
	IdentityProviders		[]IdentityProvider
	GrantConfig			GrantConfig
	SessionConfig			*SessionConfig
	TokenConfig			TokenConfig
	Templates			*OAuthTemplates
}
type OAuthTemplates struct {
	Login			string
	ProviderSelection	string
	Error			string
}
type ServiceAccountConfig struct {
	ManagedNames		[]string
	LimitSecretReferences	bool
	PrivateKeyFile		string
	PublicKeyFiles		[]string
	MasterCA		string
}
type TokenConfig struct {
	AuthorizeTokenMaxAgeSeconds		int32
	AccessTokenMaxAgeSeconds		int32
	AccessTokenInactivityTimeoutSeconds	*int32
}
type SessionConfig struct {
	SessionSecretsFile	string
	SessionMaxAgeSeconds	int32
	SessionName		string
}
type SessionSecrets struct {
	metav1.TypeMeta
	Secrets	[]SessionSecret
}
type SessionSecret struct {
	Authentication	string
	Encryption	string
}
type IdentityProvider struct {
	Name		string
	UseAsChallenger	bool
	UseAsLogin	bool
	MappingMethod	string
	Provider	runtime.Object
}
type BasicAuthPasswordIdentityProvider struct {
	metav1.TypeMeta
	RemoteConnectionInfo	RemoteConnectionInfo
}
type AllowAllPasswordIdentityProvider struct{ metav1.TypeMeta }
type DenyAllPasswordIdentityProvider struct{ metav1.TypeMeta }
type HTPasswdPasswordIdentityProvider struct {
	metav1.TypeMeta
	File	string
}
type LDAPPasswordIdentityProvider struct {
	metav1.TypeMeta
	URL		string
	BindDN		string
	BindPassword	StringSource
	Insecure	bool
	CA		string
	Attributes	LDAPAttributeMapping
}
type LDAPAttributeMapping struct {
	ID			[]string
	PreferredUsername	[]string
	Name			[]string
	Email			[]string
}
type KeystonePasswordIdentityProvider struct {
	metav1.TypeMeta
	RemoteConnectionInfo	RemoteConnectionInfo
	DomainName		string
	UseKeystoneIdentity	bool
}
type RequestHeaderIdentityProvider struct {
	metav1.TypeMeta
	LoginURL			string
	ChallengeURL			string
	ClientCA			string
	ClientCommonNames		[]string
	Headers				[]string
	PreferredUsernameHeaders	[]string
	NameHeaders			[]string
	EmailHeaders			[]string
}
type GitHubIdentityProvider struct {
	metav1.TypeMeta
	ClientID	string
	ClientSecret	StringSource
	Organizations	[]string
	Teams		[]string
	Hostname	string
	CA		string
}
type GitLabIdentityProvider struct {
	metav1.TypeMeta
	CA		string
	URL		string
	ClientID	string
	ClientSecret	StringSource
	Legacy		*bool
}
type GoogleIdentityProvider struct {
	metav1.TypeMeta
	ClientID	string
	ClientSecret	StringSource
	HostedDomain	string
}
type OpenIDIdentityProvider struct {
	metav1.TypeMeta
	CA				string
	ClientID			string
	ClientSecret			StringSource
	ExtraScopes			[]string
	ExtraAuthorizeParameters	map[string]string
	URLs				OpenIDURLs
	Claims				OpenIDClaims
}
type OpenIDURLs struct {
	Authorize	string
	Token		string
	UserInfo	string
}
type OpenIDClaims struct {
	ID			[]string
	PreferredUsername	[]string
	Name			[]string
	Email			[]string
}
type GrantConfig struct {
	Method			GrantHandlerType
	ServiceAccountMethod	GrantHandlerType
}
type GrantHandlerType string

const (
	GrantHandlerAuto	GrantHandlerType	= "auto"
	GrantHandlerPrompt	GrantHandlerType	= "prompt"
	GrantHandlerDeny	GrantHandlerType	= "deny"
)

var ValidGrantHandlerTypes = sets.NewString(string(GrantHandlerAuto), string(GrantHandlerPrompt), string(GrantHandlerDeny))
var ValidServiceAccountGrantHandlerTypes = sets.NewString(string(GrantHandlerPrompt), string(GrantHandlerDeny))

type EtcdConfig struct {
	ServingInfo	ServingInfo
	Address		string
	PeerServingInfo	ServingInfo
	PeerAddress	string
	StorageDir	string
}
type KubernetesMasterConfig struct {
	DisabledAPIGroupVersions	map[string][]string
	MasterIP			string
	MasterEndpointReconcileTTL	int
	ServicesSubnet			string
	ServicesNodePortRange		string
	SchedulerConfigFile		string
	PodEvictionTimeout		string
	ProxyClientInfo			CertInfo
	APIServerArguments		ExtendedArguments
	ControllerArguments		ExtendedArguments
	SchedulerArguments		ExtendedArguments
}
type CertInfo struct {
	CertFile	string
	KeyFile		string
}
type PodManifestConfig struct {
	Path				string
	FileCheckIntervalSeconds	int64
}

const (
	StringSourceEncryptedBlockType	= "ENCRYPTED STRING"
	StringSourceKeyBlockType	= "ENCRYPTING KEY"
)

type StringSource struct{ StringSourceSpec }
type StringSourceSpec struct {
	Value	string
	Env	string
	File	string
	KeyFile	string
}
type LDAPSyncConfig struct {
	metav1.TypeMeta
	URL					string
	BindDN					string
	BindPassword				StringSource
	Insecure				bool
	CA					string
	LDAPGroupUIDToOpenShiftGroupNameMapping	map[string]string
	RFC2307Config				*RFC2307Config
	ActiveDirectoryConfig			*ActiveDirectoryConfig
	AugmentedActiveDirectoryConfig		*AugmentedActiveDirectoryConfig
}
type RFC2307Config struct {
	AllGroupsQuery			LDAPQuery
	GroupUIDAttribute		string
	GroupNameAttributes		[]string
	GroupMembershipAttributes	[]string
	AllUsersQuery			LDAPQuery
	UserUIDAttribute		string
	UserNameAttributes		[]string
	TolerateMemberNotFoundErrors	bool
	TolerateMemberOutOfScopeErrors	bool
}
type ActiveDirectoryConfig struct {
	AllUsersQuery			LDAPQuery
	UserNameAttributes		[]string
	GroupMembershipAttributes	[]string
}
type AugmentedActiveDirectoryConfig struct {
	AllUsersQuery			LDAPQuery
	UserNameAttributes		[]string
	GroupMembershipAttributes	[]string
	AllGroupsQuery			LDAPQuery
	GroupUIDAttribute		string
	GroupNameAttributes		[]string
}
type LDAPQuery struct {
	BaseDN		string
	Scope		string
	DerefAliases	string
	TimeLimit	int
	Filter		string
	PageSize	int
}
type AdmissionPluginConfig struct {
	Location	string
	Configuration	runtime.Object
}
type AdmissionConfig struct {
	PluginConfig		map[string]*AdmissionPluginConfig
	PluginOrderOverride	[]string
}
type ControllerConfig struct {
	Controllers		[]string
	Election		*ControllerElectionConfig
	ServiceServingCert	ServiceServingCert
}
type ControllerElectionConfig struct {
	LockName	string
	LockNamespace	string
	LockResource	GroupResource
}
type GroupResource struct {
	Group		string
	Resource	string
}
type ServiceServingCert struct{ Signer *CertInfo }
type DefaultAdmissionConfig struct {
	metav1.TypeMeta
	Disable	bool
}
type OpenshiftControllerConfig struct {
	metav1.TypeMeta
	ClientConnectionOverrides	*ClientConnectionOverrides
	ServingInfo			*HTTPServingInfo
	LeaderElection			LeaderElectionConfig
	Controllers			[]string
	ResourceQuota			ResourceQuotaControllerConfig
	ServiceServingCert		ServiceServingCert
	Deployer			DeployerControllerConfig
	Build				BuildControllerConfig
	ServiceAccount			ServiceAccountControllerConfig
	DockerPullSecret		DockerPullSecretControllerConfig
	Network				NetworkControllerConfig
	Ingress				IngressControllerConfig
	ImageImport			ImageImportControllerConfig
	SecurityAllocator		SecurityAllocator
}
type DeployerControllerConfig struct{ ImageTemplateFormat ImageConfig }
type BuildControllerConfig struct {
	ImageTemplateFormat	ImageConfig
	BuildDefaults		*BuildDefaultsConfig
	BuildOverrides		*BuildOverridesConfig
}
type ResourceQuotaControllerConfig struct {
	ConcurrentSyncs	int32
	SyncPeriod	metav1.Duration
	MinResyncPeriod	metav1.Duration
}
type IngressControllerConfig struct{ IngressIPNetworkCIDR string }
type NetworkControllerConfig struct {
	NetworkPluginName	string
	ClusterNetworks		[]ClusterNetworkEntry
	ServiceNetworkCIDR	string
	VXLANPort		uint32
}
type ServiceAccountControllerConfig struct{ ManagedNames []string }
type DockerPullSecretControllerConfig struct{ RegistryURLs []string }
type ImageImportControllerConfig struct {
	MaxScheduledImageImportsPerMinute		int
	DisableScheduledImport				bool
	ScheduledImageImportMinimumIntervalSeconds	int
}
type LeaderElectionConfig struct {
	LeaseDuration	metav1.Duration
	RenewDeadline	metav1.Duration
	RetryPeriod	metav1.Duration
}
type BuildDefaultsConfig struct {
	metav1.TypeMeta
	GitHTTPProxy		string
	GitHTTPSProxy		string
	GitNoProxy		string
	Env			[]core.EnvVar
	SourceStrategyDefaults	*SourceStrategyDefaultsConfig
	ImageLabels		[]build.ImageLabel
	NodeSelector		map[string]string
	Annotations		map[string]string
	Resources		core.ResourceRequirements
}
type SourceStrategyDefaultsConfig struct{ Incremental *bool }
type BuildOverridesConfig struct {
	metav1.TypeMeta
	ForcePull	bool
	ImageLabels	[]build.ImageLabel
	NodeSelector	map[string]string
	Annotations	map[string]string
	Tolerations	[]core.Toleration
}
type OpenshiftAPIServerConfig struct {
	metav1.TypeMeta
	ServingInfo			HTTPServingInfo
	CORSAllowedOrigins		[]string
	MasterClients			MasterClients
	AuditConfig			AuditConfig
	StoragePrefix			string
	EtcdClientInfo			EtcdConnectionInfo
	ImagePolicyConfig		ServerImagePolicyConfig
	ProjectConfig			ServerProjectConfig
	RoutingConfig			RoutingConfig
	ServiceAccountOAuthGrantMethod	GrantHandlerType
	AdmissionPluginConfig		map[string]*AdmissionPluginConfig
	JenkinsPipelineConfig		JenkinsPipelineConfig
	APIServerArguments		ExtendedArguments
}
type ServerImagePolicyConfig struct {
	MaxImagesBulkImportedPerRepository	int
	AllowedRegistriesForImport		*AllowedRegistries
	InternalRegistryHostname		string
	ExternalRegistryHostname		string
	AdditionalTrustedCA			string
}
type ServerProjectConfig struct {
	DefaultNodeSelector	string
	ProjectRequestMessage	string
	ProjectRequestTemplate	string
}
type KubeAPIServerConfig struct {
	metav1.TypeMeta
	ServingInfo				HTTPServingInfo
	CORSAllowedOrigins			[]string
	OAuthConfig				*OAuthConfig
	AuthConfig				MasterAuthConfig
	AggregatorConfig			AggregatorConfig
	AuditConfig				AuditConfig
	StoragePrefix				string
	EtcdClientInfo				EtcdConnectionInfo
	KubeletClientInfo			KubeletConnectionInfo
	AdmissionPluginConfig			map[string]*AdmissionPluginConfig
	ServicesSubnet				string
	ServicesNodePortRange			string
	LegacyServiceServingCertSignerCABundle	string
	UserAgentMatchingConfig			UserAgentMatchingConfig
	ImagePolicyConfig			KubeAPIServerImagePolicyConfig
	ProjectConfig				KubeAPIServerProjectConfig
	ServiceAccountPublicKeyFiles		[]string
	APIServerArguments			ExtendedArguments
}
type KubeAPIServerImagePolicyConfig struct {
	InternalRegistryHostname	string
	ExternalRegistryHostname	string
}
type KubeAPIServerProjectConfig struct{ DefaultNodeSelector string }

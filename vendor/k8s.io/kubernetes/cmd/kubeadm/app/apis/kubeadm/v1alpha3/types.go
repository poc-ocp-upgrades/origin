package v1alpha3

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InitConfiguration struct {
	metav1.TypeMeta      `json:",inline"`
	ClusterConfiguration `json:"-"`
	BootstrapTokens      []BootstrapToken        `json:"bootstrapTokens,omitempty"`
	NodeRegistration     NodeRegistrationOptions `json:"nodeRegistration,omitempty"`
	APIEndpoint          APIEndpoint             `json:"apiEndpoint,omitempty"`
}
type ClusterConfiguration struct {
	metav1.TypeMeta               `json:",inline"`
	Etcd                          Etcd                     `json:"etcd"`
	Networking                    Networking               `json:"networking"`
	KubernetesVersion             string                   `json:"kubernetesVersion"`
	ControlPlaneEndpoint          string                   `json:"controlPlaneEndpoint"`
	APIServerExtraArgs            map[string]string        `json:"apiServerExtraArgs,omitempty"`
	ControllerManagerExtraArgs    map[string]string        `json:"controllerManagerExtraArgs,omitempty"`
	SchedulerExtraArgs            map[string]string        `json:"schedulerExtraArgs,omitempty"`
	APIServerExtraVolumes         []HostPathMount          `json:"apiServerExtraVolumes,omitempty"`
	ControllerManagerExtraVolumes []HostPathMount          `json:"controllerManagerExtraVolumes,omitempty"`
	SchedulerExtraVolumes         []HostPathMount          `json:"schedulerExtraVolumes,omitempty"`
	APIServerCertSANs             []string                 `json:"apiServerCertSANs,omitempty"`
	CertificatesDir               string                   `json:"certificatesDir"`
	ImageRepository               string                   `json:"imageRepository"`
	UnifiedControlPlaneImage      string                   `json:"unifiedControlPlaneImage"`
	AuditPolicyConfiguration      AuditPolicyConfiguration `json:"auditPolicy"`
	FeatureGates                  map[string]bool          `json:"featureGates,omitempty"`
	ClusterName                   string                   `json:"clusterName,omitempty"`
}
type ClusterStatus struct {
	metav1.TypeMeta `json:",inline"`
	APIEndpoints    map[string]APIEndpoint `json:"apiEndpoints"`
}
type APIEndpoint struct {
	AdvertiseAddress string `json:"advertiseAddress"`
	BindPort         int32  `json:"bindPort"`
}
type NodeRegistrationOptions struct {
	Name             string            `json:"name,omitempty"`
	CRISocket        string            `json:"criSocket,omitempty"`
	Taints           []v1.Taint        `json:"taints,omitempty"`
	KubeletExtraArgs map[string]string `json:"kubeletExtraArgs,omitempty"`
}
type Networking struct {
	ServiceSubnet string `json:"serviceSubnet"`
	PodSubnet     string `json:"podSubnet"`
	DNSDomain     string `json:"dnsDomain"`
}
type BootstrapToken struct {
	Token       *BootstrapTokenString `json:"token"`
	Description string                `json:"description,omitempty"`
	TTL         *metav1.Duration      `json:"ttl,omitempty"`
	Expires     *metav1.Time          `json:"expires,omitempty"`
	Usages      []string              `json:"usages,omitempty"`
	Groups      []string              `json:"groups,omitempty"`
}
type Etcd struct {
	Local    *LocalEtcd    `json:"local,omitempty"`
	External *ExternalEtcd `json:"external,omitempty"`
}
type LocalEtcd struct {
	Image          string            `json:"image"`
	DataDir        string            `json:"dataDir"`
	ExtraArgs      map[string]string `json:"extraArgs,omitempty"`
	ServerCertSANs []string          `json:"serverCertSANs,omitempty"`
	PeerCertSANs   []string          `json:"peerCertSANs,omitempty"`
}
type ExternalEtcd struct {
	Endpoints []string `json:"endpoints"`
	CAFile    string   `json:"caFile"`
	CertFile  string   `json:"certFile"`
	KeyFile   string   `json:"keyFile"`
}
type JoinConfiguration struct {
	metav1.TypeMeta                        `json:",inline"`
	NodeRegistration                       NodeRegistrationOptions `json:"nodeRegistration"`
	CACertPath                             string                  `json:"caCertPath"`
	DiscoveryFile                          string                  `json:"discoveryFile"`
	DiscoveryToken                         string                  `json:"discoveryToken"`
	DiscoveryTokenAPIServers               []string                `json:"discoveryTokenAPIServers,omitempty"`
	DiscoveryTimeout                       *metav1.Duration        `json:"discoveryTimeout,omitempty"`
	TLSBootstrapToken                      string                  `json:"tlsBootstrapToken"`
	Token                                  string                  `json:"token"`
	ClusterName                            string                  `json:"clusterName,omitempty"`
	DiscoveryTokenCACertHashes             []string                `json:"discoveryTokenCACertHashes,omitempty"`
	DiscoveryTokenUnsafeSkipCAVerification bool                    `json:"discoveryTokenUnsafeSkipCAVerification"`
	ControlPlane                           bool                    `json:"controlPlane,omitempty"`
	APIEndpoint                            APIEndpoint             `json:"apiEndpoint,omitempty"`
	FeatureGates                           map[string]bool         `json:"featureGates,omitempty"`
}
type HostPathMount struct {
	Name      string          `json:"name"`
	HostPath  string          `json:"hostPath"`
	MountPath string          `json:"mountPath"`
	Writable  bool            `json:"writable,omitempty"`
	PathType  v1.HostPathType `json:"pathType,omitempty"`
}
type AuditPolicyConfiguration struct {
	Path      string `json:"path"`
	LogDir    string `json:"logDir"`
	LogMaxAge *int32 `json:"logMaxAge,omitempty"`
}

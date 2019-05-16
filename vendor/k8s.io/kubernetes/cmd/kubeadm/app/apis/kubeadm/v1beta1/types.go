package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InitConfiguration struct {
	metav1.TypeMeta      `json:",inline"`
	ClusterConfiguration `json:"-"`
	BootstrapTokens      []BootstrapToken        `json:"bootstrapTokens,omitempty"`
	NodeRegistration     NodeRegistrationOptions `json:"nodeRegistration,omitempty"`
	LocalAPIEndpoint     APIEndpoint             `json:"localAPIEndpoint,omitempty"`
}
type ClusterConfiguration struct {
	metav1.TypeMeta      `json:",inline"`
	Etcd                 Etcd                  `json:"etcd"`
	Networking           Networking            `json:"networking"`
	KubernetesVersion    string                `json:"kubernetesVersion"`
	ControlPlaneEndpoint string                `json:"controlPlaneEndpoint"`
	APIServer            APIServer             `json:"apiServer,omitempty"`
	ControllerManager    ControlPlaneComponent `json:"controllerManager,omitempty"`
	Scheduler            ControlPlaneComponent `json:"scheduler,omitempty"`
	DNS                  DNS                   `json:"dns"`
	CertificatesDir      string                `json:"certificatesDir"`
	ImageRepository      string                `json:"imageRepository"`
	UseHyperKubeImage    bool                  `json:"useHyperKubeImage,omitempty"`
	FeatureGates         map[string]bool       `json:"featureGates,omitempty"`
	ClusterName          string                `json:"clusterName,omitempty"`
}
type ControlPlaneComponent struct {
	ExtraArgs    map[string]string `json:"extraArgs,omitempty"`
	ExtraVolumes []HostPathMount   `json:"extraVolumes,omitempty"`
}
type APIServer struct {
	ControlPlaneComponent  `json:",inline"`
	CertSANs               []string         `json:"certSANs,omitempty"`
	TimeoutForControlPlane *metav1.Duration `json:"timeoutForControlPlane,omitempty"`
}
type DNSAddOnType string

const (
	CoreDNS DNSAddOnType = "CoreDNS"
	KubeDNS DNSAddOnType = "kube-dns"
)

type DNS struct {
	Type      DNSAddOnType `json:"type"`
	ImageMeta `json:",inline"`
}
type ImageMeta struct {
	ImageRepository string `json:"imageRepository,omitempty"`
	ImageTag        string `json:"imageTag,omitempty"`
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
	ImageMeta      `json:",inline"`
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
	metav1.TypeMeta  `json:",inline"`
	NodeRegistration NodeRegistrationOptions `json:"nodeRegistration"`
	CACertPath       string                  `json:"caCertPath"`
	Discovery        Discovery               `json:"discovery"`
	ControlPlane     *JoinControlPlane       `json:"controlPlane,omitempty"`
}
type JoinControlPlane struct {
	LocalAPIEndpoint APIEndpoint `json:"localAPIEndpoint,omitempty"`
}
type Discovery struct {
	BootstrapToken    *BootstrapTokenDiscovery `json:"bootstrapToken,omitempty"`
	File              *FileDiscovery           `json:"file,omitempty"`
	TLSBootstrapToken string                   `json:"tlsBootstrapToken"`
	Timeout           *metav1.Duration         `json:"timeout,omitempty"`
}
type BootstrapTokenDiscovery struct {
	Token                    string   `json:"token"`
	APIServerEndpoint        string   `json:"apiServerEndpoint,omitempty"`
	CACertHashes             []string `json:"caCertHashes,omitempty"`
	UnsafeSkipCAVerification bool     `json:"unsafeSkipCAVerification"`
}
type FileDiscovery struct {
	KubeConfigPath string `json:"kubeConfigPath"`
}
type HostPathMount struct {
	Name      string          `json:"name"`
	HostPath  string          `json:"hostPath"`
	MountPath string          `json:"mountPath"`
	ReadOnly  bool            `json:"readOnly,omitempty"`
	PathType  v1.HostPathType `json:"pathType,omitempty"`
}

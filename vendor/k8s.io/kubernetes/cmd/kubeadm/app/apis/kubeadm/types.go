package kubeadm

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeletconfig "k8s.io/kubernetes/pkg/kubelet/apis/config"
	kubeproxyconfig "k8s.io/kubernetes/pkg/proxy/apis/config"
)

type InitConfiguration struct {
	metav1.TypeMeta
	ClusterConfiguration `json:",inline"`
	BootstrapTokens      []BootstrapToken
	NodeRegistration     NodeRegistrationOptions
	LocalAPIEndpoint     APIEndpoint
}
type ClusterConfiguration struct {
	metav1.TypeMeta
	ComponentConfigs     ComponentConfigs
	Etcd                 Etcd
	Networking           Networking
	KubernetesVersion    string
	ControlPlaneEndpoint string
	APIServer            APIServer
	ControllerManager    ControlPlaneComponent
	Scheduler            ControlPlaneComponent
	DNS                  DNS
	CertificatesDir      string
	ImageRepository      string
	CIImageRepository    string
	UseHyperKubeImage    bool
	FeatureGates         map[string]bool
	ClusterName          string
}
type ControlPlaneComponent struct {
	ExtraArgs    map[string]string
	ExtraVolumes []HostPathMount
}
type APIServer struct {
	ControlPlaneComponent
	CertSANs               []string
	TimeoutForControlPlane *metav1.Duration
}
type DNSAddOnType string

const (
	CoreDNS DNSAddOnType = "CoreDNS"
	KubeDNS DNSAddOnType = "kube-dns"
)

type DNS struct {
	Type      DNSAddOnType
	ImageMeta `json:",inline"`
}
type ImageMeta struct {
	ImageRepository string
	ImageTag        string
}
type ComponentConfigs struct {
	Kubelet   *kubeletconfig.KubeletConfiguration
	KubeProxy *kubeproxyconfig.KubeProxyConfiguration
}
type ClusterStatus struct {
	metav1.TypeMeta
	APIEndpoints map[string]APIEndpoint
}
type APIEndpoint struct {
	AdvertiseAddress string
	BindPort         int32
}
type NodeRegistrationOptions struct {
	Name             string
	CRISocket        string
	Taints           []v1.Taint
	KubeletExtraArgs map[string]string
}
type Networking struct {
	ServiceSubnet string
	PodSubnet     string
	DNSDomain     string
}
type BootstrapToken struct {
	Token       *BootstrapTokenString
	Description string
	TTL         *metav1.Duration
	Expires     *metav1.Time
	Usages      []string
	Groups      []string
}
type Etcd struct {
	Local    *LocalEtcd
	External *ExternalEtcd
}
type LocalEtcd struct {
	ImageMeta      `json:",inline"`
	DataDir        string
	ExtraArgs      map[string]string
	ServerCertSANs []string
	PeerCertSANs   []string
}
type ExternalEtcd struct {
	Endpoints []string
	CAFile    string
	CertFile  string
	KeyFile   string
}
type JoinConfiguration struct {
	metav1.TypeMeta
	NodeRegistration NodeRegistrationOptions
	CACertPath       string
	Discovery        Discovery
	ControlPlane     *JoinControlPlane
}
type JoinControlPlane struct{ LocalAPIEndpoint APIEndpoint }
type Discovery struct {
	BootstrapToken    *BootstrapTokenDiscovery
	File              *FileDiscovery
	TLSBootstrapToken string
	Timeout           *metav1.Duration
}
type BootstrapTokenDiscovery struct {
	Token                    string
	APIServerEndpoint        string
	CACertHashes             []string
	UnsafeSkipCAVerification bool
}
type FileDiscovery struct{ KubeConfigPath string }

func (cfg *ClusterConfiguration) GetControlPlaneImageRepository() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg.CIImageRepository != "" {
		return cfg.CIImageRepository
	}
	return cfg.ImageRepository
}

type HostPathMount struct {
	Name      string
	HostPath  string
	MountPath string
	ReadOnly  bool
	PathType  v1.HostPathType
}
type CommonConfiguration interface {
	GetCRISocket() string
	GetNodeName() string
	GetKubernetesVersion() string
}

func (cfg *InitConfiguration) GetCRISocket() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cfg.NodeRegistration.CRISocket
}
func (cfg *InitConfiguration) GetNodeName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cfg.NodeRegistration.Name
}
func (cfg *InitConfiguration) GetKubernetesVersion() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cfg.KubernetesVersion
}
func (cfg *JoinConfiguration) GetCRISocket() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cfg.NodeRegistration.CRISocket
}
func (cfg *JoinConfiguration) GetNodeName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cfg.NodeRegistration.Name
}
func (cfg *JoinConfiguration) GetKubernetesVersion() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}

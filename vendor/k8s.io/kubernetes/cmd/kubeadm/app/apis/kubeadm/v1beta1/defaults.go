package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"net/url"
	"time"
)

const (
	DefaultServiceDNSDomain   = "cluster.local"
	DefaultServicesSubnet     = "10.96.0.0/12"
	DefaultClusterDNSIP       = "10.96.0.10"
	DefaultKubernetesVersion  = "stable-1"
	DefaultAPIBindPort        = 6443
	DefaultCertificatesDir    = "/etc/kubernetes/pki"
	DefaultImageRepository    = "k8s.gcr.io"
	DefaultManifestsDir       = "/etc/kubernetes/manifests"
	DefaultClusterName        = "kubernetes"
	DefaultEtcdDataDir        = "/var/lib/etcd"
	DefaultProxyBindAddressv4 = "0.0.0.0"
	DefaultProxyBindAddressv6 = "::"
	DefaultDiscoveryTimeout   = 5 * time.Minute
)

var (
	DefaultAuditPolicyLogMaxAge = int32(2)
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return RegisterDefaults(scheme)
}
func SetDefaults_InitConfiguration(obj *InitConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	SetDefaults_ClusterConfiguration(&obj.ClusterConfiguration)
	SetDefaults_NodeRegistrationOptions(&obj.NodeRegistration)
	SetDefaults_BootstrapTokens(obj)
	SetDefaults_APIEndpoint(&obj.LocalAPIEndpoint)
}
func SetDefaults_ClusterConfiguration(obj *ClusterConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.KubernetesVersion == "" {
		obj.KubernetesVersion = DefaultKubernetesVersion
	}
	if obj.Networking.ServiceSubnet == "" {
		obj.Networking.ServiceSubnet = DefaultServicesSubnet
	}
	if obj.Networking.DNSDomain == "" {
		obj.Networking.DNSDomain = DefaultServiceDNSDomain
	}
	if obj.CertificatesDir == "" {
		obj.CertificatesDir = DefaultCertificatesDir
	}
	if obj.ImageRepository == "" {
		obj.ImageRepository = DefaultImageRepository
	}
	if obj.ClusterName == "" {
		obj.ClusterName = DefaultClusterName
	}
	SetDefaults_DNS(obj)
	SetDefaults_Etcd(obj)
	SetDefaults_APIServer(&obj.APIServer)
}
func SetDefaults_APIServer(obj *APIServer) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.TimeoutForControlPlane == nil {
		obj.TimeoutForControlPlane = &metav1.Duration{Duration: constants.DefaultControlPlaneTimeout}
	}
}
func SetDefaults_DNS(obj *ClusterConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.DNS.Type == "" {
		obj.DNS.Type = CoreDNS
	}
}
func SetDefaults_Etcd(obj *ClusterConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Etcd.External == nil && obj.Etcd.Local == nil {
		obj.Etcd.Local = &LocalEtcd{}
	}
	if obj.Etcd.Local != nil {
		if obj.Etcd.Local.DataDir == "" {
			obj.Etcd.Local.DataDir = DefaultEtcdDataDir
		}
	}
}
func SetDefaults_JoinConfiguration(obj *JoinConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.CACertPath == "" {
		obj.CACertPath = DefaultCACertPath
	}
	SetDefaults_NodeRegistrationOptions(&obj.NodeRegistration)
	SetDefaults_JoinControlPlane(obj.ControlPlane)
	SetDefaults_Discovery(&obj.Discovery)
}
func SetDefaults_NodeRegistrationOptions(obj *NodeRegistrationOptions) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.CRISocket == "" {
		obj.CRISocket = DefaultCRISocket
	}
}
func SetDefaults_JoinControlPlane(obj *JoinControlPlane) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj != nil {
		SetDefaults_APIEndpoint(&obj.LocalAPIEndpoint)
	}
}
func SetDefaults_Discovery(obj *Discovery) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(obj.TLSBootstrapToken) == 0 && obj.BootstrapToken != nil {
		obj.TLSBootstrapToken = obj.BootstrapToken.Token
	}
	if obj.Timeout == nil {
		obj.Timeout = &metav1.Duration{Duration: DefaultDiscoveryTimeout}
	}
	if obj.File != nil {
		SetDefaults_FileDiscovery(obj.File)
	}
}
func SetDefaults_FileDiscovery(obj *FileDiscovery) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(obj.KubeConfigPath) != 0 {
		u, err := url.Parse(obj.KubeConfigPath)
		if err == nil && u.Scheme == "file" {
			obj.KubeConfigPath = u.Path
		}
	}
}
func SetDefaults_BootstrapTokens(obj *InitConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.BootstrapTokens == nil || len(obj.BootstrapTokens) == 0 {
		obj.BootstrapTokens = []BootstrapToken{{}}
	}
	for i := range obj.BootstrapTokens {
		SetDefaults_BootstrapToken(&obj.BootstrapTokens[i])
	}
}
func SetDefaults_BootstrapToken(bt *BootstrapToken) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if bt.TTL == nil {
		bt.TTL = &metav1.Duration{Duration: constants.DefaultTokenDuration}
	}
	if len(bt.Usages) == 0 {
		bt.Usages = constants.DefaultTokenUsages
	}
	if len(bt.Groups) == 0 {
		bt.Groups = constants.DefaultTokenGroups
	}
}
func SetDefaults_APIEndpoint(obj *APIEndpoint) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.BindPort == 0 {
		obj.BindPort = DefaultAPIBindPort
	}
}

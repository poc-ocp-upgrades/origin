package constants

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/version"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
	"net"
	"os"
	goos "os"
	"path"
	"path/filepath"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

var KubernetesDir = "/etc/kubernetes"

const (
	ManifestsSubDirName                      = "manifests"
	TempDirForKubeadm                        = "tmp"
	CACertAndKeyBaseName                     = "ca"
	CACertName                               = "ca.crt"
	CAKeyName                                = "ca.key"
	APIServerCertAndKeyBaseName              = "apiserver"
	APIServerCertName                        = "apiserver.crt"
	APIServerKeyName                         = "apiserver.key"
	APIServerCertCommonName                  = "kube-apiserver"
	APIServerKubeletClientCertAndKeyBaseName = "apiserver-kubelet-client"
	APIServerKubeletClientCertName           = "apiserver-kubelet-client.crt"
	APIServerKubeletClientKeyName            = "apiserver-kubelet-client.key"
	APIServerKubeletClientCertCommonName     = "kube-apiserver-kubelet-client"
	EtcdCACertAndKeyBaseName                 = "etcd/ca"
	EtcdCACertName                           = "etcd/ca.crt"
	EtcdCAKeyName                            = "etcd/ca.key"
	EtcdServerCertAndKeyBaseName             = "etcd/server"
	EtcdServerCertName                       = "etcd/server.crt"
	EtcdServerKeyName                        = "etcd/server.key"
	EtcdListenClientPort                     = 2379
	EtcdPeerCertAndKeyBaseName               = "etcd/peer"
	EtcdPeerCertName                         = "etcd/peer.crt"
	EtcdPeerKeyName                          = "etcd/peer.key"
	EtcdListenPeerPort                       = 2380
	EtcdHealthcheckClientCertAndKeyBaseName  = "etcd/healthcheck-client"
	EtcdHealthcheckClientCertName            = "etcd/healthcheck-client.crt"
	EtcdHealthcheckClientKeyName             = "etcd/healthcheck-client.key"
	EtcdHealthcheckClientCertCommonName      = "kube-etcd-healthcheck-client"
	APIServerEtcdClientCertAndKeyBaseName    = "apiserver-etcd-client"
	APIServerEtcdClientCertName              = "apiserver-etcd-client.crt"
	APIServerEtcdClientKeyName               = "apiserver-etcd-client.key"
	APIServerEtcdClientCertCommonName        = "kube-apiserver-etcd-client"
	ServiceAccountKeyBaseName                = "sa"
	ServiceAccountPublicKeyName              = "sa.pub"
	ServiceAccountPrivateKeyName             = "sa.key"
	FrontProxyCACertAndKeyBaseName           = "front-proxy-ca"
	FrontProxyCACertName                     = "front-proxy-ca.crt"
	FrontProxyCAKeyName                      = "front-proxy-ca.key"
	FrontProxyClientCertAndKeyBaseName       = "front-proxy-client"
	FrontProxyClientCertName                 = "front-proxy-client.crt"
	FrontProxyClientKeyName                  = "front-proxy-client.key"
	FrontProxyClientCertCommonName           = "front-proxy-client"
	AdminKubeConfigFileName                  = "admin.conf"
	KubeletBootstrapKubeConfigFileName       = "bootstrap-kubelet.conf"
	KubeletKubeConfigFileName                = "kubelet.conf"
	ControllerManagerKubeConfigFileName      = "controller-manager.conf"
	SchedulerKubeConfigFileName              = "scheduler.conf"
	ControllerManagerUser                    = "system:kube-controller-manager"
	SchedulerUser                            = "system:kube-scheduler"
	MastersGroup                             = "system:masters"
	NodesGroup                               = "system:nodes"
	NodesUserPrefix                          = "system:node:"
	NodesClusterRoleBinding                  = "system:node"
	APICallRetryInterval                     = 500 * time.Millisecond
	DiscoveryRetryInterval                   = 5 * time.Second
	PatchNodeTimeout                         = 2 * time.Minute
	UpdateNodeTimeout                        = 2 * time.Minute
	TLSBootstrapTimeout                      = 2 * time.Minute
	DefaultControlPlaneTimeout               = 4 * time.Minute
	MinimumAddressesInServiceSubnet          = 10
	DefaultTokenDuration                     = 24 * time.Hour
	LabelNodeRoleMaster                      = "node-role.kubernetes.io/master"
	AnnotationKubeadmCRISocket               = "kubeadm.alpha.kubernetes.io/cri-socket"
	KubeadmConfigConfigMap                   = "kubeadm-config"
	ClusterConfigurationConfigMapKey         = "ClusterConfiguration"
	ClusterStatusConfigMapKey                = "ClusterStatus"
	KubeProxyConfigMap                       = "kube-proxy"
	KubeProxyConfigMapKey                    = "config.conf"
	KubeletBaseConfigurationConfigMapPrefix  = "kubelet-config-"
	KubeletBaseConfigurationConfigMapKey     = "kubelet"
	KubeletBaseConfigMapRolePrefix           = "kubeadm:kubelet-config-"
	KubeletRunDirectory                      = "/var/lib/kubelet"
	KubeletConfigurationFileName             = "config.yaml"
	DynamicKubeletConfigurationDirectoryName = "dynamic-config"
	KubeletEnvFileName                       = "kubeadm-flags.env"
	KubeletEnvFileVariableName               = "KUBELET_KUBEADM_ARGS"
	KubeletHealthzPort                       = 10248
	MinExternalEtcdVersion                   = "3.2.18"
	DefaultEtcdVersion                       = "3.2.24"
	PauseVersion                             = "3.1"
	Etcd                                     = "etcd"
	KubeAPIServer                            = "kube-apiserver"
	KubeControllerManager                    = "kube-controller-manager"
	KubeScheduler                            = "kube-scheduler"
	KubeProxy                                = "kube-proxy"
	HyperKube                                = "hyperkube"
	SelfHostingPrefix                        = "self-hosted-"
	KubeCertificatesVolumeName               = "k8s-certs"
	KubeConfigVolumeName                     = "kubeconfig"
	NodeBootstrapTokenAuthGroup              = "system:bootstrappers:kubeadm:default-node-token"
	DefaultCIImageRepository                 = "gcr.io/kubernetes-ci-images"
	CoreDNSConfigMap                         = "coredns"
	CoreDNSDeploymentName                    = "coredns"
	CoreDNSImageName                         = "coredns"
	KubeDNSConfigMap                         = "kube-dns"
	KubeDNSDeploymentName                    = "kube-dns"
	KubeDNSKubeDNSImageName                  = "k8s-dns-kube-dns"
	KubeDNSSidecarImageName                  = "k8s-dns-sidecar"
	KubeDNSDnsMasqNannyImageName             = "k8s-dns-dnsmasq-nanny"
	CRICtlPackage                            = "github.com/kubernetes-incubator/cri-tools/cmd/crictl"
	KubeAuditPolicyVolumeName                = "audit"
	AuditPolicyDir                           = "audit"
	AuditPolicyFile                          = "audit.yaml"
	AuditPolicyLogFile                       = "audit.log"
	KubeAuditPolicyLogVolumeName             = "audit-log"
	StaticPodAuditPolicyLogDir               = "/var/log/kubernetes/audit"
	LeaseEndpointReconcilerType              = "lease"
	KubeDNSVersion                           = "1.14.13"
	CoreDNSVersion                           = "1.2.6"
	ClusterConfigurationKind                 = "ClusterConfiguration"
	InitConfigurationKind                    = "InitConfiguration"
	JoinConfigurationKind                    = "JoinConfiguration"
	YAMLDocumentSeparator                    = "---\n"
	DefaultAPIServerBindAddress              = "0.0.0.0"
	MasterNumCPU                             = 2
)

var (
	MasterTaint                = v1.Taint{Key: LabelNodeRoleMaster, Effect: v1.TaintEffectNoSchedule}
	MasterToleration           = v1.Toleration{Key: LabelNodeRoleMaster, Effect: v1.TaintEffectNoSchedule}
	DefaultTokenUsages         = bootstrapapi.KnownTokenUsages
	DefaultTokenGroups         = []string{NodeBootstrapTokenAuthGroup}
	MasterComponents           = []string{KubeAPIServer, KubeControllerManager, KubeScheduler}
	MinimumControlPlaneVersion = version.MustParseSemantic("v1.12.0")
	MinimumKubeletVersion      = version.MustParseSemantic("v1.12.0")
	SupportedEtcdVersion       = map[uint8]string{10: "3.1.12", 11: "3.2.18", 12: "3.2.24", 13: "3.2.24"}
)

func EtcdSupportedVersion(versionString string) (*version.Version, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubernetesVersion, err := version.ParseSemantic(versionString)
	if err != nil {
		return nil, err
	}
	if etcdStringVersion, ok := SupportedEtcdVersion[uint8(kubernetesVersion.Minor())]; ok {
		etcdVersion, err := version.ParseSemantic(etcdStringVersion)
		if err != nil {
			return nil, err
		}
		return etcdVersion, nil
	}
	return nil, errors.Errorf("Unsupported or unknown Kubernetes version(%v)", kubernetesVersion)
}
func GetStaticPodDirectory() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return filepath.Join(KubernetesDir, ManifestsSubDirName)
}
func GetStaticPodFilepath(componentName, manifestsDir string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return filepath.Join(manifestsDir, componentName+".yaml")
}
func GetAdminKubeConfigPath() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return filepath.Join(KubernetesDir, AdminKubeConfigFileName)
}
func GetBootstrapKubeletKubeConfigPath() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return filepath.Join(KubernetesDir, KubeletBootstrapKubeConfigFileName)
}
func GetKubeletKubeConfigPath() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return filepath.Join(KubernetesDir, KubeletKubeConfigFileName)
}
func AddSelfHostedPrefix(componentName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s%s", SelfHostingPrefix, componentName)
}
func CreateTempDirForKubeadm(dirName string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tempDir := path.Join(KubernetesDir, TempDirForKubeadm)
	if err := os.MkdirAll(tempDir, 0700); err != nil {
		return "", errors.Wrapf(err, "failed to create directory %q", tempDir)
	}
	tempDir, err := ioutil.TempDir(tempDir, dirName)
	if err != nil {
		return "", errors.Wrap(err, "couldn't create a temporary directory")
	}
	return tempDir, nil
}
func CreateTimestampDirForKubeadm(dirName string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tempDir := path.Join(KubernetesDir, TempDirForKubeadm)
	if err := os.MkdirAll(tempDir, 0700); err != nil {
		return "", errors.Wrapf(err, "failed to create directory %q", tempDir)
	}
	timestampDirName := fmt.Sprintf("%s-%s", dirName, time.Now().Format("2006-01-02-15-04-05"))
	timestampDir := path.Join(tempDir, timestampDirName)
	if err := os.Mkdir(timestampDir, 0700); err != nil {
		return "", errors.Wrap(err, "could not create timestamp directory")
	}
	return timestampDir, nil
}
func GetDNSIP(svcSubnet string) (net.IP, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, svcSubnetCIDR, err := net.ParseCIDR(svcSubnet)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't parse service subnet CIDR %q", svcSubnet)
	}
	dnsIP, err := ipallocator.GetIndexedIP(svcSubnetCIDR, 10)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get tenth IP address from service subnet CIDR %s", svcSubnetCIDR.String())
	}
	return dnsIP, nil
}
func GetStaticPodAuditPolicyFile() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return filepath.Join(KubernetesDir, AuditPolicyDir, AuditPolicyFile)
}
func GetDNSVersion(dnsType kubeadmapi.DNSAddOnType) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch dnsType {
	case kubeadmapi.KubeDNS:
		return KubeDNSVersion
	default:
		return CoreDNSVersion
	}
}
func GetKubeletConfigMapName(k8sVersion *version.Version) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s%d.%d", KubeletBaseConfigurationConfigMapPrefix, k8sVersion.Major(), k8sVersion.Minor())
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

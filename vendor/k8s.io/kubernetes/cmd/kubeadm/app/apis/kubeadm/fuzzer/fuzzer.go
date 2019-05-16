package fuzzer

import (
	goformat "fmt"
	fuzz "github.com/google/gofuzz"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func Funcs(codecs runtimeserializer.CodecFactory) []interface{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []interface{}{fuzzInitConfiguration, fuzzClusterConfiguration, fuzzComponentConfigs, fuzzNodeRegistration, fuzzDNS, fuzzLocalEtcd, fuzzNetworking, fuzzJoinConfiguration}
}
func fuzzInitConfiguration(obj *kubeadm.InitConfiguration, c fuzz.Continue) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.FuzzNoCustom(obj)
	obj.ClusterConfiguration = kubeadm.ClusterConfiguration{APIServer: kubeadm.APIServer{TimeoutForControlPlane: &metav1.Duration{Duration: constants.DefaultControlPlaneTimeout}}, DNS: kubeadm.DNS{Type: kubeadm.CoreDNS}, CertificatesDir: v1beta1.DefaultCertificatesDir, ClusterName: v1beta1.DefaultClusterName, Etcd: kubeadm.Etcd{Local: &kubeadm.LocalEtcd{DataDir: v1beta1.DefaultEtcdDataDir}}, ImageRepository: v1beta1.DefaultImageRepository, KubernetesVersion: v1beta1.DefaultKubernetesVersion, Networking: kubeadm.Networking{ServiceSubnet: v1beta1.DefaultServicesSubnet, DNSDomain: v1beta1.DefaultServiceDNSDomain}}
	obj.BootstrapTokens = []kubeadm.BootstrapToken{{Groups: []string{"foo"}, TTL: &metav1.Duration{Duration: 1234}, Usages: []string{"foo"}}}
}
func fuzzNodeRegistration(obj *kubeadm.NodeRegistrationOptions, c fuzz.Continue) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.FuzzNoCustom(obj)
	obj.CRISocket = "foo"
}
func fuzzClusterConfiguration(obj *kubeadm.ClusterConfiguration, c fuzz.Continue) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.FuzzNoCustom(obj)
	obj.CertificatesDir = "foo"
	obj.CIImageRepository = ""
	obj.ClusterName = "bar"
	obj.ImageRepository = "baz"
	obj.KubernetesVersion = "qux"
	obj.APIServer.TimeoutForControlPlane = &metav1.Duration{Duration: constants.DefaultControlPlaneTimeout}
}
func fuzzDNS(obj *kubeadm.DNS, c fuzz.Continue) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj.Type = kubeadm.CoreDNS
}
func fuzzComponentConfigs(obj *kubeadm.ComponentConfigs, c fuzz.Continue) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func fuzzLocalEtcd(obj *kubeadm.LocalEtcd, c fuzz.Continue) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.FuzzNoCustom(obj)
	obj.DataDir = "foo"
	obj.ImageRepository = ""
	obj.ImageTag = ""
}
func fuzzNetworking(obj *kubeadm.Networking, c fuzz.Continue) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.FuzzNoCustom(obj)
	obj.DNSDomain = "foo"
	obj.ServiceSubnet = "bar"
}
func fuzzJoinConfiguration(obj *kubeadm.JoinConfiguration, c fuzz.Continue) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.FuzzNoCustom(obj)
	obj.CACertPath = "foo"
	obj.Discovery = kubeadm.Discovery{BootstrapToken: &kubeadm.BootstrapTokenDiscovery{Token: "baz"}, TLSBootstrapToken: "qux", Timeout: &metav1.Duration{Duration: 1234}}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

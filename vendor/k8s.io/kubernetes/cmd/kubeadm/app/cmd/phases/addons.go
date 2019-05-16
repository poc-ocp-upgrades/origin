package phases

import (
	goformat "fmt"
	"github.com/pkg/errors"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	dnsaddon "k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/dns"
	proxyaddon "k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/proxy"
	"k8s.io/kubernetes/pkg/util/normalizer"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var (
	coreDNSAddonLongDesc = normalizer.LongDesc(`
		Installs the CoreDNS addon components via the API server.
		Please note that although the DNS server is deployed, it will not be scheduled until CNI is installed.
		`)
	kubeProxyAddonLongDesc = normalizer.LongDesc(`
		Installs the kube-proxy addon components via the API server.
		`)
)

type addonData interface {
	Cfg() *kubeadmapi.InitConfiguration
	Client() (clientset.Interface, error)
}

func NewAddonPhase() workflow.Phase {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return workflow.Phase{Name: "addon", Short: "Installs required addons for passing Conformance tests", Long: cmdutil.MacroCommandLongDescription, Phases: []workflow.Phase{{Name: "all", Short: "Installs all the addons", InheritFlags: getAddonPhaseFlags("all"), RunAllSiblings: true}, {Name: "coredns", Short: "Installs the CoreDNS addon to a Kubernetes cluster", Long: coreDNSAddonLongDesc, InheritFlags: getAddonPhaseFlags("coredns"), Run: runCoreDNSAddon}, {Name: "kube-proxy", Short: "Installs the kube-proxy addon to a Kubernetes cluster", Long: kubeProxyAddonLongDesc, InheritFlags: getAddonPhaseFlags("kube-proxy"), Run: runKubeProxyAddon}}}
}
func getAddonData(c workflow.RunData) (*kubeadmapi.InitConfiguration, clientset.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, ok := c.(addonData)
	if !ok {
		return nil, nil, errors.New("addon phase invoked with an invalid data struct")
	}
	cfg := data.Cfg()
	client, err := data.Client()
	if err != nil {
		return nil, nil, err
	}
	return cfg, client, err
}
func runCoreDNSAddon(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg, client, err := getAddonData(c)
	if err != nil {
		return err
	}
	return dnsaddon.EnsureDNSAddon(cfg, client)
}
func runKubeProxyAddon(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg, client, err := getAddonData(c)
	if err != nil {
		return err
	}
	return proxyaddon.EnsureProxyAddon(cfg, client)
}
func getAddonPhaseFlags(name string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flags := []string{options.CfgPath, options.KubeconfigPath, options.KubernetesVersion, options.ImageRepository}
	if name == "all" || name == "kube-proxy" {
		flags = append(flags, options.APIServerAdvertiseAddress, options.APIServerBindPort, options.NetworkingPodSubnet)
	}
	if name == "all" || name == "coredns" {
		flags = append(flags, options.FeatureGatesString, options.NetworkingDNSDomain, options.NetworkingServiceSubnet)
	}
	return flags
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

package openshift_sdn

import (
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/openshift/library-go/pkg/serviceability"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	_ "github.com/openshift/origin/pkg/cmd/server/apis/config/install"
	configapilatest "github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	networkvalidation "github.com/openshift/origin/pkg/cmd/server/apis/config/validation/network"
	sdnnode "github.com/openshift/origin/pkg/network/node"
	sdnproxy "github.com/openshift/origin/pkg/network/proxy"
	"github.com/openshift/origin/pkg/version"
	"github.com/spf13/cobra"
	"io"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	kubeproxyconfig "k8s.io/kubernetes/pkg/proxy/apis/config"
	"k8s.io/kubernetes/pkg/util/interrupt"
	"net/url"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

const openshiftCNIFile string = "80-openshift-network.conf"

type OpenShiftSDN struct {
	ConfigFilePath            string
	KubeConfigFilePath        string
	URLOnlyKubeConfigFilePath string
	cniConfFile               string
	NodeConfig                *configapi.NodeConfig
	ProxyConfig               *kubeproxyconfig.KubeProxyConfiguration
	informers                 *informers
	OsdnNode                  *sdnnode.OsdnNode
	sdnRecorder               record.EventRecorder
	OsdnProxy                 *sdnproxy.OsdnProxy
}

var networkLong = `
Start OpenShift SDN node components. This includes the service proxy.

This will also read the node name from the environment variable K8S_NODE_NAME.`

func NewOpenShiftSDNCommand(basename string, errout io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sdn := &OpenShiftSDN{}
	cmd := &cobra.Command{Use: basename, Short: "Start OpenShiftSDN", Long: networkLong, Run: func(c *cobra.Command, _ []string) {
		ch := make(chan struct{})
		interrupt.New(func(s os.Signal) {
			fmt.Fprintf(errout, "interrupt: Gracefully shutting down ...\n")
			close(ch)
		}).Run(func() error {
			sdn.Run(c, errout, ch)
			return nil
		})
	}}
	flags := cmd.Flags()
	flags.StringVar(&sdn.ConfigFilePath, "config", "", "Location of the node configuration file to run from (required)")
	flags.StringVar(&sdn.KubeConfigFilePath, "kubeconfig", "", "Path to the kubeconfig file to use for requests to the Kubernetes API. Optional. When omitted, will use the in-cluster config")
	flags.StringVar(&sdn.URLOnlyKubeConfigFilePath, "url-only-kubeconfig", "", "Path to a kubeconfig file to use, but only to determine the URL to the apiserver. The in-cluster credentials will be used. Cannot use with --kubeconfig.")
	return cmd
}
func (sdn *OpenShiftSDN) Run(c *cobra.Command, errout io.Writer, stopCh chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := injectKubeAPIEnv(sdn.URLOnlyKubeConfigFilePath)
	if err != nil {
		klog.Fatal(err)
	}
	err = sdn.ValidateAndParse()
	if err != nil {
		if kerrors.IsInvalid(err) {
			if details := err.(*kerrors.StatusError).ErrStatus.Details; details != nil {
				fmt.Fprintf(errout, "Invalid %s %s\n", details.Kind, details.Name)
				for _, cause := range details.Causes {
					fmt.Fprintf(errout, "  %s: %s\n", cause.Field, cause.Message)
				}
				os.Exit(255)
			}
		}
		klog.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Base(sdn.cniConfFile), 0755); err != nil {
		klog.Fatal(err)
	}
	if err := watchForChanges(sdn.ConfigFilePath, stopCh); err != nil {
		klog.Fatalf("unable to setup configuration watch: %v", err)
	}
	err = sdn.Init()
	if err != nil {
		klog.Fatalf("Failed to initialize sdn: %v", err)
	}
	err = sdn.Start(stopCh)
	if err != nil {
		klog.Fatalf("Failed to start sdn: %v", err)
	}
	<-stopCh
	time.Sleep(500 * time.Millisecond)
}
func (sdn *OpenShiftSDN) ValidateAndParse() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(sdn.ConfigFilePath) == 0 {
		return errors.New("--config is required")
	}
	if len(sdn.KubeConfigFilePath) > 0 && len(sdn.URLOnlyKubeConfigFilePath) > 0 {
		return errors.New("cannot pass --kubeconfig and --url-only-kubeconfig")
	}
	klog.V(2).Infof("Reading node configuration from %s", sdn.ConfigFilePath)
	var err error
	sdn.NodeConfig, err = configapilatest.ReadAndResolveNodeConfig(sdn.ConfigFilePath)
	if err != nil {
		return err
	}
	if len(sdn.KubeConfigFilePath) > 0 {
		sdn.NodeConfig.MasterKubeConfig = sdn.KubeConfigFilePath
	}
	if len(sdn.NodeConfig.NodeName) == 0 {
		sdn.NodeConfig.NodeName = os.Getenv("K8S_NODE_NAME")
	}
	validationResults := networkvalidation.ValidateInClusterNetworkNodeConfig(sdn.NodeConfig, nil)
	if len(validationResults.Warnings) != 0 {
		for _, warning := range validationResults.Warnings {
			klog.Warningf("Warning: %v, node start will continue.", warning)
		}
	}
	if len(validationResults.Errors) != 0 {
		klog.V(4).Infof("Configuration is invalid: %#v", sdn.NodeConfig)
		return kerrors.NewInvalid(configapi.Kind("NodeConfig"), sdn.ConfigFilePath, validationResults.Errors)
	}
	sdn.ProxyConfig, err = ProxyConfigFromNodeConfig(*sdn.NodeConfig)
	if err != nil {
		klog.V(4).Infof("Unable to build proxy config: %v", err)
		return err
	}
	cniConfDir := "/etc/cni/net.d"
	if val, ok := sdn.NodeConfig.KubeletArguments["cni-conf-dir"]; ok && len(val) == 1 {
		cniConfDir = val[0]
	}
	sdn.cniConfFile = filepath.Join(cniConfDir, openshiftCNIFile)
	return nil
}
func (sdn *OpenShiftSDN) Init() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	err = sdn.buildInformers()
	if err != nil {
		return fmt.Errorf("failed to build informers: %v", err)
	}
	err = sdn.initSDN()
	if err != nil {
		return fmt.Errorf("failed to initialize SDN: %v", err)
	}
	err = sdn.initProxy()
	if err != nil {
		return fmt.Errorf("failed to initialize proxy: %v", err)
	}
	return nil
}
func (sdn *OpenShiftSDN) Start(stopCh <-chan struct{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Infof("Starting node networking (%s)", version.Get().String())
	serviceability.StartProfiler()
	err := sdn.runSDN()
	if err != nil {
		return err
	}
	proxyInitChan := make(chan bool)
	sdn.runProxy(proxyInitChan)
	sdn.informers.start(stopCh)
	klog.V(2).Infof("openshift-sdn network plugin waiting for proxy startup to comlete")
	<-proxyInitChan
	klog.V(2).Infof("openshift-sdn network plugin registering startup")
	if err := sdn.writeConfigFile(); err != nil {
		klog.Fatal(err)
	}
	klog.V(2).Infof("openshift-sdn network plugin ready")
	return nil
}
func injectKubeAPIEnv(kcPath string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if kcPath != "" {
		kubeconfig, err := clientcmd.LoadFromFile(kcPath)
		if err != nil {
			return err
		}
		clusterName := kubeconfig.Contexts[kubeconfig.CurrentContext].Cluster
		apiURL := kubeconfig.Clusters[clusterName].Server
		url, err := url.Parse(apiURL)
		if err != nil {
			return err
		}
		klog.V(2).Infof("Overriding kubernetes api to %s", apiURL)
		os.Setenv("KUBERNETES_SERVICE_HOST", url.Hostname())
		os.Setenv("KUBERNETES_SERVICE_PORT", url.Port())
	}
	return nil
}
func watchForChanges(configPath string, stopCh chan struct{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configPath, err := filepath.Abs(configPath)
	if err != nil {
		return err
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	p := configPath
	maxdepth := 100
	for depth := 0; depth < maxdepth; depth++ {
		if err := watcher.Add(p); err != nil {
			return err
		}
		klog.V(2).Infof("Watching config file %s for changes", p)
		stat, err := os.Lstat(p)
		if err != nil {
			return err
		}
		if stat.Mode()&os.ModeSymlink > 0 {
			p, err = filepath.EvalSymlinks(p)
			if err != nil {
				return err
			}
		} else {
			break
		}
	}
	go func() {
		for {
			select {
			case <-stopCh:
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				klog.V(2).Infof("Configuration file %s changed, exiting...", event.Name)
				close(stopCh)
				return
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				klog.V(4).Infof("fsnotify error %v", err)
			}
		}
	}()
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

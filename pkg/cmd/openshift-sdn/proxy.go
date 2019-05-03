package openshift_sdn

import (
	"fmt"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	cmdflags "github.com/openshift/origin/pkg/cmd/util/flags"
	sdnproxy "github.com/openshift/origin/pkg/network/proxy"
	"github.com/openshift/origin/pkg/proxy/hybrid"
	"github.com/openshift/origin/pkg/proxy/unidler"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/api/core/v1"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	apiserverflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/client-go/kubernetes/scheme"
	kv1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	kubeproxyoptions "k8s.io/kubernetes/cmd/kube-proxy/app"
	proxy "k8s.io/kubernetes/pkg/proxy"
	kubeproxyconfig "k8s.io/kubernetes/pkg/proxy/apis/config"
	pconfig "k8s.io/kubernetes/pkg/proxy/config"
	"k8s.io/kubernetes/pkg/proxy/healthcheck"
	"k8s.io/kubernetes/pkg/proxy/iptables"
	"k8s.io/kubernetes/pkg/proxy/metrics"
	"k8s.io/kubernetes/pkg/proxy/userspace"
	utildbus "k8s.io/kubernetes/pkg/util/dbus"
	utiliptables "k8s.io/kubernetes/pkg/util/iptables"
	utilnode "k8s.io/kubernetes/pkg/util/node"
	utilsysctl "k8s.io/kubernetes/pkg/util/sysctl"
	utilexec "k8s.io/utils/exec"
	"net"
	"net/http"
	"time"
)

func ProxyConfigFromNodeConfig(options configapi.NodeConfig) (*kubeproxyconfig.KubeProxyConfiguration, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proxyOptions := kubeproxyoptions.NewOptions()
	proxyconfig := proxyOptions.GetConfig()
	defaultedProxyConfig, err := proxyOptions.ApplyDefaults(proxyconfig)
	if err != nil {
		return nil, err
	}
	*proxyconfig = *defaultedProxyConfig
	proxyconfig.HostnameOverride = options.NodeName
	addr := options.ServingInfo.BindAddress
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, fmt.Errorf("The provided value to bind to must be an ip:port %q", addr)
	}
	ip := net.ParseIP(host)
	if ip == nil {
		return nil, fmt.Errorf("The provided value to bind to must be an ip:port: %q", addr)
	}
	proxyconfig.BindAddress = ip.String()
	proxyconfig.MetricsBindAddress = "0.0.0.0:10253"
	if arg := options.ProxyArguments["metrics-bind-address"]; len(arg) > 0 {
		proxyconfig.MetricsBindAddress = arg[0]
	}
	delete(options.ProxyArguments, "metrics-bind-address")
	oomScoreAdj := int32(0)
	proxyconfig.OOMScoreAdj = &oomScoreAdj
	proxyconfig.ResourceContainer = ""
	proxyconfig.ClientConnection.Kubeconfig = options.MasterKubeConfig
	proxyconfig.Mode = "iptables"
	syncPeriod, err := time.ParseDuration(options.IPTablesSyncPeriod)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse the provided ip-tables sync period (%s) : %v", options.IPTablesSyncPeriod, err)
	}
	proxyconfig.IPTables.SyncPeriod = metav1.Duration{Duration: syncPeriod}
	masqueradeBit := int32(0)
	proxyconfig.IPTables.MasqueradeBit = &masqueradeBit
	fss := apiserverflag.NamedFlagSets{}
	proxyOptions.AddFlags(fss.FlagSet("proxy"))
	if err := cmdflags.Resolve(options.ProxyArguments, fss); len(err) > 0 {
		return nil, kerrors.NewAggregate(err)
	}
	if err := proxyOptions.Complete(); err != nil {
		return nil, err
	}
	return proxyconfig, nil
}
func (sdn *OpenShiftSDN) initProxy() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	sdn.OsdnProxy, err = sdnproxy.New(sdn.NodeConfig.NetworkConfig.NetworkPluginName, sdn.informers.NetworkClient, sdn.informers.KubeClient, sdn.informers.NetworkInformers)
	return err
}
func (sdn *OpenShiftSDN) runProxy() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	protocol := utiliptables.ProtocolIpv4
	bindAddr := net.ParseIP(sdn.ProxyConfig.BindAddress)
	if bindAddr.To4() == nil {
		protocol = utiliptables.ProtocolIpv6
	}
	portRange := utilnet.ParsePortRangeOrDie(sdn.ProxyConfig.PortRange)
	hostname, err := utilnode.GetHostname(sdn.ProxyConfig.HostnameOverride)
	if err != nil {
		klog.Fatalf("Unable to get hostname: %v", err)
	}
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(&kv1core.EventSinkImpl{Interface: sdn.informers.KubeClient.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "kube-proxy", Host: hostname})
	execer := utilexec.New()
	dbus := utildbus.New()
	iptInterface := utiliptables.New(execer, dbus, protocol)
	var proxier proxy.ProxyProvider
	var servicesHandler pconfig.ServiceHandler
	var endpointsHandler pconfig.EndpointsHandler
	var healthzServer *healthcheck.HealthzServer
	if len(sdn.ProxyConfig.HealthzBindAddress) > 0 {
		nodeRef := &v1.ObjectReference{Kind: "Node", Name: hostname, UID: types.UID(hostname), Namespace: ""}
		healthzServer = healthcheck.NewDefaultHealthzServer(sdn.ProxyConfig.HealthzBindAddress, 2*sdn.ProxyConfig.IPTables.SyncPeriod.Duration, recorder, nodeRef)
	}
	switch sdn.ProxyConfig.Mode {
	case kubeproxyconfig.ProxyModeIPTables:
		klog.V(0).Info("Using iptables Proxier.")
		if bindAddr.Equal(net.IPv4zero) {
			var err error
			bindAddr, err = getNodeIP(sdn.informers.KubeClient.CoreV1(), hostname)
			if err != nil {
				klog.Fatalf("Unable to get a bind address: %v", err)
			}
		}
		if sdn.ProxyConfig.IPTables.MasqueradeBit == nil {
			klog.Fatalf("Unable to read IPTablesMasqueradeBit from config")
		}
		proxierIptables, err := iptables.NewProxier(iptInterface, utilsysctl.New(), execer, sdn.ProxyConfig.IPTables.SyncPeriod.Duration, sdn.ProxyConfig.IPTables.MinSyncPeriod.Duration, sdn.ProxyConfig.IPTables.MasqueradeAll, int(*sdn.ProxyConfig.IPTables.MasqueradeBit), sdn.ProxyConfig.ClusterCIDR, hostname, bindAddr, recorder, healthzServer, sdn.ProxyConfig.NodePortAddresses)
		metrics.RegisterMetrics()
		if err != nil {
			klog.Fatalf("error: Could not initialize Kubernetes Proxy. You must run this process as root (and if containerized, in the host network namespace as privileged) to use the service proxy: %v", err)
		}
		proxier = proxierIptables
		endpointsHandler = proxierIptables
		servicesHandler = proxierIptables
		klog.V(0).Info("Tearing down userspace rules.")
		userspace.CleanupLeftovers(iptInterface)
	case kubeproxyconfig.ProxyModeUserspace:
		klog.V(0).Info("Using userspace Proxier.")
		loadBalancer := userspace.NewLoadBalancerRR()
		endpointsHandler = loadBalancer
		execer := utilexec.New()
		proxierUserspace, err := userspace.NewProxier(loadBalancer, bindAddr, iptInterface, execer, *portRange, sdn.ProxyConfig.IPTables.SyncPeriod.Duration, sdn.ProxyConfig.IPTables.MinSyncPeriod.Duration, sdn.ProxyConfig.UDPIdleTimeout.Duration, sdn.ProxyConfig.NodePortAddresses)
		if err != nil {
			klog.Fatalf("error: Could not initialize Kubernetes Proxy. You must run this process as root (and if containerized, in the host network namespace as privileged) to use the service proxy: %v", err)
		}
		proxier = proxierUserspace
		servicesHandler = proxierUserspace
		klog.V(0).Info("Tearing down pure-iptables proxy rules.")
		iptables.CleanupLeftovers(iptInterface)
	default:
		klog.Fatalf("Unknown proxy mode %q", sdn.ProxyConfig.Mode)
	}
	serviceConfig := pconfig.NewServiceConfig(sdn.informers.KubeInformers.Core().V1().Services(), sdn.ProxyConfig.ConfigSyncPeriod.Duration)
	if sdn.NodeConfig.EnableUnidling {
		unidlingLoadBalancer := userspace.NewLoadBalancerRR()
		signaler := unidler.NewEventSignaler(recorder)
		unidlingUserspaceProxy, err := unidler.NewUnidlerProxier(unidlingLoadBalancer, bindAddr, iptInterface, execer, *portRange, sdn.ProxyConfig.IPTables.SyncPeriod.Duration, sdn.ProxyConfig.IPTables.MinSyncPeriod.Duration, sdn.ProxyConfig.UDPIdleTimeout.Duration, sdn.ProxyConfig.NodePortAddresses, signaler)
		if err != nil {
			klog.Fatalf("error: Could not initialize Kubernetes Proxy. You must run this process as root (and if containerized, in the host network namespace as privileged) to use the service proxy: %v", err)
		}
		hybridProxier, err := hybrid.NewHybridProxier(unidlingLoadBalancer, unidlingUserspaceProxy, endpointsHandler, servicesHandler, proxier, unidlingUserspaceProxy, sdn.ProxyConfig.IPTables.SyncPeriod.Duration, sdn.informers.KubeInformers.Core().V1().Services().Lister())
		if err != nil {
			klog.Fatalf("error: Could not initialize Kubernetes Proxy. You must run this process as root (and if containerized, in the host network namespace as privileged) to use the service proxy: %v", err)
		}
		endpointsHandler = hybridProxier
		servicesHandler = hybridProxier
		proxier = hybridProxier
	}
	iptInterface.AddReloadFunc(proxier.Sync)
	serviceConfig.RegisterEventHandler(servicesHandler)
	go serviceConfig.Run(utilwait.NeverStop)
	endpointsConfig := pconfig.NewEndpointsConfig(sdn.informers.KubeInformers.Core().V1().Endpoints(), sdn.ProxyConfig.ConfigSyncPeriod.Duration)
	if err := sdn.OsdnProxy.Start(endpointsHandler); err != nil {
		klog.Fatalf("error: node proxy plugin startup failed: %v", err)
	}
	endpointsHandler = sdn.OsdnProxy
	endpointsConfig.RegisterEventHandler(endpointsHandler)
	go endpointsConfig.Run(utilwait.NeverStop)
	if len(sdn.ProxyConfig.HealthzBindAddress) > 0 {
		healthzServer.Run()
	}
	if len(sdn.ProxyConfig.MetricsBindAddress) > 0 {
		mux := http.NewServeMux()
		mux.HandleFunc("/proxyMode", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", sdn.ProxyConfig.Mode)
		})
		mux.Handle("/metrics", prometheus.Handler())
		go utilwait.Until(func() {
			err := http.ListenAndServe(sdn.ProxyConfig.MetricsBindAddress, mux)
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("starting metrics server failed: %v", err))
			}
		}, 5*time.Second, utilwait.NeverStop)
	}
	go utilwait.Forever(proxier.SyncLoop, 0)
	klog.Infof("Started Kubernetes Proxy on %s", sdn.ProxyConfig.BindAddress)
}
func getNodeIP(client kv1core.CoreV1Interface, hostname string) (net.IP, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var node *v1.Node
	var nodeErr error
	nodeWaitBackoff := utilwait.Backoff{Duration: 2 * time.Second, Factor: 2, Steps: 7}
	utilwait.ExponentialBackoff(nodeWaitBackoff, func() (bool, error) {
		node, nodeErr = client.Nodes().Get(hostname, metav1.GetOptions{})
		if nodeErr == nil {
			return true, nil
		} else if kapierrors.IsNotFound(nodeErr) {
			klog.Warningf("waiting for node %q to be registered with master...", hostname)
			return false, nil
		} else {
			return false, nodeErr
		}
	})
	if nodeErr != nil {
		return nil, fmt.Errorf("failed to retrieve node info (after waiting): %v", nodeErr)
	}
	nodeIP, err := utilnode.GetNodeHostIP(node)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve node IP: %v", err)
	}
	return nodeIP, nil
}

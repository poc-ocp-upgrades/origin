package app

import (
	"errors"
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/proxy"
	proxyconfigapi "k8s.io/kubernetes/pkg/proxy/apis/config"
	proxyconfig "k8s.io/kubernetes/pkg/proxy/config"
	"k8s.io/kubernetes/pkg/proxy/healthcheck"
	"k8s.io/kubernetes/pkg/proxy/winkernel"
	"k8s.io/kubernetes/pkg/proxy/winuserspace"
	"k8s.io/kubernetes/pkg/util/configz"
	utilnetsh "k8s.io/kubernetes/pkg/util/netsh"
	utilnode "k8s.io/kubernetes/pkg/util/node"
	"k8s.io/utils/exec"
	"net"
	_ "net/http/pprof"
)

func NewProxyServer(o *Options) (*ProxyServer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newProxyServer(o.config, o.CleanupAndExit, o.scheme, o.master)
}
func newProxyServer(config *proxyconfigapi.KubeProxyConfiguration, cleanupAndExit bool, scheme *runtime.Scheme, master string) (*ProxyServer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if config == nil {
		return nil, errors.New("config is required")
	}
	if c, err := configz.New(proxyconfigapi.GroupName); err == nil {
		c.Set(config)
	} else {
		return nil, fmt.Errorf("unable to register configz: %s", err)
	}
	if cleanupAndExit {
		return &ProxyServer{CleanupAndExit: cleanupAndExit}, nil
	}
	client, eventClient, err := createClients(config.ClientConnection, master)
	if err != nil {
		return nil, err
	}
	hostname, err := utilnode.GetHostname(config.HostnameOverride)
	if err != nil {
		return nil, err
	}
	eventBroadcaster := record.NewBroadcaster()
	recorder := eventBroadcaster.NewRecorder(scheme, v1.EventSource{Component: "kube-proxy", Host: hostname})
	nodeRef := &v1.ObjectReference{Kind: "Node", Name: hostname, UID: types.UID(hostname), Namespace: ""}
	var healthzServer *healthcheck.HealthzServer
	var healthzUpdater healthcheck.HealthzUpdater
	if len(config.HealthzBindAddress) > 0 {
		healthzServer = healthcheck.NewDefaultHealthzServer(config.HealthzBindAddress, 2*config.IPTables.SyncPeriod.Duration, recorder, nodeRef)
		healthzUpdater = healthzServer
	}
	var proxier proxy.ProxyProvider
	var serviceEventHandler proxyconfig.ServiceHandler
	var endpointsEventHandler proxyconfig.EndpointsHandler
	proxyMode := getProxyMode(string(config.Mode), winkernel.WindowsKernelCompatTester{})
	if proxyMode == proxyModeKernelspace {
		klog.V(0).Info("Using Kernelspace Proxier.")
		proxierKernelspace, err := winkernel.NewProxier(config.IPTables.SyncPeriod.Duration, config.IPTables.MinSyncPeriod.Duration, config.IPTables.MasqueradeAll, int(*config.IPTables.MasqueradeBit), config.ClusterCIDR, hostname, utilnode.GetNodeIP(client, hostname), recorder, healthzUpdater)
		if err != nil {
			return nil, fmt.Errorf("unable to create proxier: %v", err)
		}
		proxier = proxierKernelspace
		endpointsEventHandler = proxierKernelspace
		serviceEventHandler = proxierKernelspace
	} else {
		klog.V(0).Info("Using userspace Proxier.")
		execer := exec.New()
		var netshInterface utilnetsh.Interface
		netshInterface = utilnetsh.New(execer)
		loadBalancer := winuserspace.NewLoadBalancerRR()
		endpointsEventHandler = loadBalancer
		proxierUserspace, err := winuserspace.NewProxier(loadBalancer, net.ParseIP(config.BindAddress), netshInterface, *utilnet.ParsePortRangeOrDie(config.PortRange), config.IPTables.SyncPeriod.Duration, config.UDPIdleTimeout.Duration)
		if err != nil {
			return nil, fmt.Errorf("unable to create proxier: %v", err)
		}
		proxier = proxierUserspace
		serviceEventHandler = proxierUserspace
		klog.V(0).Info("Tearing down pure-winkernel proxy rules.")
		winkernel.CleanupLeftovers()
	}
	return &ProxyServer{Client: client, EventClient: eventClient, Proxier: proxier, Broadcaster: eventBroadcaster, Recorder: recorder, ProxyMode: proxyMode, NodeRef: nodeRef, MetricsBindAddress: config.MetricsBindAddress, EnableProfiling: config.EnableProfiling, OOMScoreAdj: config.OOMScoreAdj, ResourceContainer: config.ResourceContainer, ConfigSyncPeriod: config.ConfigSyncPeriod.Duration, ServiceEventHandler: serviceEventHandler, EndpointsEventHandler: endpointsEventHandler, HealthzServer: healthzServer}, nil
}
func getProxyMode(proxyMode string, kcompat winkernel.KernelCompatTester) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if proxyMode == proxyModeUserspace {
		return proxyModeUserspace
	} else if proxyMode == proxyModeKernelspace {
		return tryWinKernelSpaceProxy(kcompat)
	}
	return proxyModeUserspace
}
func tryWinKernelSpaceProxy(kcompat winkernel.KernelCompatTester) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	useWinKerelProxy, err := winkernel.CanUseWinKernelProxier(kcompat)
	if err != nil {
		klog.Errorf("Can't determine whether to use windows kernel proxy, using userspace proxier: %v", err)
		return proxyModeUserspace
	}
	if useWinKerelProxy {
		return proxyModeKernelspace
	}
	klog.V(1).Infof("Can't use winkernel proxy, using userspace proxier")
	return proxyModeUserspace
}

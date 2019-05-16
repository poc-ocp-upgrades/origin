package config

import (
	"fmt"
	apimachineryconfig "k8s.io/apimachinery/pkg/apis/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"
)

type KubeProxyIPTablesConfiguration struct {
	MasqueradeBit *int32
	MasqueradeAll bool
	SyncPeriod    metav1.Duration
	MinSyncPeriod metav1.Duration
}
type KubeProxyIPVSConfiguration struct {
	SyncPeriod    metav1.Duration
	MinSyncPeriod metav1.Duration
	Scheduler     string
	ExcludeCIDRs  []string
}
type KubeProxyConntrackConfiguration struct {
	Max                   *int32
	MaxPerCore            *int32
	Min                   *int32
	TCPEstablishedTimeout *metav1.Duration
	TCPCloseWaitTimeout   *metav1.Duration
}
type KubeProxyConfiguration struct {
	metav1.TypeMeta
	FeatureGates       map[string]bool
	BindAddress        string
	HealthzBindAddress string
	MetricsBindAddress string
	EnableProfiling    bool
	ClusterCIDR        string
	HostnameOverride   string
	ClientConnection   apimachineryconfig.ClientConnectionConfiguration
	IPTables           KubeProxyIPTablesConfiguration
	IPVS               KubeProxyIPVSConfiguration
	OOMScoreAdj        *int32
	Mode               ProxyMode
	PortRange          string
	ResourceContainer  string
	UDPIdleTimeout     metav1.Duration
	Conntrack          KubeProxyConntrackConfiguration
	ConfigSyncPeriod   metav1.Duration
	NodePortAddresses  []string
}
type ProxyMode string

const (
	ProxyModeUserspace   ProxyMode = "userspace"
	ProxyModeIPTables    ProxyMode = "iptables"
	ProxyModeIPVS        ProxyMode = "ipvs"
	ProxyModeKernelspace ProxyMode = "kernelspace"
)

type IPVSSchedulerMethod string

const (
	RoundRobin                                  IPVSSchedulerMethod = "rr"
	WeightedRoundRobin                          IPVSSchedulerMethod = "wrr"
	LeastConnection                             IPVSSchedulerMethod = "lc"
	WeightedLeastConnection                     IPVSSchedulerMethod = "wlc"
	LocalityBasedLeastConnection                IPVSSchedulerMethod = "lblc"
	LocalityBasedLeastConnectionWithReplication IPVSSchedulerMethod = "lblcr"
	SourceHashing                               IPVSSchedulerMethod = "sh"
	DestinationHashing                          IPVSSchedulerMethod = "dh"
	ShortestExpectedDelay                       IPVSSchedulerMethod = "sed"
	NeverQueue                                  IPVSSchedulerMethod = "nq"
)

func (m *ProxyMode) Set(s string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	*m = ProxyMode(s)
	return nil
}
func (m *ProxyMode) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if m != nil {
		return string(*m)
	}
	return ""
}
func (m *ProxyMode) Type() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "ProxyMode"
}

type ConfigurationMap map[string]string

func (m *ConfigurationMap) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pairs := []string{}
	for k, v := range *m {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(pairs)
	return strings.Join(pairs, ",")
}
func (m *ConfigurationMap) Set(value string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, s := range strings.Split(value, ",") {
		if len(s) == 0 {
			continue
		}
		arr := strings.SplitN(s, "=", 2)
		if len(arr) == 2 {
			(*m)[strings.TrimSpace(arr[0])] = strings.TrimSpace(arr[1])
		} else {
			(*m)[strings.TrimSpace(arr[0])] = ""
		}
	}
	return nil
}
func (*ConfigurationMap) Type() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "mapStringString"
}

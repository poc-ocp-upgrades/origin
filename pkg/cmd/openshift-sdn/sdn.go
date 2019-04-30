package openshift_sdn

import (
	"strings"
	kclientv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	kv1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	sdnnode "github.com/openshift/origin/pkg/network/node"
)

func (sdn *OpenShiftSDN) initSDN() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	runtimeEndpoint := sdn.NodeConfig.DockerConfig.DockerShimSocket
	runtime, ok := sdn.NodeConfig.KubeletArguments["container-runtime"]
	if ok && len(runtime) == 1 && runtime[0] == "remote" {
		endpoint, ok := sdn.NodeConfig.KubeletArguments["container-runtime-endpoint"]
		if ok && len(endpoint) == 1 {
			runtimeEndpoint = endpoint[0]
		}
	}
	cniBinDir := "/opt/cni/bin"
	if val, ok := sdn.NodeConfig.KubeletArguments["cni-bin-dir"]; ok && len(val) == 1 {
		cniBinDir = val[0]
	}
	cniConfDir := "/etc/cni/net.d"
	if val, ok := sdn.NodeConfig.KubeletArguments["cni-conf-dir"]; ok && len(val) == 1 {
		cniConfDir = val[0]
	}
	enableHostports := !strings.Contains(runtimeEndpoint, "crio")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(&kv1core.EventSinkImpl{Interface: sdn.informers.KubeClient.CoreV1().Events("")})
	eventRecorder := eventBroadcaster.NewRecorder(scheme.Scheme, kclientv1.EventSource{Component: "openshift-sdn", Host: sdn.NodeConfig.NodeName})
	var err error
	sdn.OsdnNode, err = sdnnode.New(&sdnnode.OsdnNodeConfig{PluginName: sdn.NodeConfig.NetworkConfig.NetworkPluginName, Hostname: sdn.NodeConfig.NodeName, SelfIP: sdn.NodeConfig.NodeIP, RuntimeEndpoint: runtimeEndpoint, CNIBinDir: cniBinDir, CNIConfDir: cniConfDir, MTU: sdn.NodeConfig.NetworkConfig.MTU, NetworkClient: sdn.informers.NetworkClient, KClient: sdn.informers.KubeClient, KubeInformers: sdn.informers.KubeInformers, NetworkInformers: sdn.informers.NetworkInformers, IPTablesSyncPeriod: sdn.ProxyConfig.IPTables.SyncPeriod.Duration, MasqueradeBit: sdn.ProxyConfig.IPTables.MasqueradeBit, ProxyMode: sdn.ProxyConfig.Mode, EnableHostports: enableHostports, Recorder: eventRecorder})
	return err
}
func (sdn *OpenShiftSDN) runSDN() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sdn.OsdnNode.Start()
}

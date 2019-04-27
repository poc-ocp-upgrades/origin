package node

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
	"github.com/containernetworking/cni/pkg/types/current"
	networkv1 "github.com/openshift/api/network/v1"
	"github.com/openshift/origin/pkg/network/common"
	"github.com/openshift/origin/pkg/network/node/cniserver"
	"github.com/openshift/origin/pkg/util/netutils"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	kcontainer "k8s.io/kubernetes/pkg/kubelet/container"
	kubehostport "k8s.io/kubernetes/pkg/kubelet/dockershim/network/hostport"
	kbandwidth "k8s.io/kubernetes/pkg/util/bandwidth"
	utildbus "k8s.io/kubernetes/pkg/util/dbus"
	utiliptables "k8s.io/kubernetes/pkg/util/iptables"
	utilexec "k8s.io/utils/exec"
	"github.com/containernetworking/cni/pkg/invoke"
	cnitypes "github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

const (
	podInterfaceName = "eth0"
)

type podHandler interface {
	setup(req *cniserver.PodRequest) (cnitypes.Result, *runningPod, error)
	update(req *cniserver.PodRequest) (uint32, error)
	teardown(req *cniserver.PodRequest) error
}
type runningPod struct {
	podPortMapping	*kubehostport.PodPortMapping
	vnid		uint32
	ofport		int
}
type podManager struct {
	podHandler	podHandler
	cniServer	*cniserver.CNIServer
	requests	chan (*cniserver.PodRequest)
	runningPods	map[string]*runningPod
	runningPodsLock	sync.Mutex
	kClient		kubernetes.Interface
	policy		osdnPolicy
	mtu		uint32
	cniBinPath	string
	ovs		*ovsController
	enableHostports	bool
	hostportsSynced	bool
	activeHostports	bool
	ipamConfig	[]byte
	hostportSyncer	kubehostport.HostportSyncer
}

func newPodManager(kClient kubernetes.Interface, policy osdnPolicy, mtu uint32, cniBinPath string, ovs *ovsController, enableHostports bool) *podManager {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pm := newDefaultPodManager()
	pm.kClient = kClient
	pm.policy = policy
	pm.mtu = mtu
	pm.cniBinPath = cniBinPath
	pm.podHandler = pm
	pm.ovs = ovs
	pm.enableHostports = enableHostports
	return pm
}
func newDefaultPodManager() *podManager {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &podManager{runningPods: make(map[string]*runningPod), requests: make(chan *cniserver.PodRequest, 20)}
}
func getIPAMConfig(clusterNetworks []common.ClusterNetwork, localSubnet string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nodeNet, err := cnitypes.ParseCIDR(localSubnet)
	if err != nil {
		return nil, fmt.Errorf("error parsing node network '%s': %v", localSubnet, err)
	}
	type hostLocalIPAM struct {
		Type	string			`json:"type"`
		Subnet	cnitypes.IPNet		`json:"subnet"`
		Routes	[]cnitypes.Route	`json:"routes"`
		DataDir	string			`json:"dataDir"`
	}
	type cniNetworkConfig struct {
		CNIVersion	string		`json:"cniVersion"`
		Name		string		`json:"name"`
		Type		string		`json:"type"`
		IPAM		*hostLocalIPAM	`json:"ipam"`
	}
	_, mcnet, _ := net.ParseCIDR("224.0.0.0/4")
	routes := []cnitypes.Route{{Dst: net.IPNet{IP: net.IPv4zero, Mask: net.IPMask(net.IPv4zero)}, GW: netutils.GenerateDefaultGateway(nodeNet)}, {Dst: *mcnet}}
	for _, cn := range clusterNetworks {
		routes = append(routes, cnitypes.Route{Dst: *cn.ClusterCIDR})
	}
	return json.Marshal(&cniNetworkConfig{CNIVersion: "0.3.1", Name: "openshift-sdn", Type: "openshift-sdn", IPAM: &hostLocalIPAM{Type: "host-local", DataDir: hostLocalDataDir, Subnet: cnitypes.IPNet{IP: nodeNet.IP, Mask: nodeNet.Mask}, Routes: routes}})
}
func (m *podManager) Start(rundir string, localSubnetCIDR string, clusterNetworks []common.ClusterNetwork, serviceNetworkCIDR string, clearHostPorts bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.enableHostports {
		iptInterface := utiliptables.New(utilexec.New(), utildbus.New(), utiliptables.ProtocolIpv4)
		m.hostportSyncer = kubehostport.NewHostportSyncer(iptInterface)
		if clearHostPorts {
			_ = m.hostportSyncer.SyncHostports(Tun0, nil)
		}
	}
	var err error
	if m.ipamConfig, err = getIPAMConfig(clusterNetworks, localSubnetCIDR); err != nil {
		return err
	}
	go m.processCNIRequests()
	m.cniServer = cniserver.NewCNIServer(rundir, &cniserver.Config{MTU: m.mtu, ServiceNetworkCIDR: serviceNetworkCIDR})
	return m.cniServer.Start(m.handleCNIRequest)
}
func getPodKey(request *cniserver.PodRequest) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s/%s", request.PodNamespace, request.PodName)
}
func (m *podManager) getPod(request *cniserver.PodRequest) *kubehostport.PodPortMapping {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pod := m.runningPods[getPodKey(request)]; pod != nil {
		return pod.podPortMapping
	}
	return nil
}
func hasHostPorts(pod *kubehostport.PodPortMapping) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, mapping := range pod.PortMappings {
		if mapping.HostPort != 0 {
			return true
		}
	}
	return false
}
func (m *podManager) shouldSyncHostports(newPod *kubehostport.PodPortMapping) []*kubehostport.PodPortMapping {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.hostportSyncer == nil {
		return nil
	}
	newActiveHostports := false
	mappings := make([]*kubehostport.PodPortMapping, 0)
	for _, runningPod := range m.runningPods {
		mappings = append(mappings, runningPod.podPortMapping)
		if !newActiveHostports && hasHostPorts(runningPod.podPortMapping) {
			newActiveHostports = true
		}
	}
	if newPod != nil && hasHostPorts(newPod) {
		newActiveHostports = true
	}
	if !m.hostportsSynced || m.activeHostports || newActiveHostports {
		m.hostportsSynced = true
		m.activeHostports = newActiveHostports
		return mappings
	}
	return nil
}
func (m *podManager) addRequest(request *cniserver.PodRequest) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.requests <- request
}
func (m *podManager) waitRequest(request *cniserver.PodRequest) *cniserver.PodResult {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return <-request.Result
}
func (m *podManager) handleCNIRequest(request *cniserver.PodRequest) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(5).Infof("Dispatching pod network request %v", request)
	m.addRequest(request)
	result := m.waitRequest(request)
	klog.V(5).Infof("Returning pod network request %v, result %s err %v", request, string(result.Response), result.Err)
	return result.Response, result.Err
}
func (m *podManager) updateLocalMulticastRulesWithLock(vnid uint32) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ofports []int
	enabled := m.policy.GetMulticastEnabled(vnid)
	if enabled {
		for _, pod := range m.runningPods {
			if pod.vnid == vnid {
				ofports = append(ofports, pod.ofport)
			}
		}
	}
	if err := m.ovs.UpdateLocalMulticastFlows(vnid, enabled, ofports); err != nil {
		utilruntime.HandleError(fmt.Errorf("Error updating OVS multicast flows for VNID %d: %v", vnid, err))
	}
}
func (m *podManager) UpdateLocalMulticastRules(vnid uint32) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.runningPodsLock.Lock()
	defer m.runningPodsLock.Unlock()
	m.updateLocalMulticastRulesWithLock(vnid)
}
func (m *podManager) processCNIRequests() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for request := range m.requests {
		klog.V(5).Infof("Processing pod network request %v", request)
		result := m.processRequest(request)
		klog.V(5).Infof("Processed pod network request %v, result %s err %v", request, string(result.Response), result.Err)
		request.Result <- result
	}
	panic("stopped processing CNI pod requests!")
}
func (m *podManager) processRequest(request *cniserver.PodRequest) *cniserver.PodResult {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.runningPodsLock.Lock()
	defer m.runningPodsLock.Unlock()
	pk := getPodKey(request)
	result := &cniserver.PodResult{}
	switch request.Command {
	case cniserver.CNI_ADD:
		ipamResult, runningPod, err := m.podHandler.setup(request)
		if ipamResult != nil {
			result.Response, err = json.Marshal(ipamResult)
			if err == nil {
				m.runningPods[pk] = runningPod
				if m.ovs != nil {
					m.updateLocalMulticastRulesWithLock(runningPod.vnid)
				}
			}
		}
		if err != nil {
			PodOperationsErrors.WithLabelValues(PodOperationSetup).Inc()
			result.Err = err
		}
	case cniserver.CNI_UPDATE:
		vnid, err := m.podHandler.update(request)
		if err == nil {
			if runningPod, exists := m.runningPods[pk]; exists {
				runningPod.vnid = vnid
			}
		}
		result.Err = err
	case cniserver.CNI_DEL:
		if runningPod, exists := m.runningPods[pk]; exists {
			delete(m.runningPods, pk)
			if m.ovs != nil {
				m.updateLocalMulticastRulesWithLock(runningPod.vnid)
			}
		}
		result.Err = m.podHandler.teardown(request)
		if result.Err != nil {
			PodOperationsErrors.WithLabelValues(PodOperationTeardown).Inc()
		}
	default:
		result.Err = fmt.Errorf("unhandled CNI request %v", request.Command)
	}
	return result
}
func maybeAddMacvlan(pod *corev1.Pod, netns string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	annotation, ok := pod.Annotations[networkv1.AssignMacvlanAnnotation]
	if !ok || annotation == "false" {
		return nil
	}
	privileged := false
	for _, container := range append(pod.Spec.Containers, pod.Spec.InitContainers...) {
		if container.SecurityContext != nil && container.SecurityContext.Privileged != nil && *container.SecurityContext.Privileged {
			privileged = true
			break
		}
	}
	if !privileged {
		return fmt.Errorf("pod has %q annotation but is not privileged", networkv1.AssignMacvlanAnnotation)
	}
	var iface netlink.Link
	var err error
	if annotation == "true" {
		routes, err := netlink.RouteList(nil, netlink.FAMILY_V4)
		if err != nil {
			return fmt.Errorf("failed to read routes: %v", err)
		}
		for _, r := range routes {
			if r.Dst == nil {
				iface, err = netlink.LinkByIndex(r.LinkIndex)
				if err != nil {
					return fmt.Errorf("failed to get default route interface: %v", err)
				}
			}
		}
		if iface == nil {
			return fmt.Errorf("failed to find default route interface")
		}
	} else {
		iface, err = netlink.LinkByName(annotation)
		if err != nil {
			return fmt.Errorf("pod annotation %q is neither 'true' nor the name of a local network interface", networkv1.AssignMacvlanAnnotation)
		}
	}
	podNs, err := ns.GetNS(netns)
	if err != nil {
		return fmt.Errorf("could not open netns %q: %v", netns, err)
	}
	defer podNs.Close()
	err = netlink.LinkAdd(&netlink.Macvlan{LinkAttrs: netlink.LinkAttrs{MTU: iface.Attrs().MTU, Name: "macvlan0", ParentIndex: iface.Attrs().Index, Namespace: netlink.NsFd(podNs.Fd())}, Mode: netlink.MACVLAN_MODE_PRIVATE})
	if err != nil {
		return fmt.Errorf("failed to create macvlan interface: %v", err)
	}
	return nil
}
func createIPAMArgs(netnsPath, cniBinPath string, action cniserver.CNICommand, id string) *invoke.Args {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &invoke.Args{Command: string(action), ContainerID: id, NetNS: netnsPath, IfName: podInterfaceName, Path: cniBinPath}
}
func (m *podManager) ipamAdd(netnsPath string, id string) (*current.Result, net.IP, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if netnsPath == "" {
		return nil, nil, fmt.Errorf("netns required for CNI_ADD")
	}
	args := createIPAMArgs(netnsPath, m.cniBinPath, cniserver.CNI_ADD, id)
	r, err := invoke.ExecPluginWithResult(m.cniBinPath+"/host-local", m.ipamConfig, args)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to run CNI IPAM ADD: %v", err)
	}
	result, err := current.GetResult(r)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse CNI IPAM ADD result: %v", err)
	}
	if len(result.IPs) == 0 {
		return nil, nil, fmt.Errorf("failed to obtain IP address from CNI IPAM")
	}
	return result, result.IPs[0].Address.IP, nil
}
func (m *podManager) ipamDel(id string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := createIPAMArgs("", m.cniBinPath, cniserver.CNI_DEL, id)
	err := invoke.ExecPluginWithoutResult(m.cniBinPath+"/host-local", m.ipamConfig, args)
	if err != nil {
		return fmt.Errorf("failed to run CNI IPAM DEL: %v", err)
	}
	return nil
}
func setupPodBandwidth(ovs *ovsController, pod *corev1.Pod, hostVeth, sandboxID string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ingressVal, egressVal, err := kbandwidth.ExtractPodBandwidthResources(pod.Annotations)
	if err != nil {
		return fmt.Errorf("failed to parse pod bandwidth: %v", err)
	}
	ingressBPS := int64(-1)
	egressBPS := int64(-1)
	if ingressVal != nil {
		ingressBPS = ingressVal.Value()
		l, err := netlink.LinkByName(hostVeth)
		if err != nil {
			return fmt.Errorf("failed to find host veth interface %s: %v", hostVeth, err)
		}
		err = netlink.LinkSetTxQLen(l, 1000)
		if err != nil {
			return fmt.Errorf("failed to set host veth txqlen: %v", err)
		}
	}
	if egressVal != nil {
		egressBPS = egressVal.Value()
	}
	return ovs.SetPodBandwidth(hostVeth, sandboxID, ingressBPS, egressBPS)
}
func vnidToString(vnid uint32) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strconv.FormatUint(uint64(vnid), 10)
}
func podIsExited(p *kcontainer.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, c := range p.Containers {
		if c.State != kcontainer.ContainerStateExited {
			return false
		}
	}
	return true
}
func (m *podManager) setup(req *cniserver.PodRequest) (cnitypes.Result, *runningPod, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer PodOperationsLatency.WithLabelValues(PodOperationSetup).Observe(sinceInMicroseconds(time.Now()))
	var success bool
	defer func() {
		if !success {
			m.ipamDel(req.SandboxID)
			if mappings := m.shouldSyncHostports(nil); mappings != nil {
				if err := m.hostportSyncer.SyncHostports(Tun0, mappings); err != nil {
					klog.Warningf("failed syncing hostports: %v", err)
				}
			}
		}
	}()
	v1Pod, err := m.kClient.CoreV1().Pods(req.PodNamespace).Get(req.PodName, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}
	var ipamResult cnitypes.Result
	podIP := net.ParseIP(req.AssignedIP)
	if podIP == nil {
		ipamResult, podIP, err = m.ipamAdd(req.Netns, req.SandboxID)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to run IPAM for %v: %v", req.SandboxID, err)
		}
		if err := maybeAddMacvlan(v1Pod, req.Netns); err != nil {
			return nil, nil, err
		}
	}
	podPortMapping := constructPodPortMapping(v1Pod, podIP)
	if mappings := m.shouldSyncHostports(podPortMapping); mappings != nil {
		if err := m.hostportSyncer.OpenPodHostportsAndSync(podPortMapping, Tun0, mappings); err != nil {
			return nil, nil, err
		}
	}
	vnid, err := m.policy.GetVNID(req.PodNamespace)
	if err != nil {
		return nil, nil, err
	}
	ofport, err := m.ovs.SetUpPod(req.SandboxID, req.HostVeth, podIP, vnid)
	if err != nil {
		return nil, nil, err
	}
	if err := setupPodBandwidth(m.ovs, v1Pod, req.HostVeth, req.SandboxID); err != nil {
		return nil, nil, err
	}
	m.policy.EnsureVNIDRules(vnid)
	success = true
	return ipamResult, &runningPod{podPortMapping: podPortMapping, vnid: vnid, ofport: ofport}, nil
}
func (m *podManager) update(req *cniserver.PodRequest) (uint32, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vnid, err := m.policy.GetVNID(req.PodNamespace)
	if err != nil {
		return 0, err
	}
	if err := m.ovs.UpdatePod(req.SandboxID, vnid); err != nil {
		return 0, err
	}
	return vnid, nil
}
func (m *podManager) teardown(req *cniserver.PodRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer PodOperationsLatency.WithLabelValues(PodOperationTeardown).Observe(sinceInMicroseconds(time.Now()))
	errList := []error{}
	if err := m.ovs.TearDownPod(req.SandboxID); err != nil {
		errList = append(errList, err)
	}
	if err := m.ipamDel(req.SandboxID); err != nil {
		errList = append(errList, err)
	}
	if mappings := m.shouldSyncHostports(nil); mappings != nil {
		if err := m.hostportSyncer.SyncHostports(Tun0, mappings); err != nil {
			errList = append(errList, err)
		}
	}
	return kerrors.NewAggregate(errList)
}
func constructPodPortMapping(pod *corev1.Pod, podIP net.IP) *kubehostport.PodPortMapping {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	portMappings := make([]*kubehostport.PortMapping, 0)
	for _, c := range pod.Spec.Containers {
		for _, port := range c.Ports {
			portMappings = append(portMappings, &kubehostport.PortMapping{Name: port.Name, HostPort: port.HostPort, ContainerPort: port.ContainerPort, Protocol: port.Protocol, HostIP: port.HostIP})
		}
	}
	return &kubehostport.PodPortMapping{Namespace: pod.Namespace, Name: pod.Name, PortMappings: portMappings, HostNetwork: pod.Spec.HostNetwork, IP: podIP}
}

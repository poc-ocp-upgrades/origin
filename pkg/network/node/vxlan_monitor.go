package node

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"k8s.io/klog"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	"github.com/openshift/origin/pkg/network/common"
	"github.com/openshift/origin/pkg/util/ovs"
)

type egressVXLANMonitor struct {
	sync.Mutex
	ovsif		ovs.Interface
	tracker		*common.EgressIPTracker
	updates		chan<- *egressVXLANNode
	pollInterval	time.Duration
	monitorNodes	map[string]*egressVXLANNode
	stop		chan struct{}
}
type egressVXLANNode struct {
	nodeIP	string
	offline	bool
	in	uint64
	out	uint64
	retries	int
}

const (
	defaultPollInterval	= 5 * time.Second
	repollInterval		= time.Second
	maxRetries		= 2
)

func newEgressVXLANMonitor(ovsif ovs.Interface, tracker *common.EgressIPTracker, updates chan<- *egressVXLANNode) *egressVXLANMonitor {
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
	return &egressVXLANMonitor{ovsif: ovsif, tracker: tracker, updates: updates, pollInterval: defaultPollInterval, monitorNodes: make(map[string]*egressVXLANNode)}
}
func (evm *egressVXLANMonitor) AddNode(nodeIP string) {
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
	evm.Lock()
	defer evm.Unlock()
	if evm.monitorNodes[nodeIP] != nil {
		return
	}
	klog.V(4).Infof("Monitoring node %s", nodeIP)
	evm.monitorNodes[nodeIP] = &egressVXLANNode{nodeIP: nodeIP}
	if len(evm.monitorNodes) == 1 && evm.pollInterval != 0 {
		evm.stop = make(chan struct{})
		go utilwait.PollUntil(evm.pollInterval, evm.poll, evm.stop)
	}
}
func (evm *egressVXLANMonitor) RemoveNode(nodeIP string) {
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
	evm.Lock()
	defer evm.Unlock()
	if evm.monitorNodes[nodeIP] == nil {
		return
	}
	klog.V(4).Infof("Unmonitoring node %s", nodeIP)
	delete(evm.monitorNodes, nodeIP)
	if len(evm.monitorNodes) == 0 && evm.stop != nil {
		close(evm.stop)
		evm.stop = nil
	}
}
func parseNPackets(of *ovs.OvsFlow) (uint64, error) {
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
	str, _ := of.FindField("n_packets")
	if str == nil {
		return 0, fmt.Errorf("no packet count")
	}
	nPackets, err := strconv.ParseUint(str.Value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("bad packet count: %v", err)
	}
	return nPackets, nil
}
func (evm *egressVXLANMonitor) check(retryOnly bool) bool {
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
	inFlows, err := evm.ovsif.DumpFlows("table=10")
	if err != nil {
		utilruntime.HandleError(err)
		return false
	}
	outFlows, err := evm.ovsif.DumpFlows("table=100")
	if err != nil {
		utilruntime.HandleError(err)
		return false
	}
	inTraffic := make(map[string]uint64)
	for _, flow := range inFlows {
		parsed, err := ovs.ParseFlow(ovs.ParseForDump, flow)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Error parsing VXLAN input flow: %v", err))
			continue
		}
		tunSrc, _ := parsed.FindField("tun_src")
		if tunSrc == nil {
			continue
		}
		if evm.monitorNodes[tunSrc.Value] == nil {
			continue
		}
		nPackets, err := parseNPackets(parsed)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Could not parse %q: %v", flow, err))
			continue
		}
		inTraffic[tunSrc.Value] = nPackets
	}
	outTraffic := make(map[string]uint64)
	for _, flow := range outFlows {
		parsed, err := ovs.ParseFlow(ovs.ParseForDump, flow)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Error parsing VXLAN output flow: %v", err))
			continue
		}
		tunDst := ""
		for _, act := range parsed.Actions {
			if act.Name == "set_field" && strings.HasSuffix(act.Value, "->tun_dst") {
				tunDst = strings.TrimSuffix(act.Value, "->tun_dst")
				break
			}
		}
		if tunDst == "" {
			continue
		}
		if evm.monitorNodes[tunDst] == nil {
			continue
		}
		nPackets, err := parseNPackets(parsed)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Could not parse %q: %v", flow, err))
			continue
		}
		outTraffic[tunDst] += nPackets
	}
	retry := false
	for _, node := range evm.monitorNodes {
		if retryOnly && node.retries == 0 {
			continue
		}
		in := inTraffic[node.nodeIP]
		out := outTraffic[node.nodeIP]
		if node.offline {
			if in > node.in {
				klog.Infof("Node %s is back online", node.nodeIP)
				node.offline = false
				evm.updates <- node
			} else if evm.tracker != nil {
				go evm.tracker.Ping(node.nodeIP, defaultPollInterval)
			}
		} else {
			if out > node.out && in == node.in {
				node.retries++
				if node.retries > maxRetries {
					klog.Warningf("Node %s is offline", node.nodeIP)
					node.retries = 0
					node.offline = true
					evm.updates <- node
				} else {
					klog.V(2).Infof("Node %s may be offline... retrying", node.nodeIP)
					retry = true
					continue
				}
			}
		}
		node.in = in
		node.out = out
	}
	return retry
}
func (evm *egressVXLANMonitor) poll() (bool, error) {
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
	evm.Lock()
	defer evm.Unlock()
	retry := evm.check(false)
	for retry {
		time.Sleep(repollInterval)
		retry = evm.check(true)
	}
	return false, nil
}

package node

import (
	"fmt"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	utildbus "k8s.io/kubernetes/pkg/util/dbus"
	"k8s.io/kubernetes/pkg/util/iptables"
	kexec "k8s.io/utils/exec"
	"sync"
	"time"
)

type NodeIPTables struct {
	ipt                iptables.Interface
	clusterNetworkCIDR []string
	syncPeriod         time.Duration
	masqueradeServices bool
	vxlanPort          uint32
	mu                 sync.Mutex
	egressIPs          map[string]string
}

func newNodeIPTables(clusterNetworkCIDR []string, syncPeriod time.Duration, masqueradeServices bool, vxlanPort uint32) *NodeIPTables {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &NodeIPTables{ipt: iptables.New(kexec.New(), utildbus.New(), iptables.ProtocolIpv4), clusterNetworkCIDR: clusterNetworkCIDR, syncPeriod: syncPeriod, masqueradeServices: masqueradeServices, vxlanPort: vxlanPort, egressIPs: make(map[string]string)}
}
func (n *NodeIPTables) Setup() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := n.syncIPTableRules(); err != nil {
		return err
	}
	n.ipt.AddReloadFunc(func() {
		if err := n.syncIPTableRules(); err != nil {
			utilruntime.HandleError(fmt.Errorf("Reloading openshift iptables failed: %v", err))
		}
	})
	go utilwait.Forever(n.syncLoop, 0)
	return nil
}
func (n *NodeIPTables) syncLoop() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	t := time.NewTicker(n.syncPeriod)
	defer t.Stop()
	for {
		<-t.C
		klog.V(6).Infof("Periodic openshift iptables sync")
		err := n.syncIPTableRules()
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Syncing openshift iptables failed: %v", err))
		}
	}
}

type Chain struct {
	table    string
	name     string
	srcChain string
	srcRule  []string
	rules    [][]string
}

func (n *NodeIPTables) addChainRules(chain Chain) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allExisted := true
	for _, rule := range chain.rules {
		existed, err := n.ipt.EnsureRule(iptables.Append, iptables.Table(chain.table), iptables.Chain(chain.name), rule...)
		if err != nil {
			return false, fmt.Errorf("failed to ensure rule %v exists: %v", rule, err)
		}
		if !existed {
			allExisted = false
		}
	}
	return allExisted, nil
}
func (n *NodeIPTables) syncIPTableRules() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.mu.Lock()
	defer n.mu.Unlock()
	start := time.Now()
	defer func() {
		klog.V(4).Infof("syncIPTableRules took %v", time.Since(start))
	}()
	klog.V(3).Infof("Syncing openshift iptables rules")
	chains := n.getNodeIPTablesChains()
	for i := len(chains) - 1; i >= 0; i-- {
		chain := chains[i]
		chainExisted, err := n.ipt.EnsureChain(iptables.Table(chain.table), iptables.Chain(chain.name))
		if err != nil {
			return fmt.Errorf("failed to ensure chain %s exists: %v", chain.name, err)
		}
		if chain.srcChain != "" {
			_, err = n.ipt.EnsureRule(iptables.Prepend, iptables.Table(chain.table), iptables.Chain(chain.srcChain), append(chain.srcRule, "-j", chain.name)...)
			if err != nil {
				return fmt.Errorf("failed to ensure rule from %s to %s exists: %v", chain.srcChain, chain.name, err)
			}
		}
		rulesExisted, err := n.addChainRules(chain)
		if err != nil {
			return err
		}
		if chainExisted && !rulesExisted {
			if err = n.ipt.FlushChain(iptables.Table(chain.table), iptables.Chain(chain.name)); err != nil {
				return fmt.Errorf("failed to flush chain %s: %v", chain.name, err)
			}
			if _, err = n.addChainRules(chain); err != nil {
				return err
			}
		}
	}
	for egressIP, mark := range n.egressIPs {
		if err := n.ensureEgressIPRules(egressIP, mark); err != nil {
			return err
		}
	}
	return nil
}
func (n *NodeIPTables) getNodeIPTablesChains() []Chain {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var chainArray []Chain
	chainArray = append(chainArray, Chain{table: "filter", name: "OPENSHIFT-FIREWALL-ALLOW", srcChain: "INPUT", srcRule: []string{"-m", "comment", "--comment", "firewall overrides"}, rules: [][]string{{"-p", "udp", "--dport", fmt.Sprintf("%d", n.vxlanPort), "-m", "comment", "--comment", "VXLAN incoming", "-j", "ACCEPT"}, {"-i", Tun0, "-m", "comment", "--comment", "from SDN to localhost", "-j", "ACCEPT"}, {"-i", "docker0", "-m", "comment", "--comment", "from docker to localhost", "-j", "ACCEPT"}}}, Chain{table: "filter", name: "OPENSHIFT-ADMIN-OUTPUT-RULES", srcChain: "FORWARD", srcRule: []string{"-i", Tun0, "!", "-o", Tun0, "-m", "comment", "--comment", "administrator overrides"}, rules: nil})
	var masqRules [][]string
	var masq2Rules [][]string
	var filterRules [][]string
	for _, cidr := range n.clusterNetworkCIDR {
		if n.masqueradeServices {
			masqRules = append(masqRules, []string{"-s", cidr, "-m", "comment", "--comment", "masquerade pod-to-service and pod-to-external traffic", "-j", "MASQUERADE"})
		} else {
			masqRules = append(masqRules, []string{"-s", cidr, "-m", "comment", "--comment", "masquerade pod-to-external traffic", "-j", "OPENSHIFT-MASQUERADE-2"})
			masq2Rules = append(masq2Rules, []string{"-d", cidr, "-m", "comment", "--comment", "masquerade pod-to-external traffic", "-j", "RETURN"})
		}
		filterRules = append(filterRules, []string{"-s", cidr, "-m", "comment", "--comment", "attempted resend after connection close", "-m", "conntrack", "--ctstate", "INVALID", "-j", "DROP"})
		filterRules = append(filterRules, []string{"-d", cidr, "-m", "comment", "--comment", "forward traffic from SDN", "-j", "ACCEPT"})
		filterRules = append(filterRules, []string{"-s", cidr, "-m", "comment", "--comment", "forward traffic to SDN", "-j", "ACCEPT"})
	}
	chainArray = append(chainArray, Chain{table: "nat", name: "OPENSHIFT-MASQUERADE", srcChain: "POSTROUTING", srcRule: []string{"-m", "comment", "--comment", "rules for masquerading OpenShift traffic"}, rules: masqRules}, Chain{table: "filter", name: "OPENSHIFT-FIREWALL-FORWARD", srcChain: "FORWARD", srcRule: []string{"-m", "comment", "--comment", "firewall overrides"}, rules: filterRules})
	if !n.masqueradeServices {
		masq2Rules = append(masq2Rules, []string{"-j", "MASQUERADE"})
		chainArray = append(chainArray, Chain{table: "nat", name: "OPENSHIFT-MASQUERADE-2", rules: masq2Rules})
	}
	return chainArray
}
func (n *NodeIPTables) ensureEgressIPRules(egressIP, mark string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, cidr := range n.clusterNetworkCIDR {
		_, err := n.ipt.EnsureRule(iptables.Prepend, iptables.TableNAT, iptables.Chain("OPENSHIFT-MASQUERADE"), "-s", cidr, "-m", "mark", "--mark", mark, "-j", "SNAT", "--to-source", egressIP)
		if err != nil {
			return err
		}
	}
	_, err := n.ipt.EnsureRule(iptables.Append, iptables.TableFilter, iptables.Chain("OPENSHIFT-FIREWALL-ALLOW"), "-d", egressIP, "-m", "conntrack", "--ctstate", "NEW", "-j", "REJECT")
	return err
}
func (n *NodeIPTables) AddEgressIPRules(egressIP, mark string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.mu.Lock()
	defer n.mu.Unlock()
	if err := n.ensureEgressIPRules(egressIP, mark); err != nil {
		return err
	}
	n.egressIPs[egressIP] = mark
	return nil
}
func (n *NodeIPTables) DeleteEgressIPRules(egressIP, mark string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.mu.Lock()
	defer n.mu.Unlock()
	delete(n.egressIPs, egressIP)
	for _, cidr := range n.clusterNetworkCIDR {
		err := n.ipt.DeleteRule(iptables.TableNAT, iptables.Chain("OPENSHIFT-MASQUERADE"), "-s", cidr, "-m", "mark", "--mark", mark, "-j", "SNAT", "--to-source", egressIP)
		if err != nil {
			return err
		}
	}
	return n.ipt.DeleteRule(iptables.TableFilter, iptables.Chain("OPENSHIFT-FIREWALL-ALLOW"), "-d", egressIP, "-m", "conntrack", "--ctstate", "NEW", "-j", "REJECT")
}

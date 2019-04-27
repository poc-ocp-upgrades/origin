package node

import (
	"crypto/sha256"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"k8s.io/klog"
	networkapi "github.com/openshift/api/network/v1"
	"github.com/openshift/origin/pkg/network/common"
	"github.com/openshift/origin/pkg/util/ovs"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
)

type ovsController struct {
	ovs		ovs.Interface
	pluginId	int
	useConnTrack	bool
	localIP		string
	tunMAC		string
}

const (
	Br0			= "br0"
	Tun0			= "tun0"
	Vxlan0			= "vxlan0"
	ruleVersion		= 7
	ruleVersionTable	= 253
)

func NewOVSController(ovsif ovs.Interface, pluginId int, useConnTrack bool, localIP string) *ovsController {
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
	return &ovsController{ovs: ovsif, pluginId: pluginId, useConnTrack: useConnTrack, localIP: localIP}
}
func (oc *ovsController) getVersionNote() string {
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
	if ruleVersion > 254 {
		panic("Version too large!")
	}
	return fmt.Sprintf("%02X.%02X", oc.pluginId, ruleVersion)
}
func (oc *ovsController) AlreadySetUp(vxlanPort uint32) bool {
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
	flows, err := oc.ovs.DumpFlows("table=%d", ruleVersionTable)
	if err != nil || len(flows) != 1 {
		return false
	}
	port, err := oc.ovs.Get("Interface", Vxlan0, "options:dst_port")
	if err != nil || fmt.Sprintf("\"%d\"", vxlanPort) != port {
		return false
	}
	if parsed, err := ovs.ParseFlow(ovs.ParseForDump, flows[0]); err == nil {
		return parsed.NoteHasPrefix(oc.getVersionNote())
	}
	return false
}
func (oc *ovsController) SetupOVS(clusterNetworkCIDR []string, serviceNetworkCIDR, localSubnetCIDR, localSubnetGateway string, mtu uint32, vxlanPort uint32) error {
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
	err := oc.ovs.DeleteBridge(true)
	if err != nil {
		return err
	}
	err = oc.ovs.AddBridge("fail_mode=secure", "protocols=OpenFlow13")
	if err != nil {
		return err
	}
	err = oc.ovs.SetFrags("nx-match")
	if err != nil {
		return err
	}
	_ = oc.ovs.DeletePort(Vxlan0)
	_, err = oc.ovs.AddPort(Vxlan0, 1, "type=vxlan", `options:remote_ip="flow"`, `options:key="flow"`, fmt.Sprintf("options:dst_port=%d", vxlanPort))
	if err != nil {
		return err
	}
	_ = oc.ovs.DeletePort(Tun0)
	_, err = oc.ovs.AddPort(Tun0, 2, "type=internal", fmt.Sprintf("mtu_request=%d", mtu))
	if err != nil {
		return err
	}
	otx := oc.ovs.NewTransaction()
	if oc.useConnTrack {
		otx.AddFlow("table=0, priority=300, ip, ct_state=-trk, actions=ct(table=0)")
	}
	for _, clusterCIDR := range clusterNetworkCIDR {
		otx.AddFlow("table=0, priority=200, in_port=1, arp, nw_src=%s, nw_dst=%s, actions=move:NXM_NX_TUN_ID[0..31]->NXM_NX_REG0[],goto_table:10", clusterCIDR, localSubnetCIDR)
		otx.AddFlow("table=0, priority=200, in_port=1, ip, nw_src=%s, actions=move:NXM_NX_TUN_ID[0..31]->NXM_NX_REG0[],goto_table:10", clusterCIDR)
		otx.AddFlow("table=0, priority=200, in_port=1, ip, nw_dst=%s, actions=move:NXM_NX_TUN_ID[0..31]->NXM_NX_REG0[],goto_table:10", clusterCIDR)
	}
	otx.AddFlow("table=0, priority=150, in_port=1, actions=drop")
	if oc.useConnTrack {
		otx.AddFlow("table=0, priority=400, in_port=2, ip, nw_src=%s, actions=goto_table:30", localSubnetGateway)
		for _, clusterCIDR := range clusterNetworkCIDR {
			otx.AddFlow("table=0, priority=300, in_port=2, ip, nw_src=%s, nw_dst=%s, actions=goto_table:25", localSubnetCIDR, clusterCIDR)
		}
	}
	otx.AddFlow("table=0, priority=250, in_port=2, ip, nw_dst=224.0.0.0/4, actions=drop")
	for _, clusterCIDR := range clusterNetworkCIDR {
		otx.AddFlow("table=0, priority=200, in_port=2, arp, nw_src=%s, nw_dst=%s, actions=goto_table:30", localSubnetGateway, clusterCIDR)
	}
	otx.AddFlow("table=0, priority=200, in_port=2, ip, actions=goto_table:30")
	otx.AddFlow("table=0, priority=150, in_port=2, actions=drop")
	otx.AddFlow("table=0, priority=100, arp, actions=goto_table:20")
	otx.AddFlow("table=0, priority=100, ip, actions=goto_table:20")
	otx.AddFlow("table=0, priority=0, actions=drop")
	otx.AddFlow("table=10, priority=0, actions=drop")
	otx.AddFlow("table=20, priority=0, actions=drop")
	otx.AddFlow("table=21, priority=0, actions=goto_table:30")
	if oc.useConnTrack {
		otx.AddFlow("table=25, priority=0, actions=drop")
	}
	otx.AddFlow("table=30, priority=300, arp, nw_dst=%s, actions=output:2", localSubnetGateway)
	otx.AddFlow("table=30, priority=200, arp, nw_dst=%s, actions=goto_table:40", localSubnetCIDR)
	for _, clusterCIDR := range clusterNetworkCIDR {
		otx.AddFlow("table=30, priority=100, arp, nw_dst=%s, actions=goto_table:50", clusterCIDR)
	}
	otx.AddFlow("table=30, priority=300, ip, nw_dst=%s, actions=output:2", localSubnetGateway)
	otx.AddFlow("table=30, priority=100, ip, nw_dst=%s, actions=goto_table:60", serviceNetworkCIDR)
	if oc.useConnTrack {
		otx.AddFlow("table=30, priority=300, ip, nw_dst=%s, ct_state=+rpl, actions=ct(nat,table=70)", localSubnetCIDR)
	}
	otx.AddFlow("table=30, priority=200, ip, nw_dst=%s, actions=goto_table:70", localSubnetCIDR)
	for _, clusterCIDR := range clusterNetworkCIDR {
		otx.AddFlow("table=30, priority=100, ip, nw_dst=%s, actions=goto_table:90", clusterCIDR)
	}
	otx.AddFlow("table=30, priority=50, in_port=1, ip, nw_dst=224.0.0.0/4, actions=goto_table:120")
	otx.AddFlow("table=30, priority=25, ip, nw_dst=224.0.0.0/4, actions=goto_table:110")
	otx.AddFlow("table=30, priority=0, ip, actions=goto_table:100")
	otx.AddFlow("table=30, priority=0, arp, actions=drop")
	otx.AddFlow("table=40, priority=0, actions=drop")
	otx.AddFlow("table=50, priority=0, actions=drop")
	if oc.useConnTrack {
		otx.AddFlow("table=60, priority=200, actions=output:2")
	} else {
		otx.AddFlow("table=60, priority=200, reg0=0, actions=output:2")
	}
	otx.AddFlow("table=60, priority=0, actions=drop")
	otx.AddFlow("table=70, priority=0, actions=drop")
	otx.AddFlow("table=80, priority=300, ip, nw_src=%s/32, actions=output:NXM_NX_REG2[]", localSubnetGateway)
	otx.AddFlow("table=80, priority=0, actions=drop")
	otx.AddFlow("table=90, priority=0, actions=drop")
	otx.AddFlow("table=100, priority=300,udp,udp_dst=%d,actions=drop", vxlanPort)
	otx.AddFlow("table=100, priority=200,tcp,tcp_dst=53,nw_dst=%s,actions=output:2", oc.localIP)
	otx.AddFlow("table=100, priority=200,udp,udp_dst=53,nw_dst=%s,actions=output:2", oc.localIP)
	otx.AddFlow("table=100, priority=0, actions=goto_table:101")
	otx.AddFlow("table=101, priority=0, actions=output:2")
	otx.AddFlow("table=110, priority=0, actions=drop")
	otx.AddFlow("table=111, priority=100, actions=goto_table:120")
	otx.AddFlow("table=120, priority=0, actions=drop")
	otx.AddFlow("table=%d, actions=note:%s", ruleVersionTable, oc.getVersionNote())
	return otx.Commit()
}

type podNetworkInfo struct {
	vethName	string
	ip		string
}

func (oc *ovsController) GetPodNetworkInfo() (map[string]podNetworkInfo, error) {
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
	rows, err := oc.ovs.Find("interface", []string{"name", "external_ids"}, "external_ids:sandbox!=\"\"")
	if err != nil {
		return nil, err
	}
	results := make(map[string]podNetworkInfo)
	for _, row := range rows {
		if row["name"] == "" || row["external_ids"] == "" {
			utilruntime.HandleError(fmt.Errorf("ovs-vsctl output missing one or more fields: %v", row))
			continue
		}
		ids, err := ovs.ParseExternalIDs(row["external_ids"])
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Could not parse external_ids %q: %v", row["external_ids"], err))
			continue
		}
		if ids["ip"] == "" || ids["sandbox"] == "" {
			utilruntime.HandleError(fmt.Errorf("ovs-vsctl output missing one or more external_ids: %v", ids))
			continue
		}
		if net.ParseIP(ids["ip"]) == nil {
			utilruntime.HandleError(fmt.Errorf("Could not parse IP %q for sandbox %q", ids["ip"], ids["sandbox"]))
			continue
		}
		results[ids["sandbox"]] = podNetworkInfo{vethName: row["name"], ip: ids["ip"]}
	}
	return results, nil
}
func (oc *ovsController) NewTransaction() ovs.Transaction {
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
	return oc.ovs.NewTransaction()
}
func (oc *ovsController) ensureOvsPort(hostVeth, sandboxID, podIP string) (int, error) {
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
	ofport, err := oc.ovs.AddPort(hostVeth, -1, fmt.Sprintf(`external_ids=sandbox="%s",ip="%s"`, sandboxID, podIP))
	if err != nil {
		_ = oc.ovs.DeletePort(hostVeth)
	}
	return ofport, err
}
func (oc *ovsController) setupPodFlows(ofport int, podIP net.IP, vnid uint32) error {
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
	otx := oc.ovs.NewTransaction()
	ipstr := podIP.String()
	podIP = podIP.To4()
	ipmac := fmt.Sprintf("00:00:%02x:%02x:%02x:%02x/00:00:ff:ff:ff:ff", podIP[0], podIP[1], podIP[2], podIP[3])
	otx.AddFlow("table=20, priority=100, in_port=%d, arp, nw_src=%s, arp_sha=%s, actions=load:%d->NXM_NX_REG0[], goto_table:21", ofport, ipstr, ipmac, vnid)
	otx.AddFlow("table=20, priority=100, in_port=%d, ip, nw_src=%s, actions=load:%d->NXM_NX_REG0[], goto_table:21", ofport, ipstr, vnid)
	if oc.useConnTrack {
		otx.AddFlow("table=25, priority=100, ip, nw_src=%s, actions=load:%d->NXM_NX_REG0[], goto_table:30", ipstr, vnid)
	}
	otx.AddFlow("table=40, priority=100, arp, nw_dst=%s, actions=output:%d", ipstr, ofport)
	otx.AddFlow("table=70, priority=100, ip, nw_dst=%s, actions=load:%d->NXM_NX_REG1[], load:%d->NXM_NX_REG2[], goto_table:80", ipstr, vnid, ofport)
	return otx.Commit()
}
func (oc *ovsController) cleanupPodFlows(podIP net.IP) error {
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
	ipstr := podIP.String()
	otx := oc.ovs.NewTransaction()
	otx.DeleteFlows("ip, nw_dst=%s", ipstr)
	otx.DeleteFlows("ip, nw_src=%s", ipstr)
	otx.DeleteFlows("arp, nw_dst=%s", ipstr)
	otx.DeleteFlows("arp, nw_src=%s", ipstr)
	return otx.Commit()
}
func (oc *ovsController) SetUpPod(sandboxID, hostVeth string, podIP net.IP, vnid uint32) (int, error) {
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
	ofport, err := oc.ensureOvsPort(hostVeth, sandboxID, podIP.String())
	if err != nil {
		return -1, err
	}
	return ofport, oc.setupPodFlows(ofport, podIP, vnid)
}
func (oc *ovsController) getInterfacesForSandbox(sandboxID string) ([]string, error) {
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
	return oc.ovs.FindOne("interface", "name", "external_ids:sandbox="+sandboxID)
}
func (oc *ovsController) ClearPodBandwidth(portList []string, sandboxID string) error {
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
	for _, port := range portList {
		if err := oc.ovs.Clear("port", port, "qos"); err != nil {
			return err
		}
	}
	qosList, err := oc.ovs.FindOne("qos", "_uuid", "external_ids:sandbox="+sandboxID)
	if err != nil {
		return err
	}
	for _, qos := range qosList {
		if err := oc.ovs.Destroy("qos", qos); err != nil {
			return err
		}
	}
	return nil
}
func (oc *ovsController) SetPodBandwidth(hostVeth, sandboxID string, ingressBPS, egressBPS int64) error {
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
	ports, err := oc.getInterfacesForSandbox(sandboxID)
	if err != nil {
		return err
	}
	if err := oc.ClearPodBandwidth(ports, sandboxID); err != nil {
		return err
	}
	if ingressBPS > 0 {
		qos, err := oc.ovs.Create("qos", "type=linux-htb", fmt.Sprintf("other_config:max-rate=%d", ingressBPS), "external_ids=sandbox="+sandboxID)
		if err != nil {
			return err
		}
		err = oc.ovs.Set("port", hostVeth, fmt.Sprintf("qos=%s", qos))
		if err != nil {
			return err
		}
	}
	if egressBPS > 0 {
		err := oc.ovs.Set("interface", hostVeth, fmt.Sprintf("ingress_policing_rate=%d", egressBPS/1024))
		if err != nil {
			return err
		}
	}
	return nil
}
func (oc *ovsController) getPodDetailsBySandboxID(sandboxID string) (int, net.IP, error) {
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
	rows, err := oc.ovs.Find("interface", []string{"ofport", "external_ids"}, "external_ids:sandbox="+sandboxID)
	if err != nil {
		return 0, nil, err
	}
	if len(rows) == 0 {
		return 0, nil, fmt.Errorf("failed to find pod details in OVS database")
	} else if len(rows) > 1 {
		return 0, nil, fmt.Errorf("found multiple pods for sandbox ID %q: %#v", sandboxID, rows)
	}
	ofport, err := strconv.Atoi(rows[0]["ofport"])
	if err != nil {
		return 0, nil, fmt.Errorf("could not parse ofport %q: %v", rows[0]["ofport"], err)
	}
	ids, err := ovs.ParseExternalIDs(rows[0]["external_ids"])
	if err != nil {
		return 0, nil, fmt.Errorf("could not parse external_ids %q: %v", rows[0]["external_ids"], err)
	} else if ids["ip"] == "" {
		return 0, nil, fmt.Errorf("external_ids %#v does not contain IP", ids)
	}
	podIP := net.ParseIP(ids["ip"])
	if podIP == nil {
		return 0, nil, fmt.Errorf("failed to parse IP %q", ids["ip"])
	}
	return ofport, podIP, nil
}
func (oc *ovsController) UpdatePod(sandboxID string, vnid uint32) error {
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
	ofport, podIP, err := oc.getPodDetailsBySandboxID(sandboxID)
	if err != nil {
		return err
	} else if ofport == -1 {
		return fmt.Errorf("can't update pod %q with missing veth interface", sandboxID)
	}
	err = oc.cleanupPodFlows(podIP)
	if err != nil {
		return err
	}
	return oc.setupPodFlows(ofport, podIP, vnid)
}
func (oc *ovsController) TearDownPod(sandboxID string) error {
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
	_, podIP, err := oc.getPodDetailsBySandboxID(sandboxID)
	if err != nil {
		return nil
	}
	if err := oc.cleanupPodFlows(podIP); err != nil {
		return err
	}
	ports, err := oc.getInterfacesForSandbox(sandboxID)
	if err != nil {
		return err
	}
	if err := oc.ClearPodBandwidth(ports, sandboxID); err != nil {
		return err
	}
	for _, port := range ports {
		if err := oc.ovs.DeletePort(port); err != nil {
			return err
		}
	}
	return nil
}
func policyNames(policies []networkapi.EgressNetworkPolicy) string {
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
	names := make([]string, len(policies))
	for i, policy := range policies {
		names[i] = policy.Namespace + ":" + policy.Name
	}
	return strings.Join(names, ", ")
}
func (oc *ovsController) UpdateEgressNetworkPolicyRules(policies []networkapi.EgressNetworkPolicy, vnid uint32, namespaces []string, egressDNS *common.EgressDNS) error {
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
	otx := oc.ovs.NewTransaction()
	errs := []error{}
	if len(policies) == 0 {
		otx.DeleteFlows("table=101, reg0=%d", vnid)
	} else if vnid == 0 {
		errs = append(errs, fmt.Errorf("EgressNetworkPolicy in global network namespace is not allowed (%s); ignoring", policyNames(policies)))
	} else if len(namespaces) > 1 {
		errs = append(errs, fmt.Errorf("EgressNetworkPolicy not allowed in shared NetNamespace (%s); dropping all traffic", strings.Join(namespaces, ", ")))
		otx.DeleteFlows("table=101, reg0=%d", vnid)
		otx.AddFlow("table=101, reg0=%d, priority=1, actions=drop", vnid)
	} else if len(policies) > 1 {
		errs = append(errs, fmt.Errorf("multiple EgressNetworkPolicies in same network namespace (%s) is not allowed; dropping all traffic", policyNames(policies)))
		otx.DeleteFlows("table=101, reg0=%d", vnid)
		otx.AddFlow("table=101, reg0=%d, priority=1, actions=drop", vnid)
	} else {
		otx.DeleteFlows("table=101, reg0=%d", vnid)
		for i, rule := range policies[0].Spec.Egress {
			priority := len(policies[0].Spec.Egress) - i
			var action string
			if rule.Type == networkapi.EgressNetworkPolicyRuleAllow {
				action = "output:2"
			} else {
				action = "drop"
			}
			var selectors []string
			if len(rule.To.CIDRSelector) > 0 {
				selectors = append(selectors, rule.To.CIDRSelector)
			} else if len(rule.To.DNSName) > 0 {
				ips := egressDNS.GetIPs(policies[0], rule.To.DNSName)
				for _, ip := range ips {
					selectors = append(selectors, ip.String())
				}
			}
			for _, selector := range selectors {
				var dst string
				if selector == "0.0.0.0/0" {
					dst = ""
				} else if selector == "0.0.0.0/32" {
					klog.Warningf("Correcting CIDRSelector '0.0.0.0/32' to '0.0.0.0/0' in EgressNetworkPolicy %s:%s", policies[0].Namespace, policies[0].Name)
					dst = ""
				} else {
					dst = fmt.Sprintf(", nw_dst=%s", selector)
				}
				otx.AddFlow("table=101, reg0=%d, priority=%d, ip%s, actions=%s", vnid, priority, dst, action)
			}
		}
	}
	if txErr := otx.Commit(); txErr != nil {
		errs = append(errs, txErr)
	}
	return kerrors.NewAggregate(errs)
}
func hostSubnetCookie(subnet *networkapi.HostSubnet) uint32 {
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
	hash := sha256.Sum256([]byte(subnet.UID))
	return (uint32(hash[0]) << 24) | (uint32(hash[1]) << 16) | (uint32(hash[2]) << 8) | uint32(hash[3])
}
func (oc *ovsController) AddHostSubnetRules(subnet *networkapi.HostSubnet) error {
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
	cookie := hostSubnetCookie(subnet)
	otx := oc.ovs.NewTransaction()
	otx.AddFlow("table=10, priority=100, cookie=0x%08x, tun_src=%s, actions=goto_table:30", cookie, subnet.HostIP)
	if vnid, ok := subnet.Annotations[networkapi.FixedVNIDHostAnnotation]; ok {
		otx.AddFlow("table=50, priority=100, cookie=0x%08x, arp, nw_dst=%s, actions=load:%s->NXM_NX_TUN_ID[0..31],set_field:%s->tun_dst,output:1", cookie, subnet.Subnet, vnid, subnet.HostIP)
		otx.AddFlow("table=90, priority=100, cookie=0x%08x, ip, nw_dst=%s, actions=load:%s->NXM_NX_TUN_ID[0..31],set_field:%s->tun_dst,output:1", cookie, subnet.Subnet, vnid, subnet.HostIP)
	} else {
		otx.AddFlow("table=50, priority=100, cookie=0x%08x, arp, nw_dst=%s, actions=move:NXM_NX_REG0[]->NXM_NX_TUN_ID[0..31],set_field:%s->tun_dst,output:1", cookie, subnet.Subnet, subnet.HostIP)
		otx.AddFlow("table=90, priority=100, cookie=0x%08x, ip, nw_dst=%s, actions=move:NXM_NX_REG0[]->NXM_NX_TUN_ID[0..31],set_field:%s->tun_dst,output:1", cookie, subnet.Subnet, subnet.HostIP)
	}
	return otx.Commit()
}
func (oc *ovsController) DeleteHostSubnetRules(subnet *networkapi.HostSubnet) error {
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
	cookie := hostSubnetCookie(subnet)
	otx := oc.ovs.NewTransaction()
	otx.DeleteFlows("table=10, cookie=0x%08x/0xffffffff, tun_src=%s", cookie, subnet.HostIP)
	otx.DeleteFlows("table=50, cookie=0x%08x/0xffffffff, arp, nw_dst=%s", cookie, subnet.Subnet)
	otx.DeleteFlows("table=90, cookie=0x%08x/0xffffffff, ip, nw_dst=%s", cookie, subnet.Subnet)
	return otx.Commit()
}
func (oc *ovsController) AddServiceRules(service *corev1.Service, netID uint32) error {
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
	otx := oc.ovs.NewTransaction()
	action := fmt.Sprintf(", priority=100, actions=load:%d->NXM_NX_REG1[], load:2->NXM_NX_REG2[], goto_table:80", netID)
	otx.AddFlow(generateBaseServiceRule(service.Spec.ClusterIP) + ", ip_frag=later" + action)
	for _, port := range service.Spec.Ports {
		baseRule, err := generateBaseAddServiceRule(service.Spec.ClusterIP, port.Protocol, int(port.Port))
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Error creating OVS flow for service %v, netid %d: %v", service, netID, err))
		}
		otx.AddFlow(baseRule + action)
	}
	return otx.Commit()
}
func (oc *ovsController) DeleteServiceRules(service *corev1.Service) error {
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
	otx := oc.ovs.NewTransaction()
	otx.DeleteFlows(generateBaseServiceRule(service.Spec.ClusterIP))
	return otx.Commit()
}
func generateBaseServiceRule(IP string) string {
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
	return fmt.Sprintf("table=60, ip, nw_dst=%s", IP)
}
func generateBaseAddServiceRule(IP string, protocol corev1.Protocol, port int) (string, error) {
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
	var dst string
	if protocol == corev1.ProtocolUDP {
		dst = fmt.Sprintf(", udp, udp_dst=%d", port)
	} else if protocol == corev1.ProtocolTCP {
		dst = fmt.Sprintf(", tcp, tcp_dst=%d", port)
	} else {
		return "", fmt.Errorf("unhandled protocol %v", protocol)
	}
	return generateBaseServiceRule(IP) + dst, nil
}
func (oc *ovsController) UpdateLocalMulticastFlows(vnid uint32, enabled bool, ofports []int) error {
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
	otx := oc.ovs.NewTransaction()
	if enabled {
		otx.AddFlow("table=110, reg0=%d, actions=goto_table:111", vnid)
	} else {
		otx.DeleteFlows("table=110, reg0=%d", vnid)
	}
	var actions []string
	if enabled && len(ofports) > 0 {
		actions = make([]string, len(ofports))
		for i, ofport := range ofports {
			actions[i] = fmt.Sprintf("output:%d", ofport)
		}
		sort.Strings(actions)
		otx.AddFlow("table=120, priority=100, reg0=%d, actions=%s", vnid, strings.Join(actions, ","))
	} else {
		otx.DeleteFlows("table=120, reg0=%d", vnid)
	}
	return otx.Commit()
}
func (oc *ovsController) UpdateVXLANMulticastFlows(remoteIPs []string) error {
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
	otx := oc.ovs.NewTransaction()
	if len(remoteIPs) > 0 {
		actions := make([]string, len(remoteIPs))
		for i, ip := range remoteIPs {
			actions[i] = fmt.Sprintf("set_field:%s->tun_dst,output:1", ip)
		}
		sort.Strings(actions)
		otx.AddFlow("table=111, priority=100, actions=move:NXM_NX_REG0[]->NXM_NX_TUN_ID[0..31],%s,goto_table:120", strings.Join(actions, ","))
	} else {
		otx.AddFlow("table=111, priority=100, actions=goto_table:120")
	}
	return otx.Commit()
}
func (oc *ovsController) FindPolicyVNIDs() sets.Int {
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
	_, policyVNIDs := oc.findInUseAndPolicyVNIDs()
	return policyVNIDs
}
func (oc *ovsController) FindUnusedVNIDs() []int {
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
	inUseVNIDs, policyVNIDs := oc.findInUseAndPolicyVNIDs()
	return policyVNIDs.Difference(inUseVNIDs).UnsortedList()
}
func (oc *ovsController) findInUseAndPolicyVNIDs() (sets.Int, sets.Int) {
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
	inUseVNIDs := sets.NewInt(0)
	policyVNIDs := sets.NewInt()
	flows, err := oc.ovs.DumpFlows("")
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("findInUseAndPolicyVNIDs: could not DumpFlows: %v", err))
		return inUseVNIDs, policyVNIDs
	}
	for _, flow := range flows {
		parsed, err := ovs.ParseFlow(ovs.ParseForDump, flow)
		if err != nil {
			klog.Warningf("findInUseAndPolicyVNIDs: could not parse flow %q: %v", flow, err)
			continue
		}
		if parsed.Table == 60 || parsed.Table == 70 {
			for _, action := range parsed.Actions {
				if action.Name != "load" || strings.Index(action.Value, "REG1") == -1 {
					continue
				}
				vnidEnd := strings.Index(action.Value, "->")
				if vnidEnd == -1 {
					continue
				}
				vnid, err := strconv.ParseInt(action.Value[:vnidEnd], 0, 32)
				if err != nil {
					klog.Warningf("findInUseAndPolicyVNIDs: could not parse VNID in 'load:%s': %v", action.Value, err)
					continue
				}
				inUseVNIDs.Insert(int(vnid))
				break
			}
		}
		if parsed.Table == 80 {
			if field, exists := parsed.FindField("reg1"); exists {
				vnid, err := strconv.ParseInt(field.Value, 0, 32)
				if err != nil {
					klog.Warningf("findInUseAndPolicyVNIDs: could not parse VNID in 'reg1=%s': %v", field.Value, err)
					continue
				}
				policyVNIDs.Insert(int(vnid))
			}
		}
	}
	return inUseVNIDs, policyVNIDs
}
func (oc *ovsController) ensureTunMAC() error {
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
	if oc.tunMAC != "" {
		return nil
	}
	val, err := oc.ovs.Get("Interface", Tun0, "mac_in_use")
	if err != nil {
		return fmt.Errorf("could not get %s MAC address: %v", Tun0, err)
	} else if len(val) != 19 || val[0] != '"' || val[18] != '"' {
		return fmt.Errorf("bad MAC address for %s: %q", Tun0, val)
	}
	oc.tunMAC = val[1:18]
	return nil
}
func (oc *ovsController) SetNamespaceEgressNormal(vnid uint32) error {
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
	otx := oc.ovs.NewTransaction()
	otx.DeleteFlows("table=100, reg0=%d", vnid)
	return otx.Commit()
}
func (oc *ovsController) SetNamespaceEgressDropped(vnid uint32) error {
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
	otx := oc.ovs.NewTransaction()
	otx.DeleteFlows("table=100, reg0=%d", vnid)
	otx.AddFlow("table=100, priority=100, reg0=%d, actions=drop", vnid)
	return otx.Commit()
}
func (oc *ovsController) SetNamespaceEgressViaEgressIP(vnid uint32, nodeIP, mark string) error {
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
	otx := oc.ovs.NewTransaction()
	otx.DeleteFlows("table=100, reg0=%d", vnid)
	if nodeIP == "" {
		otx.AddFlow("table=100, priority=100, reg0=%d, actions=drop", vnid)
	} else if nodeIP == oc.localIP {
		if err := oc.ensureTunMAC(); err != nil {
			return err
		}
		otx.AddFlow("table=100, priority=100, reg0=%d, ip, actions=set_field:%s->eth_dst,set_field:%s->pkt_mark,goto_table:101", vnid, oc.tunMAC, mark)
	} else {
		otx.AddFlow("table=100, priority=100, reg0=%d, ip, actions=move:NXM_NX_REG0[]->NXM_NX_TUN_ID[0..31],set_field:%s->tun_dst,output:1", vnid, nodeIP)
	}
	return otx.Commit()
}

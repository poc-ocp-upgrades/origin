package node

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	networkapi "github.com/openshift/api/network/v1"
)

func assertNetlinkChange(eip *egressIPWatcher, expected ...string) error {
	_logClusterCodePath()
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
	actual := []string{}
	for range expected {
		select {
		case change := <-eip.testModeChan:
			actual = append(actual, change)
		default:
			break
		}
	}
	sort.Strings(expected)
	sort.Strings(actual)
	if reflect.DeepEqual(expected, actual) {
		return nil
	}
	return fmt.Errorf("Unexpected netlink changes: expected %#v, got %#v", expected, actual)
}
func assertNoNetlinkChanges(eip *egressIPWatcher) error {
	_logClusterCodePath()
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
	select {
	case change := <-eip.testModeChan:
		return fmt.Errorf("Unexpected netlink change %q", change)
	default:
		return nil
	}
}

type egressTrafficType string

const (
	Normal	egressTrafficType	= "normal"
	Dropped	egressTrafficType	= "dropped"
	Local	egressTrafficType	= "local"
	Remote	egressTrafficType	= "remote"
)

type egressOVSChange struct {
	vnid	uint32
	egress	egressTrafficType
	remote	string
}

func assertOVSChanges(eip *egressIPWatcher, flows *[]string, changes ...egressOVSChange) error {
	_logClusterCodePath()
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
	oldFlows := *flows
	newFlows, err := eip.oc.ovs.DumpFlows("table=100")
	if err != nil {
		return fmt.Errorf("unexpected error dumping OVS flows: %v", err)
	}
	flowChanges := []flowChange{}
	for _, change := range changes {
		vnidStr := fmt.Sprintf("reg0=%d", change.vnid)
		for _, flow := range *flows {
			if strings.Contains(flow, vnidStr) {
				flowChanges = append(flowChanges, flowChange{kind: flowRemoved, match: []string{flow}})
			}
		}
		switch change.egress {
		case Normal:
			break
		case Dropped:
			flowChanges = append(flowChanges, flowChange{kind: flowAdded, match: []string{vnidStr, "drop"}})
		case Local:
			flowChanges = append(flowChanges, flowChange{kind: flowAdded, match: []string{vnidStr, fmt.Sprintf("%s->pkt_mark", getMarkForVNID(change.vnid, eip.masqueradeBit)), "goto_table:101"}})
		case Remote:
			flowChanges = append(flowChanges, flowChange{kind: flowAdded, match: []string{vnidStr, fmt.Sprintf("%s->tun_dst", change.remote)}})
		}
	}
	err = assertFlowChanges(oldFlows, newFlows, flowChanges...)
	if err != nil {
		return fmt.Errorf("unexpected flow changes: %v\nOrig:\n%s\nNew:\n%s", err, strings.Join(oldFlows, "\n"), strings.Join(newFlows, "\n"))
	}
	*flows = newFlows
	return nil
}
func assertNoOVSChanges(eip *egressIPWatcher, flows *[]string) error {
	_logClusterCodePath()
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
	return assertOVSChanges(eip, flows)
}
func setupEgressIPWatcher(t *testing.T) (*egressIPWatcher, []string) {
	_logClusterCodePath()
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
	_, oc, _ := setupOVSController(t)
	if oc.localIP != "172.17.0.4" {
		panic("details of fake ovsController changed")
	}
	masqBit := int32(0)
	eip := newEgressIPWatcher(oc, "172.17.0.4", &masqBit)
	eip.testModeChan = make(chan string, 10)
	flows, err := eip.oc.ovs.DumpFlows("table=100")
	if err != nil {
		t.Fatalf("unexpected error dumping OVS flows: %v", err)
	}
	return eip, flows
}
func updateNodeEgress(eip *egressIPWatcher, nodeIP string, egressIPs []string) {
	_logClusterCodePath()
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
	name := "node-" + nodeIP[strings.LastIndex(nodeIP, ".")+1:]
	eip.tracker.UpdateHostSubnetEgress(&networkapi.HostSubnet{ObjectMeta: metav1.ObjectMeta{Name: name, UID: ktypes.UID(name)}, Host: name, HostIP: nodeIP, EgressIPs: egressIPs})
}
func updateNamespaceEgress(eip *egressIPWatcher, vnid uint32, egressIPs []string) {
	_logClusterCodePath()
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
	name := fmt.Sprintf("ns-%d", vnid)
	eip.tracker.UpdateNetNamespaceEgress(&networkapi.NetNamespace{ObjectMeta: metav1.ObjectMeta{Name: name, UID: ktypes.UID(name)}, NetName: name, NetID: vnid, EgressIPs: egressIPs})
}
func deleteNamespaceEgress(eip *egressIPWatcher, vnid uint32) {
	_logClusterCodePath()
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
	eip.tracker.DeleteNetNamespaceEgress(vnid)
}
func TestEgressIP(t *testing.T) {
	_logClusterCodePath()
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
	eip, flows := setupEgressIPWatcher(t)
	updateNodeEgress(eip, "172.17.0.3", []string{})
	updateNodeEgress(eip, "172.17.0.4", []string{})
	deleteNamespaceEgress(eip, 42)
	deleteNamespaceEgress(eip, 43)
	err := assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertNoOVSChanges(eip, &flows)
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 42, []string{"172.17.0.100"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.3", []string{"172.17.0.100"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.3", []string{"172.17.0.101", "172.17.0.100"})
	updateNodeEgress(eip, "172.17.0.5", []string{"172.17.0.105"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertNoOVSChanges(eip, &flows)
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 43, []string{"172.17.0.105"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 43, egress: Remote, remote: "172.17.0.5"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 43, []string{"172.17.0.101"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 43, egress: Remote, remote: "172.17.0.3"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 44, []string{"172.17.0.104"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 44, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.4", []string{"172.17.0.102", "172.17.0.104"})
	err = assertNetlinkChange(eip, "claim 172.17.0.104")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 44, egress: Local})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 44, []string{"172.17.0.102"})
	err = assertNetlinkChange(eip, "release 172.17.0.104", "claim 172.17.0.102")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertNoOVSChanges(eip, &flows)
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.4", []string{"172.17.0.102", "172.17.0.103"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertNoOVSChanges(eip, &flows)
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 45, []string{"172.17.0.103"})
	err = assertNetlinkChange(eip, "claim 172.17.0.103")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 45, egress: Local})
	if err != nil {
		t.Fatalf("%v", err)
	}
	deleteNamespaceEgress(eip, 44)
	err = assertNetlinkChange(eip, "release 172.17.0.102")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 44, egress: Normal})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 44, []string{"172.17.0.102"})
	err = assertNetlinkChange(eip, "claim 172.17.0.102")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 44, egress: Local})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.3", []string{"172.17.0.100"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 43, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.4", []string{"172.17.0.102"})
	err = assertNetlinkChange(eip, "release 172.17.0.103")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 45, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.3", []string{"172.17.0.100", "172.17.0.103"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 45, egress: Remote, remote: "172.17.0.3"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.4", []string{"172.17.0.101", "172.17.0.102"})
	err = assertNetlinkChange(eip, "claim 172.17.0.101")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 43, egress: Local})
	if err != nil {
		t.Fatalf("%v", err)
	}
}
func TestMultipleNamespaceEgressIPs(t *testing.T) {
	_logClusterCodePath()
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
	eip, flows := setupEgressIPWatcher(t)
	updateNamespaceEgress(eip, 42, []string{"172.17.0.100"})
	updateNodeEgress(eip, "172.17.0.3", []string{"172.17.0.100"})
	err := assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 42, []string{"172.17.0.101", "172.17.0.100"})
	err = assertNoOVSChanges(eip, &flows)
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.4", []string{"172.17.0.101"})
	err = assertNetlinkChange(eip, "claim 172.17.0.101")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Local})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 42, []string{"172.17.0.100", "172.17.0.101"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.4", []string{"172.17.0.200"})
	err = assertNetlinkChange(eip, "release 172.17.0.101")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertNoOVSChanges(eip, &flows)
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.3", nil)
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.3", []string{"172.17.0.100"})
	updateNodeEgress(eip, "172.17.0.4", []string{"172.17.0.101"})
	err = assertNetlinkChange(eip, "claim 172.17.0.101")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 43, []string{"172.17.0.100"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Dropped}, egressOVSChange{vnid: 43, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	deleteNamespaceEgress(eip, 43)
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"}, egressOVSChange{vnid: 43, egress: Normal})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 44, []string{"172.17.0.101"})
	err = assertNetlinkChange(eip, "release 172.17.0.101")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Dropped}, egressOVSChange{vnid: 44, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	deleteNamespaceEgress(eip, 44)
	err = assertNetlinkChange(eip, "claim 172.17.0.101")
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"}, egressOVSChange{vnid: 44, egress: Normal})
	if err != nil {
		t.Fatalf("%v", err)
	}
}
func TestNodeIPAsEgressIP(t *testing.T) {
	_logClusterCodePath()
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
	eip, flows := setupEgressIPWatcher(t)
	updateNodeEgress(eip, "172.17.0.4", []string{"172.17.0.4", "172.17.0.102"})
	err := assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertNoOVSChanges(eip, &flows)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
func TestDuplicateNodeEgressIPs(t *testing.T) {
	_logClusterCodePath()
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
	eip, flows := setupEgressIPWatcher(t)
	updateNamespaceEgress(eip, 42, []string{"172.17.0.100"})
	updateNodeEgress(eip, "172.17.0.3", []string{"172.17.0.100"})
	err := assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.4", []string{"172.17.0.100"})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.4", []string{})
	err = assertNoNetlinkChanges(eip)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.5", []string{"172.17.0.100"})
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	deleteNamespaceEgress(eip, 42)
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Normal})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 42, []string{"172.17.0.100"})
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.3", []string{})
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.5"})
	if err != nil {
		t.Fatalf("%v", err)
	}
}
func TestDuplicateNamespaceEgressIPs(t *testing.T) {
	_logClusterCodePath()
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
	eip, flows := setupEgressIPWatcher(t)
	updateNamespaceEgress(eip, 42, []string{"172.17.0.100"})
	updateNodeEgress(eip, "172.17.0.3", []string{"172.17.0.100"})
	err := assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 43, []string{"172.17.0.100"})
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Dropped}, egressOVSChange{vnid: 43, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	deleteNamespaceEgress(eip, 43)
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"}, egressOVSChange{vnid: 43, egress: Normal})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNamespaceEgress(eip, 43, []string{"172.17.0.100"})
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Dropped}, egressOVSChange{vnid: 43, egress: Dropped})
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.3", []string{})
	err = assertNoOVSChanges(eip, &flows)
	if err != nil {
		t.Fatalf("%v", err)
	}
	updateNodeEgress(eip, "172.17.0.3", []string{"172.17.0.100"})
	err = assertNoOVSChanges(eip, &flows)
	if err != nil {
		t.Fatalf("%v", err)
	}
	deleteNamespaceEgress(eip, 42)
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Normal}, egressOVSChange{vnid: 43, egress: Remote, remote: "172.17.0.3"})
	if err != nil {
		t.Fatalf("%v", err)
	}
}
func TestMarkForVNID(t *testing.T) {
	_logClusterCodePath()
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
	testcases := []struct {
		description	string
		vnid		uint32
		masqueradeBit	uint32
		result		uint32
	}{{description: "masqBit in VNID range, but not set in VNID", vnid: 0x000000aa, masqueradeBit: 0x00000001, result: 0x000000aa}, {description: "masqBit in VNID range, and set in VNID", vnid: 0x000000ab, masqueradeBit: 0x00000001, result: 0x010000aa}, {description: "masqBit in VNID range, VNID 0", vnid: 0x00000000, masqueradeBit: 0x00000001, result: 0xff000000}, {description: "masqBit outside of VNID range", vnid: 0x000000aa, masqueradeBit: 0x80000000, result: 0x000000aa}, {description: "masqBit outside of VNID range, VNID 0", vnid: 0x00000000, masqueradeBit: 0x80000000, result: 0x7f000000}, {description: "masqBit == bit 24", vnid: 0x000000aa, masqueradeBit: 0x01000000, result: 0x000000aa}, {description: "masqBit == bit 24, VNID 0", vnid: 0x00000000, masqueradeBit: 0x01000000, result: 0xfe000000}, {description: "no masqBit, ordinary VNID", vnid: 0x000000aa, masqueradeBit: 0x00000000, result: 0x000000aa}, {description: "no masqBit, VNID 0", vnid: 0x00000000, masqueradeBit: 0x00000000, result: 0xff000000}}
	for _, tc := range testcases {
		result := getMarkForVNID(tc.vnid, tc.masqueradeBit)
		if result != fmt.Sprintf("0x%08x", tc.result) {
			t.Fatalf("test %q expected %08x got %s", tc.description, tc.result, result)
		}
	}
}
func TestEgressNodeRenumbering(t *testing.T) {
	_logClusterCodePath()
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
	eip, flows := setupEgressIPWatcher(t)
	eip.tracker.UpdateHostSubnetEgress(&networkapi.HostSubnet{ObjectMeta: metav1.ObjectMeta{Name: "alpha", UID: ktypes.UID("alpha")}, Host: "alpha", HostIP: "172.17.0.3", EgressIPs: []string{"172.17.0.100"}})
	eip.tracker.UpdateHostSubnetEgress(&networkapi.HostSubnet{ObjectMeta: metav1.ObjectMeta{Name: "beta", UID: ktypes.UID("beta")}, Host: "beta", HostIP: "172.17.0.4", EgressIPs: []string{"172.17.0.101"}})
	eip.tracker.UpdateHostSubnetEgress(&networkapi.HostSubnet{ObjectMeta: metav1.ObjectMeta{Name: "gamma", UID: ktypes.UID("gamma")}, Host: "gamma", HostIP: "172.17.0.5", EgressIPs: []string{"172.17.0.102"}})
	updateNamespaceEgress(eip, 42, []string{"172.17.0.100"})
	updateNamespaceEgress(eip, 43, []string{"172.17.0.101"})
	err := assertOVSChanges(eip, &flows, egressOVSChange{vnid: 42, egress: Remote, remote: "172.17.0.3"}, egressOVSChange{vnid: 43, egress: Local})
	if err != nil {
		t.Fatalf("%v", err)
	}
	eip.tracker.UpdateHostSubnetEgress(&networkapi.HostSubnet{ObjectMeta: metav1.ObjectMeta{Name: "beta", UID: ktypes.UID("beta")}, Host: "beta", HostIP: "172.17.0.6", EgressIPs: []string{"172.17.0.101"}})
	err = assertOVSChanges(eip, &flows, egressOVSChange{vnid: 43, egress: Remote, remote: "172.17.0.6"})
	if err != nil {
		t.Fatalf("%v", err)
	}
}

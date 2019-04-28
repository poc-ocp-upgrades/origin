package common

import (
	"net"
	"sync"
	"time"
	networkapi "github.com/openshift/api/network/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

type EgressDNSUpdate struct {
	UID		ktypes.UID
	Namespace	string
}
type EgressDNS struct {
	lock		sync.Mutex
	pdMap		map[ktypes.UID]*DNS
	namespaces	map[ktypes.UID]string
	added		chan bool
	Updates		chan EgressDNSUpdate
}

func NewEgressDNS() *EgressDNS {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &EgressDNS{pdMap: map[ktypes.UID]*DNS{}, namespaces: map[ktypes.UID]string{}, added: make(chan bool), Updates: make(chan EgressDNSUpdate)}
}
func (e *EgressDNS) Add(policy networkapi.EgressNetworkPolicy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dnsInfo, err := NewDNS("/etc/resolv.conf")
	if err != nil {
		utilruntime.HandleError(err)
	}
	for _, rule := range policy.Spec.Egress {
		if len(rule.To.DNSName) > 0 {
			if err := dnsInfo.Add(rule.To.DNSName); err != nil {
				utilruntime.HandleError(err)
			}
		}
	}
	if dnsInfo.Size() > 0 {
		e.lock.Lock()
		defer e.lock.Unlock()
		e.pdMap[policy.UID] = dnsInfo
		e.namespaces[policy.UID] = policy.Namespace
		e.signalAdded()
	}
}
func (e *EgressDNS) Delete(policy networkapi.EgressNetworkPolicy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.lock.Lock()
	defer e.lock.Unlock()
	if _, ok := e.pdMap[policy.UID]; ok {
		delete(e.pdMap, policy.UID)
		delete(e.namespaces, policy.UID)
	}
}
func (e *EgressDNS) Update(policyUID ktypes.UID) (error, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.lock.Lock()
	defer e.lock.Unlock()
	if dnsInfo, ok := e.pdMap[policyUID]; ok {
		return dnsInfo.Update()
	}
	return nil, false
}
func (e *EgressDNS) Sync() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var duration time.Duration
	for {
		tm, policyUID, policyNamespace, ok := e.GetMinQueryTime()
		if !ok {
			duration = 30 * time.Minute
		} else {
			now := time.Now()
			if tm.After(now) {
				duration = tm.Sub(now)
			} else {
				err, changed := e.Update(policyUID)
				if err != nil {
					utilruntime.HandleError(err)
				}
				if changed {
					e.Updates <- EgressDNSUpdate{policyUID, policyNamespace}
				}
				continue
			}
		}
		select {
		case <-e.added:
		case <-time.After(duration):
		}
	}
}
func (e *EgressDNS) GetMinQueryTime() (time.Time, ktypes.UID, string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.lock.Lock()
	defer e.lock.Unlock()
	timeSet := false
	var minTime time.Time
	var uid ktypes.UID
	for policyUID, dnsInfo := range e.pdMap {
		tm, ok := dnsInfo.GetMinQueryTime()
		if !ok {
			continue
		}
		if (timeSet == false) || tm.Before(minTime) {
			timeSet = true
			minTime = tm
			uid = policyUID
		}
	}
	return minTime, uid, e.namespaces[uid], timeSet
}
func (e *EgressDNS) GetIPs(policy networkapi.EgressNetworkPolicy, dnsName string) []net.IP {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.lock.Lock()
	defer e.lock.Unlock()
	dnsInfo, ok := e.pdMap[policy.UID]
	if !ok {
		return []net.IP{}
	}
	return dnsInfo.Get(dnsName).ips
}
func (e *EgressDNS) GetNetCIDRs(policy networkapi.EgressNetworkPolicy, dnsName string) []net.IPNet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cidrs := []net.IPNet{}
	for _, ip := range e.GetIPs(policy, dnsName) {
		cidrs = append(cidrs, net.IPNet{IP: ip, Mask: net.CIDRMask(32, 32)})
	}
	return cidrs
}
func (e *EgressDNS) signalAdded() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case e.added <- true:
	default:
	}
}

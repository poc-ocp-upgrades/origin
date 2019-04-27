package common

import (
	"fmt"
	"net"
	"sync"
	"time"
	"github.com/miekg/dns"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	defaultTTL = 30 * time.Minute
)

type dnsValue struct {
	ips		[]net.IP
	ttl		time.Duration
	nextQueryTime	time.Time
}
type DNS struct {
	lock		sync.Mutex
	dnsMap		map[string]dnsValue
	nameservers	[]string
	port		string
}

func NewDNS(resolverConfigFile string) (*DNS, error) {
	_logClusterCodePath()
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
	config, err := dns.ClientConfigFromFile(resolverConfigFile)
	if err != nil || config == nil {
		return nil, fmt.Errorf("cannot initialize the resolver: %v", err)
	}
	return &DNS{dnsMap: map[string]dnsValue{}, nameservers: filterIPv4Servers(config.Servers), port: config.Port}, nil
}
func (d *DNS) Size() int {
	_logClusterCodePath()
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
	d.lock.Lock()
	defer d.lock.Unlock()
	return len(d.dnsMap)
}
func (d *DNS) Get(dns string) dnsValue {
	_logClusterCodePath()
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
	d.lock.Lock()
	defer d.lock.Unlock()
	data := dnsValue{}
	if res, ok := d.dnsMap[dns]; ok {
		data.ips = make([]net.IP, len(res.ips))
		copy(data.ips, res.ips)
		data.ttl = res.ttl
		data.nextQueryTime = res.nextQueryTime
	}
	return data
}
func (d *DNS) Add(dns string) error {
	_logClusterCodePath()
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
	d.lock.Lock()
	defer d.lock.Unlock()
	d.dnsMap[dns] = dnsValue{}
	err, _ := d.updateOne(dns)
	if err != nil {
		delete(d.dnsMap, dns)
	}
	return err
}
func (d *DNS) Update() (error, bool) {
	_logClusterCodePath()
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
	d.lock.Lock()
	defer d.lock.Unlock()
	errList := []error{}
	changed := false
	for dns := range d.dnsMap {
		err, updated := d.updateOne(dns)
		if err != nil {
			errList = append(errList, err)
			continue
		}
		if updated {
			changed = true
		}
	}
	return kerrors.NewAggregate(errList), changed
}
func (d *DNS) updateOne(dns string) (error, bool) {
	_logClusterCodePath()
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
	res, ok := d.dnsMap[dns]
	if !ok {
		return fmt.Errorf("DNS value not found in dnsMap for domain: %q", dns), false
	}
	ips, minTTL, err := d.getIPsAndMinTTL(dns)
	if err != nil {
		res.nextQueryTime = time.Now().Add(defaultTTL)
		d.dnsMap[dns] = res
		return err, false
	}
	changed := false
	if !ipsEqual(res.ips, ips) {
		changed = true
	}
	res.ips = ips
	res.ttl = minTTL
	res.nextQueryTime = time.Now().Add(res.ttl)
	d.dnsMap[dns] = res
	return nil, changed
}
func (d *DNS) getIPsAndMinTTL(domain string) ([]net.IP, time.Duration, error) {
	_logClusterCodePath()
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
	ips := []net.IP{}
	ttlSet := false
	var minTTL uint32
	for _, server := range d.nameservers {
		msg := new(dns.Msg)
		msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
		dialServer := server
		if _, _, err := net.SplitHostPort(server); err != nil {
			dialServer = net.JoinHostPort(server, d.port)
		}
		c := new(dns.Client)
		c.Timeout = 5 * time.Second
		in, _, err := c.Exchange(msg, dialServer)
		if err != nil {
			return nil, defaultTTL, err
		}
		if in != nil && in.Rcode != dns.RcodeSuccess {
			return nil, defaultTTL, fmt.Errorf("failed to get a valid answer: %v", in)
		}
		if in != nil && len(in.Answer) > 0 {
			for _, a := range in.Answer {
				if !ttlSet || a.Header().Ttl < minTTL {
					minTTL = a.Header().Ttl
					ttlSet = true
				}
				switch t := a.(type) {
				case *dns.A:
					ips = append(ips, t.A)
				}
			}
		}
	}
	if !ttlSet || (len(ips) == 0) {
		return nil, defaultTTL, fmt.Errorf("IPv4 addr not found for domain: %q, nameservers: %v", domain, d.nameservers)
	}
	ttl, err := time.ParseDuration(fmt.Sprintf("%ds", minTTL))
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Invalid TTL value for domain: %q, err: %v, defaulting ttl=%s", domain, err, defaultTTL.String()))
		ttl = defaultTTL
	}
	return removeDuplicateIPs(ips), ttl, nil
}
func (d *DNS) GetMinQueryTime() (time.Time, bool) {
	_logClusterCodePath()
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
	d.lock.Lock()
	defer d.lock.Unlock()
	timeSet := false
	var minTime time.Time
	for _, res := range d.dnsMap {
		if (timeSet == false) || res.nextQueryTime.Before(minTime) {
			timeSet = true
			minTime = res.nextQueryTime
		}
	}
	return minTime, timeSet
}
func ipsEqual(oldips, newips []net.IP) bool {
	_logClusterCodePath()
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
	if len(oldips) != len(newips) {
		return false
	}
	for _, oldip := range oldips {
		found := false
		for _, newip := range newips {
			if oldip.Equal(newip) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
func filterIPv4Servers(servers []string) []string {
	_logClusterCodePath()
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
	ipv4Servers := []string{}
	for _, server := range servers {
		ipString := server
		if host, _, err := net.SplitHostPort(server); err == nil {
			ipString = host
		}
		if ip := net.ParseIP(ipString); ip != nil {
			if ip.To4() != nil {
				ipv4Servers = append(ipv4Servers, server)
			}
		}
	}
	return ipv4Servers
}
func removeDuplicateIPs(ips []net.IP) []net.IP {
	_logClusterCodePath()
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
	ipSet := sets.NewString()
	for _, ip := range ips {
		ipSet.Insert(ip.String())
	}
	uniqueIPs := []net.IP{}
	for _, str := range ipSet.List() {
		ip := net.ParseIP(str)
		if ip != nil {
			uniqueIPs = append(uniqueIPs, ip)
		}
	}
	return uniqueIPs
}

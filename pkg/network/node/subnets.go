package node

import (
	"fmt"
	"time"
	"k8s.io/klog"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	networkapi "github.com/openshift/api/network/v1"
	networkinformers "github.com/openshift/client-go/network/informers/externalversions"
	"github.com/openshift/origin/pkg/network/common"
)

type hostSubnetWatcher struct {
	oc		*ovsController
	localIP		string
	networkInfo	*common.NetworkInfo
	hostSubnetMap	map[ktypes.UID]*networkapi.HostSubnet
}

func newHostSubnetWatcher(oc *ovsController, localIP string, networkInfo *common.NetworkInfo) *hostSubnetWatcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &hostSubnetWatcher{oc: oc, localIP: localIP, networkInfo: networkInfo, hostSubnetMap: make(map[ktypes.UID]*networkapi.HostSubnet)}
}
func (hsw *hostSubnetWatcher) Start(networkInformers networkinformers.SharedInformerFactory) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	funcs := common.InformerFuncs(&networkapi.HostSubnet{}, hsw.handleAddOrUpdateHostSubnet, hsw.handleDeleteHostSubnet)
	networkInformers.Network().V1().HostSubnets().Informer().AddEventHandler(funcs)
}
func (hsw *hostSubnetWatcher) handleAddOrUpdateHostSubnet(obj, _ interface{}, eventType watch.EventType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	hs := obj.(*networkapi.HostSubnet)
	klog.V(5).Infof("Watch %s event for HostSubnet %q", eventType, hs.Name)
	if err := common.ValidateHostSubnet(hs); err != nil {
		utilruntime.HandleError(fmt.Errorf("Ignoring invalid HostSubnet %s: %v", common.HostSubnetToString(hs), err))
		return
	}
	if err := hsw.updateHostSubnet(hs); err != nil {
		utilruntime.HandleError(err)
	}
}
func (hsw *hostSubnetWatcher) handleDeleteHostSubnet(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	hs := obj.(*networkapi.HostSubnet)
	klog.V(5).Infof("Watch %s event for HostSubnet %q", watch.Deleted, hs.Name)
	if err := hsw.deleteHostSubnet(hs); err != nil {
		utilruntime.HandleError(err)
	}
}
func (hsw *hostSubnetWatcher) updateHostSubnet(hs *networkapi.HostSubnet) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if hs.HostIP == hsw.localIP {
		return nil
	}
	oldSubnet, exists := hsw.hostSubnetMap[hs.UID]
	if exists {
		if oldSubnet.HostIP == hs.HostIP {
			return nil
		} else {
			hsw.oc.DeleteHostSubnetRules(oldSubnet)
		}
	}
	if err := hsw.networkInfo.ValidateNodeIP(hs.HostIP); err != nil {
		return fmt.Errorf("ignoring invalid subnet for node %s: %v", hs.HostIP, err)
	}
	hsw.hostSubnetMap[hs.UID] = hs
	errList := []error{}
	if err := hsw.oc.AddHostSubnetRules(hs); err != nil {
		errList = append(errList, fmt.Errorf("error adding OVS flows for subnet %q: %v", hs.Subnet, err))
	}
	if err := hsw.updateVXLANMulticastRules(); err != nil {
		errList = append(errList, fmt.Errorf("error updating OVS VXLAN multicast flows: %v", err))
	}
	return kerrors.NewAggregate(errList)
}
func (hsw *hostSubnetWatcher) deleteHostSubnet(hs *networkapi.HostSubnet) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if hs.HostIP == hsw.localIP {
		return nil
	}
	if _, exists := hsw.hostSubnetMap[hs.UID]; !exists {
		return nil
	}
	delete(hsw.hostSubnetMap, hs.UID)
	errList := []error{}
	if err := hsw.oc.DeleteHostSubnetRules(hs); err != nil {
		errList = append(errList, fmt.Errorf("error deleting OVS flows for subnet %q: %v", hs.Subnet, err))
	}
	if err := hsw.updateVXLANMulticastRules(); err != nil {
		errList = append(errList, fmt.Errorf("error updating OVS VXLAN multicast flows: %v", err))
	}
	return kerrors.NewAggregate(errList)
}
func (hsw *hostSubnetWatcher) updateVXLANMulticastRules() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	remoteIPs := make([]string, 0, len(hsw.hostSubnetMap))
	for _, subnet := range hsw.hostSubnetMap {
		if subnet.HostIP != hsw.localIP {
			remoteIPs = append(remoteIPs, subnet.HostIP)
		}
	}
	return hsw.oc.UpdateVXLANMulticastFlows(remoteIPs)
}
func (node *OsdnNode) getLocalSubnet() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var subnet *networkapi.HostSubnet
	backoff := utilwait.Backoff{Duration: time.Second, Factor: 1.5, Steps: 11}
	err := utilwait.ExponentialBackoff(backoff, func() (bool, error) {
		var err error
		subnet, err = node.networkClient.NetworkV1().HostSubnets().Get(node.hostName, metav1.GetOptions{})
		if err == nil {
			if err = common.ValidateHostSubnet(subnet); err != nil {
				return false, err
			} else if subnet.HostIP == node.localIP {
				return true, nil
			} else {
				klog.Warningf("HostIP %q for local subnet does not match with nodeIP %q, "+"Waiting for master to update subnet for node %q ...", subnet.HostIP, node.localIP, node.hostName)
				return false, nil
			}
		} else if kapierrors.IsNotFound(err) {
			klog.Warningf("Could not find an allocated subnet for node: %s, Waiting...", node.hostName)
			return false, nil
		} else {
			return false, err
		}
	})
	if err != nil {
		return "", fmt.Errorf("failed to get subnet for this host: %s, error: %v", node.hostName, err)
	}
	if err = node.networkInfo.ValidateNodeIP(subnet.HostIP); err != nil {
		return "", fmt.Errorf("failed to validate own HostSubnet: %v", err)
	}
	return subnet.Subnet, nil
}

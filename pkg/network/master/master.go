package master

import (
	"fmt"
	"time"
	"k8s.io/klog"
	kapi "k8s.io/api/core/v1"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	kcoreinformers "k8s.io/client-go/informers/core/v1"
	kclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	networkapi "github.com/openshift/api/network/v1"
	openshiftcontrolplanev1 "github.com/openshift/api/openshiftcontrolplane/v1"
	networkclient "github.com/openshift/client-go/network/clientset/versioned"
	networkinternalinformers "github.com/openshift/client-go/network/informers/externalversions"
	networkinformers "github.com/openshift/client-go/network/informers/externalversions/network/v1"
	"github.com/openshift/origin/pkg/network"
	"github.com/openshift/origin/pkg/network/common"
	"github.com/openshift/origin/pkg/util/netutils"
)

const (
	tun0 = "tun0"
)

type OsdnMaster struct {
	kClient			kclientset.Interface
	networkClient		networkclient.Interface
	networkInfo		*common.NetworkInfo
	vnids			*masterVNIDMap
	nodeInformer		kcoreinformers.NodeInformer
	namespaceInformer	kcoreinformers.NamespaceInformer
	hostSubnetInformer	networkinformers.HostSubnetInformer
	netNamespaceInformer	networkinformers.NetNamespaceInformer
	subnetAllocatorList	[]*SubnetAllocator
	subnetAllocatorMap	map[common.ClusterNetwork]*SubnetAllocator
	hostSubnetNodeIPs	map[ktypes.UID]string
}

func Start(networkConfig openshiftcontrolplanev1.NetworkControllerConfig, networkClient networkclient.Interface, kClient kclientset.Interface, kubeInformers informers.SharedInformerFactory, networkInformers networkinternalinformers.SharedInformerFactory) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("Initializing SDN master of type %q", networkConfig.NetworkPluginName)
	master := &OsdnMaster{kClient: kClient, networkClient: networkClient, nodeInformer: kubeInformers.Core().V1().Nodes(), namespaceInformer: kubeInformers.Core().V1().Namespaces(), hostSubnetInformer: networkInformers.Network().V1().HostSubnets(), netNamespaceInformer: networkInformers.Network().V1().NetNamespaces(), subnetAllocatorMap: map[common.ClusterNetwork]*SubnetAllocator{}, hostSubnetNodeIPs: map[ktypes.UID]string{}}
	var err error
	var clusterNetworkEntries []networkapi.ClusterNetworkEntry
	for _, entry := range networkConfig.ClusterNetworks {
		clusterNetworkEntries = append(clusterNetworkEntries, networkapi.ClusterNetworkEntry{CIDR: entry.CIDR, HostSubnetLength: entry.HostSubnetLength})
	}
	master.networkInfo, err = common.ParseNetworkInfo(clusterNetworkEntries, networkConfig.ServiceNetworkCIDR, &networkConfig.VXLANPort)
	if err != nil {
		return err
	}
	if len(clusterNetworkEntries) == 0 {
		panic("No ClusterNetworks set in networkConfig; should have been defaulted in if not configured")
	}
	var parsedClusterNetworkEntries []networkapi.ClusterNetworkEntry
	for _, entry := range master.networkInfo.ClusterNetworks {
		parsedClusterNetworkEntries = append(parsedClusterNetworkEntries, networkapi.ClusterNetworkEntry{CIDR: entry.ClusterCIDR.String(), HostSubnetLength: entry.HostSubnetLength})
	}
	configCN := &networkapi.ClusterNetwork{TypeMeta: metav1.TypeMeta{Kind: "ClusterNetwork"}, ObjectMeta: metav1.ObjectMeta{Name: networkapi.ClusterNetworkDefault}, ClusterNetworks: parsedClusterNetworkEntries, ServiceNetwork: master.networkInfo.ServiceNetwork.String(), PluginName: networkConfig.NetworkPluginName, VXLANPort: &networkConfig.VXLANPort, Network: parsedClusterNetworkEntries[0].CIDR, HostSubnetLength: parsedClusterNetworkEntries[0].HostSubnetLength}
	var getError error
	err = wait.PollImmediate(1*time.Second, time.Minute, func() (bool, error) {
		getError = nil
		existingCN, err := master.networkClient.NetworkV1().ClusterNetworks().Get(networkapi.ClusterNetworkDefault, metav1.GetOptions{})
		if err != nil {
			if !kapierrors.IsNotFound(err) {
				getError = err
				return false, nil
			}
			if err = master.checkClusterNetworkAgainstLocalNetworks(); err != nil {
				return false, err
			}
			if _, err = master.networkClient.NetworkV1().ClusterNetworks().Create(configCN); err != nil {
				return false, err
			}
			klog.Infof("Created ClusterNetwork %s", common.ClusterNetworkToString(configCN))
			if err = master.checkClusterNetworkAgainstClusterObjects(); err != nil {
				utilruntime.HandleError(fmt.Errorf("Cluster contains objects incompatible with new ClusterNetwork: %v", err))
			}
		} else {
			configChanged, err := clusterNetworkChanged(configCN, existingCN)
			if err != nil {
				return false, err
			}
			if configChanged {
				configCN.TypeMeta = existingCN.TypeMeta
				configCN.ObjectMeta = existingCN.ObjectMeta
				if err = master.checkClusterNetworkAgainstClusterObjects(); err != nil {
					utilruntime.HandleError(fmt.Errorf("Attempting to modify cluster to exclude existing objects: %v", err))
					return false, err
				}
				if _, err = master.networkClient.NetworkV1().ClusterNetworks().Update(configCN); err != nil {
					return false, err
				}
				klog.Infof("Updated ClusterNetwork %s", common.ClusterNetworkToString(configCN))
			} else {
				klog.V(5).Infof("No change to ClusterNetwork %s", common.ClusterNetworkToString(configCN))
			}
		}
		return true, nil
	})
	if err != nil {
		if getError != nil {
			return getError
		}
		return err
	}
	master.nodeInformer.Informer().GetController()
	master.namespaceInformer.Informer().GetController()
	master.hostSubnetInformer.Informer().GetController()
	master.netNamespaceInformer.Informer().GetController()
	go master.startSubSystems(networkConfig.NetworkPluginName)
	return nil
}
func (master *OsdnMaster) startSubSystems(pluginName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !cache.WaitForCacheSync(wait.NeverStop, master.nodeInformer.Informer().GetController().HasSynced, master.namespaceInformer.Informer().GetController().HasSynced, master.hostSubnetInformer.Informer().GetController().HasSynced, master.netNamespaceInformer.Informer().GetController().HasSynced) {
		klog.Fatalf("failed to sync SDN master informers")
	}
	if err := master.startSubnetMaster(); err != nil {
		klog.Fatalf("failed to start subnet master: %v", err)
	}
	switch pluginName {
	case network.MultiTenantPluginName:
		master.vnids = newMasterVNIDMap(true)
	case network.NetworkPolicyPluginName:
		master.vnids = newMasterVNIDMap(false)
	}
	if master.vnids != nil {
		if err := master.startVNIDMaster(); err != nil {
			klog.Fatalf("failed to start VNID master: %v", err)
		}
	}
	eim := newEgressIPManager()
	eim.Start(master.networkClient, master.hostSubnetInformer, master.netNamespaceInformer)
}
func (master *OsdnMaster) checkClusterNetworkAgainstLocalNetworks() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	hostIPNets, _, err := netutils.GetHostIPNetworks([]string{tun0})
	if err != nil {
		return err
	}
	return master.networkInfo.CheckHostNetworks(hostIPNets)
}
func (master *OsdnMaster) checkClusterNetworkAgainstClusterObjects() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var subnets []networkapi.HostSubnet
	var pods []kapi.Pod
	var services []kapi.Service
	if subnetList, err := master.networkClient.NetworkV1().HostSubnets().List(metav1.ListOptions{}); err == nil {
		subnets = subnetList.Items
	}
	if podList, err := master.kClient.CoreV1().Pods(metav1.NamespaceAll).List(metav1.ListOptions{}); err == nil {
		pods = podList.Items
	}
	if serviceList, err := master.kClient.CoreV1().Services(metav1.NamespaceAll).List(metav1.ListOptions{}); err == nil {
		services = serviceList.Items
	}
	return master.networkInfo.CheckClusterObjects(subnets, pods, services)
}
func clusterNetworkChanged(obj *networkapi.ClusterNetwork, old *networkapi.ClusterNetwork) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := common.ValidateClusterNetwork(old); err != nil {
		utilruntime.HandleError(fmt.Errorf("Ignoring invalid existing default ClusterNetwork (%v)", err))
		return true, nil
	}
	if old.ServiceNetwork != obj.ServiceNetwork {
		return true, fmt.Errorf("cannot change the serviceNetworkCIDR of an already-deployed cluster")
	} else if old.PluginName != obj.PluginName {
		return true, nil
	} else if len(old.ClusterNetworks) != len(obj.ClusterNetworks) {
		return true, nil
	} else {
		changed := false
		for _, oldCIDR := range old.ClusterNetworks {
			found := false
			for _, newCIDR := range obj.ClusterNetworks {
				if newCIDR.CIDR == oldCIDR.CIDR && newCIDR.HostSubnetLength == oldCIDR.HostSubnetLength {
					found = true
					break
				}
			}
			if !found {
				changed = true
				break
			}
		}
		return changed, nil
	}
}

package node

import (
	"fmt"
	"sync"
	"k8s.io/kubernetes/pkg/apis/core/v1/helper"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	networkv1 "github.com/openshift/api/network/v1"
	"github.com/openshift/origin/pkg/network"
)

type multiTenantPlugin struct {
	node		*OsdnNode
	vnids		*nodeVNIDMap
	vnidInUseLock	sync.Mutex
	vnidInUse	sets.Int
}

func NewMultiTenantPlugin() osdnPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &multiTenantPlugin{}
}
func (mp *multiTenantPlugin) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return network.MultiTenantPluginName
}
func (mp *multiTenantPlugin) SupportsVNIDs() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (mp *multiTenantPlugin) Start(node *OsdnNode) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mp.node = node
	mp.vnidInUse = node.oc.FindPolicyVNIDs()
	mp.vnids = newNodeVNIDMap(mp, node.networkClient)
	if err := mp.vnids.Start(node.networkInformers); err != nil {
		return err
	}
	otx := node.oc.NewTransaction()
	otx.AddFlow("table=80, priority=200, reg0=0, actions=output:NXM_NX_REG2[]")
	otx.AddFlow("table=80, priority=200, reg1=0, actions=output:NXM_NX_REG2[]")
	if err := otx.Commit(); err != nil {
		return err
	}
	return nil
}
func (mp *multiTenantPlugin) updatePodNetwork(namespace string, oldNetID, netID uint32) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pods, err := mp.node.GetLocalPods(namespace)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Could not get list of local pods in namespace %q: %v", namespace, err))
	}
	services, err := mp.node.kClient.CoreV1().Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Could not get list of services in namespace %q: %v", namespace, err))
		services = &corev1.ServiceList{}
	}
	if oldNetID != netID {
		for _, pod := range pods {
			err = mp.node.UpdatePod(pod)
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("Could not update pod %q in namespace %q: %v", pod.Name, namespace, err))
			}
		}
		for _, svc := range services.Items {
			if !helper.IsServiceIPSet(&svc) {
				continue
			}
			mp.node.DeleteServiceRules(&svc)
			mp.node.AddServiceRules(&svc, netID)
		}
		mp.EnsureVNIDRules(netID)
		mp.node.UpdateEgressNetworkPolicyVNID(namespace, oldNetID, netID)
	}
	mp.node.podManager.UpdateLocalMulticastRules(oldNetID)
	mp.node.podManager.UpdateLocalMulticastRules(netID)
}
func (mp *multiTenantPlugin) AddNetNamespace(netns *networkv1.NetNamespace) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mp.updatePodNetwork(netns.Name, 0, netns.NetID)
}
func (mp *multiTenantPlugin) UpdateNetNamespace(netns *networkv1.NetNamespace, oldNetID uint32) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mp.updatePodNetwork(netns.Name, oldNetID, netns.NetID)
}
func (mp *multiTenantPlugin) DeleteNetNamespace(netns *networkv1.NetNamespace) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mp.updatePodNetwork(netns.Name, netns.NetID, 0)
}
func (mp *multiTenantPlugin) GetVNID(namespace string) (uint32, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mp.vnids.WaitAndGetVNID(namespace)
}
func (mp *multiTenantPlugin) GetNamespaces(vnid uint32) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mp.vnids.GetNamespaces(vnid)
}
func (mp *multiTenantPlugin) GetMulticastEnabled(vnid uint32) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mp.vnids.GetMulticastEnabled(vnid)
}
func (mp *multiTenantPlugin) EnsureVNIDRules(vnid uint32) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if vnid == 0 {
		return
	}
	mp.vnidInUseLock.Lock()
	defer mp.vnidInUseLock.Unlock()
	if mp.vnidInUse.Has(int(vnid)) {
		return
	}
	mp.vnidInUse.Insert(int(vnid))
	klog.V(5).Infof("EnsureVNIDRules %d - adding rules", vnid)
	otx := mp.node.oc.NewTransaction()
	otx.AddFlow("table=80, priority=100, reg0=%d, reg1=%d, actions=output:NXM_NX_REG2[]", vnid, vnid)
	if err := otx.Commit(); err != nil {
		utilruntime.HandleError(fmt.Errorf("Error adding OVS flow for VNID: %v", err))
	}
}
func (mp *multiTenantPlugin) SyncVNIDRules() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mp.vnidInUseLock.Lock()
	defer mp.vnidInUseLock.Unlock()
	unused := mp.node.oc.FindUnusedVNIDs()
	klog.Infof("SyncVNIDRules: %d unused VNIDs", len(unused))
	otx := mp.node.oc.NewTransaction()
	for _, vnid := range unused {
		mp.vnidInUse.Delete(int(vnid))
		otx.DeleteFlows("table=80, reg1=%d", vnid)
	}
	if err := otx.Commit(); err != nil {
		utilruntime.HandleError(fmt.Errorf("Error deleting syncing OVS VNID rules: %v", err))
	}
}

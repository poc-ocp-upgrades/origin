package node

import (
	"fmt"
	"sync"
	"time"
	"k8s.io/klog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	networkapi "github.com/openshift/api/network/v1"
	networkclient "github.com/openshift/client-go/network/clientset/versioned"
	networkinformers "github.com/openshift/client-go/network/informers/externalversions"
	"github.com/openshift/origin/pkg/network/common"
)

type nodeVNIDMap struct {
	policy			osdnPolicy
	networkClient		networkclient.Interface
	networkInformers	networkinformers.SharedInformerFactory
	lock			sync.Mutex
	ids			map[string]uint32
	mcEnabled		map[string]bool
	namespaces		map[uint32]sets.String
}

func newNodeVNIDMap(policy osdnPolicy, networkClient networkclient.Interface) *nodeVNIDMap {
	_logClusterCodePath()
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
	return &nodeVNIDMap{policy: policy, networkClient: networkClient, ids: make(map[string]uint32), mcEnabled: make(map[string]bool), namespaces: make(map[uint32]sets.String)}
}
func (vmap *nodeVNIDMap) addNamespaceToSet(name string, vnid uint32) {
	_logClusterCodePath()
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
	set, found := vmap.namespaces[vnid]
	if !found {
		set = sets.NewString()
		vmap.namespaces[vnid] = set
	}
	set.Insert(name)
}
func (vmap *nodeVNIDMap) removeNamespaceFromSet(name string, vnid uint32) {
	_logClusterCodePath()
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
	if set, found := vmap.namespaces[vnid]; found {
		set.Delete(name)
		if set.Len() == 0 {
			delete(vmap.namespaces, vnid)
		}
	}
}
func (vmap *nodeVNIDMap) GetNamespaces(id uint32) []string {
	_logClusterCodePath()
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
	vmap.lock.Lock()
	defer vmap.lock.Unlock()
	if set, ok := vmap.namespaces[id]; ok {
		return set.List()
	} else {
		return nil
	}
}
func (vmap *nodeVNIDMap) GetMulticastEnabled(id uint32) bool {
	_logClusterCodePath()
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
	vmap.lock.Lock()
	defer vmap.lock.Unlock()
	set, exists := vmap.namespaces[id]
	if !exists || set.Len() == 0 {
		return false
	}
	for _, ns := range set.List() {
		if !vmap.mcEnabled[ns] {
			return false
		}
	}
	return true
}
func (vmap *nodeVNIDMap) WaitAndGetVNID(name string) (uint32, error) {
	_logClusterCodePath()
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
	var id uint32
	backoff := utilwait.Backoff{Duration: 400 * time.Millisecond, Factor: 1.5, Steps: 6}
	err := utilwait.ExponentialBackoff(backoff, func() (bool, error) {
		var err error
		id, err = vmap.getVNID(name)
		return err == nil, nil
	})
	if err == nil {
		return id, nil
	} else {
		VnidNotFoundErrors.Inc()
		netns, err := vmap.networkClient.NetworkV1().NetNamespaces().Get(name, metav1.GetOptions{})
		if err != nil {
			return 0, fmt.Errorf("failed to find netid for namespace: %s, %v", name, err)
		}
		klog.Warningf("Netid for namespace: %s exists but not found in vnid map", name)
		vmap.setVNID(netns.Name, netns.NetID, netnsIsMulticastEnabled(netns))
		return netns.NetID, nil
	}
}
func (vmap *nodeVNIDMap) getVNID(name string) (uint32, error) {
	_logClusterCodePath()
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
	vmap.lock.Lock()
	defer vmap.lock.Unlock()
	if id, ok := vmap.ids[name]; ok {
		return id, nil
	}
	return 0, fmt.Errorf("failed to find netid for namespace: %s in vnid map", name)
}
func (vmap *nodeVNIDMap) setVNID(name string, id uint32, mcEnabled bool) {
	_logClusterCodePath()
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
	vmap.lock.Lock()
	defer vmap.lock.Unlock()
	if oldId, found := vmap.ids[name]; found {
		vmap.removeNamespaceFromSet(name, oldId)
	}
	vmap.ids[name] = id
	vmap.mcEnabled[name] = mcEnabled
	vmap.addNamespaceToSet(name, id)
	klog.Infof("Associate netid %d to namespace %q with mcEnabled %v", id, name, mcEnabled)
}
func (vmap *nodeVNIDMap) unsetVNID(name string) (id uint32, err error) {
	_logClusterCodePath()
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
	vmap.lock.Lock()
	defer vmap.lock.Unlock()
	id, found := vmap.ids[name]
	if !found {
		return 0, fmt.Errorf("failed to find netid for namespace: %s in vnid map", name)
	}
	vmap.removeNamespaceFromSet(name, id)
	delete(vmap.ids, name)
	delete(vmap.mcEnabled, name)
	klog.Infof("Dissociate netid %d from namespace %q", id, name)
	return id, nil
}
func netnsIsMulticastEnabled(netns *networkapi.NetNamespace) bool {
	_logClusterCodePath()
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
	enabled, ok := netns.Annotations[networkapi.MulticastEnabledAnnotation]
	return enabled == "true" && ok
}
func (vmap *nodeVNIDMap) populateVNIDs() error {
	_logClusterCodePath()
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
	nets, err := vmap.networkClient.NetworkV1().NetNamespaces().List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, net := range nets.Items {
		vmap.setVNID(net.Name, net.NetID, netnsIsMulticastEnabled(&net))
	}
	return nil
}
func (vmap *nodeVNIDMap) Start(networkInformers networkinformers.SharedInformerFactory) error {
	_logClusterCodePath()
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
	vmap.networkInformers = networkInformers
	err := vmap.populateVNIDs()
	if err != nil {
		return err
	}
	vmap.watchNetNamespaces()
	return nil
}
func (vmap *nodeVNIDMap) watchNetNamespaces() {
	_logClusterCodePath()
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
	funcs := common.InformerFuncs(&networkapi.NetNamespace{}, vmap.handleAddOrUpdateNetNamespace, vmap.handleDeleteNetNamespace)
	vmap.networkInformers.Network().V1().NetNamespaces().Informer().AddEventHandler(funcs)
}
func (vmap *nodeVNIDMap) handleAddOrUpdateNetNamespace(obj, _ interface{}, eventType watch.EventType) {
	_logClusterCodePath()
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
	netns := obj.(*networkapi.NetNamespace)
	klog.V(5).Infof("Watch %s event for NetNamespace %q", eventType, netns.Name)
	oldNetID, err := vmap.getVNID(netns.NetName)
	oldMCEnabled := vmap.mcEnabled[netns.NetName]
	mcEnabled := netnsIsMulticastEnabled(netns)
	if err == nil && oldNetID == netns.NetID && oldMCEnabled == mcEnabled {
		return
	}
	vmap.setVNID(netns.NetName, netns.NetID, mcEnabled)
	if eventType == watch.Added {
		vmap.policy.AddNetNamespace(netns)
	} else {
		vmap.policy.UpdateNetNamespace(netns, oldNetID)
	}
}
func (vmap *nodeVNIDMap) handleDeleteNetNamespace(obj interface{}) {
	_logClusterCodePath()
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
	netns := obj.(*networkapi.NetNamespace)
	klog.V(5).Infof("Watch %s event for NetNamespace %q", watch.Deleted, netns.Name)
	vmap.unsetVNID(netns.NetName)
	vmap.policy.DeleteNetNamespace(netns)
}

package master

import (
	"fmt"
	networkv1 "github.com/openshift/api/network/v1"
	networkclient "github.com/openshift/client-go/network/clientset/versioned"
	"github.com/openshift/library-go/pkg/network/networkapihelpers"
	"github.com/openshift/origin/pkg/network"
	"github.com/openshift/origin/pkg/network/common"
	pnetid "github.com/openshift/origin/pkg/network/master/netid"
	kapi "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog"
	"sync"
)

type masterVNIDMap struct {
	lock             sync.Mutex
	ids              map[string]uint32
	netIDManager     *pnetid.Allocator
	adminNamespaces  sets.String
	allowRenumbering bool
}

func newMasterVNIDMap(allowRenumbering bool) *masterVNIDMap {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	netIDRange, err := pnetid.NewNetIDRange(network.MinVNID, network.MaxVNID)
	if err != nil {
		panic(err)
	}
	return &masterVNIDMap{netIDManager: pnetid.NewInMemory(netIDRange), adminNamespaces: sets.NewString(metav1.NamespaceDefault), ids: make(map[string]uint32), allowRenumbering: allowRenumbering}
}
func (vmap *masterVNIDMap) getVNID(name string) (uint32, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	id, found := vmap.ids[name]
	return id, found
}
func (vmap *masterVNIDMap) setVNID(name string, id uint32) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmap.ids[name] = id
}
func (vmap *masterVNIDMap) unsetVNID(name string) (uint32, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	id, found := vmap.ids[name]
	delete(vmap.ids, name)
	return id, found
}
func (vmap *masterVNIDMap) getVNIDCount(id uint32) int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	count := 0
	for _, netid := range vmap.ids {
		if id == netid {
			count = count + 1
		}
	}
	return count
}
func (vmap *masterVNIDMap) isAdminNamespace(nsName string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if vmap.adminNamespaces.Has(nsName) {
		return true
	}
	return false
}
func (vmap *masterVNIDMap) markAllocatedNetID(netid uint32) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if netid == network.GlobalVNID {
		return nil
	}
	switch err := vmap.netIDManager.Allocate(netid); err {
	case nil:
	case pnetid.ErrAllocated:
	default:
		return fmt.Errorf("unable to allocate netid %d: %v", netid, err)
	}
	return nil
}
func (vmap *masterVNIDMap) allocateNetID(nsName string) (uint32, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	exists := false
	if netid, found := vmap.getVNID(nsName); found {
		exists = true
		return netid, exists, nil
	}
	var netid uint32
	if vmap.isAdminNamespace(nsName) {
		netid = network.GlobalVNID
	} else {
		var err error
		netid, err = vmap.netIDManager.AllocateNext()
		if err != nil {
			return 0, exists, err
		}
	}
	vmap.setVNID(nsName, netid)
	klog.Infof("Allocated netid %d for namespace %q", netid, nsName)
	return netid, exists, nil
}
func (vmap *masterVNIDMap) releaseNetID(nsName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	netid, found := vmap.unsetVNID(nsName)
	if !found {
		return fmt.Errorf("netid not found for namespace %q", nsName)
	}
	if netid == network.GlobalVNID {
		return nil
	}
	if count := vmap.getVNIDCount(netid); count == 0 {
		if err := vmap.netIDManager.Release(netid); err != nil {
			return fmt.Errorf("error while releasing netid %d for namespace %q, %v", netid, nsName, err)
		}
		klog.Infof("Released netid %d for namespace %q", netid, nsName)
	} else {
		klog.V(5).Infof("netid %d for namespace %q is still in use", netid, nsName)
	}
	return nil
}
func (vmap *masterVNIDMap) updateNetID(nsName string, action networkapihelpers.PodNetworkAction, args string) (uint32, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var netid uint32
	allocated := false
	oldnetid, found := vmap.getVNID(nsName)
	if !found {
		return 0, fmt.Errorf("netid not found for namespace %q", nsName)
	}
	switch action {
	case networkapihelpers.GlobalPodNetwork:
		netid = network.GlobalVNID
	case networkapihelpers.JoinPodNetwork:
		joinNsName := args
		var found bool
		if netid, found = vmap.getVNID(joinNsName); !found {
			return 0, fmt.Errorf("netid not found for namespace %q", joinNsName)
		}
	case networkapihelpers.IsolatePodNetwork:
		if nsName == kapi.NamespaceDefault {
			return 0, fmt.Errorf("network isolation for namespace %q is not allowed", nsName)
		}
		if count := vmap.getVNIDCount(oldnetid); count == 1 {
			return oldnetid, nil
		}
		var err error
		netid, err = vmap.netIDManager.AllocateNext()
		if err != nil {
			return 0, err
		}
		allocated = true
	default:
		return 0, fmt.Errorf("invalid pod network action: %v", action)
	}
	if err := vmap.releaseNetID(nsName); err != nil {
		if allocated {
			vmap.netIDManager.Release(netid)
		}
		return 0, err
	}
	vmap.setVNID(nsName, netid)
	klog.Infof("Updated netid %d for namespace %q", netid, nsName)
	return netid, nil
}
func (vmap *masterVNIDMap) assignVNID(networkClient networkclient.Interface, nsName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmap.lock.Lock()
	defer vmap.lock.Unlock()
	netid, exists, err := vmap.allocateNetID(nsName)
	if err != nil {
		return err
	}
	if !exists {
		netns := &networkv1.NetNamespace{TypeMeta: metav1.TypeMeta{Kind: "NetNamespace"}, ObjectMeta: metav1.ObjectMeta{Name: nsName}, NetName: nsName, NetID: netid}
		if _, err := networkClient.NetworkV1().NetNamespaces().Create(netns); err != nil {
			if er := vmap.releaseNetID(nsName); er != nil {
				utilruntime.HandleError(er)
			}
			return err
		}
	}
	return nil
}
func (vmap *masterVNIDMap) revokeVNID(networkClient networkclient.Interface, nsName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vmap.lock.Lock()
	defer vmap.lock.Unlock()
	if err := networkClient.NetworkV1().NetNamespaces().Delete(nsName, &metav1.DeleteOptions{}); err != nil {
		return err
	}
	if err := vmap.releaseNetID(nsName); err != nil {
		return err
	}
	return nil
}
func (vmap *masterVNIDMap) updateVNID(networkClient networkclient.Interface, origNetns *networkv1.NetNamespace) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	netns := origNetns.DeepCopy()
	action, args, err := networkapihelpers.GetChangePodNetworkAnnotation(netns)
	if err == networkapihelpers.ErrorPodNetworkAnnotationNotFound {
		return nil
	} else if !vmap.allowRenumbering {
		networkapihelpers.DeleteChangePodNetworkAnnotation(netns)
		_, _ = networkClient.NetworkV1().NetNamespaces().Update(netns)
		return fmt.Errorf("network plugin does not allow NetNamespace renumbering")
	}
	vmap.lock.Lock()
	defer vmap.lock.Unlock()
	netid, err := vmap.updateNetID(netns.NetName, action, args)
	if err != nil {
		return err
	}
	netns.NetID = netid
	networkapihelpers.DeleteChangePodNetworkAnnotation(netns)
	if _, err := networkClient.NetworkV1().NetNamespaces().Update(netns); err != nil {
		return err
	}
	return nil
}
func (master *OsdnMaster) startVNIDMaster() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := master.initNetIDAllocator(); err != nil {
		return err
	}
	master.watchNamespaces()
	master.watchNetNamespaces()
	return nil
}
func (master *OsdnMaster) initNetIDAllocator() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	netnsList, err := master.networkClient.NetworkV1().NetNamespaces().List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, netns := range netnsList.Items {
		if err := master.vnids.markAllocatedNetID(netns.NetID); err != nil {
			utilruntime.HandleError(err)
		}
		master.vnids.setVNID(netns.Name, netns.NetID)
	}
	return nil
}
func (master *OsdnMaster) watchNamespaces() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	funcs := common.InformerFuncs(&kapi.Namespace{}, master.handleAddOrUpdateNamespace, master.handleDeleteNamespace)
	master.namespaceInformer.Informer().AddEventHandler(funcs)
}
func (master *OsdnMaster) handleAddOrUpdateNamespace(obj, _ interface{}, eventType watch.EventType) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ns := obj.(*kapi.Namespace)
	klog.V(5).Infof("Watch %s event for Namespace %q", eventType, ns.Name)
	if err := master.vnids.assignVNID(master.networkClient, ns.Name); err != nil {
		utilruntime.HandleError(fmt.Errorf("Error assigning netid: %v", err))
	}
}
func (master *OsdnMaster) handleDeleteNamespace(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ns := obj.(*kapi.Namespace)
	klog.V(5).Infof("Watch %s event for Namespace %q", watch.Deleted, ns.Name)
	if err := master.vnids.revokeVNID(master.networkClient, ns.Name); err != nil {
		utilruntime.HandleError(fmt.Errorf("Error revoking netid: %v", err))
	}
}
func (master *OsdnMaster) watchNetNamespaces() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	funcs := common.InformerFuncs(&networkv1.NetNamespace{}, master.handleAddOrUpdateNetNamespace, nil)
	master.netNamespaceInformer.Informer().AddEventHandler(funcs)
}
func (master *OsdnMaster) handleAddOrUpdateNetNamespace(obj, _ interface{}, eventType watch.EventType) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	netns := obj.(*networkv1.NetNamespace)
	klog.V(5).Infof("Watch %s event for NetNamespace %q", eventType, netns.Name)
	if err := master.vnids.updateVNID(master.networkClient, netns); err != nil {
		utilruntime.HandleError(fmt.Errorf("Error updating netid: %v", err))
	}
}

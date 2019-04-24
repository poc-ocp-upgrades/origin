package master

import (
	"fmt"
	"strconv"
	"k8s.io/klog"
	kapi "k8s.io/api/core/v1"
	kerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/util/retry"
	networkapi "github.com/openshift/api/network/v1"
	"github.com/openshift/origin/pkg/network"
	"github.com/openshift/origin/pkg/network/common"
)

func (master *OsdnMaster) startSubnetMaster() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := master.initSubnetAllocators(); err != nil {
		return err
	}
	master.watchNodes()
	master.watchSubnets()
	return nil
}
func (master *OsdnMaster) watchNodes() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	funcs := common.InformerFuncs(&kapi.Node{}, master.handleAddOrUpdateNode, master.handleDeleteNode)
	master.nodeInformer.Informer().AddEventHandler(funcs)
}
func (master *OsdnMaster) handleAddOrUpdateNode(obj, _ interface{}, eventType watch.EventType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := obj.(*kapi.Node)
	nodeIP := getNodeInternalIP(node)
	if len(nodeIP) == 0 {
		utilruntime.HandleError(fmt.Errorf("Node IP is not set for node %s, skipping %s event, node: %v", node.Name, eventType, node))
		return
	}
	if oldNodeIP, ok := master.hostSubnetNodeIPs[node.UID]; ok && (nodeIP == oldNodeIP) {
		return
	}
	klog.V(5).Infof("Watch %s event for Node %q", eventType, node.Name)
	master.clearInitialNodeNetworkUnavailableCondition(node)
	usedNodeIP, err := master.addNode(node.Name, string(node.UID), nodeIP, nil)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Error creating subnet for node %s, ip %s: %v", node.Name, nodeIP, err))
		return
	}
	master.hostSubnetNodeIPs[node.UID] = usedNodeIP
}
func (master *OsdnMaster) handleDeleteNode(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := obj.(*kapi.Node)
	klog.V(5).Infof("Watch %s event for Node %q", watch.Deleted, node.Name)
	if _, exists := master.hostSubnetNodeIPs[node.UID]; !exists {
		return
	}
	delete(master.hostSubnetNodeIPs, node.UID)
	if err := master.deleteNode(node.Name); err != nil {
		utilruntime.HandleError(fmt.Errorf("Error deleting node %s: %v", node.Name, err))
		return
	}
}
func (master *OsdnMaster) addNode(nodeName string, nodeUID string, nodeIP string, hsAnnotations map[string]string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := master.networkInfo.ValidateNodeIP(nodeIP); err != nil {
		return "", err
	}
	sub, err := master.networkClient.NetworkV1().HostSubnets().Get(nodeName, metav1.GetOptions{})
	if err == nil {
		if err = common.ValidateHostSubnet(sub); err != nil {
			utilruntime.HandleError(fmt.Errorf("Deleting invalid HostSubnet %q: %v", nodeName, err))
			_ = master.networkClient.NetworkV1().HostSubnets().Delete(nodeName, &metav1.DeleteOptions{})
		} else if sub.HostIP == nodeIP {
			return nodeIP, nil
		} else {
			sub.HostIP = nodeIP
			sub, err = master.networkClient.NetworkV1().HostSubnets().Update(sub)
			if err != nil {
				return "", fmt.Errorf("error updating subnet %s for node %s: %v", sub.Subnet, nodeName, err)
			}
			klog.Infof("Updated HostSubnet %s", common.HostSubnetToString(sub))
			return nodeIP, nil
		}
	}
	if len(nodeUID) != 0 {
		if hsAnnotations == nil {
			hsAnnotations = make(map[string]string)
		}
		hsAnnotations[networkapi.NodeUIDAnnotation] = nodeUID
	}
	network, err := master.allocateNetwork(nodeName)
	if err != nil {
		return "", err
	}
	sub = &networkapi.HostSubnet{TypeMeta: metav1.TypeMeta{Kind: "HostSubnet"}, ObjectMeta: metav1.ObjectMeta{Name: nodeName, Annotations: hsAnnotations}, Host: nodeName, HostIP: nodeIP, Subnet: network}
	sub, err = master.networkClient.NetworkV1().HostSubnets().Create(sub)
	if err != nil {
		if er := master.releaseNetwork(network); er != nil {
			utilruntime.HandleError(er)
		}
		return "", fmt.Errorf("error allocating subnet for node %q: %v", nodeName, err)
	}
	klog.Infof("Created HostSubnet %s", common.HostSubnetToString(sub))
	return nodeIP, nil
}
func (master *OsdnMaster) deleteNode(nodeName string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	subInfo := nodeName
	if sub, err := master.hostSubnetInformer.Lister().Get(nodeName); err == nil {
		subInfo = common.HostSubnetToString(sub)
	}
	if err := master.networkClient.NetworkV1().HostSubnets().Delete(nodeName, &metav1.DeleteOptions{}); err != nil {
		return fmt.Errorf("error deleting subnet for node %q: %v", nodeName, err)
	}
	klog.Infof("Deleted HostSubnet %s", subInfo)
	return nil
}
func (master *OsdnMaster) clearInitialNodeNetworkUnavailableCondition(origNode *kapi.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := origNode.DeepCopy()
	knode := node
	cleared := false
	resultErr := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		var err error
		if knode != node {
			knode, err = master.nodeInformer.Lister().Get(node.Name)
			if err != nil {
				return err
			}
		}
		for i := range knode.Status.Conditions {
			if knode.Status.Conditions[i].Type == kapi.NodeNetworkUnavailable {
				condition := &knode.Status.Conditions[i]
				if condition.Status != kapi.ConditionFalse && condition.Reason == "NoRouteCreated" {
					condition.Status = kapi.ConditionFalse
					condition.Reason = "RouteCreated"
					condition.Message = "openshift-sdn cleared kubelet-set NoRouteCreated"
					condition.LastTransitionTime = metav1.Now()
					if knode, err = master.kClient.CoreV1().Nodes().UpdateStatus(knode); err == nil {
						cleared = true
					}
				}
				break
			}
		}
		return err
	})
	if resultErr != nil {
		utilruntime.HandleError(fmt.Errorf("status update failed for local node: %v", resultErr))
	} else if cleared {
		klog.Infof("Cleared node NetworkUnavailable/NoRouteCreated condition for %s", node.Name)
	}
}
func getNodeInternalIP(node *kapi.Node) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var nodeIP string
	for _, addr := range node.Status.Addresses {
		if addr.Type == kapi.NodeInternalIP {
			nodeIP = addr.Address
			break
		}
	}
	return nodeIP
}
func (master *OsdnMaster) watchSubnets() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	funcs := common.InformerFuncs(&networkapi.HostSubnet{}, master.handleAddOrUpdateSubnet, master.handleDeleteSubnet)
	master.hostSubnetInformer.Informer().AddEventHandler(funcs)
}
func (master *OsdnMaster) handleAddOrUpdateSubnet(obj, _ interface{}, eventType watch.EventType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hs := obj.(*networkapi.HostSubnet)
	klog.V(5).Infof("Watch %s event for HostSubnet %q", eventType, hs.Name)
	if err := common.ValidateHostSubnet(hs); err != nil {
		utilruntime.HandleError(fmt.Errorf("Ignoring invalid HostSubnet %s: %v", common.HostSubnetToString(hs), err))
		return
	}
	if err := master.reconcileHostSubnet(hs); err != nil {
		utilruntime.HandleError(err)
	}
	if err := master.networkInfo.ValidateNodeIP(hs.HostIP); err != nil {
		utilruntime.HandleError(fmt.Errorf("Failed to validate HostSubnet %s: %v", common.HostSubnetToString(hs), err))
	}
	if _, ok := hs.Annotations[networkapi.AssignHostSubnetAnnotation]; ok {
		if err := master.handleAssignHostSubnetAnnotation(hs); err != nil {
			utilruntime.HandleError(err)
			return
		}
	}
}
func (master *OsdnMaster) handleDeleteSubnet(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hs := obj.(*networkapi.HostSubnet)
	klog.V(5).Infof("Watch %s event for HostSubnet %q", watch.Deleted, hs.Name)
	if _, ok := hs.Annotations[networkapi.AssignHostSubnetAnnotation]; ok {
		return
	}
	if err := master.releaseNetwork(hs.Subnet); err != nil {
		utilruntime.HandleError(err)
	}
}
func (master *OsdnMaster) reconcileHostSubnet(subnet *networkapi.HostSubnet) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var node *kapi.Node
	var err error
	node, err = master.nodeInformer.Lister().Get(subnet.Name)
	if err != nil {
		node, err = master.kClient.CoreV1().Nodes().Get(subnet.Name, metav1.GetOptions{})
		if err != nil {
			if kerrs.IsNotFound(err) {
				node = nil
			} else {
				return fmt.Errorf("error fetching node for subnet %q: %v", subnet.Name, err)
			}
		}
	}
	if node == nil && len(subnet.Annotations[networkapi.NodeUIDAnnotation]) == 0 {
		return nil
	} else if node != nil && len(subnet.Annotations[networkapi.NodeUIDAnnotation]) == 0 {
		sn := subnet.DeepCopy()
		if sn.Annotations == nil {
			sn.Annotations = make(map[string]string)
		}
		sn.Annotations[networkapi.NodeUIDAnnotation] = string(node.UID)
		if _, err = master.networkClient.NetworkV1().HostSubnets().Update(sn); err != nil {
			return fmt.Errorf("error updating subnet %v for node %s: %v", sn, sn.Name, err)
		}
	} else if node == nil && len(subnet.Annotations[networkapi.NodeUIDAnnotation]) > 0 {
		klog.Infof("Setup found no node associated with hostsubnet %s, deleting the hostsubnet", subnet.Name)
		if err = master.networkClient.NetworkV1().HostSubnets().Delete(subnet.Name, &metav1.DeleteOptions{}); err != nil {
			return fmt.Errorf("error deleting subnet %v: %v", subnet, err)
		}
	} else if string(node.UID) != subnet.Annotations[networkapi.NodeUIDAnnotation] {
		klog.Infof("Missed node event, hostsubnet %s has the UID of an incorrect object, deleting the hostsubnet", subnet.Name)
		if err = master.networkClient.NetworkV1().HostSubnets().Delete(subnet.Name, &metav1.DeleteOptions{}); err != nil {
			return fmt.Errorf("error deleting subnet %v: %v", subnet, err)
		}
	}
	return nil
}
func (master *OsdnMaster) handleAssignHostSubnetAnnotation(hs *networkapi.HostSubnet) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := master.networkClient.NetworkV1().HostSubnets().Delete(hs.Name, &metav1.DeleteOptions{}); err != nil {
		return fmt.Errorf("error in deleting annotated subnet: %s, %v", hs.Name, err)
	}
	klog.Infof("Deleted HostSubnet not backed by node: %s", common.HostSubnetToString(hs))
	var hsAnnotations map[string]string
	if vnid, ok := hs.Annotations[networkapi.FixedVNIDHostAnnotation]; ok {
		vnidInt, err := strconv.Atoi(vnid)
		if err == nil && vnidInt >= 0 && uint32(vnidInt) <= network.MaxVNID {
			hsAnnotations = make(map[string]string)
			hsAnnotations[networkapi.FixedVNIDHostAnnotation] = strconv.Itoa(vnidInt)
		} else {
			utilruntime.HandleError(fmt.Errorf("VNID %s is an invalid value for annotation %s. Annotation will be ignored.", vnid, networkapi.FixedVNIDHostAnnotation))
		}
	}
	if _, err := master.addNode(hs.Name, "", hs.HostIP, hsAnnotations); err != nil {
		return fmt.Errorf("error creating subnet: %s, %v", hs.Name, err)
	}
	klog.Infof("Created HostSubnet not backed by node: %s", common.HostSubnetToString(hs))
	return nil
}

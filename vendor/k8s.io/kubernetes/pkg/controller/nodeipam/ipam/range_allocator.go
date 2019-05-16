package ipam

import (
	"fmt"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	informers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/nodeipam/ipam/cidrset"
	nodeutil "k8s.io/kubernetes/pkg/controller/util/node"
	utilnode "k8s.io/kubernetes/pkg/util/node"
	"net"
	"sync"
)

type rangeAllocator struct {
	client                clientset.Interface
	cidrs                 *cidrset.CidrSet
	clusterCIDR           *net.IPNet
	maxCIDRs              int
	nodeLister            corelisters.NodeLister
	nodesSynced           cache.InformerSynced
	nodeCIDRUpdateChannel chan nodeAndCIDR
	recorder              record.EventRecorder
	lock                  sync.Mutex
	nodesInProcessing     sets.String
}

func NewCIDRRangeAllocator(client clientset.Interface, nodeInformer informers.NodeInformer, clusterCIDR *net.IPNet, serviceCIDR *net.IPNet, subNetMaskSize int, nodeList *v1.NodeList) (CIDRAllocator, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if client == nil {
		klog.Fatalf("kubeClient is nil when starting NodeController")
	}
	eventBroadcaster := record.NewBroadcaster()
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "cidrAllocator"})
	eventBroadcaster.StartLogging(klog.Infof)
	klog.V(0).Infof("Sending events to api server.")
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: client.CoreV1().Events("")})
	set, err := cidrset.NewCIDRSet(clusterCIDR, subNetMaskSize)
	if err != nil {
		return nil, err
	}
	ra := &rangeAllocator{client: client, cidrs: set, clusterCIDR: clusterCIDR, nodeLister: nodeInformer.Lister(), nodesSynced: nodeInformer.Informer().HasSynced, nodeCIDRUpdateChannel: make(chan nodeAndCIDR, cidrUpdateQueueSize), recorder: recorder, nodesInProcessing: sets.NewString()}
	if serviceCIDR != nil {
		ra.filterOutServiceRange(serviceCIDR)
	} else {
		klog.V(0).Info("No Service CIDR provided. Skipping filtering out service addresses.")
	}
	if nodeList != nil {
		for _, node := range nodeList.Items {
			if node.Spec.PodCIDR == "" {
				klog.Infof("Node %v has no CIDR, ignoring", node.Name)
				continue
			} else {
				klog.Infof("Node %v has CIDR %s, occupying it in CIDR map", node.Name, node.Spec.PodCIDR)
			}
			if err := ra.occupyCIDR(&node); err != nil {
				return nil, err
			}
		}
	}
	nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: nodeutil.CreateAddNodeHandler(ra.AllocateOrOccupyCIDR), UpdateFunc: nodeutil.CreateUpdateNodeHandler(func(_, newNode *v1.Node) error {
		if newNode.Spec.PodCIDR == "" {
			return ra.AllocateOrOccupyCIDR(newNode)
		}
		return nil
	}), DeleteFunc: nodeutil.CreateDeleteNodeHandler(ra.ReleaseCIDR)})
	return ra, nil
}
func (r *rangeAllocator) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	klog.Infof("Starting range CIDR allocator")
	defer klog.Infof("Shutting down range CIDR allocator")
	if !controller.WaitForCacheSync("cidrallocator", stopCh, r.nodesSynced) {
		return
	}
	for i := 0; i < cidrUpdateWorkers; i++ {
		go r.worker(stopCh)
	}
	<-stopCh
}
func (r *rangeAllocator) worker(stopChan <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for {
		select {
		case workItem, ok := <-r.nodeCIDRUpdateChannel:
			if !ok {
				klog.Warning("Channel nodeCIDRUpdateChannel was unexpectedly closed")
				return
			}
			if err := r.updateCIDRAllocation(workItem); err != nil {
				r.nodeCIDRUpdateChannel <- workItem
			}
		case <-stopChan:
			return
		}
	}
}
func (r *rangeAllocator) insertNodeToProcessing(nodeName string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.nodesInProcessing.Has(nodeName) {
		return false
	}
	r.nodesInProcessing.Insert(nodeName)
	return true
}
func (r *rangeAllocator) removeNodeFromProcessing(nodeName string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.lock.Lock()
	defer r.lock.Unlock()
	r.nodesInProcessing.Delete(nodeName)
}
func (r *rangeAllocator) occupyCIDR(node *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer r.removeNodeFromProcessing(node.Name)
	if node.Spec.PodCIDR == "" {
		return nil
	}
	_, podCIDR, err := net.ParseCIDR(node.Spec.PodCIDR)
	if err != nil {
		return fmt.Errorf("failed to parse node %s, CIDR %s", node.Name, node.Spec.PodCIDR)
	}
	if err := r.cidrs.Occupy(podCIDR); err != nil {
		return fmt.Errorf("failed to mark cidr as occupied: %v", err)
	}
	return nil
}
func (r *rangeAllocator) AllocateOrOccupyCIDR(node *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if node == nil {
		return nil
	}
	if !r.insertNodeToProcessing(node.Name) {
		klog.V(2).Infof("Node %v is already in a process of CIDR assignment.", node.Name)
		return nil
	}
	if node.Spec.PodCIDR != "" {
		return r.occupyCIDR(node)
	}
	podCIDR, err := r.cidrs.AllocateNext()
	if err != nil {
		r.removeNodeFromProcessing(node.Name)
		nodeutil.RecordNodeStatusChange(r.recorder, node, "CIDRNotAvailable")
		return fmt.Errorf("failed to allocate cidr: %v", err)
	}
	klog.V(4).Infof("Putting node %s with CIDR %s into the work queue", node.Name, podCIDR)
	r.nodeCIDRUpdateChannel <- nodeAndCIDR{nodeName: node.Name, cidr: podCIDR}
	return nil
}
func (r *rangeAllocator) ReleaseCIDR(node *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if node == nil || node.Spec.PodCIDR == "" {
		return nil
	}
	_, podCIDR, err := net.ParseCIDR(node.Spec.PodCIDR)
	if err != nil {
		return fmt.Errorf("Failed to parse CIDR %s on Node %v: %v", node.Spec.PodCIDR, node.Name, err)
	}
	klog.V(4).Infof("release CIDR %s", node.Spec.PodCIDR)
	if err = r.cidrs.Release(podCIDR); err != nil {
		return fmt.Errorf("Error when releasing CIDR %v: %v", node.Spec.PodCIDR, err)
	}
	return err
}
func (r *rangeAllocator) filterOutServiceRange(serviceCIDR *net.IPNet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !r.clusterCIDR.Contains(serviceCIDR.IP.Mask(r.clusterCIDR.Mask)) && !serviceCIDR.Contains(r.clusterCIDR.IP.Mask(serviceCIDR.Mask)) {
		return
	}
	if err := r.cidrs.Occupy(serviceCIDR); err != nil {
		klog.Errorf("Error filtering out service cidr %v: %v", serviceCIDR, err)
	}
}
func (r *rangeAllocator) updateCIDRAllocation(data nodeAndCIDR) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	var node *v1.Node
	defer r.removeNodeFromProcessing(data.nodeName)
	podCIDR := data.cidr.String()
	node, err = r.nodeLister.Get(data.nodeName)
	if err != nil {
		klog.Errorf("Failed while getting node %v for updating Node.Spec.PodCIDR: %v", data.nodeName, err)
		return err
	}
	if node.Spec.PodCIDR == podCIDR {
		klog.V(4).Infof("Node %v already has allocated CIDR %v. It matches the proposed one.", node.Name, podCIDR)
		return nil
	}
	if node.Spec.PodCIDR != "" {
		klog.Errorf("Node %v already has a CIDR allocated %v. Releasing the new one %v.", node.Name, node.Spec.PodCIDR, podCIDR)
		if err := r.cidrs.Release(data.cidr); err != nil {
			klog.Errorf("Error when releasing CIDR %v", podCIDR)
		}
		return nil
	}
	for i := 0; i < cidrUpdateRetries; i++ {
		if err = utilnode.PatchNodeCIDR(r.client, types.NodeName(node.Name), podCIDR); err == nil {
			klog.Infof("Set node %v PodCIDR to %v", node.Name, podCIDR)
			return nil
		}
	}
	klog.Errorf("Failed to update node %v PodCIDR to %v after multiple attempts: %v", node.Name, podCIDR, err)
	nodeutil.RecordNodeStatusChange(r.recorder, node, "CIDRAssignmentFailed")
	if !apierrors.IsServerTimeout(err) {
		klog.Errorf("CIDR assignment for node %v failed: %v. Releasing allocated CIDR", node.Name, err)
		if releaseErr := r.cidrs.Release(data.cidr); releaseErr != nil {
			klog.Errorf("Error releasing allocated CIDR for node %v: %v", node.Name, releaseErr)
		}
	}
	return err
}

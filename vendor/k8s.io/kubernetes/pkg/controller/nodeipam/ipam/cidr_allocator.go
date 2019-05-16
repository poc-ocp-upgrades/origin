package ipam

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	informers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"net"
	"time"
)

type nodeAndCIDR struct {
	cidr     *net.IPNet
	nodeName string
}
type CIDRAllocatorType string

const (
	RangeAllocatorType           CIDRAllocatorType = "RangeAllocator"
	CloudAllocatorType           CIDRAllocatorType = "CloudAllocator"
	IPAMFromClusterAllocatorType                   = "IPAMFromCluster"
	IPAMFromCloudAllocatorType                     = "IPAMFromCloud"
)
const (
	apiserverStartupGracePeriod = 10 * time.Minute
	cidrUpdateWorkers           = 30
	cidrUpdateQueueSize         = 5000
	cidrUpdateRetries           = 3
	updateRetryTimeout          = 250 * time.Millisecond
	maxUpdateRetryTimeout       = 5 * time.Second
	updateMaxRetries            = 10
)

type CIDRAllocator interface {
	AllocateOrOccupyCIDR(node *v1.Node) error
	ReleaseCIDR(node *v1.Node) error
	Run(stopCh <-chan struct{})
}

func New(kubeClient clientset.Interface, cloud cloudprovider.Interface, nodeInformer informers.NodeInformer, allocatorType CIDRAllocatorType, clusterCIDR, serviceCIDR *net.IPNet, nodeCIDRMaskSize int) (CIDRAllocator, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeList, err := listNodes(kubeClient)
	if err != nil {
		return nil, err
	}
	switch allocatorType {
	case RangeAllocatorType:
		return NewCIDRRangeAllocator(kubeClient, nodeInformer, clusterCIDR, serviceCIDR, nodeCIDRMaskSize, nodeList)
	case CloudAllocatorType:
		return NewCloudCIDRAllocator(kubeClient, cloud, nodeInformer)
	default:
		return nil, fmt.Errorf("Invalid CIDR allocator type: %v", allocatorType)
	}
}
func listNodes(kubeClient clientset.Interface) (*v1.NodeList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var nodeList *v1.NodeList
	if pollErr := wait.Poll(10*time.Second, apiserverStartupGracePeriod, func() (bool, error) {
		var err error
		nodeList, err = kubeClient.CoreV1().Nodes().List(metav1.ListOptions{FieldSelector: fields.Everything().String(), LabelSelector: labels.Everything().String()})
		if err != nil {
			klog.Errorf("Failed to list all nodes: %v", err)
			return false, nil
		}
		return true, nil
	}); pollErr != nil {
		return nil, fmt.Errorf("Failed to list all nodes in %v, cannot proceed without updating CIDR map", apiserverStartupGracePeriod)
	}
	return nodeList, nil
}

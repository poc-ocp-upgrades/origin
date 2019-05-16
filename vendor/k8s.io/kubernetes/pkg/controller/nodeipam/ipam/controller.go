package ipam

import (
	"fmt"
	"k8s.io/api/core/v1"
	informers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce"
	"k8s.io/kubernetes/pkg/controller/nodeipam/ipam/cidrset"
	nodesync "k8s.io/kubernetes/pkg/controller/nodeipam/ipam/sync"
	nodeutil "k8s.io/kubernetes/pkg/controller/util/node"
	"net"
	"sync"
	"time"
)

type Config struct {
	Resync       time.Duration
	MaxBackoff   time.Duration
	InitialRetry time.Duration
	Mode         nodesync.NodeSyncMode
}
type Controller struct {
	config  *Config
	adapter *adapter
	lock    sync.Mutex
	syncers map[string]*nodesync.NodeSync
	set     *cidrset.CidrSet
}

func NewController(config *Config, kubeClient clientset.Interface, cloud cloudprovider.Interface, clusterCIDR, serviceCIDR *net.IPNet, nodeCIDRMaskSize int) (*Controller, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !nodesync.IsValidMode(config.Mode) {
		return nil, fmt.Errorf("invalid IPAM controller mode %q", config.Mode)
	}
	gceCloud, ok := cloud.(*gce.Cloud)
	if !ok {
		return nil, fmt.Errorf("cloud IPAM controller does not support %q provider", cloud.ProviderName())
	}
	set, err := cidrset.NewCIDRSet(clusterCIDR, nodeCIDRMaskSize)
	if err != nil {
		return nil, err
	}
	c := &Controller{config: config, adapter: newAdapter(kubeClient, gceCloud), syncers: make(map[string]*nodesync.NodeSync), set: set}
	if err := occupyServiceCIDR(c.set, clusterCIDR, serviceCIDR); err != nil {
		return nil, err
	}
	return c, nil
}
func (c *Controller) Start(nodeInformer informers.NodeInformer) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(0).Infof("Starting IPAM controller (config=%+v)", c.config)
	nodes, err := listNodes(c.adapter.k8s)
	if err != nil {
		return err
	}
	for _, node := range nodes.Items {
		if node.Spec.PodCIDR != "" {
			_, cidrRange, err := net.ParseCIDR(node.Spec.PodCIDR)
			if err == nil {
				c.set.Occupy(cidrRange)
				klog.V(3).Infof("Occupying CIDR for node %q (%v)", node.Name, node.Spec.PodCIDR)
			} else {
				klog.Errorf("Node %q has an invalid CIDR (%q): %v", node.Name, node.Spec.PodCIDR, err)
			}
		}
		func() {
			c.lock.Lock()
			defer c.lock.Unlock()
			syncer := c.newSyncer(node.Name)
			c.syncers[node.Name] = syncer
			go syncer.Loop(nil)
		}()
	}
	nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: nodeutil.CreateAddNodeHandler(c.onAdd), UpdateFunc: nodeutil.CreateUpdateNodeHandler(c.onUpdate), DeleteFunc: nodeutil.CreateDeleteNodeHandler(c.onDelete)})
	return nil
}
func occupyServiceCIDR(set *cidrset.CidrSet, clusterCIDR, serviceCIDR *net.IPNet) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if clusterCIDR.Contains(serviceCIDR.IP) || serviceCIDR.Contains(clusterCIDR.IP) {
		if err := set.Occupy(serviceCIDR); err != nil {
			return err
		}
	}
	return nil
}

type nodeState struct{ t Timeout }

func (ns *nodeState) ReportResult(err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ns.t.Update(err == nil)
}
func (ns *nodeState) ResyncTimeout() time.Duration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ns.t.Next()
}
func (c *Controller) newSyncer(name string) *nodesync.NodeSync {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ns := &nodeState{Timeout{Resync: c.config.Resync, MaxBackoff: c.config.MaxBackoff, InitialRetry: c.config.InitialRetry}}
	return nodesync.New(ns, c.adapter, c.adapter, c.config.Mode, name, c.set)
}
func (c *Controller) onAdd(node *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.lock.Lock()
	defer c.lock.Unlock()
	syncer, ok := c.syncers[node.Name]
	if !ok {
		syncer = c.newSyncer(node.Name)
		c.syncers[node.Name] = syncer
		go syncer.Loop(nil)
	} else {
		klog.Warningf("Add for node %q that already exists", node.Name)
	}
	syncer.Update(node)
	return nil
}
func (c *Controller) onUpdate(_, node *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.lock.Lock()
	defer c.lock.Unlock()
	if sync, ok := c.syncers[node.Name]; ok {
		sync.Update(node)
	} else {
		klog.Errorf("Received update for non-existent node %q", node.Name)
		return fmt.Errorf("unknown node %q", node.Name)
	}
	return nil
}
func (c *Controller) onDelete(node *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.lock.Lock()
	defer c.lock.Unlock()
	if syncer, ok := c.syncers[node.Name]; ok {
		syncer.Delete(node)
		delete(c.syncers, node.Name)
	} else {
		klog.Warningf("Node %q was already deleted", node.Name)
	}
	return nil
}

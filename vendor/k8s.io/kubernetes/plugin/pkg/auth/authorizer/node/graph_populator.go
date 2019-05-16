package node

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	storagev1beta1 "k8s.io/api/storage/v1beta1"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	corev1informers "k8s.io/client-go/informers/core/v1"
	storageinformers "k8s.io/client-go/informers/storage/v1beta1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/features"
)

type graphPopulator struct{ graph *Graph }

func AddGraphEventHandlers(graph *Graph, nodes corev1informers.NodeInformer, pods corev1informers.PodInformer, pvs corev1informers.PersistentVolumeInformer, attachments storageinformers.VolumeAttachmentInformer) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g := &graphPopulator{graph: graph}
	if utilfeature.DefaultFeatureGate.Enabled(features.DynamicKubeletConfig) {
		nodes.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: g.addNode, UpdateFunc: g.updateNode, DeleteFunc: g.deleteNode})
	}
	pods.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: g.addPod, UpdateFunc: g.updatePod, DeleteFunc: g.deletePod})
	pvs.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: g.addPV, UpdateFunc: g.updatePV, DeleteFunc: g.deletePV})
	if utilfeature.DefaultFeatureGate.Enabled(features.CSIPersistentVolume) {
		attachments.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: g.addVolumeAttachment, UpdateFunc: g.updateVolumeAttachment, DeleteFunc: g.deleteVolumeAttachment})
	}
}
func (g *graphPopulator) addNode(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.updateNode(nil, obj)
}
func (g *graphPopulator) updateNode(oldObj, obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	node := obj.(*corev1.Node)
	var oldNode *corev1.Node
	if oldObj != nil {
		oldNode = oldObj.(*corev1.Node)
	}
	var name, namespace string
	if source := node.Spec.ConfigSource; source != nil && source.ConfigMap != nil {
		name = source.ConfigMap.Name
		namespace = source.ConfigMap.Namespace
	}
	var oldName, oldNamespace string
	if oldNode != nil {
		if oldSource := oldNode.Spec.ConfigSource; oldSource != nil && oldSource.ConfigMap != nil {
			oldName = oldSource.ConfigMap.Name
			oldNamespace = oldSource.ConfigMap.Namespace
		}
	}
	if name == oldName && namespace == oldNamespace {
		return
	}
	path := "nil"
	if node.Spec.ConfigSource != nil {
		path = fmt.Sprintf("%s/%s", namespace, name)
	}
	klog.V(4).Infof("updateNode configSource reference to %s for node %s", path, node.Name)
	g.graph.SetNodeConfigMap(node.Name, name, namespace)
}
func (g *graphPopulator) deleteNode(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if tombstone, ok := obj.(cache.DeletedFinalStateUnknown); ok {
		obj = tombstone.Obj
	}
	node, ok := obj.(*corev1.Node)
	if !ok {
		klog.Infof("unexpected type %T", obj)
		return
	}
	g.graph.SetNodeConfigMap(node.Name, "", "")
}
func (g *graphPopulator) addPod(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.updatePod(nil, obj)
}
func (g *graphPopulator) updatePod(oldObj, obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod := obj.(*corev1.Pod)
	if len(pod.Spec.NodeName) == 0 {
		klog.V(5).Infof("updatePod %s/%s, no node", pod.Namespace, pod.Name)
		return
	}
	if oldPod, ok := oldObj.(*corev1.Pod); ok && oldPod != nil {
		if (pod.Spec.NodeName == oldPod.Spec.NodeName) && (pod.UID == oldPod.UID) {
			klog.V(5).Infof("updatePod %s/%s, node unchanged", pod.Namespace, pod.Name)
			return
		}
	}
	klog.V(4).Infof("updatePod %s/%s for node %s", pod.Namespace, pod.Name, pod.Spec.NodeName)
	g.graph.AddPod(pod)
}
func (g *graphPopulator) deletePod(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if tombstone, ok := obj.(cache.DeletedFinalStateUnknown); ok {
		obj = tombstone.Obj
	}
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		klog.Infof("unexpected type %T", obj)
		return
	}
	if len(pod.Spec.NodeName) == 0 {
		klog.V(5).Infof("deletePod %s/%s, no node", pod.Namespace, pod.Name)
		return
	}
	klog.V(4).Infof("deletePod %s/%s for node %s", pod.Namespace, pod.Name, pod.Spec.NodeName)
	g.graph.DeletePod(pod.Name, pod.Namespace)
}
func (g *graphPopulator) addPV(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.updatePV(nil, obj)
}
func (g *graphPopulator) updatePV(oldObj, obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pv := obj.(*corev1.PersistentVolume)
	g.graph.AddPV(pv)
}
func (g *graphPopulator) deletePV(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if tombstone, ok := obj.(cache.DeletedFinalStateUnknown); ok {
		obj = tombstone.Obj
	}
	pv, ok := obj.(*corev1.PersistentVolume)
	if !ok {
		klog.Infof("unexpected type %T", obj)
		return
	}
	g.graph.DeletePV(pv.Name)
}
func (g *graphPopulator) addVolumeAttachment(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.updateVolumeAttachment(nil, obj)
}
func (g *graphPopulator) updateVolumeAttachment(oldObj, obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attachment := obj.(*storagev1beta1.VolumeAttachment)
	if oldObj != nil {
		oldAttachment := oldObj.(*storagev1beta1.VolumeAttachment)
		if oldAttachment.Spec.NodeName == attachment.Spec.NodeName {
			return
		}
	}
	g.graph.AddVolumeAttachment(attachment.Name, attachment.Spec.NodeName)
}
func (g *graphPopulator) deleteVolumeAttachment(obj interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if tombstone, ok := obj.(cache.DeletedFinalStateUnknown); ok {
		obj = tombstone.Obj
	}
	attachment, ok := obj.(*storagev1beta1.VolumeAttachment)
	if !ok {
		klog.Infof("unexpected type %T", obj)
		return
	}
	g.graph.DeleteVolumeAttachment(attachment.Name)
}

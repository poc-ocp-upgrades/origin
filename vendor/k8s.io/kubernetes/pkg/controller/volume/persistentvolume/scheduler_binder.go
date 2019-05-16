package persistentvolume

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"
	storageinformers "k8s.io/client-go/informers/storage/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
	"sort"
	"time"
)

type SchedulerVolumeBinder interface {
	FindPodVolumes(pod *v1.Pod, node *v1.Node) (unboundVolumesSatisified, boundVolumesSatisfied bool, err error)
	AssumePodVolumes(assumedPod *v1.Pod, nodeName string) (allFullyBound bool, err error)
	BindPodVolumes(assumedPod *v1.Pod) error
	GetBindingsCache() PodBindingCache
}
type volumeBinder struct {
	ctrl            *PersistentVolumeController
	pvcCache        PVCAssumeCache
	pvCache         PVAssumeCache
	podBindingCache PodBindingCache
	bindTimeout     time.Duration
}

func NewVolumeBinder(kubeClient clientset.Interface, pvcInformer coreinformers.PersistentVolumeClaimInformer, pvInformer coreinformers.PersistentVolumeInformer, storageClassInformer storageinformers.StorageClassInformer, bindTimeout time.Duration) SchedulerVolumeBinder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctrl := &PersistentVolumeController{kubeClient: kubeClient, classLister: storageClassInformer.Lister()}
	b := &volumeBinder{ctrl: ctrl, pvcCache: NewPVCAssumeCache(pvcInformer.Informer()), pvCache: NewPVAssumeCache(pvInformer.Informer()), podBindingCache: NewPodBindingCache(), bindTimeout: bindTimeout}
	return b
}
func (b *volumeBinder) GetBindingsCache() PodBindingCache {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return b.podBindingCache
}
func (b *volumeBinder) FindPodVolumes(pod *v1.Pod, node *v1.Node) (unboundVolumesSatisfied, boundVolumesSatisfied bool, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podName := getPodName(pod)
	klog.V(5).Infof("FindPodVolumes for pod %q, node %q", podName, node.Name)
	unboundVolumesSatisfied = true
	boundVolumesSatisfied = true
	start := time.Now()
	defer func() {
		VolumeSchedulingStageLatency.WithLabelValues("predicate").Observe(time.Since(start).Seconds())
		if err != nil {
			VolumeSchedulingStageFailed.WithLabelValues("predicate").Inc()
		}
	}()
	boundClaims, claimsToBind, unboundClaimsImmediate, err := b.getPodVolumes(pod)
	if err != nil {
		return false, false, err
	}
	if len(unboundClaimsImmediate) > 0 {
		return false, false, fmt.Errorf("pod has unbound immediate PersistentVolumeClaims")
	}
	if len(boundClaims) > 0 {
		boundVolumesSatisfied, err = b.checkBoundClaims(boundClaims, node, podName)
		if err != nil {
			return false, false, err
		}
	}
	if len(claimsToBind) > 0 {
		var claimsToProvision []*v1.PersistentVolumeClaim
		unboundVolumesSatisfied, claimsToProvision, err = b.findMatchingVolumes(pod, claimsToBind, node)
		if err != nil {
			return false, false, err
		}
		if !unboundVolumesSatisfied {
			unboundVolumesSatisfied, err = b.checkVolumeProvisions(pod, claimsToProvision, node)
			if err != nil {
				return false, false, err
			}
		}
	}
	return unboundVolumesSatisfied, boundVolumesSatisfied, nil
}
func (b *volumeBinder) AssumePodVolumes(assumedPod *v1.Pod, nodeName string) (allFullyBound bool, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podName := getPodName(assumedPod)
	klog.V(4).Infof("AssumePodVolumes for pod %q, node %q", podName, nodeName)
	start := time.Now()
	defer func() {
		VolumeSchedulingStageLatency.WithLabelValues("assume").Observe(time.Since(start).Seconds())
		if err != nil {
			VolumeSchedulingStageFailed.WithLabelValues("assume").Inc()
		}
	}()
	if allBound := b.arePodVolumesBound(assumedPod); allBound {
		klog.V(4).Infof("AssumePodVolumes for pod %q, node %q: all PVCs bound and nothing to do", podName, nodeName)
		return true, nil
	}
	assumedPod.Spec.NodeName = nodeName
	claimsToBind := b.podBindingCache.GetBindings(assumedPod, nodeName)
	claimsToProvision := b.podBindingCache.GetProvisionedPVCs(assumedPod, nodeName)
	newBindings := []*bindingInfo{}
	for _, binding := range claimsToBind {
		newPV, dirty, err := b.ctrl.getBindVolumeToClaim(binding.pv, binding.pvc)
		klog.V(5).Infof("AssumePodVolumes: getBindVolumeToClaim for pod %q, PV %q, PVC %q.  newPV %p, dirty %v, err: %v", podName, binding.pv.Name, binding.pvc.Name, newPV, dirty, err)
		if err != nil {
			b.revertAssumedPVs(newBindings)
			return false, err
		}
		if dirty {
			err = b.pvCache.Assume(newPV)
			if err != nil {
				b.revertAssumedPVs(newBindings)
				return false, err
			}
		}
		newBindings = append(newBindings, &bindingInfo{pv: newPV, pvc: binding.pvc})
	}
	newProvisionedPVCs := []*v1.PersistentVolumeClaim{}
	for _, claim := range claimsToProvision {
		claimClone := claim.DeepCopy()
		metav1.SetMetaDataAnnotation(&claimClone.ObjectMeta, annSelectedNode, nodeName)
		err = b.pvcCache.Assume(claimClone)
		if err != nil {
			b.revertAssumedPVs(newBindings)
			b.revertAssumedPVCs(newProvisionedPVCs)
			return
		}
		newProvisionedPVCs = append(newProvisionedPVCs, claimClone)
	}
	b.podBindingCache.UpdateBindings(assumedPod, nodeName, newBindings)
	b.podBindingCache.UpdateProvisionedPVCs(assumedPod, nodeName, newProvisionedPVCs)
	return
}
func (b *volumeBinder) BindPodVolumes(assumedPod *v1.Pod) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podName := getPodName(assumedPod)
	klog.V(4).Infof("BindPodVolumes for pod %q, node %q", podName, assumedPod.Spec.NodeName)
	start := time.Now()
	defer func() {
		VolumeSchedulingStageLatency.WithLabelValues("bind").Observe(time.Since(start).Seconds())
		if err != nil {
			VolumeSchedulingStageFailed.WithLabelValues("bind").Inc()
		}
	}()
	bindings := b.podBindingCache.GetBindings(assumedPod, assumedPod.Spec.NodeName)
	claimsToProvision := b.podBindingCache.GetProvisionedPVCs(assumedPod, assumedPod.Spec.NodeName)
	err = b.bindAPIUpdate(podName, bindings, claimsToProvision)
	if err != nil {
		return err
	}
	return wait.Poll(time.Second, b.bindTimeout, func() (bool, error) {
		bindings = b.podBindingCache.GetBindings(assumedPod, assumedPod.Spec.NodeName)
		claimsToProvision = b.podBindingCache.GetProvisionedPVCs(assumedPod, assumedPod.Spec.NodeName)
		return b.checkBindings(assumedPod, bindings, claimsToProvision)
	})
}
func getPodName(pod *v1.Pod) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pod.Namespace + "/" + pod.Name
}
func getPVCName(pvc *v1.PersistentVolumeClaim) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pvc.Namespace + "/" + pvc.Name
}
func (b *volumeBinder) bindAPIUpdate(podName string, bindings []*bindingInfo, claimsToProvision []*v1.PersistentVolumeClaim) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if bindings == nil {
		return fmt.Errorf("failed to get cached bindings for pod %q", podName)
	}
	if claimsToProvision == nil {
		return fmt.Errorf("failed to get cached claims to provision for pod %q", podName)
	}
	lastProcessedBinding := 0
	lastProcessedProvisioning := 0
	defer func() {
		if lastProcessedBinding < len(bindings) {
			b.revertAssumedPVs(bindings[lastProcessedBinding:])
		}
		if lastProcessedProvisioning < len(claimsToProvision) {
			b.revertAssumedPVCs(claimsToProvision[lastProcessedProvisioning:])
		}
	}()
	var (
		binding *bindingInfo
		claim   *v1.PersistentVolumeClaim
	)
	for _, binding = range bindings {
		klog.V(5).Infof("bindAPIUpdate: Pod %q, binding PV %q to PVC %q", podName, binding.pv.Name, binding.pvc.Name)
		if _, err := b.ctrl.updateBindVolumeToClaim(binding.pv, binding.pvc, false); err != nil {
			return err
		}
		lastProcessedBinding++
	}
	for _, claim = range claimsToProvision {
		klog.V(5).Infof("bindAPIUpdate: Pod %q, PVC %q", podName, getPVCName(claim))
		if _, err := b.ctrl.kubeClient.CoreV1().PersistentVolumeClaims(claim.Namespace).Update(claim); err != nil {
			return err
		}
		lastProcessedProvisioning++
	}
	return nil
}
func (b *volumeBinder) checkBindings(pod *v1.Pod, bindings []*bindingInfo, claimsToProvision []*v1.PersistentVolumeClaim) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podName := getPodName(pod)
	if bindings == nil {
		return false, fmt.Errorf("failed to get cached bindings for pod %q", podName)
	}
	if claimsToProvision == nil {
		return false, fmt.Errorf("failed to get cached claims to provision for pod %q", podName)
	}
	for _, binding := range bindings {
		pv, err := b.pvCache.GetPV(binding.pv.Name)
		if err != nil || pv == nil {
			return false, fmt.Errorf("failed to check pv binding: %v", err)
		}
		if pv.Spec.ClaimRef == nil || pv.Spec.ClaimRef.UID == "" {
			return false, fmt.Errorf("ClaimRef got reset for pv %q", pv.Name)
		}
		if isBound, _, err := b.isPVCBound(binding.pvc.Namespace, binding.pvc.Name); !isBound || err != nil {
			return false, err
		}
	}
	for _, claim := range claimsToProvision {
		bound, pvc, err := b.isPVCBound(claim.Namespace, claim.Name)
		if err != nil || pvc == nil {
			return false, fmt.Errorf("failed to check pvc binding: %v", err)
		}
		if pvc.Annotations == nil {
			return false, fmt.Errorf("selectedNode annotation reset for PVC %q", pvc.Name)
		}
		selectedNode := pvc.Annotations[annSelectedNode]
		if selectedNode != pod.Spec.NodeName {
			return false, fmt.Errorf("selectedNode annotation value %q not set to scheduled node %q", selectedNode, pod.Spec.NodeName)
		}
		if !bound {
			return false, nil
		}
	}
	klog.V(4).Infof("All PVCs for pod %q are bound", podName)
	return true, nil
}
func (b *volumeBinder) isVolumeBound(namespace string, vol *v1.Volume) (bool, *v1.PersistentVolumeClaim, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if vol.PersistentVolumeClaim == nil {
		return true, nil, nil
	}
	pvcName := vol.PersistentVolumeClaim.ClaimName
	return b.isPVCBound(namespace, pvcName)
}
func (b *volumeBinder) isPVCBound(namespace, pvcName string) (bool, *v1.PersistentVolumeClaim, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	claim := &v1.PersistentVolumeClaim{ObjectMeta: metav1.ObjectMeta{Name: pvcName, Namespace: namespace}}
	pvcKey := getPVCName(claim)
	pvc, err := b.pvcCache.GetPVC(pvcKey)
	if err != nil || pvc == nil {
		return false, nil, fmt.Errorf("error getting PVC %q: %v", pvcKey, err)
	}
	pvName := pvc.Spec.VolumeName
	if pvName != "" {
		if metav1.HasAnnotation(pvc.ObjectMeta, annBindCompleted) {
			klog.V(5).Infof("PVC %q is fully bound to PV %q", pvcKey, pvName)
			return true, pvc, nil
		} else {
			klog.V(5).Infof("PVC %q is not fully bound to PV %q", pvcKey, pvName)
			return false, pvc, nil
		}
	}
	klog.V(5).Infof("PVC %q is not bound", pvcKey)
	return false, pvc, nil
}
func (b *volumeBinder) arePodVolumesBound(pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, vol := range pod.Spec.Volumes {
		if isBound, _, _ := b.isVolumeBound(pod.Namespace, &vol); !isBound {
			return false
		}
	}
	return true
}
func (b *volumeBinder) getPodVolumes(pod *v1.Pod) (boundClaims []*v1.PersistentVolumeClaim, unboundClaims []*bindingInfo, unboundClaimsImmediate []*v1.PersistentVolumeClaim, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	boundClaims = []*v1.PersistentVolumeClaim{}
	unboundClaimsImmediate = []*v1.PersistentVolumeClaim{}
	unboundClaims = []*bindingInfo{}
	for _, vol := range pod.Spec.Volumes {
		volumeBound, pvc, err := b.isVolumeBound(pod.Namespace, &vol)
		if err != nil {
			return nil, nil, nil, err
		}
		if pvc == nil {
			continue
		}
		if volumeBound {
			boundClaims = append(boundClaims, pvc)
		} else {
			delayBinding, err := b.ctrl.shouldDelayBinding(pvc)
			if err != nil {
				return nil, nil, nil, err
			}
			if delayBinding && pvc.Spec.VolumeName == "" {
				unboundClaims = append(unboundClaims, &bindingInfo{pvc: pvc})
			} else {
				unboundClaimsImmediate = append(unboundClaimsImmediate, pvc)
			}
		}
	}
	return boundClaims, unboundClaims, unboundClaimsImmediate, nil
}
func (b *volumeBinder) checkBoundClaims(claims []*v1.PersistentVolumeClaim, node *v1.Node, podName string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, pvc := range claims {
		pvName := pvc.Spec.VolumeName
		pv, err := b.pvCache.GetPV(pvName)
		if err != nil {
			return false, err
		}
		err = volumeutil.CheckNodeAffinity(pv, node.Labels)
		if err != nil {
			klog.V(4).Infof("PersistentVolume %q, Node %q mismatch for Pod %q: %v", pvName, node.Name, podName, err)
			return false, nil
		}
		klog.V(5).Infof("PersistentVolume %q, Node %q matches for Pod %q", pvName, node.Name, podName)
	}
	klog.V(4).Infof("All bound volumes for Pod %q match with Node %q", podName, node.Name)
	return true, nil
}
func (b *volumeBinder) findMatchingVolumes(pod *v1.Pod, claimsToBind []*bindingInfo, node *v1.Node) (foundMatches bool, unboundClaims []*v1.PersistentVolumeClaim, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podName := getPodName(pod)
	sort.Sort(byPVCSize(claimsToBind))
	chosenPVs := map[string]*v1.PersistentVolume{}
	foundMatches = true
	matchedClaims := []*bindingInfo{}
	for _, bindingInfo := range claimsToBind {
		storageClassName := ""
		storageClass := bindingInfo.pvc.Spec.StorageClassName
		if storageClass != nil {
			storageClassName = *storageClass
		}
		allPVs := b.pvCache.ListPVs(storageClassName)
		pvcName := getPVCName(bindingInfo.pvc)
		bindingInfo.pv, err = findMatchingVolume(bindingInfo.pvc, allPVs, node, chosenPVs, true)
		if err != nil {
			return false, nil, err
		}
		if bindingInfo.pv == nil {
			klog.V(4).Infof("No matching volumes for Pod %q, PVC %q on node %q", podName, pvcName, node.Name)
			unboundClaims = append(unboundClaims, bindingInfo.pvc)
			foundMatches = false
			continue
		}
		chosenPVs[bindingInfo.pv.Name] = bindingInfo.pv
		matchedClaims = append(matchedClaims, bindingInfo)
		klog.V(5).Infof("Found matching PV %q for PVC %q on node %q for pod %q", bindingInfo.pv.Name, pvcName, node.Name, podName)
	}
	if len(matchedClaims) > 0 {
		b.podBindingCache.UpdateBindings(pod, node.Name, matchedClaims)
	}
	if foundMatches {
		klog.V(4).Infof("Found matching volumes for pod %q on node %q", podName, node.Name)
	}
	return
}
func (b *volumeBinder) checkVolumeProvisions(pod *v1.Pod, claimsToProvision []*v1.PersistentVolumeClaim, node *v1.Node) (provisionSatisfied bool, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podName := getPodName(pod)
	provisionedClaims := []*v1.PersistentVolumeClaim{}
	for _, claim := range claimsToProvision {
		pvcName := getPVCName(claim)
		className := v1helper.GetPersistentVolumeClaimClass(claim)
		if className == "" {
			return false, fmt.Errorf("no class for claim %q", pvcName)
		}
		class, err := b.ctrl.classLister.Get(className)
		if err != nil {
			return false, fmt.Errorf("failed to find storage class %q", className)
		}
		provisioner := class.Provisioner
		if provisioner == "" || provisioner == notSupportedProvisioner {
			klog.V(4).Infof("storage class %q of claim %q does not support dynamic provisioning", className, pvcName)
			return false, nil
		}
		if !v1helper.MatchTopologySelectorTerms(class.AllowedTopologies, labels.Set(node.Labels)) {
			klog.V(4).Infof("Node %q cannot satisfy provisioning topology requirements of claim %q", node.Name, pvcName)
			return false, nil
		}
		provisionedClaims = append(provisionedClaims, claim)
	}
	klog.V(4).Infof("Provisioning for claims of pod %q that has no matching volumes on node %q ...", podName, node.Name)
	b.podBindingCache.UpdateProvisionedPVCs(pod, node.Name, provisionedClaims)
	return true, nil
}
func (b *volumeBinder) revertAssumedPVs(bindings []*bindingInfo) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, bindingInfo := range bindings {
		b.pvCache.Restore(bindingInfo.pv.Name)
	}
}
func (b *volumeBinder) revertAssumedPVCs(claims []*v1.PersistentVolumeClaim) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, claim := range claims {
		b.pvcCache.Restore(getPVCName(claim))
	}
}

type bindingInfo struct {
	pvc *v1.PersistentVolumeClaim
	pv  *v1.PersistentVolume
}
type byPVCSize []*bindingInfo

func (a byPVCSize) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(a)
}
func (a byPVCSize) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a[i], a[j] = a[j], a[i]
}
func (a byPVCSize) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	iSize := a[i].pvc.Spec.Resources.Requests[v1.ResourceStorage]
	jSize := a[j].pvc.Spec.Resources.Requests[v1.ResourceStorage]
	return iSize.Cmp(jSize) == -1
}

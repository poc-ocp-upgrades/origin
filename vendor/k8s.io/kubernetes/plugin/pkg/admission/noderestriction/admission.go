package noderestriction

import (
	"fmt"
	goformat "fmt"
	"io"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/admission"
	apiserveradmission "k8s.io/apiserver/pkg/admission/initializer"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/informers"
	corev1lister "k8s.io/client-go/listers/core/v1"
	csiv1alpha1 "k8s.io/csi-api/pkg/apis/csi/v1alpha1"
	"k8s.io/klog"
	podutil "k8s.io/kubernetes/pkg/api/pod"
	authenticationapi "k8s.io/kubernetes/pkg/apis/authentication"
	coordapi "k8s.io/kubernetes/pkg/apis/coordination"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/policy"
	"k8s.io/kubernetes/pkg/auth/nodeidentifier"
	"k8s.io/kubernetes/pkg/features"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	PluginName = "NodeRestriction"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewPlugin(nodeidentifier.NewDefaultNodeIdentifier()), nil
	})
}
func NewPlugin(nodeIdentifier nodeidentifier.NodeIdentifier) *nodePlugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &nodePlugin{Handler: admission.NewHandler(admission.Create, admission.Update, admission.Delete), nodeIdentifier: nodeIdentifier, features: utilfeature.DefaultFeatureGate}
}

type nodePlugin struct {
	*admission.Handler
	nodeIdentifier nodeidentifier.NodeIdentifier
	podsGetter     corev1lister.PodLister
	features       utilfeature.FeatureGate
}

var (
	_ = admission.Interface(&nodePlugin{})
	_ = apiserveradmission.WantsExternalKubeInformerFactory(&nodePlugin{})
)

func (p *nodePlugin) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p.podsGetter = f.Core().V1().Pods().Lister()
}
func (p *nodePlugin) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if p.nodeIdentifier == nil {
		return fmt.Errorf("%s requires a node identifier", PluginName)
	}
	if p.podsGetter == nil {
		return fmt.Errorf("%s requires a pod getter", PluginName)
	}
	return nil
}

var (
	podResource         = api.Resource("pods")
	nodeResource        = api.Resource("nodes")
	pvcResource         = api.Resource("persistentvolumeclaims")
	svcacctResource     = api.Resource("serviceaccounts")
	leaseResource       = coordapi.Resource("leases")
	csiNodeInfoResource = csiv1alpha1.Resource("csinodeinfos")
)

func (c *nodePlugin) Admit(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeName, isNode := c.nodeIdentifier.NodeIdentity(a.GetUserInfo())
	if !isNode {
		return nil
	}
	if len(nodeName) == 0 {
		return admission.NewForbidden(a, fmt.Errorf("could not determine node from user %q", a.GetUserInfo().GetName()))
	}
	switch a.GetResource().GroupResource() {
	case podResource:
		switch a.GetSubresource() {
		case "":
			return c.admitPod(nodeName, a)
		case "status":
			return c.admitPodStatus(nodeName, a)
		case "eviction":
			return c.admitPodEviction(nodeName, a)
		default:
			return admission.NewForbidden(a, fmt.Errorf("unexpected pod subresource %q", a.GetSubresource()))
		}
	case nodeResource:
		return c.admitNode(nodeName, a)
	case pvcResource:
		switch a.GetSubresource() {
		case "status":
			return c.admitPVCStatus(nodeName, a)
		default:
			return admission.NewForbidden(a, fmt.Errorf("may only update PVC status"))
		}
	case svcacctResource:
		if c.features.Enabled(features.TokenRequest) {
			return c.admitServiceAccount(nodeName, a)
		}
		return nil
	case leaseResource:
		if c.features.Enabled(features.NodeLease) {
			return c.admitLease(nodeName, a)
		}
		return admission.NewForbidden(a, fmt.Errorf("disabled by feature gate %s", features.NodeLease))
	case csiNodeInfoResource:
		if c.features.Enabled(features.KubeletPluginsWatcher) && c.features.Enabled(features.CSINodeInfo) {
			return c.admitCSINodeInfo(nodeName, a)
		}
		return admission.NewForbidden(a, fmt.Errorf("disabled by feature gates %s and %s", features.KubeletPluginsWatcher, features.CSINodeInfo))
	default:
		return nil
	}
}
func (c *nodePlugin) admitPod(nodeName string, a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch a.GetOperation() {
	case admission.Create:
		pod, ok := a.GetObject().(*api.Pod)
		if !ok {
			return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetObject()))
		}
		if _, isMirrorPod := pod.Annotations[api.MirrorPodAnnotationKey]; !isMirrorPod {
			return admission.NewForbidden(a, fmt.Errorf("pod does not have %q annotation, node %q can only create mirror pods", api.MirrorPodAnnotationKey, nodeName))
		}
		if pod.Spec.NodeName != nodeName {
			return admission.NewForbidden(a, fmt.Errorf("node %q can only create pods with spec.nodeName set to itself", nodeName))
		}
		if pod.Spec.ServiceAccountName != "" {
			return admission.NewForbidden(a, fmt.Errorf("node %q can not create pods that reference a service account", nodeName))
		}
		hasSecrets := false
		podutil.VisitPodSecretNames(pod, func(name string) (shouldContinue bool) {
			hasSecrets = true
			return false
		})
		if hasSecrets {
			return admission.NewForbidden(a, fmt.Errorf("node %q can not create pods that reference secrets", nodeName))
		}
		hasConfigMaps := false
		podutil.VisitPodConfigmapNames(pod, func(name string) (shouldContinue bool) {
			hasConfigMaps = true
			return false
		})
		if hasConfigMaps {
			return admission.NewForbidden(a, fmt.Errorf("node %q can not create pods that reference configmaps", nodeName))
		}
		for _, v := range pod.Spec.Volumes {
			if v.PersistentVolumeClaim != nil {
				return admission.NewForbidden(a, fmt.Errorf("node %q can not create pods that reference persistentvolumeclaims", nodeName))
			}
		}
		return nil
	case admission.Delete:
		existingPod, err := c.podsGetter.Pods(a.GetNamespace()).Get(a.GetName())
		if errors.IsNotFound(err) {
			return err
		}
		if err != nil {
			return admission.NewForbidden(a, err)
		}
		if existingPod.Spec.NodeName != nodeName {
			return admission.NewForbidden(a, fmt.Errorf("node %q can only delete pods with spec.nodeName set to itself", nodeName))
		}
		return nil
	default:
		return admission.NewForbidden(a, fmt.Errorf("unexpected operation %q", a.GetOperation()))
	}
}
func (c *nodePlugin) admitPodStatus(nodeName string, a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch a.GetOperation() {
	case admission.Update:
		pod, ok := a.GetOldObject().(*api.Pod)
		if !ok {
			return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetOldObject()))
		}
		if pod.Spec.NodeName != nodeName {
			return admission.NewForbidden(a, fmt.Errorf("node %q can only update pod status for pods with spec.nodeName set to itself", nodeName))
		}
		return nil
	default:
		return admission.NewForbidden(a, fmt.Errorf("unexpected operation %q", a.GetOperation()))
	}
}
func (c *nodePlugin) admitPodEviction(nodeName string, a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch a.GetOperation() {
	case admission.Create:
		eviction, ok := a.GetObject().(*policy.Eviction)
		if !ok {
			return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetObject()))
		}
		podName := a.GetName()
		if len(podName) == 0 {
			if len(eviction.Name) == 0 {
				return admission.NewForbidden(a, fmt.Errorf("could not determine pod from request data"))
			}
			podName = eviction.Name
		}
		existingPod, err := c.podsGetter.Pods(a.GetNamespace()).Get(podName)
		if errors.IsNotFound(err) {
			return err
		}
		if err != nil {
			return admission.NewForbidden(a, err)
		}
		if existingPod.Spec.NodeName != nodeName {
			return admission.NewForbidden(a, fmt.Errorf("node %s can only evict pods with spec.nodeName set to itself", nodeName))
		}
		return nil
	default:
		return admission.NewForbidden(a, fmt.Errorf("unexpected operation %s", a.GetOperation()))
	}
}
func (c *nodePlugin) admitPVCStatus(nodeName string, a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch a.GetOperation() {
	case admission.Update:
		if !c.features.Enabled(features.ExpandPersistentVolumes) {
			return admission.NewForbidden(a, fmt.Errorf("node %q may not update persistentvolumeclaim metadata", nodeName))
		}
		oldPVC, ok := a.GetOldObject().(*api.PersistentVolumeClaim)
		if !ok {
			return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetOldObject()))
		}
		newPVC, ok := a.GetObject().(*api.PersistentVolumeClaim)
		if !ok {
			return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetObject()))
		}
		oldPVC = oldPVC.DeepCopy()
		newPVC = newPVC.DeepCopy()
		oldPVC.ObjectMeta.ResourceVersion = ""
		newPVC.ObjectMeta.ResourceVersion = ""
		oldPVC.Status.Capacity = nil
		newPVC.Status.Capacity = nil
		oldPVC.Status.Conditions = nil
		newPVC.Status.Conditions = nil
		if !apiequality.Semantic.DeepEqual(oldPVC, newPVC) {
			return admission.NewForbidden(a, fmt.Errorf("node %q may not update fields other than status.capacity and status.conditions: %v", nodeName, diff.ObjectReflectDiff(oldPVC, newPVC)))
		}
		return nil
	default:
		return admission.NewForbidden(a, fmt.Errorf("unexpected operation %q", a.GetOperation()))
	}
}
func (c *nodePlugin) admitNode(nodeName string, a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requestedName := a.GetName()
	if a.GetOperation() == admission.Create {
		node, ok := a.GetObject().(*api.Node)
		if !ok {
			return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetObject()))
		}
		if node.Spec.ConfigSource != nil {
			return admission.NewForbidden(a, fmt.Errorf("cannot create with non-nil configSource"))
		}
		modifiedLabels := getModifiedLabels(node.Labels, nil)
		if forbiddenLabels := c.getForbiddenCreateLabels(modifiedLabels); len(forbiddenLabels) > 0 {
			return admission.NewForbidden(a, fmt.Errorf("cannot set labels: %s", strings.Join(forbiddenLabels.List(), ", ")))
		}
		if forbiddenUpdateLabels := c.getForbiddenUpdateLabels(modifiedLabels); len(forbiddenUpdateLabels) > 0 {
			klog.Warningf("node %q added disallowed labels on node creation: %s", nodeName, strings.Join(forbiddenUpdateLabels.List(), ", "))
		}
		if len(requestedName) == 0 {
			requestedName = node.Name
		}
	}
	if requestedName != nodeName {
		return admission.NewForbidden(a, fmt.Errorf("node %q cannot modify node %q", nodeName, requestedName))
	}
	if a.GetOperation() == admission.Update {
		node, ok := a.GetObject().(*api.Node)
		if !ok {
			return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetObject()))
		}
		oldNode, ok := a.GetOldObject().(*api.Node)
		if !ok {
			return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetObject()))
		}
		if node.Spec.ConfigSource != nil && !apiequality.Semantic.DeepEqual(node.Spec.ConfigSource, oldNode.Spec.ConfigSource) {
			return admission.NewForbidden(a, fmt.Errorf("node %q cannot update configSource to a new non-nil configSource", nodeName))
		}
		if !apiequality.Semantic.DeepEqual(node.Spec.Taints, oldNode.Spec.Taints) {
			return admission.NewForbidden(a, fmt.Errorf("node %q cannot modify taints", nodeName))
		}
		modifiedLabels := getModifiedLabels(node.Labels, oldNode.Labels)
		if forbiddenUpdateLabels := c.getForbiddenUpdateLabels(modifiedLabels); len(forbiddenUpdateLabels) > 0 {
			return admission.NewForbidden(a, fmt.Errorf("cannot modify labels: %s", strings.Join(forbiddenUpdateLabels.List(), ", ")))
		}
	}
	return nil
}
func getModifiedLabels(a, b map[string]string) sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	modified := sets.NewString()
	for k, v1 := range a {
		if v2, ok := b[k]; !ok || v1 != v2 {
			modified.Insert(k)
		}
	}
	for k, v1 := range b {
		if v2, ok := a[k]; !ok || v1 != v2 {
			modified.Insert(k)
		}
	}
	return modified
}
func isKubernetesLabel(key string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespace := getLabelNamespace(key)
	if namespace == "kubernetes.io" || strings.HasSuffix(namespace, ".kubernetes.io") {
		return true
	}
	if namespace == "k8s.io" || strings.HasSuffix(namespace, ".k8s.io") {
		return true
	}
	return false
}
func getLabelNamespace(key string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if parts := strings.SplitN(key, "/", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}
func (c *nodePlugin) getForbiddenCreateLabels(modifiedLabels sets.String) sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(modifiedLabels) == 0 {
		return nil
	}
	forbiddenLabels := sets.NewString()
	for label := range modifiedLabels {
		namespace := getLabelNamespace(label)
		if namespace == kubeletapis.LabelNamespaceNodeRestriction || strings.HasSuffix(namespace, "."+kubeletapis.LabelNamespaceNodeRestriction) {
			forbiddenLabels.Insert(label)
		}
	}
	return forbiddenLabels
}
func (c *nodePlugin) getForbiddenUpdateLabels(modifiedLabels sets.String) sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(modifiedLabels) == 0 {
		return nil
	}
	forbiddenLabels := sets.NewString()
	for label := range modifiedLabels {
		namespace := getLabelNamespace(label)
		if namespace == kubeletapis.LabelNamespaceNodeRestriction || strings.HasSuffix(namespace, "."+kubeletapis.LabelNamespaceNodeRestriction) {
			forbiddenLabels.Insert(label)
		}
		if isKubernetesLabel(label) && !kubeletapis.IsKubeletLabel(label) {
			forbiddenLabels.Insert(label)
		}
	}
	return forbiddenLabels
}
func (c *nodePlugin) admitServiceAccount(nodeName string, a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetOperation() != admission.Create {
		return nil
	}
	if a.GetSubresource() != "token" {
		return nil
	}
	tr, ok := a.GetObject().(*authenticationapi.TokenRequest)
	if !ok {
		return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetObject()))
	}
	ref := tr.Spec.BoundObjectRef
	if ref == nil || ref.APIVersion != "v1" || ref.Kind != "Pod" || ref.Name == "" {
		return admission.NewForbidden(a, fmt.Errorf("node requested token not bound to a pod"))
	}
	if ref.UID == "" {
		return admission.NewForbidden(a, fmt.Errorf("node requested token with a pod binding without a uid"))
	}
	pod, err := c.podsGetter.Pods(a.GetNamespace()).Get(ref.Name)
	if errors.IsNotFound(err) {
		return err
	}
	if err != nil {
		return admission.NewForbidden(a, err)
	}
	if ref.UID != pod.UID {
		return admission.NewForbidden(a, fmt.Errorf("the UID in the bound object reference (%s) does not match the UID in record (%s). The object might have been deleted and then recreated", ref.UID, pod.UID))
	}
	if pod.Spec.NodeName != nodeName {
		return admission.NewForbidden(a, fmt.Errorf("node requested token bound to a pod scheduled on a different node"))
	}
	return nil
}
func (r *nodePlugin) admitLease(nodeName string, a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetNamespace() != api.NamespaceNodeLease {
		return admission.NewForbidden(a, fmt.Errorf("can only access leases in the %q system namespace", api.NamespaceNodeLease))
	}
	if a.GetOperation() == admission.Create {
		lease, ok := a.GetObject().(*coordapi.Lease)
		if !ok {
			return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetObject()))
		}
		if lease.Name != nodeName {
			return admission.NewForbidden(a, fmt.Errorf("can only access node lease with the same name as the requesting node"))
		}
	} else {
		if a.GetName() != nodeName {
			return admission.NewForbidden(a, fmt.Errorf("can only access node lease with the same name as the requesting node"))
		}
	}
	return nil
}
func (c *nodePlugin) admitCSINodeInfo(nodeName string, a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetOperation() == admission.Create {
		accessor, err := meta.Accessor(a.GetObject())
		if err != nil {
			return admission.NewForbidden(a, fmt.Errorf("unable to access the object name"))
		}
		if accessor.GetName() != nodeName {
			return admission.NewForbidden(a, fmt.Errorf("can only access CSINodeInfo with the same name as the requesting node"))
		}
	} else {
		if a.GetName() != nodeName {
			return admission.NewForbidden(a, fmt.Errorf("can only access CSINodeInfo with the same name as the requesting node"))
		}
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

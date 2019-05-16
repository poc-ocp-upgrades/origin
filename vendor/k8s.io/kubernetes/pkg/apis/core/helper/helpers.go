package helper

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func IsHugePageResourceName(name core.ResourceName) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.HasPrefix(string(name), core.ResourceHugePagesPrefix)
}
func IsQuotaHugePageResourceName(name core.ResourceName) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.HasPrefix(string(name), core.ResourceHugePagesPrefix) || strings.HasPrefix(string(name), core.ResourceRequestsHugePagesPrefix)
}
func HugePageResourceName(pageSize resource.Quantity) core.ResourceName {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return core.ResourceName(fmt.Sprintf("%s%s", core.ResourceHugePagesPrefix, pageSize.String()))
}
func HugePageSizeFromResourceName(name core.ResourceName) (resource.Quantity, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !IsHugePageResourceName(name) {
		return resource.Quantity{}, fmt.Errorf("resource name: %s is an invalid hugepage name", name)
	}
	pageSize := strings.TrimPrefix(string(name), core.ResourceHugePagesPrefix)
	return resource.ParseQuantity(pageSize)
}
func NonConvertibleFields(annotations map[string]string) map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nonConvertibleKeys := map[string]string{}
	for key, value := range annotations {
		if strings.HasPrefix(key, core.NonConvertibleAnnotationPrefix) {
			nonConvertibleKeys[key] = value
		}
	}
	return nonConvertibleKeys
}

var Semantic = conversion.EqualitiesOrDie(func(a, b resource.Quantity) bool {
	return a.Cmp(b) == 0
}, func(a, b metav1.MicroTime) bool {
	return a.UTC() == b.UTC()
}, func(a, b metav1.Time) bool {
	return a.UTC() == b.UTC()
}, func(a, b labels.Selector) bool {
	return a.String() == b.String()
}, func(a, b fields.Selector) bool {
	return a.String() == b.String()
})
var standardResourceQuotaScopes = sets.NewString(string(core.ResourceQuotaScopeTerminating), string(core.ResourceQuotaScopeNotTerminating), string(core.ResourceQuotaScopeBestEffort), string(core.ResourceQuotaScopeNotBestEffort), string(core.ResourceQuotaScopePriorityClass))

func IsStandardResourceQuotaScope(str string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return standardResourceQuotaScopes.Has(str)
}

var podObjectCountQuotaResources = sets.NewString(string(core.ResourcePods))
var podComputeQuotaResources = sets.NewString(string(core.ResourceCPU), string(core.ResourceMemory), string(core.ResourceLimitsCPU), string(core.ResourceLimitsMemory), string(core.ResourceRequestsCPU), string(core.ResourceRequestsMemory))

func IsResourceQuotaScopeValidForResource(scope core.ResourceQuotaScope, resource string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch scope {
	case core.ResourceQuotaScopeTerminating, core.ResourceQuotaScopeNotTerminating, core.ResourceQuotaScopeNotBestEffort, core.ResourceQuotaScopePriorityClass:
		return podObjectCountQuotaResources.Has(resource) || podComputeQuotaResources.Has(resource)
	case core.ResourceQuotaScopeBestEffort:
		return podObjectCountQuotaResources.Has(resource)
	default:
		return true
	}
}

var standardContainerResources = sets.NewString(string(core.ResourceCPU), string(core.ResourceMemory), string(core.ResourceEphemeralStorage))

func IsStandardContainerResourceName(str string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return standardContainerResources.Has(str) || IsHugePageResourceName(core.ResourceName(str))
}
func IsExtendedResourceName(name core.ResourceName) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if IsNativeResource(name) || strings.HasPrefix(string(name), core.DefaultResourceRequestsPrefix) {
		return false
	}
	nameForQuota := fmt.Sprintf("%s%s", core.DefaultResourceRequestsPrefix, string(name))
	if errs := validation.IsQualifiedName(string(nameForQuota)); len(errs) != 0 {
		return false
	}
	return true
}
func IsNativeResource(name core.ResourceName) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return !strings.Contains(string(name), "/") || strings.Contains(string(name), core.ResourceDefaultNamespacePrefix)
}
func IsOvercommitAllowed(name core.ResourceName) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return IsNativeResource(name) && !IsHugePageResourceName(name)
}

var standardLimitRangeTypes = sets.NewString(string(core.LimitTypePod), string(core.LimitTypeContainer), string(core.LimitTypePersistentVolumeClaim))

func IsStandardLimitRangeType(str string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return standardLimitRangeTypes.Has(str)
}

var standardQuotaResources = sets.NewString(string(core.ResourceCPU), string(core.ResourceMemory), string(core.ResourceEphemeralStorage), string(core.ResourceRequestsCPU), string(core.ResourceRequestsMemory), string(core.ResourceRequestsStorage), string(core.ResourceRequestsEphemeralStorage), string(core.ResourceLimitsCPU), string(core.ResourceLimitsMemory), string(core.ResourceLimitsEphemeralStorage), string(core.ResourcePods), string(core.ResourceQuotas), string(core.ResourceServices), string(core.ResourceReplicationControllers), string(core.ResourceSecrets), string(core.ResourcePersistentVolumeClaims), string(core.ResourceConfigMaps), string(core.ResourceServicesNodePorts), string(core.ResourceServicesLoadBalancers))

func IsStandardQuotaResourceName(str string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return standardQuotaResources.Has(str) || IsQuotaHugePageResourceName(core.ResourceName(str))
}

var standardResources = sets.NewString(string(core.ResourceCPU), string(core.ResourceMemory), string(core.ResourceEphemeralStorage), string(core.ResourceRequestsCPU), string(core.ResourceRequestsMemory), string(core.ResourceRequestsEphemeralStorage), string(core.ResourceLimitsCPU), string(core.ResourceLimitsMemory), string(core.ResourceLimitsEphemeralStorage), string(core.ResourcePods), string(core.ResourceQuotas), string(core.ResourceServices), string(core.ResourceReplicationControllers), string(core.ResourceSecrets), string(core.ResourceConfigMaps), string(core.ResourcePersistentVolumeClaims), string(core.ResourceStorage), string(core.ResourceRequestsStorage), string(core.ResourceServicesNodePorts), string(core.ResourceServicesLoadBalancers))

func IsStandardResourceName(str string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return standardResources.Has(str) || IsQuotaHugePageResourceName(core.ResourceName(str))
}

var integerResources = sets.NewString(string(core.ResourcePods), string(core.ResourceQuotas), string(core.ResourceServices), string(core.ResourceReplicationControllers), string(core.ResourceSecrets), string(core.ResourceConfigMaps), string(core.ResourcePersistentVolumeClaims), string(core.ResourceServicesNodePorts), string(core.ResourceServicesLoadBalancers))

func IsIntegerResourceName(str string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return integerResources.Has(str) || IsExtendedResourceName(core.ResourceName(str))
}
func IsServiceIPSet(service *core.Service) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return service.Spec.ClusterIP != core.ClusterIPNone && service.Spec.ClusterIP != ""
}

var standardFinalizers = sets.NewString(string(core.FinalizerKubernetes), metav1.FinalizerOrphanDependents, metav1.FinalizerDeleteDependents)

func IsStandardFinalizerName(str string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return standardFinalizers.Has(str)
}
func AddToNodeAddresses(addresses *[]core.NodeAddress, addAddresses ...core.NodeAddress) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, add := range addAddresses {
		exists := false
		for _, existing := range *addresses {
			if existing.Address == add.Address && existing.Type == add.Type {
				exists = true
				break
			}
		}
		if !exists {
			*addresses = append(*addresses, add)
		}
	}
}
func LoadBalancerStatusEqual(l, r *core.LoadBalancerStatus) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ingressSliceEqual(l.Ingress, r.Ingress)
}
func ingressSliceEqual(lhs, rhs []core.LoadBalancerIngress) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(lhs) != len(rhs) {
		return false
	}
	for i := range lhs {
		if !ingressEqual(&lhs[i], &rhs[i]) {
			return false
		}
	}
	return true
}
func ingressEqual(lhs, rhs *core.LoadBalancerIngress) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if lhs.IP != rhs.IP {
		return false
	}
	if lhs.Hostname != rhs.Hostname {
		return false
	}
	return true
}
func GetAccessModesAsString(modes []core.PersistentVolumeAccessMode) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	modes = removeDuplicateAccessModes(modes)
	modesStr := []string{}
	if containsAccessMode(modes, core.ReadWriteOnce) {
		modesStr = append(modesStr, "RWO")
	}
	if containsAccessMode(modes, core.ReadOnlyMany) {
		modesStr = append(modesStr, "ROX")
	}
	if containsAccessMode(modes, core.ReadWriteMany) {
		modesStr = append(modesStr, "RWX")
	}
	return strings.Join(modesStr, ",")
}
func GetAccessModesFromString(modes string) []core.PersistentVolumeAccessMode {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	strmodes := strings.Split(modes, ",")
	accessModes := []core.PersistentVolumeAccessMode{}
	for _, s := range strmodes {
		s = strings.Trim(s, " ")
		switch {
		case s == "RWO":
			accessModes = append(accessModes, core.ReadWriteOnce)
		case s == "ROX":
			accessModes = append(accessModes, core.ReadOnlyMany)
		case s == "RWX":
			accessModes = append(accessModes, core.ReadWriteMany)
		}
	}
	return accessModes
}
func removeDuplicateAccessModes(modes []core.PersistentVolumeAccessMode) []core.PersistentVolumeAccessMode {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	accessModes := []core.PersistentVolumeAccessMode{}
	for _, m := range modes {
		if !containsAccessMode(accessModes, m) {
			accessModes = append(accessModes, m)
		}
	}
	return accessModes
}
func containsAccessMode(modes []core.PersistentVolumeAccessMode, mode core.PersistentVolumeAccessMode) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, m := range modes {
		if m == mode {
			return true
		}
	}
	return false
}
func NodeSelectorRequirementsAsSelector(nsm []core.NodeSelectorRequirement) (labels.Selector, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(nsm) == 0 {
		return labels.Nothing(), nil
	}
	selector := labels.NewSelector()
	for _, expr := range nsm {
		var op selection.Operator
		switch expr.Operator {
		case core.NodeSelectorOpIn:
			op = selection.In
		case core.NodeSelectorOpNotIn:
			op = selection.NotIn
		case core.NodeSelectorOpExists:
			op = selection.Exists
		case core.NodeSelectorOpDoesNotExist:
			op = selection.DoesNotExist
		case core.NodeSelectorOpGt:
			op = selection.GreaterThan
		case core.NodeSelectorOpLt:
			op = selection.LessThan
		default:
			return nil, fmt.Errorf("%q is not a valid node selector operator", expr.Operator)
		}
		r, err := labels.NewRequirement(expr.Key, op, expr.Values)
		if err != nil {
			return nil, err
		}
		selector = selector.Add(*r)
	}
	return selector, nil
}
func NodeSelectorRequirementsAsFieldSelector(nsm []core.NodeSelectorRequirement) (fields.Selector, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(nsm) == 0 {
		return fields.Nothing(), nil
	}
	selectors := []fields.Selector{}
	for _, expr := range nsm {
		switch expr.Operator {
		case core.NodeSelectorOpIn:
			if len(expr.Values) != 1 {
				return nil, fmt.Errorf("unexpected number of value (%d) for node field selector operator %q", len(expr.Values), expr.Operator)
			}
			selectors = append(selectors, fields.OneTermEqualSelector(expr.Key, expr.Values[0]))
		case core.NodeSelectorOpNotIn:
			if len(expr.Values) != 1 {
				return nil, fmt.Errorf("unexpected number of value (%d) for node field selector operator %q", len(expr.Values), expr.Operator)
			}
			selectors = append(selectors, fields.OneTermNotEqualSelector(expr.Key, expr.Values[0]))
		default:
			return nil, fmt.Errorf("%q is not a valid node field selector operator", expr.Operator)
		}
	}
	return fields.AndSelectors(selectors...), nil
}
func GetTolerationsFromPodAnnotations(annotations map[string]string) ([]core.Toleration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var tolerations []core.Toleration
	if len(annotations) > 0 && annotations[core.TolerationsAnnotationKey] != "" {
		err := json.Unmarshal([]byte(annotations[core.TolerationsAnnotationKey]), &tolerations)
		if err != nil {
			return tolerations, err
		}
	}
	return tolerations, nil
}
func AddOrUpdateTolerationInPod(pod *core.Pod, toleration *core.Toleration) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podTolerations := pod.Spec.Tolerations
	var newTolerations []core.Toleration
	updated := false
	for i := range podTolerations {
		if toleration.MatchToleration(&podTolerations[i]) {
			if Semantic.DeepEqual(toleration, podTolerations[i]) {
				return false
			}
			newTolerations = append(newTolerations, *toleration)
			updated = true
			continue
		}
		newTolerations = append(newTolerations, podTolerations[i])
	}
	if !updated {
		newTolerations = append(newTolerations, *toleration)
	}
	pod.Spec.Tolerations = newTolerations
	return true
}
func GetTaintsFromNodeAnnotations(annotations map[string]string) ([]core.Taint, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var taints []core.Taint
	if len(annotations) > 0 && annotations[core.TaintsAnnotationKey] != "" {
		err := json.Unmarshal([]byte(annotations[core.TaintsAnnotationKey]), &taints)
		if err != nil {
			return []core.Taint{}, err
		}
	}
	return taints, nil
}
func GetPersistentVolumeClass(volume *core.PersistentVolume) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if class, found := volume.Annotations[core.BetaStorageClassAnnotation]; found {
		return class
	}
	return volume.Spec.StorageClassName
}
func GetPersistentVolumeClaimClass(claim *core.PersistentVolumeClaim) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if class, found := claim.Annotations[core.BetaStorageClassAnnotation]; found {
		return class
	}
	if claim.Spec.StorageClassName != nil {
		return *claim.Spec.StorageClassName
	}
	return ""
}
func PersistentVolumeClaimHasClass(claim *core.PersistentVolumeClaim) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, found := claim.Annotations[core.BetaStorageClassAnnotation]; found {
		return true
	}
	if claim.Spec.StorageClassName != nil {
		return true
	}
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

package limitranger

import (
	"fmt"
	goformat "fmt"
	"github.com/hashicorp/golang-lru"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitailizer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	api "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	"strings"
	"time"
	gotime "time"
)

const (
	limitRangerAnnotation = "kubernetes.io/limit-ranger"
	PluginName            = "LimitRanger"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewLimitRanger(&DefaultLimitRangerActions{})
	})
}

type LimitRanger struct {
	*admission.Handler
	client          kubernetes.Interface
	actions         LimitRangerActions
	lister          corev1listers.LimitRangeLister
	liveLookupCache *lru.Cache
	liveTTL         time.Duration
}

var _ admission.MutationInterface = &LimitRanger{}
var _ admission.ValidationInterface = &LimitRanger{}
var _ genericadmissioninitailizer.WantsExternalKubeInformerFactory = &LimitRanger{}
var _ genericadmissioninitailizer.WantsExternalKubeClientSet = &LimitRanger{}

type liveLookupEntry struct {
	expiry time.Time
	items  []*corev1.LimitRange
}

func (l *LimitRanger) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	limitRangeInformer := f.Core().V1().LimitRanges()
	l.SetReadyFunc(limitRangeInformer.Informer().HasSynced)
	l.lister = limitRangeInformer.Lister()
}
func (a *LimitRanger) SetExternalKubeClientSet(client kubernetes.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.client = client
}
func (l *LimitRanger) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if l.lister == nil {
		return fmt.Errorf("missing limitRange lister")
	}
	if l.client == nil {
		return fmt.Errorf("missing client")
	}
	return nil
}
func (l *LimitRanger) Admit(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return l.runLimitFunc(a, l.actions.MutateLimit)
}
func (l *LimitRanger) Validate(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return l.runLimitFunc(a, l.actions.ValidateLimit)
}
func (l *LimitRanger) runLimitFunc(a admission.Attributes, limitFn func(limitRange *corev1.LimitRange, kind string, obj runtime.Object) error) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !l.actions.SupportsAttributes(a) {
		return nil
	}
	oldObj := a.GetOldObject()
	if oldObj != nil {
		oldAccessor, err := meta.Accessor(oldObj)
		if err != nil {
			return admission.NewForbidden(a, err)
		}
		if oldAccessor.GetDeletionTimestamp() != nil {
			return nil
		}
	}
	items, err := l.GetLimitRanges(a)
	if err != nil {
		return err
	}
	for i := range items {
		limitRange := items[i]
		if !l.actions.SupportsLimit(limitRange) {
			continue
		}
		err = limitFn(limitRange, a.GetResource().Resource, a.GetObject())
		if err != nil {
			return admission.NewForbidden(a, err)
		}
	}
	return nil
}
func (l *LimitRanger) GetLimitRanges(a admission.Attributes) ([]*corev1.LimitRange, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	items, err := l.lister.LimitRanges(a.GetNamespace()).List(labels.Everything())
	if err != nil {
		return nil, admission.NewForbidden(a, fmt.Errorf("unable to %s %v at this time because there was an error enforcing limit ranges", a.GetOperation(), a.GetResource()))
	}
	if len(items) == 0 {
		lruItemObj, ok := l.liveLookupCache.Get(a.GetNamespace())
		if !ok || lruItemObj.(liveLookupEntry).expiry.Before(time.Now()) {
			liveList, err := l.client.CoreV1().LimitRanges(a.GetNamespace()).List(metav1.ListOptions{})
			if err != nil {
				return nil, admission.NewForbidden(a, err)
			}
			newEntry := liveLookupEntry{expiry: time.Now().Add(l.liveTTL)}
			for i := range liveList.Items {
				newEntry.items = append(newEntry.items, &liveList.Items[i])
			}
			l.liveLookupCache.Add(a.GetNamespace(), newEntry)
			lruItemObj = newEntry
		}
		lruEntry := lruItemObj.(liveLookupEntry)
		for i := range lruEntry.items {
			items = append(items, lruEntry.items[i])
		}
	}
	return items, nil
}
func NewLimitRanger(actions LimitRangerActions) (*LimitRanger, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	liveLookupCache, err := lru.New(10000)
	if err != nil {
		return nil, err
	}
	if actions == nil {
		actions = &DefaultLimitRangerActions{}
	}
	return &LimitRanger{Handler: admission.NewHandler(admission.Create, admission.Update), actions: actions, liveLookupCache: liveLookupCache, liveTTL: time.Duration(30 * time.Second)}, nil
}
func defaultContainerResourceRequirements(limitRange *corev1.LimitRange) api.ResourceRequirements {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requirements := api.ResourceRequirements{}
	requirements.Requests = api.ResourceList{}
	requirements.Limits = api.ResourceList{}
	for i := range limitRange.Spec.Limits {
		limit := limitRange.Spec.Limits[i]
		if limit.Type == corev1.LimitTypeContainer {
			for k, v := range limit.DefaultRequest {
				value := v.Copy()
				requirements.Requests[api.ResourceName(k)] = *value
			}
			for k, v := range limit.Default {
				value := v.Copy()
				requirements.Limits[api.ResourceName(k)] = *value
			}
		}
	}
	return requirements
}
func mergeContainerResources(container *api.Container, defaultRequirements *api.ResourceRequirements, annotationPrefix string, annotations []string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	setRequests := []string{}
	setLimits := []string{}
	if container.Resources.Limits == nil {
		container.Resources.Limits = api.ResourceList{}
	}
	if container.Resources.Requests == nil {
		container.Resources.Requests = api.ResourceList{}
	}
	for k, v := range defaultRequirements.Limits {
		_, found := container.Resources.Limits[k]
		if !found {
			container.Resources.Limits[k] = *v.Copy()
			setLimits = append(setLimits, string(k))
		}
	}
	for k, v := range defaultRequirements.Requests {
		_, found := container.Resources.Requests[k]
		if !found {
			container.Resources.Requests[k] = *v.Copy()
			setRequests = append(setRequests, string(k))
		}
	}
	if len(setRequests) > 0 {
		sort.Strings(setRequests)
		a := strings.Join(setRequests, ", ") + fmt.Sprintf(" request for %s %s", annotationPrefix, container.Name)
		annotations = append(annotations, a)
	}
	if len(setLimits) > 0 {
		sort.Strings(setLimits)
		a := strings.Join(setLimits, ", ") + fmt.Sprintf(" limit for %s %s", annotationPrefix, container.Name)
		annotations = append(annotations, a)
	}
	return annotations
}
func mergePodResourceRequirements(pod *api.Pod, defaultRequirements *api.ResourceRequirements) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	annotations := []string{}
	for i := range pod.Spec.Containers {
		annotations = mergeContainerResources(&pod.Spec.Containers[i], defaultRequirements, "container", annotations)
	}
	for i := range pod.Spec.InitContainers {
		annotations = mergeContainerResources(&pod.Spec.InitContainers[i], defaultRequirements, "init container", annotations)
	}
	if len(annotations) > 0 {
		if pod.ObjectMeta.Annotations == nil {
			pod.ObjectMeta.Annotations = make(map[string]string)
		}
		val := "LimitRanger plugin set: " + strings.Join(annotations, "; ")
		pod.ObjectMeta.Annotations[limitRangerAnnotation] = val
	}
}
func requestLimitEnforcedValues(requestQuantity, limitQuantity, enforcedQuantity resource.Quantity) (request, limit, enforced int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	request = requestQuantity.Value()
	limit = limitQuantity.Value()
	enforced = enforcedQuantity.Value()
	if request <= resource.MaxMilliValue && limit <= resource.MaxMilliValue && enforced <= resource.MaxMilliValue {
		request = requestQuantity.MilliValue()
		limit = limitQuantity.MilliValue()
		enforced = enforcedQuantity.MilliValue()
	}
	return
}
func minConstraint(limitType string, resourceName string, enforced resource.Quantity, request api.ResourceList, limit api.ResourceList) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, reqExists := request[api.ResourceName(resourceName)]
	lim, limExists := limit[api.ResourceName(resourceName)]
	observedReqValue, observedLimValue, enforcedValue := requestLimitEnforcedValues(req, lim, enforced)
	if !reqExists {
		return fmt.Errorf("minimum %s usage per %s is %s.  No request is specified.", resourceName, limitType, enforced.String())
	}
	if observedReqValue < enforcedValue {
		return fmt.Errorf("minimum %s usage per %s is %s, but request is %s.", resourceName, limitType, enforced.String(), req.String())
	}
	if limExists && (observedLimValue < enforcedValue) {
		return fmt.Errorf("minimum %s usage per %s is %s, but limit is %s.", resourceName, limitType, enforced.String(), lim.String())
	}
	return nil
}
func maxRequestConstraint(limitType string, resourceName string, enforced resource.Quantity, request api.ResourceList) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, reqExists := request[api.ResourceName(resourceName)]
	observedReqValue, _, enforcedValue := requestLimitEnforcedValues(req, resource.Quantity{}, enforced)
	if !reqExists {
		return fmt.Errorf("maximum %s usage per %s is %s.  No request is specified.", resourceName, limitType, enforced.String())
	}
	if observedReqValue > enforcedValue {
		return fmt.Errorf("maximum %s usage per %s is %s, but request is %s.", resourceName, limitType, enforced.String(), req.String())
	}
	return nil
}
func maxConstraint(limitType string, resourceName string, enforced resource.Quantity, request api.ResourceList, limit api.ResourceList) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, reqExists := request[api.ResourceName(resourceName)]
	lim, limExists := limit[api.ResourceName(resourceName)]
	observedReqValue, observedLimValue, enforcedValue := requestLimitEnforcedValues(req, lim, enforced)
	if !limExists {
		return fmt.Errorf("maximum %s usage per %s is %s.  No limit is specified.", resourceName, limitType, enforced.String())
	}
	if observedLimValue > enforcedValue {
		return fmt.Errorf("maximum %s usage per %s is %s, but limit is %s.", resourceName, limitType, enforced.String(), lim.String())
	}
	if reqExists && (observedReqValue > enforcedValue) {
		return fmt.Errorf("maximum %s usage per %s is %s, but request is %s.", resourceName, limitType, enforced.String(), req.String())
	}
	return nil
}
func limitRequestRatioConstraint(limitType string, resourceName string, enforced resource.Quantity, request api.ResourceList, limit api.ResourceList) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, reqExists := request[api.ResourceName(resourceName)]
	lim, limExists := limit[api.ResourceName(resourceName)]
	observedReqValue, observedLimValue, _ := requestLimitEnforcedValues(req, lim, enforced)
	if !reqExists || (observedReqValue == int64(0)) {
		return fmt.Errorf("%s max limit to request ratio per %s is %s, but no request is specified or request is 0.", resourceName, limitType, enforced.String())
	}
	if !limExists || (observedLimValue == int64(0)) {
		return fmt.Errorf("%s max limit to request ratio per %s is %s, but no limit is specified or limit is 0.", resourceName, limitType, enforced.String())
	}
	observedRatio := float64(observedLimValue) / float64(observedReqValue)
	displayObservedRatio := observedRatio
	maxLimitRequestRatio := float64(enforced.Value())
	if enforced.Value() <= resource.MaxMilliValue {
		observedRatio = observedRatio * 1000
		maxLimitRequestRatio = float64(enforced.MilliValue())
	}
	if observedRatio > maxLimitRequestRatio {
		return fmt.Errorf("%s max limit to request ratio per %s is %s, but provided ratio is %f.", resourceName, limitType, enforced.String(), displayObservedRatio)
	}
	return nil
}
func sum(inputs []api.ResourceList) api.ResourceList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := api.ResourceList{}
	keys := []api.ResourceName{}
	for i := range inputs {
		for k := range inputs[i] {
			keys = append(keys, k)
		}
	}
	for _, key := range keys {
		total, isSet := int64(0), true
		for i := range inputs {
			input := inputs[i]
			v, exists := input[key]
			if exists {
				if key == api.ResourceCPU {
					total = total + v.MilliValue()
				} else {
					total = total + v.Value()
				}
			} else {
				isSet = false
			}
		}
		if isSet {
			if key == api.ResourceCPU {
				result[key] = *(resource.NewMilliQuantity(total, resource.DecimalSI))
			} else {
				result[key] = *(resource.NewQuantity(total, resource.DecimalSI))
			}
		}
	}
	return result
}

type DefaultLimitRangerActions struct{}

var _ LimitRangerActions = &DefaultLimitRangerActions{}

func (d *DefaultLimitRangerActions) MutateLimit(limitRange *corev1.LimitRange, resourceName string, obj runtime.Object) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch resourceName {
	case "pods":
		return PodMutateLimitFunc(limitRange, obj.(*api.Pod))
	}
	return nil
}
func (d *DefaultLimitRangerActions) ValidateLimit(limitRange *corev1.LimitRange, resourceName string, obj runtime.Object) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch resourceName {
	case "pods":
		return PodValidateLimitFunc(limitRange, obj.(*api.Pod))
	case "persistentvolumeclaims":
		return PersistentVolumeClaimValidateLimitFunc(limitRange, obj.(*api.PersistentVolumeClaim))
	}
	return nil
}
func (d *DefaultLimitRangerActions) SupportsAttributes(a admission.Attributes) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetSubresource() != "" {
		return false
	}
	if a.GetKind().GroupKind() == api.Kind("Pod") && a.GetOperation() == admission.Update {
		return false
	}
	return a.GetKind().GroupKind() == api.Kind("Pod") || a.GetKind().GroupKind() == api.Kind("PersistentVolumeClaim")
}
func (d *DefaultLimitRangerActions) SupportsLimit(limitRange *corev1.LimitRange) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func PersistentVolumeClaimValidateLimitFunc(limitRange *corev1.LimitRange, pvc *api.PersistentVolumeClaim) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var errs []error
	for i := range limitRange.Spec.Limits {
		limit := limitRange.Spec.Limits[i]
		limitType := limit.Type
		if limitType == corev1.LimitTypePersistentVolumeClaim {
			for k, v := range limit.Min {
				if err := minConstraint(string(limitType), string(k), v, pvc.Spec.Resources.Requests, api.ResourceList{}); err != nil {
					errs = append(errs, err)
				}
			}
			for k, v := range limit.Max {
				if err := maxRequestConstraint(string(limitType), string(k), v, pvc.Spec.Resources.Requests); err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	return utilerrors.NewAggregate(errs)
}
func PodMutateLimitFunc(limitRange *corev1.LimitRange, pod *api.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defaultResources := defaultContainerResourceRequirements(limitRange)
	mergePodResourceRequirements(pod, &defaultResources)
	return nil
}
func PodValidateLimitFunc(limitRange *corev1.LimitRange, pod *api.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var errs []error
	for i := range limitRange.Spec.Limits {
		limit := limitRange.Spec.Limits[i]
		limitType := limit.Type
		if limitType == corev1.LimitTypeContainer {
			for j := range pod.Spec.Containers {
				container := &pod.Spec.Containers[j]
				for k, v := range limit.Min {
					if err := minConstraint(string(limitType), string(k), v, container.Resources.Requests, container.Resources.Limits); err != nil {
						errs = append(errs, err)
					}
				}
				for k, v := range limit.Max {
					if err := maxConstraint(string(limitType), string(k), v, container.Resources.Requests, container.Resources.Limits); err != nil {
						errs = append(errs, err)
					}
				}
				for k, v := range limit.MaxLimitRequestRatio {
					if err := limitRequestRatioConstraint(string(limitType), string(k), v, container.Resources.Requests, container.Resources.Limits); err != nil {
						errs = append(errs, err)
					}
				}
			}
			for j := range pod.Spec.InitContainers {
				container := &pod.Spec.InitContainers[j]
				for k, v := range limit.Min {
					if err := minConstraint(string(limitType), string(k), v, container.Resources.Requests, container.Resources.Limits); err != nil {
						errs = append(errs, err)
					}
				}
				for k, v := range limit.Max {
					if err := maxConstraint(string(limitType), string(k), v, container.Resources.Requests, container.Resources.Limits); err != nil {
						errs = append(errs, err)
					}
				}
				for k, v := range limit.MaxLimitRequestRatio {
					if err := limitRequestRatioConstraint(string(limitType), string(k), v, container.Resources.Requests, container.Resources.Limits); err != nil {
						errs = append(errs, err)
					}
				}
			}
		}
		if limitType == corev1.LimitTypePod {
			containerRequests, containerLimits := []api.ResourceList{}, []api.ResourceList{}
			for j := range pod.Spec.Containers {
				container := &pod.Spec.Containers[j]
				containerRequests = append(containerRequests, container.Resources.Requests)
				containerLimits = append(containerLimits, container.Resources.Limits)
			}
			podRequests := sum(containerRequests)
			podLimits := sum(containerLimits)
			for j := range pod.Spec.InitContainers {
				container := &pod.Spec.InitContainers[j]
				for k, v := range container.Resources.Requests {
					if v2, ok := podRequests[k]; ok {
						if v.Cmp(v2) > 0 {
							podRequests[k] = v
						}
					} else {
						podRequests[k] = v
					}
				}
				for k, v := range container.Resources.Limits {
					if v2, ok := podLimits[k]; ok {
						if v.Cmp(v2) > 0 {
							podLimits[k] = v
						}
					} else {
						podLimits[k] = v
					}
				}
			}
			for k, v := range limit.Min {
				if err := minConstraint(string(limitType), string(k), v, podRequests, podLimits); err != nil {
					errs = append(errs, err)
				}
			}
			for k, v := range limit.Max {
				if err := maxConstraint(string(limitType), string(k), v, podRequests, podLimits); err != nil {
					errs = append(errs, err)
				}
			}
			for k, v := range limit.MaxLimitRequestRatio {
				if err := limitRequestRatioConstraint(string(limitType), string(k), v, podRequests, podLimits); err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	return utilerrors.NewAggregate(errs)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

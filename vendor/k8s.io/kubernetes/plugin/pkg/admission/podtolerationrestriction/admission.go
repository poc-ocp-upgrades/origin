package podtolerationrestriction

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	"io"
	"k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	qoshelper "k8s.io/kubernetes/pkg/apis/core/helper/qos"
	k8s_api_v1 "k8s.io/kubernetes/pkg/apis/core/v1"
	"k8s.io/kubernetes/pkg/kubeapiserver/admission/util"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	"k8s.io/kubernetes/pkg/util/tolerations"
	pluginapi "k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "PodTolerationRestriction"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		pluginConfig, err := loadConfiguration(config)
		if err != nil {
			return nil, err
		}
		return NewPodTolerationsPlugin(pluginConfig), nil
	})
}

const (
	NSDefaultTolerations string = "scheduler.alpha.kubernetes.io/defaultTolerations"
	NSWLTolerations      string = "scheduler.alpha.kubernetes.io/tolerationsWhitelist"
)

var _ admission.MutationInterface = &podTolerationsPlugin{}
var _ admission.ValidationInterface = &podTolerationsPlugin{}
var _ = genericadmissioninitializer.WantsExternalKubeInformerFactory(&podTolerationsPlugin{})
var _ = genericadmissioninitializer.WantsExternalKubeClientSet(&podTolerationsPlugin{})

type podTolerationsPlugin struct {
	*admission.Handler
	client          kubernetes.Interface
	namespaceLister corev1listers.NamespaceLister
	pluginConfig    *pluginapi.Configuration
}

func (p *podTolerationsPlugin) Admit(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if shouldIgnore(a) {
		return nil
	}
	if !p.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}
	pod := a.GetObject().(*api.Pod)
	var finalTolerations []api.Toleration
	updateUninitialized, err := util.IsUpdatingUninitializedObject(a)
	if err != nil {
		return err
	}
	if a.GetOperation() == admission.Create || updateUninitialized {
		ts, err := p.getNamespaceDefaultTolerations(a.GetNamespace())
		if err != nil {
			return err
		}
		if ts == nil {
			ts = p.pluginConfig.Default
		}
		if len(ts) > 0 {
			if len(pod.Spec.Tolerations) > 0 {
				if tolerations.IsConflict(ts, pod.Spec.Tolerations) {
					return fmt.Errorf("namespace tolerations and pod tolerations conflict")
				}
				finalTolerations = tolerations.MergeTolerations(ts, pod.Spec.Tolerations)
			} else {
				finalTolerations = ts
			}
		} else {
			finalTolerations = pod.Spec.Tolerations
		}
	} else {
		finalTolerations = pod.Spec.Tolerations
	}
	if qoshelper.GetPodQOS(pod) != api.PodQOSBestEffort {
		finalTolerations = tolerations.MergeTolerations(finalTolerations, []api.Toleration{{Key: schedulerapi.TaintNodeMemoryPressure, Operator: api.TolerationOpExists, Effect: api.TaintEffectNoSchedule}})
	}
	pod.Spec.Tolerations = tolerations.MergeTolerations(finalTolerations, []api.Toleration{})
	return p.Validate(a)
}
func (p *podTolerationsPlugin) Validate(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if shouldIgnore(a) {
		return nil
	}
	if !p.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}
	pod := a.GetObject().(*api.Pod)
	if len(pod.Spec.Tolerations) > 0 {
		whitelist, err := p.getNamespaceTolerationsWhitelist(a.GetNamespace())
		if err != nil {
			return err
		}
		if whitelist == nil {
			whitelist = p.pluginConfig.Whitelist
		}
		if len(whitelist) > 0 {
			if !tolerations.VerifyAgainstWhitelist(pod.Spec.Tolerations, whitelist) {
				return fmt.Errorf("pod tolerations (possibly merged with namespace default tolerations) conflict with its namespace whitelist")
			}
		}
	}
	return nil
}
func shouldIgnore(a admission.Attributes) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resource := a.GetResource().GroupResource()
	if resource != api.Resource("pods") {
		return true
	}
	if a.GetSubresource() != "" {
		return true
	}
	obj := a.GetObject()
	_, ok := obj.(*api.Pod)
	if !ok {
		klog.Errorf("expected pod but got %s", a.GetKind().Kind)
		return true
	}
	return false
}
func NewPodTolerationsPlugin(pluginConfig *pluginapi.Configuration) *podTolerationsPlugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &podTolerationsPlugin{Handler: admission.NewHandler(admission.Create, admission.Update), pluginConfig: pluginConfig}
}
func (a *podTolerationsPlugin) SetExternalKubeClientSet(client kubernetes.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.client = client
}
func (p *podTolerationsPlugin) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespaceInformer := f.Core().V1().Namespaces()
	p.namespaceLister = namespaceInformer.Lister()
	p.SetReadyFunc(namespaceInformer.Informer().HasSynced)
}
func (p *podTolerationsPlugin) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if p.namespaceLister == nil {
		return fmt.Errorf("missing namespaceLister")
	}
	if p.client == nil {
		return fmt.Errorf("missing client")
	}
	return nil
}
func (p *podTolerationsPlugin) getNamespace(nsName string) (*corev1.Namespace, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespace, err := p.namespaceLister.Get(nsName)
	if errors.IsNotFound(err) {
		namespace, err = p.client.CoreV1().Namespaces().Get(nsName, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return nil, err
			}
			return nil, errors.NewInternalError(err)
		}
	} else if err != nil {
		return nil, errors.NewInternalError(err)
	}
	return namespace, nil
}
func (p *podTolerationsPlugin) getNamespaceDefaultTolerations(nsName string) ([]api.Toleration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ns, err := p.getNamespace(nsName)
	if err != nil {
		return nil, err
	}
	return extractNSTolerations(ns, NSDefaultTolerations)
}
func (p *podTolerationsPlugin) getNamespaceTolerationsWhitelist(nsName string) ([]api.Toleration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ns, err := p.getNamespace(nsName)
	if err != nil {
		return nil, err
	}
	return extractNSTolerations(ns, NSWLTolerations)
}
func extractNSTolerations(ns *corev1.Namespace, key string) ([]api.Toleration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(ns.Annotations) == 0 {
		return nil, nil
	}
	if _, ok := ns.Annotations[key]; !ok {
		return nil, nil
	}
	if len(ns.Annotations[key]) == 0 {
		return []api.Toleration{}, nil
	}
	var v1Tolerations []v1.Toleration
	err := json.Unmarshal([]byte(ns.Annotations[key]), &v1Tolerations)
	if err != nil {
		return nil, err
	}
	ts := make([]api.Toleration, len(v1Tolerations))
	for i := range v1Tolerations {
		if err := k8s_api_v1.Convert_v1_Toleration_To_core_Toleration(&v1Tolerations[i], &ts[i], nil); err != nil {
			return nil, err
		}
	}
	return ts, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

package podnodeselector

import (
	"fmt"
	goformat "fmt"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/kubeapiserver/admission/util"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	gotime "time"
)

var NamespaceNodeSelectors = []string{"scheduler.alpha.kubernetes.io/node-selector"}

const PluginName = "PodNodeSelector"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		pluginConfig := readConfig(config)
		plugin := NewPodNodeSelector(pluginConfig.PodNodeSelectorPluginConfig)
		return plugin, nil
	})
}

type podNodeSelector struct {
	*admission.Handler
	client               kubernetes.Interface
	namespaceLister      corev1listers.NamespaceLister
	clusterNodeSelectors map[string]string
}

var _ = genericadmissioninitializer.WantsExternalKubeClientSet(&podNodeSelector{})
var _ = genericadmissioninitializer.WantsExternalKubeInformerFactory(&podNodeSelector{})

type pluginConfig struct{ PodNodeSelectorPluginConfig map[string]string }

func readConfig(config io.Reader) *pluginConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defaultConfig := &pluginConfig{}
	if config == nil || reflect.ValueOf(config).IsNil() {
		return defaultConfig
	}
	d := yaml.NewYAMLOrJSONDecoder(config, 4096)
	for {
		if err := d.Decode(defaultConfig); err != nil {
			if err != io.EOF {
				continue
			}
		}
		break
	}
	return defaultConfig
}
func (p *podNodeSelector) Admit(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if shouldIgnore(a) {
		return nil
	}
	updateInitialized, err := util.IsUpdatingInitializedObject(a)
	if err != nil {
		return err
	}
	if updateInitialized {
		return nil
	}
	if !p.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}
	resource := a.GetResource().GroupResource()
	pod := a.GetObject().(*api.Pod)
	namespaceNodeSelector, err := p.getNamespaceNodeSelectorMap(a.GetNamespace())
	if err != nil {
		return err
	}
	if labels.Conflicts(namespaceNodeSelector, labels.Set(pod.Spec.NodeSelector)) {
		return errors.NewForbidden(resource, pod.Name, fmt.Errorf("pod node label selector conflicts with its namespace node label selector"))
	}
	podNodeSelectorLabels := labels.Merge(namespaceNodeSelector, pod.Spec.NodeSelector)
	pod.Spec.NodeSelector = map[string]string(podNodeSelectorLabels)
	return p.Validate(a)
}
func (p *podNodeSelector) Validate(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if shouldIgnore(a) {
		return nil
	}
	if !p.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}
	resource := a.GetResource().GroupResource()
	pod := a.GetObject().(*api.Pod)
	namespaceNodeSelector, err := p.getNamespaceNodeSelectorMap(a.GetNamespace())
	if err != nil {
		return err
	}
	if labels.Conflicts(namespaceNodeSelector, labels.Set(pod.Spec.NodeSelector)) {
		return errors.NewForbidden(resource, pod.Name, fmt.Errorf("pod node label selector conflicts with its namespace node label selector"))
	}
	whitelist, err := labels.ConvertSelectorToLabelsMap(p.clusterNodeSelectors[a.GetNamespace()])
	if err != nil {
		return err
	}
	if !labels.AreLabelsInWhiteList(pod.Spec.NodeSelector, whitelist) {
		return errors.NewForbidden(resource, pod.Name, fmt.Errorf("pod node label selector labels conflict with its namespace whitelist"))
	}
	return nil
}
func (p *podNodeSelector) getNamespaceNodeSelectorMap(namespaceName string) (labels.Set, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespace, err := p.namespaceLister.Get(namespaceName)
	if errors.IsNotFound(err) {
		namespace, err = p.defaultGetNamespace(namespaceName)
		if err != nil {
			if errors.IsNotFound(err) {
				return nil, err
			}
			return nil, errors.NewInternalError(err)
		}
	} else if err != nil {
		return nil, errors.NewInternalError(err)
	}
	return p.getNodeSelectorMap(namespace)
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
	_, ok := a.GetObject().(*api.Pod)
	if !ok {
		klog.Errorf("expected pod but got %s", a.GetKind().Kind)
		return true
	}
	return false
}
func NewPodNodeSelector(clusterNodeSelectors map[string]string) *podNodeSelector {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &podNodeSelector{Handler: admission.NewHandler(admission.Create, admission.Update), clusterNodeSelectors: clusterNodeSelectors}
}
func (a *podNodeSelector) SetExternalKubeClientSet(client kubernetes.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.client = client
}
func (p *podNodeSelector) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespaceInformer := f.Core().V1().Namespaces()
	p.namespaceLister = namespaceInformer.Lister()
	p.SetReadyFunc(namespaceInformer.Informer().HasSynced)
}
func (p *podNodeSelector) ValidateInitialization() error {
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
func (p *podNodeSelector) defaultGetNamespace(name string) (*corev1.Namespace, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespace, err := p.client.Core().Namespaces().Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("namespace %s does not exist", name)
	}
	return namespace, nil
}
func (p *podNodeSelector) getNodeSelectorMap(namespace *corev1.Namespace) (labels.Set, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	selector := labels.Set{}
	labelsMap := labels.Set{}
	var err error
	found := false
	if len(namespace.ObjectMeta.Annotations) > 0 {
		for _, annotation := range NamespaceNodeSelectors {
			if ns, ok := namespace.ObjectMeta.Annotations[annotation]; ok {
				labelsMap, err = labels.ConvertSelectorToLabelsMap(ns)
				if err != nil {
					return labels.Set{}, err
				}
				if labels.Conflicts(selector, labelsMap) {
					nsName := namespace.ObjectMeta.Name
					return labels.Set{}, fmt.Errorf("%s annotations' node label selectors conflict", nsName)
				}
				selector = labels.Merge(selector, labelsMap)
				found = true
			}
		}
	}
	if !found {
		selector, err = labels.ConvertSelectorToLabelsMap(p.clusterNodeSelectors["clusterDefaultNodeSelector"])
		if err != nil {
			return labels.Set{}, err
		}
	}
	return selector, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

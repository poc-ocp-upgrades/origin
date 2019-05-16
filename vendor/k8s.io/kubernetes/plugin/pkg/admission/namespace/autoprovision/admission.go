package autoprovision

import (
	"fmt"
	goformat "fmt"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	api "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "NamespaceAutoProvision"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewProvision(), nil
	})
}

type Provision struct {
	*admission.Handler
	client          kubernetes.Interface
	namespaceLister corev1listers.NamespaceLister
}

var _ admission.MutationInterface = &Provision{}
var _ = genericadmissioninitializer.WantsExternalKubeInformerFactory(&Provision{})
var _ = genericadmissioninitializer.WantsExternalKubeClientSet(&Provision{})

func (p *Provision) Admit(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.IsDryRun() {
		return nil
	}
	if len(a.GetNamespace()) == 0 || a.GetKind().GroupKind() == api.Kind("Namespace") {
		return nil
	}
	if !p.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}
	_, err := p.namespaceLister.Get(a.GetNamespace())
	if err == nil {
		return nil
	}
	if !errors.IsNotFound(err) {
		return admission.NewForbidden(a, err)
	}
	namespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: a.GetNamespace(), Namespace: ""}, Status: corev1.NamespaceStatus{}}
	_, err = p.client.Core().Namespaces().Create(namespace)
	if err != nil && !errors.IsAlreadyExists(err) {
		return admission.NewForbidden(a, err)
	}
	return nil
}
func NewProvision() *Provision {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Provision{Handler: admission.NewHandler(admission.Create)}
}
func (p *Provision) SetExternalKubeClientSet(client kubernetes.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p.client = client
}
func (p *Provision) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespaceInformer := f.Core().V1().Namespaces()
	p.namespaceLister = namespaceInformer.Lister()
	p.SetReadyFunc(namespaceInformer.Informer().HasSynced)
}
func (p *Provision) ValidateInitialization() error {
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

package resourcequota

import (
	"fmt"
	goformat "fmt"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	kubeapiserveradmission "k8s.io/kubernetes/pkg/kubeapiserver/admission"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	"k8s.io/kubernetes/pkg/quota/v1/generic"
	resourcequotaapi "k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota"
	"k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/validation"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

const PluginName = "ResourceQuota"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		configuration, err := LoadConfiguration(config)
		if err != nil {
			return nil, err
		}
		if configuration != nil {
			if errs := validation.ValidateConfiguration(configuration); len(errs) != 0 {
				return nil, errs.ToAggregate()
			}
		}
		return NewResourceQuota(configuration, 5, make(chan struct{}))
	})
}

type QuotaAdmission struct {
	*admission.Handler
	config             *resourcequotaapi.Configuration
	stopCh             <-chan struct{}
	quotaConfiguration quota.Configuration
	numEvaluators      int
	quotaAccessor      *quotaAccessor
	evaluator          Evaluator
}

var _ admission.ValidationInterface = &QuotaAdmission{}
var _ = genericadmissioninitializer.WantsExternalKubeInformerFactory(&QuotaAdmission{})
var _ = genericadmissioninitializer.WantsExternalKubeClientSet(&QuotaAdmission{})
var _ = kubeapiserveradmission.WantsQuotaConfiguration(&QuotaAdmission{})

type liveLookupEntry struct {
	expiry time.Time
	items  []*corev1.ResourceQuota
}

func NewResourceQuota(config *resourcequotaapi.Configuration, numEvaluators int, stopCh <-chan struct{}) (*QuotaAdmission, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	quotaAccessor, err := newQuotaAccessor()
	if err != nil {
		return nil, err
	}
	return &QuotaAdmission{Handler: admission.NewHandler(admission.Create, admission.Update), stopCh: stopCh, numEvaluators: numEvaluators, config: config, quotaAccessor: quotaAccessor}, nil
}
func (a *QuotaAdmission) SetExternalKubeClientSet(client kubernetes.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.quotaAccessor.client = client
}
func (a *QuotaAdmission) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.quotaAccessor.lister = f.Core().V1().ResourceQuotas().Lister()
}
func (a *QuotaAdmission) SetQuotaConfiguration(c quota.Configuration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.quotaConfiguration = c
	a.evaluator = NewQuotaEvaluator(a.quotaAccessor, a.quotaConfiguration.IgnoredResources(), generic.NewRegistry(a.quotaConfiguration.Evaluators()), nil, a.config, a.numEvaluators, a.stopCh)
}
func (a *QuotaAdmission) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.quotaAccessor == nil {
		return fmt.Errorf("missing quotaAccessor")
	}
	if a.quotaAccessor.client == nil {
		return fmt.Errorf("missing quotaAccessor.client")
	}
	if a.quotaAccessor.lister == nil {
		return fmt.Errorf("missing quotaAccessor.lister")
	}
	if a.quotaConfiguration == nil {
		return fmt.Errorf("missing quotaConfiguration")
	}
	if a.evaluator == nil {
		return fmt.Errorf("missing evaluator")
	}
	return nil
}
func (a *QuotaAdmission) Validate(attr admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attr.GetSubresource() != "" {
		return nil
	}
	if attr.GetNamespace() == "" {
		return nil
	}
	return a.evaluator.Evaluate(attr)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

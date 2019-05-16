package clusterresourcequota

import (
	"errors"
	quotatypedclient "github.com/openshift/client-go/quota/clientset/versioned/typed/quota/v1"
	quotainformer "github.com/openshift/client-go/quota/informers/externalversions/quota/v1"
	quotalister "github.com/openshift/client-go/quota/listers/quota/v1"
	oadmission "github.com/openshift/origin/pkg/cmd/server/admission"
	"github.com/openshift/origin/pkg/quota/controller/clusterquotamapping"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/apiserver/pkg/admission/plugin/namespace/lifecycle"
	"k8s.io/client-go/informers"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/rest"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	"k8s.io/kubernetes/pkg/quota/v1/install"
	"k8s.io/kubernetes/plugin/pkg/admission/resourcequota"
	resourcequotaapi "k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota"
	"sort"
	"sync"
	"time"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register("quota.openshift.io/ClusterResourceQuota", func(config io.Reader) (admission.Interface, error) {
		return NewClusterResourceQuota()
	})
}

type clusterQuotaAdmission struct {
	*admission.Handler
	clusterQuotaLister quotalister.ClusterResourceQuotaLister
	namespaceLister    corev1listers.NamespaceLister
	clusterQuotaSynced func() bool
	namespaceSynced    func() bool
	clusterQuotaClient quotatypedclient.ClusterResourceQuotasGetter
	clusterQuotaMapper clusterquotamapping.ClusterQuotaMapper
	lockFactory        LockFactory
	registry           quota.Registry
	init               sync.Once
	evaluator          resourcequota.Evaluator
}

var _ initializer.WantsExternalKubeInformerFactory = &clusterQuotaAdmission{}
var _ oadmission.WantsRESTClientConfig = &clusterQuotaAdmission{}
var _ oadmission.WantsClusterQuota = &clusterQuotaAdmission{}
var _ admission.ValidationInterface = &clusterQuotaAdmission{}

const (
	timeToWaitForCacheSync = 10 * time.Second
	numEvaluatorThreads    = 10
)

func NewClusterResourceQuota() (admission.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &clusterQuotaAdmission{Handler: admission.NewHandler(admission.Create, admission.Update), lockFactory: NewDefaultLockFactory()}, nil
}
func (q *clusterQuotaAdmission) Validate(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(a.GetSubresource()) != 0 {
		return nil
	}
	if len(a.GetNamespace()) == 0 {
		return nil
	}
	if !q.waitForSyncedStore(time.After(timeToWaitForCacheSync)) {
		return admission.NewForbidden(a, errors.New("caches not synchronized"))
	}
	q.init.Do(func() {
		clusterQuotaAccessor := newQuotaAccessor(q.clusterQuotaLister, q.namespaceLister, q.clusterQuotaClient, q.clusterQuotaMapper)
		q.evaluator = resourcequota.NewQuotaEvaluator(clusterQuotaAccessor, ignoredResources, q.registry, q.lockAquisition, &resourcequotaapi.Configuration{}, numEvaluatorThreads, utilwait.NeverStop)
	})
	return q.evaluator.Evaluate(a)
}
func (q *clusterQuotaAdmission) lockAquisition(quotas []corev1.ResourceQuota) func() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	locks := []sync.Locker{}
	sort.Sort(ByName(quotas))
	for _, quota := range quotas {
		lock := q.lockFactory.GetLock(quota.Name)
		lock.Lock()
		locks = append(locks, lock)
	}
	return func() {
		for i := len(locks) - 1; i >= 0; i-- {
			locks[i].Unlock()
		}
	}
}
func (q *clusterQuotaAdmission) waitForSyncedStore(timeout <-chan time.Time) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for !q.clusterQuotaSynced() || !q.namespaceSynced() {
		select {
		case <-time.After(100 * time.Millisecond):
		case <-timeout:
			return q.clusterQuotaSynced() && q.namespaceSynced()
		}
	}
	return true
}
func (q *clusterQuotaAdmission) SetOriginQuotaRegistry(registry quota.Registry) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	q.registry = registry
}
func (q *clusterQuotaAdmission) SetExternalKubeInformerFactory(informers informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	q.namespaceLister = informers.Core().V1().Namespaces().Lister()
	q.namespaceSynced = informers.Core().V1().Namespaces().Informer().HasSynced
}
func (q *clusterQuotaAdmission) SetRESTClientConfig(restClientConfig rest.Config) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	jsonClientConfig := rest.CopyConfig(&restClientConfig)
	jsonClientConfig.ContentConfig.AcceptContentTypes = "application/json"
	jsonClientConfig.ContentConfig.ContentType = "application/json"
	q.clusterQuotaClient, err = quotatypedclient.NewForConfig(jsonClientConfig)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
}
func (q *clusterQuotaAdmission) SetClusterQuota(clusterQuotaMapper clusterquotamapping.ClusterQuotaMapper, informers quotainformer.ClusterResourceQuotaInformer) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	q.clusterQuotaMapper = clusterQuotaMapper
	q.clusterQuotaLister = informers.Lister()
	q.clusterQuotaSynced = informers.Informer().HasSynced
}
func (q *clusterQuotaAdmission) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if q.clusterQuotaLister == nil {
		return errors.New("missing clusterQuotaLister")
	}
	if q.namespaceLister == nil {
		return errors.New("missing namespaceLister")
	}
	if q.clusterQuotaClient == nil {
		return errors.New("missing clusterQuotaClient")
	}
	if q.clusterQuotaMapper == nil {
		return errors.New("missing clusterQuotaMapper")
	}
	if q.registry == nil {
		return errors.New("missing registry")
	}
	return nil
}

type ByName []corev1.ResourceQuota

func (v ByName) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(v)
}
func (v ByName) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	v[i], v[j] = v[j], v[i]
}
func (v ByName) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v[i].Name < v[j].Name
}

var ignoredResources = map[schema.GroupResource]struct{}{}

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for k := range install.DefaultIgnoredResources() {
		ignoredResources[k] = struct{}{}
	}
	for k := range lifecycle.AccessReviewResources() {
		ignoredResources[k] = struct{}{}
	}
}

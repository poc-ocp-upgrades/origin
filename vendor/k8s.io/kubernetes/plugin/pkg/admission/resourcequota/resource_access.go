package resourcequota

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/storage/etcd"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"time"
)

type QuotaAccessor interface {
	UpdateQuotaStatus(newQuota *corev1.ResourceQuota) error
	GetQuotas(namespace string) ([]corev1.ResourceQuota, error)
}
type quotaAccessor struct {
	client          kubernetes.Interface
	lister          corev1listers.ResourceQuotaLister
	liveLookupCache *lru.Cache
	liveTTL         time.Duration
	updatedQuotas   *lru.Cache
}

func newQuotaAccessor() (*quotaAccessor, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	liveLookupCache, err := lru.New(100)
	if err != nil {
		return nil, err
	}
	updatedCache, err := lru.New(100)
	if err != nil {
		return nil, err
	}
	return &quotaAccessor{liveLookupCache: liveLookupCache, liveTTL: time.Duration(30 * time.Second), updatedQuotas: updatedCache}, nil
}
func (e *quotaAccessor) UpdateQuotaStatus(newQuota *corev1.ResourceQuota) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	updatedQuota, err := e.client.CoreV1().ResourceQuotas(newQuota.Namespace).UpdateStatus(newQuota)
	if err != nil {
		return err
	}
	key := newQuota.Namespace + "/" + newQuota.Name
	e.updatedQuotas.Add(key, updatedQuota)
	return nil
}

var etcdVersioner = etcd.APIObjectVersioner{}

func (e *quotaAccessor) checkCache(quota *corev1.ResourceQuota) *corev1.ResourceQuota {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key := quota.Namespace + "/" + quota.Name
	uncastCachedQuota, ok := e.updatedQuotas.Get(key)
	if !ok {
		return quota
	}
	cachedQuota := uncastCachedQuota.(*corev1.ResourceQuota)
	if etcdVersioner.CompareResourceVersion(quota, cachedQuota) >= 0 {
		e.updatedQuotas.Remove(key)
		return quota
	}
	return cachedQuota
}
func (e *quotaAccessor) GetQuotas(namespace string) ([]corev1.ResourceQuota, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	items, err := e.lister.ResourceQuotas(namespace).List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("error resolving quota: %v", err)
	}
	if len(items) == 0 {
		lruItemObj, ok := e.liveLookupCache.Get(namespace)
		if !ok || lruItemObj.(liveLookupEntry).expiry.Before(time.Now()) {
			liveList, err := e.client.Core().ResourceQuotas(namespace).List(metav1.ListOptions{})
			if err != nil {
				return nil, err
			}
			newEntry := liveLookupEntry{expiry: time.Now().Add(e.liveTTL)}
			for i := range liveList.Items {
				newEntry.items = append(newEntry.items, &liveList.Items[i])
			}
			e.liveLookupCache.Add(namespace, newEntry)
			lruItemObj = newEntry
		}
		lruEntry := lruItemObj.(liveLookupEntry)
		for i := range lruEntry.items {
			items = append(items, lruEntry.items[i])
		}
	}
	resourceQuotas := []corev1.ResourceQuota{}
	for i := range items {
		quota := items[i]
		quota = e.checkCache(quota)
		resourceQuotas = append(resourceQuotas, *quota)
	}
	return resourceQuotas, nil
}

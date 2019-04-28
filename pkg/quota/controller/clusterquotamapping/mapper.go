package clusterquotamapping

import (
	"reflect"
	"sync"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	quotav1 "github.com/openshift/api/quota/v1"
)

type ClusterQuotaMapper interface {
	GetClusterQuotasFor(namespaceName string) ([]string, SelectionFields)
	GetNamespacesFor(quotaName string) ([]string, quotav1.ClusterResourceQuotaSelector)
	AddListener(listener MappingChangeListener)
}
type MappingChangeListener interface {
	AddMapping(quotaName, namespaceName string)
	RemoveMapping(quotaName, namespaceName string)
}
type SelectionFields struct {
	Labels		map[string]string
	Annotations	map[string]string
}
type clusterQuotaMapper struct {
	lock				sync.RWMutex
	requiredQuotaToSelector		map[string]quotav1.ClusterResourceQuotaSelector
	requiredNamespaceToLabels	map[string]SelectionFields
	completedQuotaToSelector	map[string]quotav1.ClusterResourceQuotaSelector
	completedNamespaceToLabels	map[string]SelectionFields
	quotaToNamespaces		map[string]sets.String
	namespaceToQuota		map[string]sets.String
	listeners			[]MappingChangeListener
}

func NewClusterQuotaMapper() *clusterQuotaMapper {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterQuotaMapper{requiredQuotaToSelector: map[string]quotav1.ClusterResourceQuotaSelector{}, requiredNamespaceToLabels: map[string]SelectionFields{}, completedQuotaToSelector: map[string]quotav1.ClusterResourceQuotaSelector{}, completedNamespaceToLabels: map[string]SelectionFields{}, quotaToNamespaces: map[string]sets.String{}, namespaceToQuota: map[string]sets.String{}}
}
func (m *clusterQuotaMapper) GetClusterQuotasFor(namespaceName string) ([]string, SelectionFields) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.RLock()
	defer m.lock.RUnlock()
	quotas, ok := m.namespaceToQuota[namespaceName]
	if !ok {
		return []string{}, m.completedNamespaceToLabels[namespaceName]
	}
	return quotas.List(), m.completedNamespaceToLabels[namespaceName]
}
func (m *clusterQuotaMapper) GetNamespacesFor(quotaName string) ([]string, quotav1.ClusterResourceQuotaSelector) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.RLock()
	defer m.lock.RUnlock()
	namespaces, ok := m.quotaToNamespaces[quotaName]
	if !ok {
		return []string{}, m.completedQuotaToSelector[quotaName]
	}
	return namespaces.List(), m.completedQuotaToSelector[quotaName]
}
func (m *clusterQuotaMapper) AddListener(listener MappingChangeListener) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.Lock()
	defer m.lock.Unlock()
	m.listeners = append(m.listeners, listener)
}
func (m *clusterQuotaMapper) requireQuota(quota *quotav1.ClusterResourceQuota) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.RLock()
	selector, exists := m.requiredQuotaToSelector[quota.Name]
	m.lock.RUnlock()
	if selectorMatches(selector, exists, quota) {
		return false
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	selector, exists = m.requiredQuotaToSelector[quota.Name]
	if selectorMatches(selector, exists, quota) {
		return false
	}
	m.requiredQuotaToSelector[quota.Name] = quota.Spec.Selector
	return true
}
func (m *clusterQuotaMapper) completeQuota(quota *quotav1.ClusterResourceQuota) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.Lock()
	defer m.lock.Unlock()
	m.completedQuotaToSelector[quota.Name] = quota.Spec.Selector
}
func (m *clusterQuotaMapper) removeQuota(quotaName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.requiredQuotaToSelector, quotaName)
	delete(m.completedQuotaToSelector, quotaName)
	delete(m.quotaToNamespaces, quotaName)
	for namespaceName, quotas := range m.namespaceToQuota {
		if quotas.Has(quotaName) {
			quotas.Delete(quotaName)
			for _, listener := range m.listeners {
				listener.RemoveMapping(quotaName, namespaceName)
			}
		}
	}
}
func (m *clusterQuotaMapper) requireNamespace(namespace metav1.Object) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.RLock()
	selectionFields, exists := m.requiredNamespaceToLabels[namespace.GetName()]
	m.lock.RUnlock()
	if selectionFieldsMatch(selectionFields, exists, namespace) {
		return false
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	selectionFields, exists = m.requiredNamespaceToLabels[namespace.GetName()]
	if selectionFieldsMatch(selectionFields, exists, namespace) {
		return false
	}
	m.requiredNamespaceToLabels[namespace.GetName()] = GetSelectionFields(namespace)
	return true
}
func (m *clusterQuotaMapper) completeNamespace(namespace metav1.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.Lock()
	defer m.lock.Unlock()
	m.completedNamespaceToLabels[namespace.GetName()] = GetSelectionFields(namespace)
}
func (m *clusterQuotaMapper) removeNamespace(namespaceName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.requiredNamespaceToLabels, namespaceName)
	delete(m.completedNamespaceToLabels, namespaceName)
	delete(m.namespaceToQuota, namespaceName)
	for quotaName, namespaces := range m.quotaToNamespaces {
		if namespaces.Has(namespaceName) {
			namespaces.Delete(namespaceName)
			for _, listener := range m.listeners {
				listener.RemoveMapping(quotaName, namespaceName)
			}
		}
	}
}
func selectorMatches(selector quotav1.ClusterResourceQuotaSelector, exists bool, quota *quotav1.ClusterResourceQuota) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return exists && equality.Semantic.DeepEqual(selector, quota.Spec.Selector)
}
func selectionFieldsMatch(selectionFields SelectionFields, exists bool, namespace metav1.Object) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return exists && reflect.DeepEqual(selectionFields, GetSelectionFields(namespace))
}
func (m *clusterQuotaMapper) setMapping(quota *quotav1.ClusterResourceQuota, namespace metav1.Object, remove bool) (bool, bool, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.RLock()
	selector, selectorExists := m.requiredQuotaToSelector[quota.Name]
	selectionFields, selectionFieldsExist := m.requiredNamespaceToLabels[namespace.GetName()]
	m.lock.RUnlock()
	if !selectorMatches(selector, selectorExists, quota) {
		return false, false, selectionFieldsMatch(selectionFields, selectionFieldsExist, namespace)
	}
	if !selectionFieldsMatch(selectionFields, selectionFieldsExist, namespace) {
		return false, true, false
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	selector, selectorExists = m.requiredQuotaToSelector[quota.Name]
	selectionFields, selectionFieldsExist = m.requiredNamespaceToLabels[namespace.GetName()]
	if !selectorMatches(selector, selectorExists, quota) {
		return false, false, selectionFieldsMatch(selectionFields, selectionFieldsExist, namespace)
	}
	if !selectionFieldsMatch(selectionFields, selectionFieldsExist, namespace) {
		return false, true, false
	}
	if remove {
		mutated := false
		namespaces, ok := m.quotaToNamespaces[quota.Name]
		if !ok {
			m.quotaToNamespaces[quota.Name] = sets.String{}
		} else {
			mutated = namespaces.Has(namespace.GetName())
			namespaces.Delete(namespace.GetName())
		}
		quotas, ok := m.namespaceToQuota[namespace.GetName()]
		if !ok {
			m.namespaceToQuota[namespace.GetName()] = sets.String{}
		} else {
			mutated = mutated || quotas.Has(quota.Name)
			quotas.Delete(quota.Name)
		}
		if mutated {
			for _, listener := range m.listeners {
				listener.RemoveMapping(quota.Name, namespace.GetName())
			}
		}
		return true, true, true
	}
	mutated := false
	namespaces, ok := m.quotaToNamespaces[quota.Name]
	if !ok {
		mutated = true
		m.quotaToNamespaces[quota.Name] = sets.NewString(namespace.GetName())
	} else {
		mutated = !namespaces.Has(namespace.GetName())
		namespaces.Insert(namespace.GetName())
	}
	quotas, ok := m.namespaceToQuota[namespace.GetName()]
	if !ok {
		mutated = true
		m.namespaceToQuota[namespace.GetName()] = sets.NewString(quota.Name)
	} else {
		mutated = mutated || !quotas.Has(quota.Name)
		quotas.Insert(quota.Name)
	}
	if mutated {
		for _, listener := range m.listeners {
			listener.AddMapping(quota.Name, namespace.GetName())
		}
	}
	return true, true, true
}
func GetSelectionFields(namespace metav1.Object) SelectionFields {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return SelectionFields{Labels: namespace.GetLabels(), Annotations: namespace.GetAnnotations()}
}

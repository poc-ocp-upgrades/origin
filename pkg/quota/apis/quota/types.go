package quota

import (
	"container/list"
	"reflect"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/core"
)

type ClusterResourceQuota struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ClusterResourceQuotaSpec
	Status	ClusterResourceQuotaStatus
}
type ClusterResourceQuotaSpec struct {
	Selector	ClusterResourceQuotaSelector
	Quota		core.ResourceQuotaSpec
}
type ClusterResourceQuotaSelector struct {
	LabelSelector		*metav1.LabelSelector
	AnnotationSelector	map[string]string
}
type ClusterResourceQuotaStatus struct {
	Total		core.ResourceQuotaStatus
	Namespaces	ResourceQuotasStatusByNamespace
}
type ClusterResourceQuotaList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ClusterResourceQuota
}
type AppliedClusterResourceQuota struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ClusterResourceQuotaSpec
	Status	ClusterResourceQuotaStatus
}
type AppliedClusterResourceQuotaList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]AppliedClusterResourceQuota
}
type ResourceQuotasStatusByNamespace struct{ orderedMap orderedMap }

func (o *ResourceQuotasStatusByNamespace) Insert(key string, value core.ResourceQuotaStatus) {
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
	o.orderedMap.Insert(key, value)
}
func (o *ResourceQuotasStatusByNamespace) Get(key string) (core.ResourceQuotaStatus, bool) {
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
	ret, ok := o.orderedMap.Get(key)
	if !ok {
		return core.ResourceQuotaStatus{}, ok
	}
	return ret.(core.ResourceQuotaStatus), ok
}
func (o *ResourceQuotasStatusByNamespace) Remove(key string) {
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
	o.orderedMap.Remove(key)
}
func (o *ResourceQuotasStatusByNamespace) OrderedKeys() *list.List {
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
	return o.orderedMap.OrderedKeys()
}
func (o ResourceQuotasStatusByNamespace) DeepCopy() ResourceQuotasStatusByNamespace {
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
	out := ResourceQuotasStatusByNamespace{}
	for e := o.OrderedKeys().Front(); e != nil; e = e.Next() {
		namespace := e.Value.(string)
		instatus, _ := o.Get(namespace)
		outstatus := instatus.DeepCopy()
		out.Insert(namespace, *outstatus)
	}
	return out
}
func init() {
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
	if err := equality.Semantic.AddFuncs(func(o1, o2 ResourceQuotasStatusByNamespace) bool {
		return reflect.DeepEqual(o1.orderedMap, o2.orderedMap)
	}, func(o1, o2 *ResourceQuotasStatusByNamespace) bool {
		if o1 == nil && o2 == nil {
			return true
		}
		if (o1 == nil) != (o2 == nil) {
			return false
		}
		return reflect.DeepEqual(o1.orderedMap, o2.orderedMap)
	}); err != nil {
		panic(err)
	}
}

type orderedMap struct {
	backingMap	map[string]interface{}
	orderedKeys	*list.List
}

func (o *orderedMap) Insert(key string, value interface{}) {
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
	if o.backingMap == nil {
		o.backingMap = map[string]interface{}{}
	}
	if o.orderedKeys == nil {
		o.orderedKeys = list.New()
	}
	if _, exists := o.backingMap[key]; !exists {
		o.orderedKeys.PushBack(key)
	}
	o.backingMap[key] = value
}
func (o *orderedMap) Get(key string) (interface{}, bool) {
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
	ret, ok := o.backingMap[key]
	return ret, ok
}
func (o *orderedMap) Remove(key string) {
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
	delete(o.backingMap, key)
	if o.orderedKeys == nil {
		return
	}
	for e := o.orderedKeys.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			o.orderedKeys.Remove(e)
			break
		}
	}
}
func (o *orderedMap) OrderedKeys() *list.List {
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
	if o.orderedKeys == nil {
		return list.New()
	}
	return o.orderedKeys
}

package gce

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	"reflect"
	"sync"
	"time"
)

const (
	UIDConfigMapName    = "ingress-uid"
	UIDNamespace        = metav1.NamespaceSystem
	UIDCluster          = "uid"
	UIDProvider         = "provider-uid"
	UIDLengthBytes      = 8
	updateFuncFrequency = 10 * time.Minute
)

type ClusterID struct {
	idLock     sync.RWMutex
	client     clientset.Interface
	cfgMapKey  string
	store      cache.Store
	providerID *string
	clusterID  *string
}

func (g *Cloud) watchClusterID(stop <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	g.ClusterID = ClusterID{cfgMapKey: fmt.Sprintf("%v/%v", UIDNamespace, UIDConfigMapName), client: g.client}
	mapEventHandler := cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		m, ok := obj.(*v1.ConfigMap)
		if !ok || m == nil {
			klog.Errorf("Expected v1.ConfigMap, item=%+v, typeIsOk=%v", obj, ok)
			return
		}
		if m.Namespace != UIDNamespace || m.Name != UIDConfigMapName {
			return
		}
		klog.V(4).Infof("Observed new configmap for clusteriD: %v, %v; setting local values", m.Name, m.Data)
		g.ClusterID.update(m)
	}, UpdateFunc: func(old, cur interface{}) {
		m, ok := cur.(*v1.ConfigMap)
		if !ok || m == nil {
			klog.Errorf("Expected v1.ConfigMap, item=%+v, typeIsOk=%v", cur, ok)
			return
		}
		if m.Namespace != UIDNamespace || m.Name != UIDConfigMapName {
			return
		}
		if reflect.DeepEqual(old, cur) {
			return
		}
		klog.V(4).Infof("Observed updated configmap for clusteriD %v, %v; setting local values", m.Name, m.Data)
		g.ClusterID.update(m)
	}}
	listerWatcher := cache.NewListWatchFromClient(g.ClusterID.client.CoreV1().RESTClient(), "configmaps", UIDNamespace, fields.Everything())
	var controller cache.Controller
	g.ClusterID.store, controller = cache.NewInformer(newSingleObjectListerWatcher(listerWatcher, UIDConfigMapName), &v1.ConfigMap{}, updateFuncFrequency, mapEventHandler)
	controller.Run(stop)
}
func (ci *ClusterID) GetID() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := ci.getOrInitialize(); err != nil {
		return "", err
	}
	ci.idLock.RLock()
	defer ci.idLock.RUnlock()
	if ci.clusterID == nil {
		return "", errors.New("Could not retrieve cluster id")
	}
	if ci.providerID != nil {
		return *ci.providerID, nil
	}
	return *ci.clusterID, nil
}
func (ci *ClusterID) GetFederationID() (string, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := ci.getOrInitialize(); err != nil {
		return "", false, err
	}
	ci.idLock.RLock()
	defer ci.idLock.RUnlock()
	if ci.clusterID == nil {
		return "", false, errors.New("could not retrieve cluster id")
	}
	if ci.providerID == nil || *ci.clusterID == *ci.providerID {
		return "", false, nil
	}
	return *ci.clusterID, true, nil
}
func (ci *ClusterID) getOrInitialize() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ci.store == nil {
		return errors.New("Cloud.ClusterID is not ready. Call Initialize() before using")
	}
	if ci.clusterID != nil {
		return nil
	}
	exists, err := ci.getConfigMap()
	if err != nil {
		return err
	} else if exists {
		return nil
	}
	newID, err := makeUID()
	if err != nil {
		return err
	}
	klog.V(4).Infof("Creating clusteriD: %v", newID)
	cfg := &v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: UIDConfigMapName, Namespace: UIDNamespace}}
	cfg.Data = map[string]string{UIDCluster: newID, UIDProvider: newID}
	if _, err := ci.client.CoreV1().ConfigMaps(UIDNamespace).Create(cfg); err != nil {
		klog.Errorf("GCE cloud provider failed to create %v config map to store cluster id: %v", ci.cfgMapKey, err)
		return err
	}
	klog.V(2).Infof("Created a config map containing clusteriD: %v", newID)
	ci.update(cfg)
	return nil
}
func (ci *ClusterID) getConfigMap() (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	item, exists, err := ci.store.GetByKey(ci.cfgMapKey)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}
	m, ok := item.(*v1.ConfigMap)
	if !ok || m == nil {
		err = fmt.Errorf("Expected v1.ConfigMap, item=%+v, typeIsOk=%v", item, ok)
		klog.Error(err)
		return false, err
	}
	ci.update(m)
	return true, nil
}
func (ci *ClusterID) update(m *v1.ConfigMap) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ci.idLock.Lock()
	defer ci.idLock.Unlock()
	if clusterID, exists := m.Data[UIDCluster]; exists {
		ci.clusterID = &clusterID
	}
	if provID, exists := m.Data[UIDProvider]; exists {
		ci.providerID = &provID
	}
}
func makeUID() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b := make([]byte, UIDLengthBytes)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
func newSingleObjectListerWatcher(lw cache.ListerWatcher, objectName string) *singleObjListerWatcher {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &singleObjListerWatcher{lw: lw, objectName: objectName}
}

type singleObjListerWatcher struct {
	lw         cache.ListerWatcher
	objectName string
}

func (sow *singleObjListerWatcher) List(options metav1.ListOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	options.FieldSelector = "metadata.name=" + sow.objectName
	return sow.lw.List(options)
}
func (sow *singleObjListerWatcher) Watch(options metav1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	options.FieldSelector = "metadata.name=" + sow.objectName
	return sow.lw.Watch(options)
}

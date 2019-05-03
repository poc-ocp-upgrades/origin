package persistentvolume

import (
 "fmt"
 "strconv"
 "sync"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/api/meta"
 "k8s.io/client-go/tools/cache"
)

type AssumeCache interface {
 Assume(obj interface{}) error
 Restore(objName string)
 Get(objName string) (interface{}, error)
 List(indexObj interface{}) []interface{}
}
type errWrongType struct {
 typeName string
 object   interface{}
}

func (e *errWrongType) Error() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("could not convert object to type %v: %+v", e.typeName, e.object)
}

type errNotFound struct {
 typeName   string
 objectName string
}

func (e *errNotFound) Error() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("could not find %v %q", e.typeName, e.objectName)
}

type errObjectName struct{ detailedErr error }

func (e *errObjectName) Error() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("failed to get object name: %v", e.detailedErr)
}

type assumeCache struct {
 rwMutex     sync.RWMutex
 description string
 store       cache.Indexer
 indexFunc   cache.IndexFunc
 indexName   string
}
type objInfo struct {
 name      string
 latestObj interface{}
 apiObj    interface{}
}

func objInfoKeyFunc(obj interface{}) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objInfo, ok := obj.(*objInfo)
 if !ok {
  return "", &errWrongType{"objInfo", obj}
 }
 return objInfo.name, nil
}
func (c *assumeCache) objInfoIndexFunc(obj interface{}) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objInfo, ok := obj.(*objInfo)
 if !ok {
  return []string{""}, &errWrongType{"objInfo", obj}
 }
 return c.indexFunc(objInfo.latestObj)
}
func NewAssumeCache(informer cache.SharedIndexInformer, description, indexName string, indexFunc cache.IndexFunc) *assumeCache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c := &assumeCache{description: description, indexFunc: indexFunc, indexName: indexName}
 c.store = cache.NewIndexer(objInfoKeyFunc, cache.Indexers{indexName: c.objInfoIndexFunc})
 if informer != nil {
  informer.AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.add, UpdateFunc: c.update, DeleteFunc: c.delete})
 }
 return c
}
func (c *assumeCache) add(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj == nil {
  return
 }
 name, err := cache.MetaNamespaceKeyFunc(obj)
 if err != nil {
  klog.Errorf("add failed: %v", &errObjectName{err})
  return
 }
 c.rwMutex.Lock()
 defer c.rwMutex.Unlock()
 if objInfo, _ := c.getObjInfo(name); objInfo != nil {
  newVersion, err := c.getObjVersion(name, obj)
  if err != nil {
   klog.Errorf("add: couldn't get object version: %v", err)
   return
  }
  storedVersion, err := c.getObjVersion(name, objInfo.latestObj)
  if err != nil {
   klog.Errorf("add: couldn't get stored object version: %v", err)
   return
  }
  if newVersion <= storedVersion {
   klog.V(10).Infof("Skip adding %v %v to assume cache because version %v is not newer than %v", c.description, name, newVersion, storedVersion)
   return
  }
 }
 objInfo := &objInfo{name: name, latestObj: obj, apiObj: obj}
 c.store.Update(objInfo)
 klog.V(10).Infof("Adding %v %v to assume cache: %+v ", c.description, name, obj)
}
func (c *assumeCache) update(oldObj interface{}, newObj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.add(newObj)
}
func (c *assumeCache) delete(obj interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj == nil {
  return
 }
 name, err := cache.MetaNamespaceKeyFunc(obj)
 if err != nil {
  klog.Errorf("delete failed: %v", &errObjectName{err})
  return
 }
 c.rwMutex.Lock()
 defer c.rwMutex.Unlock()
 objInfo := &objInfo{name: name}
 err = c.store.Delete(objInfo)
 if err != nil {
  klog.Errorf("delete: failed to delete %v %v: %v", c.description, name, err)
 }
}
func (c *assumeCache) getObjVersion(name string, obj interface{}) (int64, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objAccessor, err := meta.Accessor(obj)
 if err != nil {
  return -1, err
 }
 objResourceVersion, err := strconv.ParseInt(objAccessor.GetResourceVersion(), 10, 64)
 if err != nil {
  return -1, fmt.Errorf("error parsing ResourceVersion %q for %v %q: %s", objAccessor.GetResourceVersion(), c.description, name, err)
 }
 return objResourceVersion, nil
}
func (c *assumeCache) getObjInfo(name string) (*objInfo, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, ok, err := c.store.GetByKey(name)
 if err != nil {
  return nil, err
 }
 if !ok {
  return nil, &errNotFound{c.description, name}
 }
 objInfo, ok := obj.(*objInfo)
 if !ok {
  return nil, &errWrongType{"objInfo", obj}
 }
 return objInfo, nil
}
func (c *assumeCache) Get(objName string) (interface{}, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.rwMutex.RLock()
 defer c.rwMutex.RUnlock()
 objInfo, err := c.getObjInfo(objName)
 if err != nil {
  return nil, err
 }
 return objInfo.latestObj, nil
}
func (c *assumeCache) List(indexObj interface{}) []interface{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.rwMutex.RLock()
 defer c.rwMutex.RUnlock()
 allObjs := []interface{}{}
 objs, err := c.store.Index(c.indexName, &objInfo{latestObj: indexObj})
 if err != nil {
  klog.Errorf("list index error: %v", err)
  return nil
 }
 for _, obj := range objs {
  objInfo, ok := obj.(*objInfo)
  if !ok {
   klog.Errorf("list error: %v", &errWrongType{"objInfo", obj})
   continue
  }
  allObjs = append(allObjs, objInfo.latestObj)
 }
 return allObjs
}
func (c *assumeCache) Assume(obj interface{}) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 name, err := cache.MetaNamespaceKeyFunc(obj)
 if err != nil {
  return &errObjectName{err}
 }
 c.rwMutex.Lock()
 defer c.rwMutex.Unlock()
 objInfo, err := c.getObjInfo(name)
 if err != nil {
  return err
 }
 newVersion, err := c.getObjVersion(name, obj)
 if err != nil {
  return err
 }
 storedVersion, err := c.getObjVersion(name, objInfo.latestObj)
 if err != nil {
  return err
 }
 if newVersion < storedVersion {
  return fmt.Errorf("%v %q is out of sync", c.description, name)
 }
 objInfo.latestObj = obj
 klog.V(4).Infof("Assumed %v %q, version %v", c.description, name, newVersion)
 return nil
}
func (c *assumeCache) Restore(objName string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.rwMutex.Lock()
 defer c.rwMutex.Unlock()
 objInfo, err := c.getObjInfo(objName)
 if err != nil {
  klog.V(5).Infof("Restore %v %q warning: %v", c.description, objName, err)
 } else {
  objInfo.latestObj = objInfo.apiObj
  klog.V(4).Infof("Restored %v %q", c.description, objName)
 }
}

type PVAssumeCache interface {
 AssumeCache
 GetPV(pvName string) (*v1.PersistentVolume, error)
 ListPVs(storageClassName string) []*v1.PersistentVolume
}
type pvAssumeCache struct{ *assumeCache }

func pvStorageClassIndexFunc(obj interface{}) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if pv, ok := obj.(*v1.PersistentVolume); ok {
  return []string{pv.Spec.StorageClassName}, nil
 }
 return []string{""}, fmt.Errorf("object is not a v1.PersistentVolume: %v", obj)
}
func NewPVAssumeCache(informer cache.SharedIndexInformer) PVAssumeCache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &pvAssumeCache{assumeCache: NewAssumeCache(informer, "v1.PersistentVolume", "storageclass", pvStorageClassIndexFunc)}
}
func (c *pvAssumeCache) GetPV(pvName string) (*v1.PersistentVolume, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Get(pvName)
 if err != nil {
  return nil, err
 }
 pv, ok := obj.(*v1.PersistentVolume)
 if !ok {
  return nil, &errWrongType{"v1.PersistentVolume", obj}
 }
 return pv, nil
}
func (c *pvAssumeCache) ListPVs(storageClassName string) []*v1.PersistentVolume {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objs := c.List(&v1.PersistentVolume{Spec: v1.PersistentVolumeSpec{StorageClassName: storageClassName}})
 pvs := []*v1.PersistentVolume{}
 for _, obj := range objs {
  pv, ok := obj.(*v1.PersistentVolume)
  if !ok {
   klog.Errorf("ListPVs: %v", &errWrongType{"v1.PersistentVolume", obj})
  }
  pvs = append(pvs, pv)
 }
 return pvs
}

type PVCAssumeCache interface {
 AssumeCache
 GetPVC(pvcKey string) (*v1.PersistentVolumeClaim, error)
}
type pvcAssumeCache struct{ *assumeCache }

func NewPVCAssumeCache(informer cache.SharedIndexInformer) PVCAssumeCache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &pvcAssumeCache{assumeCache: NewAssumeCache(informer, "v1.PersistentVolumeClaim", "namespace", cache.MetaNamespaceIndexFunc)}
}
func (c *pvcAssumeCache) GetPVC(pvcKey string) (*v1.PersistentVolumeClaim, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Get(pvcKey)
 if err != nil {
  return nil, err
 }
 pvc, ok := obj.(*v1.PersistentVolumeClaim)
 if !ok {
  return nil, &errWrongType{"v1.PersistentVolumeClaim", obj}
 }
 return pvc, nil
}

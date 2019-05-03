package garbagecollector

import (
 "encoding/json"
 "fmt"
 "strings"
 "k8s.io/apimachinery/pkg/api/errors"
 "k8s.io/apimachinery/pkg/api/meta"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/kubernetes/pkg/controller/garbagecollector/metaonly"
)

func deleteOwnerRefStrategicMergePatch(dependentUID types.UID, ownerUIDs ...types.UID) []byte {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var pieces []string
 for _, ownerUID := range ownerUIDs {
  pieces = append(pieces, fmt.Sprintf(`{"$patch":"delete","uid":"%s"}`, ownerUID))
 }
 patch := fmt.Sprintf(`{"metadata":{"ownerReferences":[%s],"uid":"%s"}}`, strings.Join(pieces, ","), dependentUID)
 return []byte(patch)
}
func (gc *GarbageCollector) getMetadata(apiVersion, kind, namespace, name string) (metav1.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiResource, _, err := gc.apiResource(apiVersion, kind)
 if err != nil {
  return nil, err
 }
 gc.dependencyGraphBuilder.monitorLock.RLock()
 defer gc.dependencyGraphBuilder.monitorLock.RUnlock()
 m, ok := gc.dependencyGraphBuilder.monitors[apiResource]
 if !ok || m == nil {
  return gc.dynamicClient.Resource(apiResource).Namespace(namespace).Get(name, metav1.GetOptions{})
 }
 key := name
 if len(namespace) != 0 {
  key = namespace + "/" + name
 }
 raw, exist, err := m.store.GetByKey(key)
 if err != nil {
  return nil, err
 }
 if !exist {
  return gc.dynamicClient.Resource(apiResource).Namespace(namespace).Get(name, metav1.GetOptions{})
 }
 obj, ok := raw.(runtime.Object)
 if !ok {
  return nil, fmt.Errorf("expect a runtime.Object, got %v", raw)
 }
 return meta.Accessor(obj)
}

type objectForPatch struct {
 ObjectMetaForPatch `json:"metadata"`
}
type ObjectMetaForPatch struct {
 ResourceVersion string                  `json:"resourceVersion"`
 OwnerReferences []metav1.OwnerReference `json:"ownerReferences"`
}
type jsonMergePatchFunc func(*node) ([]byte, error)

func (gc *GarbageCollector) patch(item *node, smp []byte, jmp jsonMergePatchFunc) (*unstructured.Unstructured, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 smpResult, err := gc.patchObject(item.identity, smp, types.StrategicMergePatchType)
 if err == nil {
  return smpResult, nil
 }
 if !errors.IsUnsupportedMediaType(err) {
  return nil, err
 }
 patch, err := jmp(item)
 if err != nil {
  return nil, err
 }
 return gc.patchObject(item.identity, patch, types.MergePatchType)
}
func (gc *GarbageCollector) deleteOwnerRefJSONMergePatch(item *node, ownerUIDs ...types.UID) ([]byte, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 accessor, err := gc.getMetadata(item.identity.APIVersion, item.identity.Kind, item.identity.Namespace, item.identity.Name)
 if err != nil {
  return nil, err
 }
 expectedObjectMeta := ObjectMetaForPatch{}
 expectedObjectMeta.ResourceVersion = accessor.GetResourceVersion()
 refs := accessor.GetOwnerReferences()
 for _, ref := range refs {
  var skip bool
  for _, ownerUID := range ownerUIDs {
   if ref.UID == ownerUID {
    skip = true
    break
   }
  }
  if !skip {
   expectedObjectMeta.OwnerReferences = append(expectedObjectMeta.OwnerReferences, ref)
  }
 }
 return json.Marshal(objectForPatch{expectedObjectMeta})
}
func (n *node) unblockOwnerReferencesStrategicMergePatch() ([]byte, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var dummy metaonly.MetadataOnlyObject
 var blockingRefs []metav1.OwnerReference
 falseVar := false
 for _, owner := range n.owners {
  if owner.BlockOwnerDeletion != nil && *owner.BlockOwnerDeletion {
   ref := owner
   ref.BlockOwnerDeletion = &falseVar
   blockingRefs = append(blockingRefs, ref)
  }
 }
 dummy.ObjectMeta.SetOwnerReferences(blockingRefs)
 dummy.ObjectMeta.UID = n.identity.UID
 return json.Marshal(dummy)
}
func (gc *GarbageCollector) unblockOwnerReferencesJSONMergePatch(n *node) ([]byte, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 accessor, err := gc.getMetadata(n.identity.APIVersion, n.identity.Kind, n.identity.Namespace, n.identity.Name)
 if err != nil {
  return nil, err
 }
 expectedObjectMeta := ObjectMetaForPatch{}
 expectedObjectMeta.ResourceVersion = accessor.GetResourceVersion()
 var expectedOwners []metav1.OwnerReference
 falseVar := false
 for _, owner := range n.owners {
  owner.BlockOwnerDeletion = &falseVar
  expectedOwners = append(expectedOwners, owner)
 }
 expectedObjectMeta.OwnerReferences = expectedOwners
 return json.Marshal(objectForPatch{expectedObjectMeta})
}

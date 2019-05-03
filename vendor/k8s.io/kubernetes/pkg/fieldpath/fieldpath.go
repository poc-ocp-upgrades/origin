package fieldpath

import (
 "fmt"
 "strings"
 "k8s.io/apimachinery/pkg/api/meta"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apimachinery/pkg/util/validation"
)

func FormatMap(m map[string]string) (fmtStr string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 keys := sets.NewString()
 for key := range m {
  keys.Insert(key)
 }
 for _, key := range keys.List() {
  fmtStr += fmt.Sprintf("%v=%q\n", key, m[key])
 }
 fmtStr = strings.TrimSuffix(fmtStr, "\n")
 return
}
func ExtractFieldPathAsString(obj interface{}, fieldPath string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 accessor, err := meta.Accessor(obj)
 if err != nil {
  return "", nil
 }
 if path, subscript, ok := SplitMaybeSubscriptedPath(fieldPath); ok {
  switch path {
  case "metadata.annotations":
   if errs := validation.IsQualifiedName(strings.ToLower(subscript)); len(errs) != 0 {
    return "", fmt.Errorf("invalid key subscript in %s: %s", fieldPath, strings.Join(errs, ";"))
   }
   return accessor.GetAnnotations()[subscript], nil
  case "metadata.labels":
   if errs := validation.IsQualifiedName(subscript); len(errs) != 0 {
    return "", fmt.Errorf("invalid key subscript in %s: %s", fieldPath, strings.Join(errs, ";"))
   }
   return accessor.GetLabels()[subscript], nil
  default:
   return "", fmt.Errorf("fieldPath %q does not support subscript", fieldPath)
  }
 }
 switch fieldPath {
 case "metadata.annotations":
  return FormatMap(accessor.GetAnnotations()), nil
 case "metadata.labels":
  return FormatMap(accessor.GetLabels()), nil
 case "metadata.name":
  return accessor.GetName(), nil
 case "metadata.namespace":
  return accessor.GetNamespace(), nil
 case "metadata.uid":
  return string(accessor.GetUID()), nil
 }
 return "", fmt.Errorf("unsupported fieldPath: %v", fieldPath)
}
func SplitMaybeSubscriptedPath(fieldPath string) (string, string, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !strings.HasSuffix(fieldPath, "']") {
  return fieldPath, "", false
 }
 s := strings.TrimSuffix(fieldPath, "']")
 parts := strings.SplitN(s, "['", 2)
 if len(parts) < 2 {
  return fieldPath, "", false
 }
 if len(parts[0]) == 0 {
  return fieldPath, "", false
 }
 return parts[0], parts[1], true
}

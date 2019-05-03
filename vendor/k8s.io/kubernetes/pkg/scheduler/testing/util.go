package testing

import (
 "fmt"
 "mime"
 "os"
 "reflect"
 "strings"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 api "k8s.io/kubernetes/pkg/apis/core"
 _ "k8s.io/kubernetes/pkg/apis/core/install"
)

type TestGroup struct {
 externalGroupVersion schema.GroupVersion
 internalGroupVersion schema.GroupVersion
 internalTypes        map[string]reflect.Type
 externalTypes        map[string]reflect.Type
}

var (
 Groups     = make(map[string]TestGroup)
 Test       TestGroup
 serializer runtime.SerializerInfo
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if apiMediaType := os.Getenv("KUBE_TEST_API_TYPE"); len(apiMediaType) > 0 {
  var ok bool
  mediaType, _, err := mime.ParseMediaType(apiMediaType)
  if err != nil {
   panic(err)
  }
  serializer, ok = runtime.SerializerInfoForMediaType(legacyscheme.Codecs.SupportedMediaTypes(), mediaType)
  if !ok {
   panic(fmt.Sprintf("no serializer for %s", apiMediaType))
  }
 }
 kubeTestAPI := os.Getenv("KUBE_TEST_API")
 if len(kubeTestAPI) != 0 {
  testGroupVersions := strings.Split(kubeTestAPI, ",")
  for i := len(testGroupVersions) - 1; i >= 0; i-- {
   gvString := testGroupVersions[i]
   groupVersion, err := schema.ParseGroupVersion(gvString)
   if err != nil {
    panic(fmt.Sprintf("Error parsing groupversion %v: %v", gvString, err))
   }
   internalGroupVersion := schema.GroupVersion{Group: groupVersion.Group, Version: runtime.APIVersionInternal}
   Groups[groupVersion.Group] = TestGroup{externalGroupVersion: groupVersion, internalGroupVersion: internalGroupVersion, internalTypes: legacyscheme.Scheme.KnownTypes(internalGroupVersion), externalTypes: legacyscheme.Scheme.KnownTypes(groupVersion)}
  }
 }
 if _, ok := Groups[api.GroupName]; !ok {
  externalGroupVersion := schema.GroupVersion{Group: api.GroupName, Version: "v1"}
  Groups[api.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion, internalGroupVersion: api.SchemeGroupVersion, internalTypes: legacyscheme.Scheme.KnownTypes(api.SchemeGroupVersion), externalTypes: legacyscheme.Scheme.KnownTypes(externalGroupVersion)}
 }
 Test = Groups[api.GroupName]
}
func (g TestGroup) Codec() runtime.Codec {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if serializer.Serializer == nil {
  return legacyscheme.Codecs.LegacyCodec(g.externalGroupVersion)
 }
 return legacyscheme.Codecs.CodecForVersions(serializer.Serializer, legacyscheme.Codecs.UniversalDeserializer(), schema.GroupVersions{g.externalGroupVersion}, nil)
}
func (g TestGroup) SelfLink(resource, name string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if g.externalGroupVersion.Group == api.GroupName {
  if name == "" {
   return fmt.Sprintf("/api/%s/%s", g.externalGroupVersion.Version, resource)
  }
  return fmt.Sprintf("/api/%s/%s/%s", g.externalGroupVersion.Version, resource, name)
 }
 if name == "" {
  return fmt.Sprintf("/apis/%s/%s/%s", g.externalGroupVersion.Group, g.externalGroupVersion.Version, resource)
 }
 return fmt.Sprintf("/apis/%s/%s/%s/%s", g.externalGroupVersion.Group, g.externalGroupVersion.Version, resource, name)
}
func (g TestGroup) ResourcePathWithPrefix(prefix, resource, namespace, name string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var path string
 if g.externalGroupVersion.Group == api.GroupName {
  path = "/api/" + g.externalGroupVersion.Version
 } else {
  path = "/apis/" + g.externalGroupVersion.Group + "/" + g.externalGroupVersion.Version
 }
 if prefix != "" {
  path = path + "/" + prefix
 }
 if namespace != "" {
  path = path + "/namespaces/" + namespace
 }
 resource = strings.ToLower(resource)
 if resource != "" {
  path = path + "/" + resource
 }
 if name != "" {
  path = path + "/" + name
 }
 return path
}
func (g TestGroup) ResourcePath(resource, namespace, name string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return g.ResourcePathWithPrefix("", resource, namespace, name)
}
func (g TestGroup) SubResourcePath(resource, namespace, name, sub string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 path := g.ResourcePathWithPrefix("", resource, namespace, name)
 if sub != "" {
  path = path + "/" + sub
 }
 return path
}

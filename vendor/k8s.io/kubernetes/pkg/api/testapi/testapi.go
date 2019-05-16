package testapi

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/recognizer"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/admission"
	_ "k8s.io/kubernetes/pkg/apis/admission/install"
	"k8s.io/kubernetes/pkg/apis/admissionregistration"
	_ "k8s.io/kubernetes/pkg/apis/admissionregistration/install"
	"k8s.io/kubernetes/pkg/apis/apps"
	_ "k8s.io/kubernetes/pkg/apis/apps/install"
	"k8s.io/kubernetes/pkg/apis/auditregistration"
	_ "k8s.io/kubernetes/pkg/apis/auditregistration/install"
	_ "k8s.io/kubernetes/pkg/apis/authentication/install"
	"k8s.io/kubernetes/pkg/apis/authorization"
	_ "k8s.io/kubernetes/pkg/apis/authorization/install"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	_ "k8s.io/kubernetes/pkg/apis/autoscaling/install"
	"k8s.io/kubernetes/pkg/apis/batch"
	_ "k8s.io/kubernetes/pkg/apis/batch/install"
	"k8s.io/kubernetes/pkg/apis/certificates"
	_ "k8s.io/kubernetes/pkg/apis/certificates/install"
	"k8s.io/kubernetes/pkg/apis/coordination"
	_ "k8s.io/kubernetes/pkg/apis/coordination/install"
	api "k8s.io/kubernetes/pkg/apis/core"
	_ "k8s.io/kubernetes/pkg/apis/core/install"
	"k8s.io/kubernetes/pkg/apis/events"
	_ "k8s.io/kubernetes/pkg/apis/events/install"
	"k8s.io/kubernetes/pkg/apis/extensions"
	_ "k8s.io/kubernetes/pkg/apis/extensions/install"
	"k8s.io/kubernetes/pkg/apis/imagepolicy"
	_ "k8s.io/kubernetes/pkg/apis/imagepolicy/install"
	"k8s.io/kubernetes/pkg/apis/networking"
	_ "k8s.io/kubernetes/pkg/apis/networking/install"
	"k8s.io/kubernetes/pkg/apis/policy"
	_ "k8s.io/kubernetes/pkg/apis/policy/install"
	"k8s.io/kubernetes/pkg/apis/rbac"
	_ "k8s.io/kubernetes/pkg/apis/rbac/install"
	"k8s.io/kubernetes/pkg/apis/scheduling"
	_ "k8s.io/kubernetes/pkg/apis/scheduling/install"
	"k8s.io/kubernetes/pkg/apis/settings"
	_ "k8s.io/kubernetes/pkg/apis/settings/install"
	"k8s.io/kubernetes/pkg/apis/storage"
	_ "k8s.io/kubernetes/pkg/apis/storage/install"
	"mime"
	"os"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

var (
	Groups            = make(map[string]TestGroup)
	Default           TestGroup
	Autoscaling       TestGroup
	Batch             TestGroup
	Extensions        TestGroup
	Apps              TestGroup
	Policy            TestGroup
	Rbac              TestGroup
	Storage           TestGroup
	Admission         TestGroup
	serializer        runtime.SerializerInfo
	storageSerializer runtime.SerializerInfo
)

type TestGroup struct{ externalGroupVersion schema.GroupVersion }

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	if storageMediaType := StorageMediaType(); len(storageMediaType) > 0 {
		var ok bool
		mediaType, _, err := mime.ParseMediaType(storageMediaType)
		if err != nil {
			panic(err)
		}
		storageSerializer, ok = runtime.SerializerInfoForMediaType(legacyscheme.Codecs.SupportedMediaTypes(), mediaType)
		if !ok {
			panic(fmt.Sprintf("no serializer for %s", storageMediaType))
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
			Groups[groupVersion.Group] = TestGroup{externalGroupVersion: groupVersion}
		}
	}
	if _, ok := Groups[api.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: api.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(api.GroupName)[0].Version}
		Groups[api.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[extensions.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: extensions.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(extensions.GroupName)[0].Version}
		Groups[extensions.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[autoscaling.GroupName]; !ok {
		internalTypes := make(map[string]reflect.Type)
		for k, t := range legacyscheme.Scheme.KnownTypes(extensions.SchemeGroupVersion) {
			if k == "Scale" {
				continue
			}
			internalTypes[k] = t
		}
		externalGroupVersion := schema.GroupVersion{Group: autoscaling.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(autoscaling.GroupName)[0].Version}
		Groups[autoscaling.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[autoscaling.GroupName+"IntraGroup"]; !ok {
		internalTypes := make(map[string]reflect.Type)
		for k, t := range legacyscheme.Scheme.KnownTypes(extensions.SchemeGroupVersion) {
			if k == "Scale" {
				internalTypes[k] = t
				break
			}
		}
		externalGroupVersion := schema.GroupVersion{Group: autoscaling.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(autoscaling.GroupName)[0].Version}
		Groups[autoscaling.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[batch.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: batch.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(batch.GroupName)[0].Version}
		Groups[batch.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[apps.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: apps.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(apps.GroupName)[0].Version}
		Groups[apps.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[policy.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: policy.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(policy.GroupName)[0].Version}
		Groups[policy.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[rbac.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: rbac.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(rbac.GroupName)[0].Version}
		Groups[rbac.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[scheduling.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: scheduling.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(scheduling.GroupName)[0].Version}
		Groups[scheduling.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[settings.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: settings.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(settings.GroupName)[0].Version}
		Groups[settings.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[storage.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: storage.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(storage.GroupName)[0].Version}
		Groups[storage.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[certificates.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: certificates.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(certificates.GroupName)[0].Version}
		Groups[certificates.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[imagepolicy.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: imagepolicy.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(imagepolicy.GroupName)[0].Version}
		Groups[imagepolicy.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[authorization.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: authorization.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(authorization.GroupName)[0].Version}
		Groups[authorization.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[admissionregistration.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: admissionregistration.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(admissionregistration.GroupName)[0].Version}
		Groups[admissionregistration.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[admission.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: admission.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(admission.GroupName)[0].Version}
		Groups[admission.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[networking.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: networking.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(networking.GroupName)[0].Version}
		Groups[networking.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[events.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: events.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(events.GroupName)[0].Version}
		Groups[events.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[coordination.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: coordination.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(coordination.GroupName)[0].Version}
		Groups[coordination.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	if _, ok := Groups[auditregistration.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: auditregistration.GroupName, Version: legacyscheme.Scheme.PrioritizedVersionsForGroup(auditregistration.GroupName)[0].Version}
		Groups[auditregistration.GroupName] = TestGroup{externalGroupVersion: externalGroupVersion}
	}
	Default = Groups[api.GroupName]
	Autoscaling = Groups[autoscaling.GroupName]
	Batch = Groups[batch.GroupName]
	Apps = Groups[apps.GroupName]
	Policy = Groups[policy.GroupName]
	Extensions = Groups[extensions.GroupName]
	Rbac = Groups[rbac.GroupName]
	Storage = Groups[storage.GroupName]
	Admission = Groups[admission.GroupName]
}
func (g TestGroup) GroupVersion() *schema.GroupVersion {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	copyOfGroupVersion := g.externalGroupVersion
	return &copyOfGroupVersion
}
func (g TestGroup) Codec() runtime.Codec {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if serializer.Serializer == nil {
		return legacyscheme.Codecs.LegacyCodec(g.externalGroupVersion)
	}
	return legacyscheme.Codecs.CodecForVersions(serializer.Serializer, legacyscheme.Codecs.UniversalDeserializer(), schema.GroupVersions{g.externalGroupVersion}, nil)
}
func StorageMediaType() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return os.Getenv("KUBE_TEST_API_STORAGE_TYPE")
}
func (g TestGroup) StorageCodec() runtime.Codec {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s := storageSerializer.Serializer
	if s == nil {
		return legacyscheme.Codecs.LegacyCodec(g.externalGroupVersion)
	}
	if !storageSerializer.EncodesAsText {
		s = runtime.NewBase64Serializer(s, s)
	}
	ds := recognizer.NewDecoder(s, legacyscheme.Codecs.UniversalDeserializer())
	return legacyscheme.Codecs.CodecForVersions(s, ds, schema.GroupVersions{g.externalGroupVersion}, nil)
}
func (g TestGroup) SelfLink(resource, name string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return g.ResourcePathWithPrefix("", resource, namespace, name)
}
func (g TestGroup) SubResourcePath(resource, namespace, name, sub string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	path := g.ResourcePathWithPrefix("", resource, namespace, name)
	if sub != "" {
		path = path + "/" + sub
	}
	return path
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

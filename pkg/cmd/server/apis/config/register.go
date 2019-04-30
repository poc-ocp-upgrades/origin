package config

import (
	"github.com/openshift/origin/pkg/build/apis/build"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kubernetes/pkg/apis/core"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)

const GroupName = ""

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}

func Kind(kind string) schema.GroupKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}
func Resource(resource string) schema.GroupResource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder	= runtime.NewSchemeBuilder(addKnownTypes, core.AddToScheme, build.AddToScheme)
	InstallLegacy	= SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddIgnoredConversionType(&metav1.TypeMeta{}, &metav1.TypeMeta{}); err != nil {
		return err
	}
	scheme.AddKnownTypes(SchemeGroupVersion, KnownTypes...)
	return nil
}

var KnownTypes = []runtime.Object{&MasterConfig{}, &NodeConfig{}, &SessionSecrets{}, &BasicAuthPasswordIdentityProvider{}, &AllowAllPasswordIdentityProvider{}, &DenyAllPasswordIdentityProvider{}, &HTPasswdPasswordIdentityProvider{}, &LDAPPasswordIdentityProvider{}, &KeystonePasswordIdentityProvider{}, &RequestHeaderIdentityProvider{}, &GitHubIdentityProvider{}, &GitLabIdentityProvider{}, &GoogleIdentityProvider{}, &OpenIDIdentityProvider{}, &LDAPSyncConfig{}, &DefaultAdmissionConfig{}, &BuildDefaultsConfig{}, &BuildOverridesConfig{}}

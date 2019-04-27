package legacy

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/apis/core"
	templatev1 "github.com/openshift/api/template/v1"
	"github.com/openshift/origin/pkg/template/apis/template"
	templatev1helpers "github.com/openshift/origin/pkg/template/apis/template/v1"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"
)

func InstallInternalLegacyTemplate(scheme *runtime.Scheme) {
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
	InstallExternalLegacyTemplate(scheme)
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedInternalTemplateTypes, core.AddToScheme, corev1conversions.AddToScheme, templatev1helpers.RegisterDefaults, templatev1helpers.RegisterConversions)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func InstallExternalLegacyTemplate(scheme *runtime.Scheme) {
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
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedTemplateTypes, corev1.AddToScheme)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func addUngroupifiedTemplateTypes(scheme *runtime.Scheme) error {
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
	types := []runtime.Object{&templatev1.Template{}, &templatev1.TemplateList{}}
	scheme.AddKnownTypes(GroupVersion, types...)
	scheme.AddKnownTypeWithName(GroupVersion.WithKind("TemplateConfig"), &templatev1.Template{})
	scheme.AddKnownTypeWithName(GroupVersion.WithKind("ProcessedTemplate"), &templatev1.Template{})
	return nil
}
func addUngroupifiedInternalTemplateTypes(scheme *runtime.Scheme) error {
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
	types := []runtime.Object{&template.Template{}, &template.TemplateList{}}
	scheme.AddKnownTypes(InternalGroupVersion, types...)
	scheme.AddKnownTypeWithName(InternalGroupVersion.WithKind("TemplateConfig"), &template.Template{})
	scheme.AddKnownTypeWithName(InternalGroupVersion.WithKind("ProcessedTemplate"), &template.Template{})
	return nil
}

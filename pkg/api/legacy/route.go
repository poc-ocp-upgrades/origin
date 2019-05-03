package legacy

import (
	routev1 "github.com/openshift/api/route/v1"
	"github.com/openshift/origin/pkg/route/apis/route"
	routev1helpers "github.com/openshift/origin/pkg/route/apis/route/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/apis/core"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"
)

func InstallInternalLegacyRoute(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	InstallExternalLegacyRoute(scheme)
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedInternalRouteTypes, core.AddToScheme, corev1conversions.AddToScheme, addLegacyRouteFieldSelectorKeyConversions, routev1helpers.RegisterDefaults, routev1helpers.RegisterConversions)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func InstallExternalLegacyRoute(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedRouteTypes, corev1.AddToScheme)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func addUngroupifiedRouteTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	types := []runtime.Object{&routev1.Route{}, &routev1.RouteList{}}
	scheme.AddKnownTypes(GroupVersion, types...)
	return nil
}
func addUngroupifiedInternalRouteTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(InternalGroupVersion, &route.Route{}, &route.RouteList{})
	return nil
}
func addLegacyRouteFieldSelectorKeyConversions(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddFieldLabelConversionFunc(GroupVersion.WithKind("Route"), legacyRouteFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	return nil
}
func legacyRouteFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "spec.path", "spec.host", "spec.to.name":
		return label, value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}

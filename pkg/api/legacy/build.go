package legacy

import (
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/origin/pkg/api/apihelpers"
	"github.com/openshift/origin/pkg/build/apis/build"
	buildv1helpers "github.com/openshift/origin/pkg/build/apis/build/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/apis/core"
)

func InstallInternalLegacyBuild(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	InstallExternalLegacyBuild(scheme)
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedInternalBuildTypes, core.AddToScheme, addLegacyBuildFieldSelectorKeyConversions, buildv1helpers.AddConversionFuncs, buildv1helpers.RegisterDefaults, buildv1helpers.RegisterConversions)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func InstallExternalLegacyBuild(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedBuildTypes, corev1.AddToScheme)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func addUngroupifiedBuildTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	types := []runtime.Object{&buildv1.Build{}, &buildv1.BuildList{}, &buildv1.BuildConfig{}, &buildv1.BuildConfigList{}, &buildv1.BuildLog{}, &buildv1.BuildRequest{}, &buildv1.BuildLogOptions{}, &buildv1.BinaryBuildRequestOptions{}}
	scheme.AddKnownTypes(GroupVersion, types...)
	return nil
}
func addUngroupifiedInternalBuildTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(InternalGroupVersion, &build.Build{}, &build.BuildList{}, &build.BuildConfig{}, &build.BuildConfigList{}, &build.BuildLog{}, &build.BuildRequest{}, &build.BuildLogOptions{}, &build.BinaryBuildRequestOptions{})
	return nil
}
func addLegacyBuildFieldSelectorKeyConversions(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddFieldLabelConversionFunc(GroupVersion.WithKind("Build"), legacyBuildFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	if err := scheme.AddFieldLabelConversionFunc(GroupVersion.WithKind("BuildConfig"), apihelpers.LegacyMetaV1FieldSelectorConversionWithName); err != nil {
		return err
	}
	return nil
}
func legacyBuildFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "status", "podName":
		return label, value, nil
	default:
		return apihelpers.LegacyMetaV1FieldSelectorConversionWithName(label, value)
	}
}

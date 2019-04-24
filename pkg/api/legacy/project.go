package legacy

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	projectv1 "github.com/openshift/api/project/v1"
	"github.com/openshift/origin/pkg/api/apihelpers"
	"github.com/openshift/origin/pkg/project/apis/project"
	projectv1helpers "github.com/openshift/origin/pkg/project/apis/project/v1"
)

func InstallInternalLegacyProject(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	InstallExternalLegacyProject(scheme)
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedInternalProjectTypes, addLegacyProjectFieldSelectorKeyConversions, projectv1helpers.RegisterDefaults, projectv1helpers.RegisterConversions)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func InstallExternalLegacyProject(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedProjectTypes)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func addUngroupifiedProjectTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	types := []runtime.Object{&projectv1.Project{}, &projectv1.ProjectList{}, &projectv1.ProjectRequest{}}
	scheme.AddKnownTypes(GroupVersion, types...)
	return nil
}
func addUngroupifiedInternalProjectTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(InternalGroupVersion, &project.Project{}, &project.ProjectList{}, &project.ProjectRequest{})
	return nil
}
func addLegacyProjectFieldSelectorKeyConversions(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddFieldLabelConversionFunc(GroupVersion.WithKind("Project"), legacyProjectFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	return nil
}
func legacyProjectFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "status.phase":
		return label, value, nil
	default:
		return apihelpers.LegacyMetaV1FieldSelectorConversionWithName(label, value)
	}
}

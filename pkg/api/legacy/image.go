package legacy

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/apis/core"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"
	"github.com/openshift/api/image/docker10"
	"github.com/openshift/api/image/dockerpre012"
	imagev1 "github.com/openshift/api/image/v1"
	"github.com/openshift/origin/pkg/api/apihelpers"
	"github.com/openshift/origin/pkg/image/apis/image"
	imagev1helpers "github.com/openshift/origin/pkg/image/apis/image/v1"
)

func InstallInternalLegacyImage(scheme *runtime.Scheme) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	InstallExternalLegacyImage(scheme)
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedInternalImageTypes, core.AddToScheme, corev1conversions.AddToScheme, addLegacyImageFieldSelectorKeyConversions, imagev1helpers.RegisterDefaults, imagev1helpers.RegisterConversions)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func InstallExternalLegacyImage(scheme *runtime.Scheme) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedImageTypes, docker10.AddToSchemeInCoreGroup, dockerpre012.AddToSchemeInCoreGroup, corev1.AddToScheme)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func addUngroupifiedImageTypes(scheme *runtime.Scheme) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	types := []runtime.Object{&imagev1.Image{}, &imagev1.ImageList{}, &imagev1.ImageSignature{}, &imagev1.ImageStream{}, &imagev1.ImageStreamList{}, &imagev1.ImageStreamMapping{}, &imagev1.ImageStreamTag{}, &imagev1.ImageStreamTagList{}, &imagev1.ImageStreamImage{}, &imagev1.ImageStreamImport{}}
	scheme.AddKnownTypes(GroupVersion, types...)
	return nil
}
func addUngroupifiedInternalImageTypes(scheme *runtime.Scheme) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(InternalGroupVersion, &image.Image{}, &image.ImageList{}, &image.DockerImage{}, &image.ImageSignature{}, &image.ImageStream{}, &image.ImageStreamList{}, &image.ImageStreamMapping{}, &image.ImageStreamTag{}, &image.ImageStreamTagList{}, &image.ImageStreamImage{}, &image.ImageStreamImport{})
	return nil
}
func addLegacyImageFieldSelectorKeyConversions(scheme *runtime.Scheme) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddFieldLabelConversionFunc(GroupVersion.WithKind("ImageStream"), legacyImageStreamFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	return nil
}
func legacyImageStreamFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "spec.dockerImageRepository", "status.dockerImageRepository":
		return label, value, nil
	default:
		return apihelpers.LegacyMetaV1FieldSelectorConversionWithName(label, value)
	}
}

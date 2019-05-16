package image

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/apis/core"
)

const (
	GroupName = "image.openshift.io"
)

var (
	GroupVersion       = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	schemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes, core.AddToScheme)
	Install            = schemeBuilder.AddToScheme
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	AddToScheme        = schemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme.AddKnownTypes(SchemeGroupVersion, &Image{}, &ImageList{}, &DockerImage{}, &ImageSignature{}, &ImageStream{}, &ImageStreamList{}, &ImageStreamMapping{}, &ImageStreamTag{}, &ImageStreamTagList{}, &ImageStreamImage{}, &ImageStreamLayers{}, &ImageStreamImport{}, &core.SecretList{})
	return nil
}

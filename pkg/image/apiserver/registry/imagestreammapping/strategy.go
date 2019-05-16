package imagestreammapping

import (
	"context"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apis/image/validation"
	"github.com/openshift/origin/pkg/image/apiserver/registryhostname"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

type Strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
	registryHostRetriever registryhostname.RegistryHostnameRetriever
}

func NewStrategy(registryHost registryhostname.RegistryHostnameRetriever) Strategy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return Strategy{ObjectTyper: legacyscheme.Scheme, NameGenerator: names.SimpleNameGenerator, registryHostRetriever: registryHost}
}
func (s Strategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (s Strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ism := obj.(*imageapi.ImageStreamMapping)
	if len(ism.Image.DockerImageReference) == 0 {
		internalRegistry, ok := s.registryHostRetriever.InternalRegistryHostname()
		if ok {
			ism.Image.DockerImageReference = imageapi.DockerImageReference{Registry: internalRegistry, Namespace: ism.Namespace, Name: ism.Name, ID: ism.Image.Name}.Exact()
		}
	}
	ism.Image.Signatures = nil
}
func (s Strategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (s Strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mapping := obj.(*imageapi.ImageStreamMapping)
	return validation.ValidateImageStreamMapping(mapping)
}

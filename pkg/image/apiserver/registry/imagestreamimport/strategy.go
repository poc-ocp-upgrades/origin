package imagestreamimport

import (
	"context"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apis/image/validation"
	"github.com/openshift/origin/pkg/image/apis/image/validation/whitelist"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

type strategy struct {
	runtime.ObjectTyper
	registryWhitelister whitelist.RegistryWhitelister
}

func NewStrategy(rw whitelist.RegistryWhitelister) *strategy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &strategy{ObjectTyper: legacyscheme.Scheme, registryWhitelister: rw}
}
func (s *strategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (s *strategy) GenerateName(string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func (s *strategy) Canonicalize(runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (s *strategy) ValidateAllowedRegistries(isi *imageapi.ImageStreamImport) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := field.ErrorList{}
	validate := func(path *field.Path, name string, insecure bool) field.ErrorList {
		ref, _ := imageapi.ParseDockerImageReference(name)
		registryHost, registryPort := ref.RegistryHostPort(insecure)
		return validation.ValidateRegistryAllowedForImport(s.registryWhitelister, path.Child("from", "name"), ref.Name, registryHost, registryPort)
	}
	if spec := isi.Spec.Repository; spec != nil && spec.From.Kind == "DockerImage" {
		errs = append(errs, validate(field.NewPath("spec").Child("repository"), spec.From.Name, spec.ImportPolicy.Insecure)...)
	}
	if len(isi.Spec.Images) > 0 {
		for i, image := range isi.Spec.Images {
			errs = append(errs, validate(field.NewPath("spec").Child("images").Index(i), image.From.Name, image.ImportPolicy.Insecure)...)
		}
	}
	return errs
}
func (s *strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newIST := obj.(*imageapi.ImageStreamImport)
	newIST.Status = imageapi.ImageStreamImportStatus{}
}
func (s *strategy) PrepareImageForCreate(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	image := obj.(*imageapi.Image)
	image.Signatures = nil
	image.DockerImageManifest = ""
	image.DockerImageConfig = ""
}
func (s *strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	isi := obj.(*imageapi.ImageStreamImport)
	return validation.ValidateImageStreamImport(isi)
}

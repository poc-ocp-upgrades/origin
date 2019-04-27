package image

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apis/image/validation"
	"github.com/openshift/origin/pkg/image/util"
)

const managedSignatureAnnotation = "image.openshift.io/managed-signature"

type imageStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = imageStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}
var _ rest.GarbageCollectionDeleteStrategy = imageStrategy{}

func (imageStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
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
	return rest.Unsupported
}
func (imageStrategy) NamespaceScoped() bool {
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
	return false
}
func (s imageStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
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
	newImage := obj.(*imageapi.Image)
	if err := util.InternalImageWithMetadata(newImage); err != nil {
		utilruntime.HandleError(fmt.Errorf("Unable to update image metadata for %q: %v", newImage.Name, err))
	}
	if newImage.Annotations[imageapi.ImageManifestBlobStoredAnnotation] == "true" {
		newImage.DockerImageManifest = ""
		newImage.DockerImageConfig = ""
	}
	removeManagedSignatureAnnotation(newImage)
}
func (imageStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
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
	image := obj.(*imageapi.Image)
	return validation.ValidateImage(image)
}
func (imageStrategy) AllowCreateOnUpdate() bool {
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
	return false
}
func (imageStrategy) AllowUnconditionalUpdate() bool {
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
	return false
}
func (imageStrategy) Canonicalize(obj runtime.Object) {
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
}
func (s imageStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
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
	newImage := obj.(*imageapi.Image)
	oldImage := old.(*imageapi.Image)
	newImage.DockerImageMetadata = oldImage.DockerImageMetadata
	newImage.DockerImageMetadataVersion = oldImage.DockerImageMetadataVersion
	newImage.DockerImageLayers = oldImage.DockerImageLayers
	if oldImage.DockerImageSignatures != nil {
		newImage.DockerImageSignatures = make([][]byte, 0, len(oldImage.DockerImageSignatures))
		newImage.DockerImageSignatures = append(newImage.DockerImageSignatures, oldImage.DockerImageSignatures...)
	}
	var err error
	if newImage.DockerImageManifest != oldImage.DockerImageManifest {
		ok := true
		if len(newImage.DockerImageManifest) > 0 {
			ok, err = util.ManifestMatchesImage(oldImage, []byte(newImage.DockerImageManifest))
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("attempted to validate that a manifest change to %q matched the signature, but failed: %v", oldImage.Name, err))
			}
		}
		if !ok {
			newImage.DockerImageManifest = oldImage.DockerImageManifest
		}
	}
	if newImage.DockerImageConfig != oldImage.DockerImageConfig {
		ok := true
		if len(newImage.DockerImageConfig) > 0 {
			ok, err = util.ImageConfigMatchesImage(newImage, []byte(newImage.DockerImageConfig))
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("attempted to validate that a new config for %q mentioned in the manifest, but failed: %v", oldImage.Name, err))
			}
		}
		if !ok {
			newImage.DockerImageConfig = oldImage.DockerImageConfig
		}
	}
	if err = util.InternalImageWithMetadata(newImage); err != nil {
		utilruntime.HandleError(fmt.Errorf("Unable to update image metadata for %q: %v", newImage.Name, err))
	}
	if newImage.Annotations[imageapi.ImageManifestBlobStoredAnnotation] == "true" {
		newImage.DockerImageManifest = ""
		newImage.DockerImageConfig = ""
	}
	removeManagedSignatureAnnotation(newImage)
}
func (imageStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
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
	return validation.ValidateImageUpdate(obj.(*imageapi.Image), old.(*imageapi.Image))
}
func removeManagedSignatureAnnotation(img *imageapi.Image) {
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
	for i := range img.Signatures {
		delete(img.Signatures[i].Annotations, managedSignatureAnnotation)
	}
}

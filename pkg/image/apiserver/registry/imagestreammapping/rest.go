package imagestreammapping

import (
	"context"
	imagegroup "github.com/openshift/api/image"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apiserver/registry/image"
	"github.com/openshift/origin/pkg/image/apiserver/registry/imagestream"
	"github.com/openshift/origin/pkg/image/apiserver/registryhostname"
	"k8s.io/apimachinery/pkg/api/errors"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/klog"
	kapihelper "k8s.io/kubernetes/pkg/apis/core/helper"
)

const maxRetriesOnConflict = 10

type REST struct {
	imageRegistry       image.Registry
	imageStreamRegistry imagestream.Registry
	strategy            Strategy
}

var _ rest.Creater = &REST{}
var _ rest.Scoper = &REST{}

func NewREST(imageRegistry image.Registry, imageStreamRegistry imagestream.Registry, registry registryhostname.RegistryHostnameRetriever) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &REST{imageRegistry: imageRegistry, imageStreamRegistry: imageStreamRegistry, strategy: NewStrategy(registry)}
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &imageapi.ImageStreamMapping{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (s *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := rest.BeforeCreate(s.strategy, ctx, obj); err != nil {
		return nil, err
	}
	if err := createValidation(obj.DeepCopyObject()); err != nil {
		return nil, err
	}
	mapping := obj.(*imageapi.ImageStreamMapping)
	stream, err := s.findStreamForMapping(ctx, mapping)
	if err != nil {
		return nil, err
	}
	image := mapping.Image
	tag := mapping.Tag
	if len(tag) == 0 {
		tag = imageapi.DefaultImageTag
	}
	imageCreateErr := s.imageRegistry.CreateImage(ctx, &image)
	if imageCreateErr != nil && !errors.IsAlreadyExists(imageCreateErr) {
		return nil, imageCreateErr
	}
	ref := image.DockerImageReference
	if errors.IsAlreadyExists(imageCreateErr) && image.Annotations[imageapi.ManagedByOpenShiftAnnotation] == "true" {
		if streamRef, err := imageapi.DockerImageReferenceForStream(stream); err == nil {
			streamRef.ID = image.Name
			ref = streamRef.Exact()
		} else {
			klog.V(4).Infof("Failed to get dockerImageReference for stream %s/%s: %v", stream.Namespace, stream.Name, err)
		}
	}
	next := imageapi.TagEvent{Created: metav1.Now(), DockerImageReference: ref, Image: image.Name}
	err = wait.ExponentialBackoff(wait.Backoff{Steps: maxRetriesOnConflict}, func() (bool, error) {
		lastEvent := imageapi.LatestTaggedImage(stream, tag)
		next.Generation = stream.Generation
		if !imageapi.AddTagEventToImageStream(stream, tag, next) {
			return true, nil
		}
		imageapi.UpdateTrackingTags(stream, tag, next)
		_, err := s.imageStreamRegistry.UpdateImageStreamStatus(ctx, stream, false, &metav1.UpdateOptions{})
		if err == nil {
			return true, nil
		}
		if !errors.IsConflict(err) {
			return false, err
		}
		latestStream, findLatestErr := s.findStreamForMapping(ctx, mapping)
		if findLatestErr != nil {
			return false, findLatestErr
		}
		if lastEvent == nil {
			stream = latestStream
			return false, nil
		}
		newerEvent := imageapi.LatestTaggedImage(latestStream, tag)
		lastEvent.Generation = newerEvent.Generation
		lastEvent.Created = newerEvent.Created
		if kapihelper.Semantic.DeepEqual(lastEvent, newerEvent) {
			stream = latestStream
			return false, nil
		}
		return false, err
	})
	if err != nil {
		return nil, err
	}
	return &metav1.Status{Status: metav1.StatusSuccess}, nil
}
func (s *REST) findStreamForMapping(ctx context.Context, mapping *imageapi.ImageStreamMapping) (*imageapi.ImageStream, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(mapping.Name) > 0 {
		return s.imageStreamRegistry.GetImageStream(ctx, mapping.Name, &metav1.GetOptions{})
	}
	if len(mapping.DockerImageRepository) != 0 {
		list, err := s.imageStreamRegistry.ListImageStreams(ctx, &metainternal.ListOptions{})
		if err != nil {
			return nil, err
		}
		for i := range list.Items {
			if mapping.DockerImageRepository == list.Items[i].Spec.DockerImageRepository {
				return &list.Items[i], nil
			}
		}
		return nil, errors.NewInvalid(imagegroup.Kind("ImageStreamMapping"), "", field.ErrorList{field.NotFound(field.NewPath("dockerImageStream"), mapping.DockerImageRepository)})
	}
	return nil, errors.NewNotFound(imagegroup.Resource("imagestream"), "")
}

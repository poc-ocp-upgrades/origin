package imagestream

import (
	"context"
	"fmt"
	"strings"
	authorizationapi "k8s.io/api/authorization/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/authentication/user"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/storage/names"
	authorizationclient "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	kapihelper "k8s.io/kubernetes/pkg/apis/core/helper"
	authorizationutil "github.com/openshift/origin/pkg/authorization/util"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apis/image/validation"
	"github.com/openshift/origin/pkg/image/apis/image/validation/whitelist"
	imageadmission "github.com/openshift/origin/pkg/image/apiserver/admission/limitrange"
	"github.com/openshift/origin/pkg/image/apiserver/registryhostname"
)

type ResourceGetter interface {
	Get(context.Context, string, *metav1.GetOptions) (runtime.Object, error)
}
type Strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
	registryHostnameRetriever	registryhostname.RegistryHostnameRetriever
	tagVerifier			*TagVerifier
	limitVerifier			imageadmission.LimitVerifier
	registryWhitelister		whitelist.RegistryWhitelister
	imageStreamGetter		ResourceGetter
}

func NewStrategy(registryHostname registryhostname.RegistryHostnameRetriever, subjectAccessReviewClient authorizationclient.SubjectAccessReviewInterface, limitVerifier imageadmission.LimitVerifier, registryWhitelister whitelist.RegistryWhitelister, imageStreamGetter ResourceGetter) Strategy {
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
	return Strategy{ObjectTyper: legacyscheme.Scheme, NameGenerator: names.SimpleNameGenerator, registryHostnameRetriever: registryHostname, tagVerifier: &TagVerifier{subjectAccessReviewClient}, limitVerifier: limitVerifier, registryWhitelister: registryWhitelister, imageStreamGetter: imageStreamGetter}
}
func (s Strategy) NamespaceScoped() bool {
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
	return true
}
func collapseEmptyStatusTags(stream *imageapi.ImageStream) {
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
	for tag, ref := range stream.Status.Tags {
		if len(ref.Items) == 0 && len(ref.Conditions) == 0 {
			delete(stream.Status.Tags, tag)
		}
	}
}
func (s Strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
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
	stream := obj.(*imageapi.ImageStream)
	stream.Status = imageapi.ImageStreamStatus{DockerImageRepository: s.dockerImageRepository(stream, false), Tags: make(map[string]imageapi.TagEventList)}
	stream.Generation = 1
	for tag, ref := range stream.Spec.Tags {
		ref.Generation = &stream.Generation
		stream.Spec.Tags[tag] = ref
	}
	collapseEmptyStatusTags(stream)
}
func (s Strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
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
	stream := obj.(*imageapi.ImageStream)
	var errs field.ErrorList
	if err := s.validateTagsAndLimits(ctx, nil, stream); err != nil {
		errs = append(errs, field.InternalError(field.NewPath(""), err))
	}
	errs = append(errs, validation.ValidateImageStreamWithWhitelister(s.registryWhitelister, stream)...)
	return errs
}
func (s Strategy) validateTagsAndLimits(ctx context.Context, oldStream, newStream *imageapi.ImageStream) error {
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
	user, ok := apirequest.UserFrom(ctx)
	if !ok {
		return kerrors.NewForbidden(schema.GroupResource{Resource: "imagestreams"}, newStream.Name, fmt.Errorf("no user context available"))
	}
	errs := s.tagVerifier.Verify(oldStream, newStream, user)
	errs = append(errs, s.tagsChanged(oldStream, newStream)...)
	if len(errs) > 0 {
		return kerrors.NewInvalid(schema.GroupKind{Kind: "imagestreams"}, newStream.Name, errs)
	}
	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		ns = newStream.Namespace
	}
	return s.limitVerifier.VerifyLimits(ns, newStream)
}
func (s Strategy) AllowCreateOnUpdate() bool {
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
func (Strategy) AllowUnconditionalUpdate() bool {
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
func (s Strategy) dockerImageRepository(stream *imageapi.ImageStream, allowNamespaceDefaulting bool) string {
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
	registry, ok := s.registryHostnameRetriever.InternalRegistryHostname()
	if !ok {
		return stream.Spec.DockerImageRepository
	}
	if len(stream.Namespace) == 0 && allowNamespaceDefaulting {
		stream.Namespace = metav1.NamespaceDefault
	}
	ref := imageapi.DockerImageReference{Registry: registry, Namespace: stream.Namespace, Name: stream.Name}
	return ref.String()
}
func (s Strategy) publicDockerImageRepository(stream *imageapi.ImageStream) string {
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
	externalHostname, ok := s.registryHostnameRetriever.ExternalRegistryHostname()
	if !ok {
		return ""
	}
	ref := imageapi.DockerImageReference{Registry: externalHostname, Namespace: stream.Namespace, Name: stream.Name}
	return ref.String()
}
func parseFromReference(stream *imageapi.ImageStream, from *kapi.ObjectReference) (string, string, error) {
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
	splitChar := ""
	refType := ""
	switch from.Kind {
	case "ImageStreamTag":
		splitChar = ":"
		refType = "tag"
	case "ImageStreamImage":
		splitChar = "@"
		refType = "id"
	default:
		return "", "", fmt.Errorf("invalid from.kind %q - only ImageStreamTag and ImageStreamImage are allowed", from.Kind)
	}
	parts := strings.Split(from.Name, splitChar)
	switch len(parts) {
	case 1:
		return stream.Name, from.Name, nil
	case 2:
		return parts[0], parts[1], nil
	default:
		return "", "", fmt.Errorf("invalid from.name %q - it must be of the form <%s> or <stream>%s<%s>", from.Name, refType, splitChar, refType)
	}
}
func (s Strategy) tagsChanged(old, stream *imageapi.ImageStream) field.ErrorList {
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
	internalRegistry, hasInternalRegistry := s.registryHostnameRetriever.InternalRegistryHostname()
	var errs field.ErrorList
	oldTags := map[string]imageapi.TagReference{}
	if old != nil && old.Spec.Tags != nil {
		oldTags = old.Spec.Tags
	}
	for tag, tagRef := range stream.Spec.Tags {
		if oldRef, ok := oldTags[tag]; ok && !tagRefChanged(oldRef, tagRef, stream.Namespace) {
			continue
		}
		if tagRef.From == nil {
			continue
		}
		klog.V(5).Infof("Detected changed tag %s in %s/%s", tag, stream.Namespace, stream.Name)
		generation := stream.Generation
		tagRef.Generation = &generation
		fromPath := field.NewPath("spec", "tags").Key(tag).Child("from")
		if tagRef.From.Kind == "DockerImage" && len(tagRef.From.Name) > 0 {
			if tagRef.Reference {
				event, err := tagReferenceToTagEvent(stream, tagRef, "")
				if err != nil {
					errs = append(errs, field.Invalid(fromPath, tagRef.From, err.Error()))
					continue
				}
				stream.Spec.Tags[tag] = tagRef
				imageapi.AddTagEventToImageStream(stream, tag, *event)
			}
			continue
		}
		tagRefStreamName, tagOrID, err := parseFromReference(stream, tagRef.From)
		if err != nil {
			errs = append(errs, field.Invalid(fromPath.Child("name"), tagRef.From.Name, "must be of the form <tag>, <repo>:<tag>, <id>, or <repo>@<id>"))
			continue
		}
		streamRef := stream
		streamRefNamespace := tagRef.From.Namespace
		if len(streamRefNamespace) == 0 {
			streamRefNamespace = stream.Namespace
		}
		if streamRefNamespace != stream.Namespace || tagRefStreamName != stream.Name {
			obj, err := s.imageStreamGetter.Get(apirequest.WithNamespace(apirequest.NewContext(), streamRefNamespace), tagRefStreamName, &metav1.GetOptions{})
			if err != nil {
				if kerrors.IsNotFound(err) {
					errs = append(errs, field.NotFound(fromPath.Child("name"), tagRef.From.Name))
				} else {
					errs = append(errs, field.Invalid(fromPath.Child("name"), tagRef.From.Name, fmt.Sprintf("unable to retrieve image stream: %v", err)))
				}
				continue
			}
			streamRef = obj.(*imageapi.ImageStream)
		}
		event, err := tagReferenceToTagEvent(streamRef, tagRef, tagOrID)
		if err != nil {
			errs = append(errs, field.Invalid(fromPath.Child("name"), tagRef.From.Name, fmt.Sprintf("error generating tag event: %v", err)))
			continue
		}
		if event == nil {
			continue
		}
		if !tagRef.Reference {
			if ref, err := imageapi.ParseDockerImageReference(event.DockerImageReference); err == nil {
				if hasInternalRegistry && ref.Registry == internalRegistry && ref.Namespace == streamRef.Namespace && ref.Name == streamRef.Name {
					ref.Namespace = stream.Namespace
					ref.Name = stream.Name
					event.DockerImageReference = ref.Exact()
				}
			}
		}
		stream.Spec.Tags[tag] = tagRef
		imageapi.AddTagEventToImageStream(stream, tag, *event)
	}
	imageapi.UpdateChangedTrackingTags(stream, old)
	if old == nil && !stream.CreationTimestamp.IsZero() {
		for tag, list := range stream.Status.Tags {
			for _, event := range list.Items {
				event.Created = stream.CreationTimestamp
			}
			stream.Status.Tags[tag] = list
		}
	}
	return errs
}
func tagReferenceToTagEvent(stream *imageapi.ImageStream, tagRef imageapi.TagReference, tagOrID string) (*imageapi.TagEvent, error) {
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
	var (
		event	*imageapi.TagEvent
		err	error
	)
	switch tagRef.From.Kind {
	case "DockerImage":
		event = &imageapi.TagEvent{Created: metav1.Now(), DockerImageReference: tagRef.From.Name}
	case "ImageStreamImage":
		event, err = imageapi.ResolveImageID(stream, tagOrID)
	case "ImageStreamTag":
		event, err = imageapi.LatestTaggedImage(stream, tagOrID), nil
	default:
		err = fmt.Errorf("invalid from.kind %q: it must be DockerImage, ImageStreamImage or ImageStreamTag", tagRef.From.Kind)
	}
	if err != nil {
		return nil, err
	}
	if event != nil && tagRef.Generation != nil {
		event.Generation = *tagRef.Generation
	}
	return event, nil
}
func tagRefChanged(old, next imageapi.TagReference, streamNamespace string) bool {
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
	if next.From == nil {
		return false
	}
	if len(next.From.Kind) == 0 && len(next.From.Name) == 0 {
		return false
	}
	oldFrom := old.From
	if oldFrom == nil {
		oldFrom = &kapi.ObjectReference{}
	}
	oldNamespace := oldFrom.Namespace
	if len(oldNamespace) == 0 {
		oldNamespace = streamNamespace
	}
	nextNamespace := next.From.Namespace
	if len(nextNamespace) == 0 {
		nextNamespace = streamNamespace
	}
	if oldNamespace != nextNamespace {
		return true
	}
	if oldFrom.Name != next.From.Name {
		return true
	}
	if old.Reference != next.Reference {
		return true
	}
	return tagRefGenerationChanged(old, next)
}
func tagRefGenerationChanged(old, next imageapi.TagReference) bool {
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
	switch {
	case old.Generation != nil && next.Generation != nil:
		if *old.Generation == *next.Generation {
			return false
		}
		if *next.Generation == 0 {
			return true
		}
		return false
	default:
		return false
	}
}
func tagEventChanged(old, next imageapi.TagEvent) bool {
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
	return old.Image != next.Image || old.DockerImageReference != next.DockerImageReference || old.Generation > next.Generation
}
func updateSpecTagGenerationsForUpdate(stream, oldStream *imageapi.ImageStream) {
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
	for tag, ref := range stream.Spec.Tags {
		if ref.Generation != nil && *ref.Generation == 0 {
			continue
		}
		if oldRef, ok := oldStream.Spec.Tags[tag]; ok {
			ref.Generation = oldRef.Generation
			stream.Spec.Tags[tag] = ref
		}
	}
}
func ensureSpecTagGenerationsAreSet(stream, oldStream *imageapi.ImageStream) {
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
	oldTags := map[string]imageapi.TagReference{}
	if oldStream != nil && oldStream.Spec.Tags != nil {
		oldTags = oldStream.Spec.Tags
	}
	for tag, ref := range stream.Spec.Tags {
		if oldRef, ok := oldTags[tag]; !ok || tagRefChanged(oldRef, ref, stream.Namespace) {
			ref.Generation = nil
		}
		if ref.Generation != nil && *ref.Generation != 0 {
			continue
		}
		ref.Generation = &stream.Generation
		stream.Spec.Tags[tag] = ref
	}
}
func updateObservedGenerationForStatusUpdate(stream, oldStream *imageapi.ImageStream) {
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
	for tag, newer := range stream.Status.Tags {
		if len(newer.Items) == 0 || newer.Items[0].Generation != 0 {
			continue
		}
		older := oldStream.Status.Tags[tag]
		if len(older.Items) == 0 || !tagEventChanged(older.Items[0], newer.Items[0]) {
			newer.Items[0].Generation = stream.Generation
			stream.Status.Tags[tag] = newer
			continue
		}
		spec, ok := stream.Spec.Tags[tag]
		if !ok || spec.Generation == nil {
			newer.Items[0].Generation = stream.Generation
			stream.Status.Tags[tag] = newer
			continue
		}
		newer.Items[0].Generation = *spec.Generation
		stream.Status.Tags[tag] = newer
	}
}

type TagVerifier struct {
	subjectAccessReviewClient authorizationclient.SubjectAccessReviewInterface
}

func (v *TagVerifier) Verify(old, stream *imageapi.ImageStream, user user.Info) field.ErrorList {
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
	var errors field.ErrorList
	oldTags := map[string]imageapi.TagReference{}
	if old != nil && old.Spec.Tags != nil {
		oldTags = old.Spec.Tags
	}
	for tag, tagRef := range stream.Spec.Tags {
		if tagRef.From == nil {
			continue
		}
		if len(tagRef.From.Namespace) == 0 {
			continue
		}
		if stream.Namespace == tagRef.From.Namespace {
			continue
		}
		if oldRef, ok := oldTags[tag]; ok && !tagRefChanged(oldRef, tagRef, stream.Namespace) {
			continue
		}
		streamName, _, err := parseFromReference(stream, tagRef.From)
		fromPath := field.NewPath("spec", "tags").Key(tag).Child("from")
		if err != nil {
			errors = append(errors, field.Invalid(fromPath.Child("name"), tagRef.From.Name, "must be of the form <tag>, <repo>:<tag>, <id>, or <repo>@<id>"))
			continue
		}
		subjectAccessReview := authorizationutil.AddUserToSAR(user, &authorizationapi.SubjectAccessReview{Spec: authorizationapi.SubjectAccessReviewSpec{ResourceAttributes: &authorizationapi.ResourceAttributes{Namespace: tagRef.From.Namespace, Verb: "get", Group: imageapi.GroupName, Resource: "imagestreams", Subresource: "layers", Name: streamName}}})
		klog.V(4).Infof("Performing SubjectAccessReview for user=%s, groups=%v to %s/%s", user.GetName(), user.GetGroups(), tagRef.From.Namespace, streamName)
		resp, err := v.subjectAccessReviewClient.Create(subjectAccessReview)
		if err != nil || resp == nil || (resp != nil && !resp.Status.Allowed) {
			message := fmt.Sprintf("%s/%s", tagRef.From.Namespace, streamName)
			if resp != nil {
				message = message + fmt.Sprintf(": %q %q", resp.Status.Reason, resp.Status.EvaluationError)
			}
			if err != nil {
				message = message + fmt.Sprintf("- %v", err)
			}
			errors = append(errors, field.Forbidden(fromPath, message))
			continue
		}
	}
	return errors
}
func (Strategy) Canonicalize(obj runtime.Object) {
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
func (s Strategy) prepareForUpdate(obj, old runtime.Object, resetStatus bool) {
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
	oldStream := old.(*imageapi.ImageStream)
	stream := obj.(*imageapi.ImageStream)
	collapseEmptyStatusTags(stream)
	stream.Generation = oldStream.Generation
	if resetStatus {
		stream.Status = oldStream.Status
	}
	stream.Status.DockerImageRepository = s.dockerImageRepository(stream, true)
	updateSpecTagGenerationsForUpdate(stream, oldStream)
	if !kapihelper.Semantic.DeepEqual(oldStream.Spec, stream.Spec) || stream.Generation == 0 {
		stream.Generation = oldStream.Generation + 1
	}
	ensureSpecTagGenerationsAreSet(stream, oldStream)
}
func (s Strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
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
	s.prepareForUpdate(obj, old, true)
}
func (s Strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
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
	stream := obj.(*imageapi.ImageStream)
	oldStream := old.(*imageapi.ImageStream)
	var errs field.ErrorList
	if err := s.validateTagsAndLimits(ctx, oldStream, stream); err != nil {
		errs = append(errs, field.InternalError(field.NewPath(""), err))
	}
	errs = append(errs, validation.ValidateImageStreamUpdateWithWhitelister(s.registryWhitelister, stream, oldStream)...)
	return errs
}
func (s Strategy) Decorate(obj runtime.Object) error {
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
	switch t := obj.(type) {
	case *imageapi.ImageStream:
		t.Status.DockerImageRepository = s.dockerImageRepository(t, true)
		t.Status.PublicDockerImageRepository = s.publicDockerImageRepository(t)
	case *imageapi.ImageStreamList:
		for i := range t.Items {
			is := &t.Items[i]
			is.Status.DockerImageRepository = s.dockerImageRepository(is, true)
			is.Status.PublicDockerImageRepository = s.publicDockerImageRepository(is)
		}
	default:
		return kerrors.NewBadRequest(fmt.Sprintf("not an ImageStream nor ImageStreamList: %v", obj))
	}
	return nil
}

type StatusStrategy struct{ Strategy }

func NewStatusStrategy(strategy Strategy) StatusStrategy {
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
	return StatusStrategy{strategy}
}
func (StatusStrategy) Canonicalize(obj runtime.Object) {
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
func (StatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
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
	oldStream := old.(*imageapi.ImageStream)
	stream := obj.(*imageapi.ImageStream)
	stream.Spec.Tags = oldStream.Spec.Tags
	stream.Spec.DockerImageRepository = oldStream.Spec.DockerImageRepository
	updateObservedGenerationForStatusUpdate(stream, oldStream)
}
func (s StatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
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
	newIS := obj.(*imageapi.ImageStream)
	errs := field.ErrorList{}
	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		ns = newIS.Namespace
	}
	err := s.limitVerifier.VerifyLimits(ns, newIS)
	if err != nil {
		errs = append(errs, field.Forbidden(field.NewPath("imageStream"), err.Error()))
	}
	errs = append(errs, validation.ValidateImageStreamStatusUpdateWithWhitelister(s.registryWhitelister, newIS, old.(*imageapi.ImageStream))...)
	return errs
}

type InternalStrategy struct{ Strategy }

func NewInternalStrategy(strategy Strategy) InternalStrategy {
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
	return InternalStrategy{strategy}
}
func (InternalStrategy) Canonicalize(obj runtime.Object) {
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
func (s InternalStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
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
	stream := obj.(*imageapi.ImageStream)
	stream.Status.DockerImageRepository = s.dockerImageRepository(stream, false)
	stream.Generation = 1
	for tag, ref := range stream.Spec.Tags {
		ref.Generation = &stream.Generation
		stream.Spec.Tags[tag] = ref
	}
}
func (s InternalStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
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
	s.prepareForUpdate(obj, old, false)
}

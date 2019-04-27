package image

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strings"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"github.com/blang/semver"
	"k8s.io/klog"
	"github.com/openshift/api/image"
	"github.com/openshift/origin/pkg/image/apis/image/reference"
	"github.com/openshift/origin/pkg/image/internal/digest"
)

const (
	DockerDefaultNamespace		= "library"
	DockerDefaultRegistry		= "docker.io"
	DockerDefaultV1Registry		= "index." + DockerDefaultRegistry
	DockerDefaultV2Registry		= "registry-1." + DockerDefaultRegistry
	TagReferenceAnnotationTagHidden	= "hidden"
)

var errNoRegistryURLPathAllowed = errors.New("no path after <host>[:<port>] is allowed")
var errNoRegistryURLQueryAllowed = errors.New("no query arguments are allowed after <host>[:<port>]")
var errRegistryURLHostEmpty = errors.New("no host name specified")
var ErrImageStreamImportUnsupported = errors.New("the server does not support directly importing images - create an image stream with tags or the dockerImageRepository field set")
var ErrCircularReference = errors.New("reference tag is circular")
var ErrNotFoundReference = errors.New("reference tag is not found")
var ErrCrossImageStreamReference = errors.New("reference tag points to another imagestream")
var ErrInvalidReference = errors.New("reference tag is invalid")

func ParseImageStreamImageName(input string) (name string, id string, err error) {
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
	segments := strings.SplitN(input, "@", 3)
	switch len(segments) {
	case 2:
		name = segments[0]
		id = segments[1]
		if len(name) == 0 || len(id) == 0 {
			err = fmt.Errorf("image stream image name %q must have a name and ID", input)
		}
	default:
		err = fmt.Errorf("expected exactly one @ in the isimage name %q", input)
	}
	return
}
func ParseImageStreamTagName(istag string) (name string, tag string, err error) {
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
	if strings.Contains(istag, "@") {
		err = fmt.Errorf("%q is an image stream image, not an image stream tag", istag)
		return
	}
	segments := strings.SplitN(istag, ":", 3)
	switch len(segments) {
	case 2:
		name = segments[0]
		tag = segments[1]
		if len(name) == 0 || len(tag) == 0 {
			err = fmt.Errorf("image stream tag name %q must have a name and a tag", istag)
		}
	default:
		err = fmt.Errorf("expected exactly one : delimiter in the istag %q", istag)
	}
	return
}
func ParseDockerImageReference(spec string) (reference.DockerImageReference, error) {
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
	ref, err := reference.Parse(spec)
	if err != nil {
		return ref, err
	}
	return ref, nil
}
func SplitImageStreamTag(nameAndTag string) (name string, tag string, ok bool) {
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
	parts := strings.SplitN(nameAndTag, ":", 2)
	name = parts[0]
	if len(parts) > 1 {
		tag = parts[1]
	}
	if len(tag) == 0 {
		tag = DefaultImageTag
	}
	return name, tag, len(parts) == 2
}
func SplitImageStreamImage(nameAndID string) (name string, id string, ok bool) {
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
	parts := strings.SplitN(nameAndID, "@", 2)
	name = parts[0]
	if len(parts) > 1 {
		id = parts[1]
	}
	return name, id, len(parts) == 2
}
func JoinImageStreamTag(name, tag string) string {
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
	if len(tag) == 0 {
		tag = DefaultImageTag
	}
	return fmt.Sprintf("%s:%s", name, tag)
}
func JoinImageStreamImage(name, id string) string {
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
	return fmt.Sprintf("%s@%s", name, id)
}
func NormalizeImageStreamTag(name string) string {
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
	stripped, tag, ok := SplitImageStreamTag(name)
	if !ok {
		return JoinImageStreamTag(stripped, tag)
	}
	return name
}
func DockerImageReferenceForStream(stream *ImageStream) (DockerImageReference, error) {
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
	spec := stream.Status.DockerImageRepository
	if len(spec) == 0 {
		spec = stream.Spec.DockerImageRepository
	}
	if len(spec) == 0 {
		return DockerImageReference{}, fmt.Errorf("no possible pull spec for %s/%s", stream.Namespace, stream.Name)
	}
	return ParseDockerImageReference(spec)
}
func FollowTagReference(stream *ImageStream, tag string) (finalTag string, ref *TagReference, multiple bool, err error) {
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
	seen := sets.NewString()
	for {
		if seen.Has(tag) {
			return tag, nil, multiple, ErrCircularReference
		}
		seen.Insert(tag)
		tagRef, ok := stream.Spec.Tags[tag]
		if !ok {
			return tag, nil, multiple, ErrNotFoundReference
		}
		if tagRef.From == nil || tagRef.From.Kind != "ImageStreamTag" {
			return tag, &tagRef, multiple, nil
		}
		if tagRef.From.Namespace != "" && tagRef.From.Namespace != stream.ObjectMeta.Namespace {
			return tag, nil, multiple, ErrCrossImageStreamReference
		}
		if strings.Contains(tagRef.From.Name, ":") {
			name, tagref, ok := SplitImageStreamTag(tagRef.From.Name)
			if !ok {
				return tag, nil, multiple, ErrInvalidReference
			}
			if name != stream.ObjectMeta.Name {
				return tag, nil, multiple, ErrCrossImageStreamReference
			}
			tag = tagref
		} else {
			tag = tagRef.From.Name
		}
		multiple = true
	}
}
func LatestImageTagEvent(stream *ImageStream, imageID string) (string, *TagEvent) {
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
		latestTagEvent	*TagEvent
		latestTag	string
	)
	for tag, events := range stream.Status.Tags {
		if len(events.Items) == 0 {
			continue
		}
		for i, event := range events.Items {
			if DigestOrImageMatch(event.Image, imageID) && (latestTagEvent == nil || latestTagEvent != nil && event.Created.After(latestTagEvent.Created.Time)) {
				latestTagEvent = &events.Items[i]
				latestTag = tag
			}
		}
	}
	return latestTag, latestTagEvent
}
func LatestTaggedImage(stream *ImageStream, tag string) *TagEvent {
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
	if len(tag) == 0 {
		tag = DefaultImageTag
	}
	if stream.Status.Tags != nil {
		if history, ok := stream.Status.Tags[tag]; ok {
			if len(history.Items) == 0 {
				return nil
			}
			return &history.Items[0]
		}
	}
	return nil
}
func ResolveLatestTaggedImage(stream *ImageStream, tag string) (string, bool) {
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
	if len(tag) == 0 {
		tag = DefaultImageTag
	}
	return ResolveTagReference(stream, tag, LatestTaggedImage(stream, tag))
}
func ResolveTagReference(stream *ImageStream, tag string, latest *TagEvent) (string, bool) {
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
	if latest == nil {
		return "", false
	}
	return ResolveReferenceForTagEvent(stream, tag, latest), true
}
func ResolveReferenceForTagEvent(stream *ImageStream, tag string, latest *TagEvent) string {
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
	ref, ok := stream.Spec.Tags[tag]
	if !ok {
		return latest.DockerImageReference
	}
	switch ref.ReferencePolicy.Type {
	case LocalTagReferencePolicy:
		local := stream.Status.DockerImageRepository
		if len(local) == 0 || len(latest.Image) == 0 {
			return latest.DockerImageReference
		}
		ref, err := ParseDockerImageReference(local)
		if err != nil {
			return latest.DockerImageReference
		}
		ref.Tag = ""
		ref.ID = latest.Image
		return ref.Exact()
	default:
		return latest.DockerImageReference
	}
}
func DockerImageReferenceForImage(stream *ImageStream, imageID string) (string, bool) {
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
	tag, event := LatestImageTagEvent(stream, imageID)
	if len(tag) == 0 {
		return "", false
	}
	ref, ok := stream.Spec.Tags[tag]
	if !ok {
		return event.DockerImageReference, true
	}
	switch ref.ReferencePolicy.Type {
	case LocalTagReferencePolicy:
		ref, err := ParseDockerImageReference(stream.Status.DockerImageRepository)
		if err != nil {
			return event.DockerImageReference, true
		}
		ref.Tag = ""
		ref.ID = event.Image
		return ref.Exact(), true
	default:
		return event.DockerImageReference, true
	}
}
func DifferentTagEvent(stream *ImageStream, tag string, next TagEvent) bool {
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
	tags, ok := stream.Status.Tags[tag]
	if !ok || len(tags.Items) == 0 {
		return true
	}
	previous := &tags.Items[0]
	sameRef := previous.DockerImageReference == next.DockerImageReference
	sameImage := previous.Image == next.Image
	return !(sameRef && sameImage)
}
func DifferentTagGeneration(stream *ImageStream, tag string) bool {
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
	specTag, ok := stream.Spec.Tags[tag]
	if !ok || specTag.Generation == nil {
		return true
	}
	statusTag, ok := stream.Status.Tags[tag]
	if !ok || len(statusTag.Items) == 0 {
		return true
	}
	return *specTag.Generation > statusTag.Items[0].Generation
}
func AddTagEventToImageStream(stream *ImageStream, tag string, next TagEvent) bool {
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
	if stream.Status.Tags == nil {
		stream.Status.Tags = make(map[string]TagEventList)
	}
	tags, ok := stream.Status.Tags[tag]
	if !ok || len(tags.Items) == 0 {
		stream.Status.Tags[tag] = TagEventList{Items: []TagEvent{next}}
		return true
	}
	previous := &tags.Items[0]
	sameRef := previous.DockerImageReference == next.DockerImageReference
	sameImage := previous.Image == next.Image
	sameGen := previous.Generation == next.Generation
	switch {
	case sameRef && sameImage && sameGen:
		return false
	case sameImage && sameRef:
	case sameRef:
		previous.Image = next.Image
	case sameImage:
		previous.DockerImageReference = next.DockerImageReference
	default:
		tags.Conditions = nil
		tags.Items = append([]TagEvent{next}, tags.Items...)
		stream.Status.Tags[tag] = tags
		return true
	}
	previous.Generation = next.Generation
	tags.Conditions = nil
	stream.Status.Tags[tag] = tags
	return true
}
func UpdateChangedTrackingTags(new, old *ImageStream) int {
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
	changes := 0
	for newTag, newImages := range new.Status.Tags {
		if len(newImages.Items) == 0 {
			continue
		}
		if old != nil {
			oldImages := old.Status.Tags[newTag]
			changed, deleted := tagsChanged(newImages.Items, oldImages.Items)
			if !changed || deleted {
				continue
			}
		}
		changes += UpdateTrackingTags(new, newTag, newImages.Items[0])
	}
	return changes
}
func tagsChanged(new, old []TagEvent) (changed bool, deleted bool) {
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
	case len(old) == 0 && len(new) == 0:
		return false, false
	case len(new) == 0:
		return true, true
	case len(old) == 0:
		return true, false
	default:
		return new[0] != old[0], false
	}
}
func UpdateTrackingTags(stream *ImageStream, updatedTag string, updatedImage TagEvent) int {
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
	updated := 0
	klog.V(5).Infof("UpdateTrackingTags: stream=%s/%s, updatedTag=%s, updatedImage.dockerImageReference=%s, updatedImage.image=%s", stream.Namespace, stream.Name, updatedTag, updatedImage.DockerImageReference, updatedImage.Image)
	for specTag, tagRef := range stream.Spec.Tags {
		klog.V(5).Infof("Examining spec tag %q, tagRef=%#v", specTag, tagRef)
		if tagRef.From == nil {
			klog.V(5).Infof("tagRef.From is nil, skipping")
			continue
		}
		if tagRef.From.Kind != "ImageStreamTag" {
			klog.V(5).Infof("tagRef.Kind %q isn't ImageStreamTag, skipping", tagRef.From.Kind)
			continue
		}
		tagRefNamespace := tagRef.From.Namespace
		if len(tagRefNamespace) == 0 {
			tagRefNamespace = stream.Namespace
		}
		if tagRefNamespace != stream.Namespace {
			klog.V(5).Infof("tagRefNamespace %q doesn't match stream namespace %q - skipping", tagRefNamespace, stream.Namespace)
			continue
		}
		tag := ""
		tagRefName := ""
		if strings.Contains(tagRef.From.Name, ":") {
			ok := true
			tagRefName, tag, ok = SplitImageStreamTag(tagRef.From.Name)
			if !ok {
				klog.V(5).Infof("tagRefName %q contains invalid reference - skipping", tagRef.From.Name)
				continue
			}
		} else {
			tagRefName = stream.Name
			tag = tagRef.From.Name
		}
		klog.V(5).Infof("tagRefName=%q, tag=%q", tagRefName, tag)
		if tagRefName != stream.Name {
			klog.V(5).Infof("tagRefName %q doesn't match stream name %q - skipping", tagRefName, stream.Name)
			continue
		}
		if tag != updatedTag {
			klog.V(5).Infof("tag %q doesn't match updated tag %q - skipping", tag, updatedTag)
			continue
		}
		if AddTagEventToImageStream(stream, specTag, updatedImage) {
			klog.V(5).Infof("stream updated")
			updated++
		}
	}
	return updated
}
func DigestOrImageMatch(image, imageID string) bool {
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
	if d, err := digest.ParseDigest(image); err == nil {
		return strings.HasPrefix(d.Hex(), imageID) || strings.HasPrefix(image, imageID)
	}
	return strings.HasPrefix(image, imageID)
}
func ResolveImageID(stream *ImageStream, imageID string) (*TagEvent, error) {
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
	var event *TagEvent
	set := sets.NewString()
	for _, history := range stream.Status.Tags {
		for i := range history.Items {
			tagging := &history.Items[i]
			if DigestOrImageMatch(tagging.Image, imageID) {
				event = tagging
				set.Insert(tagging.Image)
			}
		}
	}
	switch len(set) {
	case 1:
		return &TagEvent{Created: metav1.Now(), DockerImageReference: event.DockerImageReference, Image: event.Image}, nil
	case 0:
		return nil, kerrors.NewNotFound(image.Resource("imagestreamimage"), imageID)
	default:
		return nil, kerrors.NewConflict(image.Resource("imagestreamimage"), imageID, fmt.Errorf("multiple images match the prefix %q: %s", imageID, strings.Join(set.List(), ", ")))
	}
}
func MostAccuratePullSpec(pullSpec string, id, tag string) (string, bool) {
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
	ref, err := reference.Parse(pullSpec)
	if err != nil {
		return pullSpec, false
	}
	if len(id) > 0 {
		ref.ID = id
	}
	if len(tag) > 0 {
		ref.Tag = tag
	}
	return ref.MostSpecific().Exact(), true
}
func ShortDockerImageID(image *DockerImage, length int) string {
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
	id := image.ID
	if s, err := digest.ParseDigest(id); err == nil {
		id = s.Hex()
	}
	if len(id) > length {
		id = id[:length]
	}
	return id
}
func HasTagCondition(stream *ImageStream, tag string, condition TagEventCondition) bool {
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
	for _, existing := range stream.Status.Tags[tag].Conditions {
		if condition.Type == existing.Type && condition.Status == existing.Status && condition.Reason == existing.Reason {
			return true
		}
	}
	return false
}
func SetTagConditions(stream *ImageStream, tag string, conditions ...TagEventCondition) {
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
	tagEvents := stream.Status.Tags[tag]
	tagEvents.Conditions = conditions
	if stream.Status.Tags == nil {
		stream.Status.Tags = make(map[string]TagEventList)
	}
	stream.Status.Tags[tag] = tagEvents
}
func LatestObservedTagGeneration(stream *ImageStream, tag string) int64 {
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
	tagEvents, ok := stream.Status.Tags[tag]
	if !ok {
		return 0
	}
	lastGen := int64(0)
	if items := tagEvents.Items; len(items) > 0 {
		tagEvent := items[0]
		if tagEvent.Generation > lastGen {
			lastGen = tagEvent.Generation
		}
	}
	for _, condition := range tagEvents.Conditions {
		if condition.Type != ImportSuccess {
			continue
		}
		if condition.Generation > lastGen {
			lastGen = condition.Generation
		}
		break
	}
	return lastGen
}

var (
	reMinorSemantic		= regexp.MustCompile(`^[\d]+\.[\d]+$`)
	reMinorWithPatch	= regexp.MustCompile(`^([\d]+\.[\d]+)-\w+$`)
)

type tagPriority int

const (
	tagPriorityLatest	tagPriority	= iota
	tagPriorityMinor
	tagPriorityFull
	tagPriorityOther
)

type prioritizedTag struct {
	tag		string
	priority	tagPriority
	semver		semver.Version
	prefix		string
}

func prioritizeTag(tag string) prioritizedTag {
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
	if tag == DefaultImageTag {
		return prioritizedTag{tag: tag, priority: tagPriorityLatest}
	}
	short := tag
	prefix := ""
	if strings.HasPrefix(tag, "v") {
		prefix = "v"
		short = tag[1:]
	}
	if v, err := semver.Parse(short); err == nil {
		return prioritizedTag{tag: tag, priority: tagPriorityFull, semver: v, prefix: prefix}
	}
	if reMinorSemantic.MatchString(short) {
		if v, err := semver.Parse(short + ".0"); err == nil {
			return prioritizedTag{tag: tag, priority: tagPriorityMinor, semver: v, prefix: prefix}
		}
	}
	if match := reMinorWithPatch.FindStringSubmatch(short); match != nil {
		if v, err := semver.Parse(strings.Replace(short, match[1], match[1]+".0", 1)); err == nil {
			return prioritizedTag{tag: tag, priority: tagPriorityMinor, semver: v, prefix: prefix}
		}
	}
	return prioritizedTag{tag: tag, priority: tagPriorityOther, prefix: prefix}
}

type prioritizedTags []prioritizedTag

func (t prioritizedTags) Len() int {
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
	return len(t)
}
func (t prioritizedTags) Swap(i, j int) {
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
	t[i], t[j] = t[j], t[i]
}
func (t prioritizedTags) Less(i, j int) bool {
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
	if t[i].priority != t[j].priority {
		return t[i].priority < t[j].priority
	}
	if t[i].priority == tagPriorityOther {
		return t[i].tag < t[j].tag
	}
	cmp := t[i].semver.Compare(t[j].semver)
	if cmp > 0 {
		return true
	}
	return cmp == 0 && t[i].prefix < t[j].prefix
}
func PrioritizeTags(tags []string) {
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
	ptags := make(prioritizedTags, len(tags))
	for i, tag := range tags {
		ptags[i] = prioritizeTag(tag)
	}
	sort.Sort(ptags)
	for i, pt := range ptags {
		tags[i] = pt.tag
	}
}
func LabelForStream(stream *ImageStream) string {
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
	return fmt.Sprintf("%s/%s", stream.Namespace, stream.Name)
}
func JoinImageSignatureName(imageName, signatureName string) (string, error) {
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
	if len(imageName) == 0 {
		return "", fmt.Errorf("imageName may not be empty")
	}
	if len(signatureName) == 0 {
		return "", fmt.Errorf("signatureName may not be empty")
	}
	if strings.Count(imageName, "@") > 0 || strings.Count(signatureName, "@") > 0 {
		return "", fmt.Errorf("neither imageName nor signatureName can contain '@'")
	}
	return fmt.Sprintf("%s@%s", imageName, signatureName), nil
}
func SplitImageSignatureName(imageSignatureName string) (imageName, signatureName string, err error) {
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
	segments := strings.Split(imageSignatureName, "@")
	switch len(segments) {
	case 2:
		signatureName = segments[1]
		imageName = segments[0]
		if len(imageName) == 0 || len(signatureName) == 0 {
			err = fmt.Errorf("image signature name %q must have an image name and signature name", imageSignatureName)
		}
	default:
		err = fmt.Errorf("expected exactly one @ in the image signature name %q", imageSignatureName)
	}
	return
}
func IndexOfImageSignatureByName(signatures []ImageSignature, name string) int {
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
	for i := range signatures {
		if signatures[i].Name == name {
			return i
		}
	}
	return -1
}
func IndexOfImageSignature(signatures []ImageSignature, sType string, sContent []byte) int {
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
	for i := range signatures {
		if signatures[i].Type == sType && bytes.Equal(signatures[i].Content, sContent) {
			return i
		}
	}
	return -1
}
func (tagref TagReference) HasAnnotationTag(searchTag string) bool {
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
	for _, tag := range strings.Split(tagref.Annotations["tags"], ",") {
		if tag == searchTag {
			return true
		}
	}
	return false
}
func ValidateRegistryURL(registryURL string) error {
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
		u	*url.URL
		err	error
		parts	= strings.SplitN(registryURL, "://", 2)
	)
	switch len(parts) {
	case 2:
		u, err = url.Parse(registryURL)
		if err != nil {
			return err
		}
		switch u.Scheme {
		case "http", "https":
		default:
			return fmt.Errorf("unsupported scheme: %s", u.Scheme)
		}
	case 1:
		u, err = url.Parse("https://" + registryURL)
		if err != nil {
			return err
		}
	}
	if len(u.Path) > 0 && u.Path != "/" {
		return errNoRegistryURLPathAllowed
	}
	if len(u.RawQuery) > 0 {
		return errNoRegistryURLQueryAllowed
	}
	if len(u.Host) == 0 {
		return errRegistryURLHostEmpty
	}
	return nil
}

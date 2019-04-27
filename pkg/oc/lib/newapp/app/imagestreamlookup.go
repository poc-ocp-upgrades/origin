package app

import (
	"fmt"
	"strings"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	dockerv10 "github.com/openshift/api/image/docker10"
	imagev1 "github.com/openshift/api/image/v1"
	imagev1typedclient "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imageutil "github.com/openshift/origin/pkg/image/util"
)

type ImageStreamSearcher struct {
	Client			imagev1typedclient.ImageV1Interface
	Namespaces		[]string
	AllowMissingTags	bool
}

func (r ImageStreamSearcher) Type() string {
	_logClusterCodePath()
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
	return "images in image streams"
}
func (r ImageStreamSearcher) Search(precise bool, terms ...string) (ComponentMatches, []error) {
	_logClusterCodePath()
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
	componentMatches := ComponentMatches{}
	var errs []error
	for _, term := range terms {
		var (
			ref	imageapi.DockerImageReference
			err	error
		)
		switch term {
		case "__imagestream_fail":
			errs = append(errs, fmt.Errorf("unable to find the specified image: %s", term))
			continue
		case "*":
			ref = imageapi.DockerImageReference{Name: term}
		default:
			ref, err = imageapi.ParseDockerImageReference(term)
			if err != nil || len(ref.Registry) != 0 {
				klog.V(2).Infof("image streams must be of the form [<namespace>/]<name>[:<tag>|@<digest>], term %q did not qualify", term)
				continue
			}
		}
		namespaces := r.Namespaces
		if len(ref.Namespace) != 0 {
			namespaces = []string{ref.Namespace}
		}
		followTag := false
		searchTag := ref.Tag
		if len(searchTag) == 0 {
			searchTag = imageapi.DefaultImageTag
			followTag = true
		}
		for _, namespace := range namespaces {
			klog.V(4).Infof("checking ImageStreams %s/%s with ref %q", namespace, ref.Name, searchTag)
			exact := false
			streams, err := r.Client.ImageStreams(namespace).List(metav1.ListOptions{})
			if err != nil {
				if errors.IsNotFound(err) || errors.IsForbidden(err) {
					continue
				}
				errs = append(errs, err)
				continue
			}
			original := ref
			ref.Namespace = namespace
			for i := range streams.Items {
				stream := &streams.Items[i]
				score, scored := imageStreamScorer(*stream, ref.Name)
				if !scored {
					klog.V(2).Infof("unscored %s: %v", stream.Name, score)
					continue
				}
				var meta map[string]string
				if stream.Generation > 0 {
					meta = map[string]string{"direct-tag": "1"}
				}
				imageref := original
				imageref.Name = stream.Name
				imageref.Registry = ""
				matchName := fmt.Sprintf("%s/%s", stream.Namespace, stream.Name)
				addMatch := func(tag string, matchScore float32, image *dockerv10.DockerImage, notFound bool) {
					name := matchName
					var description, argument string
					if len(tag) > 0 {
						name = fmt.Sprintf("%s:%s", name, tag)
						argument = fmt.Sprintf("--image-stream=%q", name)
						description = fmt.Sprintf("Image stream %q (tag %q) in project %q", stream.Name, tag, stream.Namespace)
					} else {
						argument = fmt.Sprintf("--image-stream=%q --allow-missing-imagestream-tags", name)
						description = fmt.Sprintf("Image stream %q in project %q", stream.Name, stream.Namespace)
					}
					match := &ComponentMatch{Value: term, Argument: argument, Name: name, Description: description, Score: matchScore, ImageStream: stream, DockerImage: image, ImageTag: tag, Meta: meta, NoTagsFound: notFound}
					klog.V(2).Infof("Adding %s as component match for %q with score %v", match.Description, term, matchScore)
					componentMatches = append(componentMatches, match)
				}
				finalTag := searchTag
				var specTag *imagev1.TagReference
				if t, hasTag := imageutil.SpecHasTag(stream, searchTag); hasTag {
					specTag = &t
				}
				if specTag != nil && followTag && !imageutil.HasAnnotationTag(specTag, imageapi.TagReferenceAnnotationTagHidden) {
					if specTag.From != nil && specTag.From.Kind == "ImageStreamTag" && !strings.Contains(specTag.From.Name, ":") {
						if t, hasTag := imageutil.SpecHasTag(stream, specTag.From.Name); hasTag && !imageutil.HasAnnotationTag(&t, imageapi.TagReferenceAnnotationTagHidden) {
							if imageutil.LatestTaggedImage(stream, specTag.From.Name) != nil {
								finalTag = specTag.From.Name
							}
						}
					}
				}
				latest := imageutil.LatestTaggedImage(stream, finalTag)
				if (specTag != nil && followTag && imageutil.HasAnnotationTag(specTag, imageapi.TagReferenceAnnotationTagHidden)) || latest == nil || len(latest.Image) == 0 {
					klog.V(2).Infof("no image recorded for %s/%s:%s", stream.Namespace, stream.Name, finalTag)
					if r.AllowMissingTags {
						addMatch(finalTag, score, nil, false)
						continue
					}
					foundOtherTags := false
					for _, tag := range stream.Status.Tags {
						latest := imageutil.LatestTaggedImage(stream, tag.Tag)
						if latest == nil || len(latest.Image) == 0 {
							continue
						}
						foundOtherTags = true
						t, hasTag := imageutil.SpecHasTag(stream, tag.Tag)
						if hasTag && imageutil.HasAnnotationTag(&t, imageapi.TagReferenceAnnotationTagHidden) {
							continue
						}
						tagScore := score + 0.5
						addMatch(tag.Tag, tagScore, nil, false)
					}
					if !foundOtherTags {
						addMatch("", 0.5+score, nil, true)
					}
					continue
				}
				imageStreamImage, err := r.Client.ImageStreamImages(namespace).Get(imageapi.JoinImageStreamImage(stream.Name, latest.Image), metav1.GetOptions{})
				if err != nil {
					if errors.IsNotFound(err) {
						klog.V(2).Infof("tag %q is set, but image %q has been removed", finalTag, latest.Image)
						continue
					}
					errs = append(errs, err)
					continue
				}
				if err := imageutil.ImageWithMetadata(&imageStreamImage.Image); err != nil {
					errs = append(errs, err)
					continue
				}
				dockerImage, ok := imageStreamImage.Image.DockerImageMetadata.Object.(*dockerv10.DockerImage)
				if !ok {
					continue
				}
				addMatch(finalTag, score, dockerImage, false)
				if score == 0.0 {
					exact = true
				}
			}
			if exact && precise {
				break
			}
		}
	}
	return componentMatches, errs
}
func InputImageFromMatch(match *ComponentMatch) (*ImageRef, error) {
	_logClusterCodePath()
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
	g := NewImageRefGenerator()
	switch {
	case match.ImageStream != nil:
		input, err := g.FromStream(match.ImageStream, match.ImageTag)
		if err != nil {
			return nil, err
		}
		if match.Meta["direct-tag"] == "1" {
			input.TagDirectly = true
		}
		input.AsImageStream = true
		input.Info = match.DockerImage
		return input, nil
	case match.DockerImage != nil:
		input, err := g.FromName(match.Value)
		if err != nil {
			return nil, err
		}
		if match.Meta["direct-tag"] == "1" {
			input.TagDirectly = true
			input.AsResolvedImage = true
		}
		input.AsImageStream = !match.LocalOnly
		input.Info = match.DockerImage
		input.Insecure = match.Insecure
		return input, nil
	default:
		input, err := g.FromName(match.Value)
		if err != nil {
			return nil, err
		}
		return input, nil
	}
}

type ImageStreamByAnnotationSearcher struct {
	Client			imagev1typedclient.ImageStreamsGetter
	ImageStreamImages	imagev1typedclient.ImageStreamImagesGetter
	Namespaces		[]string
	imageStreams		map[string]*imagev1.ImageStreamList
}

const supportsAnnotationKey = "supports"

func NewImageStreamByAnnotationSearcher(streamClient imagev1typedclient.ImageStreamsGetter, imageClient imagev1typedclient.ImageStreamImagesGetter, namespaces []string) Searcher {
	_logClusterCodePath()
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
	return &ImageStreamByAnnotationSearcher{Client: streamClient, ImageStreamImages: imageClient, Namespaces: namespaces, imageStreams: make(map[string]*imagev1.ImageStreamList)}
}
func (r *ImageStreamByAnnotationSearcher) getImageStreams(namespace string) ([]imagev1.ImageStream, error) {
	_logClusterCodePath()
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
	imageStreamList, ok := r.imageStreams[namespace]
	if !ok {
		var err error
		imageStreamList, err = r.Client.ImageStreams(namespace).List(metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		r.imageStreams[namespace] = imageStreamList
	}
	return imageStreamList.Items, nil
}
func matchSupportsAnnotation(value, annotation string) (float32, bool) {
	_logClusterCodePath()
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
	valueBase := strings.Split(value, ":")[0]
	parts := strings.Split(annotation, ",")
	for _, p := range parts {
		if value == p {
			return 0.0, true
		}
	}
	for _, p := range parts {
		partBase := strings.Split(p, ":")[0]
		if valueBase == partBase {
			return 0.5, true
		}
	}
	return 0, false
}
func (r *ImageStreamByAnnotationSearcher) annotationMatches(stream *imagev1.ImageStream, value string) []*ComponentMatch {
	_logClusterCodePath()
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
	if stream.Spec.Tags == nil {
		klog.Infof("No tags found on image, returning nil")
		return nil
	}
	matches := []*ComponentMatch{}
	for _, tagref := range stream.Spec.Tags {
		if tagref.Annotations == nil {
			continue
		}
		supports, ok := tagref.Annotations[supportsAnnotationKey]
		if !ok {
			continue
		}
		score, ok := matchSupportsAnnotation(value, supports)
		if !ok {
			continue
		}
		latest := imageutil.LatestTaggedImage(stream, tagref.Name)
		if latest == nil {
			continue
		}
		imageStream, err := r.ImageStreamImages.ImageStreamImages(stream.Namespace).Get(imageapi.JoinImageStreamImage(stream.Name, latest.Image), metav1.GetOptions{})
		if err != nil {
			klog.V(2).Infof("Could not retrieve image stream image for stream %q, tag %q: %v", stream.Name, tagref.Name, err)
			continue
		}
		if imageStream == nil {
			continue
		}
		var meta map[string]string
		if imageStream.Generation > 0 {
			meta = map[string]string{"direct-tag": "1"}
		}
		imageData := imageStream.Image
		if err := imageutil.ImageWithMetadata(&imageData); err != nil {
			klog.V(5).Infof("error obtaining docker image metadata: %v", err)
			return nil
		}
		dockerImage, ok := imageData.DockerImageMetadata.Object.(*dockerv10.DockerImage)
		if !ok {
			continue
		}
		matchName := fmt.Sprintf("%s/%s", stream.Namespace, stream.Name)
		description := fmt.Sprintf("Image stream %q in project %q", stream.Name, stream.Namespace)
		if len(tagref.Name) > 0 {
			matchName = fmt.Sprintf("%s:%s", matchName, tagref.Name)
			description = fmt.Sprintf("Image stream %q (tag %q) in project %q", stream.Name, tagref.Name, stream.Namespace)
		}
		klog.V(5).Infof("ImageStreamAnnotationSearcher match found: %s for %s with score %f", matchName, value, score)
		match := &ComponentMatch{Value: value, Name: fmt.Sprintf("%s", matchName), Argument: fmt.Sprintf("--image-stream=%q", matchName), Description: description, Score: score, ImageStream: stream, DockerImage: dockerImage, ImageTag: tagref.Name, Meta: meta}
		matches = append(matches, match)
	}
	return matches
}
func (r *ImageStreamByAnnotationSearcher) Type() string {
	_logClusterCodePath()
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
	return "image stream images with a 'supports' annotation"
}
func (r *ImageStreamByAnnotationSearcher) Search(precise bool, terms ...string) (ComponentMatches, []error) {
	_logClusterCodePath()
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
	matches := ComponentMatches{}
	var errs []error
	for _, namespace := range r.Namespaces {
		streams, err := r.getImageStreams(namespace)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		for i := range streams {
			for _, term := range terms {
				if term == "__imagestreamannotation_fail" {
					errs = append(errs, fmt.Errorf("unable to find the specified image: %s", term))
					continue
				}
				klog.V(5).Infof("Checking imagestream %s/%s for supports annotation %q", namespace, streams[i].Name, term)
				matches = append(matches, r.annotationMatches(&streams[i], term)...)
			}
		}
		if precise {
			for _, m := range matches {
				if m.Score == 0.0 {
					return matches, errs
				}
			}
		}
	}
	return matches, errs
}

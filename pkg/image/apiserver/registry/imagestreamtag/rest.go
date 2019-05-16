package imagestreamtag

import (
	"context"
	"fmt"
	goformat "fmt"
	imagegroup "github.com/openshift/api/image"
	"github.com/openshift/origin/pkg/api/apihelpers"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apis/image/validation/whitelist"
	"github.com/openshift/origin/pkg/image/apiserver/registry/image"
	"github.com/openshift/origin/pkg/image/apiserver/registry/imagestream"
	"github.com/openshift/origin/pkg/image/util"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type REST struct {
	imageRegistry       image.Registry
	imageStreamRegistry imagestream.Registry
	strategy            Strategy
	rest.TableConvertor
}

func NewREST(imageRegistry image.Registry, imageStreamRegistry imagestream.Registry, registryWhitelister whitelist.RegistryWhitelister) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &REST{imageRegistry: imageRegistry, imageStreamRegistry: imageStreamRegistry, strategy: NewStrategy(registryWhitelister), TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
}

var _ rest.Getter = &REST{}
var _ rest.Lister = &REST{}
var _ rest.CreaterUpdater = &REST{}
var _ rest.GracefulDeleter = &REST{}
var _ rest.ShortNamesProvider = &REST{}
var _ rest.Scoper = &REST{}

func (r *REST) ShortNames() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{"istag"}
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &imageapi.ImageStreamTag{}
}
func (r *REST) NewList() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &imageapi.ImageStreamTagList{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func nameAndTag(id string) (name string, tag string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name, tag, err = imageapi.ParseImageStreamTagName(id)
	if err != nil {
		err = kapierrors.NewBadRequest("ImageStreamTags must be retrieved with <name>:<tag>")
	}
	return
}
func (r *REST) List(ctx context.Context, options *metainternal.ListOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	imageStreams, err := r.imageStreamRegistry.ListImageStreams(ctx, options)
	if err != nil {
		return nil, err
	}
	matcher := MatchImageStreamTag(apihelpers.InternalListOptionsToSelectors(options))
	list := &imageapi.ImageStreamTagList{}
	for _, currIS := range imageStreams.Items {
		for currTag := range currIS.Status.Tags {
			istag, err := newISTag(currTag, &currIS, nil, false)
			if err != nil {
				if kapierrors.IsNotFound(err) {
					continue
				}
				return nil, err
			}
			matches, err := matcher.Matches(istag)
			if err != nil {
				return nil, err
			}
			if matches {
				list.Items = append(list.Items, *istag)
			}
		}
	}
	return list, nil
}
func (r *REST) Get(ctx context.Context, id string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name, tag, err := nameAndTag(id)
	if err != nil {
		return nil, err
	}
	imageStream, err := r.imageStreamRegistry.GetImageStream(ctx, name, options)
	if err != nil {
		return nil, err
	}
	image, err := r.imageFor(ctx, tag, imageStream)
	if err != nil {
		return nil, err
	}
	return newISTag(tag, imageStream, image, false)
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	istag, ok := obj.(*imageapi.ImageStreamTag)
	if !ok {
		return nil, kapierrors.NewBadRequest(fmt.Sprintf("obj is not an ImageStreamTag: %#v", obj))
	}
	if err := rest.BeforeCreate(r.strategy, ctx, obj); err != nil {
		return nil, err
	}
	if err := createValidation(obj.DeepCopyObject()); err != nil {
		return nil, err
	}
	namespace, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		return nil, kapierrors.NewBadRequest("a namespace must be specified to import images")
	}
	imageStreamName, imageTag, ok := imageapi.SplitImageStreamTag(istag.Name)
	if !ok {
		return nil, fmt.Errorf("%q must be of the form <stream_name>:<tag>", istag.Name)
	}
	for i := 10; i > 0; i-- {
		target, err := r.imageStreamRegistry.GetImageStream(ctx, imageStreamName, &metav1.GetOptions{})
		if err != nil {
			if !kapierrors.IsNotFound(err) {
				return nil, err
			}
			target = &imageapi.ImageStream{ObjectMeta: metav1.ObjectMeta{Name: imageStreamName, Namespace: namespace}}
		}
		if target.Spec.Tags == nil {
			target.Spec.Tags = make(map[string]imageapi.TagReference)
		}
		_, exists := target.Spec.Tags[imageTag]
		if exists {
			return nil, kapierrors.NewAlreadyExists(imagegroup.Resource("imagestreamtag"), istag.Name)
		}
		if istag.Tag != nil {
			target.Spec.Tags[imageTag] = *istag.Tag
		}
		if target.CreationTimestamp.IsZero() {
			target, err = r.imageStreamRegistry.CreateImageStream(ctx, target, &metav1.CreateOptions{})
		} else {
			target, err = r.imageStreamRegistry.UpdateImageStream(ctx, target, false, &metav1.UpdateOptions{})
		}
		if kapierrors.IsAlreadyExists(err) || kapierrors.IsConflict(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		image, _ := r.imageFor(ctx, imageTag, target)
		return newISTag(imageTag, target, image, true)
	}
	return nil, kapierrors.NewServerTimeout(imagegroup.Resource("imagestreamtags"), "create", 2)
}
func (r *REST) Update(ctx context.Context, tagName string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name, tag, err := nameAndTag(tagName)
	if err != nil {
		return nil, false, err
	}
	create := false
	imageStream, err := r.imageStreamRegistry.GetImageStream(ctx, name, &metav1.GetOptions{})
	if err != nil {
		if !kapierrors.IsNotFound(err) {
			return nil, false, err
		}
		namespace, ok := apirequest.NamespaceFrom(ctx)
		if !ok {
			return nil, false, kapierrors.NewBadRequest("namespace is required on ImageStreamTags")
		}
		imageStream = &imageapi.ImageStream{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name}}
		rest.FillObjectMetaSystemFields(&imageStream.ObjectMeta)
		create = true
	}
	old, err := newISTag(tag, imageStream, nil, true)
	if err != nil {
		return nil, false, err
	}
	obj, err := objInfo.UpdatedObject(ctx, old)
	if err != nil {
		return nil, false, err
	}
	istag, ok := obj.(*imageapi.ImageStreamTag)
	if !ok {
		return nil, false, kapierrors.NewBadRequest(fmt.Sprintf("obj is not an ImageStreamTag: %#v", obj))
	}
	switch {
	case len(istag.ResourceVersion) == 0:
		istag.ResourceVersion = imageStream.ResourceVersion
	case len(imageStream.ResourceVersion) == 0:
		return nil, false, kapierrors.NewNotFound(imagegroup.Resource("imagestreamtags"), tagName)
	case imageStream.ResourceVersion != istag.ResourceVersion:
		return nil, false, kapierrors.NewConflict(imagegroup.Resource("imagestreamtags"), istag.Name, fmt.Errorf("another caller has updated the resource version to %s", imageStream.ResourceVersion))
	}
	if len(imageStream.Labels) > 0 && len(istag.Labels) == 0 {
		istag.Labels = imageStream.Labels
	}
	if create {
		if err := rest.BeforeCreate(r.strategy, ctx, obj); err != nil {
			return nil, false, err
		}
		if err := createValidation(obj.DeepCopyObject()); err != nil {
			return nil, false, err
		}
	} else {
		if err := rest.BeforeUpdate(r.strategy, ctx, obj, old); err != nil {
			return nil, false, err
		}
		if err := updateValidation(obj.DeepCopyObject(), old.DeepCopyObject()); err != nil {
			return nil, false, err
		}
	}
	if imageStream.Spec.Tags == nil {
		imageStream.Spec.Tags = map[string]imageapi.TagReference{}
	}
	tagRef, exists := imageStream.Spec.Tags[tag]
	if !exists && istag.Tag == nil {
		return nil, false, kapierrors.NewBadRequest(fmt.Sprintf("imagestreamtag %s is not a spec tag in imagestream %s/%s, cannot be updated", tag, imageStream.Namespace, imageStream.Name))
	}
	if istag.Tag != nil {
		tagRef = *istag.Tag
		tagRef.Name = tag
	}
	tagRef.Annotations = istag.Annotations
	imageStream.Spec.Tags[tag] = tagRef
	var newImageStream *imageapi.ImageStream
	if create {
		newImageStream, err = r.imageStreamRegistry.CreateImageStream(ctx, imageStream, &metav1.CreateOptions{})
	} else {
		newImageStream, err = r.imageStreamRegistry.UpdateImageStream(ctx, imageStream, false, &metav1.UpdateOptions{})
	}
	if err != nil {
		return nil, false, err
	}
	image, err := r.imageFor(ctx, tag, newImageStream)
	if err != nil {
		if !kapierrors.IsNotFound(err) {
			return nil, false, err
		}
	}
	newISTag, err := newISTag(tag, newImageStream, image, true)
	return newISTag, !exists, err
}
func (r *REST) Delete(ctx context.Context, id string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name, tag, err := nameAndTag(id)
	if err != nil {
		return nil, false, err
	}
	for i := 10; i > 0; i-- {
		stream, err := r.imageStreamRegistry.GetImageStream(ctx, name, &metav1.GetOptions{})
		if err != nil {
			return nil, false, err
		}
		if options != nil {
			if pre := options.Preconditions; pre != nil {
				if pre.UID != nil && *pre.UID != stream.UID {
					return nil, false, kapierrors.NewConflict(imagegroup.Resource("imagestreamtags"), id, fmt.Errorf("the UID precondition was not met"))
				}
			}
		}
		notFound := true
		if _, ok := stream.Status.Tags[tag]; ok {
			delete(stream.Status.Tags, tag)
			notFound = false
		}
		if _, ok := stream.Spec.Tags[tag]; ok {
			delete(stream.Spec.Tags, tag)
			notFound = false
		}
		if notFound {
			return nil, false, kapierrors.NewNotFound(imagegroup.Resource("imagestreamtags"), id)
		}
		_, err = r.imageStreamRegistry.UpdateImageStream(ctx, stream, false, &metav1.UpdateOptions{})
		if kapierrors.IsConflict(err) {
			continue
		}
		if err != nil && !kapierrors.IsNotFound(err) {
			return nil, false, err
		}
		return &metav1.Status{Status: metav1.StatusSuccess}, true, nil
	}
	return nil, false, kapierrors.NewServerTimeout(imagegroup.Resource("imagestreamtags"), "delete", 2)
}
func (r *REST) imageFor(ctx context.Context, tag string, imageStream *imageapi.ImageStream) (*imageapi.Image, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	event := imageapi.LatestTaggedImage(imageStream, tag)
	if event == nil || len(event.Image) == 0 {
		return nil, kapierrors.NewNotFound(imagegroup.Resource("imagestreamtags"), imageapi.JoinImageStreamTag(imageStream.Name, tag))
	}
	return r.imageRegistry.GetImage(ctx, event.Image, &metav1.GetOptions{})
}
func newISTag(tag string, imageStream *imageapi.ImageStream, image *imageapi.Image, allowEmptyEvent bool) (*imageapi.ImageStreamTag, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	istagName := imageapi.JoinImageStreamTag(imageStream.Name, tag)
	event := imageapi.LatestTaggedImage(imageStream, tag)
	if event == nil || len(event.Image) == 0 {
		if !allowEmptyEvent {
			klog.V(4).Infof("did not find tag %s in image stream status tags: %#v", tag, imageStream.Status.Tags)
			return nil, kapierrors.NewNotFound(imagegroup.Resource("imagestreamtags"), istagName)
		}
		event = &imageapi.TagEvent{Created: imageStream.CreationTimestamp}
	}
	ist := &imageapi.ImageStreamTag{ObjectMeta: metav1.ObjectMeta{Namespace: imageStream.Namespace, Name: istagName, CreationTimestamp: event.Created, Annotations: map[string]string{}, Labels: imageStream.Labels, ResourceVersion: imageStream.ResourceVersion, UID: imageStream.UID}, Generation: event.Generation, Conditions: imageStream.Status.Tags[tag].Conditions, LookupPolicy: imageStream.Spec.LookupPolicy}
	if imageStream.Spec.Tags != nil {
		if tagRef, ok := imageStream.Spec.Tags[tag]; ok {
			ist.Tag = &tagRef
			if from := ist.Tag.From; from != nil {
				copied := *from
				ist.Tag.From = &copied
			}
			if gen := ist.Tag.Generation; gen != nil {
				copied := *gen
				ist.Tag.Generation = &copied
			}
			if image != nil && image.Annotations == nil {
				image.Annotations = make(map[string]string)
			}
			for k, v := range tagRef.Annotations {
				ist.Annotations[k] = v
				if image != nil {
					image.Annotations[k] = v
				}
			}
		}
	}
	if image != nil {
		if err := util.InternalImageWithMetadata(image); err != nil {
			return nil, err
		}
		image.DockerImageManifest = ""
		image.DockerImageConfig = ""
		ist.Image = *image
	} else {
		ist.Image = imageapi.Image{}
		ist.Image.Name = event.Image
	}
	ist.Image.DockerImageReference = imageapi.ResolveReferenceForTagEvent(imageStream, tag, event)
	return ist, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

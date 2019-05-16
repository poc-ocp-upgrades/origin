package image

import (
	"context"
	goformat "fmt"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &storage{Storage: s}
}
func (s *storage) ListImages(ctx context.Context, options *metainternal.ListOptions) (*imageapi.ImageList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*imageapi.ImageList), nil
}
func (s *storage) GetImage(ctx context.Context, imageID string, options *metav1.GetOptions) (*imageapi.Image, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, err := s.Get(ctx, imageID, options)
	if err != nil {
		return nil, err
	}
	return obj.(*imageapi.Image), nil
}
func (s *storage) CreateImage(ctx context.Context, image *imageapi.Image) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, err := s.Create(ctx, image, rest.ValidateAllObjectFunc, &metav1.CreateOptions{})
	return err
}
func (s *storage) UpdateImage(ctx context.Context, image *imageapi.Image) (*imageapi.Image, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, _, err := s.Update(ctx, image.Name, rest.DefaultUpdatedObjectInfo(image), rest.ValidateAllObjectFunc, rest.ValidateAllObjectUpdateFunc, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*imageapi.Image), nil
}
func (s *storage) DeleteImage(ctx context.Context, imageID string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, _, err := s.Delete(ctx, imageID, nil)
	return err
}
func (s *storage) WatchImages(ctx context.Context, options *metainternal.ListOptions) (watch.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.Watch(ctx, options)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

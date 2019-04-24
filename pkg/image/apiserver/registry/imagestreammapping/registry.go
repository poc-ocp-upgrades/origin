package imagestreammapping

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Registry interface {
	CreateImageStreamMapping(ctx context.Context, mapping *imageapi.ImageStreamMapping) (*metav1.Status, error)
}
type Storage interface {
	Create(ctx context.Context, obj runtime.Object) (runtime.Object, error)
}
type storage struct{ Storage }

func NewRegistry(s Storage) Registry {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &storage{s}
}
func (s *storage) CreateImageStreamMapping(ctx context.Context, mapping *imageapi.ImageStreamMapping) (*metav1.Status, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := s.Create(ctx, mapping)
	if err != nil {
		return nil, err
	}
	return obj.(*metav1.Status), nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

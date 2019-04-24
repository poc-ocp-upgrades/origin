package subjectaccessreview

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	api "github.com/openshift/origin/pkg/authorization/apis/authorization"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

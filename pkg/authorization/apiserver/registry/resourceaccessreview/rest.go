package resourceaccessreview

import (
	"context"
	"errors"
	"fmt"
	authorization "github.com/openshift/api/authorization"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	authorizationvalidation "github.com/openshift/origin/pkg/authorization/apis/authorization/validation"
	"github.com/openshift/origin/pkg/authorization/apiserver/registry/util"
	authorizationutil "github.com/openshift/origin/pkg/authorization/util"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/authentication/user"
	kauthorizer "k8s.io/apiserver/pkg/authorization/authorizer"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
)

type REST struct {
	authorizer     kauthorizer.Authorizer
	subjectLocator rbac.SubjectLocator
}

var _ rest.Creater = &REST{}
var _ rest.Scoper = &REST{}

func NewREST(authorizer kauthorizer.Authorizer, subjectLocator rbac.SubjectLocator) *REST {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &REST{authorizer, subjectLocator}
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &authorizationapi.ResourceAccessReview{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resourceAccessReview, ok := obj.(*authorizationapi.ResourceAccessReview)
	if !ok {
		return nil, kapierrors.NewBadRequest(fmt.Sprintf("not a resourceAccessReview: %#v", obj))
	}
	if errs := authorizationvalidation.ValidateResourceAccessReview(resourceAccessReview); len(errs) > 0 {
		return nil, kapierrors.NewInvalid(authorization.Kind(resourceAccessReview.Kind), "", errs)
	}
	user, ok := apirequest.UserFrom(ctx)
	if !ok {
		return nil, kapierrors.NewInternalError(errors.New("missing user on request"))
	}
	if namespace := apirequest.NamespaceValue(ctx); len(namespace) > 0 {
		resourceAccessReview.Action.Namespace = namespace
	} else if err := r.isAllowed(user, resourceAccessReview); err != nil {
		return nil, err
	}
	attributes := util.ToDefaultAuthorizationAttributes(nil, resourceAccessReview.Action.Namespace, resourceAccessReview.Action)
	subjects, err := r.subjectLocator.AllowedSubjects(attributes)
	users, groups := authorizationutil.RBACSubjectsToUsersAndGroups(subjects, attributes.GetNamespace())
	response := &authorizationapi.ResourceAccessReviewResponse{Namespace: resourceAccessReview.Action.Namespace, Users: sets.NewString(users...), Groups: sets.NewString(groups...)}
	if err != nil {
		response.EvaluationError = err.Error()
	}
	return response, nil
}
func (r *REST) isAllowed(user user.Info, rar *authorizationapi.ResourceAccessReview) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localRARAttributes := kauthorizer.AttributesRecord{User: user, Verb: "create", Namespace: rar.Action.Namespace, Resource: "localresourceaccessreviews", ResourceRequest: true}
	authorized, reason, err := r.authorizer.Authorize(localRARAttributes)
	if err != nil {
		return kapierrors.NewForbidden(authorization.Resource(localRARAttributes.GetResource()), localRARAttributes.GetName(), err)
	}
	if authorized != kauthorizer.DecisionAllow {
		forbiddenError := kapierrors.NewForbidden(authorization.Resource(localRARAttributes.GetResource()), localRARAttributes.GetName(), errors.New(""))
		forbiddenError.ErrStatus.Message = reason
		return forbiddenError
	}
	return nil
}

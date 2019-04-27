package subjectaccessreview

import (
	"context"
	"errors"
	"fmt"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authentication/user"
	kauthorizer "k8s.io/apiserver/pkg/authorization/authorizer"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	authorization "github.com/openshift/api/authorization"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	authorizationvalidation "github.com/openshift/origin/pkg/authorization/apis/authorization/validation"
	"github.com/openshift/origin/pkg/authorization/apiserver/registry/util"
	"github.com/openshift/origin/pkg/authorization/authorizer"
)

type REST struct{ authorizer kauthorizer.Authorizer }

var _ rest.Creater = &REST{}
var _ rest.Scoper = &REST{}

func NewREST(authorizer kauthorizer.Authorizer) *REST {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &REST{authorizer}
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &authorizationapi.SubjectAccessReview{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	subjectAccessReview, ok := obj.(*authorizationapi.SubjectAccessReview)
	if !ok {
		return nil, kapierrors.NewBadRequest(fmt.Sprintf("not a subjectAccessReview: %#v", obj))
	}
	if errs := authorizationvalidation.ValidateSubjectAccessReview(subjectAccessReview); len(errs) > 0 {
		return nil, kapierrors.NewInvalid(authorization.Kind(subjectAccessReview.Kind), "", errs)
	}
	requestingUser, ok := apirequest.UserFrom(ctx)
	if !ok {
		return nil, kapierrors.NewInternalError(errors.New("missing user on request"))
	}
	if namespace := apirequest.NamespaceValue(ctx); len(namespace) > 0 {
		subjectAccessReview.Action.Namespace = namespace
	} else if err := r.isAllowed(requestingUser, subjectAccessReview); err != nil {
		return nil, err
	}
	var userToCheck *user.DefaultInfo
	if (len(subjectAccessReview.User) == 0) && (len(subjectAccessReview.Groups) == 0) {
		ctxUser, exists := apirequest.UserFrom(ctx)
		if !exists {
			return nil, kapierrors.NewBadRequest("user missing from context")
		}
		newExtra := map[string][]string{}
		for k, v := range ctxUser.GetExtra() {
			if v == nil {
				newExtra[k] = nil
				continue
			}
			newSlice := make([]string, len(v))
			copy(newSlice, v)
			newExtra[k] = newSlice
		}
		userToCheck = &user.DefaultInfo{Name: ctxUser.GetName(), Groups: ctxUser.GetGroups(), UID: ctxUser.GetUID(), Extra: newExtra}
	} else {
		userToCheck = &user.DefaultInfo{Name: subjectAccessReview.User, Groups: subjectAccessReview.Groups.List(), Extra: map[string][]string{}}
	}
	switch {
	case subjectAccessReview.Scopes == nil:
	case len(subjectAccessReview.Scopes) == 0:
		delete(userToCheck.Extra, authorizationapi.ScopesKey)
	case len(subjectAccessReview.Scopes) > 0:
		userToCheck.Extra[authorizationapi.ScopesKey] = subjectAccessReview.Scopes
	}
	attributes := util.ToDefaultAuthorizationAttributes(userToCheck, subjectAccessReview.Action.Namespace, subjectAccessReview.Action)
	authorized, reason, err := r.authorizer.Authorize(attributes)
	response := &authorizationapi.SubjectAccessReviewResponse{Namespace: subjectAccessReview.Action.Namespace, Allowed: authorized == kauthorizer.DecisionAllow, Reason: reason}
	if err != nil {
		response.EvaluationError = err.Error()
	}
	return response, nil
}
func (r *REST) isAllowed(user user.Info, sar *authorizationapi.SubjectAccessReview) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var localSARAttributes kauthorizer.AttributesRecord
	if authorizer.IsPersonalAccessReviewFromSAR(sar) {
		localSARAttributes = kauthorizer.AttributesRecord{User: user, Verb: "create", Namespace: sar.Action.Namespace, APIGroup: "authorization.k8s.io", Resource: "selfsubjectaccessreviews", ResourceRequest: true}
	} else {
		localSARAttributes = kauthorizer.AttributesRecord{User: user, Verb: "create", Namespace: sar.Action.Namespace, Resource: "localsubjectaccessreviews", ResourceRequest: true}
	}
	authorized, reason, err := r.authorizer.Authorize(localSARAttributes)
	if err != nil {
		return kapierrors.NewForbidden(authorization.Resource(localSARAttributes.GetResource()), localSARAttributes.GetName(), err)
	}
	if authorized != kauthorizer.DecisionAllow {
		forbiddenError := kapierrors.NewForbidden(authorization.Resource(localSARAttributes.GetResource()), localSARAttributes.GetName(), errors.New(""))
		forbiddenError.ErrStatus.Message = reason
		return forbiddenError
	}
	return nil
}

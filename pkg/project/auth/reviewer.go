package auth

import (
	authorizationutil "github.com/openshift/origin/pkg/authorization/util"
	kauthorizer "k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
)

type Review interface {
	Users() []string
	Groups() []string
	EvaluationError() string
}
type defaultReview struct {
	users           []string
	groups          []string
	evaluationError string
}

func (r *defaultReview) Users() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.users
}
func (r *defaultReview) Groups() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.groups
}
func (r *defaultReview) EvaluationError() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.evaluationError
}

type Reviewer interface {
	Review(name string) (Review, error)
}
type authorizerReviewer struct{ policyChecker rbac.SubjectLocator }

func NewAuthorizerReviewer(policyChecker rbac.SubjectLocator) Reviewer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &authorizerReviewer{policyChecker: policyChecker}
}
func (r *authorizerReviewer) Review(namespaceName string) (Review, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attributes := kauthorizer.AttributesRecord{Verb: "get", Namespace: namespaceName, Resource: "namespaces", Name: namespaceName, ResourceRequest: true}
	subjects, err := r.policyChecker.AllowedSubjects(attributes)
	review := &defaultReview{}
	review.users, review.groups = authorizationutil.RBACSubjectsToUsersAndGroups(subjects, attributes.GetNamespace())
	if err != nil {
		review.evaluationError = err.Error()
	}
	return review, nil
}

package selfsubjectrulesreview

import (
	"context"
	"fmt"
	goformat "fmt"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type REST struct{ ruleResolver authorizer.RuleResolver }

func NewREST(ruleResolver authorizer.RuleResolver) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &REST{ruleResolver}
}
func (r *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &authorizationapi.SelfSubjectRulesReview{}
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	selfSRR, ok := obj.(*authorizationapi.SelfSubjectRulesReview)
	if !ok {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("not a SelfSubjectRulesReview: %#v", obj))
	}
	user, ok := genericapirequest.UserFrom(ctx)
	if !ok {
		return nil, apierrors.NewBadRequest("no user present on request")
	}
	namespace := selfSRR.Spec.Namespace
	if namespace == "" {
		return nil, apierrors.NewBadRequest("no namespace on request")
	}
	resourceInfo, nonResourceInfo, incomplete, err := r.ruleResolver.RulesFor(user, namespace)
	ret := &authorizationapi.SelfSubjectRulesReview{Status: authorizationapi.SubjectRulesReviewStatus{ResourceRules: getResourceRules(resourceInfo), NonResourceRules: getNonResourceRules(nonResourceInfo), Incomplete: incomplete}}
	if err != nil {
		ret.Status.EvaluationError = err.Error()
	}
	return ret, nil
}
func getResourceRules(infos []authorizer.ResourceRuleInfo) []authorizationapi.ResourceRule {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rules := make([]authorizationapi.ResourceRule, len(infos))
	for i, info := range infos {
		rules[i] = authorizationapi.ResourceRule{Verbs: info.GetVerbs(), APIGroups: info.GetAPIGroups(), Resources: info.GetResources(), ResourceNames: info.GetResourceNames()}
	}
	return rules
}
func getNonResourceRules(infos []authorizer.NonResourceRuleInfo) []authorizationapi.NonResourceRule {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rules := make([]authorizationapi.NonResourceRule, len(infos))
	for i, info := range infos {
		rules[i] = authorizationapi.NonResourceRule{Verbs: info.GetVerbs(), NonResourceURLs: info.GetNonResourceURLs()}
	}
	return rules
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

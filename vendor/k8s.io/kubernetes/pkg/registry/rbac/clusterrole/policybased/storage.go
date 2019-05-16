package policybased

import (
	"context"
	"errors"
	goformat "fmt"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/apiserver/pkg/registry/rest"
	kapihelper "k8s.io/kubernetes/pkg/apis/core/helper"
	"k8s.io/kubernetes/pkg/apis/rbac"
	rbacregistry "k8s.io/kubernetes/pkg/registry/rbac"
	rbacregistryvalidation "k8s.io/kubernetes/pkg/registry/rbac/validation"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var groupResource = rbac.Resource("clusterroles")

type Storage struct {
	rest.StandardStorage
	authorizer   authorizer.Authorizer
	ruleResolver rbacregistryvalidation.AuthorizationRuleResolver
}

func NewStorage(s rest.StandardStorage, authorizer authorizer.Authorizer, ruleResolver rbacregistryvalidation.AuthorizationRuleResolver) *Storage {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Storage{s, authorizer, ruleResolver}
}
func (r *Storage) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}

var fullAuthority = []rbac.PolicyRule{rbac.NewRule("*").Groups("*").Resources("*").RuleOrDie(), rbac.NewRule("*").URLs("*").RuleOrDie()}

func (s *Storage) Create(ctx context.Context, obj runtime.Object, createValidatingAdmission rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if rbacregistry.EscalationAllowed(ctx) || rbacregistry.RoleEscalationAuthorized(ctx, s.authorizer) {
		return s.StandardStorage.Create(ctx, obj, createValidatingAdmission, options)
	}
	clusterRole := obj.(*rbac.ClusterRole)
	rules := clusterRole.Rules
	if err := rbacregistryvalidation.ConfirmNoEscalationInternal(ctx, s.ruleResolver, rules); err != nil {
		return nil, apierrors.NewForbidden(groupResource, clusterRole.Name, err)
	}
	if hasAggregationRule(clusterRole) {
		if err := rbacregistryvalidation.ConfirmNoEscalationInternal(ctx, s.ruleResolver, fullAuthority); err != nil {
			return nil, apierrors.NewForbidden(groupResource, clusterRole.Name, errors.New("must have cluster-admin privileges to use the aggregationRule"))
		}
	}
	return s.StandardStorage.Create(ctx, obj, createValidatingAdmission, options)
}
func (s *Storage) Update(ctx context.Context, name string, obj rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if rbacregistry.EscalationAllowed(ctx) || rbacregistry.RoleEscalationAuthorized(ctx, s.authorizer) {
		return s.StandardStorage.Update(ctx, name, obj, createValidation, updateValidation, forceAllowCreate, options)
	}
	nonEscalatingInfo := rest.WrapUpdatedObjectInfo(obj, func(ctx context.Context, obj runtime.Object, oldObj runtime.Object) (runtime.Object, error) {
		clusterRole := obj.(*rbac.ClusterRole)
		oldClusterRole := oldObj.(*rbac.ClusterRole)
		if rbacregistry.IsOnlyMutatingGCFields(obj, oldObj, kapihelper.Semantic) {
			return obj, nil
		}
		rules := clusterRole.Rules
		if err := rbacregistryvalidation.ConfirmNoEscalationInternal(ctx, s.ruleResolver, rules); err != nil {
			return nil, apierrors.NewForbidden(groupResource, clusterRole.Name, err)
		}
		if hasAggregationRule(clusterRole) || hasAggregationRule(oldClusterRole) {
			if err := rbacregistryvalidation.ConfirmNoEscalationInternal(ctx, s.ruleResolver, fullAuthority); err != nil {
				return nil, apierrors.NewForbidden(groupResource, clusterRole.Name, errors.New("must have cluster-admin privileges to use the aggregationRule"))
			}
		}
		return obj, nil
	})
	return s.StandardStorage.Update(ctx, name, nonEscalatingInfo, createValidation, updateValidation, forceAllowCreate, options)
}
func hasAggregationRule(clusterRole *rbac.ClusterRole) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return clusterRole.AggregationRule != nil && len(clusterRole.AggregationRule.ClusterRoleSelectors) > 0
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

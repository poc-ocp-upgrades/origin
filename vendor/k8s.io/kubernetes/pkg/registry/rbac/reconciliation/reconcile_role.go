package reconciliation

import (
	"fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/registry/rbac/validation"
	"reflect"
)

type ReconcileOperation string

var (
	ReconcileCreate   ReconcileOperation = "create"
	ReconcileUpdate   ReconcileOperation = "update"
	ReconcileRecreate ReconcileOperation = "recreate"
	ReconcileNone     ReconcileOperation = "none"
)

type RuleOwnerModifier interface {
	Get(namespace, name string) (RuleOwner, error)
	Create(RuleOwner) (RuleOwner, error)
	Update(RuleOwner) (RuleOwner, error)
}
type RuleOwner interface {
	GetObject() runtime.Object
	GetNamespace() string
	GetName() string
	GetLabels() map[string]string
	SetLabels(map[string]string)
	GetAnnotations() map[string]string
	SetAnnotations(map[string]string)
	GetRules() []rbacv1.PolicyRule
	SetRules([]rbacv1.PolicyRule)
	GetAggregationRule() *rbacv1.AggregationRule
	SetAggregationRule(*rbacv1.AggregationRule)
	DeepCopyRuleOwner() RuleOwner
}
type ReconcileRoleOptions struct {
	Role                   RuleOwner
	Confirm                bool
	RemoveExtraPermissions bool
	Client                 RuleOwnerModifier
}
type ReconcileClusterRoleResult struct {
	Role                            RuleOwner
	MissingRules                    []rbacv1.PolicyRule
	ExtraRules                      []rbacv1.PolicyRule
	MissingAggregationRuleSelectors []metav1.LabelSelector
	ExtraAggregationRuleSelectors   []metav1.LabelSelector
	Operation                       ReconcileOperation
	Protected                       bool
}

func (o *ReconcileRoleOptions) Run() (*ReconcileClusterRoleResult, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.run(0)
}
func (o *ReconcileRoleOptions) run(attempts int) (*ReconcileClusterRoleResult, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attempts > 2 {
		return nil, fmt.Errorf("exceeded maximum attempts")
	}
	var result *ReconcileClusterRoleResult
	existing, err := o.Client.Get(o.Role.GetNamespace(), o.Role.GetName())
	switch {
	case errors.IsNotFound(err):
		aggregationRule := o.Role.GetAggregationRule()
		if aggregationRule == nil {
			aggregationRule = &rbacv1.AggregationRule{}
		}
		result = &ReconcileClusterRoleResult{Role: o.Role, MissingRules: o.Role.GetRules(), MissingAggregationRuleSelectors: aggregationRule.ClusterRoleSelectors, Operation: ReconcileCreate}
	case err != nil:
		return nil, err
	default:
		result, err = computeReconciledRole(existing, o.Role, o.RemoveExtraPermissions)
		if err != nil {
			return nil, err
		}
	}
	if result.Protected {
		return result, nil
	}
	if !o.Confirm {
		return result, nil
	}
	switch result.Operation {
	case ReconcileCreate:
		created, err := o.Client.Create(result.Role)
		if errors.IsAlreadyExists(err) {
			return o.run(attempts + 1)
		}
		if err != nil {
			return nil, err
		}
		result.Role = created
	case ReconcileUpdate:
		updated, err := o.Client.Update(result.Role)
		if errors.IsNotFound(err) {
			return o.run(attempts + 1)
		}
		if err != nil {
			return nil, err
		}
		result.Role = updated
	case ReconcileNone:
	default:
		return nil, fmt.Errorf("invalid operation: %v", result.Operation)
	}
	return result, nil
}
func computeReconciledRole(existing, expected RuleOwner, removeExtraPermissions bool) (*ReconcileClusterRoleResult, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := &ReconcileClusterRoleResult{Operation: ReconcileNone}
	result.Protected = (existing.GetAnnotations()[rbacv1.AutoUpdateAnnotationKey] == "false")
	result.Role = existing.DeepCopyRuleOwner()
	result.Role.SetAnnotations(merge(expected.GetAnnotations(), result.Role.GetAnnotations()))
	if !reflect.DeepEqual(result.Role.GetAnnotations(), existing.GetAnnotations()) {
		result.Operation = ReconcileUpdate
	}
	result.Role.SetLabels(merge(expected.GetLabels(), result.Role.GetLabels()))
	if !reflect.DeepEqual(result.Role.GetLabels(), existing.GetLabels()) {
		result.Operation = ReconcileUpdate
	}
	_, result.ExtraRules = validation.Covers(expected.GetRules(), existing.GetRules())
	_, result.MissingRules = validation.Covers(existing.GetRules(), expected.GetRules())
	switch {
	case !removeExtraPermissions && len(result.MissingRules) > 0:
		result.Role.SetRules(append(result.Role.GetRules(), result.MissingRules...))
		result.Operation = ReconcileUpdate
	case removeExtraPermissions && (len(result.MissingRules) > 0 || len(result.ExtraRules) > 0):
		result.Role.SetRules(expected.GetRules())
		result.Operation = ReconcileUpdate
	}
	_, result.ExtraAggregationRuleSelectors = aggregationRuleCovers(expected.GetAggregationRule(), existing.GetAggregationRule())
	_, result.MissingAggregationRuleSelectors = aggregationRuleCovers(existing.GetAggregationRule(), expected.GetAggregationRule())
	switch {
	case expected.GetAggregationRule() == nil && existing.GetAggregationRule() != nil:
		result.Role.SetAggregationRule(nil)
		result.Operation = ReconcileUpdate
	case !removeExtraPermissions && len(result.MissingAggregationRuleSelectors) > 0:
		aggregationRule := result.Role.GetAggregationRule()
		if aggregationRule == nil {
			aggregationRule = &rbacv1.AggregationRule{}
		}
		aggregationRule.ClusterRoleSelectors = append(aggregationRule.ClusterRoleSelectors, result.MissingAggregationRuleSelectors...)
		result.Role.SetAggregationRule(aggregationRule)
		result.Operation = ReconcileUpdate
	case removeExtraPermissions && (len(result.MissingAggregationRuleSelectors) > 0 || len(result.ExtraAggregationRuleSelectors) > 0):
		result.Role.SetAggregationRule(expected.GetAggregationRule())
		result.Operation = ReconcileUpdate
	}
	return result, nil
}
func merge(maps ...map[string]string) map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var output map[string]string = nil
	for _, m := range maps {
		if m != nil && output == nil {
			output = map[string]string{}
		}
		for k, v := range m {
			output[k] = v
		}
	}
	return output
}
func aggregationRuleCovers(ownerRule, servantRule *rbacv1.AggregationRule) (bool, []metav1.LabelSelector) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case ownerRule == nil && servantRule == nil:
		return true, []metav1.LabelSelector{}
	case ownerRule == nil && servantRule != nil:
		return false, servantRule.ClusterRoleSelectors
	case ownerRule != nil && servantRule == nil:
		return true, []metav1.LabelSelector{}
	}
	ownerSelectors := ownerRule.ClusterRoleSelectors
	servantSelectors := servantRule.ClusterRoleSelectors
	uncoveredSelectors := []metav1.LabelSelector{}
	for _, servantSelector := range servantSelectors {
		covered := false
		for _, ownerSelector := range ownerSelectors {
			if equality.Semantic.DeepEqual(ownerSelector, servantSelector) {
				covered = true
				break
			}
		}
		if !covered {
			uncoveredSelectors = append(uncoveredSelectors, servantSelector)
		}
	}
	return (len(uncoveredSelectors) == 0), uncoveredSelectors
}

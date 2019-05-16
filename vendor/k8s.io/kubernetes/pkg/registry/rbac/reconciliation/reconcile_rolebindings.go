package reconciliation

import (
	"fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"reflect"
)

type RoleBindingModifier interface {
	Get(namespace, name string) (RoleBinding, error)
	Delete(namespace, name string, uid types.UID) error
	Create(RoleBinding) (RoleBinding, error)
	Update(RoleBinding) (RoleBinding, error)
}
type RoleBinding interface {
	GetObject() runtime.Object
	GetNamespace() string
	GetName() string
	GetUID() types.UID
	GetLabels() map[string]string
	SetLabels(map[string]string)
	GetAnnotations() map[string]string
	SetAnnotations(map[string]string)
	GetRoleRef() rbacv1.RoleRef
	GetSubjects() []rbacv1.Subject
	SetSubjects([]rbacv1.Subject)
	DeepCopyRoleBinding() RoleBinding
}
type ReconcileRoleBindingOptions struct {
	RoleBinding         RoleBinding
	Confirm             bool
	RemoveExtraSubjects bool
	Client              RoleBindingModifier
}
type ReconcileClusterRoleBindingResult struct {
	RoleBinding     RoleBinding
	MissingSubjects []rbacv1.Subject
	ExtraSubjects   []rbacv1.Subject
	Operation       ReconcileOperation
	Protected       bool
}

func (o *ReconcileRoleBindingOptions) Run() (*ReconcileClusterRoleBindingResult, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.run(0)
}
func (o *ReconcileRoleBindingOptions) run(attempts int) (*ReconcileClusterRoleBindingResult, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attempts > 3 {
		return nil, fmt.Errorf("exceeded maximum attempts")
	}
	var result *ReconcileClusterRoleBindingResult
	existingBinding, err := o.Client.Get(o.RoleBinding.GetNamespace(), o.RoleBinding.GetName())
	switch {
	case errors.IsNotFound(err):
		result = &ReconcileClusterRoleBindingResult{RoleBinding: o.RoleBinding, MissingSubjects: o.RoleBinding.GetSubjects(), Operation: ReconcileCreate}
	case err != nil:
		return nil, err
	default:
		result, err = computeReconciledRoleBinding(existingBinding, o.RoleBinding, o.RemoveExtraSubjects)
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
	case ReconcileRecreate:
		err := o.Client.Delete(existingBinding.GetNamespace(), existingBinding.GetName(), existingBinding.GetUID())
		switch {
		case err == nil, errors.IsNotFound(err):
		case errors.IsConflict(err):
			return o.run(attempts + 1)
		default:
			return nil, err
		}
		fallthrough
	case ReconcileCreate:
		created, err := o.Client.Create(result.RoleBinding)
		if errors.IsAlreadyExists(err) {
			return o.run(attempts + 1)
		}
		if err != nil {
			return nil, err
		}
		result.RoleBinding = created
	case ReconcileUpdate:
		updated, err := o.Client.Update(result.RoleBinding)
		if errors.IsNotFound(err) {
			return o.run(attempts + 1)
		}
		if err != nil {
			return nil, err
		}
		result.RoleBinding = updated
	case ReconcileNone:
	default:
		return nil, fmt.Errorf("invalid operation: %v", result.Operation)
	}
	return result, nil
}
func computeReconciledRoleBinding(existing, expected RoleBinding, removeExtraSubjects bool) (*ReconcileClusterRoleBindingResult, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := &ReconcileClusterRoleBindingResult{Operation: ReconcileNone}
	result.Protected = (existing.GetAnnotations()[rbacv1.AutoUpdateAnnotationKey] == "false")
	if expected.GetRoleRef() != existing.GetRoleRef() {
		result.RoleBinding = expected
		result.Operation = ReconcileRecreate
		return result, nil
	}
	result.RoleBinding = existing.DeepCopyRoleBinding()
	result.RoleBinding.SetAnnotations(merge(expected.GetAnnotations(), result.RoleBinding.GetAnnotations()))
	if !reflect.DeepEqual(result.RoleBinding.GetAnnotations(), existing.GetAnnotations()) {
		result.Operation = ReconcileUpdate
	}
	result.RoleBinding.SetLabels(merge(expected.GetLabels(), result.RoleBinding.GetLabels()))
	if !reflect.DeepEqual(result.RoleBinding.GetLabels(), existing.GetLabels()) {
		result.Operation = ReconcileUpdate
	}
	result.MissingSubjects, result.ExtraSubjects = diffSubjectLists(expected.GetSubjects(), existing.GetSubjects())
	switch {
	case !removeExtraSubjects && len(result.MissingSubjects) > 0:
		result.RoleBinding.SetSubjects(append(result.RoleBinding.GetSubjects(), result.MissingSubjects...))
		result.Operation = ReconcileUpdate
	case removeExtraSubjects && (len(result.MissingSubjects) > 0 || len(result.ExtraSubjects) > 0):
		result.RoleBinding.SetSubjects(expected.GetSubjects())
		result.Operation = ReconcileUpdate
	}
	return result, nil
}
func contains(list []rbacv1.Subject, item rbacv1.Subject) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, listItem := range list {
		if listItem == item {
			return true
		}
	}
	return false
}
func diffSubjectLists(list1 []rbacv1.Subject, list2 []rbacv1.Subject) (list1Only []rbacv1.Subject, list2Only []rbacv1.Subject) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, list1Item := range list1 {
		if !contains(list2, list1Item) {
			if !contains(list1Only, list1Item) {
				list1Only = append(list1Only, list1Item)
			}
		}
	}
	for _, list2Item := range list2 {
		if !contains(list1, list2Item) {
			if !contains(list2Only, list2Item) {
				list2Only = append(list2Only, list2Item)
			}
		}
	}
	return
}

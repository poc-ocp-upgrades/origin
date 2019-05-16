package reconciliation

import (
	goformat "fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type ClusterRoleRuleOwner struct{ ClusterRole *rbacv1.ClusterRole }

func (o ClusterRoleRuleOwner) GetObject() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRole
}
func (o ClusterRoleRuleOwner) GetNamespace() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRole.Namespace
}
func (o ClusterRoleRuleOwner) GetName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRole.Name
}
func (o ClusterRoleRuleOwner) GetLabels() map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRole.Labels
}
func (o ClusterRoleRuleOwner) SetLabels(in map[string]string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o.ClusterRole.Labels = in
}
func (o ClusterRoleRuleOwner) GetAnnotations() map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRole.Annotations
}
func (o ClusterRoleRuleOwner) SetAnnotations(in map[string]string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o.ClusterRole.Annotations = in
}
func (o ClusterRoleRuleOwner) GetRules() []rbacv1.PolicyRule {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRole.Rules
}
func (o ClusterRoleRuleOwner) SetRules(in []rbacv1.PolicyRule) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o.ClusterRole.Rules = in
}
func (o ClusterRoleRuleOwner) GetAggregationRule() *rbacv1.AggregationRule {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return o.ClusterRole.AggregationRule
}
func (o ClusterRoleRuleOwner) SetAggregationRule(in *rbacv1.AggregationRule) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o.ClusterRole.AggregationRule = in
}

type ClusterRoleModifier struct {
	Client rbacv1client.ClusterRoleInterface
}

func (c ClusterRoleModifier) Get(namespace, name string) (RuleOwner, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret, err := c.Client.Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return ClusterRoleRuleOwner{ClusterRole: ret}, err
}
func (c ClusterRoleModifier) Create(in RuleOwner) (RuleOwner, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret, err := c.Client.Create(in.(ClusterRoleRuleOwner).ClusterRole)
	if err != nil {
		return nil, err
	}
	return ClusterRoleRuleOwner{ClusterRole: ret}, err
}
func (c ClusterRoleModifier) Update(in RuleOwner) (RuleOwner, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret, err := c.Client.Update(in.(ClusterRoleRuleOwner).ClusterRole)
	if err != nil {
		return nil, err
	}
	return ClusterRoleRuleOwner{ClusterRole: ret}, err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

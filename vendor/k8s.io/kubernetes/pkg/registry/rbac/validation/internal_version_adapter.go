package validation

import (
	"context"
	goformat "fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/kubernetes/pkg/apis/rbac"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ConfirmNoEscalationInternal(ctx context.Context, ruleResolver AuthorizationRuleResolver, inRules []rbac.PolicyRule) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rules := []rbacv1.PolicyRule{}
	for i := range inRules {
		v1Rule := rbacv1.PolicyRule{}
		err := rbacv1helpers.Convert_rbac_PolicyRule_To_v1_PolicyRule(&inRules[i], &v1Rule, nil)
		if err != nil {
			return err
		}
		rules = append(rules, v1Rule)
	}
	return ConfirmNoEscalation(ctx, ruleResolver, rules)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

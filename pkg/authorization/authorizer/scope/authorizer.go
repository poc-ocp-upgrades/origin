package scope

import (
	"fmt"
	goformat "fmt"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	rbaclisters "k8s.io/client-go/listers/rbac/v1"
	authorizerrbac "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type scopeAuthorizer struct{ clusterRoleGetter rbaclisters.ClusterRoleLister }

func NewAuthorizer(clusterRoleGetter rbaclisters.ClusterRoleLister) authorizer.Authorizer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &scopeAuthorizer{clusterRoleGetter: clusterRoleGetter}
}
func (a *scopeAuthorizer) Authorize(attributes authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	user := attributes.GetUser()
	if user == nil {
		return authorizer.DecisionNoOpinion, "", fmt.Errorf("user missing from context")
	}
	scopes := user.GetExtra()[authorizationapi.ScopesKey]
	if len(scopes) == 0 {
		return authorizer.DecisionNoOpinion, "", nil
	}
	nonFatalErrors := ""
	rules, err := ScopesToRules(scopes, attributes.GetNamespace(), a.clusterRoleGetter)
	if err != nil {
		nonFatalErrors = fmt.Sprintf(", additionally the following non-fatal errors were reported: %v", err)
	}
	if authorizerrbac.RulesAllow(attributes, rules...) {
		return authorizer.DecisionNoOpinion, "", nil
	}
	return authorizer.DecisionDeny, fmt.Sprintf("scopes %v prevent this action%s", scopes, nonFatalErrors), nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

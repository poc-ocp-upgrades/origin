package util

import (
	goformat "fmt"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func ToDefaultAuthorizationAttributes(user user.Info, namespace string, in authorizationapi.Action) authorizer.Attributes {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tokens := strings.SplitN(in.Resource, "/", 2)
	resource := ""
	subresource := ""
	switch {
	case len(tokens) == 2:
		subresource = tokens[1]
		fallthrough
	case len(tokens) == 1:
		resource = tokens[0]
	}
	return &authorizer.AttributesRecord{User: user, Verb: in.Verb, Namespace: namespace, APIGroup: in.Group, APIVersion: in.Version, Resource: resource, Subresource: subresource, Name: in.ResourceName, ResourceRequest: !in.IsNonResourceURL, Path: in.Path}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

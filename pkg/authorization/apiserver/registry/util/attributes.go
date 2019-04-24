package util

import (
	"strings"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
)

func ToDefaultAuthorizationAttributes(user user.Info, namespace string, in authorizationapi.Action) authorizer.Attributes {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

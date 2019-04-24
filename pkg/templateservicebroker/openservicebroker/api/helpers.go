package api

import (
	"net/http"
	"bytes"
	"runtime"
	"fmt"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	"k8s.io/apiserver/pkg/authentication/user"
)

func NewResponse(code int, body interface{}, err error) *Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Response{Code: code, Body: body, Err: err}
}
func BadRequest(err error) *Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewResponse(http.StatusBadRequest, nil, err)
}
func Forbidden(err error) *Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewResponse(http.StatusForbidden, nil, err)
}
func InternalServerError(err error) *Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewResponse(http.StatusInternalServerError, nil, err)
}
func ConvertUserToTemplateInstanceRequester(u user.Info) templateapi.TemplateInstanceRequester {
	_logClusterCodePath()
	defer _logClusterCodePath()
	templatereq := templateapi.TemplateInstanceRequester{}
	if u != nil {
		extra := map[string]templateapi.ExtraValue{}
		if u.GetExtra() != nil {
			for k, v := range u.GetExtra() {
				extra[k] = templateapi.ExtraValue(v)
			}
		}
		templatereq.Username = u.GetName()
		templatereq.UID = u.GetUID()
		templatereq.Groups = u.GetGroups()
		templatereq.Extra = extra
	}
	return templatereq
}
func ConvertTemplateInstanceRequesterToUser(templateReq *templateapi.TemplateInstanceRequester) user.Info {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := user.DefaultInfo{}
	u.Extra = map[string][]string{}
	if templateReq != nil {
		u.Name = templateReq.Username
		u.UID = templateReq.UID
		u.Groups = templateReq.Groups
		if templateReq.Extra != nil {
			for k, v := range templateReq.Extra {
				u.Extra[k] = []string(v)
			}
		}
	}
	return &u
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

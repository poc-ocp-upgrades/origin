package api

import (
	goformat "fmt"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	"k8s.io/apiserver/pkg/authentication/user"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func NewResponse(code int, body interface{}, err error) *Response {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Response{Code: code, Body: body, Err: err}
}
func BadRequest(err error) *Response {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewResponse(http.StatusBadRequest, nil, err)
}
func Forbidden(err error) *Response {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewResponse(http.StatusForbidden, nil, err)
}
func InternalServerError(err error) *Response {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewResponse(http.StatusInternalServerError, nil, err)
}
func ConvertUserToTemplateInstanceRequester(u user.Info) templateapi.TemplateInstanceRequester {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

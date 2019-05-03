package errorpage

import (
	"bytes"
	"fmt"
	"github.com/openshift/origin/pkg/util/httprequest"
	"html/template"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog"
	"net/http"
)

type ErrorPage struct{ render ErrorPageRenderer }

func NewErrorPageHandler(renderer ErrorPageRenderer) *ErrorPage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ErrorPage{render: renderer}
}
func (p *ErrorPage) AuthenticationError(err error, w http.ResponseWriter, req *http.Request) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Errorf("AuthenticationError: %v", err)
	if !httprequest.PrefersHTML(req) {
		return false, err
	}
	errorData := ErrorData{}
	errorData.ErrorCode = AuthenticationErrorCode(err)
	errorData.Error = AuthenticationErrorMessage(errorData.ErrorCode)
	p.render.Render(errorData, w, req)
	return true, nil
}
func (p *ErrorPage) GrantError(err error, w http.ResponseWriter, req *http.Request) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Errorf("GrantError: %v", err)
	if !httprequest.PrefersHTML(req) {
		return false, err
	}
	errorData := ErrorData{}
	errorData.ErrorCode = GrantErrorCode(err)
	errorData.Error = GrantErrorMessage(errorData.ErrorCode)
	p.render.Render(errorData, w, req)
	return true, nil
}

type ErrorData struct {
	Error     string
	ErrorCode string
}
type ErrorPageRenderer interface {
	Render(data ErrorData, w http.ResponseWriter, req *http.Request)
}
type errorPageTemplateRenderer struct{ errorPageTemplate *template.Template }

func NewErrorPageTemplateRenderer(templateFile string) (ErrorPageRenderer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := &errorPageTemplateRenderer{}
	if len(templateFile) > 0 {
		customTemplate, err := template.ParseFiles(templateFile)
		if err != nil {
			return nil, err
		}
		r.errorPageTemplate = customTemplate
	} else {
		r.errorPageTemplate = defaultErrorPageTemplate
	}
	return r, nil
}
func (r *errorPageTemplateRenderer) Render(data ErrorData, w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := r.errorPageTemplate.Execute(w, data); err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to render error page template: %v", err))
	}
}
func ValidateErrorPageTemplate(templateContent []byte) []error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var allErrs []error
	template, err := template.New("errorPageTemplateTest").Parse(string(templateContent))
	if err != nil {
		return append(allErrs, err)
	}
	var buffer bytes.Buffer
	err = template.Execute(&buffer, ErrorData{})
	if err != nil {
		return append(allErrs, err)
	}
	return allErrs
}

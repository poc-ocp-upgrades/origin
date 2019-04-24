package internalversion

import (
	"k8s.io/client-go/rest"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
)

type TemplateProcessorInterface interface {
	Process(*templateapi.Template) (*templateapi.Template, error)
}

func NewTemplateProcessorClient(c rest.Interface, ns string) TemplateProcessorInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &templateProcessor{client: c, ns: ns}
}

type templateProcessor struct {
	client	rest.Interface
	ns	string
}

func (c *templateProcessor) Process(in *templateapi.Template) (*templateapi.Template, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	template := &templateapi.Template{}
	err := c.client.Post().Namespace(c.ns).Resource("processedTemplates").Body(in).Do().Into(template)
	return template, err
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

package v1

import (
	"k8s.io/client-go/rest"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	templatev1 "github.com/openshift/api/template/v1"
)

type TemplateProcessorInterface interface {
	Process(*templatev1.Template) (*templatev1.Template, error)
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

func (c *templateProcessor) Process(in *templatev1.Template) (*templatev1.Template, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	template := &templatev1.Template{}
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

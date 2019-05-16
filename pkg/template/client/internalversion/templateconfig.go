package internalversion

import (
	goformat "fmt"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	"k8s.io/client-go/rest"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type TemplateProcessorInterface interface {
	Process(*templateapi.Template) (*templateapi.Template, error)
}

func NewTemplateProcessorClient(c rest.Interface, ns string) TemplateProcessorInterface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &templateProcessor{client: c, ns: ns}
}

type templateProcessor struct {
	client rest.Interface
	ns     string
}

func (c *templateProcessor) Process(in *templateapi.Template) (*templateapi.Template, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	template := &templateapi.Template{}
	err := c.client.Post().Namespace(c.ns).Resource("processedTemplates").Body(in).Do().Into(template)
	return template, err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

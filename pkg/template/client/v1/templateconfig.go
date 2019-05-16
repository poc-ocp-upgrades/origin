package v1

import (
	goformat "fmt"
	templatev1 "github.com/openshift/api/template/v1"
	"k8s.io/client-go/rest"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type TemplateProcessorInterface interface {
	Process(*templatev1.Template) (*templatev1.Template, error)
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

func (c *templateProcessor) Process(in *templatev1.Template) (*templatev1.Template, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	template := &templatev1.Template{}
	err := c.client.Post().Namespace(c.ns).Resource("processedTemplates").Body(in).Do().Into(template)
	return template, err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

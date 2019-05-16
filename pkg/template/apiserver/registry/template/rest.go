package template

import (
	"context"
	goformat "fmt"
	"github.com/openshift/api/template"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	templatevalidation "github.com/openshift/origin/pkg/template/apis/template/validation"
	"github.com/openshift/origin/pkg/template/generator"
	"github.com/openshift/origin/pkg/template/templateprocessing"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/klog"
	"math/rand"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type REST struct{}

var _ rest.Creater = &REST{}
var _ rest.Scoper = &REST{}

func NewREST() *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &REST{}
}
func (s *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &templateapi.Template{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (s *REST) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tpl, ok := obj.(*templateapi.Template)
	if !ok {
		return nil, errors.NewBadRequest("not a template")
	}
	if errs := templatevalidation.ValidateProcessedTemplate(tpl); len(errs) > 0 {
		return nil, errors.NewInvalid(template.Kind("Template"), tpl.Name, errs)
	}
	generators := map[string]generator.Generator{"expression": generator.NewExpressionValueGenerator(rand.New(rand.NewSource(time.Now().UnixNano())))}
	processor := templateprocessing.NewProcessor(generators)
	if errs := processor.Process(tpl); len(errs) > 0 {
		klog.V(1).Infof(errs.ToAggregate().Error())
		return nil, errors.NewInvalid(template.Kind("Template"), tpl.Name, errs)
	}
	for i := range tpl.Objects {
		tpl.Objects[i] = runtime.NewEncodable(unstructured.UnstructuredJSONScheme, tpl.Objects[i])
	}
	return tpl, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

package templateprocessing

import (
	goformat "fmt"
	templatev1 "github.com/openshift/api/template/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/dynamic"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type DynamicTemplateProcessor interface {
	ProcessToList(template *templatev1.Template) (*unstructured.UnstructuredList, error)
	ProcessToListFromUnstructured(unstructuredTemplate *unstructured.Unstructured) (*unstructured.UnstructuredList, error)
}
type dynamicTemplateProcessor struct{ client dynamic.Interface }

func NewDynamicTemplateProcessor(client dynamic.Interface) DynamicTemplateProcessor {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &dynamicTemplateProcessor{client: client}
}
func (c *dynamicTemplateProcessor) ProcessToList(template *templatev1.Template) (*unstructured.UnstructuredList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	versionedTemplate, err := scheme.ConvertToVersion(template, templatev1.GroupVersion)
	if err != nil {
		return nil, err
	}
	unstructuredTemplate, err := runtime.DefaultUnstructuredConverter.ToUnstructured(versionedTemplate)
	if err != nil {
		return nil, err
	}
	return c.ProcessToListFromUnstructured(&unstructured.Unstructured{Object: unstructuredTemplate})
}
func (c *dynamicTemplateProcessor) ProcessToListFromUnstructured(unstructuredTemplate *unstructured.Unstructured) (*unstructured.UnstructuredList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	processedTemplate, err := c.client.Resource(templatev1.GroupVersion.WithResource("processedtemplates")).Namespace("default").Create(unstructuredTemplate, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	if err := unstructured.SetNestedField(processedTemplate.Object, processedTemplate.Object["objects"], "items"); err != nil {
		return nil, err
	}
	return processedTemplate.ToList()
}

var (
	scheme = runtime.NewScheme()
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	utilruntime.Must(templatev1.Install(scheme))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

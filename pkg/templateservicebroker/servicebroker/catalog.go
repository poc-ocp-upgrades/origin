package servicebroker

import (
	"net/http"
	"strings"
	"k8s.io/apimachinery/pkg/labels"
	jsschema "github.com/lestrrat/go-jsschema"
	"k8s.io/klog"
	templateapiv1 "github.com/openshift/api/template/v1"
	oapi "github.com/openshift/origin/pkg/api"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	"github.com/openshift/origin/pkg/templateservicebroker/openservicebroker/api"
)

const (
	noDescriptionProvided = "No description provided."
)

var annotationMap = map[string]string{oapi.OpenShiftDisplayName: api.ServiceMetadataDisplayName, oapi.OpenShiftLongDescriptionAnnotation: api.ServiceMetadataLongDescription, oapi.OpenShiftProviderDisplayNameAnnotation: api.ServiceMetadataProviderDisplayName, oapi.OpenShiftDocumentationURLAnnotation: api.ServiceMetadataDocumentationURL, oapi.OpenShiftSupportURLAnnotation: api.ServiceMetadataSupportURL, templateapi.IconClassAnnotation: templateapi.ServiceMetadataIconClass}

func serviceFromTemplate(template *templateapiv1.Template) *api.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	metadata := make(map[string]interface{})
	for srcname, dstname := range annotationMap {
		if value, ok := template.Annotations[srcname]; ok {
			metadata[dstname] = value
		}
	}
	properties := map[string]*jsschema.Schema{}
	paramOrdering := []string{}
	required := []string{}
	for _, param := range template.Parameters {
		properties[param.Name] = &jsschema.Schema{Title: param.DisplayName, Description: param.Description, Default: param.Value, Type: []jsschema.PrimitiveType{jsschema.StringType}}
		if param.Required && param.Generate == "" {
			required = append(required, param.Name)
		}
		paramOrdering = append(paramOrdering, param.Name)
	}
	bindable := strings.ToLower(template.Annotations[templateapi.BindableAnnotation]) != "false"
	plan := api.Plan{ID: string(template.UID), Name: "default", Description: "Default plan", Free: true, Bindable: bindable, Schemas: api.Schema{ServiceInstance: api.ServiceInstances{Create: map[string]*jsschema.Schema{"parameters": {SchemaRef: jsschema.SchemaURL, Type: []jsschema.PrimitiveType{jsschema.ObjectType}, Properties: properties, Required: required}}}, ServiceBinding: api.ServiceBindings{Create: map[string]*jsschema.Schema{"parameters": {SchemaRef: jsschema.SchemaURL, Type: []jsschema.PrimitiveType{jsschema.ObjectType}, Properties: map[string]*jsschema.Schema{}, Required: []string{}}}}}}
	plan.Metadata = make(map[string]interface{})
	plan.Metadata["schemas"] = api.ParameterSchemas{ServiceInstance: api.ParameterSchema{Create: api.OpenShiftMetadata{OpenShiftFormDefinition: paramOrdering}}}
	description := template.Annotations["description"]
	if description == "" {
		description = noDescriptionProvided
	}
	return &api.Service{Name: template.Name, ID: string(template.UID), Description: description, Tags: strings.Split(template.Annotations["tags"], ","), Bindable: bindable, Metadata: metadata, Plans: []api.Plan{plan}}
}
func (b *Broker) Catalog() *api.Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("Template service broker: Catalog")
	var services []*api.Service
	for namespace := range b.templateNamespaces {
		templates, err := b.lister.Templates(namespace).List(labels.Everything())
		if err != nil {
			return api.InternalServerError(err)
		}
		for _, template := range templates {
			services = append(services, serviceFromTemplate(template))
		}
	}
	return api.NewResponse(http.StatusOK, &api.CatalogResponse{Services: services}, nil)
}

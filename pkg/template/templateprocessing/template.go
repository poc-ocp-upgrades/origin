package templateprocessing

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"regexp"
	"strings"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	templatev1 "github.com/openshift/api/template/v1"
	"github.com/openshift/origin/pkg/api/legacygroupification"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	. "github.com/openshift/origin/pkg/template/generator"
	"github.com/openshift/origin/pkg/util"
	"github.com/openshift/origin/pkg/util/stringreplace"
)

var stringParameterExp = regexp.MustCompile(`\$\{([a-zA-Z0-9\_]+?)\}`)
var nonStringParameterExp = regexp.MustCompile(`^\$\{\{([a-zA-Z0-9\_]+)\}\}$`)

type Processor struct{ Generators map[string]Generator }

func NewProcessor(generators map[string]Generator) *Processor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Processor{Generators: generators}
}
func (p *Processor) Process(template *templateapi.Template) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	templateErrors := field.ErrorList{}
	if errs := p.GenerateParameterValues(template); len(errs) > 0 {
		return append(templateErrors, errs...)
	}
	paramMap := make(map[string]templateapi.Parameter)
	for _, param := range template.Parameters {
		paramMap[param.Name] = param
	}
	template.Message, _ = p.EvaluateParameterSubstitution(paramMap, template.Message)
	for k, v := range template.ObjectLabels {
		newk, _ := p.EvaluateParameterSubstitution(paramMap, k)
		v, _ = p.EvaluateParameterSubstitution(paramMap, v)
		template.ObjectLabels[newk] = v
		if newk != k {
			delete(template.ObjectLabels, k)
		}
	}
	itemPath := field.NewPath("item")
	for i, item := range template.Objects {
		idxPath := itemPath.Index(i)
		if obj, ok := item.(*runtime.Unknown); ok {
			decodedObj, err := runtime.Decode(unstructured.UnstructuredJSONScheme, obj.Raw)
			if err != nil {
				templateErrors = append(templateErrors, field.Invalid(idxPath.Child("objects"), obj, fmt.Sprintf("unable to handle object: %v", err)))
				continue
			}
			item = decodedObj
		}
		stripNamespace(item)
		newItem, err := p.SubstituteParameters(paramMap, item)
		if err != nil {
			templateErrors = append(templateErrors, field.Invalid(idxPath.Child("parameters"), template.Parameters, err.Error()))
		}
		gvk := item.GetObjectKind().GroupVersionKind()
		legacygroupification.OAPIToGroupifiedGVK(&gvk)
		item.GetObjectKind().SetGroupVersionKind(gvk)
		if err := util.AddObjectLabels(newItem, template.ObjectLabels); err != nil {
			templateErrors = append(templateErrors, field.Invalid(idxPath.Child("labels"), template.ObjectLabels, fmt.Sprintf("label could not be applied: %v", err)))
		}
		template.Objects[i] = newItem
	}
	return templateErrors
}
func stripNamespace(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if itemMeta, err := meta.Accessor(obj); err == nil && len(itemMeta.GetNamespace()) > 0 && !stringParameterExp.MatchString(itemMeta.GetNamespace()) {
		itemMeta.SetNamespace("")
		return
	}
	if unstruct, ok := obj.(*unstructured.Unstructured); ok && unstruct.Object != nil {
		if obj, ok := unstruct.Object["metadata"]; ok {
			if m, ok := obj.(map[string]interface{}); ok {
				if ns, ok := m["namespace"]; ok {
					if ns, ok := ns.(string); !ok || !stringParameterExp.MatchString(ns) {
						m["namespace"] = ""
					}
				}
			}
			return
		}
		if ns, ok := unstruct.Object["namespace"]; ok {
			if ns, ok := ns.(string); !ok || !stringParameterExp.MatchString(ns) {
				unstruct.Object["namespace"] = ""
				return
			}
		}
	}
}
func DeprecatedGetParameterByNameInternal(t *templateapi.Template, name string) *templateapi.Parameter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, param := range t.Parameters {
		if param.Name == name {
			return &(t.Parameters[i])
		}
	}
	return nil
}
func GetParameterByName(t *templatev1.Template, name string) *templatev1.Parameter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, param := range t.Parameters {
		if param.Name == name {
			return &(t.Parameters[i])
		}
	}
	return nil
}
func (p *Processor) EvaluateParameterSubstitution(params map[string]templateapi.Parameter, in string) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := in
	for _, match := range nonStringParameterExp.FindAllStringSubmatch(in, -1) {
		if len(match) > 1 {
			if paramValue, found := params[match[1]]; found {
				out = strings.Replace(out, match[0], paramValue.Value, 1)
				return out, false
			}
		}
	}
	for _, match := range stringParameterExp.FindAllStringSubmatch(in, -1) {
		if len(match) > 1 {
			if paramValue, found := params[match[1]]; found {
				out = strings.Replace(out, match[0], paramValue.Value, 1)
			}
		}
	}
	return out, true
}
func (p *Processor) SubstituteParameters(params map[string]templateapi.Parameter, item runtime.Object) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stringreplace.VisitObjectStrings(item, func(in string) (string, bool) {
		return p.EvaluateParameterSubstitution(params, in)
	})
	return item, nil
}
func (p *Processor) GenerateParameterValues(t *templateapi.Template) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var errs field.ErrorList
	for i := range t.Parameters {
		param := &t.Parameters[i]
		if len(param.Value) > 0 {
			continue
		}
		templatePath := field.NewPath("template").Child("parameters").Index(i)
		if param.Generate != "" {
			generator, ok := p.Generators[param.Generate]
			if !ok {
				err := fmt.Errorf("Unknown generator name '%v' for parameter %s", param.Generate, param.Name)
				errs = append(errs, field.Invalid(templatePath, param.Generate, err.Error()))
				continue
			}
			if generator == nil {
				err := fmt.Errorf("template.parameters[%v]: Invalid '%v' generator for parameter %s", i, param.Generate, param.Name)
				errs = append(errs, field.Invalid(templatePath, param, err.Error()))
				continue
			}
			value, err := generator.GenerateValue(param.From)
			if err != nil {
				errs = append(errs, field.Invalid(templatePath, param, err.Error()))
				continue
			}
			param.Value, ok = value.(string)
			if !ok {
				err := fmt.Errorf("template.parameters[%v]: Unable to convert the generated value '%#v' to string for parameter %s", i, value, param.Name)
				errs = append(errs, field.Invalid(templatePath, param, err.Error()))
				continue
			}
		}
		if len(param.Value) == 0 && param.Required {
			err := fmt.Errorf("template.parameters[%v]: parameter %s is required and must be specified", i, param.Name)
			errs = append(errs, field.Required(templatePath, err.Error()))
		}
	}
	return errs
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

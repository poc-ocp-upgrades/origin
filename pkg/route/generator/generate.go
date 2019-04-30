package generator

import (
	"fmt"
	"strconv"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/kubernetes/pkg/kubectl/generate"
	routev1 "github.com/openshift/api/route/v1"
)

type RouteGenerator struct{}

var _ generate.Generator = RouteGenerator{}

func (RouteGenerator) ParamNames() []generate.GeneratorParam {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []generate.GeneratorParam{{Name: "labels", Required: false}, {Name: "default-name", Required: true}, {Name: "port", Required: false}, {Name: "name", Required: false}, {Name: "hostname", Required: false}, {Name: "path", Required: false}, {Name: "wildcard-policy", Required: false}}
}
func (RouteGenerator) Generate(genericParams map[string]interface{}) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		labels	map[string]string
		err	error
	)
	params := map[string]string{}
	for key, value := range genericParams {
		strVal, isString := value.(string)
		if !isString {
			return nil, fmt.Errorf("expected string, saw %v for '%s'", value, key)
		}
		params[key] = strVal
	}
	labelString, found := params["labels"]
	if found && len(labelString) > 0 {
		labels, err = generate.ParseLabels(labelString)
		if err != nil {
			return nil, err
		}
	}
	name, found := params["name"]
	if !found || len(name) == 0 {
		name, found = params["default-name"]
		if !found || len(name) == 0 {
			return nil, fmt.Errorf("'name' is a required parameter.")
		}
	}
	route := &routev1.Route{ObjectMeta: metav1.ObjectMeta{Name: name, Labels: labels}, Spec: routev1.RouteSpec{Host: params["hostname"], WildcardPolicy: routev1.WildcardPolicyType(params["wildcard-policy"]), Path: params["path"], To: routev1.RouteTargetReference{Name: params["default-name"]}}}
	portString := params["port"]
	if len(portString) > 0 {
		var targetPort intstr.IntOrString
		if port, err := strconv.Atoi(portString); err == nil {
			targetPort = intstr.FromInt(port)
		} else {
			targetPort = intstr.FromString(portString)
		}
		route.Spec.Port = &routev1.RoutePort{TargetPort: targetPort}
	}
	return route, nil
}

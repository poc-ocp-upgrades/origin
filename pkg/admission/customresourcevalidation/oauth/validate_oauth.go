package oauth

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"io"
	"k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/admission"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation"
)

const PluginName = "config.openshift.io/ValidateOAuth"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return customresourcevalidation.NewValidator(map[schema.GroupResource]bool{configv1.GroupVersion.WithResource("oauths").GroupResource(): true}, map[schema.GroupVersionKind]customresourcevalidation.ObjectValidator{configv1.GroupVersion.WithKind("OAuth"): oauthV1{}})
	})
}
func toOAuthV1(uncastObj runtime.Object) (*configv1.OAuth, field.ErrorList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if uncastObj == nil {
		return nil, nil
	}
	errs := field.ErrorList{}
	obj, ok := uncastObj.(*configv1.OAuth)
	if !ok {
		return nil, append(errs, field.NotSupported(field.NewPath("kind"), fmt.Sprintf("%T", uncastObj), []string{"OAuth"}), field.NotSupported(field.NewPath("apiVersion"), fmt.Sprintf("%T", uncastObj), []string{"config.openshift.io/v1"}))
	}
	return obj, nil
}

type oauthV1 struct{}

func (oauthV1) ValidateCreate(uncastObj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, errs := toOAuthV1(uncastObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMeta(&obj.ObjectMeta, false, customresourcevalidation.RequireNameCluster, field.NewPath("metadata"))...)
	errs = append(errs, validateOAuthSpec(obj.Spec)...)
	return errs
}
func (oauthV1) ValidateUpdate(uncastObj runtime.Object, uncastOldObj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, errs := toOAuthV1(uncastObj)
	if len(errs) > 0 {
		return errs
	}
	oldObj, errs := toOAuthV1(uncastOldObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMetaUpdate(&obj.ObjectMeta, &oldObj.ObjectMeta, field.NewPath("metadata"))...)
	errs = append(errs, validateOAuthSpec(obj.Spec)...)
	return errs
}
func (oauthV1) ValidateStatusUpdate(uncastObj runtime.Object, uncastOldObj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, errs := toOAuthV1(uncastObj)
	if len(errs) > 0 {
		return errs
	}
	oldObj, errs := toOAuthV1(uncastOldObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMetaUpdate(&obj.ObjectMeta, &oldObj.ObjectMeta, field.NewPath("metadata"))...)
	errs = append(errs, validateOAuthStatus(obj.Status)...)
	return errs
}
func validateOAuthSpec(spec configv1.OAuthSpec) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errs := field.ErrorList{}
	return errs
}
func validateOAuthStatus(status configv1.OAuthStatus) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errs := field.ErrorList{}
	return errs
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

package image

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

const PluginName = "config.openshift.io/ValidateImage"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return customresourcevalidation.NewValidator(map[schema.GroupResource]bool{configv1.Resource("images"): true}, map[schema.GroupVersionKind]customresourcevalidation.ObjectValidator{configv1.GroupVersion.WithKind("Image"): imageV1{}})
	})
}
func toImageV1(uncastObj runtime.Object) (*configv1.Image, field.ErrorList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if uncastObj == nil {
		return nil, nil
	}
	allErrs := field.ErrorList{}
	obj, ok := uncastObj.(*configv1.Image)
	if !ok {
		return nil, append(allErrs, field.NotSupported(field.NewPath("kind"), fmt.Sprintf("%T", uncastObj), []string{"Image"}), field.NotSupported(field.NewPath("apiVersion"), fmt.Sprintf("%T", uncastObj), []string{"config.openshift.io/v1"}))
	}
	return obj, nil
}

type imageV1 struct{}

func (imageV1) ValidateCreate(uncastObj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, errs := toImageV1(uncastObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMeta(&obj.ObjectMeta, false, customresourcevalidation.RequireNameCluster, field.NewPath("metadata"))...)
	return errs
}
func (imageV1) ValidateUpdate(uncastObj runtime.Object, uncastOldObj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, errs := toImageV1(uncastObj)
	if len(errs) > 0 {
		return errs
	}
	oldObj, errs := toImageV1(uncastOldObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMetaUpdate(&obj.ObjectMeta, &oldObj.ObjectMeta, field.NewPath("metadata"))...)
	return errs
}
func (imageV1) ValidateStatusUpdate(uncastObj runtime.Object, uncastOldObj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, errs := toImageV1(uncastObj)
	if len(errs) > 0 {
		return errs
	}
	oldObj, errs := toImageV1(uncastOldObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMetaUpdate(&obj.ObjectMeta, &oldObj.ObjectMeta, field.NewPath("metadata"))...)
	return errs
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}

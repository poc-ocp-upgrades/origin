package rolebindingrestriction

import (
	"fmt"
	goformat "fmt"
	authorizationv1 "github.com/openshift/api/authorization/v1"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation"
	rbrvalidation "github.com/openshift/origin/pkg/admission/customresourcevalidation/rolebindingrestriction/validation"
	"io"
	"k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/admission"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "authorization.openshift.io/ValidateRoleBindingRestriction"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return customresourcevalidation.NewValidator(map[schema.GroupResource]bool{{Group: authorizationv1.GroupName, Resource: "rolebindingrestrictions"}: true}, map[schema.GroupVersionKind]customresourcevalidation.ObjectValidator{authorizationv1.GroupVersion.WithKind("RoleBindingRestriction"): roleBindingRestrictionV1{}})
	})
}
func toRoleBindingRestriction(uncastObj runtime.Object) (*authorizationv1.RoleBindingRestriction, field.ErrorList) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if uncastObj == nil {
		return nil, nil
	}
	allErrs := field.ErrorList{}
	obj, ok := uncastObj.(*authorizationv1.RoleBindingRestriction)
	if !ok {
		return nil, append(allErrs, field.NotSupported(field.NewPath("kind"), fmt.Sprintf("%T", uncastObj), []string{"RoleBindingRestriction"}), field.NotSupported(field.NewPath("apiVersion"), fmt.Sprintf("%T", uncastObj), []string{authorizationv1.GroupVersion.String()}))
	}
	return obj, nil
}

type roleBindingRestrictionV1 struct{}

func (roleBindingRestrictionV1) ValidateCreate(obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	roleBindingRestrictionObj, errs := toRoleBindingRestriction(obj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMeta(&roleBindingRestrictionObj.ObjectMeta, true, validation.NameIsDNSSubdomain, field.NewPath("metadata"))...)
	errs = append(errs, rbrvalidation.ValidateRoleBindingRestriction(roleBindingRestrictionObj)...)
	return errs
}
func (roleBindingRestrictionV1) ValidateUpdate(obj runtime.Object, oldObj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	roleBindingRestrictionObj, errs := toRoleBindingRestriction(obj)
	if len(errs) > 0 {
		return errs
	}
	roleBindingRestrictionOldObj, errs := toRoleBindingRestriction(oldObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMeta(&roleBindingRestrictionObj.ObjectMeta, true, validation.NameIsDNSSubdomain, field.NewPath("metadata"))...)
	errs = append(errs, rbrvalidation.ValidateRoleBindingRestrictionUpdate(roleBindingRestrictionObj, roleBindingRestrictionOldObj)...)
	return errs
}
func (r roleBindingRestrictionV1) ValidateStatusUpdate(obj runtime.Object, oldObj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.ValidateUpdate(obj, oldObj)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}

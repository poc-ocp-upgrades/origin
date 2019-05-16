package securitycontextconstraints

import (
	"fmt"
	securityv1 "github.com/openshift/api/security/v1"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation"
	sccvalidation "github.com/openshift/origin/pkg/admission/customresourcevalidation/securitycontextconstraints/validation"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/admission"
)

const PluginName = "security.openshift.io/ValidateSecurityContextConstraints"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return customresourcevalidation.NewValidator(map[schema.GroupResource]bool{{Group: securityv1.GroupName, Resource: "securitycontextconstraints"}: true}, map[schema.GroupVersionKind]customresourcevalidation.ObjectValidator{securityv1.GroupVersion.WithKind("SecurityContextConstraints"): securityContextConstraintsV1{}})
	})
}
func toSecurityContextConstraints(uncastObj runtime.Object) (*securityv1.SecurityContextConstraints, field.ErrorList) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if uncastObj == nil {
		return nil, nil
	}
	obj, ok := uncastObj.(*securityv1.SecurityContextConstraints)
	if !ok {
		return nil, field.ErrorList{field.NotSupported(field.NewPath("kind"), fmt.Sprintf("%T", uncastObj), []string{"SecurityContextConstraints"}), field.NotSupported(field.NewPath("apiVersion"), fmt.Sprintf("%T", uncastObj), []string{securityv1.GroupVersion.String()})}
	}
	return obj, nil
}

type securityContextConstraintsV1 struct{}

func (securityContextConstraintsV1) ValidateCreate(obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	securityContextConstraintsObj, errs := toSecurityContextConstraints(obj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, sccvalidation.ValidateSecurityContextConstraints(securityContextConstraintsObj)...)
	return errs
}
func (securityContextConstraintsV1) ValidateUpdate(obj runtime.Object, oldObj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	securityContextConstraintsObj, errs := toSecurityContextConstraints(obj)
	if len(errs) > 0 {
		return errs
	}
	securityContextConstraintsOldObj, errs := toSecurityContextConstraints(oldObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, sccvalidation.ValidateSecurityContextConstraintsUpdate(securityContextConstraintsObj, securityContextConstraintsOldObj)...)
	return errs
}
func (c securityContextConstraintsV1) ValidateStatusUpdate(obj runtime.Object, oldObj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.ValidateUpdate(obj, oldObj)
}

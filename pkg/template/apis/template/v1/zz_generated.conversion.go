package v1

import (
	unsafe "unsafe"
	v1 "github.com/openshift/api/template/v1"
	template "github.com/openshift/origin/pkg/template/apis/template"
	apicorev1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func init() {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*v1.BrokerTemplateInstance)(nil), (*template.BrokerTemplateInstance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance(a.(*v1.BrokerTemplateInstance), b.(*template.BrokerTemplateInstance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.BrokerTemplateInstance)(nil), (*v1.BrokerTemplateInstance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(a.(*template.BrokerTemplateInstance), b.(*v1.BrokerTemplateInstance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BrokerTemplateInstanceList)(nil), (*template.BrokerTemplateInstanceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BrokerTemplateInstanceList_To_template_BrokerTemplateInstanceList(a.(*v1.BrokerTemplateInstanceList), b.(*template.BrokerTemplateInstanceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.BrokerTemplateInstanceList)(nil), (*v1.BrokerTemplateInstanceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList(a.(*template.BrokerTemplateInstanceList), b.(*v1.BrokerTemplateInstanceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BrokerTemplateInstanceSpec)(nil), (*template.BrokerTemplateInstanceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec(a.(*v1.BrokerTemplateInstanceSpec), b.(*template.BrokerTemplateInstanceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.BrokerTemplateInstanceSpec)(nil), (*v1.BrokerTemplateInstanceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(a.(*template.BrokerTemplateInstanceSpec), b.(*v1.BrokerTemplateInstanceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.Parameter)(nil), (*template.Parameter)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Parameter_To_template_Parameter(a.(*v1.Parameter), b.(*template.Parameter), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.Parameter)(nil), (*v1.Parameter)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_Parameter_To_v1_Parameter(a.(*template.Parameter), b.(*v1.Parameter), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.Template)(nil), (*template.Template)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Template_To_template_Template(a.(*v1.Template), b.(*template.Template), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.Template)(nil), (*v1.Template)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_Template_To_v1_Template(a.(*template.Template), b.(*v1.Template), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TemplateInstance)(nil), (*template.TemplateInstance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TemplateInstance_To_template_TemplateInstance(a.(*v1.TemplateInstance), b.(*template.TemplateInstance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.TemplateInstance)(nil), (*v1.TemplateInstance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_TemplateInstance_To_v1_TemplateInstance(a.(*template.TemplateInstance), b.(*v1.TemplateInstance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TemplateInstanceCondition)(nil), (*template.TemplateInstanceCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TemplateInstanceCondition_To_template_TemplateInstanceCondition(a.(*v1.TemplateInstanceCondition), b.(*template.TemplateInstanceCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.TemplateInstanceCondition)(nil), (*v1.TemplateInstanceCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_TemplateInstanceCondition_To_v1_TemplateInstanceCondition(a.(*template.TemplateInstanceCondition), b.(*v1.TemplateInstanceCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TemplateInstanceList)(nil), (*template.TemplateInstanceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TemplateInstanceList_To_template_TemplateInstanceList(a.(*v1.TemplateInstanceList), b.(*template.TemplateInstanceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.TemplateInstanceList)(nil), (*v1.TemplateInstanceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_TemplateInstanceList_To_v1_TemplateInstanceList(a.(*template.TemplateInstanceList), b.(*v1.TemplateInstanceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TemplateInstanceObject)(nil), (*template.TemplateInstanceObject)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject(a.(*v1.TemplateInstanceObject), b.(*template.TemplateInstanceObject), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.TemplateInstanceObject)(nil), (*v1.TemplateInstanceObject)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject(a.(*template.TemplateInstanceObject), b.(*v1.TemplateInstanceObject), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TemplateInstanceRequester)(nil), (*template.TemplateInstanceRequester)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TemplateInstanceRequester_To_template_TemplateInstanceRequester(a.(*v1.TemplateInstanceRequester), b.(*template.TemplateInstanceRequester), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.TemplateInstanceRequester)(nil), (*v1.TemplateInstanceRequester)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_TemplateInstanceRequester_To_v1_TemplateInstanceRequester(a.(*template.TemplateInstanceRequester), b.(*v1.TemplateInstanceRequester), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TemplateInstanceSpec)(nil), (*template.TemplateInstanceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec(a.(*v1.TemplateInstanceSpec), b.(*template.TemplateInstanceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.TemplateInstanceSpec)(nil), (*v1.TemplateInstanceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(a.(*template.TemplateInstanceSpec), b.(*v1.TemplateInstanceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TemplateInstanceStatus)(nil), (*template.TemplateInstanceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus(a.(*v1.TemplateInstanceStatus), b.(*template.TemplateInstanceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.TemplateInstanceStatus)(nil), (*v1.TemplateInstanceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(a.(*template.TemplateInstanceStatus), b.(*v1.TemplateInstanceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TemplateList)(nil), (*template.TemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TemplateList_To_template_TemplateList(a.(*v1.TemplateList), b.(*template.TemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*template.TemplateList)(nil), (*v1.TemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_template_TemplateList_To_v1_TemplateList(a.(*template.TemplateList), b.(*v1.TemplateList), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance(in *v1.BrokerTemplateInstance, out *template.BrokerTemplateInstance, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance(in *v1.BrokerTemplateInstance, out *template.BrokerTemplateInstance, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance(in, out, s)
}
func autoConvert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(in *template.BrokerTemplateInstance, out *v1.BrokerTemplateInstance, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}
func Convert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(in *template.BrokerTemplateInstance, out *v1.BrokerTemplateInstance, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(in, out, s)
}
func autoConvert_v1_BrokerTemplateInstanceList_To_template_BrokerTemplateInstanceList(in *v1.BrokerTemplateInstanceList, out *template.BrokerTemplateInstanceList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]template.BrokerTemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_BrokerTemplateInstanceList_To_template_BrokerTemplateInstanceList(in *v1.BrokerTemplateInstanceList, out *template.BrokerTemplateInstanceList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BrokerTemplateInstanceList_To_template_BrokerTemplateInstanceList(in, out, s)
}
func autoConvert_template_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList(in *template.BrokerTemplateInstanceList, out *v1.BrokerTemplateInstanceList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.BrokerTemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_template_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList(in *template.BrokerTemplateInstanceList, out *v1.BrokerTemplateInstanceList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList(in, out, s)
}
func autoConvert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec(in *v1.BrokerTemplateInstanceSpec, out *template.BrokerTemplateInstanceSpec, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.TemplateInstance, &out.TemplateInstance, s); err != nil {
		return err
	}
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.Secret, &out.Secret, s); err != nil {
		return err
	}
	out.BindingIDs = *(*[]string)(unsafe.Pointer(&in.BindingIDs))
	return nil
}
func Convert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec(in *v1.BrokerTemplateInstanceSpec, out *template.BrokerTemplateInstanceSpec, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec(in, out, s)
}
func autoConvert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(in *template.BrokerTemplateInstanceSpec, out *v1.BrokerTemplateInstanceSpec, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.TemplateInstance, &out.TemplateInstance, s); err != nil {
		return err
	}
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.Secret, &out.Secret, s); err != nil {
		return err
	}
	out.BindingIDs = *(*[]string)(unsafe.Pointer(&in.BindingIDs))
	return nil
}
func Convert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(in *template.BrokerTemplateInstanceSpec, out *v1.BrokerTemplateInstanceSpec, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(in, out, s)
}
func autoConvert_v1_Parameter_To_template_Parameter(in *v1.Parameter, out *template.Parameter, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	out.Required = in.Required
	return nil
}
func Convert_v1_Parameter_To_template_Parameter(in *v1.Parameter, out *template.Parameter, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_Parameter_To_template_Parameter(in, out, s)
}
func autoConvert_template_Parameter_To_v1_Parameter(in *template.Parameter, out *v1.Parameter, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	out.Required = in.Required
	return nil
}
func Convert_template_Parameter_To_v1_Parameter(in *template.Parameter, out *v1.Parameter, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_Parameter_To_v1_Parameter(in, out, s)
}
func autoConvert_v1_Template_To_template_Template(in *v1.Template, out *template.Template, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Message = in.Message
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]runtime.Object, len(*in))
		for i := range *in {
			if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Objects = nil
	}
	out.Parameters = *(*[]template.Parameter)(unsafe.Pointer(&in.Parameters))
	out.ObjectLabels = *(*map[string]string)(unsafe.Pointer(&in.ObjectLabels))
	return nil
}
func Convert_v1_Template_To_template_Template(in *v1.Template, out *template.Template, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_Template_To_template_Template(in, out, s)
}
func autoConvert_template_Template_To_v1_Template(in *template.Template, out *v1.Template, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Message = in.Message
	out.Parameters = *(*[]v1.Parameter)(unsafe.Pointer(&in.Parameters))
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Objects = nil
	}
	out.ObjectLabels = *(*map[string]string)(unsafe.Pointer(&in.ObjectLabels))
	return nil
}
func Convert_template_Template_To_v1_Template(in *template.Template, out *v1.Template, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_Template_To_v1_Template(in, out, s)
}
func autoConvert_v1_TemplateInstance_To_template_TemplateInstance(in *v1.TemplateInstance, out *template.TemplateInstance, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_TemplateInstance_To_template_TemplateInstance(in *v1.TemplateInstance, out *template.TemplateInstance, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TemplateInstance_To_template_TemplateInstance(in, out, s)
}
func autoConvert_template_TemplateInstance_To_v1_TemplateInstance(in *template.TemplateInstance, out *v1.TemplateInstance, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_template_TemplateInstance_To_v1_TemplateInstance(in *template.TemplateInstance, out *v1.TemplateInstance, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_TemplateInstance_To_v1_TemplateInstance(in, out, s)
}
func autoConvert_v1_TemplateInstanceCondition_To_template_TemplateInstanceCondition(in *v1.TemplateInstanceCondition, out *template.TemplateInstanceCondition, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = template.TemplateInstanceConditionType(in.Type)
	out.Status = core.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}
func Convert_v1_TemplateInstanceCondition_To_template_TemplateInstanceCondition(in *v1.TemplateInstanceCondition, out *template.TemplateInstanceCondition, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TemplateInstanceCondition_To_template_TemplateInstanceCondition(in, out, s)
}
func autoConvert_template_TemplateInstanceCondition_To_v1_TemplateInstanceCondition(in *template.TemplateInstanceCondition, out *v1.TemplateInstanceCondition, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.TemplateInstanceConditionType(in.Type)
	out.Status = apicorev1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}
func Convert_template_TemplateInstanceCondition_To_v1_TemplateInstanceCondition(in *template.TemplateInstanceCondition, out *v1.TemplateInstanceCondition, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_TemplateInstanceCondition_To_v1_TemplateInstanceCondition(in, out, s)
}
func autoConvert_v1_TemplateInstanceList_To_template_TemplateInstanceList(in *v1.TemplateInstanceList, out *template.TemplateInstanceList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]template.TemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_v1_TemplateInstance_To_template_TemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_TemplateInstanceList_To_template_TemplateInstanceList(in *v1.TemplateInstanceList, out *template.TemplateInstanceList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TemplateInstanceList_To_template_TemplateInstanceList(in, out, s)
}
func autoConvert_template_TemplateInstanceList_To_v1_TemplateInstanceList(in *template.TemplateInstanceList, out *v1.TemplateInstanceList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.TemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_template_TemplateInstance_To_v1_TemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_template_TemplateInstanceList_To_v1_TemplateInstanceList(in *template.TemplateInstanceList, out *v1.TemplateInstanceList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_TemplateInstanceList_To_v1_TemplateInstanceList(in, out, s)
}
func autoConvert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject(in *v1.TemplateInstanceObject, out *template.TemplateInstanceObject, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.Ref, &out.Ref, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject(in *v1.TemplateInstanceObject, out *template.TemplateInstanceObject, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject(in, out, s)
}
func autoConvert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject(in *template.TemplateInstanceObject, out *v1.TemplateInstanceObject, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.Ref, &out.Ref, s); err != nil {
		return err
	}
	return nil
}
func Convert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject(in *template.TemplateInstanceObject, out *v1.TemplateInstanceObject, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject(in, out, s)
}
func autoConvert_v1_TemplateInstanceRequester_To_template_TemplateInstanceRequester(in *v1.TemplateInstanceRequester, out *template.TemplateInstanceRequester, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]template.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}
func Convert_v1_TemplateInstanceRequester_To_template_TemplateInstanceRequester(in *v1.TemplateInstanceRequester, out *template.TemplateInstanceRequester, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TemplateInstanceRequester_To_template_TemplateInstanceRequester(in, out, s)
}
func autoConvert_template_TemplateInstanceRequester_To_v1_TemplateInstanceRequester(in *template.TemplateInstanceRequester, out *v1.TemplateInstanceRequester, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]v1.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}
func Convert_template_TemplateInstanceRequester_To_v1_TemplateInstanceRequester(in *template.TemplateInstanceRequester, out *v1.TemplateInstanceRequester, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_TemplateInstanceRequester_To_v1_TemplateInstanceRequester(in, out, s)
}
func autoConvert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec(in *v1.TemplateInstanceSpec, out *template.TemplateInstanceSpec, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_Template_To_template_Template(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(core.LocalObjectReference)
		if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Secret = nil
	}
	out.Requester = (*template.TemplateInstanceRequester)(unsafe.Pointer(in.Requester))
	return nil
}
func Convert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec(in *v1.TemplateInstanceSpec, out *template.TemplateInstanceSpec, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec(in, out, s)
}
func autoConvert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(in *template.TemplateInstanceSpec, out *v1.TemplateInstanceSpec, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_template_Template_To_v1_Template(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(apicorev1.LocalObjectReference)
		if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Secret = nil
	}
	out.Requester = (*v1.TemplateInstanceRequester)(unsafe.Pointer(in.Requester))
	return nil
}
func Convert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(in *template.TemplateInstanceSpec, out *v1.TemplateInstanceSpec, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(in, out, s)
}
func autoConvert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus(in *v1.TemplateInstanceStatus, out *template.TemplateInstanceStatus, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Conditions = *(*[]template.TemplateInstanceCondition)(unsafe.Pointer(&in.Conditions))
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]template.TemplateInstanceObject, len(*in))
		for i := range *in {
			if err := Convert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Objects = nil
	}
	return nil
}
func Convert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus(in *v1.TemplateInstanceStatus, out *template.TemplateInstanceStatus, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus(in, out, s)
}
func autoConvert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(in *template.TemplateInstanceStatus, out *v1.TemplateInstanceStatus, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Conditions = *(*[]v1.TemplateInstanceCondition)(unsafe.Pointer(&in.Conditions))
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]v1.TemplateInstanceObject, len(*in))
		for i := range *in {
			if err := Convert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Objects = nil
	}
	return nil
}
func Convert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(in *template.TemplateInstanceStatus, out *v1.TemplateInstanceStatus, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(in, out, s)
}
func autoConvert_v1_TemplateList_To_template_TemplateList(in *v1.TemplateList, out *template.TemplateList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]template.Template, len(*in))
		for i := range *in {
			if err := Convert_v1_Template_To_template_Template(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_TemplateList_To_template_TemplateList(in *v1.TemplateList, out *template.TemplateList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TemplateList_To_template_TemplateList(in, out, s)
}
func autoConvert_template_TemplateList_To_v1_TemplateList(in *template.TemplateList, out *v1.TemplateList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.Template, len(*in))
		for i := range *in {
			if err := Convert_template_Template_To_v1_Template(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_template_TemplateList_To_v1_TemplateList(in *template.TemplateList, out *v1.TemplateList, s conversion.Scope) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_template_TemplateList_To_v1_TemplateList(in, out, s)
}

package template

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
)

func (in *BrokerTemplateInstance) DeepCopyInto(out *BrokerTemplateInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}
func (in *BrokerTemplateInstance) DeepCopy() *BrokerTemplateInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BrokerTemplateInstance)
	in.DeepCopyInto(out)
	return out
}
func (in *BrokerTemplateInstance) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BrokerTemplateInstanceList) DeepCopyInto(out *BrokerTemplateInstanceList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BrokerTemplateInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *BrokerTemplateInstanceList) DeepCopy() *BrokerTemplateInstanceList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BrokerTemplateInstanceList)
	in.DeepCopyInto(out)
	return out
}
func (in *BrokerTemplateInstanceList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BrokerTemplateInstanceSpec) DeepCopyInto(out *BrokerTemplateInstanceSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TemplateInstance = in.TemplateInstance
	out.Secret = in.Secret
	if in.BindingIDs != nil {
		in, out := &in.BindingIDs, &out.BindingIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *BrokerTemplateInstanceSpec) DeepCopy() *BrokerTemplateInstanceSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BrokerTemplateInstanceSpec)
	in.DeepCopyInto(out)
	return out
}
func (in ExtraValue) DeepCopyInto(out *ExtraValue) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	{
		in := &in
		*out = make(ExtraValue, len(*in))
		copy(*out, *in)
		return
	}
}
func (in ExtraValue) DeepCopy() ExtraValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ExtraValue)
	in.DeepCopyInto(out)
	return *out
}
func (in *Parameter) DeepCopyInto(out *Parameter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *Parameter) DeepCopy() *Parameter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}
func (in *Template) DeepCopyInto(out *Template) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]Parameter, len(*in))
		copy(*out, *in)
	}
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]runtime.Object, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				(*out)[i] = (*in)[i].DeepCopyObject()
			}
		}
	}
	if in.ObjectLabels != nil {
		in, out := &in.ObjectLabels, &out.ObjectLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}
func (in *Template) DeepCopy() *Template {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}
func (in *Template) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *TemplateInstance) DeepCopyInto(out *TemplateInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *TemplateInstance) DeepCopy() *TemplateInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TemplateInstance)
	in.DeepCopyInto(out)
	return out
}
func (in *TemplateInstance) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *TemplateInstanceCondition) DeepCopyInto(out *TemplateInstanceCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *TemplateInstanceCondition) DeepCopy() *TemplateInstanceCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TemplateInstanceCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *TemplateInstanceList) DeepCopyInto(out *TemplateInstanceList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemplateInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *TemplateInstanceList) DeepCopy() *TemplateInstanceList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TemplateInstanceList)
	in.DeepCopyInto(out)
	return out
}
func (in *TemplateInstanceList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *TemplateInstanceObject) DeepCopyInto(out *TemplateInstanceObject) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.Ref = in.Ref
	return
}
func (in *TemplateInstanceObject) DeepCopy() *TemplateInstanceObject {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TemplateInstanceObject)
	in.DeepCopyInto(out)
	return out
}
func (in *TemplateInstanceRequester) DeepCopyInto(out *TemplateInstanceRequester) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make(map[string]ExtraValue, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ExtraValue, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}
func (in *TemplateInstanceRequester) DeepCopy() *TemplateInstanceRequester {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TemplateInstanceRequester)
	in.DeepCopyInto(out)
	return out
}
func (in *TemplateInstanceSpec) DeepCopyInto(out *TemplateInstanceSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(core.LocalObjectReference)
		**out = **in
	}
	if in.Requester != nil {
		in, out := &in.Requester, &out.Requester
		*out = new(TemplateInstanceRequester)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *TemplateInstanceSpec) DeepCopy() *TemplateInstanceSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TemplateInstanceSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *TemplateInstanceStatus) DeepCopyInto(out *TemplateInstanceStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TemplateInstanceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]TemplateInstanceObject, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *TemplateInstanceStatus) DeepCopy() *TemplateInstanceStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TemplateInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *TemplateList) DeepCopyInto(out *TemplateList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Template, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *TemplateList) DeepCopy() *TemplateList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TemplateList)
	in.DeepCopyInto(out)
	return out
}
func (in *TemplateList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

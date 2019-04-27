package v1

import (
	unsafe "unsafe"
	v1 "github.com/openshift/api/project/v1"
	project "github.com/openshift/origin/pkg/project/apis/project"
	corev1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
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
	if err := s.AddGeneratedConversionFunc((*v1.Project)(nil), (*project.Project)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Project_To_project_Project(a.(*v1.Project), b.(*project.Project), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*project.Project)(nil), (*v1.Project)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_project_Project_To_v1_Project(a.(*project.Project), b.(*v1.Project), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ProjectList)(nil), (*project.ProjectList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ProjectList_To_project_ProjectList(a.(*v1.ProjectList), b.(*project.ProjectList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*project.ProjectList)(nil), (*v1.ProjectList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_project_ProjectList_To_v1_ProjectList(a.(*project.ProjectList), b.(*v1.ProjectList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ProjectRequest)(nil), (*project.ProjectRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ProjectRequest_To_project_ProjectRequest(a.(*v1.ProjectRequest), b.(*project.ProjectRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*project.ProjectRequest)(nil), (*v1.ProjectRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_project_ProjectRequest_To_v1_ProjectRequest(a.(*project.ProjectRequest), b.(*v1.ProjectRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ProjectSpec)(nil), (*project.ProjectSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ProjectSpec_To_project_ProjectSpec(a.(*v1.ProjectSpec), b.(*project.ProjectSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*project.ProjectSpec)(nil), (*v1.ProjectSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_project_ProjectSpec_To_v1_ProjectSpec(a.(*project.ProjectSpec), b.(*v1.ProjectSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ProjectStatus)(nil), (*project.ProjectStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ProjectStatus_To_project_ProjectStatus(a.(*v1.ProjectStatus), b.(*project.ProjectStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*project.ProjectStatus)(nil), (*v1.ProjectStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_project_ProjectStatus_To_v1_ProjectStatus(a.(*project.ProjectStatus), b.(*v1.ProjectStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_Project_To_project_Project(in *v1.Project, out *project.Project, s conversion.Scope) error {
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
	if err := Convert_v1_ProjectSpec_To_project_ProjectSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ProjectStatus_To_project_ProjectStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_Project_To_project_Project(in *v1.Project, out *project.Project, s conversion.Scope) error {
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
	return autoConvert_v1_Project_To_project_Project(in, out, s)
}
func autoConvert_project_Project_To_v1_Project(in *project.Project, out *v1.Project, s conversion.Scope) error {
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
	if err := Convert_project_ProjectSpec_To_v1_ProjectSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_project_ProjectStatus_To_v1_ProjectStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_project_Project_To_v1_Project(in *project.Project, out *v1.Project, s conversion.Scope) error {
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
	return autoConvert_project_Project_To_v1_Project(in, out, s)
}
func autoConvert_v1_ProjectList_To_project_ProjectList(in *v1.ProjectList, out *project.ProjectList, s conversion.Scope) error {
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
	out.Items = *(*[]project.Project)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_ProjectList_To_project_ProjectList(in *v1.ProjectList, out *project.ProjectList, s conversion.Scope) error {
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
	return autoConvert_v1_ProjectList_To_project_ProjectList(in, out, s)
}
func autoConvert_project_ProjectList_To_v1_ProjectList(in *project.ProjectList, out *v1.ProjectList, s conversion.Scope) error {
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
	out.Items = *(*[]v1.Project)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_project_ProjectList_To_v1_ProjectList(in *project.ProjectList, out *v1.ProjectList, s conversion.Scope) error {
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
	return autoConvert_project_ProjectList_To_v1_ProjectList(in, out, s)
}
func autoConvert_v1_ProjectRequest_To_project_ProjectRequest(in *v1.ProjectRequest, out *project.ProjectRequest, s conversion.Scope) error {
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
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}
func Convert_v1_ProjectRequest_To_project_ProjectRequest(in *v1.ProjectRequest, out *project.ProjectRequest, s conversion.Scope) error {
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
	return autoConvert_v1_ProjectRequest_To_project_ProjectRequest(in, out, s)
}
func autoConvert_project_ProjectRequest_To_v1_ProjectRequest(in *project.ProjectRequest, out *v1.ProjectRequest, s conversion.Scope) error {
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
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}
func Convert_project_ProjectRequest_To_v1_ProjectRequest(in *project.ProjectRequest, out *v1.ProjectRequest, s conversion.Scope) error {
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
	return autoConvert_project_ProjectRequest_To_v1_ProjectRequest(in, out, s)
}
func autoConvert_v1_ProjectSpec_To_project_ProjectSpec(in *v1.ProjectSpec, out *project.ProjectSpec, s conversion.Scope) error {
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
	out.Finalizers = *(*[]core.FinalizerName)(unsafe.Pointer(&in.Finalizers))
	return nil
}
func Convert_v1_ProjectSpec_To_project_ProjectSpec(in *v1.ProjectSpec, out *project.ProjectSpec, s conversion.Scope) error {
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
	return autoConvert_v1_ProjectSpec_To_project_ProjectSpec(in, out, s)
}
func autoConvert_project_ProjectSpec_To_v1_ProjectSpec(in *project.ProjectSpec, out *v1.ProjectSpec, s conversion.Scope) error {
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
	out.Finalizers = *(*[]corev1.FinalizerName)(unsafe.Pointer(&in.Finalizers))
	return nil
}
func Convert_project_ProjectSpec_To_v1_ProjectSpec(in *project.ProjectSpec, out *v1.ProjectSpec, s conversion.Scope) error {
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
	return autoConvert_project_ProjectSpec_To_v1_ProjectSpec(in, out, s)
}
func autoConvert_v1_ProjectStatus_To_project_ProjectStatus(in *v1.ProjectStatus, out *project.ProjectStatus, s conversion.Scope) error {
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
	out.Phase = core.NamespacePhase(in.Phase)
	return nil
}
func Convert_v1_ProjectStatus_To_project_ProjectStatus(in *v1.ProjectStatus, out *project.ProjectStatus, s conversion.Scope) error {
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
	return autoConvert_v1_ProjectStatus_To_project_ProjectStatus(in, out, s)
}
func autoConvert_project_ProjectStatus_To_v1_ProjectStatus(in *project.ProjectStatus, out *v1.ProjectStatus, s conversion.Scope) error {
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
	out.Phase = corev1.NamespacePhase(in.Phase)
	return nil
}
func Convert_project_ProjectStatus_To_v1_ProjectStatus(in *project.ProjectStatus, out *v1.ProjectStatus, s conversion.Scope) error {
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
	return autoConvert_project_ProjectStatus_To_v1_ProjectStatus(in, out, s)
}

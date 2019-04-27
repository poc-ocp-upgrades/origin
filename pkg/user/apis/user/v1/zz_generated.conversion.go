package v1

import (
	unsafe "unsafe"
	v1 "github.com/openshift/api/user/v1"
	user "github.com/openshift/origin/pkg/user/apis/user"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
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
	if err := s.AddGeneratedConversionFunc((*v1.Group)(nil), (*user.Group)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Group_To_user_Group(a.(*v1.Group), b.(*user.Group), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*user.Group)(nil), (*v1.Group)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_user_Group_To_v1_Group(a.(*user.Group), b.(*v1.Group), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GroupList)(nil), (*user.GroupList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GroupList_To_user_GroupList(a.(*v1.GroupList), b.(*user.GroupList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*user.GroupList)(nil), (*v1.GroupList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_user_GroupList_To_v1_GroupList(a.(*user.GroupList), b.(*v1.GroupList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.Identity)(nil), (*user.Identity)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Identity_To_user_Identity(a.(*v1.Identity), b.(*user.Identity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*user.Identity)(nil), (*v1.Identity)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_user_Identity_To_v1_Identity(a.(*user.Identity), b.(*v1.Identity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.IdentityList)(nil), (*user.IdentityList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_IdentityList_To_user_IdentityList(a.(*v1.IdentityList), b.(*user.IdentityList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*user.IdentityList)(nil), (*v1.IdentityList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_user_IdentityList_To_v1_IdentityList(a.(*user.IdentityList), b.(*v1.IdentityList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.User)(nil), (*user.User)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_User_To_user_User(a.(*v1.User), b.(*user.User), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*user.User)(nil), (*v1.User)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_user_User_To_v1_User(a.(*user.User), b.(*v1.User), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.UserIdentityMapping)(nil), (*user.UserIdentityMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_UserIdentityMapping_To_user_UserIdentityMapping(a.(*v1.UserIdentityMapping), b.(*user.UserIdentityMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*user.UserIdentityMapping)(nil), (*v1.UserIdentityMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_user_UserIdentityMapping_To_v1_UserIdentityMapping(a.(*user.UserIdentityMapping), b.(*v1.UserIdentityMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.UserList)(nil), (*user.UserList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_UserList_To_user_UserList(a.(*v1.UserList), b.(*user.UserList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*user.UserList)(nil), (*v1.UserList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_user_UserList_To_v1_UserList(a.(*user.UserList), b.(*v1.UserList), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_Group_To_user_Group(in *v1.Group, out *user.Group, s conversion.Scope) error {
	_logClusterCodePath()
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
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	return nil
}
func Convert_v1_Group_To_user_Group(in *v1.Group, out *user.Group, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_v1_Group_To_user_Group(in, out, s)
}
func autoConvert_user_Group_To_v1_Group(in *user.Group, out *v1.Group, s conversion.Scope) error {
	_logClusterCodePath()
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
	out.Users = *(*v1.OptionalNames)(unsafe.Pointer(&in.Users))
	return nil
}
func Convert_user_Group_To_v1_Group(in *user.Group, out *v1.Group, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_user_Group_To_v1_Group(in, out, s)
}
func autoConvert_v1_GroupList_To_user_GroupList(in *v1.GroupList, out *user.GroupList, s conversion.Scope) error {
	_logClusterCodePath()
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
	out.Items = *(*[]user.Group)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_GroupList_To_user_GroupList(in *v1.GroupList, out *user.GroupList, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_v1_GroupList_To_user_GroupList(in, out, s)
}
func autoConvert_user_GroupList_To_v1_GroupList(in *user.GroupList, out *v1.GroupList, s conversion.Scope) error {
	_logClusterCodePath()
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
	out.Items = *(*[]v1.Group)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_user_GroupList_To_v1_GroupList(in *user.GroupList, out *v1.GroupList, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_user_GroupList_To_v1_GroupList(in, out, s)
}
func autoConvert_v1_Identity_To_user_Identity(in *v1.Identity, out *user.Identity, s conversion.Scope) error {
	_logClusterCodePath()
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
	out.ProviderName = in.ProviderName
	out.ProviderUserName = in.ProviderUserName
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Extra = *(*map[string]string)(unsafe.Pointer(&in.Extra))
	return nil
}
func Convert_v1_Identity_To_user_Identity(in *v1.Identity, out *user.Identity, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_v1_Identity_To_user_Identity(in, out, s)
}
func autoConvert_user_Identity_To_v1_Identity(in *user.Identity, out *v1.Identity, s conversion.Scope) error {
	_logClusterCodePath()
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
	out.ProviderName = in.ProviderName
	out.ProviderUserName = in.ProviderUserName
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Extra = *(*map[string]string)(unsafe.Pointer(&in.Extra))
	return nil
}
func Convert_user_Identity_To_v1_Identity(in *user.Identity, out *v1.Identity, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_user_Identity_To_v1_Identity(in, out, s)
}
func autoConvert_v1_IdentityList_To_user_IdentityList(in *v1.IdentityList, out *user.IdentityList, s conversion.Scope) error {
	_logClusterCodePath()
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
		*out = make([]user.Identity, len(*in))
		for i := range *in {
			if err := Convert_v1_Identity_To_user_Identity(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_IdentityList_To_user_IdentityList(in *v1.IdentityList, out *user.IdentityList, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_v1_IdentityList_To_user_IdentityList(in, out, s)
}
func autoConvert_user_IdentityList_To_v1_IdentityList(in *user.IdentityList, out *v1.IdentityList, s conversion.Scope) error {
	_logClusterCodePath()
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
		*out = make([]v1.Identity, len(*in))
		for i := range *in {
			if err := Convert_user_Identity_To_v1_Identity(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_user_IdentityList_To_v1_IdentityList(in *user.IdentityList, out *v1.IdentityList, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_user_IdentityList_To_v1_IdentityList(in, out, s)
}
func autoConvert_v1_User_To_user_User(in *v1.User, out *user.User, s conversion.Scope) error {
	_logClusterCodePath()
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
	out.FullName = in.FullName
	out.Identities = *(*[]string)(unsafe.Pointer(&in.Identities))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}
func Convert_v1_User_To_user_User(in *v1.User, out *user.User, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_v1_User_To_user_User(in, out, s)
}
func autoConvert_user_User_To_v1_User(in *user.User, out *v1.User, s conversion.Scope) error {
	_logClusterCodePath()
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
	out.FullName = in.FullName
	out.Identities = *(*[]string)(unsafe.Pointer(&in.Identities))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}
func Convert_user_User_To_v1_User(in *user.User, out *v1.User, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_user_User_To_v1_User(in, out, s)
}
func autoConvert_v1_UserIdentityMapping_To_user_UserIdentityMapping(in *v1.UserIdentityMapping, out *user.UserIdentityMapping, s conversion.Scope) error {
	_logClusterCodePath()
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
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.Identity, &out.Identity, s); err != nil {
		return err
	}
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_UserIdentityMapping_To_user_UserIdentityMapping(in *v1.UserIdentityMapping, out *user.UserIdentityMapping, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_v1_UserIdentityMapping_To_user_UserIdentityMapping(in, out, s)
}
func autoConvert_user_UserIdentityMapping_To_v1_UserIdentityMapping(in *user.UserIdentityMapping, out *v1.UserIdentityMapping, s conversion.Scope) error {
	_logClusterCodePath()
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
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.Identity, &out.Identity, s); err != nil {
		return err
	}
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	return nil
}
func Convert_user_UserIdentityMapping_To_v1_UserIdentityMapping(in *user.UserIdentityMapping, out *v1.UserIdentityMapping, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_user_UserIdentityMapping_To_v1_UserIdentityMapping(in, out, s)
}
func autoConvert_v1_UserList_To_user_UserList(in *v1.UserList, out *user.UserList, s conversion.Scope) error {
	_logClusterCodePath()
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
	out.Items = *(*[]user.User)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_UserList_To_user_UserList(in *v1.UserList, out *user.UserList, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_v1_UserList_To_user_UserList(in, out, s)
}
func autoConvert_user_UserList_To_v1_UserList(in *user.UserList, out *v1.UserList, s conversion.Scope) error {
	_logClusterCodePath()
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
	out.Items = *(*[]v1.User)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_user_UserList_To_v1_UserList(in *user.UserList, out *v1.UserList, s conversion.Scope) error {
	_logClusterCodePath()
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
	return autoConvert_user_UserList_To_v1_UserList(in, out, s)
}

package oauth

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *ClusterRoleScopeRestriction) DeepCopyInto(out *ClusterRoleScopeRestriction) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.RoleNames != nil {
		in, out := &in.RoleNames, &out.RoleNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *ClusterRoleScopeRestriction) DeepCopy() *ClusterRoleScopeRestriction {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterRoleScopeRestriction)
	in.DeepCopyInto(out)
	return out
}
func (in *OAuthAccessToken) DeepCopyInto(out *OAuthAccessToken) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *OAuthAccessToken) DeepCopy() *OAuthAccessToken {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OAuthAccessToken)
	in.DeepCopyInto(out)
	return out
}
func (in *OAuthAccessToken) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *OAuthAccessTokenList) DeepCopyInto(out *OAuthAccessTokenList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthAccessToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *OAuthAccessTokenList) DeepCopy() *OAuthAccessTokenList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OAuthAccessTokenList)
	in.DeepCopyInto(out)
	return out
}
func (in *OAuthAccessTokenList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *OAuthAuthorizeToken) DeepCopyInto(out *OAuthAuthorizeToken) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *OAuthAuthorizeToken) DeepCopy() *OAuthAuthorizeToken {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OAuthAuthorizeToken)
	in.DeepCopyInto(out)
	return out
}
func (in *OAuthAuthorizeToken) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *OAuthAuthorizeTokenList) DeepCopyInto(out *OAuthAuthorizeTokenList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthAuthorizeToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *OAuthAuthorizeTokenList) DeepCopy() *OAuthAuthorizeTokenList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OAuthAuthorizeTokenList)
	in.DeepCopyInto(out)
	return out
}
func (in *OAuthAuthorizeTokenList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *OAuthClient) DeepCopyInto(out *OAuthClient) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AdditionalSecrets != nil {
		in, out := &in.AdditionalSecrets, &out.AdditionalSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RedirectURIs != nil {
		in, out := &in.RedirectURIs, &out.RedirectURIs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ScopeRestrictions != nil {
		in, out := &in.ScopeRestrictions, &out.ScopeRestrictions
		*out = make([]ScopeRestriction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AccessTokenMaxAgeSeconds != nil {
		in, out := &in.AccessTokenMaxAgeSeconds, &out.AccessTokenMaxAgeSeconds
		*out = new(int32)
		**out = **in
	}
	if in.AccessTokenInactivityTimeoutSeconds != nil {
		in, out := &in.AccessTokenInactivityTimeoutSeconds, &out.AccessTokenInactivityTimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	return
}
func (in *OAuthClient) DeepCopy() *OAuthClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OAuthClient)
	in.DeepCopyInto(out)
	return out
}
func (in *OAuthClient) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *OAuthClientAuthorization) DeepCopyInto(out *OAuthClientAuthorization) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *OAuthClientAuthorization) DeepCopy() *OAuthClientAuthorization {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OAuthClientAuthorization)
	in.DeepCopyInto(out)
	return out
}
func (in *OAuthClientAuthorization) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *OAuthClientAuthorizationList) DeepCopyInto(out *OAuthClientAuthorizationList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthClientAuthorization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *OAuthClientAuthorizationList) DeepCopy() *OAuthClientAuthorizationList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OAuthClientAuthorizationList)
	in.DeepCopyInto(out)
	return out
}
func (in *OAuthClientAuthorizationList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *OAuthClientList) DeepCopyInto(out *OAuthClientList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *OAuthClientList) DeepCopy() *OAuthClientList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OAuthClientList)
	in.DeepCopyInto(out)
	return out
}
func (in *OAuthClientList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *OAuthRedirectReference) DeepCopyInto(out *OAuthRedirectReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Reference = in.Reference
	return
}
func (in *OAuthRedirectReference) DeepCopy() *OAuthRedirectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OAuthRedirectReference)
	in.DeepCopyInto(out)
	return out
}
func (in *OAuthRedirectReference) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *RedirectReference) DeepCopyInto(out *RedirectReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *RedirectReference) DeepCopy() *RedirectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RedirectReference)
	in.DeepCopyInto(out)
	return out
}
func (in *ScopeRestriction) DeepCopyInto(out *ScopeRestriction) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.ExactValues != nil {
		in, out := &in.ExactValues, &out.ExactValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterRole != nil {
		in, out := &in.ClusterRole, &out.ClusterRole
		*out = new(ClusterRoleScopeRestriction)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ScopeRestriction) DeepCopy() *ScopeRestriction {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ScopeRestriction)
	in.DeepCopyInto(out)
	return out
}

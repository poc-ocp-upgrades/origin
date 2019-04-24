package v1

import (
	unsafe "unsafe"
	v1 "github.com/openshift/api/route/v1"
	route "github.com/openshift/origin/pkg/route/apis/route"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*v1.Route)(nil), (*route.Route)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Route_To_route_Route(a.(*v1.Route), b.(*route.Route), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*route.Route)(nil), (*v1.Route)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_route_Route_To_v1_Route(a.(*route.Route), b.(*v1.Route), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RouteIngress)(nil), (*route.RouteIngress)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RouteIngress_To_route_RouteIngress(a.(*v1.RouteIngress), b.(*route.RouteIngress), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*route.RouteIngress)(nil), (*v1.RouteIngress)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_route_RouteIngress_To_v1_RouteIngress(a.(*route.RouteIngress), b.(*v1.RouteIngress), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RouteIngressCondition)(nil), (*route.RouteIngressCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RouteIngressCondition_To_route_RouteIngressCondition(a.(*v1.RouteIngressCondition), b.(*route.RouteIngressCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*route.RouteIngressCondition)(nil), (*v1.RouteIngressCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_route_RouteIngressCondition_To_v1_RouteIngressCondition(a.(*route.RouteIngressCondition), b.(*v1.RouteIngressCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RouteList)(nil), (*route.RouteList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RouteList_To_route_RouteList(a.(*v1.RouteList), b.(*route.RouteList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*route.RouteList)(nil), (*v1.RouteList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_route_RouteList_To_v1_RouteList(a.(*route.RouteList), b.(*v1.RouteList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RoutePort)(nil), (*route.RoutePort)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoutePort_To_route_RoutePort(a.(*v1.RoutePort), b.(*route.RoutePort), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*route.RoutePort)(nil), (*v1.RoutePort)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_route_RoutePort_To_v1_RoutePort(a.(*route.RoutePort), b.(*v1.RoutePort), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RouteSpec)(nil), (*route.RouteSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RouteSpec_To_route_RouteSpec(a.(*v1.RouteSpec), b.(*route.RouteSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*route.RouteSpec)(nil), (*v1.RouteSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_route_RouteSpec_To_v1_RouteSpec(a.(*route.RouteSpec), b.(*v1.RouteSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RouteStatus)(nil), (*route.RouteStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RouteStatus_To_route_RouteStatus(a.(*v1.RouteStatus), b.(*route.RouteStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*route.RouteStatus)(nil), (*v1.RouteStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_route_RouteStatus_To_v1_RouteStatus(a.(*route.RouteStatus), b.(*v1.RouteStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RouteTargetReference)(nil), (*route.RouteTargetReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RouteTargetReference_To_route_RouteTargetReference(a.(*v1.RouteTargetReference), b.(*route.RouteTargetReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*route.RouteTargetReference)(nil), (*v1.RouteTargetReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_route_RouteTargetReference_To_v1_RouteTargetReference(a.(*route.RouteTargetReference), b.(*v1.RouteTargetReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RouterShard)(nil), (*route.RouterShard)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RouterShard_To_route_RouterShard(a.(*v1.RouterShard), b.(*route.RouterShard), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*route.RouterShard)(nil), (*v1.RouterShard)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_route_RouterShard_To_v1_RouterShard(a.(*route.RouterShard), b.(*v1.RouterShard), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TLSConfig)(nil), (*route.TLSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TLSConfig_To_route_TLSConfig(a.(*v1.TLSConfig), b.(*route.TLSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*route.TLSConfig)(nil), (*v1.TLSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_route_TLSConfig_To_v1_TLSConfig(a.(*route.TLSConfig), b.(*v1.TLSConfig), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_Route_To_route_Route(in *v1.Route, out *route.Route, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_RouteSpec_To_route_RouteSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_RouteStatus_To_route_RouteStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_Route_To_route_Route(in *v1.Route, out *route.Route, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_Route_To_route_Route(in, out, s)
}
func autoConvert_route_Route_To_v1_Route(in *route.Route, out *v1.Route, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_route_RouteSpec_To_v1_RouteSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_route_RouteStatus_To_v1_RouteStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_route_Route_To_v1_Route(in *route.Route, out *v1.Route, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_route_Route_To_v1_Route(in, out, s)
}
func autoConvert_v1_RouteIngress_To_route_RouteIngress(in *v1.RouteIngress, out *route.RouteIngress, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Host = in.Host
	out.RouterName = in.RouterName
	out.Conditions = *(*[]route.RouteIngressCondition)(unsafe.Pointer(&in.Conditions))
	out.WildcardPolicy = route.WildcardPolicyType(in.WildcardPolicy)
	out.RouterCanonicalHostname = in.RouterCanonicalHostname
	return nil
}
func Convert_v1_RouteIngress_To_route_RouteIngress(in *v1.RouteIngress, out *route.RouteIngress, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RouteIngress_To_route_RouteIngress(in, out, s)
}
func autoConvert_route_RouteIngress_To_v1_RouteIngress(in *route.RouteIngress, out *v1.RouteIngress, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Host = in.Host
	out.RouterName = in.RouterName
	out.Conditions = *(*[]v1.RouteIngressCondition)(unsafe.Pointer(&in.Conditions))
	out.WildcardPolicy = v1.WildcardPolicyType(in.WildcardPolicy)
	out.RouterCanonicalHostname = in.RouterCanonicalHostname
	return nil
}
func Convert_route_RouteIngress_To_v1_RouteIngress(in *route.RouteIngress, out *v1.RouteIngress, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_route_RouteIngress_To_v1_RouteIngress(in, out, s)
}
func autoConvert_v1_RouteIngressCondition_To_route_RouteIngressCondition(in *v1.RouteIngressCondition, out *route.RouteIngressCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = route.RouteIngressConditionType(in.Type)
	out.Status = core.ConditionStatus(in.Status)
	out.Reason = in.Reason
	out.Message = in.Message
	out.LastTransitionTime = (*metav1.Time)(unsafe.Pointer(in.LastTransitionTime))
	return nil
}
func Convert_v1_RouteIngressCondition_To_route_RouteIngressCondition(in *v1.RouteIngressCondition, out *route.RouteIngressCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RouteIngressCondition_To_route_RouteIngressCondition(in, out, s)
}
func autoConvert_route_RouteIngressCondition_To_v1_RouteIngressCondition(in *route.RouteIngressCondition, out *v1.RouteIngressCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.RouteIngressConditionType(in.Type)
	out.Status = corev1.ConditionStatus(in.Status)
	out.Reason = in.Reason
	out.Message = in.Message
	out.LastTransitionTime = (*metav1.Time)(unsafe.Pointer(in.LastTransitionTime))
	return nil
}
func Convert_route_RouteIngressCondition_To_v1_RouteIngressCondition(in *route.RouteIngressCondition, out *v1.RouteIngressCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_route_RouteIngressCondition_To_v1_RouteIngressCondition(in, out, s)
}
func autoConvert_v1_RouteList_To_route_RouteList(in *v1.RouteList, out *route.RouteList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]route.Route)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_RouteList_To_route_RouteList(in *v1.RouteList, out *route.RouteList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RouteList_To_route_RouteList(in, out, s)
}
func autoConvert_route_RouteList_To_v1_RouteList(in *route.RouteList, out *v1.RouteList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.Route)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_route_RouteList_To_v1_RouteList(in *route.RouteList, out *v1.RouteList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_route_RouteList_To_v1_RouteList(in, out, s)
}
func autoConvert_v1_RoutePort_To_route_RoutePort(in *v1.RoutePort, out *route.RoutePort, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.TargetPort = in.TargetPort
	return nil
}
func Convert_v1_RoutePort_To_route_RoutePort(in *v1.RoutePort, out *route.RoutePort, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RoutePort_To_route_RoutePort(in, out, s)
}
func autoConvert_route_RoutePort_To_v1_RoutePort(in *route.RoutePort, out *v1.RoutePort, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.TargetPort = in.TargetPort
	return nil
}
func Convert_route_RoutePort_To_v1_RoutePort(in *route.RoutePort, out *v1.RoutePort, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_route_RoutePort_To_v1_RoutePort(in, out, s)
}
func autoConvert_v1_RouteSpec_To_route_RouteSpec(in *v1.RouteSpec, out *route.RouteSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Host = in.Host
	out.Path = in.Path
	if err := Convert_v1_RouteTargetReference_To_route_RouteTargetReference(&in.To, &out.To, s); err != nil {
		return err
	}
	out.AlternateBackends = *(*[]route.RouteTargetReference)(unsafe.Pointer(&in.AlternateBackends))
	out.Port = (*route.RoutePort)(unsafe.Pointer(in.Port))
	out.TLS = (*route.TLSConfig)(unsafe.Pointer(in.TLS))
	out.WildcardPolicy = route.WildcardPolicyType(in.WildcardPolicy)
	return nil
}
func Convert_v1_RouteSpec_To_route_RouteSpec(in *v1.RouteSpec, out *route.RouteSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RouteSpec_To_route_RouteSpec(in, out, s)
}
func autoConvert_route_RouteSpec_To_v1_RouteSpec(in *route.RouteSpec, out *v1.RouteSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Host = in.Host
	out.Path = in.Path
	if err := Convert_route_RouteTargetReference_To_v1_RouteTargetReference(&in.To, &out.To, s); err != nil {
		return err
	}
	out.AlternateBackends = *(*[]v1.RouteTargetReference)(unsafe.Pointer(&in.AlternateBackends))
	out.Port = (*v1.RoutePort)(unsafe.Pointer(in.Port))
	out.TLS = (*v1.TLSConfig)(unsafe.Pointer(in.TLS))
	out.WildcardPolicy = v1.WildcardPolicyType(in.WildcardPolicy)
	return nil
}
func Convert_route_RouteSpec_To_v1_RouteSpec(in *route.RouteSpec, out *v1.RouteSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_route_RouteSpec_To_v1_RouteSpec(in, out, s)
}
func autoConvert_v1_RouteStatus_To_route_RouteStatus(in *v1.RouteStatus, out *route.RouteStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Ingress = *(*[]route.RouteIngress)(unsafe.Pointer(&in.Ingress))
	return nil
}
func Convert_v1_RouteStatus_To_route_RouteStatus(in *v1.RouteStatus, out *route.RouteStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RouteStatus_To_route_RouteStatus(in, out, s)
}
func autoConvert_route_RouteStatus_To_v1_RouteStatus(in *route.RouteStatus, out *v1.RouteStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Ingress = *(*[]v1.RouteIngress)(unsafe.Pointer(&in.Ingress))
	return nil
}
func Convert_route_RouteStatus_To_v1_RouteStatus(in *route.RouteStatus, out *v1.RouteStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_route_RouteStatus_To_v1_RouteStatus(in, out, s)
}
func autoConvert_v1_RouteTargetReference_To_route_RouteTargetReference(in *v1.RouteTargetReference, out *route.RouteTargetReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Kind = in.Kind
	out.Name = in.Name
	out.Weight = (*int32)(unsafe.Pointer(in.Weight))
	return nil
}
func Convert_v1_RouteTargetReference_To_route_RouteTargetReference(in *v1.RouteTargetReference, out *route.RouteTargetReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RouteTargetReference_To_route_RouteTargetReference(in, out, s)
}
func autoConvert_route_RouteTargetReference_To_v1_RouteTargetReference(in *route.RouteTargetReference, out *v1.RouteTargetReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Kind = in.Kind
	out.Name = in.Name
	out.Weight = (*int32)(unsafe.Pointer(in.Weight))
	return nil
}
func Convert_route_RouteTargetReference_To_v1_RouteTargetReference(in *route.RouteTargetReference, out *v1.RouteTargetReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_route_RouteTargetReference_To_v1_RouteTargetReference(in, out, s)
}
func autoConvert_v1_RouterShard_To_route_RouterShard(in *v1.RouterShard, out *route.RouterShard, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ShardName = in.ShardName
	out.DNSSuffix = in.DNSSuffix
	return nil
}
func Convert_v1_RouterShard_To_route_RouterShard(in *v1.RouterShard, out *route.RouterShard, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RouterShard_To_route_RouterShard(in, out, s)
}
func autoConvert_route_RouterShard_To_v1_RouterShard(in *route.RouterShard, out *v1.RouterShard, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ShardName = in.ShardName
	out.DNSSuffix = in.DNSSuffix
	return nil
}
func Convert_route_RouterShard_To_v1_RouterShard(in *route.RouterShard, out *v1.RouterShard, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_route_RouterShard_To_v1_RouterShard(in, out, s)
}
func autoConvert_v1_TLSConfig_To_route_TLSConfig(in *v1.TLSConfig, out *route.TLSConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Termination = route.TLSTerminationType(in.Termination)
	out.Certificate = in.Certificate
	out.Key = in.Key
	out.CACertificate = in.CACertificate
	out.DestinationCACertificate = in.DestinationCACertificate
	out.InsecureEdgeTerminationPolicy = route.InsecureEdgeTerminationPolicyType(in.InsecureEdgeTerminationPolicy)
	return nil
}
func Convert_v1_TLSConfig_To_route_TLSConfig(in *v1.TLSConfig, out *route.TLSConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TLSConfig_To_route_TLSConfig(in, out, s)
}
func autoConvert_route_TLSConfig_To_v1_TLSConfig(in *route.TLSConfig, out *v1.TLSConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Termination = v1.TLSTerminationType(in.Termination)
	out.Certificate = in.Certificate
	out.Key = in.Key
	out.CACertificate = in.CACertificate
	out.DestinationCACertificate = in.DestinationCACertificate
	out.InsecureEdgeTerminationPolicy = v1.InsecureEdgeTerminationPolicyType(in.InsecureEdgeTerminationPolicy)
	return nil
}
func Convert_route_TLSConfig_To_v1_TLSConfig(in *route.TLSConfig, out *v1.TLSConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_route_TLSConfig_To_v1_TLSConfig(in, out, s)
}

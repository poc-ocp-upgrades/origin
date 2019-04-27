package v1

import (
	v1 "github.com/openshift/api/route/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddTypeDefaultingFunc(&v1.Route{}, func(obj interface{}) {
		SetObjectDefaults_Route(obj.(*v1.Route))
	})
	scheme.AddTypeDefaultingFunc(&v1.RouteList{}, func(obj interface{}) {
		SetObjectDefaults_RouteList(obj.(*v1.RouteList))
	})
	return nil
}
func SetObjectDefaults_Route(in *v1.Route) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_RouteSpec(&in.Spec)
	SetDefaults_RouteTargetReference(&in.Spec.To)
	for i := range in.Spec.AlternateBackends {
		a := &in.Spec.AlternateBackends[i]
		SetDefaults_RouteTargetReference(a)
	}
	if in.Spec.TLS != nil {
		SetDefaults_TLSConfig(in.Spec.TLS)
	}
	for i := range in.Status.Ingress {
		a := &in.Status.Ingress[i]
		SetDefaults_RouteIngress(a)
	}
}
func SetObjectDefaults_RouteList(in *v1.RouteList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Route(a)
	}
}

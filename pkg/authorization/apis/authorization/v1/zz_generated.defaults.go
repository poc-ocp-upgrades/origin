package v1

import (
	v1 "github.com/openshift/api/authorization/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddTypeDefaultingFunc(&v1.ClusterRole{}, func(obj interface{}) {
		SetObjectDefaults_ClusterRole(obj.(*v1.ClusterRole))
	})
	scheme.AddTypeDefaultingFunc(&v1.ClusterRoleList{}, func(obj interface{}) {
		SetObjectDefaults_ClusterRoleList(obj.(*v1.ClusterRoleList))
	})
	scheme.AddTypeDefaultingFunc(&v1.Role{}, func(obj interface{}) {
		SetObjectDefaults_Role(obj.(*v1.Role))
	})
	scheme.AddTypeDefaultingFunc(&v1.RoleList{}, func(obj interface{}) {
		SetObjectDefaults_RoleList(obj.(*v1.RoleList))
	})
	scheme.AddTypeDefaultingFunc(&v1.SelfSubjectRulesReview{}, func(obj interface{}) {
		SetObjectDefaults_SelfSubjectRulesReview(obj.(*v1.SelfSubjectRulesReview))
	})
	scheme.AddTypeDefaultingFunc(&v1.SubjectRulesReview{}, func(obj interface{}) {
		SetObjectDefaults_SubjectRulesReview(obj.(*v1.SubjectRulesReview))
	})
	return nil
}
func SetObjectDefaults_ClusterRole(in *v1.ClusterRole) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Rules {
		a := &in.Rules[i]
		SetDefaults_PolicyRule(a)
	}
}
func SetObjectDefaults_ClusterRoleList(in *v1.ClusterRoleList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ClusterRole(a)
	}
}
func SetObjectDefaults_Role(in *v1.Role) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Rules {
		a := &in.Rules[i]
		SetDefaults_PolicyRule(a)
	}
}
func SetObjectDefaults_RoleList(in *v1.RoleList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Role(a)
	}
}
func SetObjectDefaults_SelfSubjectRulesReview(in *v1.SelfSubjectRulesReview) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Status.Rules {
		a := &in.Status.Rules[i]
		SetDefaults_PolicyRule(a)
	}
}
func SetObjectDefaults_SubjectRulesReview(in *v1.SubjectRulesReview) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Status.Rules {
		a := &in.Status.Rules[i]
		SetDefaults_PolicyRule(a)
	}
}

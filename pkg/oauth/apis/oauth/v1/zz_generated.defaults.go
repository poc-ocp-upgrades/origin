package v1

import (
	v1 "github.com/openshift/api/oauth/v1"
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
	scheme.AddTypeDefaultingFunc(&v1.OAuthAuthorizeToken{}, func(obj interface{}) {
		SetObjectDefaults_OAuthAuthorizeToken(obj.(*v1.OAuthAuthorizeToken))
	})
	scheme.AddTypeDefaultingFunc(&v1.OAuthAuthorizeTokenList{}, func(obj interface{}) {
		SetObjectDefaults_OAuthAuthorizeTokenList(obj.(*v1.OAuthAuthorizeTokenList))
	})
	return nil
}
func SetObjectDefaults_OAuthAuthorizeToken(in *v1.OAuthAuthorizeToken) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_OAuthAuthorizeToken(in)
}
func SetObjectDefaults_OAuthAuthorizeTokenList(in *v1.OAuthAuthorizeTokenList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
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
		SetObjectDefaults_OAuthAuthorizeToken(a)
	}
}

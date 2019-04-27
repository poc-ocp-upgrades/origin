package identitymapper

import (
	kuser "k8s.io/apiserver/pkg/authentication/user"
	userapi "github.com/openshift/api/user/v1"
)

func userToInfo(user *userapi.User) kuser.Info {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &kuser.DefaultInfo{Name: user.Name, UID: string(user.UID)}
}

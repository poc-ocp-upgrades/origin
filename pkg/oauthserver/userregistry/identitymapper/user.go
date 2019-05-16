package identitymapper

import (
	userapi "github.com/openshift/api/user/v1"
	kuser "k8s.io/apiserver/pkg/authentication/user"
)

func userToInfo(user *userapi.User) kuser.Info {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &kuser.DefaultInfo{Name: user.Name, UID: string(user.UID)}
}
